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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIZrRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"QUORUM_NUMBER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNDELEGATION_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"oprAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectValidators\",\"inputs\":[{\"name\":\"inactiveValidatorSet\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllOperatorsAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structZrServiceManagerLib.Validator[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEigenStake\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakedBalanceForValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Validator\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAssigned\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDeregistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorsEjected\",\"inputs\":[{\"name\":\"validatorHashes\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6101806040523480156200001257600080fd5b50604051620093b9380380620093b98339810160408190526200003591620002de565b8383838383838383896001600081905550806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200009f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000c5919062000346565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200011d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000143919062000346565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200019d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001c3919062000346565b6001600160a01b0390811660e0529485166101005250918316610120528216610140521661016052620001f562000203565b50505050505050506200036d565b606554610100900460ff1615620002705760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60655460ff9081161015620002c3576065805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114620002db57600080fd5b50565b60008060008060808587031215620002f557600080fd5b84516200030281620002c5565b60208601519094506200031581620002c5565b60408601519093506200032881620002c5565b60608601519092506200033b81620002c5565b939692955090935050565b6000602082840312156200035957600080fd5b81516200036681620002c5565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051618eff620004ba60003960008181610e3c01528181610f970152818161102e015281816146f0015281816148c201528181614a450152614ae4015260008181610c6701528181610cf601528181610d76015281816137e30152818161393501528181613a2201528181613ca20152818161471a015281816147fd015281816149a001528181614c44015281816150890152616c30015260008181615201015281816152bd01526153a90152600081816104d401528181613a7601528181613d0701528181613d860152614ca10152600081816106c70152612f9a01526000818161047101528181611cba015261317c0152600081816104b00152818161335201526135140152600081816104fd015281816115a901528181612c8401528181612dfd01526130370152618eff6000f3fe608060405234801561001057600080fd5b50600436106102bb5760003560e01c8063886f119511610182578063d604902f116100e9578063f63c5bab116100a2578063fabc1cbc1161007c578063fabc1cbc14610760578063fb558ed514610773578063fc299dee1461077b578063fce36c7d1461078e57600080fd5b8063f63c5bab14610731578063f891899d1461073a578063f8a2334c1461074d57600080fd5b8063d604902f14610684578063da4dc49c146106af578063df5cf723146106c2578063e481af9d146106e9578063f2fde38b146106f1578063f5c9899d1461070457600080fd5b8063a98fb3551161013b578063a98fb355146105fd578063b98d090814610610578063c63e3c501461061d578063caf73aa014610630578063cefdc1d414610643578063d5f20ff61461066457600080fd5b8063886f119514610598578063889e2c6e146105ab5780638da5cb5b146105b35780639926ee7d146105c45780639d3a0f2d146105d7578063a364f4da146105ea57600080fd5b8063595c6a67116102265780636b3aa72e116101df5780636b3aa72e146104d25780636d14a987146104f85780636efb46361461051f578063715018a6146105405780637533f901146105485780637b654c5d1461055257600080fd5b8063595c6a67146103ff5780635ac86ab7146104075780635c1556621461043a5780635c975abb1461045a5780635df459461461046c57806368304835146104ab57600080fd5b80633563b0d1116102785780633563b0d114610371578063416c7e5e146103915780634a91a2f8146103a45780634c71c165146103b95780634d2b57fe146103cc5780634f739f74146103df57600080fd5b806310d67a2f146102c0578063136439dd146102d5578063171f1d5b146102e85780632dda9fc61461031757806331b36bd91461033157806333cfb7b714610351575b600080fd5b6102d36102ce366004616f29565b6107a1565b005b6102d36102e3366004616f46565b61085d565b6102fb6102f63660046170b0565b61099c565b6040805192151583529015156020830152015b60405180910390f35b61031f600081565b60405160ff909116815260200161030e565b61034461033f366004617124565b610b26565b60405161030e9190617212565b61036461035f366004616f29565b610c42565b60405161030e919061725e565b61038461037f3660046172fd565b611111565b60405161030e9190617404565b6102d361039f366004617425565b6115a7565b6103ac6116de565b60405161030e91906174d3565b6102d36103c73660046175ce565b6117e4565b6103646103da3660046176c2565b611e03565b6103f26103ed36600461779d565b611f18565b60405161030e9190617867565b6102d3612640565b61042a610415366004617931565b600254600160ff9092169190911b9081161490565b604051901515815260200161030e565b61044d61044836600461794e565b612707565b60405161030e9190617995565b6002545b60405190815260200161030e565b6104937f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161030e565b6104937f000000000000000000000000000000000000000000000000000000000000000081565b7f0000000000000000000000000000000000000000000000000000000000000000610493565b6104937f000000000000000000000000000000000000000000000000000000000000000081565b61053261052d366004617be6565b6128cf565b60405161030e929190617ca6565b6102d36137c9565b61045e62024ea081565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0454640100000000900463ffffffff165b60405163ffffffff909116815260200161030e565b600154610493906001600160a01b031681565b6103646137dd565b6098546001600160a01b0316610493565b6102d36105d2366004617d67565b613a17565b6102d36105e5366004617dac565b613ae3565b6102d36105f8366004616f29565b613c97565b6102d361060b366004617e10565b613d67565b60335461042a9060ff1681565b6102d361062b366004617ebf565b613dbb565b6102d361063e366004617f07565b613e51565b610656610651366004617f8e565b614353565b60405161030e929190617fc5565b610677610672366004616f46565b6144e5565b60405161030e9190617fde565b610697610692366004617e10565b614510565b6040516001600160601b03909116815260200161030e565b6106976106bd366004617ff1565b6146ce565b6104937f000000000000000000000000000000000000000000000000000000000000000081565b6103646147f7565b6102d36106ff366004616f29565b614bc3565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f045463ffffffff16610583565b61058361271081565b6102d361074836600461802a565b614c39565b6102d361075b366004618095565b614d45565b6102d361076e366004616f46565b614f0a565b6102d3615066565b60ca54610493906001600160a01b031681565b6102d361079c366004618106565b6150bd565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108189190618147565b6001600160a01b0316336001600160a01b0316146108515760405162461bcd60e51b815260040161084890618164565b60405180910390fd5b61085a816153e0565b50565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156108a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c991906181ae565b6108e55760405162461bcd60e51b8152600401610848906181cb565b6002548181161461095e5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610848565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106109e4576109e4618213565b60200201518951600160200201518a60200151600060028110610a0957610a09618213565b60200201518b60200151600160028110610a2557610a25618213565b602090810291909101518c518d830151604051610a829a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610aa59190618229565b9050610b18610abe610ab788846154d7565b8690615568565b610ac66155fd565b610b0e610aff85610af9604080518082018252600080825260209182015281518083019092526001825260029082015290565b906154d7565b610b088c6156bd565b90615568565b886201d4c061574c565b909890975095505050505050565b606081516001600160401b03811115610b4157610b41616f5f565b604051908082528060200260200182016040528015610b6a578160200160208202803683370190505b50905060005b8251811015610c3b57836001600160a01b03166313542a4e848381518110610b9a57610b9a618213565b60200260200101516040518263ffffffff1660e01b8152600401610bcd91906001600160a01b0391909116815260200190565b602060405180830381865afa158015610bea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0e919061824b565b828281518110610c2057610c20618213565b6020908102919091010152610c348161827a565b9050610b70565b5092915050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610cae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd2919061824b565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610d3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d619190618293565b90506001600160c01b0381161580610dfb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df691906182bc565b60ff16155b15610e1757505060408051600081526020810190915292915050565b6000610e2b826001600160c01b0316615970565b90506000805b8251811015610f01577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610e7b57610e7b618213565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610ebf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee3919061824b565b610eed90836182d9565b915080610ef98161827a565b915050610e31565b506000816001600160401b03811115610f1c57610f1c616f5f565b604051908082528060200260200182016040528015610f45578160200160208202803683370190505b5090506000805b8451811015611104576000858281518110610f6957610f69618213565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610fde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611002919061824b565b905060005b818110156110ee576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561107c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a09190618301565b600001518686815181106110b6576110b6618213565b6001600160a01b0390921660209283029190910190910152846110d88161827a565b95505080806110e69061827a565b915050611007565b50505080806110fc9061827a565b915050610f4c565b5090979650505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611153573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111779190618147565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111dd9190618147565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561121f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112439190618147565b9050600086516001600160401b0381111561126057611260616f5f565b60405190808252806020026020018201604052801561129357816020015b606081526020019060019003908161127e5790505b50905060005b875181101561159b5760008882815181106112b6576112b6618213565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015611317573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261133f9190810190618342565b905080516001600160401b0381111561135a5761135a616f5f565b6040519080825280602002602001820160405280156113a557816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816113785790505b508484815181106113b8576113b8618213565b602002602001018190525060005b8151811015611585576040518060600160405280876001600160a01b03166347b314e88585815181106113fb576113fb618213565b60200260200101516040518263ffffffff1660e01b815260040161142191815260200190565b602060405180830381865afa15801561143e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114629190618147565b6001600160a01b0316815260200183838151811061148257611482618213565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106114b0576114b0618213565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa15801561150c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153091906183d2565b6001600160601b031681525085858151811061154e5761154e618213565b6020026020010151828151811061156757611567618213565b6020026020010181905250808061157d9061827a565b9150506113c6565b50505080806115939061827a565b915050611299565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611605573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116299190618147565b6001600160a01b0316336001600160a01b0316146116d55760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610848565b61085a81615a32565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0354606090600080516020618e8a833981519152906000906001600160401b0381111561172d5761172d616f5f565b60405190808252806020026020018201604052801561177a57816020015b6040805160608082018352808252600060208301529181019190915281526020019060019003908161174b5790505b50905060005b6003830154811015610c3b576117b48360030182815481106117a4576117a4618213565b9060005260206000200154615a8b565b8282815181106117c6576117c6618213565b602002602001018190525080806117dc9061827a565b915050611780565b60006117f360208501856183ef565b63ffffffff811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f066020526040902054909150600080516020618e8a833981519152906118915760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610848565b84846040516020016118a492919061853b565b60408051601f19818403018152918152815160209283012063ffffffff8516600090815260068501909352912054146119455760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610848565b63ffffffff8216600090815260078201602052604090205460ff16156119df5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610848565b6127106119ef60208601866183ef565b6119f99190618579565b63ffffffff164363ffffffff161115611a7a5760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610848565b611a8760208601866183ef565b63ffffffff16611a9a60208801886183ef565b63ffffffff1614611ae05760405162461bcd60e51b815260206004820152601060248201526f0a8c2e6d640928840dad2e6dac2e8c6d60831b6044820152606401610848565b600083516001600160401b03811115611afb57611afb616f5f565b604051908082528060200260200182016040528015611b24578160200160208202803683370190505b50905060005b8451811015611b9657611b67858281518110611b4857611b48618213565b6020026020010151805160009081526020918201519091526040902090565b828281518110611b7957611b79618213565b602090810291909101015280611b8e8161827a565b915050611b2a565b506000611ba96040890160208a016183ef565b82604051602001611bbb929190618596565b60405160208183030381529060405280519060200120905085602001358114611c655760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610848565b600085516001600160401b03811115611c8057611c80616f5f565b604051908082528060200260200182016040528015611ca9578160200160208202803683370190505b50905060005b8651811015611d9c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110611cf957611cf9618213565b60200260200101516040518263ffffffff1660e01b8152600401611d1f91815260200190565b602060405180830381865afa158015611d3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d609190618147565b828281518110611d7257611d72618213565b6001600160a01b039092166020928302919091019091015280611d948161827a565b915050611caf565b5063ffffffff85166000818152600786016020908152604091829020805460ff19166001179055815192835233908301527fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec910160405180910390a1505050505050505050565b606081516001600160401b03811115611e1e57611e1e616f5f565b604051908082528060200260200182016040528015611e47578160200160208202803683370190505b50905060005b8251811015610c3b57836001600160a01b031663296bb064848381518110611e7757611e77618213565b60200260200101516040518263ffffffff1660e01b8152600401611e9d91815260200190565b602060405180830381865afa158015611eba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ede9190618147565b828281518110611ef057611ef0618213565b6001600160a01b0390921660209283029190910190910152611f118161827a565b9050611e4d565b611f436040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fa79190618147565b9050611fd46040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90612004908b90899089906004016185d1565b600060405180830381865afa158015612021573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120499190810190618618565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061207b908b908b908b906004016186a6565b600060405180830381865afa158015612098573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120c09190810190618618565b6040820152856001600160401b038111156120dd576120dd616f5f565b60405190808252806020026020018201604052801561211057816020015b60608152602001906001900390816120fb5790505b50606082015260005b60ff8116871115612551576000856001600160401b0381111561213e5761213e616f5f565b604051908082528060200260200182016040528015612167578160200160208202803683370190505b5083606001518360ff168151811061218157612181618213565b602002602001018190525060005b868110156124515760008c6001600160a01b03166304ec63518a8a858181106121ba576121ba618213565b905060200201358e886000015186815181106121d8576121d8618213565b60200260200101516040518463ffffffff1660e01b81526004016122159392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612232573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122569190618293565b9050806001600160c01b03166000036122fd5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610848565b8a8a8560ff1681811061231257612312618213565b60016001600160c01b038516919093013560f81c1c8216909103905061243e57856001600160a01b031663dd9846b98a8a8581811061235357612353618213565b905060200201358d8d8860ff1681811061236f5761236f618213565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156123c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e991906186c6565b85606001518560ff168151811061240257612402618213565b6020026020010151848151811061241b5761241b618213565b63ffffffff909216602092830291909101909101528261243a8161827a565b9350505b50806124498161827a565b91505061218f565b506000816001600160401b0381111561246c5761246c616f5f565b604051908082528060200260200182016040528015612495578160200160208202803683370190505b50905060005b828110156125165784606001518460ff16815181106124bc576124bc618213565b602002602001015181815181106124d5576124d5618213565b60200260200101518282815181106124ef576124ef618213565b63ffffffff909216602092830291909101909101528061250e8161827a565b91505061249b565b508084606001518460ff168151811061253157612531618213565b602002602001018190525050508080612549906186e3565b915050612119565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015612592573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125b69190618147565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906125e9908b908b908e90600401618702565b600060405180830381865afa158015612606573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261262e9190810190618618565b60208301525098975050505050505050565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015612688573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ac91906181ae565b6126c85760405162461bcd60e51b8152600401610848906181cb565b600019600281905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161273992919061872c565b600060405180830381865afa158015612756573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261277e9190810190618618565b9050600084516001600160401b0381111561279b5761279b616f5f565b6040519080825280602002602001820160405280156127c4578160200160208202803683370190505b50905060005b85518110156128c557866001600160a01b03166304ec63518783815181106127f4576127f4618213565b60200260200101518786858151811061280f5761280f618213565b60200260200101516040518463ffffffff1660e01b815260040161284c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612869573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061288d9190618293565b6001600160c01b03168282815181106128a8576128a8618213565b6020908102919091010152806128bd8161827a565b9150506127ca565b5095945050505050565b604080518082019091526060808252602082015260008481036129485760405162461bcd60e51b81526020600482015260376024820152600080516020618eaa83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610848565b60408301515185148015612960575060a08301515185145b8015612970575060c08301515185145b8015612980575060e08301515185145b6129ea5760405162461bcd60e51b81526020600482015260416024820152600080516020618eaa83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610848565b82515160208401515114612a625760405162461bcd60e51b815260206004820152604460248201819052600080516020618eaa833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610848565b4363ffffffff168463ffffffff1610612ad15760405162461bcd60e51b815260206004820152603c6024820152600080516020618eaa83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610848565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115612b1257612b12616f5f565b604051908082528060200260200182016040528015612b3b578160200160208202803683370190505b506020820152866001600160401b03811115612b5957612b59616f5f565b604051908082528060200260200182016040528015612b82578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115612bb657612bb6616f5f565b604051908082528060200260200182016040528015612bdf578160200160208202803683370190505b5081526020860151516001600160401b03811115612bff57612bff616f5f565b604051908082528060200260200182016040528015612c28578160200160208202803683370190505b5081602001819052506000612cfa8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612cd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cf591906182bc565b615bd4565b905060005b876020015151811015612f7657612d2588602001518281518110611b4857611b48618213565b83602001518281518110612d3b57612d3b618213565b60209081029190910101528015612dfb576020830151612d5c60018361874b565b81518110612d6c57612d6c618213565b602002602001015160001c83602001518281518110612d8d57612d8d618213565b602002602001015160001c11612dfb576040805162461bcd60e51b8152602060048201526024810191909152600080516020618eaa83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610848565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612e4057612e40618213565b60200260200101518b8b600001518581518110612e5f57612e5f618213565b60200260200101516040518463ffffffff1660e01b8152600401612e9c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612eb9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612edd9190618293565b6001600160c01b031683600001518281518110612efc57612efc618213565b602002602001018181525050612f62610ab7612f368486600001518581518110612f2857612f28618213565b602002602001015116615c5e565b8a602001518481518110612f4c57612f4c618213565b6020026020010151615c8990919063ffffffff16565b945080612f6e8161827a565b915050612cff565b5050612f8183615d6c565b60335490935060ff16600081612f9857600061301a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ff6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061301a919061824b565b905060005b8a81101561369857821561317a578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061307657613076618213565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156130b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130da919061824b565b6130e491906182d9565b1161317a5760405162461bcd60e51b81526020600482015260666024820152600080516020618eaa83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610848565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106131bb576131bb618213565b9050013560f81c60f81b60f81c8c8c60a0015185815181106131df576131df618213565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561323b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061325f919061875e565b6001600160401b0319166132828a604001518381518110611b4857611b48618213565b67ffffffffffffffff19161461331e5760405162461bcd60e51b81526020600482015260616024820152600080516020618eaa83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610848565b61334e8960400151828151811061333757613337618213565b60200260200101518761556890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061339157613391618213565b9050013560f81c60f81b60f81c8c8c60c0015185815181106133b5576133b5618213565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015613411573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061343591906183d2565b8560200151828151811061344b5761344b618213565b6001600160601b0390921660209283029190910182015285015180518290811061347757613477618213565b60200260200101518560000151828151811061349557613495618213565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156136835761350d866000015182815181106134df576134df618213565b60200260200101518f8f868181106134f9576134f9618213565b600192013560f81c9290921c811614919050565b15613671577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061355357613553618213565b9050013560f81c60f81b60f81c8e8960200151858151811061357757613577618213565b60200260200101518f60e00151888151811061359557613595618213565b602002602001015187815181106135ae576135ae618213565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015613612573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061363691906183d2565b875180518590811061364a5761364a618213565b6020026020010181815161365e9190618789565b6001600160601b03169052506001909101905b8061367b8161827a565b9150506134b9565b505080806136909061827a565b91505061301f565b5050506000806136b28c868a606001518b6080015161099c565b91509150816137235760405162461bcd60e51b81526020600482015260436024820152600080516020618eaa83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610848565b806137845760405162461bcd60e51b81526020600482015260396024820152600080516020618eaa83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610848565b5050600087826020015160405160200161379f929190618596565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6137d1615e07565b6137db6000615e61565b565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa15801561383f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138639190618147565b604051638902624560e01b81526000600482015263ffffffff431660248201526001600160a01b039190911690638902624590604401600060405180830381865afa1580156138b6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526138de9190810190618342565b9050600081516001600160401b038111156138fb576138fb616f5f565b604051908082528060200260200182016040528015613924578160200160208202803683370190505b50905060005b8251811015610c3b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663296bb06484838151811061397457613974618213565b60200260200101516040518263ffffffff1660e01b815260040161399a91815260200190565b602060405180830381865afa1580156139b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139db9190618147565b8282815181106139ed576139ed618213565b6001600160a01b039092166020928302919091019091015280613a0f8161827a565b91505061392a565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614613a5f5760405162461bcd60e51b8152600401610848906187a9565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90613aad9085908590600401618821565b600060405180830381600087803b158015613ac757600080fd5b505af1158015613adb573d6000803e3d6000fd5b505050505050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0954600080516020618e8a833981519152906001600160a01b03163314613b765760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610848565b60408051608081018252606081830181905263ffffffff8881168352438116602080850191909152908816918301919091528251601f860182900482028101820190935284835290919085908590819084018382808284376000920182905250604086019490945250613beb9150615a799050565b905081604051602001613bfe91906188b7565b60408051601f19818403018152828252805160209182012063ffffffff8b16600090815260058601909252919020557ff31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca590613c5c90899085906188ca565b60405180910390a160048101805467ffffffff00000000191664010000000063ffffffff8a1602179055613c8e615066565b50505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614613cdf5760405162461bcd60e51b8152600401610848906187a9565b613ce881615eb3565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015613d4c57600080fd5b505af1158015613d60573d6000803e3d6000fd5b5050505050565b613d6f615e07565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590613d329084906004016188e9565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0854600080516020618e8a833981519152906001600160a01b03163314613e445760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610848565b613e4d82616189565b5050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0854600080516020618e8a833981519152906001600160a01b03163314613eda5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610848565b6000613eec60408601602087016183ef565b9050366000613efe60408801886188fc565b90925090506000613f156080890160608a016183ef565b9050600080516020618e8a8339815191527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f056000613f5660208c018c6183ef565b63ffffffff1663ffffffff1681526020019081526020016000205489604051602001613f829190618942565b604051602081830303815290604052805190602001201461400b5760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610848565b6000600682018161401f60208d018d6183ef565b63ffffffff1663ffffffff168152602001908152602001600020541461409c5760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610848565b60048101546140b19063ffffffff1686618579565b63ffffffff164363ffffffff1611156141225760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610848565b60008860405160200161413591906189be565b60405160208183030381529060405280519060200120905060008061415d8388888b8e6128cf565b9150915060005b8681101561425c578560ff168360200151828151811061418657614186618213565b602002602001015161419891906189d1565b6001600160601b03166043846000015183815181106141b9576141b9618213565b60200260200101516001600160601b03166141d491906189f4565b101561424a576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610848565b806142548161827a565b915050614164565b50600061426c60208d018d618a0b565b905011156142925761429261428460208d018d618a0b565b61428d91618a54565b616189565b60408051808201825263ffffffff431681526020808201849052915190916142be918e91849101618a61565b604051602081830303815290604052805190602001208560060160008f60000160208101906142ed91906183ef565b63ffffffff1663ffffffff168152602001908152602001600020819055507fb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce35678c8260405161433c929190618a61565b60405180910390a150505050505050505050505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061438e5761438e618213565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906143ca908890869060040161872c565b600060405180830381865afa1580156143e7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261440f9190810190618618565b60008151811061442157614421618213565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561448d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144b19190618293565b6001600160c01b0316905060006144c782615970565b9050816144d58a838a611111565b9550955050505050935093915050565b6040805160608082018352808252600060208301529181019190915261450a82615a8b565b92915050565b600080826040516020016145249190618a94565b6040516020818303038152906040528051906020012090506000614553600080516020618e8a83398151915290565b9050600081600001600084815260200190815260200160002060405180606001604052908160008201805461458790618ab0565b80601f01602080910402602001604051908101604052809291908181526020018280546145b390618ab0565b80156146005780601f106145d557610100808354040283529160200191614600565b820191906000526020600020905b8154815290600101906020018083116145e357829003601f168201915b50505050508152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561466c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161464e575b50505050508152505090506000805b8260400151518110156128c5576146b0836040015182815181106146a1576146a1618213565b602002602001015160006146ce565b6146ba9083618ae4565b9150806146c68161827a565b91505061467b565b604051631619718360e21b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000821691635401ed27917f000000000000000000000000000000000000000000000000000000000000000090911690635865c60c906024016040805180830381865afa158015614762573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147869190618b04565b5160405160e083901b6001600160e01b0319168152600481019190915260ff85166024820152604401602060405180830381865afa1580156147cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147f091906183d2565b9392505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614859573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061487d91906182bc565b60ff1690508060000361489e57505060408051600081526020810190915290565b6000805b8281101561495357604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015614911573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614935919061824b565b61493f90836182d9565b91508061494b8161827a565b9150506148a2565b506000816001600160401b0381111561496e5761496e616f5f565b604051908082528060200260200182016040528015614997578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156149fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a2091906182bc565b60ff16811015614bb957604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015614a94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ab8919061824b565b905060005b81811015614ba4576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015614b32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b569190618301565b60000151858581518110614b6c57614b6c618213565b6001600160a01b039092166020928302919091019091015283614b8e8161827a565b9450508080614b9c9061827a565b915050614abd565b50508080614bb19061827a565b91505061499e565b5090949350505050565b614bcb615e07565b6001600160a01b038116614c305760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610848565b61085a81615e61565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614614c815760405162461bcd60e51b8152600401610848906187a9565b614c8a82616297565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90614cd89086908590600401618821565b600060405180830381600087803b158015614cf257600080fd5b505af1158015614d06573d6000803e3d6000fd5b50505050600082604051602001614d1d9190618a94565b604051602081830303815290604052805190602001209050614d3f84826162e2565b50505050565b606554610100900460ff1615808015614d655750606554600160ff909116105b80614d7f5750303b158015614d7f575060655460ff166001145b614de25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610848565b6065805460ff191660011790558015614e05576065805461ff0019166101001790555b614e1086600061650f565b614e19856165f5565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0880546001600160a01b03199081166001600160a01b03878116919091179092557fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0980549091169185169190911790557fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f04805463ffffffff191663ffffffff84161790558015613adb576065805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614f5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614f819190618147565b6001600160a01b0316336001600160a01b031614614fb15760405162461bcd60e51b815260040161084890618164565b60025419811960025419161461502f5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610848565b600281905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610991565b60006150706137dd565b60405162cf2ab560e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062cf2ab590613d3290849060040161725e565b60005b81811015615391578282828181106150da576150da618213565b90506020028101906150ec9190618b34565b6150fd906040810190602001616f29565b6001600160a01b03166323b872dd333086868681811061511f5761511f618213565b90506020028101906151319190618b34565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af1158015615188573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151ac91906181ae565b5060008383838181106151c1576151c1618213565b90506020028101906151d39190618b34565b6151e4906040810190602001616f29565b604051636eb1769f60e11b81523060048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166024830152919091169063dd62ed3e90604401602060405180830381865afa158015615252573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615276919061824b565b905083838381811061528a5761528a618213565b905060200281019061529c9190618b34565b6152ad906040810190602001616f29565b6001600160a01b031663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000838787878181106152ef576152ef618213565b90506020028101906153019190618b34565b6040013561530f91906182d9565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561535a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061537e91906181ae565b50508061538a9061827a565b90506150c0565b5060405163fce36c7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063fce36c7d90613aad9085908590600401618bb0565b6001600160a01b03811661546e5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610848565b600154604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526154f3616d68565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808061552257fe5b50806155605760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610848565b505092915050565b6040805180820190915260008082526020820152615584616d86565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080806155bf57fe5b50806155605760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610848565b615605616da4565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806156ed600080516020618e6a83398151915286618229565b90505b6156f981616625565b9093509150600080516020618e6a8339815191528283098303615732576040805180820190915290815260208101919091529392505050565b600080516020618e6a8339815191526001820890506156f0565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061577e616dc9565b60005b60028110156159435760006157978260066189f4565b90508482600281106157ab576157ab618213565b602002015151836157bd8360006182d9565b600c81106157cd576157cd618213565b60200201528482600281106157e4576157e4618213565b602002015160200151838260016157fb91906182d9565b600c811061580b5761580b618213565b602002015283826002811061582257615822618213565b60200201515151836158358360026182d9565b600c811061584557615845618213565b602002015283826002811061585c5761585c618213565b60200201515160016020020151836158758360036182d9565b600c811061588557615885618213565b602002015283826002811061589c5761589c618213565b6020020151602001516000600281106158b7576158b7618213565b6020020151836158c88360046182d9565b600c81106158d8576158d8618213565b60200201528382600281106158ef576158ef618213565b60200201516020015160016002811061590a5761590a618213565b60200201518361591b8360056182d9565b600c811061592b5761592b618213565b6020020152508061593b8161827a565b915050615781565b5061594c616de8565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b606060008061597e84615c5e565b61ffff166001600160401b0381111561599957615999616f5f565b6040519080825280601f01601f1916602001820160405280156159c3576020820181803683370190505b5090506000805b8251821080156159db575061010081105b15614bb9576001811b935085841615615a22578060f81b838381518110615a0457615a04618213565b60200101906001600160f81b031916908160001a9053508160010191505b615a2b8161827a565b90506159ca565b6033805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b600080516020618e8a83398151915290565b604080516060808201835280825260006020808401829052838501839052858252600080516020618e8a8339815191529081905290849020845192830190945283549293909282908290615ade90618ab0565b80601f0160208091040260200160405190810160405280929190818152602001828054615b0a90618ab0565b8015615b575780601f10615b2c57610100808354040283529160200191615b57565b820191906000526020600020905b815481529060010190602001808311615b3a57829003601f168201915b505050505081526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015615bc357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311615ba5575b505050505081525050915050919050565b600080615be0846166a7565b9050808360ff166001901b116147f05760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610848565b6000805b821561450a57615c7360018461874b565b9092169180615c8181618cbe565b915050615c62565b60408051808201909152600080825260208201526102008261ffff1610615ce55760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610848565b8161ffff16600103615cf857508161450a565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610615d6157600161ffff871660ff83161c81169003615d4457615d418484615568565b93505b615d4e8384615568565b92506201fffe600192831b169101615d14565b509195945050505050565b60408051808201909152600080825260208201528151158015615d9157506020820151155b15615daf575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020618e6a8339815191528460200151615de29190618229565b615dfa90600080516020618e6a83398151915261874b565b905292915050565b919050565b6098546001600160a01b031633146137db5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610848565b609880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038116615f015760405162461bcd60e51b81526020600482015260156024820152747a65726f206f70657261746f72206164647265737360581b6044820152606401610848565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f016020526040902054600080516020618e8a8339815191529080615f8f5760405162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840ecc2d8d2c8c2e8dee440d0c2e6d60531b6044820152606401610848565b6001600160a01b038316600090815260028084016020908152604080842054858552918690529092209081015482106160005760405162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f6620626f756e647360681b6044820152606401610848565b846001600160a01b031681600201838154811061601f5761601f618213565b6000918252602090912001546001600160a01b0316146160815760405162461bcd60e51b815260206004820152601960248201527f4f70657261746f722061646472657373206d69736d61746368000000000000006044820152606401610848565b60028101546000906160959060019061874b565b6002830180549192506000916160ad9060019061874b565b815481106160bd576160bd618213565b6000918252602090912001546001600160a01b0316905083821461612057808360020185815481106160f1576160f1618213565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b8260020180548061613357616133618cdf565b600082815260208120820160001990810180546001600160a01b031916905590910190915560405186917fec3ecc6d7c19c6fc03c8b840fabed1a1b70c73f74bdb2d71f18c2f5a23639d3291a250505050505050565b600081516001600160401b038111156161a4576161a4616f5f565b6040519080825280602002602001820160405280156161cd578160200160208202803683370190505b50905060005b825181101561625b5760008382815181106161f0576161f0618213565b60200260200101516040516020016162089190618a94565b6040516020818303038152906040528051906020012090508083838151811061623357616233618213565b60200260200101818152505061624881616837565b50806162538161827a565b9150506161d3565b507fb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a8160405161628b9190617212565b60405180910390a15050565b6000816040516020016162aa9190618a94565b6040516020818303038152906040528051906020012090506162cb81616937565b1515600003613e4d576162dd8261694c565b505050565b6001600160a01b0382166163305760405162461bcd60e51b81526020600482015260156024820152747a65726f206f70657261746f72206164647265737360581b6044820152606401610848565b8061637d5760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642076616c696461746f722061646472657373000000000000006044820152606401610848565b61638681616937565b6163ca5760405162461bcd60e51b81526020600482015260156024820152740eadcd6deeedc40ecc2d8d2c8c2e8dee440d0c2e6d605b1b6044820152606401610848565b60006163d78360006146ce565b90506000816001600160601b0316116164225760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964207374616b6560981b6044820152606401610848565b6001600160a01b03831660008181527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0160209081526040808320869055858352600080516020618e8a83398151915280835281842060020180548686527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0285528386208190558285526001810182559085529383902090930180546001600160a01b0319168517905580516001600160601b038616815290519293869390927f5ea87e0bfaffd2b7f2901f087737e2b175f237068c6123ca9ea5240076bb5668928290030190a350505050565b6001546001600160a01b031615801561653057506001600160a01b03821615155b6165b25760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610848565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613e4d826153e0565b606554610100900460ff1661661c5760405162461bcd60e51b815260040161084890618cf5565b61085a81616ac5565b60008080600080516020618e6a8339815191526003600080516020618e6a83398151915286600080516020618e6a83398151915288890909089050600061669b827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020618e6a833981519152616aec565b91959194509092505050565b6000610100825111156167305760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610848565b815160000361674157506000919050565b6000808360008151811061675757616757618213565b0160200151600160f89190911c81901b92505b845181101561682e5784818151811061678557616785618213565b0160200151600160f89190911c1b915082821161681a5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610848565b918117916168278161827a565b905061676a565b50909392505050565b6000818152600080516020618e8a833981519152602081905260409091206001810154156162dd5760005b6002820154811015616905576168a082600201828154811061688657616886618213565b6000918252602090912001546001600160a01b0316616b95565b838260020182815481106168b6576168b6618213565b60009182526020822001546040516001600160a01b03909116917f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a3806168fd8161827a565b915050616862565b5060008381526020839052604081209061691f8282616e06565b6001820160009055600282016000613d609190616e40565b600061694282615a8b565b5151151592915050565b6040805160608082018352808252600060208301529181019190915260008260405160200161697b9190618a94565b60405160208183030381529060405280519060200120905061699c81616937565b156169e95760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f7220616c72656164792072656769737465726564000000006044820152606401610848565b604080516060810182528481526020808201849052825160008082529181018452909282015290506000600080516020618e8a8339815191526000848152602082905260409020835191925083918190616a439082618d86565b50602082810151600183015560408301518051616a669260028501920190616e5e565b505050600381018054600181018255600091825260209091200183905560405183907f1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea390616ab59088906188e9565b60405180910390a2509392505050565b606554610100900460ff16614c305760405162461bcd60e51b815260040161084890618cf5565b600080616af7616de8565b616aff616ec3565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280616b3c57fe5b5082616b8a5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610848565b505195945050505050565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f016020526040902054600080516020618e8a8339815191529080156162dd5760408051600180825281830190925260009160208083019080368337019050509050600081600081518110616c1a57616c1a618213565b602002602001019060ff16908160ff16815250507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e3b17db85616c6784616cbc565b6040518363ffffffff1660e01b8152600401616c84929190618e45565b600060405180830381600087803b158015616c9e57600080fd5b505af1158015616cb2573d6000803e3d6000fd5b5050505050505050565b6060600082516001600160401b03811115616cd957616cd9616f5f565b6040519080825280601f01601f191660200182016040528015616d03576020820181803683370190505b50905060005b8351811015610c3b57838181518110616d2457616d24618213565b602002602001015160f81b828281518110616d4157616d41618213565b60200101906001600160f81b031916908160001a905350616d618161827a565b9050616d09565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280616db7616ee1565b8152602001616dc4616ee1565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b508054616e1290618ab0565b6000825580601f10616e22575050565b601f01602090049060005260206000209081019061085a9190616eff565b508054600082559060005260206000209081019061085a9190616eff565b828054828255906000526020600020908101928215616eb3579160200282015b82811115616eb357825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190616e7e565b50616ebf929150616eff565b5090565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b5b80821115616ebf5760008155600101616f00565b6001600160a01b038116811461085a57600080fd5b600060208284031215616f3b57600080fd5b81356147f081616f14565b600060208284031215616f5857600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715616f9757616f97616f5f565b60405290565b60405161010081016001600160401b0381118282101715616f9757616f97616f5f565b604051601f8201601f191681016001600160401b0381118282101715616fe857616fe8616f5f565b604052919050565b60006040828403121561700257600080fd5b61700a616f75565b9050813581526020820135602082015292915050565b600082601f83011261703157600080fd5b617039616f75565b80604084018581111561704b57600080fd5b845b8181101561706557803584526020938401930161704d565b509095945050505050565b60006080828403121561708257600080fd5b61708a616f75565b90506170968383617020565b81526170a58360408401617020565b602082015292915050565b60008060008061012085870312156170c757600080fd5b843593506170d88660208701616ff0565b92506170e78660608701617070565b91506170f68660e08701616ff0565b905092959194509250565b60006001600160401b0382111561711a5761711a616f5f565b5060051b60200190565b6000806040838503121561713757600080fd5b823561714281616f14565b91506020838101356001600160401b0381111561715e57600080fd5b8401601f8101861361716f57600080fd5b803561718261717d82617101565b616fc0565b81815260059190911b820183019083810190888311156171a157600080fd5b928401925b828410156171c85783356171b981616f14565b825292840192908401906171a6565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015617207578151875295820195908201906001016171eb565b509495945050505050565b6020815260006147f060208301846171d7565b600081518084526020808501945080840160005b838110156172075781516001600160a01b031687529582019590820190600101617239565b6020815260006147f06020830184617225565b600082601f83011261728257600080fd5b81356001600160401b0381111561729b5761729b616f5f565b6172ae601f8201601f1916602001616fc0565b8181528460208386010111156172c357600080fd5b816020850160208301376000918101602001919091529392505050565b63ffffffff8116811461085a57600080fd5b8035615e02816172e0565b60008060006060848603121561731257600080fd5b833561731d81616f14565b925060208401356001600160401b0381111561733857600080fd5b61734486828701617271565b9250506040840135617355816172e0565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156173f6578385038a52825180518087529087019087870190845b818110156173e157835180516001600160a01b031684528a8101518b8501526040908101516001600160601b0316908401529289019260609092019160010161739d565b50509a87019a9550509185019160010161737f565b509298975050505050505050565b6020815260006147f06020830184617360565b801515811461085a57600080fd5b60006020828403121561743757600080fd5b81356147f081617417565b60005b8381101561745d578181015183820152602001617445565b50506000910152565b6000815180845261747e816020860160208601617442565b601f01601f19169290920160200192915050565b60008151606084526174a76060850182617466565b905060208301516020850152604083015184820360408601526174ca8282617225565b95945050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561752857603f19888603018452617516858351617492565b945092850192908501906001016174fa565b5092979650505050505050565b60006080828403121561754757600080fd5b50919050565b60006040828403121561754757600080fd5b600082601f83011261757057600080fd5b8135602061758061717d83617101565b82815260069290921b8401810191818101908684111561759f57600080fd5b8286015b848110156175c3576175b58882616ff0565b8352918301916040016175a3565b509695505050505050565b60008060008060a085870312156175e457600080fd5b84356001600160401b03808211156175fb57600080fd5b61760788838901617535565b9550602087013591508082111561761d57600080fd5b6176298883890161754d565b9450617638886040890161754d565b9350608087013591508082111561764e57600080fd5b5061765b8782880161755f565b91505092959194509250565b600082601f83011261767857600080fd5b8135602061768861717d83617101565b82815260059290921b840181019181810190868411156176a757600080fd5b8286015b848110156175c357803583529183019183016176ab565b600080604083850312156176d557600080fd5b82356176e081616f14565b915060208301356001600160401b038111156176fb57600080fd5b61770785828601617667565b9150509250929050565b60008083601f84011261772357600080fd5b5081356001600160401b0381111561773a57600080fd5b60208301915083602082850101111561775257600080fd5b9250929050565b60008083601f84011261776b57600080fd5b5081356001600160401b0381111561778257600080fd5b6020830191508360208260051b850101111561775257600080fd5b600080600080600080608087890312156177b657600080fd5b86356177c181616f14565b955060208701356177d1816172e0565b945060408701356001600160401b03808211156177ed57600080fd5b6177f98a838b01617711565b9096509450606089013591508082111561781257600080fd5b5061781f89828a01617759565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561720757815163ffffffff1687529582019590820190600101617845565b60006020808352835160808285015261788360a0850182617831565b905081850151601f19808684030160408701526178a08383617831565b925060408701519150808684030160608701526178bd8383617831565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156179145784878303018452617902828751617831565b958801959388019391506001016178e8565b509998505050505050505050565b60ff8116811461085a57600080fd5b60006020828403121561794357600080fd5b81356147f081617922565b60008060006060848603121561796357600080fd5b833561796e81616f14565b925060208401356001600160401b0381111561798957600080fd5b61734486828701617667565b6020808252825182820181905260009190848201906040850190845b818110156179cd578351835292840192918401916001016179b1565b50909695505050505050565b600082601f8301126179ea57600080fd5b813560206179fa61717d83617101565b82815260059290921b84018101918181019086841115617a1957600080fd5b8286015b848110156175c3578035617a30816172e0565b8352918301918301617a1d565b600082601f830112617a4e57600080fd5b81356020617a5e61717d83617101565b82815260059290921b84018101918181019086841115617a7d57600080fd5b8286015b848110156175c35780356001600160401b03811115617aa05760008081fd5b617aae8986838b01016179d9565b845250918301918301617a81565b60006101808284031215617acf57600080fd5b617ad7616f9d565b905081356001600160401b0380821115617af057600080fd5b617afc858386016179d9565b83526020840135915080821115617b1257600080fd5b617b1e8583860161755f565b60208401526040840135915080821115617b3757600080fd5b617b438583860161755f565b6040840152617b558560608601617070565b6060840152617b678560e08601616ff0565b6080840152610120840135915080821115617b8157600080fd5b617b8d858386016179d9565b60a0840152610140840135915080821115617ba757600080fd5b617bb3858386016179d9565b60c0840152610160840135915080821115617bcd57600080fd5b50617bda84828501617a3d565b60e08301525092915050565b600080600080600060808688031215617bfe57600080fd5b8535945060208601356001600160401b0380821115617c1c57600080fd5b617c2889838a01617711565b909650945060408801359150617c3d826172e0565b90925060608701359080821115617c5357600080fd5b50617c6088828901617abc565b9150509295509295909350565b600081518084526020808501945080840160005b838110156172075781516001600160601b031687529582019590820190600101617c81565b6040815260008351604080840152617cc16080840182617c6d565b90506020850151603f19848303016060850152617cde8282617c6d565b925050508260208301529392505050565b600060608284031215617d0157600080fd5b604051606081016001600160401b038282108183111715617d2457617d24616f5f565b816040528293508435915080821115617d3c57600080fd5b50617d4985828601617271565b82525060208301356020820152604083013560408201525092915050565b60008060408385031215617d7a57600080fd5b8235617d8581616f14565b915060208301356001600160401b03811115617da057600080fd5b61770785828601617cef565b60008060008060608587031215617dc257600080fd5b8435617dcd816172e0565b93506020850135617ddd816172e0565b925060408501356001600160401b03811115617df857600080fd5b617e0487828801617711565b95989497509550505050565b600060208284031215617e2257600080fd5b81356001600160401b03811115617e3857600080fd5b617e4484828501617271565b949350505050565b6000617e5a61717d84617101565b8381529050602080820190600585901b840186811115617e7957600080fd5b845b81811015617eb45780356001600160401b03811115617e9a5760008081fd5b617ea689828901617271565b855250928201928201617e7b565b505050509392505050565b600060208284031215617ed157600080fd5b81356001600160401b03811115617ee757600080fd5b8201601f81018413617ef857600080fd5b617e4484823560208401617e4c565b600080600060608486031215617f1c57600080fd5b83356001600160401b0380821115617f3357600080fd5b617f3f87838801617535565b94506020860135915080821115617f5557600080fd5b617f618783880161754d565b93506040860135915080821115617f7757600080fd5b50617f8486828701617abc565b9150509250925092565b600080600060608486031215617fa357600080fd5b8335617fae81616f14565b9250602084013591506040840135617355816172e0565b828152604060208201526000617e446040830184617360565b6020815260006147f06020830184617492565b6000806040838503121561800457600080fd5b823561800f81616f14565b9150602083013561801f81617922565b809150509250929050565b60008060006060848603121561803f57600080fd5b833561804a81616f14565b925060208401356001600160401b038082111561806657600080fd5b61807287838801617271565b9350604086013591508082111561808857600080fd5b50617f8486828701617cef565b600080600080600060a086880312156180ad57600080fd5b85356180b881616f14565b945060208601356180c881616f14565b935060408601356180d881616f14565b925060608601356180e881616f14565b915060808601356180f8816172e0565b809150509295509295909350565b6000806020838503121561811957600080fd5b82356001600160401b0381111561812f57600080fd5b61813b85828601617759565b90969095509350505050565b60006020828403121561815957600080fd5b81516147f081616f14565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156181c057600080fd5b81516147f081617417565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261824657634e487b7160e01b600052601260045260246000fd5b500690565b60006020828403121561825d57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60006001820161828c5761828c618264565b5060010190565b6000602082840312156182a557600080fd5b81516001600160c01b03811681146147f057600080fd5b6000602082840312156182ce57600080fd5b81516147f081617922565b8082018082111561450a5761450a618264565b6001600160601b038116811461085a57600080fd5b60006040828403121561831357600080fd5b61831b616f75565b825161832681616f14565b81526020830151618336816182ec565b60208201529392505050565b6000602080838503121561835557600080fd5b82516001600160401b0381111561836b57600080fd5b8301601f8101851361837c57600080fd5b805161838a61717d82617101565b81815260059190911b820183019083810190878311156183a957600080fd5b928401925b828410156183c7578351825292840192908401906183ae565b979650505050505050565b6000602082840312156183e457600080fd5b81516147f0816182ec565b60006020828403121561840157600080fd5b81356147f0816172e0565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e1984360301811261844c57600080fd5b83016020810192503590506001600160401b0381111561846b57600080fd5b80360382131561775257600080fd5b600060408301823561848b816172e0565b63ffffffff16845260208381013536859003601e190181126184ac57600080fd5b840181810190356001600160401b038111156184c757600080fd5b8060051b8036038313156184da57600080fd5b6040848901529381905260609387018401938290880160005b8381101561852d57898703605f1901825261850e8386618435565b61851989828461840c565b9850505091850191908501906001016184f3565b509498975050505050505050565b60608152600061854e606083018561847a565b9050823561855b816172e0565b63ffffffff8116602084015250602083013560408301529392505050565b63ffffffff818116838216019080821115610c3b57610c3b618264565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015617528578151855293820193908201906001016185b5565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156185fe57600080fd5b8260051b8085606085013791909101606001949350505050565b6000602080838503121561862b57600080fd5b82516001600160401b0381111561864157600080fd5b8301601f8101851361865257600080fd5b805161866061717d82617101565b81815260059190911b8201830190838101908783111561867f57600080fd5b928401925b828410156183c7578351618697816172e0565b82529284019290840190618684565b63ffffffff841681526040602082015260006174ca60408301848661840c565b6000602082840312156186d857600080fd5b81516147f0816172e0565b600060ff821660ff81036186f9576186f9618264565b60010192915050565b60408152600061871660408301858761840c565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201526000617e4460408301846171d7565b8181038181111561450a5761450a618264565b60006020828403121561877057600080fd5b815167ffffffffffffffff19811681146147f057600080fd5b6001600160601b03828116828216039080821115610c3b57610c3b618264565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261884b60a0840182617466565b90506020840151606084015260408401516080840152809150509392505050565b600063ffffffff80835116845280602084015116602085015260408301516080604086015261889e6080860182617466565b9050816060850151166060860152809250505092915050565b6020815260006147f0602083018461886c565b63ffffffff83168152604060208201526000617e44604083018461886c565b6020815260006147f06020830184617466565b6000808335601e1984360301811261891357600080fd5b8301803591506001600160401b0382111561892d57600080fd5b60200191503681900382131561775257600080fd5b6020815260008235618953816172e0565b63ffffffff808216602085015260208501359150618970826172e0565b80821660408501526189856040860186618435565b92506080606086015261899c60a08601848361840c565b92505060608501356189ad816172e0565b166080939093019290925250919050565b6020815260006147f0602083018461847a565b6001600160601b0381811683821602808216919082811461556057615560618264565b808202811582820484141761450a5761450a618264565b6000808335601e19843603018112618a2257600080fd5b8301803591506001600160401b03821115618a3c57600080fd5b6020019150600581901b360382131561775257600080fd5b60006147f0368484617e4c565b606081526000618a74606083018561847a565b905063ffffffff8351166020830152602083015160408301529392505050565b60008251618aa6818460208701617442565b9190910192915050565b600181811c90821680618ac457607f821691505b60208210810361754757634e487b7160e01b600052602260045260246000fd5b6001600160601b03818116838216019080821115610c3b57610c3b618264565b600060408284031215618b1657600080fd5b618b1e616f75565b8251815260208301516003811061833657600080fd5b60008235609e19833603018112618aa657600080fd5b8035615e0281616f14565b8183526000602080850194508260005b85811015617207578135618b7881616f14565b6001600160a01b0316875281830135618b90816182ec565b6001600160601b0316878401526040968701969190910190600101618b65565b60208082528181018390526000906040808401600586901b8501820187855b88811015618cb057878303603f190184528135368b9003609e19018112618bf557600080fd5b8a0160a0813536839003601e19018112618c0e57600080fd5b820188810190356001600160401b03811115618c2957600080fd5b8060061b3603821315618c3b57600080fd5b828752618c4b8388018284618b55565b92505050618c5a888301618b4a565b6001600160a01b03168886015281870135878601526060618c7c8184016172f2565b63ffffffff16908601526080618c938382016172f2565b63ffffffff16950194909452509285019290850190600101618bcf565b509098975050505050505050565b600061ffff808316818103618cd557618cd5618264565b6001019392505050565b634e487b7160e01b600052603160045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b601f8211156162dd57600081815260208120601f850160051c81016020861015618d675750805b601f850160051c820191505b81811015613adb57828155600101618d73565b81516001600160401b03811115618d9f57618d9f616f5f565b618db381618dad8454618ab0565b84618d40565b602080601f831160018114618de85760008415618dd05750858301515b600019600386901b1c1916600185901b178555613adb565b600085815260208120601f198616915b82811015618e1757888601518255948401946001909101908401618df8565b5085821015618e355787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b0383168152604060208201819052600090617e449083018461746656fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47d4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220be1dc8b0f5814ff8b18476952ce58706f831d66e20a0d450e0aaac5d859afebb64736f6c63430008150033",
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

// Initialize is a paid mutator transaction binding the contract method 0xf8a2334c.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _aggregator, address _generator, uint32 _taskResponseWindowBlock) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, _initialOwner common.Address, _aggregator common.Address, _generator common.Address, _taskResponseWindowBlock uint32) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "initialize", _pauserRegistry, _initialOwner, _aggregator, _generator, _taskResponseWindowBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8a2334c.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _aggregator, address _generator, uint32 _taskResponseWindowBlock) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Initialize(_pauserRegistry common.Address, _initialOwner common.Address, _aggregator common.Address, _generator common.Address, _taskResponseWindowBlock uint32) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Initialize(&_ContractZrServiceManager.TransactOpts, _pauserRegistry, _initialOwner, _aggregator, _generator, _taskResponseWindowBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8a2334c.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _aggregator, address _generator, uint32 _taskResponseWindowBlock) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) Initialize(_pauserRegistry common.Address, _initialOwner common.Address, _aggregator common.Address, _generator common.Address, _taskResponseWindowBlock uint32) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Initialize(&_ContractZrServiceManager.TransactOpts, _pauserRegistry, _initialOwner, _aggregator, _generator, _taskResponseWindowBlock)
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
