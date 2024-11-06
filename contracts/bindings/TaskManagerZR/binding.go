// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTaskManagerZR

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

// ITaskManagerZRTask is an auto generated low-level Go binding around an user-defined struct.
type ITaskManagerZRTask struct {
	TaskId                    uint32
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// ITaskManagerZRTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type ITaskManagerZRTaskResponse struct {
	ReferenceTaskId  uint32
	ActiveSetZRChain []string
}

// ITaskManagerZRTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type ITaskManagerZRTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
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

// ContractTaskManagerZRMetaData contains all meta data concerning the ContractTaskManagerZR contract.
var ContractTaskManagerZRMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveSet\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestActiveSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidatorInTaskSet\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"validatorAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskValidatorAddressesByIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskValidatorCount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAddressesStored\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620066cc380380620066cc8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e051610100516163d5620002f7600039600081816102b40152818161064e01526133a6015260008181610617015261273d0152600081816104a50152818161158a015261291f0152600081816104cc01528181612af50152612cb70152600081816104f301528181610f7401528181612427015281816125a001526127da01526163d56000f3fe608060405234801561001057600080fd5b50600436106102535760003560e01c8063683048351161014657806390a54d14116100c3578063df5cf72311610087578063df5cf72314610612578063f2fde38b14610639578063f5c9899d1461064c578063f63c5bab14610672578063f8c8765e1461067b578063fabc1cbc1461068e57600080fd5b806390a54d141461059e5780639d3a0f2d146105be578063b98d0908146105d1578063caf73aa0146105de578063cefdc1d4146105f157600080fd5b8063767917c61161010a578063767917c61461054c57806379e2c40f146105545780637afa1eed14610567578063886f11951461057a5780638da5cb5b1461058d57600080fd5b806368304835146104c75780636d14a987146104ee5780636efb463614610515578063715018a61461053657806372d18e8d1461053e57600080fd5b80633ec54dcf116101d45780635ac86ab7116101985780635ac86ab7146104225780635c155662146104555780635c975abb146104755780635decc3f51461047d5780635df45946146104a057600080fd5b80633ec54dcf146103b4578063416c7e5e146103d45780634c71c165146103e75780634f739f74146103fa578063595c6a671461041a57600080fd5b806324ad60c51161021b57806324ad60c5146103165780632cb223d5146103365780632d3de569146103645780632d89f6fc146103745780633563b0d11461039457600080fd5b806310d67a2f14610258578063136439dd1461026d578063171f1d5b146102805780631ad43189146102af578063245a7bfc146102eb575b600080fd5b61026b610266366004614b97565b6106a1565b005b61026b61027b366004614bb4565b61075d565b61029361028e366004614d32565b61089c565b6040805192151583529015156020830152015b60405180910390f35b6102d67f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102a6565b60cf546102fe906001600160a01b031681565b6040516001600160a01b0390911681526020016102a6565b610329610324366004614da0565b610a26565b6040516102a69190614e28565b610356610344366004614e3b565b60cb6020526000908152604090205481565b6040519081526020016102a6565b60c9546102d69063ffffffff1681565b610356610382366004614e3b565b60ca6020526000908152604090205481565b6103a76103a2366004614eaf565b610acb565b6040516102a69190614fca565b6103c76103c2366004614e3b565b610f61565b6040516102a69190614fdd565b61026b6103e236600461504d565b610f72565b61026b6103f536600461512b565b6110e7565b61040d61040836600461520c565b6116bf565b6040516102a69190615310565b61026b611de5565b6104456104303660046153da565b606654600160ff9092169190911b9081161490565b60405190151581526020016102a6565b6104686104633660046153f7565b611eac565b6040516102a691906154ac565b606654610356565b61044561048b366004614e3b565b60ce6020526000908152604090205460ff1681565b6102fe7f000000000000000000000000000000000000000000000000000000000000000081565b6102fe7f000000000000000000000000000000000000000000000000000000000000000081565b6102fe7f000000000000000000000000000000000000000000000000000000000000000081565b6105286105233660046156fd565b612074565b6040516102a69291906157bd565b61026b612f6c565b60c95463ffffffff166102d6565b6103c7612f80565b610445610562366004615806565b612f9a565b60d0546102fe906001600160a01b031681565b6065546102fe906001600160a01b031681565b6033546001600160a01b03166102fe565b6103566105ac366004614e3b565b60cd6020526000908152604090205481565b61026b6105cc366004615869565b613050565b6097546104459060ff1681565b61026b6105ec3660046158cd565b6131b5565b6106046105ff366004615954565b613656565b6040516102a692919061598b565b6102fe7f000000000000000000000000000000000000000000000000000000000000000081565b61026b610647366004614b97565b6137e8565b7f00000000000000000000000000000000000000000000000000000000000000006102d6565b6102d661271081565b61026b6106893660046159ac565b61385e565b61026b61069c366004614bb4565b6139af565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107189190615a08565b6001600160a01b0316336001600160a01b0316146107515760405162461bcd60e51b815260040161074890615a25565b60405180910390fd5b61075a81613b0b565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156107a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c99190615a6f565b6107e55760405162461bcd60e51b815260040161074890615a8c565b6066548181161461085e5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610748565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106108e4576108e4615ad4565b60200201518951600160200201518a6020015160006002811061090957610909615ad4565b60200201518b6020015160016002811061092557610925615ad4565b602090810291909101518c518d8301516040516109829a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6109a59190615aea565b9050610a186109be6109b78884613c02565b8690613c99565b6109c6613d2d565b610a0e6109ff856109f9604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613c02565b610a088c613ded565b90613c99565b886201d4c0613e7d565b909890975095505050505050565b60cc60209081526000928352604080842090915290825290208054610a4a90615b0c565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7690615b0c565b8015610ac35780601f10610a9857610100808354040283529160200191610ac3565b820191906000526020600020905b815481529060010190602001808311610aa657829003601f168201915b505050505081565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b319190615a08565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b979190615a08565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bfd9190615a08565b9050600086516001600160401b03811115610c1a57610c1a614bcd565b604051908082528060200260200182016040528015610c4d57816020015b6060815260200190600190039081610c385790505b50905060005b8751811015610f55576000888281518110610c7057610c70615ad4565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610cd1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610cf99190810190615b41565b905080516001600160401b03811115610d1457610d14614bcd565b604051908082528060200260200182016040528015610d5f57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610d325790505b50848481518110610d7257610d72615ad4565b602002602001018190525060005b8151811015610f3f576040518060600160405280876001600160a01b03166347b314e8858581518110610db557610db5615ad4565b60200260200101516040518263ffffffff1660e01b8152600401610ddb91815260200190565b602060405180830381865afa158015610df8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1c9190615a08565b6001600160a01b03168152602001838381518110610e3c57610e3c615ad4565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610e6a57610e6a615ad4565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610ec6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eea9190615bd1565b6001600160601b0316815250858581518110610f0857610f08615ad4565b60200260200101518281518110610f2157610f21615ad4565b60200260200101819052508080610f3790615c10565b915050610d80565b5050508080610f4d90615c10565b915050610c53565b50979650505050505050565b6060610f6c826140a1565b92915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff49190615a08565b6001600160a01b0316336001600160a01b0316146110a05760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610748565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60006110f66020850185614e3b565b63ffffffff8116600090815260cb60205260409020549091506111655760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610748565b8383604051602001611178929190615d5a565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb909352912054146112175760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610748565b63ffffffff8116600090815260ce602052604090205460ff16156112af5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610748565b6127106112bf6020850185614e3b565b6112c99190615d98565b63ffffffff164363ffffffff16111561134a5760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610748565b6113576020850185614e3b565b63ffffffff1661136a6020870187614e3b565b63ffffffff16146113b05760405162461bcd60e51b815260206004820152601060248201526f0a8c2e6d640928840dad2e6dac2e8c6d60831b6044820152606401610748565b600082516001600160401b038111156113cb576113cb614bcd565b6040519080825280602002602001820160405280156113f4578160200160208202803683370190505b50905060005b83518110156114665761143784828151811061141857611418615ad4565b6020026020010151805160009081526020918201519091526040902090565b82828151811061144957611449615ad4565b60209081029190910101528061145e81615c10565b9150506113fa565b5060006114796040880160208901614e3b565b8260405160200161148b929190615dc0565b604051602081830303815290604052805190602001209050846020013581146115355760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610748565b600084516001600160401b0381111561155057611550614bcd565b604051908082528060200260200182016040528015611579578160200160208202803683370190505b50905060005b855181101561166c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106115c9576115c9615ad4565b60200260200101516040518263ffffffff1660e01b81526004016115ef91815260200190565b602060405180830381865afa15801561160c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116309190615a08565b82828151811061164257611642615ad4565b6001600160a01b03909216602092830291909101909101528061166481615c10565b91505061157f565b5063ffffffff8416600081815260ce6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505050565b6116ea6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561172a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061174e9190615a08565b905061177b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906117ab908b9089908990600401615dfb565b600060405180830381865afa1580156117c8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117f09190810190615e45565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611822908b908b908b90600401615ed3565b600060405180830381865afa15801561183f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118679190810190615e45565b6040820152856001600160401b0381111561188457611884614bcd565b6040519080825280602002602001820160405280156118b757816020015b60608152602001906001900390816118a25790505b50606082015260005b60ff8116871115611cf6576000856001600160401b038111156118e5576118e5614bcd565b60405190808252806020026020018201604052801561190e578160200160208202803683370190505b5083606001518360ff168151811061192857611928615ad4565b602002602001018190525060005b86811015611bf65760008c6001600160a01b03166304ec63518a8a8581811061196157611961615ad4565b905060200201358e8860000151868151811061197f5761197f615ad4565b60200260200101516040518463ffffffff1660e01b81526004016119bc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156119d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119fd9190615efc565b90506001600160c01b038116611aa15760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610748565b8a8a8560ff16818110611ab657611ab6615ad4565b6001600160c01b03841692013560f81c9190911c600190811614159050611be357856001600160a01b031663dd9846b98a8a85818110611af857611af8615ad4565b905060200201358d8d8860ff16818110611b1457611b14615ad4565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611b6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8e9190615f25565b85606001518560ff1681518110611ba757611ba7615ad4565b60200260200101518481518110611bc057611bc0615ad4565b63ffffffff9092166020928302919091019091015282611bdf81615c10565b9350505b5080611bee81615c10565b915050611936565b506000816001600160401b03811115611c1157611c11614bcd565b604051908082528060200260200182016040528015611c3a578160200160208202803683370190505b50905060005b82811015611cbb5784606001518460ff1681518110611c6157611c61615ad4565b60200260200101518181518110611c7a57611c7a615ad4565b6020026020010151828281518110611c9457611c94615ad4565b63ffffffff9092166020928302919091019091015280611cb381615c10565b915050611c40565b508084606001518460ff1681518110611cd657611cd6615ad4565b602002602001018190525050508080611cee90615f42565b9150506118c0565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d5b9190615a08565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611d8e908b908b908e90600401615f62565b600060405180830381865afa158015611dab573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611dd39190810190615e45565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611e2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e519190615a6f565b611e6d5760405162461bcd60e51b815260040161074890615a8c565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611ede929190615f8c565b600060405180830381865afa158015611efb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f239190810190615e45565b9050600084516001600160401b03811115611f4057611f40614bcd565b604051908082528060200260200182016040528015611f69578160200160208202803683370190505b50905060005b855181101561206a57866001600160a01b03166304ec6351878381518110611f9957611f99615ad4565b602002602001015187868581518110611fb457611fb4615ad4565b60200260200101516040518463ffffffff1660e01b8152600401611ff19392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561200e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120329190615efc565b6001600160c01b031682828151811061204d5761204d615ad4565b60209081029190910101528061206281615c10565b915050611f6f565b5095945050505050565b60408051808201909152606080825260208201526000846120eb5760405162461bcd60e51b8152602060048201526037602482015260008051602061638083398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610748565b60408301515185148015612103575060a08301515185145b8015612113575060c08301515185145b8015612123575060e08301515185145b61218d5760405162461bcd60e51b8152602060048201526041602482015260008051602061638083398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610748565b825151602084015151146122055760405162461bcd60e51b815260206004820152604460248201819052600080516020616380833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610748565b4363ffffffff168463ffffffff16106122745760405162461bcd60e51b815260206004820152603c602482015260008051602061638083398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610748565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156122b5576122b5614bcd565b6040519080825280602002602001820160405280156122de578160200160208202803683370190505b506020820152866001600160401b038111156122fc576122fc614bcd565b604051908082528060200260200182016040528015612325578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561235957612359614bcd565b604051908082528060200260200182016040528015612382578160200160208202803683370190505b5081526020860151516001600160401b038111156123a2576123a2614bcd565b6040519080825280602002602001820160405280156123cb578160200160208202803683370190505b508160200181905250600061249d8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612474573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124989190615fe0565b6141f4565b905060005b876020015151811015612719576124c88860200151828151811061141857611418615ad4565b836020015182815181106124de576124de615ad4565b6020908102919091010152801561259e5760208301516124ff600183615ffd565b8151811061250f5761250f615ad4565b602002602001015160001c8360200151828151811061253057612530615ad4565b602002602001015160001c1161259e576040805162461bcd60e51b815260206004820152602481019190915260008051602061638083398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610748565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106125e3576125e3615ad4565b60200260200101518b8b60000151858151811061260257612602615ad4565b60200260200101516040518463ffffffff1660e01b815260040161263f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561265c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126809190615efc565b6001600160c01b03168360000151828151811061269f5761269f615ad4565b6020026020010181815250506127056109b76126d984866000015185815181106126cb576126cb615ad4565b602002602001015116614285565b8a6020015184815181106126ef576126ef615ad4565b60200260200101516142b090919063ffffffff16565b94508061271181615c10565b9150506124a2565b505061272483614394565b60975490935060ff1660008161273b5760006127bd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612799573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127bd9190616014565b905060005b8a811015612e3b57821561291d578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061281957612819615ad4565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612859573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061287d9190616014565b612887919061602d565b1161291d5760405162461bcd60e51b8152602060048201526066602482015260008051602061638083398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610748565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061295e5761295e615ad4565b9050013560f81c60f81b60f81c8c8c60a00151858151811061298257612982615ad4565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156129de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a029190616045565b6001600160401b031916612a258a60400151838151811061141857611418615ad4565b67ffffffffffffffff191614612ac15760405162461bcd60e51b8152602060048201526061602482015260008051602061638083398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610748565b612af189604001518281518110612ada57612ada615ad4565b602002602001015187613c9990919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612b3457612b34615ad4565b9050013560f81c60f81b60f81c8c8c60c001518581518110612b5857612b58615ad4565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612bb4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bd89190615bd1565b85602001518281518110612bee57612bee615ad4565b6001600160601b03909216602092830291909101820152850151805182908110612c1a57612c1a615ad4565b602002602001015185600001518281518110612c3857612c38615ad4565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612e2657612cb086600001518281518110612c8257612c82615ad4565b60200260200101518f8f86818110612c9c57612c9c615ad4565b600192013560f81c9290921c811614919050565b15612e14577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612cf657612cf6615ad4565b9050013560f81c60f81b60f81c8e89602001518581518110612d1a57612d1a615ad4565b60200260200101518f60e001518881518110612d3857612d38615ad4565b60200260200101518781518110612d5157612d51615ad4565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612db5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dd99190615bd1565b8751805185908110612ded57612ded615ad4565b60200260200101818151612e019190616070565b6001600160601b03169052506001909101905b80612e1e81615c10565b915050612c5c565b50508080612e3390615c10565b9150506127c2565b505050600080612e558c868a606001518b6080015161089c565b9150915081612ec65760405162461bcd60e51b8152602060048201526043602482015260008051602061638083398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610748565b80612f275760405162461bcd60e51b8152602060048201526039602482015260008051602061638083398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610748565b50506000878260200151604051602001612f42929190615dc0565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612f7461442f565b612f7e6000614489565b565b60c954606090612f959063ffffffff166140a1565b905090565b63ffffffff8216600090815260cd6020526040812054815b818110156130455783604051602001612fcb9190616098565b60408051601f19818403018152828252805160209182012063ffffffff8916600090815260cc8352838120868252835292909220919261300c9291016160b4565b60405160208183030381529060405280519060200120141561303357600192505050610f6c565b8061303d81615c10565b915050612fb2565b506000949350505050565b60d0546001600160a01b031633146130b45760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610748565b60408051608081018252606081830181905263ffffffff8781168352438116602080850191909152908716918301919091528251601f850182900482028101820190935283835290919084908490819084018382808284376000920191909152505050506040808301919091525161313090829060200161614c565b60408051601f19818403018152828252805160209182012063ffffffff8916600081815260ca9093529290912055907ff31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca59061318c90849061614c565b60405180910390a2505060c9805463ffffffff191663ffffffff94909416939093179092555050565b60cf546001600160a01b0316331461320f5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610748565b60006132216040850160208601614e3b565b9050366000613233604087018761619e565b9092509050600061324a6080880160608901614e3b565b905060ca600061325d60208a018a614e3b565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161328991906161e4565b60405160208183030381529060405280519060200120146133125760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610748565b600060cb8161332460208b018b614e3b565b63ffffffff1663ffffffff16815260200190815260200160002054146133a15760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610748565b6133cb7f000000000000000000000000000000000000000000000000000000000000000085615d98565b63ffffffff164363ffffffff16111561343c5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610748565b60008660405160200161344f9190616260565b6040516020818303038152906040528051906020012090506000806134778387878a8c612074565b9150915060005b85811015613576578460ff16836020015182815181106134a0576134a0615ad4565b60200260200101516134b29190616273565b6001600160601b03166043846000015183815181106134d3576134d3615ad4565b60200260200101516001600160601b03166134ee91906162a2565b1015613564576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610748565b8061356e81615c10565b91505061347e565b5061359961358760208c018c614e3b565b61359460208c018c6162c1565b6144db565b60408051808201825263ffffffff431681526020808201849052915190916135c5918c9184910161630a565b6040516020818303038152906040528051906020012060cb60008d60000160208101906135f29190614e3b565b63ffffffff1663ffffffff168152602001908152602001600020819055507fb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce35678a8260405161364192919061630a565b60405180910390a15050505050505050505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061369157613691615ad4565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906136cd9088908690600401615f8c565b600060405180830381865afa1580156136ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526137129190810190615e45565b60008151811061372457613724615ad4565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613790573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137b49190615efc565b6001600160c01b0316905060006137ca826145a2565b9050816137d88a838a610acb565b9550955050505050935093915050565b6137f061442f565b6001600160a01b0381166138555760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610748565b61075a81614489565b600054610100900460ff161580801561387e5750600054600160ff909116105b806138985750303b158015613898575060005460ff166001145b6138fb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610748565b6000805460ff19166001179055801561391e576000805461ff0019166101001790555b61392985600061466e565b61393284614489565b60cf80546001600160a01b038086166001600160a01b03199283161790925560d080549285169290911691909117905580156139a8576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a269190615a08565b6001600160a01b0316336001600160a01b031614613a565760405162461bcd60e51b815260040161074890615a25565b606654198119606654191614613ad45760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610748565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610891565b6001600160a01b038116613b995760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610748565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613c1e614a0f565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613c5157613c53565bfe5b5080613c915760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610748565b505092915050565b6040805180820190915260008082526020820152613cb5614a2d565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613c51575080613c915760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610748565b613d35614a4b565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613e1d60008051602061636083398151915286615aea565b90505b613e2981614758565b9093509150600080516020616360833981519152828309831415613e63576040805180820190915290815260208101919091529392505050565b600080516020616360833981519152600182089050613e20565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613eaf614a70565b60005b6002811015614074576000613ec88260066162a2565b9050848260028110613edc57613edc615ad4565b60200201515183613eee83600061602d565b600c8110613efe57613efe615ad4565b6020020152848260028110613f1557613f15615ad4565b60200201516020015183826001613f2c919061602d565b600c8110613f3c57613f3c615ad4565b6020020152838260028110613f5357613f53615ad4565b6020020151515183613f6683600261602d565b600c8110613f7657613f76615ad4565b6020020152838260028110613f8d57613f8d615ad4565b6020020151516001602002015183613fa683600361602d565b600c8110613fb657613fb6615ad4565b6020020152838260028110613fcd57613fcd615ad4565b602002015160200151600060028110613fe857613fe8615ad4565b602002015183613ff983600461602d565b600c811061400957614009615ad4565b602002015283826002811061402057614020615ad4565b60200201516020015160016002811061403b5761403b615ad4565b60200201518361404c83600561602d565b600c811061405c5761405c615ad4565b6020020152508061406c81615c10565b915050613eb2565b5061407d614a8f565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b63ffffffff8116600090815260cd6020526040812054606091816001600160401b038111156140d2576140d2614bcd565b60405190808252806020026020018201604052801561410557816020015b60608152602001906001900390816140f05790505b50905060005b828110156141ec5763ffffffff8516600090815260cc602090815260408083208484529091529020805461413e90615b0c565b80601f016020809104026020016040519081016040528092919081815260200182805461416a90615b0c565b80156141b75780601f1061418c576101008083540402835291602001916141b7565b820191906000526020600020905b81548152906001019060200180831161419a57829003601f168201915b50505050508282815181106141ce576141ce615ad4565b602002602001018190525080806141e490615c10565b91505061410b565b509392505050565b600080614200846147da565b9050808360ff166001901b1161427e5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610748565b9392505050565b6000805b8215610f6c5761429a600184615ffd565b90921691806142a88161633d565b915050614289565b60408051808201909152600080825260208201526102008261ffff161061430c5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610748565b8161ffff1660011415614320575081610f6c565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061438957600161ffff871660ff83161c8116141561436c576143698484613c99565b93505b6143768384613c99565b92506201fffe600192831b16910161433c565b509195945050505050565b604080518082019091526000808252602082015281511580156143b957506020820151155b156143d7575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020616360833981519152846020015161440a9190615aea565b61442290600080516020616360833981519152615ffd565b905292915050565b919050565b6033546001600160a01b03163314612f7e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610748565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b8060005b81811015614549578383828181106144f9576144f9615ad4565b905060200281019061450b919061619e565b63ffffffff8716600090815260cc602090815260408083208684529091529020614536929091614aad565b508061454181615c10565b9150506144df565b5063ffffffff8416600081815260cd602052604090819020839055517f5a04eb24e83275653ac5d71ef5205698d4ca126250977a126ab5b124d17f2c39906145949084815260200190565b60405180910390a250505050565b60606000806145b084614285565b61ffff166001600160401b038111156145cb576145cb614bcd565b6040519080825280601f01601f1916602001820160405280156145f5576020820181803683370190505b5090506000805b82518210801561460d575061010081105b15614664576001811b935085841615614654578060f81b83838151811061463657614636615ad4565b60200101906001600160f81b031916908160001a9053508160010191505b61465d81615c10565b90506145fc565b5090949350505050565b6065546001600160a01b031615801561468f57506001600160a01b03821615155b6147115760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610748565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261475482613b0b565b5050565b600080806000805160206163608339815191526003600080516020616360833981519152866000805160206163608339815191528889090908905060006147ce827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020616360833981519152614967565b91959194509092505050565b6000610100825111156148635760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610748565b815161487157506000919050565b6000808360008151811061488757614887615ad4565b0160200151600160f89190911c81901b92505b845181101561495e578481815181106148b5576148b5615ad4565b0160200151600160f89190911c1b915082821161494a5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610748565b9181179161495781615c10565b905061489a565b50909392505050565b600080614972614a8f565b61497a614b31565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613c51575082614a045760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610748565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280614a5e614b4f565b8152602001614a6b614b4f565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b828054614ab990615b0c565b90600052602060002090601f016020900481019282614adb5760008555614b21565b82601f10614af45782800160ff19823516178555614b21565b82800160010185558215614b21579182015b82811115614b21578235825591602001919060010190614b06565b50614b2d929150614b6d565b5090565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b5b80821115614b2d5760008155600101614b6e565b6001600160a01b038116811461075a57600080fd5b600060208284031215614ba957600080fd5b813561427e81614b82565b600060208284031215614bc657600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614c0557614c05614bcd565b60405290565b60405161010081016001600160401b0381118282101715614c0557614c05614bcd565b604051601f8201601f191681016001600160401b0381118282101715614c5657614c56614bcd565b604052919050565b600060408284031215614c7057600080fd5b614c78614be3565b9050813581526020820135602082015292915050565b600082601f830112614c9f57600080fd5b604051604081018181106001600160401b0382111715614cc157614cc1614bcd565b8060405250806040840185811115614cd857600080fd5b845b81811015614389578035835260209283019201614cda565b600060808284031215614d0457600080fd5b614d0c614be3565b9050614d188383614c8e565b8152614d278360408401614c8e565b602082015292915050565b6000806000806101208587031215614d4957600080fd5b84359350614d5a8660208701614c5e565b9250614d698660608701614cf2565b9150614d788660e08701614c5e565b905092959194509250565b63ffffffff8116811461075a57600080fd5b803561442a81614d83565b60008060408385031215614db357600080fd5b8235614dbe81614d83565b946020939093013593505050565b60005b83811015614de7578181015183820152602001614dcf565b83811115614df6576000848401525b50505050565b60008151808452614e14816020860160208601614dcc565b601f01601f19169290920160200192915050565b60208152600061427e6020830184614dfc565b600060208284031215614e4d57600080fd5b813561427e81614d83565b60006001600160401b03831115614e7157614e71614bcd565b614e84601f8401601f1916602001614c2e565b9050828152838383011115614e9857600080fd5b828260208301376000602084830101529392505050565b600080600060608486031215614ec457600080fd5b8335614ecf81614b82565b925060208401356001600160401b03811115614eea57600080fd5b8401601f81018613614efb57600080fd5b614f0a86823560208401614e58565b9250506040840135614f1b81614d83565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614fbc578385038a52825180518087529087019087870190845b81811015614fa757835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614f63565b50509a87019a95505091850191600101614f45565b509298975050505050505050565b60208152600061427e6020830184614f26565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561503257603f19888603018452615020858351614dfc565b94509285019290850190600101615004565b5092979650505050505050565b801515811461075a57600080fd5b60006020828403121561505f57600080fd5b813561427e8161503f565b60006080828403121561507c57600080fd5b50919050565b60006040828403121561507c57600080fd5b60006001600160401b038211156150ad576150ad614bcd565b5060051b60200190565b600082601f8301126150c857600080fd5b813560206150dd6150d883615094565b614c2e565b82815260069290921b840181019181810190868411156150fc57600080fd5b8286015b84811015615120576151128882614c5e565b835291830191604001615100565b509695505050505050565b60008060008060a0858703121561514157600080fd5b84356001600160401b038082111561515857600080fd5b6151648883890161506a565b9550602087013591508082111561517a57600080fd5b61518688838901615082565b94506151958860408901615082565b935060808701359150808211156151ab57600080fd5b506151b8878288016150b7565b91505092959194509250565b60008083601f8401126151d657600080fd5b5081356001600160401b038111156151ed57600080fd5b60208301915083602082850101111561520557600080fd5b9250929050565b6000806000806000806080878903121561522557600080fd5b863561523081614b82565b9550602087013561524081614d83565b945060408701356001600160401b038082111561525c57600080fd5b6152688a838b016151c4565b9096509450606089013591508082111561528157600080fd5b818901915089601f83011261529557600080fd5b8135818111156152a457600080fd5b8a60208260051b85010111156152b957600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561530557815163ffffffff16875295820195908201906001016152e3565b509495945050505050565b60006020808352835160808285015261532c60a08501826152cf565b905081850151601f198086840301604087015261534983836152cf565b9250604087015191508086840301606087015261536683836152cf565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156153bd57848783030184526153ab8287516152cf565b95880195938801939150600101615391565b509998505050505050505050565b60ff8116811461075a57600080fd5b6000602082840312156153ec57600080fd5b813561427e816153cb565b60008060006060848603121561540c57600080fd5b833561541781614b82565b92506020848101356001600160401b0381111561543357600080fd5b8501601f8101871361544457600080fd5b80356154526150d882615094565b81815260059190911b8201830190838101908983111561547157600080fd5b928401925b8284101561548f57833582529284019290840190615476565b80965050505050506154a360408501614d95565b90509250925092565b6020808252825182820181905260009190848201906040850190845b818110156154e4578351835292840192918401916001016154c8565b50909695505050505050565b600082601f83011261550157600080fd5b813560206155116150d883615094565b82815260059290921b8401810191818101908684111561553057600080fd5b8286015b8481101561512057803561554781614d83565b8352918301918301615534565b600082601f83011261556557600080fd5b813560206155756150d883615094565b82815260059290921b8401810191818101908684111561559457600080fd5b8286015b848110156151205780356001600160401b038111156155b75760008081fd5b6155c58986838b01016154f0565b845250918301918301615598565b600061018082840312156155e657600080fd5b6155ee614c0b565b905081356001600160401b038082111561560757600080fd5b615613858386016154f0565b8352602084013591508082111561562957600080fd5b615635858386016150b7565b6020840152604084013591508082111561564e57600080fd5b61565a858386016150b7565b604084015261566c8560608601614cf2565b606084015261567e8560e08601614c5e565b608084015261012084013591508082111561569857600080fd5b6156a4858386016154f0565b60a08401526101408401359150808211156156be57600080fd5b6156ca858386016154f0565b60c08401526101608401359150808211156156e457600080fd5b506156f184828501615554565b60e08301525092915050565b60008060008060006080868803121561571557600080fd5b8535945060208601356001600160401b038082111561573357600080fd5b61573f89838a016151c4565b90965094506040880135915061575482614d83565b9092506060870135908082111561576a57600080fd5b50615777888289016155d3565b9150509295509295909350565b600081518084526020808501945080840160005b838110156153055781516001600160601b031687529582019590820190600101615798565b60408152600083516040808401526157d86080840182615784565b90506020850151603f198483030160608501526157f58282615784565b925050508260208301529392505050565b6000806040838503121561581957600080fd5b823561582481614d83565b915060208301356001600160401b0381111561583f57600080fd5b8301601f8101851361585057600080fd5b61585f85823560208401614e58565b9150509250929050565b6000806000806060858703121561587f57600080fd5b843561588a81614d83565b9350602085013561589a81614d83565b925060408501356001600160401b038111156158b557600080fd5b6158c1878288016151c4565b95989497509550505050565b6000806000606084860312156158e257600080fd5b83356001600160401b03808211156158f957600080fd5b6159058783880161506a565b9450602086013591508082111561591b57600080fd5b61592787838801615082565b9350604086013591508082111561593d57600080fd5b5061594a868287016155d3565b9150509250925092565b60008060006060848603121561596957600080fd5b833561597481614b82565b9250602084013591506040840135614f1b81614d83565b8281526040602082015260006159a46040830184614f26565b949350505050565b600080600080608085870312156159c257600080fd5b84356159cd81614b82565b935060208501356159dd81614b82565b925060408501356159ed81614b82565b915060608501356159fd81614b82565b939692955090935050565b600060208284031215615a1a57600080fd5b815161427e81614b82565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215615a8157600080fd5b815161427e8161503f565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082615b0757634e487b7160e01b600052601260045260246000fd5b500690565b600181811c90821680615b2057607f821691505b6020821081141561507c57634e487b7160e01b600052602260045260246000fd5b60006020808385031215615b5457600080fd5b82516001600160401b03811115615b6a57600080fd5b8301601f81018513615b7b57600080fd5b8051615b896150d882615094565b81815260059190911b82018301908381019087831115615ba857600080fd5b928401925b82841015615bc657835182529284019290840190615bad565b979650505050505050565b600060208284031215615be357600080fd5b81516001600160601b038116811461427e57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415615c2457615c24615bfa565b5060010190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112615c6b57600080fd5b83016020810192503590506001600160401b03811115615c8a57600080fd5b80360383131561520557600080fd5b6000604083018235615caa81614d83565b63ffffffff16845260208381013536859003601e19018112615ccb57600080fd5b840181810190356001600160401b03811115615ce657600080fd5b8060051b803603871315615cf957600080fd5b6040848901529381905260609387018401938290880160005b83811015615d4c57898703605f19018252615d2d8386615c54565b615d38898284615c2b565b985050509185019190850190600101615d12565b509498975050505050505050565b606081526000615d6d6060830185615c99565b90508235615d7a81614d83565b63ffffffff8116602084015250602083013560408301529392505050565b600063ffffffff808316818516808303821115615db757615db7615bfa565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561503257815185529382019390820190600101615ddf565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115615e2857600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215615e5857600080fd5b82516001600160401b03811115615e6e57600080fd5b8301601f81018513615e7f57600080fd5b8051615e8d6150d882615094565b81815260059190911b82018301908381019087831115615eac57600080fd5b928401925b82841015615bc6578351615ec481614d83565b82529284019290840190615eb1565b63ffffffff84168152604060208201526000615ef3604083018486615c2b565b95945050505050565b600060208284031215615f0e57600080fd5b81516001600160c01b038116811461427e57600080fd5b600060208284031215615f3757600080fd5b815161427e81614d83565b600060ff821660ff811415615f5957615f59615bfa565b60010192915050565b604081526000615f76604083018587615c2b565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015615fd357845183529383019391830191600101615fb7565b5090979650505050505050565b600060208284031215615ff257600080fd5b815161427e816153cb565b60008282101561600f5761600f615bfa565b500390565b60006020828403121561602657600080fd5b5051919050565b6000821982111561604057616040615bfa565b500190565b60006020828403121561605757600080fd5b815167ffffffffffffffff198116811461427e57600080fd5b60006001600160601b038381169083168181101561609057616090615bfa565b039392505050565b600082516160aa818460208701614dcc565b9190910192915050565b600080835481600182811c9150808316806160d057607f831692505b60208084108214156160f057634e487b7160e01b86526022600452602486fd5b818015616104576001811461611557615d4c565b60ff19861689528489019650615d4c565b60008a81526020902060005b8681101561613a5781548b820152908501908301616121565b50505096909201979650505050505050565b60208152600063ffffffff80845116602084015280602085015116604084015260408401516080606085015261618560a0850182614dfc565b9050816060860151166080850152809250505092915050565b6000808335601e198436030181126161b557600080fd5b8301803591506001600160401b038211156161cf57600080fd5b60200191503681900382131561520557600080fd5b60208152600082356161f581614d83565b63ffffffff80821660208501526020850135915061621282614d83565b80821660408501526162276040860186615c54565b92506080606086015261623e60a086018483615c2b565b925050606085013561624f81614d83565b166080939093019290925250919050565b60208152600061427e6020830184615c99565b60006001600160601b038083168185168183048111821515161561629957616299615bfa565b02949350505050565b60008160001904831182151516156162bc576162bc615bfa565b500290565b6000808335601e198436030181126162d857600080fd5b8301803591506001600160401b038211156162f257600080fd5b6020019150600581901b360382131561520557600080fd5b60608152600061631d6060830185615c99565b905063ffffffff8351166020830152602083015160408301529392505050565b600061ffff8083168181141561635557616355615bfa565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212204734343f08fb489a7105a2fd71d7be0a29a94f87387a01c59f79d295018a516564736f6c634300080c0033",
}

// ContractTaskManagerZRABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTaskManagerZRMetaData.ABI instead.
var ContractTaskManagerZRABI = ContractTaskManagerZRMetaData.ABI

// ContractTaskManagerZRBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTaskManagerZRMetaData.Bin instead.
var ContractTaskManagerZRBin = ContractTaskManagerZRMetaData.Bin

// DeployContractTaskManagerZR deploys a new Ethereum contract, binding an instance of ContractTaskManagerZR to it.
func DeployContractTaskManagerZR(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractTaskManagerZR, error) {
	parsed, err := ContractTaskManagerZRMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTaskManagerZRBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTaskManagerZR{ContractTaskManagerZRCaller: ContractTaskManagerZRCaller{contract: contract}, ContractTaskManagerZRTransactor: ContractTaskManagerZRTransactor{contract: contract}, ContractTaskManagerZRFilterer: ContractTaskManagerZRFilterer{contract: contract}}, nil
}

// ContractTaskManagerZR is an auto generated Go binding around an Ethereum contract.
type ContractTaskManagerZR struct {
	ContractTaskManagerZRCaller     // Read-only binding to the contract
	ContractTaskManagerZRTransactor // Write-only binding to the contract
	ContractTaskManagerZRFilterer   // Log filterer for contract events
}

// ContractTaskManagerZRCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTaskManagerZRCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskManagerZRTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTaskManagerZRTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskManagerZRFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTaskManagerZRFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskManagerZRSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTaskManagerZRSession struct {
	Contract     *ContractTaskManagerZR // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractTaskManagerZRCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTaskManagerZRCallerSession struct {
	Contract *ContractTaskManagerZRCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractTaskManagerZRTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTaskManagerZRTransactorSession struct {
	Contract     *ContractTaskManagerZRTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractTaskManagerZRRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTaskManagerZRRaw struct {
	Contract *ContractTaskManagerZR // Generic contract binding to access the raw methods on
}

// ContractTaskManagerZRCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTaskManagerZRCallerRaw struct {
	Contract *ContractTaskManagerZRCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTaskManagerZRTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTaskManagerZRTransactorRaw struct {
	Contract *ContractTaskManagerZRTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTaskManagerZR creates a new instance of ContractTaskManagerZR, bound to a specific deployed contract.
func NewContractTaskManagerZR(address common.Address, backend bind.ContractBackend) (*ContractTaskManagerZR, error) {
	contract, err := bindContractTaskManagerZR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZR{ContractTaskManagerZRCaller: ContractTaskManagerZRCaller{contract: contract}, ContractTaskManagerZRTransactor: ContractTaskManagerZRTransactor{contract: contract}, ContractTaskManagerZRFilterer: ContractTaskManagerZRFilterer{contract: contract}}, nil
}

// NewContractTaskManagerZRCaller creates a new read-only instance of ContractTaskManagerZR, bound to a specific deployed contract.
func NewContractTaskManagerZRCaller(address common.Address, caller bind.ContractCaller) (*ContractTaskManagerZRCaller, error) {
	contract, err := bindContractTaskManagerZR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRCaller{contract: contract}, nil
}

// NewContractTaskManagerZRTransactor creates a new write-only instance of ContractTaskManagerZR, bound to a specific deployed contract.
func NewContractTaskManagerZRTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTaskManagerZRTransactor, error) {
	contract, err := bindContractTaskManagerZR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRTransactor{contract: contract}, nil
}

// NewContractTaskManagerZRFilterer creates a new log filterer instance of ContractTaskManagerZR, bound to a specific deployed contract.
func NewContractTaskManagerZRFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTaskManagerZRFilterer, error) {
	contract, err := bindContractTaskManagerZR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRFilterer{contract: contract}, nil
}

// bindContractTaskManagerZR binds a generic wrapper to an already deployed contract.
func bindContractTaskManagerZR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTaskManagerZRMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskManagerZR *ContractTaskManagerZRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskManagerZR.Contract.ContractTaskManagerZRCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskManagerZR *ContractTaskManagerZRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.ContractTaskManagerZRTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskManagerZR *ContractTaskManagerZRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.ContractTaskManagerZRTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskManagerZR.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractTaskManagerZR.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractTaskManagerZR.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractTaskManagerZR.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractTaskManagerZR.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractTaskManagerZR.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractTaskManagerZR.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractTaskManagerZR.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractTaskManagerZR.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Aggregator() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Aggregator(&_ContractTaskManagerZR.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) Aggregator() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Aggregator(&_ContractTaskManagerZR.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractTaskManagerZR.Contract.AllTaskHashes(&_ContractTaskManagerZR.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractTaskManagerZR.Contract.AllTaskHashes(&_ContractTaskManagerZR.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractTaskManagerZR.Contract.AllTaskResponses(&_ContractTaskManagerZR.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractTaskManagerZR.Contract.AllTaskResponses(&_ContractTaskManagerZR.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) BlsApkRegistry() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.BlsApkRegistry(&_ContractTaskManagerZR.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.BlsApkRegistry(&_ContractTaskManagerZR.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractTaskManagerZR.Contract.CheckSignatures(&_ContractTaskManagerZR.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractTaskManagerZR.Contract.CheckSignatures(&_ContractTaskManagerZR.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Delegation() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Delegation(&_ContractTaskManagerZR.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) Delegation() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Delegation(&_ContractTaskManagerZR.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Generator() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Generator(&_ContractTaskManagerZR.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) Generator() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Generator(&_ContractTaskManagerZR.CallOpts)
}

// GetActiveSet is a free data retrieval call binding the contract method 0x3ec54dcf.
//
// Solidity: function getActiveSet(uint32 taskId) view returns(string[])
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) GetActiveSet(opts *bind.CallOpts, taskId uint32) ([]string, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "getActiveSet", taskId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetActiveSet is a free data retrieval call binding the contract method 0x3ec54dcf.
//
// Solidity: function getActiveSet(uint32 taskId) view returns(string[])
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) GetActiveSet(taskId uint32) ([]string, error) {
	return _ContractTaskManagerZR.Contract.GetActiveSet(&_ContractTaskManagerZR.CallOpts, taskId)
}

// GetActiveSet is a free data retrieval call binding the contract method 0x3ec54dcf.
//
// Solidity: function getActiveSet(uint32 taskId) view returns(string[])
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) GetActiveSet(taskId uint32) ([]string, error) {
	return _ContractTaskManagerZR.Contract.GetActiveSet(&_ContractTaskManagerZR.CallOpts, taskId)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractTaskManagerZR.Contract.GetCheckSignaturesIndices(&_ContractTaskManagerZR.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractTaskManagerZR.Contract.GetCheckSignaturesIndices(&_ContractTaskManagerZR.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetLatestActiveSet is a free data retrieval call binding the contract method 0x767917c6.
//
// Solidity: function getLatestActiveSet() view returns(string[])
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) GetLatestActiveSet(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "getLatestActiveSet")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetLatestActiveSet is a free data retrieval call binding the contract method 0x767917c6.
//
// Solidity: function getLatestActiveSet() view returns(string[])
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) GetLatestActiveSet() ([]string, error) {
	return _ContractTaskManagerZR.Contract.GetLatestActiveSet(&_ContractTaskManagerZR.CallOpts)
}

// GetLatestActiveSet is a free data retrieval call binding the contract method 0x767917c6.
//
// Solidity: function getLatestActiveSet() view returns(string[])
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) GetLatestActiveSet() ([]string, error) {
	return _ContractTaskManagerZR.Contract.GetLatestActiveSet(&_ContractTaskManagerZR.CallOpts)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractTaskManagerZR.Contract.GetOperatorState(&_ContractTaskManagerZR.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractTaskManagerZR.Contract.GetOperatorState(&_ContractTaskManagerZR.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractTaskManagerZR.Contract.GetOperatorState0(&_ContractTaskManagerZR.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractTaskManagerZR.Contract.GetOperatorState0(&_ContractTaskManagerZR.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractTaskManagerZR.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractTaskManagerZR.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractTaskManagerZR.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractTaskManagerZR.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractTaskManagerZR.Contract.GetTaskResponseWindowBlock(&_ContractTaskManagerZR.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractTaskManagerZR.Contract.GetTaskResponseWindowBlock(&_ContractTaskManagerZR.CallOpts)
}

// IsValidatorInTaskSet is a free data retrieval call binding the contract method 0x79e2c40f.
//
// Solidity: function isValidatorInTaskSet(uint32 taskId, string validatorAddress) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) IsValidatorInTaskSet(opts *bind.CallOpts, taskId uint32, validatorAddress string) (bool, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "isValidatorInTaskSet", taskId, validatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorInTaskSet is a free data retrieval call binding the contract method 0x79e2c40f.
//
// Solidity: function isValidatorInTaskSet(uint32 taskId, string validatorAddress) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) IsValidatorInTaskSet(taskId uint32, validatorAddress string) (bool, error) {
	return _ContractTaskManagerZR.Contract.IsValidatorInTaskSet(&_ContractTaskManagerZR.CallOpts, taskId, validatorAddress)
}

// IsValidatorInTaskSet is a free data retrieval call binding the contract method 0x79e2c40f.
//
// Solidity: function isValidatorInTaskSet(uint32 taskId, string validatorAddress) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) IsValidatorInTaskSet(taskId uint32, validatorAddress string) (bool, error) {
	return _ContractTaskManagerZR.Contract.IsValidatorInTaskSet(&_ContractTaskManagerZR.CallOpts, taskId, validatorAddress)
}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) LatestTaskId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "latestTaskId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) LatestTaskId() (uint32, error) {
	return _ContractTaskManagerZR.Contract.LatestTaskId(&_ContractTaskManagerZR.CallOpts)
}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) LatestTaskId() (uint32, error) {
	return _ContractTaskManagerZR.Contract.LatestTaskId(&_ContractTaskManagerZR.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Owner() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Owner(&_ContractTaskManagerZR.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) Owner() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.Owner(&_ContractTaskManagerZR.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Paused(index uint8) (bool, error) {
	return _ContractTaskManagerZR.Contract.Paused(&_ContractTaskManagerZR.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) Paused(index uint8) (bool, error) {
	return _ContractTaskManagerZR.Contract.Paused(&_ContractTaskManagerZR.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Paused0() (*big.Int, error) {
	return _ContractTaskManagerZR.Contract.Paused0(&_ContractTaskManagerZR.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) Paused0() (*big.Int, error) {
	return _ContractTaskManagerZR.Contract.Paused0(&_ContractTaskManagerZR.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) PauserRegistry() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.PauserRegistry(&_ContractTaskManagerZR.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.PauserRegistry(&_ContractTaskManagerZR.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) RegistryCoordinator() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.RegistryCoordinator(&_ContractTaskManagerZR.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.RegistryCoordinator(&_ContractTaskManagerZR.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) StakeRegistry() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.StakeRegistry(&_ContractTaskManagerZR.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractTaskManagerZR.Contract.StakeRegistry(&_ContractTaskManagerZR.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) StaleStakesForbidden() (bool, error) {
	return _ContractTaskManagerZR.Contract.StaleStakesForbidden(&_ContractTaskManagerZR.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractTaskManagerZR.Contract.StaleStakesForbidden(&_ContractTaskManagerZR.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TaskNumber() (uint32, error) {
	return _ContractTaskManagerZR.Contract.TaskNumber(&_ContractTaskManagerZR.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TaskNumber() (uint32, error) {
	return _ContractTaskManagerZR.Contract.TaskNumber(&_ContractTaskManagerZR.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractTaskManagerZR.Contract.TaskSuccesfullyChallenged(&_ContractTaskManagerZR.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractTaskManagerZR.Contract.TaskSuccesfullyChallenged(&_ContractTaskManagerZR.CallOpts, arg0)
}

// TaskValidatorAddressesByIndex is a free data retrieval call binding the contract method 0x24ad60c5.
//
// Solidity: function taskValidatorAddressesByIndex(uint32 , uint256 ) view returns(string)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TaskValidatorAddressesByIndex(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "taskValidatorAddressesByIndex", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TaskValidatorAddressesByIndex is a free data retrieval call binding the contract method 0x24ad60c5.
//
// Solidity: function taskValidatorAddressesByIndex(uint32 , uint256 ) view returns(string)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TaskValidatorAddressesByIndex(arg0 uint32, arg1 *big.Int) (string, error) {
	return _ContractTaskManagerZR.Contract.TaskValidatorAddressesByIndex(&_ContractTaskManagerZR.CallOpts, arg0, arg1)
}

// TaskValidatorAddressesByIndex is a free data retrieval call binding the contract method 0x24ad60c5.
//
// Solidity: function taskValidatorAddressesByIndex(uint32 , uint256 ) view returns(string)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TaskValidatorAddressesByIndex(arg0 uint32, arg1 *big.Int) (string, error) {
	return _ContractTaskManagerZR.Contract.TaskValidatorAddressesByIndex(&_ContractTaskManagerZR.CallOpts, arg0, arg1)
}

// TaskValidatorCount is a free data retrieval call binding the contract method 0x90a54d14.
//
// Solidity: function taskValidatorCount(uint32 ) view returns(uint256)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TaskValidatorCount(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "taskValidatorCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskValidatorCount is a free data retrieval call binding the contract method 0x90a54d14.
//
// Solidity: function taskValidatorCount(uint32 ) view returns(uint256)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TaskValidatorCount(arg0 uint32) (*big.Int, error) {
	return _ContractTaskManagerZR.Contract.TaskValidatorCount(&_ContractTaskManagerZR.CallOpts, arg0)
}

// TaskValidatorCount is a free data retrieval call binding the contract method 0x90a54d14.
//
// Solidity: function taskValidatorCount(uint32 ) view returns(uint256)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TaskValidatorCount(arg0 uint32) (*big.Int, error) {
	return _ContractTaskManagerZR.Contract.TaskValidatorCount(&_ContractTaskManagerZR.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractTaskManagerZR.Contract.TrySignatureAndApkVerification(&_ContractTaskManagerZR.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractTaskManagerZR.Contract.TrySignatureAndApkVerification(&_ContractTaskManagerZR.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) CreateNewTask(opts *bind.TransactOpts, taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "createNewTask", taskId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) CreateNewTask(taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.CreateNewTask(&_ContractTaskManagerZR.TransactOpts, taskId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) CreateNewTask(taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.CreateNewTask(&_ContractTaskManagerZR.TransactOpts, taskId, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.Initialize(&_ContractTaskManagerZR.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.Initialize(&_ContractTaskManagerZR.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.Pause(&_ContractTaskManagerZR.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.Pause(&_ContractTaskManagerZR.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) PauseAll() (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.PauseAll(&_ContractTaskManagerZR.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.PauseAll(&_ContractTaskManagerZR.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c71c165.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task ITaskManagerZRTask, taskResponse ITaskManagerZRTaskResponse, taskResponseMetadata ITaskManagerZRTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c71c165.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) RaiseAndResolveChallenge(task ITaskManagerZRTask, taskResponse ITaskManagerZRTaskResponse, taskResponseMetadata ITaskManagerZRTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.RaiseAndResolveChallenge(&_ContractTaskManagerZR.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c71c165.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) RaiseAndResolveChallenge(task ITaskManagerZRTask, taskResponse ITaskManagerZRTaskResponse, taskResponseMetadata ITaskManagerZRTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.RaiseAndResolveChallenge(&_ContractTaskManagerZR.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.RenounceOwnership(&_ContractTaskManagerZR.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.RenounceOwnership(&_ContractTaskManagerZR.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) RespondToTask(opts *bind.TransactOpts, task ITaskManagerZRTask, taskResponse ITaskManagerZRTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) RespondToTask(task ITaskManagerZRTask, taskResponse ITaskManagerZRTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.RespondToTask(&_ContractTaskManagerZR.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) RespondToTask(task ITaskManagerZRTask, taskResponse ITaskManagerZRTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.RespondToTask(&_ContractTaskManagerZR.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.SetPauserRegistry(&_ContractTaskManagerZR.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.SetPauserRegistry(&_ContractTaskManagerZR.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.SetStaleStakesForbidden(&_ContractTaskManagerZR.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.SetStaleStakesForbidden(&_ContractTaskManagerZR.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.TransferOwnership(&_ContractTaskManagerZR.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.TransferOwnership(&_ContractTaskManagerZR.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTaskManagerZR.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.Unpause(&_ContractTaskManagerZR.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractTaskManagerZR *ContractTaskManagerZRTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTaskManagerZR.Contract.Unpause(&_ContractTaskManagerZR.TransactOpts, newPausedStatus)
}

// ContractTaskManagerZRInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRInitializedIterator struct {
	Event *ContractTaskManagerZRInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRInitialized)
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
		it.Event = new(ContractTaskManagerZRInitialized)
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
func (it *ContractTaskManagerZRInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRInitialized represents a Initialized event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTaskManagerZRInitializedIterator, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRInitializedIterator{contract: _ContractTaskManagerZR.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRInitialized)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseInitialized(log types.Log) (*ContractTaskManagerZRInitialized, error) {
	event := new(ContractTaskManagerZRInitialized)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRNewTaskCreatedIterator struct {
	Event *ContractTaskManagerZRNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRNewTaskCreated)
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
		it.Event = new(ContractTaskManagerZRNewTaskCreated)
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
func (it *ContractTaskManagerZRNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRNewTaskCreated represents a NewTaskCreated event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRNewTaskCreated struct {
	TaskId uint32
	Task   ITaskManagerZRTask
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xf31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskId []uint32) (*ContractTaskManagerZRNewTaskCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRNewTaskCreatedIterator{contract: _ContractTaskManagerZR.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xf31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRNewTaskCreated, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRNewTaskCreated)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseNewTaskCreated(log types.Log) (*ContractTaskManagerZRNewTaskCreated, error) {
	event := new(ContractTaskManagerZRNewTaskCreated)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZROwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZROwnershipTransferredIterator struct {
	Event *ContractTaskManagerZROwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZROwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZROwnershipTransferred)
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
		it.Event = new(ContractTaskManagerZROwnershipTransferred)
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
func (it *ContractTaskManagerZROwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZROwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZROwnershipTransferred represents a OwnershipTransferred event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZROwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTaskManagerZROwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZROwnershipTransferredIterator{contract: _ContractTaskManagerZR.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZROwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZROwnershipTransferred)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTaskManagerZROwnershipTransferred, error) {
	event := new(ContractTaskManagerZROwnershipTransferred)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRPausedIterator struct {
	Event *ContractTaskManagerZRPaused // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRPaused)
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
		it.Event = new(ContractTaskManagerZRPaused)
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
func (it *ContractTaskManagerZRPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRPaused represents a Paused event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractTaskManagerZRPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRPausedIterator{contract: _ContractTaskManagerZR.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRPaused)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParsePaused(log types.Log) (*ContractTaskManagerZRPaused, error) {
	event := new(ContractTaskManagerZRPaused)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRPauserRegistrySetIterator struct {
	Event *ContractTaskManagerZRPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRPauserRegistrySet)
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
		it.Event = new(ContractTaskManagerZRPauserRegistrySet)
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
func (it *ContractTaskManagerZRPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRPauserRegistrySet represents a PauserRegistrySet event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractTaskManagerZRPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRPauserRegistrySetIterator{contract: _ContractTaskManagerZR.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRPauserRegistrySet)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParsePauserRegistrySet(log types.Log) (*ContractTaskManagerZRPauserRegistrySet, error) {
	event := new(ContractTaskManagerZRPauserRegistrySet)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRStaleStakesForbiddenUpdateIterator struct {
	Event *ContractTaskManagerZRStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractTaskManagerZRStaleStakesForbiddenUpdate)
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
func (it *ContractTaskManagerZRStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractTaskManagerZRStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRStaleStakesForbiddenUpdateIterator{contract: _ContractTaskManagerZR.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRStaleStakesForbiddenUpdate)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractTaskManagerZRStaleStakesForbiddenUpdate, error) {
	event := new(ContractTaskManagerZRStaleStakesForbiddenUpdate)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskChallengedSuccessfullyIterator struct {
	Event *ContractTaskManagerZRTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRTaskChallengedSuccessfully)
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
		it.Event = new(ContractTaskManagerZRTaskChallengedSuccessfully)
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
func (it *ContractTaskManagerZRTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskChallengedSuccessfully struct {
	TaskId     uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskId []uint32, challenger []common.Address) (*ContractTaskManagerZRTaskChallengedSuccessfullyIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRTaskChallengedSuccessfullyIterator{contract: _ContractTaskManagerZR.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRTaskChallengedSuccessfully, taskId []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRTaskChallengedSuccessfully)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractTaskManagerZRTaskChallengedSuccessfully, error) {
	event := new(ContractTaskManagerZRTaskChallengedSuccessfully)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractTaskManagerZRTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractTaskManagerZRTaskChallengedUnsuccessfully)
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
func (it *ContractTaskManagerZRTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskChallengedUnsuccessfully struct {
	TaskId     uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskId []uint32, challenger []common.Address) (*ContractTaskManagerZRTaskChallengedUnsuccessfullyIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRTaskChallengedUnsuccessfullyIterator{contract: _ContractTaskManagerZR.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRTaskChallengedUnsuccessfully, taskId []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRTaskChallengedUnsuccessfully)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractTaskManagerZRTaskChallengedUnsuccessfully, error) {
	event := new(ContractTaskManagerZRTaskChallengedUnsuccessfully)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskCompletedIterator struct {
	Event *ContractTaskManagerZRTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRTaskCompleted)
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
		it.Event = new(ContractTaskManagerZRTaskCompleted)
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
func (it *ContractTaskManagerZRTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRTaskCompleted represents a TaskCompleted event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskCompleted struct {
	TaskId uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskId)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []uint32) (*ContractTaskManagerZRTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRTaskCompletedIterator{contract: _ContractTaskManagerZR.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskId)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRTaskCompleted, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRTaskCompleted)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskId)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseTaskCompleted(log types.Log) (*ContractTaskManagerZRTaskCompleted, error) {
	event := new(ContractTaskManagerZRTaskCompleted)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskRespondedIterator struct {
	Event *ContractTaskManagerZRTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRTaskResponded)
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
		it.Event = new(ContractTaskManagerZRTaskResponded)
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
func (it *ContractTaskManagerZRTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRTaskResponded represents a TaskResponded event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRTaskResponded struct {
	TaskResponse         ITaskManagerZRTaskResponse
	TaskResponseMetadata ITaskManagerZRTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567.
//
// Solidity: event TaskResponded((uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractTaskManagerZRTaskRespondedIterator, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRTaskRespondedIterator{contract: _ContractTaskManagerZR.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567.
//
// Solidity: event TaskResponded((uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRTaskResponded)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseTaskResponded(log types.Log) (*ContractTaskManagerZRTaskResponded, error) {
	event := new(ContractTaskManagerZRTaskResponded)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRUnpausedIterator struct {
	Event *ContractTaskManagerZRUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRUnpaused)
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
		it.Event = new(ContractTaskManagerZRUnpaused)
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
func (it *ContractTaskManagerZRUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRUnpaused represents a Unpaused event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractTaskManagerZRUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRUnpausedIterator{contract: _ContractTaskManagerZR.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRUnpaused)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseUnpaused(log types.Log) (*ContractTaskManagerZRUnpaused, error) {
	event := new(ContractTaskManagerZRUnpaused)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskManagerZRValidatorAddressesStoredIterator is returned from FilterValidatorAddressesStored and is used to iterate over the raw logs and unpacked data for ValidatorAddressesStored events raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRValidatorAddressesStoredIterator struct {
	Event *ContractTaskManagerZRValidatorAddressesStored // Event containing the contract specifics and raw log

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
func (it *ContractTaskManagerZRValidatorAddressesStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskManagerZRValidatorAddressesStored)
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
		it.Event = new(ContractTaskManagerZRValidatorAddressesStored)
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
func (it *ContractTaskManagerZRValidatorAddressesStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskManagerZRValidatorAddressesStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskManagerZRValidatorAddressesStored represents a ValidatorAddressesStored event raised by the ContractTaskManagerZR contract.
type ContractTaskManagerZRValidatorAddressesStored struct {
	TaskId uint32
	Count  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorAddressesStored is a free log retrieval operation binding the contract event 0x5a04eb24e83275653ac5d71ef5205698d4ca126250977a126ab5b124d17f2c39.
//
// Solidity: event ValidatorAddressesStored(uint32 indexed taskId, uint256 count)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) FilterValidatorAddressesStored(opts *bind.FilterOpts, taskId []uint32) (*ContractTaskManagerZRValidatorAddressesStoredIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.FilterLogs(opts, "ValidatorAddressesStored", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskManagerZRValidatorAddressesStoredIterator{contract: _ContractTaskManagerZR.contract, event: "ValidatorAddressesStored", logs: logs, sub: sub}, nil
}

// WatchValidatorAddressesStored is a free log subscription operation binding the contract event 0x5a04eb24e83275653ac5d71ef5205698d4ca126250977a126ab5b124d17f2c39.
//
// Solidity: event ValidatorAddressesStored(uint32 indexed taskId, uint256 count)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) WatchValidatorAddressesStored(opts *bind.WatchOpts, sink chan<- *ContractTaskManagerZRValidatorAddressesStored, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractTaskManagerZR.contract.WatchLogs(opts, "ValidatorAddressesStored", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskManagerZRValidatorAddressesStored)
				if err := _ContractTaskManagerZR.contract.UnpackLog(event, "ValidatorAddressesStored", log); err != nil {
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

// ParseValidatorAddressesStored is a log parse operation binding the contract event 0x5a04eb24e83275653ac5d71ef5205698d4ca126250977a126ab5b124d17f2c39.
//
// Solidity: event ValidatorAddressesStored(uint32 indexed taskId, uint256 count)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseValidatorAddressesStored(log types.Log) (*ContractTaskManagerZRValidatorAddressesStored, error) {
	event := new(ContractTaskManagerZRValidatorAddressesStored)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "ValidatorAddressesStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
