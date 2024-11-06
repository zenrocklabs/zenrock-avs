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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveSet\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestActiveSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidatorInTaskSet\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"validatorAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskValidatorAddresses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAddressesStored\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620068aa380380620068aa8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e051610100516165b3620002f7600039600081816102a901528181610623015261360b0152600081816105ec015261274801526000818161047a01528181611595015261292a0152600081816104a101528181612b000152612cc20152600081816104c801528181610f7f01528181612432015281816125ab01526127e501526165b36000f3fe608060405234801561001057600080fd5b50600436106102485760003560e01c80636d14a9871161013b578063b75d69a2116100b8578063f2fde38b1161007c578063f2fde38b1461060e578063f5c9899d14610621578063f63c5bab14610647578063f8c8765e14610650578063fabc1cbc1461066357600080fd5b8063b75d69a214610586578063b98d0908146105a6578063caf73aa0146105b3578063cefdc1d4146105c6578063df5cf723146105e757600080fd5b806379e2c40f116100ff57806379e2c40f146105295780637afa1eed1461053c578063886f11951461054f5780638da5cb5b146105625780639d3a0f2d1461057357600080fd5b80636d14a987146104c35780636efb4636146104ea578063715018a61461050b57806372d18e8d14610513578063767917c61461052157600080fd5b8063416c7e5e116101c95780635c1556621161018d5780635c1556621461042a5780635c975abb1461044a5780635decc3f5146104525780635df4594614610475578063683048351461049c57600080fd5b8063416c7e5e146103a95780634c71c165146103bc5780634f739f74146103cf578063595c6a67146103ef5780635ac86ab7146103f757600080fd5b80632cb223d5116102105780632cb223d51461030b5780632d3de569146103395780632d89f6fc146103495780633563b0d1146103695780633ec54dcf1461038957600080fd5b806310d67a2f1461024d578063136439dd14610262578063171f1d5b146102755780631ad43189146102a4578063245a7bfc146102e0575b600080fd5b61026061025b366004614da7565b610676565b005b610260610270366004614dc4565b610732565b610288610283366004614f42565b610871565b6040805192151583529015156020830152015b60405180910390f35b6102cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161029b565b60ce546102f3906001600160a01b031681565b6040516001600160a01b03909116815260200161029b565b61032b610319366004614fb0565b60cb6020526000908152604090205481565b60405190815260200161029b565b60c9546102cb9063ffffffff1681565b61032b610357366004614fb0565b60ca6020526000908152604090205481565b61037c610377366004615024565b6109fb565b60405161029b919061513f565b61039c610397366004614fb0565b610e91565b60405161029b91906151ae565b6102606103b736600461521e565b610f7d565b6102606103ca3660046152fc565b6110f2565b6103e26103dd3660046153dd565b6116ca565b60405161029b91906154e1565b610260611df0565b61041a6104053660046155ab565b606654600160ff9092169190911b9081161490565b604051901515815260200161029b565b61043d6104383660046155c8565b611eb7565b60405161029b919061567d565b60665461032b565b61041a610460366004614fb0565b60cd6020526000908152604090205460ff1681565b6102f37f000000000000000000000000000000000000000000000000000000000000000081565b6102f37f000000000000000000000000000000000000000000000000000000000000000081565b6102f37f000000000000000000000000000000000000000000000000000000000000000081565b6104fd6104f83660046158ce565b61207f565b60405161029b92919061598e565b610260612f77565b60c95463ffffffff166102cb565b61039c612f8b565b61041a6105373660046159d7565b613077565b60cf546102f3906001600160a01b031681565b6065546102f3906001600160a01b031681565b6033546001600160a01b03166102f3565b610260610581366004615a3a565b6131fc565b610599610594366004615a9e565b613361565b60405161029b9190615aca565b60975461041a9060ff1681565b6102606105c1366004615add565b61341a565b6105d96105d4366004615b64565b613a0b565b60405161029b929190615b9b565b6102f37f000000000000000000000000000000000000000000000000000000000000000081565b61026061061c366004614da7565b613b9d565b7f00000000000000000000000000000000000000000000000000000000000000006102cb565b6102cb61271081565b61026061065e366004615bbc565b613c13565b610260610671366004614dc4565b613d64565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ed9190615c18565b6001600160a01b0316336001600160a01b0316146107265760405162461bcd60e51b815260040161071d90615c35565b60405180910390fd5b61072f81613ec0565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561077a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079e9190615c7f565b6107ba5760405162461bcd60e51b815260040161071d90615c9c565b606654818116146108335760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161071d565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106108b9576108b9615ce4565b60200201518951600160200201518a602001516000600281106108de576108de615ce4565b60200201518b602001516001600281106108fa576108fa615ce4565b602090810291909101518c518d8301516040516109579a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61097a9190615cfa565b90506109ed61099361098c8884613fb7565b869061404e565b61099b6140e2565b6109e36109d4856109ce604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613fb7565b6109dd8c6141a2565b9061404e565b886201d4c0614232565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a619190615c18565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aa3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac79190615c18565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2d9190615c18565b9050600086516001600160401b03811115610b4a57610b4a614ddd565b604051908082528060200260200182016040528015610b7d57816020015b6060815260200190600190039081610b685790505b50905060005b8751811015610e85576000888281518110610ba057610ba0615ce4565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c01573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c299190810190615d1c565b905080516001600160401b03811115610c4457610c44614ddd565b604051908082528060200260200182016040528015610c8f57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610c625790505b50848481518110610ca257610ca2615ce4565b602002602001018190525060005b8151811015610e6f576040518060600160405280876001600160a01b03166347b314e8858581518110610ce557610ce5615ce4565b60200260200101516040518263ffffffff1660e01b8152600401610d0b91815260200190565b602060405180830381865afa158015610d28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4c9190615c18565b6001600160a01b03168152602001838381518110610d6c57610d6c615ce4565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610d9a57610d9a615ce4565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610df6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1a9190615dac565b6001600160601b0316815250858581518110610e3857610e38615ce4565b60200260200101518281518110610e5157610e51615ce4565b60200260200101819052508080610e6790615deb565b915050610cb0565b5050508080610e7d90615deb565b915050610b83565b50979650505050505050565b63ffffffff8116600090815260cc60209081526040808320805482518185028101850190935280835260609492939192909184015b82821015610f72578382906000526020600020018054610ee590615e06565b80601f0160208091040260200160405190810160405280929190818152602001828054610f1190615e06565b8015610f5e5780601f10610f3357610100808354040283529160200191610f5e565b820191906000526020600020905b815481529060010190602001808311610f4157829003601f168201915b505050505081526020019060010190610ec6565b505050509050919050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fdb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fff9190615c18565b6001600160a01b0316336001600160a01b0316146110ab5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a40161071d565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60006111016020850185614fb0565b63ffffffff8116600090815260cb60205260409020549091506111705760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b606482015260840161071d565b8383604051602001611183929190615f6a565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb909352912054146112225760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606482015260840161071d565b63ffffffff8116600090815260cd602052604090205460ff16156112ba5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a40161071d565b6127106112ca6020850185614fb0565b6112d49190615fa8565b63ffffffff164363ffffffff1611156113555760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e000000000000000000606482015260840161071d565b6113626020850185614fb0565b63ffffffff166113756020870187614fb0565b63ffffffff16146113bb5760405162461bcd60e51b815260206004820152601060248201526f0a8c2e6d640928840dad2e6dac2e8c6d60831b604482015260640161071d565b600082516001600160401b038111156113d6576113d6614ddd565b6040519080825280602002602001820160405280156113ff578160200160208202803683370190505b50905060005b83518110156114715761144284828151811061142357611423615ce4565b6020026020010151805160009081526020918201519091526040902090565b82828151811061145457611454615ce4565b60209081029190910101528061146981615deb565b915050611405565b5060006114846040880160208901614fb0565b82604051602001611496929190615fd0565b604051602081830303815290604052805190602001209050846020013581146115405760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a40161071d565b600084516001600160401b0381111561155b5761155b614ddd565b604051908082528060200260200182016040528015611584578160200160208202803683370190505b50905060005b8551811015611677577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106115d4576115d4615ce4565b60200260200101516040518263ffffffff1660e01b81526004016115fa91815260200190565b602060405180830381865afa158015611617573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163b9190615c18565b82828151811061164d5761164d615ce4565b6001600160a01b03909216602092830291909101909101528061166f81615deb565b91505061158a565b5063ffffffff8416600081815260cd6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505050565b6116f56040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611735573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117599190615c18565b90506117866040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906117b6908b908990899060040161600b565b600060405180830381865afa1580156117d3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117fb9190810190616055565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061182d908b908b908b906004016160e3565b600060405180830381865afa15801561184a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118729190810190616055565b6040820152856001600160401b0381111561188f5761188f614ddd565b6040519080825280602002602001820160405280156118c257816020015b60608152602001906001900390816118ad5790505b50606082015260005b60ff8116871115611d01576000856001600160401b038111156118f0576118f0614ddd565b604051908082528060200260200182016040528015611919578160200160208202803683370190505b5083606001518360ff168151811061193357611933615ce4565b602002602001018190525060005b86811015611c015760008c6001600160a01b03166304ec63518a8a8581811061196c5761196c615ce4565b905060200201358e8860000151868151811061198a5761198a615ce4565b60200260200101516040518463ffffffff1660e01b81526004016119c79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156119e4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a08919061610c565b90506001600160c01b038116611aac5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40161071d565b8a8a8560ff16818110611ac157611ac1615ce4565b6001600160c01b03841692013560f81c9190911c600190811614159050611bee57856001600160a01b031663dd9846b98a8a85818110611b0357611b03615ce4565b905060200201358d8d8860ff16818110611b1f57611b1f615ce4565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611b75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b999190616135565b85606001518560ff1681518110611bb257611bb2615ce4565b60200260200101518481518110611bcb57611bcb615ce4565b63ffffffff9092166020928302919091019091015282611bea81615deb565b9350505b5080611bf981615deb565b915050611941565b506000816001600160401b03811115611c1c57611c1c614ddd565b604051908082528060200260200182016040528015611c45578160200160208202803683370190505b50905060005b82811015611cc65784606001518460ff1681518110611c6c57611c6c615ce4565b60200260200101518181518110611c8557611c85615ce4565b6020026020010151828281518110611c9f57611c9f615ce4565b63ffffffff9092166020928302919091019091015280611cbe81615deb565b915050611c4b565b508084606001518460ff1681518110611ce157611ce1615ce4565b602002602001018190525050508080611cf990616152565b9150506118cb565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d669190615c18565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611d99908b908b908e90600401616172565b600060405180830381865afa158015611db6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611dde9190810190616055565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611e38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e5c9190615c7f565b611e785760405162461bcd60e51b815260040161071d90615c9c565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611ee992919061619c565b600060405180830381865afa158015611f06573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f2e9190810190616055565b9050600084516001600160401b03811115611f4b57611f4b614ddd565b604051908082528060200260200182016040528015611f74578160200160208202803683370190505b50905060005b855181101561207557866001600160a01b03166304ec6351878381518110611fa457611fa4615ce4565b602002602001015187868581518110611fbf57611fbf615ce4565b60200260200101516040518463ffffffff1660e01b8152600401611ffc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612019573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203d919061610c565b6001600160c01b031682828151811061205857612058615ce4565b60209081029190910101528061206d81615deb565b915050611f7a565b5095945050505050565b60408051808201909152606080825260208201526000846120f65760405162461bcd60e51b8152602060048201526037602482015260008051602061655e83398151915260448201527f7265733a20656d7074792071756f72756d20696e707574000000000000000000606482015260840161071d565b6040830151518514801561210e575060a08301515185145b801561211e575060c08301515185145b801561212e575060e08301515185145b6121985760405162461bcd60e51b8152602060048201526041602482015260008051602061655e83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a40161071d565b825151602084015151146122105760405162461bcd60e51b81526020600482015260446024820181905260008051602061655e833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a40161071d565b4363ffffffff168463ffffffff161061227f5760405162461bcd60e51b815260206004820152603c602482015260008051602061655e83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b00000000606482015260840161071d565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156122c0576122c0614ddd565b6040519080825280602002602001820160405280156122e9578160200160208202803683370190505b506020820152866001600160401b0381111561230757612307614ddd565b604051908082528060200260200182016040528015612330578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561236457612364614ddd565b60405190808252806020026020018201604052801561238d578160200160208202803683370190505b5081526020860151516001600160401b038111156123ad576123ad614ddd565b6040519080825280602002602001820160405280156123d6578160200160208202803683370190505b50816020018190525060006124a88a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa15801561247f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124a391906161f0565b614456565b905060005b876020015151811015612724576124d38860200151828151811061142357611423615ce4565b836020015182815181106124e9576124e9615ce4565b602090810291909101015280156125a957602083015161250a60018361620d565b8151811061251a5761251a615ce4565b602002602001015160001c8360200151828151811061253b5761253b615ce4565b602002602001015160001c116125a9576040805162461bcd60e51b815260206004820152602481019190915260008051602061655e83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f72746564606482015260840161071d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106125ee576125ee615ce4565b60200260200101518b8b60000151858151811061260d5761260d615ce4565b60200260200101516040518463ffffffff1660e01b815260040161264a9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612667573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061268b919061610c565b6001600160c01b0316836000015182815181106126aa576126aa615ce4565b60200260200101818152505061271061098c6126e484866000015185815181106126d6576126d6615ce4565b6020026020010151166144e7565b8a6020015184815181106126fa576126fa615ce4565b602002602001015161451290919063ffffffff16565b94508061271c81615deb565b9150506124ad565b505061272f836145f6565b60975490935060ff166000816127465760006127c8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127c89190616224565b905060005b8a811015612e46578215612928578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061282457612824615ce4565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612864573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128889190616224565b612892919061623d565b116129285760405162461bcd60e51b8152602060048201526066602482015260008051602061655e83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c40161071d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061296957612969615ce4565b9050013560f81c60f81b60f81c8c8c60a00151858151811061298d5761298d615ce4565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156129e9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a0d9190616255565b6001600160401b031916612a308a60400151838151811061142357611423615ce4565b67ffffffffffffffff191614612acc5760405162461bcd60e51b8152602060048201526061602482015260008051602061655e83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c40161071d565b612afc89604001518281518110612ae557612ae5615ce4565b60200260200101518761404e90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612b3f57612b3f615ce4565b9050013560f81c60f81b60f81c8c8c60c001518581518110612b6357612b63615ce4565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612bbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612be39190615dac565b85602001518281518110612bf957612bf9615ce4565b6001600160601b03909216602092830291909101820152850151805182908110612c2557612c25615ce4565b602002602001015185600001518281518110612c4357612c43615ce4565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612e3157612cbb86600001518281518110612c8d57612c8d615ce4565b60200260200101518f8f86818110612ca757612ca7615ce4565b600192013560f81c9290921c811614919050565b15612e1f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612d0157612d01615ce4565b9050013560f81c60f81b60f81c8e89602001518581518110612d2557612d25615ce4565b60200260200101518f60e001518881518110612d4357612d43615ce4565b60200260200101518781518110612d5c57612d5c615ce4565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612dc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612de49190615dac565b8751805185908110612df857612df8615ce4565b60200260200101818151612e0c9190616280565b6001600160601b03169052506001909101905b80612e2981615deb565b915050612c67565b50508080612e3e90615deb565b9150506127cd565b505050600080612e608c868a606001518b60800151610871565b9150915081612ed15760405162461bcd60e51b8152602060048201526043602482015260008051602061655e83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a40161071d565b80612f325760405162461bcd60e51b8152602060048201526039602482015260008051602061655e83398151915260448201527f7265733a207369676e617475726520697320696e76616c696400000000000000606482015260840161071d565b50506000878260200151604051602001612f4d929190615fd0565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612f7f614691565b612f8960006146eb565b565b60c95463ffffffff16600090815260cc60209081526040808320805482518185028101850190935280835260609492939192909184015b8282101561306e578382906000526020600020018054612fe190615e06565b80601f016020809104026020016040519081016040528092919081815260200182805461300d90615e06565b801561305a5780601f1061302f5761010080835404028352916020019161305a565b820191906000526020600020905b81548152906001019060200180831161303d57829003601f168201915b505050505081526020019060010190612fc2565b50505050905090565b63ffffffff8216600090815260cc6020908152604080832080548251818502810185019093528083528493849084015b828210156131535783829060005260206000200180546130c690615e06565b80601f01602080910402602001604051908101604052809291908181526020018280546130f290615e06565b801561313f5780601f106131145761010080835404028352916020019161313f565b820191906000526020600020905b81548152906001019060200180831161312257829003601f168201915b5050505050815260200190600101906130a7565b50505050905060005b81518110156131ef578360405160200161317691906162a8565b6040516020818303038152906040528051906020012082828151811061319e5761319e615ce4565b60200260200101516040516020016131b691906162a8565b6040516020818303038152906040528051906020012014156131dd576001925050506131f6565b806131e781615deb565b91505061315c565b5060009150505b92915050565b60cf546001600160a01b031633146132605760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b606482015260840161071d565b60408051608081018252606081830181905263ffffffff8781168352438116602080850191909152908716918301919091528251601f85018290048202810182019093528383529091908490849081908401838280828437600092019190915250505050604080830191909152516132dc9082906020016162c4565b60408051601f19818403018152828252805160209182012063ffffffff8916600081815260ca9093529290912055907ff31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5906133389084906162c4565b60405180910390a2505060c9805463ffffffff191663ffffffff94909416939093179092555050565b60cc602052816000526040600020818154811061337d57600080fd5b9060005260206000200160009150915050805461339990615e06565b80601f01602080910402602001604051908101604052809291908181526020018280546133c590615e06565b80156134125780601f106133e757610100808354040283529160200191613412565b820191906000526020600020905b8154815290600101906020018083116133f557829003601f168201915b505050505081565b60ce546001600160a01b031633146134745760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c6572000000604482015260640161071d565b60006134866040850160208601614fb0565b90503660006134986040870187616316565b909250905060006134af6080880160608901614fb0565b905060ca60006134c260208a018a614fb0565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016134ee919061635c565b60405160208183030381529060405280519060200120146135775760405162461bcd60e51b815260206004820152603d60248201527f537570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606482015260840161071d565b600060cb8161358960208b018b614fb0565b63ffffffff1663ffffffff16815260200190815260200160002054146136065760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b606482015260840161071d565b6136307f000000000000000000000000000000000000000000000000000000000000000085615fa8565b63ffffffff164363ffffffff1611156136a15760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b606482015260840161071d565b6000866040516020016136b491906163d8565b6040516020818303038152906040528051906020012090506000806136dc8387878a8c61207f565b9150915060005b858110156137db578460ff168360200151828151811061370557613705615ce4565b602002602001015161371791906163eb565b6001600160601b031660438460000151838151811061373857613738615ce4565b60200260200101516001600160601b0316613753919061641a565b10156137c9576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d606482015260840161071d565b806137d381615deb565b9150506136e3565b5060cc60006137ed60208d018d614fb0565b63ffffffff1663ffffffff16815260200190815260200160002060006138139190614baa565b60005b61382360208b018b616439565b90508110156138f657600061383b60208c018c616439565b8381811061384b5761384b615ce4565b905060200281019061385d9190616316565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093945060cc9392506138a591505060208f018f614fb0565b63ffffffff168152602080820192909252604001600090812080546001810182559082529082902083516138e193919092019190840190614bc8565b505080806138ee90615deb565b915050613816565b5061390460208b018b614fb0565b63ffffffff167f48d8dea4bc1eac7c64776d6c2d268a0c0dc0c42b5cc20c85e7126337fc9532b361393860208c018c616439565b604051613946929190616482565b60405180910390a260408051808201825263ffffffff4316815260208082018490529151909161397a918c918491016164e8565b6040516020818303038152906040528051906020012060cb60008d60000160208101906139a79190614fb0565b63ffffffff1663ffffffff168152602001908152602001600020819055507fb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce35678a826040516139f69291906164e8565b60405180910390a15050505050505050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110613a4657613a46615ce4565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90613a82908890869060040161619c565b600060405180830381865afa158015613a9f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613ac79190810190616055565b600081518110613ad957613ad9615ce4565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613b45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b69919061610c565b6001600160c01b031690506000613b7f8261473d565b905081613b8d8a838a6109fb565b9550955050505050935093915050565b613ba5614691565b6001600160a01b038116613c0a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161071d565b61072f816146eb565b600054610100900460ff1615808015613c335750600054600160ff909116105b80613c4d5750303b158015613c4d575060005460ff166001145b613cb05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161071d565b6000805460ff191660011790558015613cd3576000805461ff0019166101001790555b613cde856000614809565b613ce7846146eb565b60ce80546001600160a01b038086166001600160a01b03199283161790925560cf8054928516929091169190911790558015613d5d576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613db7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ddb9190615c18565b6001600160a01b0316336001600160a01b031614613e0b5760405162461bcd60e51b815260040161071d90615c35565b606654198119606654191614613e895760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161071d565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610866565b6001600160a01b038116613f4e5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161071d565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613fd3614c4c565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561400657614008565bfe5b50806140465760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640161071d565b505092915050565b604080518082019091526000808252602082015261406a614c6a565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156140065750806140465760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161071d565b6140ea614c88565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806141d260008051602061653e83398151915286615cfa565b90505b6141de816148f3565b909350915060008051602061653e833981519152828309831415614218576040805180820190915290815260208101919091529392505050565b60008051602061653e8339815191526001820890506141d5565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190614264614cad565b60005b600281101561442957600061427d82600661641a565b905084826002811061429157614291615ce4565b602002015151836142a383600061623d565b600c81106142b3576142b3615ce4565b60200201528482600281106142ca576142ca615ce4565b602002015160200151838260016142e1919061623d565b600c81106142f1576142f1615ce4565b602002015283826002811061430857614308615ce4565b602002015151518361431b83600261623d565b600c811061432b5761432b615ce4565b602002015283826002811061434257614342615ce4565b602002015151600160200201518361435b83600361623d565b600c811061436b5761436b615ce4565b602002015283826002811061438257614382615ce4565b60200201516020015160006002811061439d5761439d615ce4565b6020020151836143ae83600461623d565b600c81106143be576143be615ce4565b60200201528382600281106143d5576143d5615ce4565b6020020151602001516001600281106143f0576143f0615ce4565b60200201518361440183600561623d565b600c811061441157614411615ce4565b6020020152508061442181615deb565b915050614267565b50614432614ccc565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008061446284614975565b9050808360ff166001901b116144e05760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c756500606482015260840161071d565b9392505050565b6000805b82156131f6576144fc60018461620d565b909216918061450a8161651b565b9150506144eb565b60408051808201909152600080825260208201526102008261ffff161061456e5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b604482015260640161071d565b8161ffff16600114156145825750816131f6565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16106145eb57600161ffff871660ff83161c811614156145ce576145cb848461404e565b93505b6145d8838461404e565b92506201fffe600192831b16910161459e565b509195945050505050565b6040805180820190915260008082526020820152815115801561461b57506020820151155b15614639575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061653e833981519152846020015161466c9190615cfa565b6146849060008051602061653e83398151915261620d565b905292915050565b919050565b6033546001600160a01b03163314612f895760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161071d565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606060008061474b846144e7565b61ffff166001600160401b0381111561476657614766614ddd565b6040519080825280601f01601f191660200182016040528015614790576020820181803683370190505b5090506000805b8251821080156147a8575061010081105b156147ff576001811b9350858416156147ef578060f81b8383815181106147d1576147d1615ce4565b60200101906001600160f81b031916908160001a9053508160010191505b6147f881615deb565b9050614797565b5090949350505050565b6065546001600160a01b031615801561482a57506001600160a01b03821615155b6148ac5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161071d565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26148ef82613ec0565b5050565b6000808060008051602061653e833981519152600360008051602061653e8339815191528660008051602061653e833981519152888909090890506000614969827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061653e833981519152614b02565b91959194509092505050565b6000610100825111156149fe5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a40161071d565b8151614a0c57506000919050565b60008083600081518110614a2257614a22615ce4565b0160200151600160f89190911c81901b92505b8451811015614af957848181518110614a5057614a50615ce4565b0160200151600160f89190911c1b9150828211614ae55760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a40161071d565b91811791614af281615deb565b9050614a35565b50909392505050565b600080614b0d614ccc565b614b15614cea565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015614006575082614b9f5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161071d565b505195945050505050565b508054600082559060005260206000209081019061072f9190614d08565b828054614bd490615e06565b90600052602060002090601f016020900481019282614bf65760008555614c3c565b82601f10614c0f57805160ff1916838001178555614c3c565b82800160010185558215614c3c579182015b82811115614c3c578251825591602001919060010190614c21565b50614c48929150614d25565b5090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280614c9b614d3a565b8152602001614ca8614d3a565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b80821115614c48576000614d1c8282614d58565b50600101614d08565b5b80821115614c485760008155600101614d26565b60405180604001604052806002906020820280368337509192915050565b508054614d6490615e06565b6000825580601f10614d74575050565b601f01602090049060005260206000209081019061072f9190614d25565b6001600160a01b038116811461072f57600080fd5b600060208284031215614db957600080fd5b81356144e081614d92565b600060208284031215614dd657600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614e1557614e15614ddd565b60405290565b60405161010081016001600160401b0381118282101715614e1557614e15614ddd565b604051601f8201601f191681016001600160401b0381118282101715614e6657614e66614ddd565b604052919050565b600060408284031215614e8057600080fd5b614e88614df3565b9050813581526020820135602082015292915050565b600082601f830112614eaf57600080fd5b604051604081018181106001600160401b0382111715614ed157614ed1614ddd565b8060405250806040840185811115614ee857600080fd5b845b818110156145eb578035835260209283019201614eea565b600060808284031215614f1457600080fd5b614f1c614df3565b9050614f288383614e9e565b8152614f378360408401614e9e565b602082015292915050565b6000806000806101208587031215614f5957600080fd5b84359350614f6a8660208701614e6e565b9250614f798660608701614f02565b9150614f888660e08701614e6e565b905092959194509250565b63ffffffff8116811461072f57600080fd5b803561468c81614f93565b600060208284031215614fc257600080fd5b81356144e081614f93565b60006001600160401b03831115614fe657614fe6614ddd565b614ff9601f8401601f1916602001614e3e565b905082815283838301111561500d57600080fd5b828260208301376000602084830101529392505050565b60008060006060848603121561503957600080fd5b833561504481614d92565b925060208401356001600160401b0381111561505f57600080fd5b8401601f8101861361507057600080fd5b61507f86823560208401614fcd565b925050604084013561509081614f93565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015615131578385038a52825180518087529087019087870190845b8181101561511c57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016150d8565b50509a87019a955050918501916001016150ba565b509298975050505050505050565b6020815260006144e0602083018461509b565b60005b8381101561516d578181015183820152602001615155565b8381111561517c576000848401525b50505050565b6000815180845261519a816020860160208601615152565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561520357603f198886030184526151f1858351615182565b945092850192908501906001016151d5565b5092979650505050505050565b801515811461072f57600080fd5b60006020828403121561523057600080fd5b81356144e081615210565b60006080828403121561524d57600080fd5b50919050565b60006040828403121561524d57600080fd5b60006001600160401b0382111561527e5761527e614ddd565b5060051b60200190565b600082601f83011261529957600080fd5b813560206152ae6152a983615265565b614e3e565b82815260069290921b840181019181810190868411156152cd57600080fd5b8286015b848110156152f1576152e38882614e6e565b8352918301916040016152d1565b509695505050505050565b60008060008060a0858703121561531257600080fd5b84356001600160401b038082111561532957600080fd5b6153358883890161523b565b9550602087013591508082111561534b57600080fd5b61535788838901615253565b94506153668860408901615253565b9350608087013591508082111561537c57600080fd5b5061538987828801615288565b91505092959194509250565b60008083601f8401126153a757600080fd5b5081356001600160401b038111156153be57600080fd5b6020830191508360208285010111156153d657600080fd5b9250929050565b600080600080600080608087890312156153f657600080fd5b863561540181614d92565b9550602087013561541181614f93565b945060408701356001600160401b038082111561542d57600080fd5b6154398a838b01615395565b9096509450606089013591508082111561545257600080fd5b818901915089601f83011261546657600080fd5b81358181111561547557600080fd5b8a60208260051b850101111561548a57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b838110156154d657815163ffffffff16875295820195908201906001016154b4565b509495945050505050565b6000602080835283516080828501526154fd60a08501826154a0565b905081850151601f198086840301604087015261551a83836154a0565b9250604087015191508086840301606087015261553783836154a0565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561558e578487830301845261557c8287516154a0565b95880195938801939150600101615562565b509998505050505050505050565b60ff8116811461072f57600080fd5b6000602082840312156155bd57600080fd5b81356144e08161559c565b6000806000606084860312156155dd57600080fd5b83356155e881614d92565b92506020848101356001600160401b0381111561560457600080fd5b8501601f8101871361561557600080fd5b80356156236152a982615265565b81815260059190911b8201830190838101908983111561564257600080fd5b928401925b8284101561566057833582529284019290840190615647565b809650505050505061567460408501614fa5565b90509250925092565b6020808252825182820181905260009190848201906040850190845b818110156156b557835183529284019291840191600101615699565b50909695505050505050565b600082601f8301126156d257600080fd5b813560206156e26152a983615265565b82815260059290921b8401810191818101908684111561570157600080fd5b8286015b848110156152f157803561571881614f93565b8352918301918301615705565b600082601f83011261573657600080fd5b813560206157466152a983615265565b82815260059290921b8401810191818101908684111561576557600080fd5b8286015b848110156152f15780356001600160401b038111156157885760008081fd5b6157968986838b01016156c1565b845250918301918301615769565b600061018082840312156157b757600080fd5b6157bf614e1b565b905081356001600160401b03808211156157d857600080fd5b6157e4858386016156c1565b835260208401359150808211156157fa57600080fd5b61580685838601615288565b6020840152604084013591508082111561581f57600080fd5b61582b85838601615288565b604084015261583d8560608601614f02565b606084015261584f8560e08601614e6e565b608084015261012084013591508082111561586957600080fd5b615875858386016156c1565b60a084015261014084013591508082111561588f57600080fd5b61589b858386016156c1565b60c08401526101608401359150808211156158b557600080fd5b506158c284828501615725565b60e08301525092915050565b6000806000806000608086880312156158e657600080fd5b8535945060208601356001600160401b038082111561590457600080fd5b61591089838a01615395565b90965094506040880135915061592582614f93565b9092506060870135908082111561593b57600080fd5b50615948888289016157a4565b9150509295509295909350565b600081518084526020808501945080840160005b838110156154d65781516001600160601b031687529582019590820190600101615969565b60408152600083516040808401526159a96080840182615955565b90506020850151603f198483030160608501526159c68282615955565b925050508260208301529392505050565b600080604083850312156159ea57600080fd5b82356159f581614f93565b915060208301356001600160401b03811115615a1057600080fd5b8301601f81018513615a2157600080fd5b615a3085823560208401614fcd565b9150509250929050565b60008060008060608587031215615a5057600080fd5b8435615a5b81614f93565b93506020850135615a6b81614f93565b925060408501356001600160401b03811115615a8657600080fd5b615a9287828801615395565b95989497509550505050565b60008060408385031215615ab157600080fd5b8235615abc81614f93565b946020939093013593505050565b6020815260006144e06020830184615182565b600080600060608486031215615af257600080fd5b83356001600160401b0380821115615b0957600080fd5b615b158783880161523b565b94506020860135915080821115615b2b57600080fd5b615b3787838801615253565b93506040860135915080821115615b4d57600080fd5b50615b5a868287016157a4565b9150509250925092565b600080600060608486031215615b7957600080fd5b8335615b8481614d92565b925060208401359150604084013561509081614f93565b828152604060208201526000615bb4604083018461509b565b949350505050565b60008060008060808587031215615bd257600080fd5b8435615bdd81614d92565b93506020850135615bed81614d92565b92506040850135615bfd81614d92565b91506060850135615c0d81614d92565b939692955090935050565b600060208284031215615c2a57600080fd5b81516144e081614d92565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215615c9157600080fd5b81516144e081615210565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082615d1757634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215615d2f57600080fd5b82516001600160401b03811115615d4557600080fd5b8301601f81018513615d5657600080fd5b8051615d646152a982615265565b81815260059190911b82018301908381019087831115615d8357600080fd5b928401925b82841015615da157835182529284019290840190615d88565b979650505050505050565b600060208284031215615dbe57600080fd5b81516001600160601b03811681146144e057600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415615dff57615dff615dd5565b5060010190565b600181811c90821680615e1a57607f821691505b6020821081141561524d57634e487b7160e01b600052602260045260246000fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112615e7b57600080fd5b83016020810192503590506001600160401b03811115615e9a57600080fd5b8036038313156153d657600080fd5b6000604083018235615eba81614f93565b63ffffffff16845260208381013536859003601e19018112615edb57600080fd5b840181810190356001600160401b03811115615ef657600080fd5b8060051b803603871315615f0957600080fd5b6040848901529381905260609387018401938290880160005b83811015615f5c57898703605f19018252615f3d8386615e64565b615f48898284615e3b565b985050509185019190850190600101615f22565b509498975050505050505050565b606081526000615f7d6060830185615ea9565b90508235615f8a81614f93565b63ffffffff8116602084015250602083013560408301529392505050565b600063ffffffff808316818516808303821115615fc757615fc7615dd5565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561520357815185529382019390820190600101615fef565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561603857600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561606857600080fd5b82516001600160401b0381111561607e57600080fd5b8301601f8101851361608f57600080fd5b805161609d6152a982615265565b81815260059190911b820183019083810190878311156160bc57600080fd5b928401925b82841015615da15783516160d481614f93565b825292840192908401906160c1565b63ffffffff84168152604060208201526000616103604083018486615e3b565b95945050505050565b60006020828403121561611e57600080fd5b81516001600160c01b03811681146144e057600080fd5b60006020828403121561614757600080fd5b81516144e081614f93565b600060ff821660ff81141561616957616169615dd5565b60010192915050565b604081526000616186604083018587615e3b565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156161e3578451835293830193918301916001016161c7565b5090979650505050505050565b60006020828403121561620257600080fd5b81516144e08161559c565b60008282101561621f5761621f615dd5565b500390565b60006020828403121561623657600080fd5b5051919050565b6000821982111561625057616250615dd5565b500190565b60006020828403121561626757600080fd5b815167ffffffffffffffff19811681146144e057600080fd5b60006001600160601b03838116908316818110156162a0576162a0615dd5565b039392505050565b600082516162ba818460208701615152565b9190910192915050565b60208152600063ffffffff8084511660208401528060208501511660408401526040840151608060608501526162fd60a0850182615182565b9050816060860151166080850152809250505092915050565b6000808335601e1984360301811261632d57600080fd5b8301803591506001600160401b0382111561634757600080fd5b6020019150368190038213156153d657600080fd5b602081526000823561636d81614f93565b63ffffffff80821660208501526020850135915061638a82614f93565b808216604085015261639f6040860186615e64565b9250608060608601526163b660a086018483615e3b565b92505060608501356163c781614f93565b166080939093019290925250919050565b6020815260006144e06020830184615ea9565b60006001600160601b038083168185168183048111821515161561641157616411615dd5565b02949350505050565b600081600019048311821515161561643457616434615dd5565b500290565b6000808335601e1984360301811261645057600080fd5b8301803591506001600160401b0382111561646a57600080fd5b6020019150600581901b36038213156153d657600080fd5b60208082528181018390526000906040600585901b8401810190840186845b878110156164db57868403603f190183526164bc828a615e64565b6164c7868284615e3b565b9550505091840191908401906001016164a1565b5091979650505050505050565b6060815260006164fb6060830185615ea9565b905063ffffffff8351166020830152602083015160408301529392505050565b600061ffff8083168181141561653357616533615dd5565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220dcc39acfcc3380b7c8eaf68ca07235ec005cdd410ddd51398313178069989d5564736f6c634300080c0033",
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

// TaskValidatorAddresses is a free data retrieval call binding the contract method 0xb75d69a2.
//
// Solidity: function taskValidatorAddresses(uint32 , uint256 ) view returns(string)
func (_ContractTaskManagerZR *ContractTaskManagerZRCaller) TaskValidatorAddresses(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _ContractTaskManagerZR.contract.Call(opts, &out, "taskValidatorAddresses", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TaskValidatorAddresses is a free data retrieval call binding the contract method 0xb75d69a2.
//
// Solidity: function taskValidatorAddresses(uint32 , uint256 ) view returns(string)
func (_ContractTaskManagerZR *ContractTaskManagerZRSession) TaskValidatorAddresses(arg0 uint32, arg1 *big.Int) (string, error) {
	return _ContractTaskManagerZR.Contract.TaskValidatorAddresses(&_ContractTaskManagerZR.CallOpts, arg0, arg1)
}

// TaskValidatorAddresses is a free data retrieval call binding the contract method 0xb75d69a2.
//
// Solidity: function taskValidatorAddresses(uint32 , uint256 ) view returns(string)
func (_ContractTaskManagerZR *ContractTaskManagerZRCallerSession) TaskValidatorAddresses(arg0 uint32, arg1 *big.Int) (string, error) {
	return _ContractTaskManagerZR.Contract.TaskValidatorAddresses(&_ContractTaskManagerZR.CallOpts, arg0, arg1)
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
	TaskId    uint32
	Addresses []string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAddressesStored is a free log retrieval operation binding the contract event 0x48d8dea4bc1eac7c64776d6c2d268a0c0dc0c42b5cc20c85e7126337fc9532b3.
//
// Solidity: event ValidatorAddressesStored(uint32 indexed taskId, string[] addresses)
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

// WatchValidatorAddressesStored is a free log subscription operation binding the contract event 0x48d8dea4bc1eac7c64776d6c2d268a0c0dc0c42b5cc20c85e7126337fc9532b3.
//
// Solidity: event ValidatorAddressesStored(uint32 indexed taskId, string[] addresses)
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

// ParseValidatorAddressesStored is a log parse operation binding the contract event 0x48d8dea4bc1eac7c64776d6c2d268a0c0dc0c42b5cc20c85e7126337fc9532b3.
//
// Solidity: event ValidatorAddressesStored(uint32 indexed taskId, string[] addresses)
func (_ContractTaskManagerZR *ContractTaskManagerZRFilterer) ParseValidatorAddressesStored(log types.Log) (*ContractTaskManagerZRValidatorAddressesStored, error) {
	event := new(ContractTaskManagerZRValidatorAddressesStored)
	if err := _ContractTaskManagerZR.contract.UnpackLog(event, "ValidatorAddressesStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
