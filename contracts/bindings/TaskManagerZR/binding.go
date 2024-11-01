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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005f6b38038062005f6b8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615c74620002f76000396000818161027d0152818161059c015261316e01526000818161056501526125d501526000818161042e0152818161142201526127b70152600081816104550152818161298d0152612b4f01526000818161047c01528181610e0c015281816122bf0152818161243801526126720152615c746000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80636830483511610125578063b98d0908116100ad578063f2fde38b1161007c578063f2fde38b14610587578063f5c9899d1461059a578063f63c5bab146105c0578063f8c8765e146105c9578063fabc1cbc146105dc57600080fd5b8063b98d09081461051f578063caf73aa01461052c578063cefdc1d41461053f578063df5cf7231461056057600080fd5b806372d18e8d116100f457806372d18e8d146104c75780637afa1eed146104d5578063886f1195146104e85780638da5cb5b146104fb5780639d3a0f2d1461050c57600080fd5b806368304835146104505780636d14a987146104775780636efb46361461049e578063715018a6146104bf57600080fd5b8063416c7e5e116101a85780635ac86ab7116101775780635ac86ab7146103ab5780635c155662146103de5780635c975abb146103fe5780635decc3f5146104065780635df459461461042957600080fd5b8063416c7e5e1461035d5780634c71c165146103705780634f739f7414610383578063595c6a67146103a357600080fd5b8063245a7bfc116101ef578063245a7bfc146102b45780632cb223d5146102df5780632d3de5691461030d5780632d89f6fc1461031d5780633563b0d11461033d57600080fd5b806310d67a2f14610221578063136439dd14610236578063171f1d5b146102495780631ad4318914610278575b600080fd5b61023461022f36600461468c565b6105ef565b005b6102346102443660046146a9565b6106ab565b61025c610257366004614827565b6107ea565b6040805192151583529015156020830152015b60405180910390f35b61029f7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026f565b60cd546102c7906001600160a01b031681565b6040516001600160a01b03909116815260200161026f565b6102ff6102ed366004614895565b60cb6020526000908152604090205481565b60405190815260200161026f565b60c95461029f9063ffffffff1681565b6102ff61032b366004614895565b60ca6020526000908152604090205481565b61035061034b3660046148b2565b610974565b60405161026f9190614a0d565b61023461036b366004614a35565b610e0a565b61023461037e366004614b13565b610f7f565b610396610391366004614bf4565b611557565b60405161026f9190614cf8565b610234611c7d565b6103ce6103b9366004614dc2565b606654600160ff9092169190911b9081161490565b604051901515815260200161026f565b6103f16103ec366004614ddf565b611d44565b60405161026f9190614e8b565b6066546102ff565b6103ce610414366004614895565b60cc6020526000908152604090205460ff1681565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6104b16104ac3660046150dc565b611f0c565b60405161026f92919061519c565b610234612e04565b60c95463ffffffff1661029f565b60ce546102c7906001600160a01b031681565b6065546102c7906001600160a01b031681565b6033546001600160a01b03166102c7565b61023461051a3660046151e5565b612e18565b6097546103ce9060ff1681565b61023461053a366004615249565b612f7d565b61055261054d3660046152d0565b6133fc565b60405161026f929190615312565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b61023461059536600461468c565b61358e565b7f000000000000000000000000000000000000000000000000000000000000000061029f565b61029f61271081565b6102346105d7366004615333565b613604565b6102346105ea3660046146a9565b613755565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610642573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610666919061538f565b6001600160a01b0316336001600160a01b03161461069f5760405162461bcd60e51b8152600401610696906153ac565b60405180910390fd5b6106a8816138b1565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071791906153f6565b6107335760405162461bcd60e51b815260040161069690615413565b606654818116146107ac5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610696565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106108325761083261545b565b60200201518951600160200201518a602001516000600281106108575761085761545b565b60200201518b602001516001600281106108735761087361545b565b602090810291909101518c518d8301516040516108d09a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108f39190615471565b905061096661090c61090588846139a8565b8690613a3f565b610914613ad3565b61095c61094d85610947604080518082018252600080825260209182015281518083019092526001825260029082015290565b906139a8565b6109568c613b93565b90613a3f565b886201d4c0613c23565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109da919061538f565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a40919061538f565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa6919061538f565b9050600086516001600160401b03811115610ac357610ac36146c2565b604051908082528060200260200182016040528015610af657816020015b6060815260200190600190039081610ae15790505b50905060005b8751811015610dfe576000888281518110610b1957610b1961545b565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b7a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ba29190810190615493565b905080516001600160401b03811115610bbd57610bbd6146c2565b604051908082528060200260200182016040528015610c0857816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bdb5790505b50848481518110610c1b57610c1b61545b565b602002602001018190525060005b8151811015610de8576040518060600160405280876001600160a01b03166347b314e8858581518110610c5e57610c5e61545b565b60200260200101516040518263ffffffff1660e01b8152600401610c8491815260200190565b602060405180830381865afa158015610ca1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc5919061538f565b6001600160a01b03168152602001838381518110610ce557610ce561545b565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610d1357610d1361545b565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d6f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d939190615523565b6001600160601b0316815250858581518110610db157610db161545b565b60200260200101518281518110610dca57610dca61545b565b60200260200101819052508080610de090615562565b915050610c29565b5050508080610df690615562565b915050610afc565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8c919061538f565b6001600160a01b0316336001600160a01b031614610f385760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610696565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6000610f8e6020850185614895565b63ffffffff8116600090815260cb6020526040902054909150610ffd5760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610696565b83836040516020016110109291906156ac565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb909352912054146110af5760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610696565b63ffffffff8116600090815260cc602052604090205460ff16156111475760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610696565b6127106111576020850185614895565b61116191906156ea565b63ffffffff164363ffffffff1611156111e25760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610696565b6111ef6020850185614895565b63ffffffff166112026020870187614895565b63ffffffff16146112485760405162461bcd60e51b815260206004820152601060248201526f0a8c2e6d640928840dad2e6dac2e8c6d60831b6044820152606401610696565b600082516001600160401b03811115611263576112636146c2565b60405190808252806020026020018201604052801561128c578160200160208202803683370190505b50905060005b83518110156112fe576112cf8482815181106112b0576112b061545b565b6020026020010151805160009081526020918201519091526040902090565b8282815181106112e1576112e161545b565b6020908102919091010152806112f681615562565b915050611292565b5060006113116040880160208901614895565b82604051602001611323929190615712565b604051602081830303815290604052805190602001209050846020013581146113cd5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610696565b600084516001600160401b038111156113e8576113e86146c2565b604051908082528060200260200182016040528015611411578160200160208202803683370190505b50905060005b8551811015611504577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106114615761146161545b565b60200260200101516040518263ffffffff1660e01b815260040161148791815260200190565b602060405180830381865afa1580156114a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114c8919061538f565b8282815181106114da576114da61545b565b6001600160a01b0390921660209283029190910190910152806114fc81615562565b915050611417565b5063ffffffff8416600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505050565b6115826040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115e6919061538f565b90506116136040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611643908b908990899060040161575a565b600060405180830381865afa158015611660573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261168891908101906157a4565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906116ba908b908b908b90600401615832565b600060405180830381865afa1580156116d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116ff91908101906157a4565b6040820152856001600160401b0381111561171c5761171c6146c2565b60405190808252806020026020018201604052801561174f57816020015b606081526020019060019003908161173a5790505b50606082015260005b60ff8116871115611b8e576000856001600160401b0381111561177d5761177d6146c2565b6040519080825280602002602001820160405280156117a6578160200160208202803683370190505b5083606001518360ff16815181106117c0576117c061545b565b602002602001018190525060005b86811015611a8e5760008c6001600160a01b03166304ec63518a8a858181106117f9576117f961545b565b905060200201358e886000015186815181106118175761181761545b565b60200260200101516040518463ffffffff1660e01b81526004016118549392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611871573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611895919061585b565b90506001600160c01b0381166119395760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610696565b8a8a8560ff1681811061194e5761194e61545b565b6001600160c01b03841692013560f81c9190911c600190811614159050611a7b57856001600160a01b031663dd9846b98a8a858181106119905761199061545b565b905060200201358d8d8860ff168181106119ac576119ac61545b565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611a02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a269190615884565b85606001518560ff1681518110611a3f57611a3f61545b565b60200260200101518481518110611a5857611a5861545b565b63ffffffff9092166020928302919091019091015282611a7781615562565b9350505b5080611a8681615562565b9150506117ce565b506000816001600160401b03811115611aa957611aa96146c2565b604051908082528060200260200182016040528015611ad2578160200160208202803683370190505b50905060005b82811015611b535784606001518460ff1681518110611af957611af961545b565b60200260200101518181518110611b1257611b1261545b565b6020026020010151828281518110611b2c57611b2c61545b565b63ffffffff9092166020928302919091019091015280611b4b81615562565b915050611ad8565b508084606001518460ff1681518110611b6e57611b6e61545b565b602002602001018190525050508080611b86906158a1565b915050611758565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bcf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bf3919061538f565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611c26908b908b908e906004016158c1565b600060405180830381865afa158015611c43573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c6b91908101906157a4565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611cc5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ce991906153f6565b611d055760405162461bcd60e51b815260040161069690615413565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611d769291906158eb565b600060405180830381865afa158015611d93573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611dbb91908101906157a4565b9050600084516001600160401b03811115611dd857611dd86146c2565b604051908082528060200260200182016040528015611e01578160200160208202803683370190505b50905060005b8551811015611f0257866001600160a01b03166304ec6351878381518110611e3157611e3161545b565b602002602001015187868581518110611e4c57611e4c61545b565b60200260200101516040518463ffffffff1660e01b8152600401611e899392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611ea6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eca919061585b565b6001600160c01b0316828281518110611ee557611ee561545b565b602090810291909101015280611efa81615562565b915050611e07565b5095945050505050565b6040805180820190915260608082526020820152600084611f835760405162461bcd60e51b81526020600482015260376024820152600080516020615c1f83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610696565b60408301515185148015611f9b575060a08301515185145b8015611fab575060c08301515185145b8015611fbb575060e08301515185145b6120255760405162461bcd60e51b81526020600482015260416024820152600080516020615c1f83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610696565b8251516020840151511461209d5760405162461bcd60e51b815260206004820152604460248201819052600080516020615c1f833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610696565b4363ffffffff168463ffffffff161061210c5760405162461bcd60e51b815260206004820152603c6024820152600080516020615c1f83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610696565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561214d5761214d6146c2565b604051908082528060200260200182016040528015612176578160200160208202803683370190505b506020820152866001600160401b03811115612194576121946146c2565b6040519080825280602002602001820160405280156121bd578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156121f1576121f16146c2565b60405190808252806020026020018201604052801561221a578160200160208202803683370190505b5081526020860151516001600160401b0381111561223a5761223a6146c2565b604051908082528060200260200182016040528015612263578160200160208202803683370190505b50816020018190525060006123358a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa15801561230c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612330919061593f565b613e47565b905060005b8760200151518110156125b157612360886020015182815181106112b0576112b061545b565b836020015182815181106123765761237661545b565b6020908102919091010152801561243657602083015161239760018361595c565b815181106123a7576123a761545b565b602002602001015160001c836020015182815181106123c8576123c861545b565b602002602001015160001c11612436576040805162461bcd60e51b8152602060048201526024810191909152600080516020615c1f83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610696565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061247b5761247b61545b565b60200260200101518b8b60000151858151811061249a5761249a61545b565b60200260200101516040518463ffffffff1660e01b81526004016124d79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156124f4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612518919061585b565b6001600160c01b0316836000015182815181106125375761253761545b565b60200260200101818152505061259d61090561257184866000015185815181106125635761256361545b565b602002602001015116613eda565b8a6020015184815181106125875761258761545b565b6020026020010151613f0590919063ffffffff16565b9450806125a981615562565b91505061233a565b50506125bc83613fe9565b60975490935060ff166000816125d3576000612655565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612631573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126559190615973565b905060005b8a811015612cd35782156127b5578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106126b1576126b161545b565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156126f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127159190615973565b61271f919061598c565b116127b55760405162461bcd60e51b81526020600482015260666024820152600080516020615c1f83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610696565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106127f6576127f661545b565b9050013560f81c60f81b60f81c8c8c60a00151858151811061281a5761281a61545b565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612876573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061289a91906159a4565b6001600160401b0319166128bd8a6040015183815181106112b0576112b061545b565b67ffffffffffffffff1916146129595760405162461bcd60e51b81526020600482015260616024820152600080516020615c1f83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610696565b612989896040015182815181106129725761297261545b565b602002602001015187613a3f90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106129cc576129cc61545b565b9050013560f81c60f81b60f81c8c8c60c0015185815181106129f0576129f061545b565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612a4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a709190615523565b85602001518281518110612a8657612a8661545b565b6001600160601b03909216602092830291909101820152850151805182908110612ab257612ab261545b565b602002602001015185600001518281518110612ad057612ad061545b565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612cbe57612b4886600001518281518110612b1a57612b1a61545b565b60200260200101518f8f86818110612b3457612b3461545b565b600192013560f81c9290921c811614919050565b15612cac577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612b8e57612b8e61545b565b9050013560f81c60f81b60f81c8e89602001518581518110612bb257612bb261545b565b60200260200101518f60e001518881518110612bd057612bd061545b565b60200260200101518781518110612be957612be961545b565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612c4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c719190615523565b8751805185908110612c8557612c8561545b565b60200260200101818151612c9991906159cf565b6001600160601b03169052506001909101905b80612cb681615562565b915050612af4565b50508080612ccb90615562565b91505061265a565b505050600080612ced8c868a606001518b608001516107ea565b9150915081612d5e5760405162461bcd60e51b81526020600482015260436024820152600080516020615c1f83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610696565b80612dbf5760405162461bcd60e51b81526020600482015260396024820152600080516020615c1f83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610696565b50506000878260200151604051602001612dda929190615712565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612e0c614084565b612e1660006140de565b565b60ce546001600160a01b03163314612e7c5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610696565b60408051608081018252606081830181905263ffffffff8781168352438116602080850191909152908716918301919091528251601f8501829004820281018201909352838352909190849084908190840183828082843760009201919091525050505060408083019190915251612ef89082906020016159f7565b60408051601f19818403018152828252805160209182012063ffffffff8916600081815260ca9093529290912055907ff31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca590612f549084906159f7565b60405180910390a2505060c9805463ffffffff191663ffffffff94909416939093179092555050565b60cd546001600160a01b03163314612fd75760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610696565b6000612fe96040850160208601614895565b9050366000612ffb6040870187615a86565b909250905060006130126080880160608901614895565b905060ca600061302560208a018a614895565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016130519190615acc565b60405160208183030381529060405280519060200120146130da5760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610696565b600060cb816130ec60208b018b614895565b63ffffffff1663ffffffff16815260200190815260200160002054146131695760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610696565b6131937f0000000000000000000000000000000000000000000000000000000000000000856156ea565b63ffffffff164363ffffffff1611156132045760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610696565b6000866040516020016132179190615b48565b60405160208183030381529060405280519060200120905060008061323f8387878a8c611f0c565b9150915060005b8581101561333e578460ff16836020015182815181106132685761326861545b565b602002602001015161327a9190615b5b565b6001600160601b031660438460000151838151811061329b5761329b61545b565b60200260200101516001600160601b03166132b69190615b8a565b101561332c576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610696565b8061333681615562565b915050613246565b5060408051808201825263ffffffff4316815260208082018490529151909161336b918c91849101615ba9565b6040516020818303038152906040528051906020012060cb60008d60000160208101906133989190614895565b63ffffffff1663ffffffff168152602001908152602001600020819055507fb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce35678a826040516133e7929190615ba9565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106134375761343761545b565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061347390889086906004016158eb565b600060405180830381865afa158015613490573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526134b891908101906157a4565b6000815181106134ca576134ca61545b565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613536573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061355a919061585b565b6001600160c01b03169050600061357082614130565b90508161357e8a838a610974565b9550955050505050935093915050565b613596614084565b6001600160a01b0381166135fb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610696565b6106a8816140de565b600054610100900460ff16158080156136245750600054600160ff909116105b8061363e5750303b15801561363e575060005460ff166001145b6136a15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610696565b6000805460ff1916600117905580156136c4576000805461ff0019166101001790555b6136cf8560006141fc565b6136d8846140de565b60cd80546001600160a01b038086166001600160a01b03199283161790925560ce805492851692909116919091179055801561374e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137cc919061538f565b6001600160a01b0316336001600160a01b0316146137fc5760405162461bcd60e51b8152600401610696906153ac565b60665419811960665419161461387a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610696565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107df565b6001600160a01b03811661393f5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610696565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526139c461459d565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156139f7576139f9565bfe5b5080613a375760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610696565b505092915050565b6040805180820190915260008082526020820152613a5b6145bb565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156139f7575080613a375760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610696565b613adb6145d9565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613bc3600080516020615bff83398151915286615471565b90505b613bcf816142e6565b9093509150600080516020615bff833981519152828309831415613c09576040805180820190915290815260208101919091529392505050565b600080516020615bff833981519152600182089050613bc6565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613c556145fe565b60005b6002811015613e1a576000613c6e826006615b8a565b9050848260028110613c8257613c8261545b565b60200201515183613c9483600061598c565b600c8110613ca457613ca461545b565b6020020152848260028110613cbb57613cbb61545b565b60200201516020015183826001613cd2919061598c565b600c8110613ce257613ce261545b565b6020020152838260028110613cf957613cf961545b565b6020020151515183613d0c83600261598c565b600c8110613d1c57613d1c61545b565b6020020152838260028110613d3357613d3361545b565b6020020151516001602002015183613d4c83600361598c565b600c8110613d5c57613d5c61545b565b6020020152838260028110613d7357613d7361545b565b602002015160200151600060028110613d8e57613d8e61545b565b602002015183613d9f83600461598c565b600c8110613daf57613daf61545b565b6020020152838260028110613dc657613dc661545b565b602002015160200151600160028110613de157613de161545b565b602002015183613df283600561598c565b600c8110613e0257613e0261545b565b60200201525080613e1281615562565b915050613c58565b50613e2361461d565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613e5384614368565b9050808360ff166001901b11613ed15760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610696565b90505b92915050565b6000805b8215613ed457613eef60018461595c565b9092169180613efd81615bdc565b915050613ede565b60408051808201909152600080825260208201526102008261ffff1610613f615760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610696565b8161ffff1660011415613f75575081613ed4565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613fde57600161ffff871660ff83161c81161415613fc157613fbe8484613a3f565b93505b613fcb8384613a3f565b92506201fffe600192831b169101613f91565b509195945050505050565b6040805180820190915260008082526020820152815115801561400e57506020820151155b1561402c575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615bff833981519152846020015161405f9190615471565b61407790600080516020615bff83398151915261595c565b905292915050565b919050565b6033546001600160a01b03163314612e165760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610696565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606060008061413e84613eda565b61ffff166001600160401b03811115614159576141596146c2565b6040519080825280601f01601f191660200182016040528015614183576020820181803683370190505b5090506000805b82518210801561419b575061010081105b156141f2576001811b9350858416156141e2578060f81b8383815181106141c4576141c461545b565b60200101906001600160f81b031916908160001a9053508160010191505b6141eb81615562565b905061418a565b5090949350505050565b6065546001600160a01b031615801561421d57506001600160a01b03821615155b61429f5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610696565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26142e2826138b1565b5050565b60008080600080516020615bff8339815191526003600080516020615bff83398151915286600080516020615bff83398151915288890909089050600061435c827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615bff8339815191526144f5565b91959194509092505050565b6000610100825111156143f15760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610696565b81516143ff57506000919050565b600080836000815181106144155761441561545b565b0160200151600160f89190911c81901b92505b84518110156144ec578481815181106144435761444361545b565b0160200151600160f89190911c1b91508282116144d85760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610696565b918117916144e581615562565b9050614428565b50909392505050565b60008061450061461d565b61450861463b565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156139f75750826145925760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610696565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806145ec614659565b81526020016145f9614659565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106a857600080fd5b60006020828403121561469e57600080fd5b8135613ed181614677565b6000602082840312156146bb57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156146fa576146fa6146c2565b60405290565b60405161010081016001600160401b03811182821017156146fa576146fa6146c2565b604051601f8201601f191681016001600160401b038111828210171561474b5761474b6146c2565b604052919050565b60006040828403121561476557600080fd5b61476d6146d8565b9050813581526020820135602082015292915050565b600082601f83011261479457600080fd5b604051604081018181106001600160401b03821117156147b6576147b66146c2565b80604052508060408401858111156147cd57600080fd5b845b81811015613fde5780358352602092830192016147cf565b6000608082840312156147f957600080fd5b6148016146d8565b905061480d8383614783565b815261481c8360408401614783565b602082015292915050565b600080600080610120858703121561483e57600080fd5b8435935061484f8660208701614753565b925061485e86606087016147e7565b915061486d8660e08701614753565b905092959194509250565b63ffffffff811681146106a857600080fd5b803561407f81614878565b6000602082840312156148a757600080fd5b8135613ed181614878565b6000806000606084860312156148c757600080fd5b83356148d281614677565b92506020848101356001600160401b03808211156148ef57600080fd5b818701915087601f83011261490357600080fd5b813581811115614915576149156146c2565b614927601f8201601f19168501614723565b9150808252888482850101111561493d57600080fd5b80848401858401376000848284010152508094505050506149606040850161488a565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156149ff578385038a52825180518087529087019087870190845b818110156149ea57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016149a6565b50509a87019a95505091850191600101614988565b509298975050505050505050565b602081526000614a206020830184614969565b9392505050565b80151581146106a857600080fd5b600060208284031215614a4757600080fd5b8135613ed181614a27565b600060808284031215614a6457600080fd5b50919050565b600060408284031215614a6457600080fd5b60006001600160401b03821115614a9557614a956146c2565b5060051b60200190565b600082601f830112614ab057600080fd5b81356020614ac5614ac083614a7c565b614723565b82815260069290921b84018101918181019086841115614ae457600080fd5b8286015b84811015614b0857614afa8882614753565b835291830191604001614ae8565b509695505050505050565b60008060008060a08587031215614b2957600080fd5b84356001600160401b0380821115614b4057600080fd5b614b4c88838901614a52565b95506020870135915080821115614b6257600080fd5b614b6e88838901614a6a565b9450614b7d8860408901614a6a565b93506080870135915080821115614b9357600080fd5b50614ba087828801614a9f565b91505092959194509250565b60008083601f840112614bbe57600080fd5b5081356001600160401b03811115614bd557600080fd5b602083019150836020828501011115614bed57600080fd5b9250929050565b60008060008060008060808789031215614c0d57600080fd5b8635614c1881614677565b95506020870135614c2881614878565b945060408701356001600160401b0380821115614c4457600080fd5b614c508a838b01614bac565b90965094506060890135915080821115614c6957600080fd5b818901915089601f830112614c7d57600080fd5b813581811115614c8c57600080fd5b8a60208260051b8501011115614ca157600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614ced57815163ffffffff1687529582019590820190600101614ccb565b509495945050505050565b600060208083528351608082850152614d1460a0850182614cb7565b905081850151601f1980868403016040870152614d318383614cb7565b92506040870151915080868403016060870152614d4e8383614cb7565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614da55784878303018452614d93828751614cb7565b95880195938801939150600101614d79565b509998505050505050505050565b60ff811681146106a857600080fd5b600060208284031215614dd457600080fd5b8135613ed181614db3565b600080600060608486031215614df457600080fd5b8335614dff81614677565b92506020848101356001600160401b03811115614e1b57600080fd5b8501601f81018713614e2c57600080fd5b8035614e3a614ac082614a7c565b81815260059190911b82018301908381019089831115614e5957600080fd5b928401925b82841015614e7757833582529284019290840190614e5e565b80965050505050506149606040850161488a565b6020808252825182820181905260009190848201906040850190845b81811015614ec357835183529284019291840191600101614ea7565b50909695505050505050565b600082601f830112614ee057600080fd5b81356020614ef0614ac083614a7c565b82815260059290921b84018101918181019086841115614f0f57600080fd5b8286015b84811015614b08578035614f2681614878565b8352918301918301614f13565b600082601f830112614f4457600080fd5b81356020614f54614ac083614a7c565b82815260059290921b84018101918181019086841115614f7357600080fd5b8286015b84811015614b085780356001600160401b03811115614f965760008081fd5b614fa48986838b0101614ecf565b845250918301918301614f77565b60006101808284031215614fc557600080fd5b614fcd614700565b905081356001600160401b0380821115614fe657600080fd5b614ff285838601614ecf565b8352602084013591508082111561500857600080fd5b61501485838601614a9f565b6020840152604084013591508082111561502d57600080fd5b61503985838601614a9f565b604084015261504b85606086016147e7565b606084015261505d8560e08601614753565b608084015261012084013591508082111561507757600080fd5b61508385838601614ecf565b60a084015261014084013591508082111561509d57600080fd5b6150a985838601614ecf565b60c08401526101608401359150808211156150c357600080fd5b506150d084828501614f33565b60e08301525092915050565b6000806000806000608086880312156150f457600080fd5b8535945060208601356001600160401b038082111561511257600080fd5b61511e89838a01614bac565b90965094506040880135915061513382614878565b9092506060870135908082111561514957600080fd5b5061515688828901614fb2565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614ced5781516001600160601b031687529582019590820190600101615177565b60408152600083516040808401526151b76080840182615163565b90506020850151603f198483030160608501526151d48282615163565b925050508260208301529392505050565b600080600080606085870312156151fb57600080fd5b843561520681614878565b9350602085013561521681614878565b925060408501356001600160401b0381111561523157600080fd5b61523d87828801614bac565b95989497509550505050565b60008060006060848603121561525e57600080fd5b83356001600160401b038082111561527557600080fd5b61528187838801614a52565b9450602086013591508082111561529757600080fd5b6152a387838801614a6a565b935060408601359150808211156152b957600080fd5b506152c686828701614fb2565b9150509250925092565b6000806000606084860312156152e557600080fd5b83356152f081614677565b925060208401359150604084013561530781614878565b809150509250925092565b82815260406020820152600061532b6040830184614969565b949350505050565b6000806000806080858703121561534957600080fd5b843561535481614677565b9350602085013561536481614677565b9250604085013561537481614677565b9150606085013561538481614677565b939692955090935050565b6000602082840312156153a157600080fd5b8151613ed181614677565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561540857600080fd5b8151613ed181614a27565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261548e57634e487b7160e01b600052601260045260246000fd5b500690565b600060208083850312156154a657600080fd5b82516001600160401b038111156154bc57600080fd5b8301601f810185136154cd57600080fd5b80516154db614ac082614a7c565b81815260059190911b820183019083810190878311156154fa57600080fd5b928401925b82841015615518578351825292840192908401906154ff565b979650505050505050565b60006020828403121561553557600080fd5b81516001600160601b0381168114613ed157600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156155765761557661554c565b5060010190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e198436030181126155bd57600080fd5b83016020810192503590506001600160401b038111156155dc57600080fd5b803603831315614bed57600080fd5b60006040830182356155fc81614878565b63ffffffff16845260208381013536859003601e1901811261561d57600080fd5b840181810190356001600160401b0381111561563857600080fd5b8060051b80360387131561564b57600080fd5b6040848901529381905260609387018401938290880160005b8381101561569e57898703605f1901825261567f83866155a6565b61568a89828461557d565b985050509185019190850190600101615664565b509498975050505050505050565b6060815260006156bf60608301856155eb565b905082356156cc81614878565b63ffffffff8116602084015250602083013560408301529392505050565b600063ffffffff8083168185168083038211156157095761570961554c565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561574d57815185529382019390820190600101615731565b5092979650505050505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561578757600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156157b757600080fd5b82516001600160401b038111156157cd57600080fd5b8301601f810185136157de57600080fd5b80516157ec614ac082614a7c565b81815260059190911b8201830190838101908783111561580b57600080fd5b928401925b8284101561551857835161582381614878565b82529284019290840190615810565b63ffffffff8416815260406020820152600061585260408301848661557d565b95945050505050565b60006020828403121561586d57600080fd5b81516001600160c01b0381168114613ed157600080fd5b60006020828403121561589657600080fd5b8151613ed181614878565b600060ff821660ff8114156158b8576158b861554c565b60010192915050565b6040815260006158d560408301858761557d565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561593257845183529383019391830191600101615916565b5090979650505050505050565b60006020828403121561595157600080fd5b8151613ed181614db3565b60008282101561596e5761596e61554c565b500390565b60006020828403121561598557600080fd5b5051919050565b6000821982111561599f5761599f61554c565b500190565b6000602082840312156159b657600080fd5b815167ffffffffffffffff1981168114613ed157600080fd5b60006001600160601b03838116908316818110156159ef576159ef61554c565b039392505050565b6000602080835263ffffffff808551168285015280828601511660408501525060408401516080606085015280518060a086015260005b81811015615a4a5782810184015186820160c001528301615a2e565b81811115615a5c57600060c083880101525b50606086015163ffffffff811660808701529250601f01601f19169390930160c001949350505050565b6000808335601e19843603018112615a9d57600080fd5b8301803591506001600160401b03821115615ab757600080fd5b602001915036819003821315614bed57600080fd5b6020815260008235615add81614878565b63ffffffff808216602085015260208501359150615afa82614878565b8082166040850152615b0f60408601866155a6565b925060806060860152615b2660a08601848361557d565b9250506060850135615b3781614878565b166080939093019290925250919050565b602081526000614a2060208301846155eb565b60006001600160601b0380831681851681830481118215151615615b8157615b8161554c565b02949350505050565b6000816000190483118215151615615ba457615ba461554c565b500290565b606081526000615bbc60608301856155eb565b905063ffffffff8351166020830152602083015160408301529392505050565b600061ffff80831681811415615bf457615bf461554c565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212200256cec39de1e7f5b242cf5971b354edfe0ebde06b4144c9e4bd3df247c5efa164736f6c634300080c0033",
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
