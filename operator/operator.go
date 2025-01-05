package operator

import (
	"context"
	"fmt"
	"slices"
	"sync/atomic"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zenrocklabs/zenrock-avs/aggregator"
	cstaskmanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/ZrTaskManager"
	"github.com/zenrocklabs/zenrock-avs/core"
	"github.com/zenrocklabs/zenrock-avs/core/chainio"
	"github.com/zenrocklabs/zenrock-avs/metrics"
	"github.com/zenrocklabs/zenrock-avs/types"

	"github.com/Zenrock-Foundation/zrchain/v5/go-client"

	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	sidecartypes "github.com/Zenrock-Foundation/zrchain/v5/sidecar/shared"
)

const AVS_NAME = "zenrock-avs"
const SEM_VER = "0.0.1"

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient eth.Client
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg       *prometheus.Registry
	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader sdkelcontracts.ELReader
	eigenlayerWriter sdkelcontracts.ELWriter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
	operatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *cstaskmanager.ContractZrTaskManagerNewTaskCreated
	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	credibleSquaringServiceManagerAddr common.Address
	// zenrock chain client
	zrChainClient   *client.QueryClient
	sidecarStatePtr *atomic.Value
}

// TODO(samlaf): config is a mess right now, since the chainio client constructors
//
//	take the config in core (which is shared with aggregator and challenger)
func NewOperatorFromConfig(c types.NodeConfig, sidecarStatePtr *atomic.Value) (*Operator, error) {
	var logLevel logging.LogLevel
	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}
	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}
	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient eth.Client
	if c.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewInstrumentedClient(c.EthWsUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	} else {
		ethRpcClient, err = eth.NewClient(c.EthRpcUrl)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewClient(c.EthWsUrl)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	}

	// blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	// if !ok {
	// 	logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	// }
	// blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	// if err != nil {
	// 	logger.Errorf("Cannot parse bls private key", "err", err)
	// 	return nil, err
	// }

	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	// chainId, err := ethRpcClient.ChainID(context.Background())
	// if err != nil {
	// 	logger.Error("Cannot get chainId", "err", err)
	// 	return nil, err
	// }

	// ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	// if !ok {
	// 	logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	// }

	// signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
	// 	KeystorePath: c.EcdsaPrivateKeyStorePath,
	// 	Password:     ecdsaKeyPassword,
	// }, chainId)
	// if err != nil {
	// 	panic(err)
	// }

	// chainioConfig := clients.BuildAllConfig{
	// 	EthHttpUrl:                 c.EthRpcUrl,
	// 	EthWsUrl:                   c.EthWsUrl,
	// 	RegistryCoordinatorAddr:    c.AVSRegistryCoordinatorAddress,
	// 	OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress,
	// 	AvsName:                    AVS_NAME,
	// 	PromMetricsIpPortAddress:   c.EigenMetricsIpPortAddress,
	// }

	// operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
	// 	c.EcdsaPrivateKeyStorePath,
	// 	ecdsaKeyPassword,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// sdkClients, err := clients.BuildAll(chainioConfig, operatorEcdsaPrivateKey, logger)
	// if err != nil {
	// 	panic(err)
	// }
	// skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	// if err != nil {
	// 	panic(err)
	// }
	// txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, common.HexToAddress(c.OperatorAddress))

	// avsWriter, err := chainio.BuildAvsWriter(
	// 	txMgr, common.HexToAddress(c.AVSRegistryCoordinatorAddress),
	// 	common.HexToAddress(c.OperatorStateRetrieverAddress), ethRpcClient, logger,
	// )
	// if err != nil {
	// 	logger.Error("Cannot create AvsWriter", "err", err)
	// 	return nil, err
	// }

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.BuildAvsSubscriber(common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress), ethWsClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	// quorumNames := map[sdktypes.QuorumNum]string{
	// 	0: "quorum0",
	// }
	// economicMetricsCollector := economic.NewCollector(
	// 	sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
	// 	AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	// reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	zrChainClient, err := client.NewQueryClient(c.ZRChainRPCAddress, true)
	if err != nil {
		return nil, fmt.Errorf("Refresh Address Client: failed to get new client: %w", err)
	}

	operator := &Operator{
		config:     c,
		logger:     logger,
		metricsReg: reg,
		metrics:    avsAndEigenMetrics,
		nodeApi:    nodeApi,
		ethClient:  ethRpcClient,
		// avsWriter:                          avsWriter,
		avsReader:     avsReader,
		avsSubscriber: avsSubscriber,
		// eigenlayerReader:                   sdkClients.ElChainReader,
		// eigenlayerWriter:                   sdkClients.ElChainWriter,
		// blsKeypair:                         blsKeyPair,
		operatorAddr:                       common.HexToAddress(c.OperatorAddress),
		aggregatorServerIpPortAddr:         c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:                aggregatorRpcClient,
		newTaskCreatedChan:                 make(chan *cstaskmanager.ContractZrTaskManagerNewTaskCreated),
		credibleSquaringServiceManagerAddr: common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		operatorId:                         [32]byte{0}, // this is set below
		zrChainClient:                      zrChainClient,
		sidecarStatePtr:                    sidecarStatePtr,
	}

	// operator.registerOperatorOnStartup(operatorEcdsaPrivateKey)

	// OperatorId is set in contract during registration so we get it after registering operator.
	// operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	// if err != nil {
	// 	logger.Error("Cannot get operator id", "err", err)
	// 	return nil, err
	// }
	// operator.operatorId = operatorId
	// logger.Info("Operator info",
	// 	"operatorId", operatorId,
	// 	"operatorAddr", c.OperatorAddress,
	// 	"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
	// 	"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	// )

	return operator, nil

}

func (o *Operator) Start(ctx context.Context) error {
	o.logger.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}
	var metricsErrChan <-chan error
	if o.config.EnableMetrics {
		metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
	sub := o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-metricsErrChan:
			// TODO(samlaf); we should also register the service as unhealthy in the node api
			// https://eigen.nethermind.io/docs/spec/api/
			o.logger.Fatal("Error in metrics server", "err", err)
		case err := <-sub.Err():
			o.logger.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			sub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			sub = o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
		case newTaskCreatedLog := <-o.newTaskCreatedChan:
			o.metrics.IncNumTasksReceived()
			taskResponse := o.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			signedTaskResponse, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)
		}
	}
}

// ProcessNewTaskCreatedLog takes a NewTaskCreatedLog struct as input and returns a TaskResponse struct.
// The TaskResponse struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractZrTaskManagerNewTaskCreated) *cstaskmanager.ZrServiceManagerLibTaskResponse {
	o.logger.Debug("Received new task", "task", newTaskCreatedLog)
	o.logger.Info("Received new task",
		"taskId", newTaskCreatedLog.Task.TaskId,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)

	var validatorAddresses []string
	pageReq := &query.PageRequest{}

	for {
		resp, err := o.zrChainClient.ValidationQueryClient.BondedValidators(context.Background(), pageReq)
		if err != nil {
			o.logger.Error("Error getting unbonded validators", "err", err)
		}

		for _, validator := range resp.Validators {
			validatorAddresses = append(validatorAddresses, validator.OperatorAddress)
		}

		if resp.Pagination == nil || len(resp.Pagination.NextKey) == 0 {
			break
		}

		pageReq.Key = resp.Pagination.NextKey
	}

	sidecarState := o.sidecarStatePtr.Load().(*sidecartypes.OracleState)
	delegatedValidators := make([]string, 0, len(sidecarState.EigenDelegations))
	for k := range sidecarState.EigenDelegations {
		delegatedValidators = append(delegatedValidators, k)
	}

	// Get subset of delegatedValidators that are not in validatorAddresses (bonded validators on zrChain)
	validatorsToRemove := []string{}
	for _, validator := range delegatedValidators {
		if !slices.Contains(validatorAddresses, validator) {
			validatorsToRemove = append(validatorsToRemove, validator)
		}
	}

	// Create the TaskResponse with the new structure
	taskResponse := &cstaskmanager.ZrServiceManagerLibTaskResponse{
		ReferenceTaskId:    newTaskCreatedLog.Task.TaskId,
		InactiveSetZRChain: validatorsToRemove,
	}

	return taskResponse
}

func (o *Operator) SignTaskResponse(taskResponse *cstaskmanager.ZrServiceManagerLibTaskResponse) (*aggregator.SignedTaskResponse, error) {
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}
	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)
	signedTaskResponse := &aggregator.SignedTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
