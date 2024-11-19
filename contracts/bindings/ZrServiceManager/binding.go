// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractZrServiceManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IRewardsCoordinatorRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ZrServiceManagerLibTask is an auto generated low-level Go binding around an user-defined struct.
type ZrServiceManagerLibTask struct {
	TaskId                    uint32
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// ZrServiceManagerLibTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type ZrServiceManagerLibTaskResponse struct {
	ReferenceTaskId    uint32
	InactiveSetZRChain []string
}

// ZrServiceManagerLibTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type ZrServiceManagerLibTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// ZrServiceManagerLibValidator is an auto generated low-level Go binding around an user-defined struct.
type ZrServiceManagerLibValidator struct {
	ValidatorAddr string
	ValidatorHash [32]byte
	Operators     []common.Address
}

// ContractZrServiceManagerMetaData contains all meta data concerning the ContractZrServiceManager contract.
var ContractZrServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIZrRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"QUORUM_NUMBER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNDELEGATION_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"oprAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectValidators\",\"inputs\":[{\"name\":\"inactiveValidatorSet\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllOperatorsAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structZrServiceManagerLib.Validator[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEigenStake\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakedBalanceForValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Validator\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardsInitiator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsInitiator\",\"inputs\":[{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAssigned\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDeregistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorsEjected\",\"inputs\":[{\"name\":\"validatorHashes\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6101806040523480156200001257600080fd5b506040516200951c3803806200951c8339810160408190526200003591620002de565b8383838383838383896001600081905550806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200009f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000c5919062000346565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200011d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000143919062000346565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200019d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001c3919062000346565b6001600160a01b0390811660e0529485166101005250918316610120528216610140521661016052620001f562000203565b50505050505050506200036d565b606554610100900460ff1615620002705760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60655460ff9081161015620002c3576065805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114620002db57600080fd5b50565b60008060008060808587031215620002f557600080fd5b84516200030281620002c5565b60208601519094506200031581620002c5565b60408601519093506200032881620002c5565b60608601519092506200033b81620002c5565b939692955090935050565b6000602082840312156200035957600080fd5b81516200036681620002c5565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051619062620004ba60003960008181610e6a01528181610fc50152818161105c015281816148ee01528181614ac001528181614c430152614ce2015260008181610c9501528181610d2401528181610da4015281816138220152818161397401528181613c2901528181613ea001528181614918015281816149fb01528181614b9e01528181614e42015281816150c20152616d82015260008181615242015281816152fe01526153ea01526000818161050201528181613c7d01528181613f0501528181613f840152614e9f0152600081816107080152612fd901526000818161049f01528181611cf901526131bb0152600081816104de01528181613391015261355301526000818161052b015281816115e801528181612cc301528181612e3c015261307601526190626000f3fe608060405234801561001057600080fd5b50600436106102d65760003560e01c8063886f119511610182578063d5f20ff6116100e9578063f5c9899d116100a2578063fabc1cbc1161007c578063fabc1cbc1461078e578063fb558ed5146107a1578063fc299dee146107a9578063fce36c7d146107bc57600080fd5b8063f5c9899d14610745578063f63c5bab14610772578063f891899d1461077b57600080fd5b8063d5f20ff6146106a5578063d604902f146106c5578063da4dc49c146106f0578063df5cf72314610703578063e481af9d1461072a578063f2fde38b1461073257600080fd5b8063a364f4da1161013b578063a364f4da1461062b578063a98fb3551461063e578063b98d090814610651578063c63e3c501461065e578063caf73aa014610671578063cefdc1d41461068457600080fd5b8063886f1195146105c6578063889e2c6e146105d95780638c5178cf146105e15780638da5cb5b146105f45780639926ee7d146106055780639d3a0f2d1461061857600080fd5b80634f739f741161024157806368304835116101fa5780636efb4636116101d45780636efb46361461054d578063715018a61461056e5780637533f901146105765780637b654c5d1461058057600080fd5b806368304835146104d95780636b3aa72e146105005780636d14a9871461052657600080fd5b80634f739f741461040d578063595c6a671461042d5780635ac86ab7146104355780635c155662146104685780635c975abb146104885780635df459461461049a57600080fd5b80633563b0d1116102935780633563b0d11461038c5780633bc28c8c146103ac578063416c7e5e146103bf5780634a91a2f8146103d25780634c71c165146103e75780634d2b57fe146103fa57600080fd5b806310d67a2f146102db578063136439dd146102f0578063171f1d5b146103035780632dda9fc61461033257806331b36bd91461034c57806333cfb7b71461036c575b600080fd5b6102ee6102e936600461707b565b6107cf565b005b6102ee6102fe366004617098565b61088b565b610316610311366004617202565b6109ca565b6040805192151583529015156020830152015b60405180910390f35b61033a600081565b60405160ff9091168152602001610329565b61035f61035a366004617276565b610b54565b6040516103299190617364565b61037f61037a36600461707b565b610c70565b60405161032991906173b0565b61039f61039a36600461744f565b61113f565b6040516103299190617556565b6102ee6103ba36600461707b565b6115d5565b6102ee6103cd366004617577565b6115e6565b6103da61171d565b6040516103299190617625565b6102ee6103f5366004617720565b611823565b61037f610408366004617814565b611e42565b61042061041b3660046178ef565b611f57565b60405161032991906179b9565b6102ee61267f565b610458610443366004617a83565b600254600160ff9092169190911b9081161490565b6040519015158152602001610329565b61047b610476366004617aa0565b612746565b6040516103299190617ae7565b6002545b604051908152602001610329565b6104c17f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610329565b6104c17f000000000000000000000000000000000000000000000000000000000000000081565b7f00000000000000000000000000000000000000000000000000000000000000006104c1565b6104c17f000000000000000000000000000000000000000000000000000000000000000081565b61056061055b366004617d38565b61290e565b604051610329929190617df8565b6102ee613808565b61048c62024ea081565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0454640100000000900463ffffffff165b60405163ffffffff9091168152602001610329565b6001546104c1906001600160a01b031681565b61037f61381c565b6102ee6105ef366004617e41565b613a56565b6098546001600160a01b03166104c1565b6102ee610613366004617f3b565b613c1e565b6102ee610626366004617f80565b613cea565b6102ee61063936600461707b565b613e95565b6102ee61064c366004617fe4565b613f65565b6033546104589060ff1681565b6102ee61066c366004618093565b613fb9565b6102ee61067f3660046180db565b61404f565b610697610692366004618162565b614551565b604051610329929190618199565b6106b86106b3366004617098565b6146e3565b60405161032991906181b2565b6106d86106d3366004617fe4565b61470e565b6040516001600160601b039091168152602001610329565b6106d86106fe3660046181c5565b6148cc565b6104c17f000000000000000000000000000000000000000000000000000000000000000081565b61037f6149f5565b6102ee61074036600461707b565b614dc1565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f045463ffffffff166105b1565b6105b161271081565b6102ee6107893660046181fe565b614e37565b6102ee61079c366004617098565b614f43565b6102ee61509f565b60ca546104c1906001600160a01b031681565b6102ee6107ca366004618269565b6150f6565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610822573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061084691906182aa565b6001600160a01b0316336001600160a01b03161461087f5760405162461bcd60e51b8152600401610876906182c7565b60405180910390fd5b61088881615421565b50565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156108d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f79190618311565b6109135760405162461bcd60e51b81526004016108769061832e565b6002548181161461098c5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610876565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610a1257610a12618376565b60200201518951600160200201518a60200151600060028110610a3757610a37618376565b60200201518b60200151600160028110610a5357610a53618376565b602090810291909101518c518d830151604051610ab09a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610ad3919061838c565b9050610b46610aec610ae58884615518565b86906155a9565b610af461563e565b610b3c610b2d85610b27604080518082018252600080825260209182015281518083019092526001825260029082015290565b90615518565b610b368c6156fe565b906155a9565b886201d4c061578d565b909890975095505050505050565b606081516001600160401b03811115610b6f57610b6f6170b1565b604051908082528060200260200182016040528015610b98578160200160208202803683370190505b50905060005b8251811015610c6957836001600160a01b03166313542a4e848381518110610bc857610bc8618376565b60200260200101516040518263ffffffff1660e01b8152600401610bfb91906001600160a01b0391909116815260200190565b602060405180830381865afa158015610c18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3c91906183ae565b828281518110610c4e57610c4e618376565b6020908102919091010152610c62816183dd565b9050610b9e565b5092915050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610cdc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0091906183ae565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610d6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8f91906183f6565b90506001600160c01b0381161580610e2957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e24919061841f565b60ff16155b15610e4557505060408051600081526020810190915292915050565b6000610e59826001600160c01b03166159b1565b90506000805b8251811015610f2f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610ea957610ea9618376565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610eed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1191906183ae565b610f1b908361843c565b915080610f27816183dd565b915050610e5f565b506000816001600160401b03811115610f4a57610f4a6170b1565b604051908082528060200260200182016040528015610f73578160200160208202803683370190505b5090506000805b8451811015611132576000858281518110610f9757610f97618376565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa15801561100c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103091906183ae565b905060005b8181101561111c576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156110aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ce9190618464565b600001518686815181106110e4576110e4618376565b6001600160a01b039092166020928302919091019091015284611106816183dd565b9550508080611114906183dd565b915050611035565b505050808061112a906183dd565b915050610f7a565b5090979650505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611181573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a591906182aa565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120b91906182aa565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561124d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061127191906182aa565b9050600086516001600160401b0381111561128e5761128e6170b1565b6040519080825280602002602001820160405280156112c157816020015b60608152602001906001900390816112ac5790505b50905060005b87518110156115c95760008882815181106112e4576112e4618376565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015611345573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261136d91908101906184a5565b905080516001600160401b03811115611388576113886170b1565b6040519080825280602002602001820160405280156113d357816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816113a65790505b508484815181106113e6576113e6618376565b602002602001018190525060005b81518110156115b3576040518060600160405280876001600160a01b03166347b314e885858151811061142957611429618376565b60200260200101516040518263ffffffff1660e01b815260040161144f91815260200190565b602060405180830381865afa15801561146c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061149091906182aa565b6001600160a01b031681526020018383815181106114b0576114b0618376565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106114de576114de618376565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa15801561153a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061155e9190618535565b6001600160601b031681525085858151811061157c5761157c618376565b6020026020010151828151811061159557611595618376565b602002602001018190525080806115ab906183dd565b9150506113f4565b50505080806115c1906183dd565b9150506112c7565b50979650505050505050565b6115dd615a73565b61088881615acd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611644573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061166891906182aa565b6001600160a01b0316336001600160a01b0316146117145760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610876565b61088881615b36565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0354606090600080516020618fed833981519152906000906001600160401b0381111561176c5761176c6170b1565b6040519080825280602002602001820160405280156117b957816020015b6040805160608082018352808252600060208301529181019190915281526020019060019003908161178a5790505b50905060005b6003830154811015610c69576117f38360030182815481106117e3576117e3618376565b9060005260206000200154615b8f565b82828151811061180557611805618376565b6020026020010181905250808061181b906183dd565b9150506117bf565b60006118326020850185618552565b63ffffffff811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f066020526040902054909150600080516020618fed833981519152906118d05760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610876565b84846040516020016118e392919061869e565b60408051601f19818403018152918152815160209283012063ffffffff8516600090815260068501909352912054146119845760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610876565b63ffffffff8216600090815260078201602052604090205460ff1615611a1e5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610876565b612710611a2e6020860186618552565b611a3891906186dc565b63ffffffff164363ffffffff161115611ab95760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610876565b611ac66020860186618552565b63ffffffff16611ad96020880188618552565b63ffffffff1614611b1f5760405162461bcd60e51b815260206004820152601060248201526f0a8c2e6d640928840dad2e6dac2e8c6d60831b6044820152606401610876565b600083516001600160401b03811115611b3a57611b3a6170b1565b604051908082528060200260200182016040528015611b63578160200160208202803683370190505b50905060005b8451811015611bd557611ba6858281518110611b8757611b87618376565b6020026020010151805160009081526020918201519091526040902090565b828281518110611bb857611bb8618376565b602090810291909101015280611bcd816183dd565b915050611b69565b506000611be86040890160208a01618552565b82604051602001611bfa9291906186f9565b60405160208183030381529060405280519060200120905085602001358114611ca45760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610876565b600085516001600160401b03811115611cbf57611cbf6170b1565b604051908082528060200260200182016040528015611ce8578160200160208202803683370190505b50905060005b8651811015611ddb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110611d3857611d38618376565b60200260200101516040518263ffffffff1660e01b8152600401611d5e91815260200190565b602060405180830381865afa158015611d7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d9f91906182aa565b828281518110611db157611db1618376565b6001600160a01b039092166020928302919091019091015280611dd3816183dd565b915050611cee565b5063ffffffff85166000818152600786016020908152604091829020805460ff19166001179055815192835233908301527fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec910160405180910390a1505050505050505050565b606081516001600160401b03811115611e5d57611e5d6170b1565b604051908082528060200260200182016040528015611e86578160200160208202803683370190505b50905060005b8251811015610c6957836001600160a01b031663296bb064848381518110611eb657611eb6618376565b60200260200101516040518263ffffffff1660e01b8152600401611edc91815260200190565b602060405180830381865afa158015611ef9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f1d91906182aa565b828281518110611f2f57611f2f618376565b6001600160a01b0390921660209283029190910190910152611f50816183dd565b9050611e8c565b611f826040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fe691906182aa565b90506120136040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90612043908b9089908990600401618734565b600060405180830381865afa158015612060573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612088919081019061877b565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906120ba908b908b908b90600401618809565b600060405180830381865afa1580156120d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120ff919081019061877b565b6040820152856001600160401b0381111561211c5761211c6170b1565b60405190808252806020026020018201604052801561214f57816020015b606081526020019060019003908161213a5790505b50606082015260005b60ff8116871115612590576000856001600160401b0381111561217d5761217d6170b1565b6040519080825280602002602001820160405280156121a6578160200160208202803683370190505b5083606001518360ff16815181106121c0576121c0618376565b602002602001018190525060005b868110156124905760008c6001600160a01b03166304ec63518a8a858181106121f9576121f9618376565b905060200201358e8860000151868151811061221757612217618376565b60200260200101516040518463ffffffff1660e01b81526004016122549392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612271573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061229591906183f6565b9050806001600160c01b031660000361233c5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610876565b8a8a8560ff1681811061235157612351618376565b60016001600160c01b038516919093013560f81c1c8216909103905061247d57856001600160a01b031663dd9846b98a8a8581811061239257612392618376565b905060200201358d8d8860ff168181106123ae576123ae618376565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015612404573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124289190618829565b85606001518560ff168151811061244157612441618376565b6020026020010151848151811061245a5761245a618376565b63ffffffff9092166020928302919091019091015282612479816183dd565b9350505b5080612488816183dd565b9150506121ce565b506000816001600160401b038111156124ab576124ab6170b1565b6040519080825280602002602001820160405280156124d4578160200160208202803683370190505b50905060005b828110156125555784606001518460ff16815181106124fb576124fb618376565b6020026020010151818151811061251457612514618376565b602002602001015182828151811061252e5761252e618376565b63ffffffff909216602092830291909101909101528061254d816183dd565b9150506124da565b508084606001518460ff168151811061257057612570618376565b60200260200101819052505050808061258890618846565b915050612158565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125f591906182aa565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90612628908b908b908e90600401618865565b600060405180830381865afa158015612645573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261266d919081019061877b565b60208301525098975050505050505050565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156126c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126eb9190618311565b6127075760405162461bcd60e51b81526004016108769061832e565b600019600281905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161277892919061888f565b600060405180830381865afa158015612795573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526127bd919081019061877b565b9050600084516001600160401b038111156127da576127da6170b1565b604051908082528060200260200182016040528015612803578160200160208202803683370190505b50905060005b855181101561290457866001600160a01b03166304ec635187838151811061283357612833618376565b60200260200101518786858151811061284e5761284e618376565b60200260200101516040518463ffffffff1660e01b815260040161288b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156128a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128cc91906183f6565b6001600160c01b03168282815181106128e7576128e7618376565b6020908102919091010152806128fc816183dd565b915050612809565b5095945050505050565b604080518082019091526060808252602082015260008481036129875760405162461bcd60e51b8152602060048201526037602482015260008051602061900d83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610876565b6040830151518514801561299f575060a08301515185145b80156129af575060c08301515185145b80156129bf575060e08301515185145b612a295760405162461bcd60e51b8152602060048201526041602482015260008051602061900d83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610876565b82515160208401515114612aa15760405162461bcd60e51b81526020600482015260446024820181905260008051602061900d833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610876565b4363ffffffff168463ffffffff1610612b105760405162461bcd60e51b815260206004820152603c602482015260008051602061900d83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610876565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115612b5157612b516170b1565b604051908082528060200260200182016040528015612b7a578160200160208202803683370190505b506020820152866001600160401b03811115612b9857612b986170b1565b604051908082528060200260200182016040528015612bc1578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115612bf557612bf56170b1565b604051908082528060200260200182016040528015612c1e578160200160208202803683370190505b5081526020860151516001600160401b03811115612c3e57612c3e6170b1565b604051908082528060200260200182016040528015612c67578160200160208202803683370190505b5081602001819052506000612d398a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612d10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d34919061841f565b615cd8565b905060005b876020015151811015612fb557612d6488602001518281518110611b8757611b87618376565b83602001518281518110612d7a57612d7a618376565b60209081029190910101528015612e3a576020830151612d9b6001836188ae565b81518110612dab57612dab618376565b602002602001015160001c83602001518281518110612dcc57612dcc618376565b602002602001015160001c11612e3a576040805162461bcd60e51b815260206004820152602481019190915260008051602061900d83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610876565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612e7f57612e7f618376565b60200260200101518b8b600001518581518110612e9e57612e9e618376565b60200260200101516040518463ffffffff1660e01b8152600401612edb9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612ef8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f1c91906183f6565b6001600160c01b031683600001518281518110612f3b57612f3b618376565b602002602001018181525050612fa1610ae5612f758486600001518581518110612f6757612f67618376565b602002602001015116615d62565b8a602001518481518110612f8b57612f8b618376565b6020026020010151615d8d90919063ffffffff16565b945080612fad816183dd565b915050612d3e565b5050612fc083615e70565b60335490935060ff16600081612fd7576000613059565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015613035573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061305991906183ae565b905060005b8a8110156136d75782156131b9578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106130b5576130b5618376565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156130f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061311991906183ae565b613123919061843c565b116131b95760405162461bcd60e51b8152602060048201526066602482015260008051602061900d83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610876565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106131fa576131fa618376565b9050013560f81c60f81b60f81c8c8c60a00151858151811061321e5761321e618376565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561327a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061329e91906188c1565b6001600160401b0319166132c18a604001518381518110611b8757611b87618376565b67ffffffffffffffff19161461335d5760405162461bcd60e51b8152602060048201526061602482015260008051602061900d83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610876565b61338d8960400151828151811061337657613376618376565b6020026020010151876155a990919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106133d0576133d0618376565b9050013560f81c60f81b60f81c8c8c60c0015185815181106133f4576133f4618376565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015613450573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134749190618535565b8560200151828151811061348a5761348a618376565b6001600160601b039092166020928302919091018201528501518051829081106134b6576134b6618376565b6020026020010151856000015182815181106134d4576134d4618376565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156136c25761354c8660000151828151811061351e5761351e618376565b60200260200101518f8f8681811061353857613538618376565b600192013560f81c9290921c811614919050565b156136b0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061359257613592618376565b9050013560f81c60f81b60f81c8e896020015185815181106135b6576135b6618376565b60200260200101518f60e0015188815181106135d4576135d4618376565b602002602001015187815181106135ed576135ed618376565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015613651573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136759190618535565b875180518590811061368957613689618376565b6020026020010181815161369d91906188ec565b6001600160601b03169052506001909101905b806136ba816183dd565b9150506134f8565b505080806136cf906183dd565b91505061305e565b5050506000806136f18c868a606001518b608001516109ca565b91509150816137625760405162461bcd60e51b8152602060048201526043602482015260008051602061900d83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610876565b806137c35760405162461bcd60e51b8152602060048201526039602482015260008051602061900d83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610876565b505060008782602001516040516020016137de9291906186f9565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b613810615a73565b61381a6000615f0b565b565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa15801561387e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138a291906182aa565b604051638902624560e01b81526000600482015263ffffffff431660248201526001600160a01b039190911690638902624590604401600060405180830381865afa1580156138f5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261391d91908101906184a5565b9050600081516001600160401b0381111561393a5761393a6170b1565b604051908082528060200260200182016040528015613963578160200160208202803683370190505b50905060005b8251811015610c69577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663296bb0648483815181106139b3576139b3618376565b60200260200101516040518263ffffffff1660e01b81526004016139d991815260200190565b602060405180830381865afa1580156139f6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a1a91906182aa565b828281518110613a2c57613a2c618376565b6001600160a01b039092166020928302919091019091015280613a4e816183dd565b915050613969565b606554610100900460ff1615808015613a765750606554600160ff909116105b80613a905750303b158015613a90575060655460ff166001145b613af35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610876565b6065805460ff191660011790558015613b16576065805461ff0019166101001790555b613b21876000615f5d565b613b2b8684616043565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0880546001600160a01b03199081166001600160a01b03888116919091179092557fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0980549091169186169190911790557fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f04805463ffffffff191663ffffffff84161790558015613c15576065805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614613c665760405162461bcd60e51b81526004016108769061890c565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90613cb49085908590600401618984565b600060405180830381600087803b158015613cce57600080fd5b505af1158015613ce2573d6000803e3d6000fd5b505050505050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0954600080516020618fed833981519152906001600160a01b03163314613d7d5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610876565b60408051608081018252606081830181905263ffffffff8881168352438116602080850191909152908816918301919091528251601f860182900482028101820190935284835290919085908590819084018382808284376000920182905250604086019490945250613df29150615b7d9050565b905081604051602001613e059190618a1a565b60408051601f19818403018152828252805160209182012063ffffffff8b16600090815260058601909252919020557ff31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca590613e639089908590618a2d565b60405180910390a160048101805467ffffffff00000000191664010000000063ffffffff8a1602179055613c1561509f565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614613edd5760405162461bcd60e51b81526004016108769061890c565b613ee681616074565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015613f4a57600080fd5b505af1158015613f5e573d6000803e3d6000fd5b5050505050565b613f6d615a73565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590613f30908490600401618a4c565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0854600080516020618fed833981519152906001600160a01b031633146140425760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610876565b61404b8261634a565b5050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0854600080516020618fed833981519152906001600160a01b031633146140d85760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610876565b60006140ea6040860160208701618552565b90503660006140fc6040880188618a5f565b909250905060006141136080890160608a01618552565b9050600080516020618fed8339815191527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f05600061415460208c018c618552565b63ffffffff1663ffffffff16815260200190815260200160002054896040516020016141809190618aa5565b60405160208183030381529060405280519060200120146142095760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610876565b6000600682018161421d60208d018d618552565b63ffffffff1663ffffffff168152602001908152602001600020541461429a5760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610876565b60048101546142af9063ffffffff16866186dc565b63ffffffff164363ffffffff1611156143205760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610876565b6000886040516020016143339190618b21565b60405160208183030381529060405280519060200120905060008061435b8388888b8e61290e565b9150915060005b8681101561445a578560ff168360200151828151811061438457614384618376565b60200260200101516143969190618b34565b6001600160601b03166043846000015183815181106143b7576143b7618376565b60200260200101516001600160601b03166143d29190618b57565b1015614448576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610876565b80614452816183dd565b915050614362565b50600061446a60208d018d618b6e565b905011156144905761449061448260208d018d618b6e565b61448b91618bb7565b61634a565b60408051808201825263ffffffff431681526020808201849052915190916144bc918e91849101618bc4565b604051602081830303815290604052805190602001208560060160008f60000160208101906144eb9190618552565b63ffffffff1663ffffffff168152602001908152602001600020819055507fb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce35678c8260405161453a929190618bc4565b60405180910390a150505050505050505050505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061458c5761458c618376565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906145c8908890869060040161888f565b600060405180830381865afa1580156145e5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261460d919081019061877b565b60008151811061461f5761461f618376565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561468b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146af91906183f6565b6001600160c01b0316905060006146c5826159b1565b9050816146d38a838a61113f565b9550955050505050935093915050565b6040805160608082018352808252600060208301529181019190915261470882615b8f565b92915050565b600080826040516020016147229190618bf7565b6040516020818303038152906040528051906020012090506000614751600080516020618fed83398151915290565b9050600081600001600084815260200190815260200160002060405180606001604052908160008201805461478590618c13565b80601f01602080910402602001604051908101604052809291908181526020018280546147b190618c13565b80156147fe5780601f106147d3576101008083540402835291602001916147fe565b820191906000526020600020905b8154815290600101906020018083116147e157829003601f168201915b50505050508152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561486a57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161484c575b50505050508152505090506000805b826040015151811015612904576148ae8360400151828151811061489f5761489f618376565b602002602001015160006148cc565b6148b89083618c47565b9150806148c4816183dd565b915050614879565b604051631619718360e21b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000821691635401ed27917f000000000000000000000000000000000000000000000000000000000000000090911690635865c60c906024016040805180830381865afa158015614960573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149849190618c67565b5160405160e083901b6001600160e01b0319168152600481019190915260ff85166024820152604401602060405180830381865afa1580156149ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149ee9190618535565b9392505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614a57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a7b919061841f565b60ff16905080600003614a9c57505060408051600081526020810190915290565b6000805b82811015614b5157604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015614b0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b3391906183ae565b614b3d908361843c565b915080614b49816183dd565b915050614aa0565b506000816001600160401b03811115614b6c57614b6c6170b1565b604051908082528060200260200182016040528015614b95578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614bfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c1e919061841f565b60ff16811015614db757604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015614c92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614cb691906183ae565b905060005b81811015614da2576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015614d30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614d549190618464565b60000151858581518110614d6a57614d6a618376565b6001600160a01b039092166020928302919091019091015283614d8c816183dd565b9450508080614d9a906183dd565b915050614cbb565b50508080614daf906183dd565b915050614b9c565b5090949350505050565b614dc9615a73565b6001600160a01b038116614e2e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610876565b61088881615f0b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614614e7f5760405162461bcd60e51b81526004016108769061890c565b614e8882616458565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90614ed69086908590600401618984565b600060405180830381600087803b158015614ef057600080fd5b505af1158015614f04573d6000803e3d6000fd5b50505050600082604051602001614f1b9190618bf7565b604051602081830303815290604052805190602001209050614f3d84826164a3565b50505050565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614f96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614fba91906182aa565b6001600160a01b0316336001600160a01b031614614fea5760405162461bcd60e51b8152600401610876906182c7565b6002541981196002541916146150685760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610876565b600281905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016109bf565b60006150a961381c565b60405162cf2ab560e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062cf2ab590613f309084906004016173b0565b6150fe6166d0565b60005b818110156153d25782828281811061511b5761511b618376565b905060200281019061512d9190618c97565b61513e90604081019060200161707b565b6001600160a01b03166323b872dd333086868681811061516057615160618376565b90506020028101906151729190618c97565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156151c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151ed9190618311565b50600083838381811061520257615202618376565b90506020028101906152149190618c97565b61522590604081019060200161707b565b604051636eb1769f60e11b81523060048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166024830152919091169063dd62ed3e90604401602060405180830381865afa158015615293573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906152b791906183ae565b90508383838181106152cb576152cb618376565b90506020028101906152dd9190618c97565b6152ee90604081019060200161707b565b6001600160a01b031663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008387878781811061533057615330618376565b90506020028101906153429190618c97565b60400135615350919061843c565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561539b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906153bf9190618311565b5050806153cb906183dd565b9050615101565b5060405163fce36c7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063fce36c7d90613cb49085908590600401618d13565b6001600160a01b0381166154af5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610876565b600154604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152615534616eba565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808061556357fe5b50806155a15760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610876565b505092915050565b60408051808201909152600080825260208201526155c5616ed8565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808061560057fe5b50806155a15760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610876565b615646616ef6565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061572e600080516020618fcd8339815191528661838c565b90505b61573a81616765565b9093509150600080516020618fcd8339815191528283098303615773576040805180820190915290815260208101919091529392505050565b600080516020618fcd833981519152600182089050615731565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906157bf616f1b565b60005b60028110156159845760006157d8826006618b57565b90508482600281106157ec576157ec618376565b602002015151836157fe83600061843c565b600c811061580e5761580e618376565b602002015284826002811061582557615825618376565b6020020151602001518382600161583c919061843c565b600c811061584c5761584c618376565b602002015283826002811061586357615863618376565b602002015151518361587683600261843c565b600c811061588657615886618376565b602002015283826002811061589d5761589d618376565b60200201515160016020020151836158b683600361843c565b600c81106158c6576158c6618376565b60200201528382600281106158dd576158dd618376565b6020020151602001516000600281106158f8576158f8618376565b60200201518361590983600461843c565b600c811061591957615919618376565b602002015283826002811061593057615930618376565b60200201516020015160016002811061594b5761594b618376565b60200201518361595c83600561843c565b600c811061596c5761596c618376565b6020020152508061597c816183dd565b9150506157c2565b5061598d616f3a565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60606000806159bf84615d62565b61ffff166001600160401b038111156159da576159da6170b1565b6040519080825280601f01601f191660200182016040528015615a04576020820181803683370190505b5090506000805b825182108015615a1c575061010081105b15614db7576001811b935085841615615a63578060f81b838381518110615a4557615a45618376565b60200101906001600160f81b031916908160001a9053508160010191505b615a6c816183dd565b9050615a0b565b6098546001600160a01b0316331461381a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610876565b60ca54604080516001600160a01b03928316815291831660208301527fe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3910160405180910390a160ca80546001600160a01b0319166001600160a01b0392909216919091179055565b6033805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b600080516020618fed83398151915290565b604080516060808201835280825260006020808401829052838501839052858252600080516020618fed8339815191529081905290849020845192830190945283549293909282908290615be290618c13565b80601f0160208091040260200160405190810160405280929190818152602001828054615c0e90618c13565b8015615c5b5780601f10615c3057610100808354040283529160200191615c5b565b820191906000526020600020905b815481529060010190602001808311615c3e57829003601f168201915b505050505081526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015615cc757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311615ca9575b505050505081525050915050919050565b600080615ce4846167e7565b9050808360ff166001901b116149ee5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610876565b6000805b821561470857615d776001846188ae565b9092169180615d8581618e21565b915050615d66565b60408051808201909152600080825260208201526102008261ffff1610615de95760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610876565b8161ffff16600103615dfc575081614708565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610615e6557600161ffff871660ff83161c81169003615e4857615e4584846155a9565b93505b615e5283846155a9565b92506201fffe600192831b169101615e18565b509195945050505050565b60408051808201909152600080825260208201528151158015615e9557506020820151155b15615eb3575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020618fcd8339815191528460200151615ee6919061838c565b615efe90600080516020618fcd8339815191526188ae565b905292915050565b919050565b609880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001546001600160a01b0316158015615f7e57506001600160a01b03821615155b6160005760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610876565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261404b82615421565b606554610100900460ff1661606a5760405162461bcd60e51b815260040161087690618e42565b61404b8282616977565b6001600160a01b0381166160c25760405162461bcd60e51b81526020600482015260156024820152747a65726f206f70657261746f72206164647265737360581b6044820152606401610876565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f016020526040902054600080516020618fed83398151915290806161505760405162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840ecc2d8d2c8c2e8dee440d0c2e6d60531b6044820152606401610876565b6001600160a01b038316600090815260028084016020908152604080842054858552918690529092209081015482106161c15760405162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f6620626f756e647360681b6044820152606401610876565b846001600160a01b03168160020183815481106161e0576161e0618376565b6000918252602090912001546001600160a01b0316146162425760405162461bcd60e51b815260206004820152601960248201527f4f70657261746f722061646472657373206d69736d61746368000000000000006044820152606401610876565b6002810154600090616256906001906188ae565b60028301805491925060009161626e906001906188ae565b8154811061627e5761627e618376565b6000918252602090912001546001600160a01b031690508382146162e157808360020185815481106162b2576162b2618376565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b826002018054806162f4576162f4618e8d565b600082815260208120820160001990810180546001600160a01b031916905590910190915560405186917fec3ecc6d7c19c6fc03c8b840fabed1a1b70c73f74bdb2d71f18c2f5a23639d3291a250505050505050565b600081516001600160401b03811115616365576163656170b1565b60405190808252806020026020018201604052801561638e578160200160208202803683370190505b50905060005b825181101561641c5760008382815181106163b1576163b1618376565b60200260200101516040516020016163c99190618bf7565b604051602081830303815290604052805190602001209050808383815181106163f4576163f4618376565b602002602001018181525050616409816169b0565b5080616414816183dd565b915050616394565b507fb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a8160405161644c9190617364565b60405180910390a15050565b60008160405160200161646b9190618bf7565b60405160208183030381529060405280519060200120905061648c81616ab0565b151560000361404b5761649e82616ac5565b505050565b6001600160a01b0382166164f15760405162461bcd60e51b81526020600482015260156024820152747a65726f206f70657261746f72206164647265737360581b6044820152606401610876565b8061653e5760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642076616c696461746f722061646472657373000000000000006044820152606401610876565b61654781616ab0565b61658b5760405162461bcd60e51b81526020600482015260156024820152740eadcd6deeedc40ecc2d8d2c8c2e8dee440d0c2e6d605b1b6044820152606401610876565b60006165988360006148cc565b90506000816001600160601b0316116165e35760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964207374616b6560981b6044820152606401610876565b6001600160a01b03831660008181527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0160209081526040808320869055858352600080516020618fed83398151915280835281842060020180548686527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0285528386208190558285526001810182559085529383902090930180546001600160a01b0319168517905580516001600160601b038616815290519293869390927f5ea87e0bfaffd2b7f2901f087737e2b175f237068c6123ca9ea5240076bb5668928290030190a350505050565b60ca546001600160a01b0316331461381a5760405162461bcd60e51b815260206004820152604c60248201527f536572766963654d616e61676572426173652e6f6e6c7952657761726473496e60448201527f69746961746f723a2063616c6c6572206973206e6f742074686520726577617260648201526b32399034b734ba34b0ba37b960a11b608482015260a401610876565b60008080600080516020618fcd8339815191526003600080516020618fcd83398151915286600080516020618fcd8339815191528889090908905060006167db827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020618fcd833981519152616c3e565b91959194509092505050565b6000610100825111156168705760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610876565b815160000361688157506000919050565b6000808360008151811061689757616897618376565b0160200151600160f89190911c81901b92505b845181101561696e578481815181106168c5576168c5618376565b0160200151600160f89190911c1b915082821161695a5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610876565b91811791616967816183dd565b90506168aa565b50909392505050565b606554610100900460ff1661699e5760405162461bcd60e51b815260040161087690618e42565b6169a782615f0b565b61404b81615acd565b6000818152600080516020618fed8339815191526020819052604090912060018101541561649e5760005b6002820154811015616a7e57616a198260020182815481106169ff576169ff618376565b6000918252602090912001546001600160a01b0316616ce7565b83826002018281548110616a2f57616a2f618376565b60009182526020822001546040516001600160a01b03909116917f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a380616a76816183dd565b9150506169db565b50600083815260208390526040812090616a988282616f58565b6001820160009055600282016000613f5e9190616f92565b6000616abb82615b8f565b5151151592915050565b60408051606080820183528082526000602083015291810191909152600082604051602001616af49190618bf7565b604051602081830303815290604052805190602001209050616b1581616ab0565b15616b625760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f7220616c72656164792072656769737465726564000000006044820152606401610876565b604080516060810182528481526020808201849052825160008082529181018452909282015290506000600080516020618fed8339815191526000848152602082905260409020835191925083918190616bbc9082618ee9565b50602082810151600183015560408301518051616bdf9260028501920190616fb0565b505050600381018054600181018255600091825260209091200183905560405183907f1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea390616c2e908890618a4c565b60405180910390a2509392505050565b600080616c49616f3a565b616c51617015565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280616c8e57fe5b5082616cdc5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610876565b505195945050505050565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f016020526040902054600080516020618fed83398151915290801561649e5760408051600180825281830190925260009160208083019080368337019050509050600081600081518110616d6c57616d6c618376565b602002602001019060ff16908160ff16815250507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e3b17db85616db984616e0e565b6040518363ffffffff1660e01b8152600401616dd6929190618fa8565b600060405180830381600087803b158015616df057600080fd5b505af1158015616e04573d6000803e3d6000fd5b5050505050505050565b6060600082516001600160401b03811115616e2b57616e2b6170b1565b6040519080825280601f01601f191660200182016040528015616e55576020820181803683370190505b50905060005b8351811015610c6957838181518110616e7657616e76618376565b602002602001015160f81b828281518110616e9357616e93618376565b60200101906001600160f81b031916908160001a905350616eb3816183dd565b9050616e5b565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280616f09617033565b8152602001616f16617033565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b508054616f6490618c13565b6000825580601f10616f74575050565b601f0160209004906000526020600020908101906108889190617051565b50805460008255906000526020600020908101906108889190617051565b828054828255906000526020600020908101928215617005579160200282015b8281111561700557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190616fd0565b50617011929150617051565b5090565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b5b808211156170115760008155600101617052565b6001600160a01b038116811461088857600080fd5b60006020828403121561708d57600080fd5b81356149ee81617066565b6000602082840312156170aa57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156170e9576170e96170b1565b60405290565b60405161010081016001600160401b03811182821017156170e9576170e96170b1565b604051601f8201601f191681016001600160401b038111828210171561713a5761713a6170b1565b604052919050565b60006040828403121561715457600080fd5b61715c6170c7565b9050813581526020820135602082015292915050565b600082601f83011261718357600080fd5b61718b6170c7565b80604084018581111561719d57600080fd5b845b818110156171b757803584526020938401930161719f565b509095945050505050565b6000608082840312156171d457600080fd5b6171dc6170c7565b90506171e88383617172565b81526171f78360408401617172565b602082015292915050565b600080600080610120858703121561721957600080fd5b8435935061722a8660208701617142565b925061723986606087016171c2565b91506172488660e08701617142565b905092959194509250565b60006001600160401b0382111561726c5761726c6170b1565b5060051b60200190565b6000806040838503121561728957600080fd5b823561729481617066565b91506020838101356001600160401b038111156172b057600080fd5b8401601f810186136172c157600080fd5b80356172d46172cf82617253565b617112565b81815260059190911b820183019083810190888311156172f357600080fd5b928401925b8284101561731a57833561730b81617066565b825292840192908401906172f8565b80955050505050509250929050565b600081518084526020808501945080840160005b838110156173595781518752958201959082019060010161733d565b509495945050505050565b6020815260006149ee6020830184617329565b600081518084526020808501945080840160005b838110156173595781516001600160a01b03168752958201959082019060010161738b565b6020815260006149ee6020830184617377565b600082601f8301126173d457600080fd5b81356001600160401b038111156173ed576173ed6170b1565b617400601f8201601f1916602001617112565b81815284602083860101111561741557600080fd5b816020850160208301376000918101602001919091529392505050565b63ffffffff8116811461088857600080fd5b8035615f0681617432565b60008060006060848603121561746457600080fd5b833561746f81617066565b925060208401356001600160401b0381111561748a57600080fd5b617496868287016173c3565b92505060408401356174a781617432565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015617548578385038a52825180518087529087019087870190845b8181101561753357835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016174ef565b50509a87019a955050918501916001016174d1565b509298975050505050505050565b6020815260006149ee60208301846174b2565b801515811461088857600080fd5b60006020828403121561758957600080fd5b81356149ee81617569565b60005b838110156175af578181015183820152602001617597565b50506000910152565b600081518084526175d0816020860160208601617594565b601f01601f19169290920160200192915050565b60008151606084526175f960608501826175b8565b9050602083015160208501526040830151848203604086015261761c8282617377565b95945050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561767a57603f198886030184526176688583516175e4565b9450928501929085019060010161764c565b5092979650505050505050565b60006080828403121561769957600080fd5b50919050565b60006040828403121561769957600080fd5b600082601f8301126176c257600080fd5b813560206176d26172cf83617253565b82815260069290921b840181019181810190868411156176f157600080fd5b8286015b84811015617715576177078882617142565b8352918301916040016176f5565b509695505050505050565b60008060008060a0858703121561773657600080fd5b84356001600160401b038082111561774d57600080fd5b61775988838901617687565b9550602087013591508082111561776f57600080fd5b61777b8883890161769f565b945061778a886040890161769f565b935060808701359150808211156177a057600080fd5b506177ad878288016176b1565b91505092959194509250565b600082601f8301126177ca57600080fd5b813560206177da6172cf83617253565b82815260059290921b840181019181810190868411156177f957600080fd5b8286015b8481101561771557803583529183019183016177fd565b6000806040838503121561782757600080fd5b823561783281617066565b915060208301356001600160401b0381111561784d57600080fd5b617859858286016177b9565b9150509250929050565b60008083601f84011261787557600080fd5b5081356001600160401b0381111561788c57600080fd5b6020830191508360208285010111156178a457600080fd5b9250929050565b60008083601f8401126178bd57600080fd5b5081356001600160401b038111156178d457600080fd5b6020830191508360208260051b85010111156178a457600080fd5b6000806000806000806080878903121561790857600080fd5b863561791381617066565b9550602087013561792381617432565b945060408701356001600160401b038082111561793f57600080fd5b61794b8a838b01617863565b9096509450606089013591508082111561796457600080fd5b5061797189828a016178ab565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561735957815163ffffffff1687529582019590820190600101617997565b6000602080835283516080828501526179d560a0850182617983565b905081850151601f19808684030160408701526179f28383617983565b92506040870151915080868403016060870152617a0f8383617983565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015617a665784878303018452617a54828751617983565b95880195938801939150600101617a3a565b509998505050505050505050565b60ff8116811461088857600080fd5b600060208284031215617a9557600080fd5b81356149ee81617a74565b600080600060608486031215617ab557600080fd5b8335617ac081617066565b925060208401356001600160401b03811115617adb57600080fd5b617496868287016177b9565b6020808252825182820181905260009190848201906040850190845b81811015617b1f57835183529284019291840191600101617b03565b50909695505050505050565b600082601f830112617b3c57600080fd5b81356020617b4c6172cf83617253565b82815260059290921b84018101918181019086841115617b6b57600080fd5b8286015b84811015617715578035617b8281617432565b8352918301918301617b6f565b600082601f830112617ba057600080fd5b81356020617bb06172cf83617253565b82815260059290921b84018101918181019086841115617bcf57600080fd5b8286015b848110156177155780356001600160401b03811115617bf25760008081fd5b617c008986838b0101617b2b565b845250918301918301617bd3565b60006101808284031215617c2157600080fd5b617c296170ef565b905081356001600160401b0380821115617c4257600080fd5b617c4e85838601617b2b565b83526020840135915080821115617c6457600080fd5b617c70858386016176b1565b60208401526040840135915080821115617c8957600080fd5b617c95858386016176b1565b6040840152617ca785606086016171c2565b6060840152617cb98560e08601617142565b6080840152610120840135915080821115617cd357600080fd5b617cdf85838601617b2b565b60a0840152610140840135915080821115617cf957600080fd5b617d0585838601617b2b565b60c0840152610160840135915080821115617d1f57600080fd5b50617d2c84828501617b8f565b60e08301525092915050565b600080600080600060808688031215617d5057600080fd5b8535945060208601356001600160401b0380821115617d6e57600080fd5b617d7a89838a01617863565b909650945060408801359150617d8f82617432565b90925060608701359080821115617da557600080fd5b50617db288828901617c0e565b9150509295509295909350565b600081518084526020808501945080840160005b838110156173595781516001600160601b031687529582019590820190600101617dd3565b6040815260008351604080840152617e136080840182617dbf565b90506020850151603f19848303016060850152617e308282617dbf565b925050508260208301529392505050565b60008060008060008060c08789031215617e5a57600080fd5b8635617e6581617066565b95506020870135617e7581617066565b94506040870135617e8581617066565b93506060870135617e9581617066565b92506080870135617ea581617066565b915060a0870135617eb581617432565b809150509295509295509295565b600060608284031215617ed557600080fd5b604051606081016001600160401b038282108183111715617ef857617ef86170b1565b816040528293508435915080821115617f1057600080fd5b50617f1d858286016173c3565b82525060208301356020820152604083013560408201525092915050565b60008060408385031215617f4e57600080fd5b8235617f5981617066565b915060208301356001600160401b03811115617f7457600080fd5b61785985828601617ec3565b60008060008060608587031215617f9657600080fd5b8435617fa181617432565b93506020850135617fb181617432565b925060408501356001600160401b03811115617fcc57600080fd5b617fd887828801617863565b95989497509550505050565b600060208284031215617ff657600080fd5b81356001600160401b0381111561800c57600080fd5b618018848285016173c3565b949350505050565b600061802e6172cf84617253565b8381529050602080820190600585901b84018681111561804d57600080fd5b845b818110156180885780356001600160401b0381111561806e5760008081fd5b61807a898289016173c3565b85525092820192820161804f565b505050509392505050565b6000602082840312156180a557600080fd5b81356001600160401b038111156180bb57600080fd5b8201601f810184136180cc57600080fd5b61801884823560208401618020565b6000806000606084860312156180f057600080fd5b83356001600160401b038082111561810757600080fd5b61811387838801617687565b9450602086013591508082111561812957600080fd5b6181358783880161769f565b9350604086013591508082111561814b57600080fd5b5061815886828701617c0e565b9150509250925092565b60008060006060848603121561817757600080fd5b833561818281617066565b92506020840135915060408401356174a781617432565b82815260406020820152600061801860408301846174b2565b6020815260006149ee60208301846175e4565b600080604083850312156181d857600080fd5b82356181e381617066565b915060208301356181f381617a74565b809150509250929050565b60008060006060848603121561821357600080fd5b833561821e81617066565b925060208401356001600160401b038082111561823a57600080fd5b618246878388016173c3565b9350604086013591508082111561825c57600080fd5b5061815886828701617ec3565b6000806020838503121561827c57600080fd5b82356001600160401b0381111561829257600080fd5b61829e858286016178ab565b90969095509350505050565b6000602082840312156182bc57600080fd5b81516149ee81617066565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561832357600080fd5b81516149ee81617569565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826183a957634e487b7160e01b600052601260045260246000fd5b500690565b6000602082840312156183c057600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6000600182016183ef576183ef6183c7565b5060010190565b60006020828403121561840857600080fd5b81516001600160c01b03811681146149ee57600080fd5b60006020828403121561843157600080fd5b81516149ee81617a74565b80820180821115614708576147086183c7565b6001600160601b038116811461088857600080fd5b60006040828403121561847657600080fd5b61847e6170c7565b825161848981617066565b815260208301516184998161844f565b60208201529392505050565b600060208083850312156184b857600080fd5b82516001600160401b038111156184ce57600080fd5b8301601f810185136184df57600080fd5b80516184ed6172cf82617253565b81815260059190911b8201830190838101908783111561850c57600080fd5b928401925b8284101561852a57835182529284019290840190618511565b979650505050505050565b60006020828403121561854757600080fd5b81516149ee8161844f565b60006020828403121561856457600080fd5b81356149ee81617432565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e198436030181126185af57600080fd5b83016020810192503590506001600160401b038111156185ce57600080fd5b8036038213156178a457600080fd5b60006040830182356185ee81617432565b63ffffffff16845260208381013536859003601e1901811261860f57600080fd5b840181810190356001600160401b0381111561862a57600080fd5b8060051b80360383131561863d57600080fd5b6040848901529381905260609387018401938290880160005b8381101561869057898703605f190182526186718386618598565b61867c89828461856f565b985050509185019190850190600101618656565b509498975050505050505050565b6060815260006186b160608301856185dd565b905082356186be81617432565b63ffffffff8116602084015250602083013560408301529392505050565b63ffffffff818116838216019080821115610c6957610c696183c7565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561767a57815185529382019390820190600101618718565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561876157600080fd5b8260051b8085606085013791909101606001949350505050565b6000602080838503121561878e57600080fd5b82516001600160401b038111156187a457600080fd5b8301601f810185136187b557600080fd5b80516187c36172cf82617253565b81815260059190911b820183019083810190878311156187e257600080fd5b928401925b8284101561852a5783516187fa81617432565b825292840192908401906187e7565b63ffffffff8416815260406020820152600061761c60408301848661856f565b60006020828403121561883b57600080fd5b81516149ee81617432565b600060ff821660ff810361885c5761885c6183c7565b60010192915050565b60408152600061887960408301858761856f565b905063ffffffff83166020830152949350505050565b63ffffffff831681526040602082015260006180186040830184617329565b81810381811115614708576147086183c7565b6000602082840312156188d357600080fd5b815167ffffffffffffffff19811681146149ee57600080fd5b6001600160601b03828116828216039080821115610c6957610c696183c7565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526189ae60a08401826175b8565b90506020840151606084015260408401516080840152809150509392505050565b600063ffffffff808351168452806020840151166020850152604083015160806040860152618a0160808601826175b8565b9050816060850151166060860152809250505092915050565b6020815260006149ee60208301846189cf565b63ffffffff8316815260406020820152600061801860408301846189cf565b6020815260006149ee60208301846175b8565b6000808335601e19843603018112618a7657600080fd5b8301803591506001600160401b03821115618a9057600080fd5b6020019150368190038213156178a457600080fd5b6020815260008235618ab681617432565b63ffffffff808216602085015260208501359150618ad382617432565b8082166040850152618ae86040860186618598565b925060806060860152618aff60a08601848361856f565b9250506060850135618b1081617432565b166080939093019290925250919050565b6020815260006149ee60208301846185dd565b6001600160601b038181168382160280821691908281146155a1576155a16183c7565b8082028115828204841417614708576147086183c7565b6000808335601e19843603018112618b8557600080fd5b8301803591506001600160401b03821115618b9f57600080fd5b6020019150600581901b36038213156178a457600080fd5b60006149ee368484618020565b606081526000618bd760608301856185dd565b905063ffffffff8351166020830152602083015160408301529392505050565b60008251618c09818460208701617594565b9190910192915050565b600181811c90821680618c2757607f821691505b60208210810361769957634e487b7160e01b600052602260045260246000fd5b6001600160601b03818116838216019080821115610c6957610c696183c7565b600060408284031215618c7957600080fd5b618c816170c7565b8251815260208301516003811061849957600080fd5b60008235609e19833603018112618c0957600080fd5b8035615f0681617066565b8183526000602080850194508260005b85811015617359578135618cdb81617066565b6001600160a01b0316875281830135618cf38161844f565b6001600160601b0316878401526040968701969190910190600101618cc8565b60208082528181018390526000906040808401600586901b8501820187855b88811015618e1357878303603f190184528135368b9003609e19018112618d5857600080fd5b8a0160a0813536839003601e19018112618d7157600080fd5b820188810190356001600160401b03811115618d8c57600080fd5b8060061b3603821315618d9e57600080fd5b828752618dae8388018284618cb8565b92505050618dbd888301618cad565b6001600160a01b03168886015281870135878601526060618ddf818401617444565b63ffffffff16908601526080618df6838201617444565b63ffffffff16950194909452509285019290850190600101618d32565b509098975050505050505050565b600061ffff808316818103618e3857618e386183c7565b6001019392505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b634e487b7160e01b600052603160045260246000fd5b601f82111561649e57600081815260208120601f850160051c81016020861015618eca5750805b601f850160051c820191505b81811015613ce257828155600101618ed6565b81516001600160401b03811115618f0257618f026170b1565b618f1681618f108454618c13565b84618ea3565b602080601f831160018114618f4b5760008415618f335750858301515b600019600386901b1c1916600185901b178555613ce2565b600085815260208120601f198616915b82811015618f7a57888601518255948401946001909101908401618f5b565b5085821015618f985787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b0383168152604060208201819052600090618018908301846175b856fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47d4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220896a983522d519f1205c866a91ff78ffb89ac0c994c5eecae6da077a05c5aa5564736f6c63430008150033",
}

// ContractZrServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractZrServiceManagerMetaData.ABI instead.
var ContractZrServiceManagerABI = ContractZrServiceManagerMetaData.ABI

// ContractZrServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractZrServiceManagerMetaData.Bin instead.
var ContractZrServiceManagerBin = ContractZrServiceManagerMetaData.Bin

// DeployContractZrServiceManager deploys a new Ethereum contract, binding an instance of ContractZrServiceManager to it.
func DeployContractZrServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _rewardsCoordinator common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address) (common.Address, *types.Transaction, *ContractZrServiceManager, error) {
	parsed, err := ContractZrServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractZrServiceManagerBin), backend, _avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractZrServiceManager{ContractZrServiceManagerCaller: ContractZrServiceManagerCaller{contract: contract}, ContractZrServiceManagerTransactor: ContractZrServiceManagerTransactor{contract: contract}, ContractZrServiceManagerFilterer: ContractZrServiceManagerFilterer{contract: contract}}, nil
}

// ContractZrServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractZrServiceManager struct {
	ContractZrServiceManagerCaller     // Read-only binding to the contract
	ContractZrServiceManagerTransactor // Write-only binding to the contract
	ContractZrServiceManagerFilterer   // Log filterer for contract events
}

// ContractZrServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractZrServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractZrServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractZrServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractZrServiceManagerSession struct {
	Contract     *ContractZrServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractZrServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractZrServiceManagerCallerSession struct {
	Contract *ContractZrServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractZrServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractZrServiceManagerTransactorSession struct {
	Contract     *ContractZrServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractZrServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractZrServiceManagerRaw struct {
	Contract *ContractZrServiceManager // Generic contract binding to access the raw methods on
}

// ContractZrServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractZrServiceManagerCallerRaw struct {
	Contract *ContractZrServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractZrServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractZrServiceManagerTransactorRaw struct {
	Contract *ContractZrServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractZrServiceManager creates a new instance of ContractZrServiceManager, bound to a specific deployed contract.
func NewContractZrServiceManager(address common.Address, backend bind.ContractBackend) (*ContractZrServiceManager, error) {
	contract, err := bindContractZrServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManager{ContractZrServiceManagerCaller: ContractZrServiceManagerCaller{contract: contract}, ContractZrServiceManagerTransactor: ContractZrServiceManagerTransactor{contract: contract}, ContractZrServiceManagerFilterer: ContractZrServiceManagerFilterer{contract: contract}}, nil
}

// NewContractZrServiceManagerCaller creates a new read-only instance of ContractZrServiceManager, bound to a specific deployed contract.
func NewContractZrServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractZrServiceManagerCaller, error) {
	contract, err := bindContractZrServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerCaller{contract: contract}, nil
}

// NewContractZrServiceManagerTransactor creates a new write-only instance of ContractZrServiceManager, bound to a specific deployed contract.
func NewContractZrServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractZrServiceManagerTransactor, error) {
	contract, err := bindContractZrServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerTransactor{contract: contract}, nil
}

// NewContractZrServiceManagerFilterer creates a new log filterer instance of ContractZrServiceManager, bound to a specific deployed contract.
func NewContractZrServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractZrServiceManagerFilterer, error) {
	contract, err := bindContractZrServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerFilterer{contract: contract}, nil
}

// bindContractZrServiceManager binds a generic wrapper to an already deployed contract.
func bindContractZrServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractZrServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZrServiceManager *ContractZrServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZrServiceManager.Contract.ContractZrServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZrServiceManager *ContractZrServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.ContractZrServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZrServiceManager *ContractZrServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.ContractZrServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZrServiceManager *ContractZrServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZrServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.contract.Transact(opts, method, params...)
}

// QUORUMNUMBER is a free data retrieval call binding the contract method 0x2dda9fc6.
//
// Solidity: function QUORUM_NUMBER() view returns(uint8)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) QUORUMNUMBER(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "QUORUM_NUMBER")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QUORUMNUMBER is a free data retrieval call binding the contract method 0x2dda9fc6.
//
// Solidity: function QUORUM_NUMBER() view returns(uint8)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) QUORUMNUMBER() (uint8, error) {
	return _ContractZrServiceManager.Contract.QUORUMNUMBER(&_ContractZrServiceManager.CallOpts)
}

// QUORUMNUMBER is a free data retrieval call binding the contract method 0x2dda9fc6.
//
// Solidity: function QUORUM_NUMBER() view returns(uint8)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) QUORUMNUMBER() (uint8, error) {
	return _ContractZrServiceManager.Contract.QUORUMNUMBER(&_ContractZrServiceManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZrServiceManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZrServiceManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZrServiceManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZrServiceManager.CallOpts)
}

// UNDELEGATIONPERIOD is a free data retrieval call binding the contract method 0x7533f901.
//
// Solidity: function UNDELEGATION_PERIOD() view returns(uint256)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) UNDELEGATIONPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "UNDELEGATION_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNDELEGATIONPERIOD is a free data retrieval call binding the contract method 0x7533f901.
//
// Solidity: function UNDELEGATION_PERIOD() view returns(uint256)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) UNDELEGATIONPERIOD() (*big.Int, error) {
	return _ContractZrServiceManager.Contract.UNDELEGATIONPERIOD(&_ContractZrServiceManager.CallOpts)
}

// UNDELEGATIONPERIOD is a free data retrieval call binding the contract method 0x7533f901.
//
// Solidity: function UNDELEGATION_PERIOD() view returns(uint256)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) UNDELEGATIONPERIOD() (*big.Int, error) {
	return _ContractZrServiceManager.Contract.UNDELEGATIONPERIOD(&_ContractZrServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractZrServiceManager.Contract.AvsDirectory(&_ContractZrServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractZrServiceManager.Contract.AvsDirectory(&_ContractZrServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZrServiceManager.Contract.BlsApkRegistry(&_ContractZrServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZrServiceManager.Contract.BlsApkRegistry(&_ContractZrServiceManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZrServiceManager.Contract.CheckSignatures(&_ContractZrServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZrServiceManager.Contract.CheckSignatures(&_ContractZrServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Delegation() (common.Address, error) {
	return _ContractZrServiceManager.Contract.Delegation(&_ContractZrServiceManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractZrServiceManager.Contract.Delegation(&_ContractZrServiceManager.CallOpts)
}

// GetAllOperatorsAddresses is a free data retrieval call binding the contract method 0x889e2c6e.
//
// Solidity: function getAllOperatorsAddresses() view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetAllOperatorsAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getAllOperatorsAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllOperatorsAddresses is a free data retrieval call binding the contract method 0x889e2c6e.
//
// Solidity: function getAllOperatorsAddresses() view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetAllOperatorsAddresses() ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetAllOperatorsAddresses(&_ContractZrServiceManager.CallOpts)
}

// GetAllOperatorsAddresses is a free data retrieval call binding the contract method 0x889e2c6e.
//
// Solidity: function getAllOperatorsAddresses() view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetAllOperatorsAddresses() ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetAllOperatorsAddresses(&_ContractZrServiceManager.CallOpts)
}

// GetAllValidator is a free data retrieval call binding the contract method 0x4a91a2f8.
//
// Solidity: function getAllValidator() view returns((string,bytes32,address[])[])
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetAllValidator(opts *bind.CallOpts) ([]ZrServiceManagerLibValidator, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getAllValidator")

	if err != nil {
		return *new([]ZrServiceManagerLibValidator), err
	}

	out0 := *abi.ConvertType(out[0], new([]ZrServiceManagerLibValidator)).(*[]ZrServiceManagerLibValidator)

	return out0, err

}

// GetAllValidator is a free data retrieval call binding the contract method 0x4a91a2f8.
//
// Solidity: function getAllValidator() view returns((string,bytes32,address[])[])
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetAllValidator() ([]ZrServiceManagerLibValidator, error) {
	return _ContractZrServiceManager.Contract.GetAllValidator(&_ContractZrServiceManager.CallOpts)
}

// GetAllValidator is a free data retrieval call binding the contract method 0x4a91a2f8.
//
// Solidity: function getAllValidator() view returns((string,bytes32,address[])[])
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetAllValidator() ([]ZrServiceManagerLibValidator, error) {
	return _ContractZrServiceManager.Contract.GetAllValidator(&_ContractZrServiceManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetBatchOperatorFromId(&_ContractZrServiceManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetBatchOperatorFromId(&_ContractZrServiceManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractZrServiceManager.Contract.GetBatchOperatorId(&_ContractZrServiceManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractZrServiceManager.Contract.GetBatchOperatorId(&_ContractZrServiceManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZrServiceManager.Contract.GetCheckSignaturesIndices(&_ContractZrServiceManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZrServiceManager.Contract.GetCheckSignaturesIndices(&_ContractZrServiceManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetEigenStake is a free data retrieval call binding the contract method 0xda4dc49c.
//
// Solidity: function getEigenStake(address operator, uint8 quorumNumber) view returns(uint96)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetEigenStake(opts *bind.CallOpts, operator common.Address, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getEigenStake", operator, quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEigenStake is a free data retrieval call binding the contract method 0xda4dc49c.
//
// Solidity: function getEigenStake(address operator, uint8 quorumNumber) view returns(uint96)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetEigenStake(operator common.Address, quorumNumber uint8) (*big.Int, error) {
	return _ContractZrServiceManager.Contract.GetEigenStake(&_ContractZrServiceManager.CallOpts, operator, quorumNumber)
}

// GetEigenStake is a free data retrieval call binding the contract method 0xda4dc49c.
//
// Solidity: function getEigenStake(address operator, uint8 quorumNumber) view returns(uint96)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetEigenStake(operator common.Address, quorumNumber uint8) (*big.Int, error) {
	return _ContractZrServiceManager.Contract.GetEigenStake(&_ContractZrServiceManager.CallOpts, operator, quorumNumber)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractZrServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractZrServiceManager.CallOpts, operator)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZrServiceManager.Contract.GetOperatorState(&_ContractZrServiceManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZrServiceManager.Contract.GetOperatorState(&_ContractZrServiceManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZrServiceManager.Contract.GetOperatorState0(&_ContractZrServiceManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZrServiceManager.Contract.GetOperatorState0(&_ContractZrServiceManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZrServiceManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZrServiceManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZrServiceManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZrServiceManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetRestakeableStrategies(&_ContractZrServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractZrServiceManager.Contract.GetRestakeableStrategies(&_ContractZrServiceManager.CallOpts)
}

// GetStakedBalanceForValidator is a free data retrieval call binding the contract method 0xd604902f.
//
// Solidity: function getStakedBalanceForValidator(string validatorAddr) view returns(uint96)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetStakedBalanceForValidator(opts *bind.CallOpts, validatorAddr string) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getStakedBalanceForValidator", validatorAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakedBalanceForValidator is a free data retrieval call binding the contract method 0xd604902f.
//
// Solidity: function getStakedBalanceForValidator(string validatorAddr) view returns(uint96)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetStakedBalanceForValidator(validatorAddr string) (*big.Int, error) {
	return _ContractZrServiceManager.Contract.GetStakedBalanceForValidator(&_ContractZrServiceManager.CallOpts, validatorAddr)
}

// GetStakedBalanceForValidator is a free data retrieval call binding the contract method 0xd604902f.
//
// Solidity: function getStakedBalanceForValidator(string validatorAddr) view returns(uint96)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetStakedBalanceForValidator(validatorAddr string) (*big.Int, error) {
	return _ContractZrServiceManager.Contract.GetStakedBalanceForValidator(&_ContractZrServiceManager.CallOpts, validatorAddr)
}

// GetTaskNumber is a free data retrieval call binding the contract method 0x7b654c5d.
//
// Solidity: function getTaskNumber() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetTaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getTaskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskNumber is a free data retrieval call binding the contract method 0x7b654c5d.
//
// Solidity: function getTaskNumber() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetTaskNumber() (uint32, error) {
	return _ContractZrServiceManager.Contract.GetTaskNumber(&_ContractZrServiceManager.CallOpts)
}

// GetTaskNumber is a free data retrieval call binding the contract method 0x7b654c5d.
//
// Solidity: function getTaskNumber() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetTaskNumber() (uint32, error) {
	return _ContractZrServiceManager.Contract.GetTaskNumber(&_ContractZrServiceManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZrServiceManager.Contract.GetTaskResponseWindowBlock(&_ContractZrServiceManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZrServiceManager.Contract.GetTaskResponseWindowBlock(&_ContractZrServiceManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validatorHash) view returns((string,bytes32,address[]))
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) GetValidator(opts *bind.CallOpts, validatorHash [32]byte) (ZrServiceManagerLibValidator, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "getValidator", validatorHash)

	if err != nil {
		return *new(ZrServiceManagerLibValidator), err
	}

	out0 := *abi.ConvertType(out[0], new(ZrServiceManagerLibValidator)).(*ZrServiceManagerLibValidator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validatorHash) view returns((string,bytes32,address[]))
func (_ContractZrServiceManager *ContractZrServiceManagerSession) GetValidator(validatorHash [32]byte) (ZrServiceManagerLibValidator, error) {
	return _ContractZrServiceManager.Contract.GetValidator(&_ContractZrServiceManager.CallOpts, validatorHash)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validatorHash) view returns((string,bytes32,address[]))
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) GetValidator(validatorHash [32]byte) (ZrServiceManagerLibValidator, error) {
	return _ContractZrServiceManager.Contract.GetValidator(&_ContractZrServiceManager.CallOpts, validatorHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Owner() (common.Address, error) {
	return _ContractZrServiceManager.Contract.Owner(&_ContractZrServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractZrServiceManager.Contract.Owner(&_ContractZrServiceManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Paused(index uint8) (bool, error) {
	return _ContractZrServiceManager.Contract.Paused(&_ContractZrServiceManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractZrServiceManager.Contract.Paused(&_ContractZrServiceManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Paused0() (*big.Int, error) {
	return _ContractZrServiceManager.Contract.Paused0(&_ContractZrServiceManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractZrServiceManager.Contract.Paused0(&_ContractZrServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractZrServiceManager.Contract.PauserRegistry(&_ContractZrServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractZrServiceManager.Contract.PauserRegistry(&_ContractZrServiceManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZrServiceManager.Contract.RegistryCoordinator(&_ContractZrServiceManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZrServiceManager.Contract.RegistryCoordinator(&_ContractZrServiceManager.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) RewardsInitiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "rewardsInitiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) RewardsInitiator() (common.Address, error) {
	return _ContractZrServiceManager.Contract.RewardsInitiator(&_ContractZrServiceManager.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) RewardsInitiator() (common.Address, error) {
	return _ContractZrServiceManager.Contract.RewardsInitiator(&_ContractZrServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractZrServiceManager.Contract.StakeRegistry(&_ContractZrServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractZrServiceManager.Contract.StakeRegistry(&_ContractZrServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZrServiceManager.Contract.StaleStakesForbidden(&_ContractZrServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZrServiceManager.Contract.StaleStakesForbidden(&_ContractZrServiceManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZrServiceManager.Contract.TrySignatureAndApkVerification(&_ContractZrServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZrServiceManager.Contract.TrySignatureAndApkVerification(&_ContractZrServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.CreateAVSRewardsSubmission(&_ContractZrServiceManager.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.CreateAVSRewardsSubmission(&_ContractZrServiceManager.TransactOpts, rewardsSubmissions)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) CreateNewTask(opts *bind.TransactOpts, taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "createNewTask", taskId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) CreateNewTask(taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.CreateNewTask(&_ContractZrServiceManager.TransactOpts, taskId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) CreateNewTask(taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.CreateNewTask(&_ContractZrServiceManager.TransactOpts, taskId, quorumThresholdPercentage, quorumNumbers)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address oprAddr) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, oprAddr common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", oprAddr)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address oprAddr) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) DeregisterOperatorFromAVS(oprAddr common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractZrServiceManager.TransactOpts, oprAddr)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address oprAddr) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) DeregisterOperatorFromAVS(oprAddr common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractZrServiceManager.TransactOpts, oprAddr)
}

// EjectValidators is a paid mutator transaction binding the contract method 0xc63e3c50.
//
// Solidity: function ejectValidators(string[] inactiveValidatorSet) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) EjectValidators(opts *bind.TransactOpts, inactiveValidatorSet []string) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "ejectValidators", inactiveValidatorSet)
}

// EjectValidators is a paid mutator transaction binding the contract method 0xc63e3c50.
//
// Solidity: function ejectValidators(string[] inactiveValidatorSet) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) EjectValidators(inactiveValidatorSet []string) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.EjectValidators(&_ContractZrServiceManager.TransactOpts, inactiveValidatorSet)
}

// EjectValidators is a paid mutator transaction binding the contract method 0xc63e3c50.
//
// Solidity: function ejectValidators(string[] inactiveValidatorSet) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) EjectValidators(inactiveValidatorSet []string) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.EjectValidators(&_ContractZrServiceManager.TransactOpts, inactiveValidatorSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c5178cf.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _aggregator, address _generator, address _rewardsInitiator, uint32 _taskResponseWindowBlock) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, _initialOwner common.Address, _aggregator common.Address, _generator common.Address, _rewardsInitiator common.Address, _taskResponseWindowBlock uint32) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "initialize", _pauserRegistry, _initialOwner, _aggregator, _generator, _rewardsInitiator, _taskResponseWindowBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c5178cf.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _aggregator, address _generator, address _rewardsInitiator, uint32 _taskResponseWindowBlock) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Initialize(_pauserRegistry common.Address, _initialOwner common.Address, _aggregator common.Address, _generator common.Address, _rewardsInitiator common.Address, _taskResponseWindowBlock uint32) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Initialize(&_ContractZrServiceManager.TransactOpts, _pauserRegistry, _initialOwner, _aggregator, _generator, _rewardsInitiator, _taskResponseWindowBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c5178cf.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _aggregator, address _generator, address _rewardsInitiator, uint32 _taskResponseWindowBlock) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) Initialize(_pauserRegistry common.Address, _initialOwner common.Address, _aggregator common.Address, _generator common.Address, _rewardsInitiator common.Address, _taskResponseWindowBlock uint32) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Initialize(&_ContractZrServiceManager.TransactOpts, _pauserRegistry, _initialOwner, _aggregator, _generator, _rewardsInitiator, _taskResponseWindowBlock)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Pause(&_ContractZrServiceManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Pause(&_ContractZrServiceManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.PauseAll(&_ContractZrServiceManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.PauseAll(&_ContractZrServiceManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c71c165.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task ZrServiceManagerLibTask, taskResponse ZrServiceManagerLibTaskResponse, taskResponseMetadata ZrServiceManagerLibTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c71c165.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) RaiseAndResolveChallenge(task ZrServiceManagerLibTask, taskResponse ZrServiceManagerLibTaskResponse, taskResponseMetadata ZrServiceManagerLibTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RaiseAndResolveChallenge(&_ContractZrServiceManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c71c165.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) RaiseAndResolveChallenge(task ZrServiceManagerLibTask, taskResponse ZrServiceManagerLibTaskResponse, taskResponseMetadata ZrServiceManagerLibTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RaiseAndResolveChallenge(&_ContractZrServiceManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RegisterOperatorToAVS(&_ContractZrServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RegisterOperatorToAVS(&_ContractZrServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS0 is a paid mutator transaction binding the contract method 0xf891899d.
//
// Solidity: function registerOperatorToAVS(address operatorAddr, string validatorAddr, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) RegisterOperatorToAVS0(opts *bind.TransactOpts, operatorAddr common.Address, validatorAddr string, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "registerOperatorToAVS0", operatorAddr, validatorAddr, operatorSignature)
}

// RegisterOperatorToAVS0 is a paid mutator transaction binding the contract method 0xf891899d.
//
// Solidity: function registerOperatorToAVS(address operatorAddr, string validatorAddr, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) RegisterOperatorToAVS0(operatorAddr common.Address, validatorAddr string, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RegisterOperatorToAVS0(&_ContractZrServiceManager.TransactOpts, operatorAddr, validatorAddr, operatorSignature)
}

// RegisterOperatorToAVS0 is a paid mutator transaction binding the contract method 0xf891899d.
//
// Solidity: function registerOperatorToAVS(address operatorAddr, string validatorAddr, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) RegisterOperatorToAVS0(operatorAddr common.Address, validatorAddr string, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RegisterOperatorToAVS0(&_ContractZrServiceManager.TransactOpts, operatorAddr, validatorAddr, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RenounceOwnership(&_ContractZrServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RenounceOwnership(&_ContractZrServiceManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) RespondToTask(opts *bind.TransactOpts, task ZrServiceManagerLibTask, taskResponse ZrServiceManagerLibTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) RespondToTask(task ZrServiceManagerLibTask, taskResponse ZrServiceManagerLibTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RespondToTask(&_ContractZrServiceManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) RespondToTask(task ZrServiceManagerLibTask, taskResponse ZrServiceManagerLibTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.RespondToTask(&_ContractZrServiceManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.SetPauserRegistry(&_ContractZrServiceManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.SetPauserRegistry(&_ContractZrServiceManager.TransactOpts, newPauserRegistry)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) SetRewardsInitiator(opts *bind.TransactOpts, newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "setRewardsInitiator", newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.SetRewardsInitiator(&_ContractZrServiceManager.TransactOpts, newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.SetRewardsInitiator(&_ContractZrServiceManager.TransactOpts, newRewardsInitiator)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.SetStaleStakesForbidden(&_ContractZrServiceManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.SetStaleStakesForbidden(&_ContractZrServiceManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.TransferOwnership(&_ContractZrServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.TransferOwnership(&_ContractZrServiceManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Unpause(&_ContractZrServiceManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Unpause(&_ContractZrServiceManager.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.UpdateAVSMetadataURI(&_ContractZrServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.UpdateAVSMetadataURI(&_ContractZrServiceManager.TransactOpts, _metadataURI)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0xfb558ed5.
//
// Solidity: function updateOperators() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) UpdateOperators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "updateOperators")
}

// UpdateOperators is a paid mutator transaction binding the contract method 0xfb558ed5.
//
// Solidity: function updateOperators() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) UpdateOperators() (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.UpdateOperators(&_ContractZrServiceManager.TransactOpts)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0xfb558ed5.
//
// Solidity: function updateOperators() returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) UpdateOperators() (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.UpdateOperators(&_ContractZrServiceManager.TransactOpts)
}

// ContractZrServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerInitializedIterator struct {
	Event *ContractZrServiceManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerInitialized represents a Initialized event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractZrServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerInitializedIterator{contract: _ContractZrServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerInitialized)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractZrServiceManagerInitialized, error) {
	event := new(ContractZrServiceManagerInitialized)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerNewTaskCreatedIterator struct {
	Event *ContractZrServiceManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerNewTaskCreated struct {
	TaskId uint32
	Task   ZrServiceManagerLibTask
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xf31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5.
//
// Solidity: event NewTaskCreated(uint32 taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts) (*ContractZrServiceManagerNewTaskCreatedIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "NewTaskCreated")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerNewTaskCreatedIterator{contract: _ContractZrServiceManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xf31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5.
//
// Solidity: event NewTaskCreated(uint32 taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerNewTaskCreated) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "NewTaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerNewTaskCreated)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0xf31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5.
//
// Solidity: event NewTaskCreated(uint32 taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractZrServiceManagerNewTaskCreated, error) {
	event := new(ContractZrServiceManagerNewTaskCreated)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerOperatorAssignedIterator is returned from FilterOperatorAssigned and is used to iterate over the raw logs and unpacked data for OperatorAssigned events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerOperatorAssignedIterator struct {
	Event *ContractZrServiceManagerOperatorAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerOperatorAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerOperatorAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerOperatorAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerOperatorAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerOperatorAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerOperatorAssigned represents a OperatorAssigned event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerOperatorAssigned struct {
	OperatorAddr  common.Address
	ValidatorHash [32]byte
	Stake         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAssigned is a free log retrieval operation binding the contract event 0x5ea87e0bfaffd2b7f2901f087737e2b175f237068c6123ca9ea5240076bb5668.
//
// Solidity: event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash, uint96 stake)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterOperatorAssigned(opts *bind.FilterOpts, operatorAddr []common.Address, validatorHash [][32]byte) (*ContractZrServiceManagerOperatorAssignedIterator, error) {

	var operatorAddrRule []interface{}
	for _, operatorAddrItem := range operatorAddr {
		operatorAddrRule = append(operatorAddrRule, operatorAddrItem)
	}
	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "OperatorAssigned", operatorAddrRule, validatorHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerOperatorAssignedIterator{contract: _ContractZrServiceManager.contract, event: "OperatorAssigned", logs: logs, sub: sub}, nil
}

// WatchOperatorAssigned is a free log subscription operation binding the contract event 0x5ea87e0bfaffd2b7f2901f087737e2b175f237068c6123ca9ea5240076bb5668.
//
// Solidity: event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash, uint96 stake)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchOperatorAssigned(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerOperatorAssigned, operatorAddr []common.Address, validatorHash [][32]byte) (event.Subscription, error) {

	var operatorAddrRule []interface{}
	for _, operatorAddrItem := range operatorAddr {
		operatorAddrRule = append(operatorAddrRule, operatorAddrItem)
	}
	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "OperatorAssigned", operatorAddrRule, validatorHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerOperatorAssigned)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "OperatorAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorAssigned is a log parse operation binding the contract event 0x5ea87e0bfaffd2b7f2901f087737e2b175f237068c6123ca9ea5240076bb5668.
//
// Solidity: event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash, uint96 stake)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseOperatorAssigned(log types.Log) (*ContractZrServiceManagerOperatorAssigned, error) {
	event := new(ContractZrServiceManagerOperatorAssigned)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "OperatorAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerOperatorDeregisteredIterator struct {
	Event *ContractZrServiceManagerOperatorDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerOperatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerOperatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerOperatorDeregistered represents a OperatorDeregistered event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerOperatorDeregistered struct {
	OperatorAddr  common.Address
	ValidatorHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operatorAddr, bytes32 indexed validatorHash)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operatorAddr []common.Address, validatorHash [][32]byte) (*ContractZrServiceManagerOperatorDeregisteredIterator, error) {

	var operatorAddrRule []interface{}
	for _, operatorAddrItem := range operatorAddr {
		operatorAddrRule = append(operatorAddrRule, operatorAddrItem)
	}
	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "OperatorDeregistered", operatorAddrRule, validatorHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerOperatorDeregisteredIterator{contract: _ContractZrServiceManager.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operatorAddr, bytes32 indexed validatorHash)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerOperatorDeregistered, operatorAddr []common.Address, validatorHash [][32]byte) (event.Subscription, error) {

	var operatorAddrRule []interface{}
	for _, operatorAddrItem := range operatorAddr {
		operatorAddrRule = append(operatorAddrRule, operatorAddrItem)
	}
	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "OperatorDeregistered", operatorAddrRule, validatorHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerOperatorDeregistered)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operatorAddr, bytes32 indexed validatorHash)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseOperatorDeregistered(log types.Log) (*ContractZrServiceManagerOperatorDeregistered, error) {
	event := new(ContractZrServiceManagerOperatorDeregistered)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerOwnershipTransferredIterator struct {
	Event *ContractZrServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractZrServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerOwnershipTransferredIterator{contract: _ContractZrServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerOwnershipTransferred)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractZrServiceManagerOwnershipTransferred, error) {
	event := new(ContractZrServiceManagerOwnershipTransferred)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerPausedIterator struct {
	Event *ContractZrServiceManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerPaused represents a Paused event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractZrServiceManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerPausedIterator{contract: _ContractZrServiceManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerPaused)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParsePaused(log types.Log) (*ContractZrServiceManagerPaused, error) {
	event := new(ContractZrServiceManagerPaused)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerPauserRegistrySetIterator struct {
	Event *ContractZrServiceManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractZrServiceManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerPauserRegistrySetIterator{contract: _ContractZrServiceManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerPauserRegistrySet)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractZrServiceManagerPauserRegistrySet, error) {
	event := new(ContractZrServiceManagerPauserRegistrySet)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerRewardsInitiatorUpdatedIterator is returned from FilterRewardsInitiatorUpdated and is used to iterate over the raw logs and unpacked data for RewardsInitiatorUpdated events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerRewardsInitiatorUpdatedIterator struct {
	Event *ContractZrServiceManagerRewardsInitiatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerRewardsInitiatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerRewardsInitiatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerRewardsInitiatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerRewardsInitiatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerRewardsInitiatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerRewardsInitiatorUpdated represents a RewardsInitiatorUpdated event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerRewardsInitiatorUpdated struct {
	PrevRewardsInitiator common.Address
	NewRewardsInitiator  common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRewardsInitiatorUpdated is a free log retrieval operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterRewardsInitiatorUpdated(opts *bind.FilterOpts) (*ContractZrServiceManagerRewardsInitiatorUpdatedIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerRewardsInitiatorUpdatedIterator{contract: _ContractZrServiceManager.contract, event: "RewardsInitiatorUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsInitiatorUpdated is a free log subscription operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchRewardsInitiatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerRewardsInitiatorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerRewardsInitiatorUpdated)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsInitiatorUpdated is a log parse operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseRewardsInitiatorUpdated(log types.Log) (*ContractZrServiceManagerRewardsInitiatorUpdated, error) {
	event := new(ContractZrServiceManagerRewardsInitiatorUpdated)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractZrServiceManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractZrServiceManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerStaleStakesForbiddenUpdateIterator{contract: _ContractZrServiceManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerStaleStakesForbiddenUpdate)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractZrServiceManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractZrServiceManagerStaleStakesForbiddenUpdate)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractZrServiceManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerTaskChallengedSuccessfully struct {
	TaskId     uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 taskId, address challenger)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts) (*ContractZrServiceManagerTaskChallengedSuccessfullyIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerTaskChallengedSuccessfullyIterator{contract: _ContractZrServiceManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 taskId, address challenger)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerTaskChallengedSuccessfully) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerTaskChallengedSuccessfully)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 taskId, address challenger)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractZrServiceManagerTaskChallengedSuccessfully, error) {
	event := new(ContractZrServiceManagerTaskChallengedSuccessfully)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerTaskRespondedIterator struct {
	Event *ContractZrServiceManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerTaskResponded represents a TaskResponded event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerTaskResponded struct {
	TaskResponse         ZrServiceManagerLibTaskResponse
	TaskResponseMetadata ZrServiceManagerLibTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567.
//
// Solidity: event TaskResponded((uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractZrServiceManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerTaskRespondedIterator{contract: _ContractZrServiceManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567.
//
// Solidity: event TaskResponded((uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerTaskResponded)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0xb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567.
//
// Solidity: event TaskResponded((uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseTaskResponded(log types.Log) (*ContractZrServiceManagerTaskResponded, error) {
	event := new(ContractZrServiceManagerTaskResponded)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerUnpausedIterator struct {
	Event *ContractZrServiceManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerUnpaused represents a Unpaused event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractZrServiceManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerUnpausedIterator{contract: _ContractZrServiceManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerUnpaused)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseUnpaused(log types.Log) (*ContractZrServiceManagerUnpaused, error) {
	event := new(ContractZrServiceManagerUnpaused)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerValidatorDeregisteredIterator is returned from FilterValidatorDeregistered and is used to iterate over the raw logs and unpacked data for ValidatorDeregistered events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerValidatorDeregisteredIterator struct {
	Event *ContractZrServiceManagerValidatorDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerValidatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerValidatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerValidatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerValidatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerValidatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerValidatorDeregistered represents a ValidatorDeregistered event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerValidatorDeregistered struct {
	ValidatorHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeregistered is a free log retrieval operation binding the contract event 0xec3ecc6d7c19c6fc03c8b840fabed1a1b70c73f74bdb2d71f18c2f5a23639d32.
//
// Solidity: event ValidatorDeregistered(bytes32 indexed validatorHash)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterValidatorDeregistered(opts *bind.FilterOpts, validatorHash [][32]byte) (*ContractZrServiceManagerValidatorDeregisteredIterator, error) {

	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "ValidatorDeregistered", validatorHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerValidatorDeregisteredIterator{contract: _ContractZrServiceManager.contract, event: "ValidatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchValidatorDeregistered is a free log subscription operation binding the contract event 0xec3ecc6d7c19c6fc03c8b840fabed1a1b70c73f74bdb2d71f18c2f5a23639d32.
//
// Solidity: event ValidatorDeregistered(bytes32 indexed validatorHash)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchValidatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerValidatorDeregistered, validatorHash [][32]byte) (event.Subscription, error) {

	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "ValidatorDeregistered", validatorHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerValidatorDeregistered)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeregistered is a log parse operation binding the contract event 0xec3ecc6d7c19c6fc03c8b840fabed1a1b70c73f74bdb2d71f18c2f5a23639d32.
//
// Solidity: event ValidatorDeregistered(bytes32 indexed validatorHash)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseValidatorDeregistered(log types.Log) (*ContractZrServiceManagerValidatorDeregistered, error) {
	event := new(ContractZrServiceManagerValidatorDeregistered)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerValidatorRegisteredIterator struct {
	Event *ContractZrServiceManagerValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerValidatorRegistered represents a ValidatorRegistered event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerValidatorRegistered struct {
	ValidatorHash [32]byte
	ValidatorAddr string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea3.
//
// Solidity: event ValidatorRegistered(bytes32 indexed validatorHash, string validatorAddr)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validatorHash [][32]byte) (*ContractZrServiceManagerValidatorRegisteredIterator, error) {

	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "ValidatorRegistered", validatorHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerValidatorRegisteredIterator{contract: _ContractZrServiceManager.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea3.
//
// Solidity: event ValidatorRegistered(bytes32 indexed validatorHash, string validatorAddr)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerValidatorRegistered, validatorHash [][32]byte) (event.Subscription, error) {

	var validatorHashRule []interface{}
	for _, validatorHashItem := range validatorHash {
		validatorHashRule = append(validatorHashRule, validatorHashItem)
	}

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "ValidatorRegistered", validatorHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerValidatorRegistered)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0x1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea3.
//
// Solidity: event ValidatorRegistered(bytes32 indexed validatorHash, string validatorAddr)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseValidatorRegistered(log types.Log) (*ContractZrServiceManagerValidatorRegistered, error) {
	event := new(ContractZrServiceManagerValidatorRegistered)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrServiceManagerValidatorsEjectedIterator is returned from FilterValidatorsEjected and is used to iterate over the raw logs and unpacked data for ValidatorsEjected events raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerValidatorsEjectedIterator struct {
	Event *ContractZrServiceManagerValidatorsEjected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractZrServiceManagerValidatorsEjectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrServiceManagerValidatorsEjected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractZrServiceManagerValidatorsEjected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractZrServiceManagerValidatorsEjectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrServiceManagerValidatorsEjectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrServiceManagerValidatorsEjected represents a ValidatorsEjected event raised by the ContractZrServiceManager contract.
type ContractZrServiceManagerValidatorsEjected struct {
	ValidatorHashes [][32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorsEjected is a free log retrieval operation binding the contract event 0xb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a.
//
// Solidity: event ValidatorsEjected(bytes32[] validatorHashes)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) FilterValidatorsEjected(opts *bind.FilterOpts) (*ContractZrServiceManagerValidatorsEjectedIterator, error) {

	logs, sub, err := _ContractZrServiceManager.contract.FilterLogs(opts, "ValidatorsEjected")
	if err != nil {
		return nil, err
	}
	return &ContractZrServiceManagerValidatorsEjectedIterator{contract: _ContractZrServiceManager.contract, event: "ValidatorsEjected", logs: logs, sub: sub}, nil
}

// WatchValidatorsEjected is a free log subscription operation binding the contract event 0xb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a.
//
// Solidity: event ValidatorsEjected(bytes32[] validatorHashes)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) WatchValidatorsEjected(opts *bind.WatchOpts, sink chan<- *ContractZrServiceManagerValidatorsEjected) (event.Subscription, error) {

	logs, sub, err := _ContractZrServiceManager.contract.WatchLogs(opts, "ValidatorsEjected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrServiceManagerValidatorsEjected)
				if err := _ContractZrServiceManager.contract.UnpackLog(event, "ValidatorsEjected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsEjected is a log parse operation binding the contract event 0xb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a.
//
// Solidity: event ValidatorsEjected(bytes32[] validatorHashes)
func (_ContractZrServiceManager *ContractZrServiceManagerFilterer) ParseValidatorsEjected(log types.Log) (*ContractZrServiceManagerValidatorsEjected, error) {
	event := new(ContractZrServiceManagerValidatorsEjected)
	if err := _ContractZrServiceManager.contract.UnpackLog(event, "ValidatorsEjected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
