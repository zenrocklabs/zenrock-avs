package aggregator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/zenrocklabs/zenrock-avs/aggregator/types"
	"github.com/zenrocklabs/zenrock-avs/core"
	"github.com/zenrocklabs/zenrock-avs/core/chainio"
	"github.com/zenrocklabs/zenrock-avs/core/config"

	cstaskmanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/ZrTaskManager"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 10000
	blockTimeSeconds         = 12 * time.Second
	avsName                  = "zenrock"

	taskCadence               = 7 * 24 * time.Hour
	responseDelay             = 30 * time.Second
	quorumThresholdPercentage = 67
)

// Aggregator sends tasks (numbers to square) onchain, then listens for operator signed TaskResponses.
// It aggregates responses signatures, and if any of the TaskResponses reaches the QuorumThresholdPercentage for each quorum
// (currently we only use a single quorum of the ERC20Mock token), it sends the aggregated TaskResponse and signature onchain.
//
// The signature is checked in the BLSSignatureChecker.sol contract, which expects a
//
//	struct NonSignerStakesAndSignature {
//		uint32[] nonSignerQuorumBitmapIndices;
//		BN254.G1Point[] nonSignerPubkeys;
//		BN254.G1Point[] quorumApks;
//		BN254.G2Point apkG2;
//		BN254.G1Point sigma;
//		uint32[] quorumApkIndices;
//		uint32[] totalStakeIndices;
//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
//	}
//
// A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
// In order to verify this onchain, the Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
// updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
// after having received new delegated shares or having delegated shares removed by stakers queuing withdrawals. Each of these pushes to their respective datatype array a new entry.
//
// This is true for quorumBitmaps (represent the quorums each operator is opted into), quorumApks (apks per quorum), totalStakes (total stake per quorum), and nonSignerStakes (stake per quorum per operator).
// The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
// Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
// Thankfully, we have deployed a helper contract BLSOperatorStateRetriever.sol whose function getCheckSignaturesIndices() can be used to fetch the indices given a block number.
//
// The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
// apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
// quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
// nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
// Upon sending a task onchain (or receiving a NewTaskCreated Event if the tasks were sent by an external task generator), the aggregator can get the list of all operators opted into each quorum at that
// block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.
type Aggregator struct {
	logger           logging.Logger
	serverIpPortAddr string
	avsWriter        chainio.AvsWriterer
	avsReader        chainio.AvsReaderer
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	tasks                 map[types.TaskId]cstaskmanager.ZrServiceManagerLibTask
	tasksMu               sync.RWMutex
	taskResponses         map[types.TaskId]map[sdktypes.TaskResponseDigest]cstaskmanager.ZrServiceManagerLibTaskResponse
	taskResponsesMu       sync.RWMutex
	currentTaskId         uint32
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {

	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.IncredibleSquaringRegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
		PromMetricsIpPortAddress:   ":9090",
	}
	clients, err := clients.BuildAll(chainioConfig, c.EcdsaPrivateKey, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create sdk clients", "err", err)
		return nil, err
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(context.Background(), clients.AvsRegistryChainSubscriber, clients.AvsRegistryChainReader, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)

	return &Aggregator{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		avsWriter:             avsWriter,
		avsReader:             avsReader,
		blsAggregationService: blsAggregationService,
		tasks:                 make(map[types.TaskId]cstaskmanager.ZrServiceManagerLibTask),
		taskResponses:         make(map[types.TaskId]map[sdktypes.TaskResponseDigest]cstaskmanager.ZrServiceManagerLibTaskResponse),
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting aggregator.")
	agg.logger.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)

	// TODO(soubhik): refactor task generation/sending into a separate function that we can run as goroutine
	ticker := time.NewTicker(taskCadence)
	agg.logger.Infof("Aggregator set to send new task every %s...", formatDuration(taskCadence))
	defer ticker.Stop()

	// Initialize currentTaskId
	latestTaskNumber, err := agg.avsReader.GetLatestTaskNumber(ctx)
	if err != nil {
		agg.logger.Error("Failed to get latest task number", "err", err)
		return err
	}
	agg.currentTaskId = latestTaskNumber + 1

	agg.sendNewTask()
	agg.currentTaskId++

	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			agg.sendAggregatedResponseToContract(blsAggServiceResp)
		case <-ticker.C:
			if err := agg.sendNewTask(); err != nil {
				continue
			}
			agg.currentTaskId++
		}
	}
}

func (agg *Aggregator) sendAggregatedResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	// TODO: check if blsAggServiceResp contains an err
	if blsAggServiceResp.Err != nil {
		agg.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// panicking to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(blsAggServiceResp.Err)
	}
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []cstaskmanager.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	agg.logger.Infof("Threshold reached for task %d. Sending aggregated response onchain in %s...",
		blsAggServiceResp.TaskIndex,
		responseDelay,
	)
	agg.tasksMu.RLock()
	task := agg.tasks[blsAggServiceResp.TaskIndex]
	agg.tasksMu.RUnlock()
	agg.taskResponsesMu.RLock()
	taskResponse := agg.taskResponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	agg.taskResponsesMu.RUnlock()
	time.Sleep(responseDelay)
	_, err := agg.avsWriter.SendAggregatedResponse(context.Background(), task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		agg.logger.Error("Aggregator failed to respond to task", "err", err)
	}
}

// sendNewTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (agg *Aggregator) sendNewTask() error {
	agg.logger.Info("Aggregator sending new task", "taskId", agg.currentTaskId)

	// Send task to the task manager contract
	newTask, taskIndex, err := agg.avsWriter.SendNewTask(context.Background(), agg.currentTaskId, quorumThresholdPercentage, types.QUORUM_NUMBERS)
	if err != nil {
		agg.logger.Error("Aggregator failed to send new task", "err", err)
		return err
	}

	agg.tasksMu.Lock()
	agg.tasks[taskIndex] = newTask
	agg.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}

	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}

	agg.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, quorumNums, quorumThresholdPercentages, taskTimeToExpiry)
	return nil
}

// formatDuration formats time.Duration in the most appropriate units for human readability
func formatDuration(d time.Duration) string {
	switch {
	case d.Hours() >= 24:
		return fmt.Sprintf("%d days", int(d.Hours()/24))
	case d.Hours() >= 1:
		return fmt.Sprintf("%d hours", int(d.Hours()))
	case d.Minutes() >= 1:
		return fmt.Sprintf("%d minutes", int(d.Minutes()))
	default:
		return fmt.Sprintf("%d seconds", int(d.Seconds()))
	}
}
