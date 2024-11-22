// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractZrTaskManager

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

// IZRTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IZRTaskManagerTask struct {
	TaskId                    uint32
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IZRTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IZRTaskManagerTaskResponse struct {
	ReferenceTaskId    uint32
	InactiveSetZRChain []string
}

// IZRTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IZRTaskManagerTaskResponseMetadata struct {
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

// ContractZrTaskManagerMetaData contains all meta data concerning the ContractZrTaskManager contract.
var ContractZrTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_zrServiceManager\",\"type\":\"address\",\"internalType\":\"contractIZrServiceManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIZRTaskManager.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIZRTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zrServiceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIZrServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIZRTaskManager.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIZRTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIZRTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005d4e38038062005d4e8339810160408190526200003591620001ea565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b5919062000231565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000133919062000231565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b3919062000231565b6001600160a01b031660e0525063ffffffff16610100525062000258565b6001600160a01b0381168114620001e757600080fd5b50565b60008060408385031215620001fe57600080fd5b82516200020b81620001d1565b602084015190925063ffffffff811681146200022657600080fd5b809150509250929050565b6000602082840312156200024457600080fd5b81516200025181620001d1565b9392505050565b60805160a05160c05160e05161010051615a6b620002e3600039600081816102a6015281816106050152612f610152600081816105ce01526123c801526000818161048401526125aa0152600081816104ab0152818161278001526129420152600081816104d2015281816110dd015281816120930152818161222b01526124650152615a6b6000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c80635df45946116101305780638da5cb5b116100b8578063df5cf7231161007c578063df5cf723146105c9578063f2fde38b146105f0578063f5c9899d14610603578063f63c5bab14610629578063fabc1cbc1461063257600080fd5b80638da5cb5b146105645780639d3a0f2d14610575578063b98d090814610588578063caf73aa014610595578063cefdc1d4146105a857600080fd5b8063715018a6116100ff578063715018a61461051557806372d18e8d1461051d5780637afa1eed1461052b57806387712db51461053e578063886f11951461055157600080fd5b80635df459461461047f57806368304835146104a65780636d14a987146104cd5780636efb4636146104f457600080fd5b806331b36bd9116101be578063595c6a6711610182578063595c6a67146103f95780635ac86ab7146104015780635c155662146104345780635c975abb146104545780635decc3f51461045c57600080fd5b806331b36bd9146103665780633563b0d114610386578063416c7e5e146103a65780634d2b57fe146103b95780634f739f74146103d957600080fd5b80631ad43189116102055780631ad43189146102a1578063245a7bfc146102dd5780632cb223d5146103085780632d3de569146103365780632d89f6fc1461034657600080fd5b806310d67a2f14610237578063136439dd1461024c5780631459457a1461025f578063171f1d5b14610272575b600080fd5b61024a6102453660046143f7565b610645565b005b61024a61025a366004614414565b610701565b61024a61026d36600461442d565b610840565b6102856102803660046145ef565b61099f565b6040805192151583529015156020830152015b60405180910390f35b6102c87f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610298565b60cd546102f0906001600160a01b031681565b6040516001600160a01b039091168152602001610298565b61032861031636600461465d565b60cb6020526000908152604090205481565b604051908152602001610298565b60c9546102c89063ffffffff1681565b61032861035436600461465d565b60ca6020526000908152604090205481565b61037961037436600461469d565b610b29565b604051610298919061478b565b6103996103943660046147a5565b610c45565b6040516102989190614900565b61024a6103b4366004614921565b6110db565b6103cc6103c73660046149a4565b611212565b60405161029891906149f3565b6103ec6103e7366004614a88565b611327565b6040516102989190614b81565b61024a611a4f565b61042461040f366004614c4b565b606654600160ff9092169190911b9081161490565b6040519015158152602001610298565b610447610442366004614c68565b611b16565b6040516102989190614ccb565b606654610328565b61042461046a36600461465d565b60cc6020526000908152604090205460ff1681565b6102f07f000000000000000000000000000000000000000000000000000000000000000081565b6102f07f000000000000000000000000000000000000000000000000000000000000000081565b6102f07f000000000000000000000000000000000000000000000000000000000000000081565b610507610502366004614f74565b611cde565b604051610298929190615034565b61024a612bf7565b60c95463ffffffff166102c8565b60ce546102f0906001600160a01b031681565b60cf546102f0906001600160a01b031681565b6065546102f0906001600160a01b031681565b6033546001600160a01b03166102f0565b61024a61058336600461507d565b612c0b565b6097546104249060ff1681565b61024a6105a33660046150e1565b612d70565b6105bb6105b6366004615178565b613274565b6040516102989291906151af565b6102f07f000000000000000000000000000000000000000000000000000000000000000081565b61024a6105fe3660046143f7565b613406565b7f00000000000000000000000000000000000000000000000000000000000000006102c8565b6102c86103e881565b61024a610640366004614414565b61347c565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610698573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106bc91906151d0565b6001600160a01b0316336001600160a01b0316146106f55760405162461bcd60e51b81526004016106ec906151ed565b60405180910390fd5b6106fe816135d8565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076d9190615237565b6107895760405162461bcd60e51b81526004016106ec90615254565b606654818116146108025760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016106ec565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156108605750600054600160ff909116105b8061087a5750303b15801561087a575060005460ff166001145b6108dd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106ec565b6000805460ff191660011790558015610900576000805461ff0019166101001790555b61090b8660006136cf565b610914856137b9565b60cd80546001600160a01b038087166001600160a01b03199283161790925560ce805486841690831617905560cf8054928516929091169190911790558015610997576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106109e7576109e761529c565b60200201518951600160200201518a60200151600060028110610a0c57610a0c61529c565b60200201518b60200151600160028110610a2857610a2861529c565b602090810291909101518c518d830151604051610a859a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610aa891906152b2565b9050610b1b610ac1610aba888461380b565b869061389c565b610ac9613931565b610b11610b0285610afc604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061380b565b610b0b8c6139f1565b9061389c565b886201d4c0613a80565b909890975095505050505050565b606081516001600160401b03811115610b4457610b4461449e565b604051908082528060200260200182016040528015610b6d578160200160208202803683370190505b50905060005b8251811015610c3e57836001600160a01b03166313542a4e848381518110610b9d57610b9d61529c565b60200260200101516040518263ffffffff1660e01b8152600401610bd091906001600160a01b0391909116815260200190565b602060405180830381865afa158015610bed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1191906152d4565b828281518110610c2357610c2361529c565b6020908102919091010152610c3781615303565b9050610b73565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cab91906151d0565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ced573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1191906151d0565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7791906151d0565b9050600086516001600160401b03811115610d9457610d9461449e565b604051908082528060200260200182016040528015610dc757816020015b6060815260200190600190039081610db25790505b50905060005b87518110156110cf576000888281518110610dea57610dea61529c565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610e4b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e73919081019061531c565b905080516001600160401b03811115610e8e57610e8e61449e565b604051908082528060200260200182016040528015610ed957816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610eac5790505b50848481518110610eec57610eec61529c565b602002602001018190525060005b81518110156110b9576040518060600160405280876001600160a01b03166347b314e8858581518110610f2f57610f2f61529c565b60200260200101516040518263ffffffff1660e01b8152600401610f5591815260200190565b602060405180830381865afa158015610f72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f9691906151d0565b6001600160a01b03168152602001838381518110610fb657610fb661529c565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610fe457610fe461529c565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611040573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061106491906153ac565b6001600160601b03168152508585815181106110825761108261529c565b6020026020010151828151811061109b5761109b61529c565b602002602001018190525080806110b190615303565b915050610efa565b50505080806110c790615303565b915050610dcd565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611139573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115d91906151d0565b6001600160a01b0316336001600160a01b0316146112095760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016106ec565b6106fe81613ca4565b606081516001600160401b0381111561122d5761122d61449e565b604051908082528060200260200182016040528015611256578160200160208202803683370190505b50905060005b8251811015610c3e57836001600160a01b031663296bb0648483815181106112865761128661529c565b60200260200101516040518263ffffffff1660e01b81526004016112ac91815260200190565b602060405180830381865afa1580156112c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ed91906151d0565b8282815181106112ff576112ff61529c565b6001600160a01b039092166020928302919091019091015261132081615303565b905061125c565b6113526040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611392573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113b691906151d0565b90506113e36040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611413908b90899089906004016153d5565b600060405180830381865afa158015611430573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611458919081019061541c565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061148a908b908b908b906004016154d3565b600060405180830381865afa1580156114a7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114cf919081019061541c565b6040820152856001600160401b038111156114ec576114ec61449e565b60405190808252806020026020018201604052801561151f57816020015b606081526020019060019003908161150a5790505b50606082015260005b60ff8116871115611960576000856001600160401b0381111561154d5761154d61449e565b604051908082528060200260200182016040528015611576578160200160208202803683370190505b5083606001518360ff16815181106115905761159061529c565b602002602001018190525060005b868110156118605760008c6001600160a01b03166304ec63518a8a858181106115c9576115c961529c565b905060200201358e886000015186815181106115e7576115e761529c565b60200260200101516040518463ffffffff1660e01b81526004016116249392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611641573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061166591906154fc565b9050806001600160c01b031660000361170c5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a4016106ec565b8a8a8560ff168181106117215761172161529c565b60016001600160c01b038516919093013560f81c1c8216909103905061184d57856001600160a01b031663dd9846b98a8a858181106117625761176261529c565b905060200201358d8d8860ff1681811061177e5761177e61529c565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156117d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f89190615525565b85606001518560ff16815181106118115761181161529c565b6020026020010151848151811061182a5761182a61529c565b63ffffffff909216602092830291909101909101528261184981615303565b9350505b508061185881615303565b91505061159e565b506000816001600160401b0381111561187b5761187b61449e565b6040519080825280602002602001820160405280156118a4578160200160208202803683370190505b50905060005b828110156119255784606001518460ff16815181106118cb576118cb61529c565b602002602001015181815181106118e4576118e461529c565b60200260200101518282815181106118fe576118fe61529c565b63ffffffff909216602092830291909101909101528061191d81615303565b9150506118aa565b508084606001518460ff16815181106119405761194061529c565b60200260200101819052505050808061195890615542565b915050611528565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c591906151d0565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906119f8908b908b908e90600401615561565b600060405180830381865afa158015611a15573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a3d919081019061541c565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611a97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611abb9190615237565b611ad75760405162461bcd60e51b81526004016106ec90615254565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611b4892919061558b565b600060405180830381865afa158015611b65573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b8d919081019061541c565b9050600084516001600160401b03811115611baa57611baa61449e565b604051908082528060200260200182016040528015611bd3578160200160208202803683370190505b50905060005b8551811015611cd457866001600160a01b03166304ec6351878381518110611c0357611c0361529c565b602002602001015187868581518110611c1e57611c1e61529c565b60200260200101516040518463ffffffff1660e01b8152600401611c5b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611c78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c9c91906154fc565b6001600160c01b0316828281518110611cb757611cb761529c565b602090810291909101015280611ccc81615303565b915050611bd9565b5095945050505050565b60408051808201909152606080825260208201526000848103611d575760405162461bcd60e51b81526020600482015260376024820152600080516020615a1683398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016106ec565b60408301515185148015611d6f575060a08301515185145b8015611d7f575060c08301515185145b8015611d8f575060e08301515185145b611df95760405162461bcd60e51b81526020600482015260416024820152600080516020615a1683398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016106ec565b82515160208401515114611e715760405162461bcd60e51b815260206004820152604460248201819052600080516020615a16833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016106ec565b4363ffffffff168463ffffffff1610611ee05760405162461bcd60e51b815260206004820152603c6024820152600080516020615a1683398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016106ec565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611f2157611f2161449e565b604051908082528060200260200182016040528015611f4a578160200160208202803683370190505b506020820152866001600160401b03811115611f6857611f6861449e565b604051908082528060200260200182016040528015611f91578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611fc557611fc561449e565b604051908082528060200260200182016040528015611fee578160200160208202803683370190505b5081526020860151516001600160401b0381111561200e5761200e61449e565b604051908082528060200260200182016040528015612037578160200160208202803683370190505b50816020018190525060006121098a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156120e0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061210491906155aa565b613ceb565b905060005b8760200151518110156123a457612153886020015182815181106121345761213461529c565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106121695761216961529c565b6020908102919091010152801561222957602083015161218a6001836155c7565b8151811061219a5761219a61529c565b602002602001015160001c836020015182815181106121bb576121bb61529c565b602002602001015160001c11612229576040805162461bcd60e51b8152602060048201526024810191909152600080516020615a1683398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016106ec565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061226e5761226e61529c565b60200260200101518b8b60000151858151811061228d5761228d61529c565b60200260200101516040518463ffffffff1660e01b81526004016122ca9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156122e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061230b91906154fc565b6001600160c01b03168360000151828151811061232a5761232a61529c565b602002602001018181525050612390610aba61236484866000015185815181106123565761235661529c565b602002602001015116613d7e565b8a60200151848151811061237a5761237a61529c565b6020026020010151613da990919063ffffffff16565b94508061239c81615303565b91505061210e565b50506123af83613e8c565b60975490935060ff166000816123c6576000612448565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612424573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061244891906152d4565b905060005b8a811015612ac65782156125a8578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106124a4576124a461529c565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156124e4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061250891906152d4565b61251291906155da565b116125a85760405162461bcd60e51b81526020600482015260666024820152600080516020615a1683398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016106ec565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106125e9576125e961529c565b9050013560f81c60f81b60f81c8c8c60a00151858151811061260d5761260d61529c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612669573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061268d91906155ed565b6001600160401b0319166126b08a6040015183815181106121345761213461529c565b67ffffffffffffffff19161461274c5760405162461bcd60e51b81526020600482015260616024820152600080516020615a1683398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016106ec565b61277c896040015182815181106127655761276561529c565b60200260200101518761389c90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106127bf576127bf61529c565b9050013560f81c60f81b60f81c8c8c60c0015185815181106127e3576127e361529c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561283f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061286391906153ac565b856020015182815181106128795761287961529c565b6001600160601b039092166020928302919091018201528501518051829081106128a5576128a561529c565b6020026020010151856000015182815181106128c3576128c361529c565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612ab15761293b8660000151828151811061290d5761290d61529c565b60200260200101518f8f868181106129275761292761529c565b600192013560f81c9290921c811614919050565b15612a9f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106129815761298161529c565b9050013560f81c60f81b60f81c8e896020015185815181106129a5576129a561529c565b60200260200101518f60e0015188815181106129c3576129c361529c565b602002602001015187815181106129dc576129dc61529c565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612a40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a6491906153ac565b8751805185908110612a7857612a7861529c565b60200260200101818151612a8c9190615618565b6001600160601b03169052506001909101905b80612aa981615303565b9150506128e7565b50508080612abe90615303565b91505061244d565b505050600080612ae08c868a606001518b6080015161099f565b9150915081612b515760405162461bcd60e51b81526020600482015260436024820152600080516020615a1683398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016106ec565b80612bb25760405162461bcd60e51b81526020600482015260396024820152600080516020615a1683398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016106ec565b50506000878260200151604051602001612bcd929190615638565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612bff613f27565b612c0960006137b9565b565b60ce546001600160a01b03163314612c6f5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084016106ec565b60408051608081018252606081830181905263ffffffff8781168352438116602080850191909152908716918301919091528251601f8501829004820281018201909352838352909190849084908190840183828082843760009201919091525050505060408083019190915251612ceb908290602001615680565b60408051601f19818403018152828252805160209182012063ffffffff8916600081815260ca9093529290912055907ff31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca590612d47908490615680565b60405180910390a2505060c9805463ffffffff191663ffffffff94909416939093179092555050565b60cd546001600160a01b03163314612dca5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016106ec565b6000612ddc604085016020860161465d565b9050366000612dee6040870187615706565b90925090506000612e05608088016060890161465d565b905060ca6000612e1860208a018a61465d565b63ffffffff1663ffffffff1681526020019081526020016000205487604051602001612e449190615791565b6040516020818303038152906040528051906020012014612ecd5760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e747261637400000060648201526084016106ec565b600060cb81612edf60208b018b61465d565b63ffffffff1663ffffffff1681526020019081526020016000205414612f5c5760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016106ec565b612f867f00000000000000000000000000000000000000000000000000000000000000008561580d565b63ffffffff164363ffffffff161115612ff75760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b60648201526084016106ec565b60008660405160200161300a91906158f7565b6040516020818303038152906040528051906020012090506000806130328387878a8c611cde565b9150915060005b85811015613131578460ff168360200151828151811061305b5761305b61529c565b602002602001015161306d919061590a565b6001600160601b031660438460000151838151811061308e5761308e61529c565b60200260200101516001600160601b03166130a9919061592d565b101561311f576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d60648201526084016106ec565b8061312981615303565b915050613039565b50600061314160208b018b615944565b905011156131b75760cf546001600160a01b031663c63e3c5061316760208c018c615944565b6040518363ffffffff1660e01b815260040161318492919061598d565b600060405180830381600087803b15801561319e57600080fd5b505af11580156131b2573d6000803e3d6000fd5b505050505b60408051808201825263ffffffff431681526020808201849052915190916131e3918c918491016159a1565b6040516020818303038152906040528051906020012060cb60008d6000016020810190613210919061465d565b63ffffffff1663ffffffff168152602001908152602001600020819055507fb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce35678a8260405161325f9291906159a1565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106132af576132af61529c565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906132eb908890869060040161558b565b600060405180830381865afa158015613308573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613330919081019061541c565b6000815181106133425761334261529c565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156133ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133d291906154fc565b6001600160c01b0316905060006133e882613f81565b9050816133f68a838a610c45565b9550955050505050935093915050565b61340e613f27565b6001600160a01b0381166134735760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106ec565b6106fe816137b9565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134f391906151d0565b6001600160a01b0316336001600160a01b0316146135235760405162461bcd60e51b81526004016106ec906151ed565b6066541981196066541916146135a15760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016106ec565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610835565b6001600160a01b0381166136665760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106ec565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b03161580156136f057506001600160a01b03821615155b6137725760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106ec565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26137b5826135d8565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613827614308565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808061385657fe5b50806138945760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016106ec565b505092915050565b60408051808201909152600080825260208201526138b8614326565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080806138f357fe5b50806138945760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016106ec565b613939614344565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613a216000805160206159f6833981519152866152b2565b90505b613a2d8161404d565b90935091506000805160206159f68339815191528283098303613a66576040805180820190915290815260208101919091529392505050565b6000805160206159f6833981519152600182089050613a24565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613ab2614369565b60005b6002811015613c77576000613acb82600661592d565b9050848260028110613adf57613adf61529c565b60200201515183613af18360006155da565b600c8110613b0157613b0161529c565b6020020152848260028110613b1857613b1861529c565b60200201516020015183826001613b2f91906155da565b600c8110613b3f57613b3f61529c565b6020020152838260028110613b5657613b5661529c565b6020020151515183613b698360026155da565b600c8110613b7957613b7961529c565b6020020152838260028110613b9057613b9061529c565b6020020151516001602002015183613ba98360036155da565b600c8110613bb957613bb961529c565b6020020152838260028110613bd057613bd061529c565b602002015160200151600060028110613beb57613beb61529c565b602002015183613bfc8360046155da565b600c8110613c0c57613c0c61529c565b6020020152838260028110613c2357613c2361529c565b602002015160200151600160028110613c3e57613c3e61529c565b602002015183613c4f8360056155da565b600c8110613c5f57613c5f61529c565b60200201525080613c6f81615303565b915050613ab5565b50613c80614388565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b600080613cf7846140cf565b9050808360ff166001901b11613d755760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016106ec565b90505b92915050565b6000805b8215613d7857613d936001846155c7565b9092169180613da1816159d4565b915050613d82565b60408051808201909152600080825260208201526102008261ffff1610613e055760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016106ec565b8161ffff16600103613e18575081613d78565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613e8157600161ffff871660ff83161c81169003613e6457613e61848461389c565b93505b613e6e838461389c565b92506201fffe600192831b169101613e34565b509195945050505050565b60408051808201909152600080825260208201528151158015613eb157506020820151155b15613ecf575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206159f68339815191528460200151613f0291906152b2565b613f1a906000805160206159f68339815191526155c7565b905292915050565b919050565b6033546001600160a01b03163314612c095760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106ec565b6060600080613f8f84613d7e565b61ffff166001600160401b03811115613faa57613faa61449e565b6040519080825280601f01601f191660200182016040528015613fd4576020820181803683370190505b5090506000805b825182108015613fec575061010081105b15614043576001811b935085841615614033578060f81b8383815181106140155761401561529c565b60200101906001600160f81b031916908160001a9053508160010191505b61403c81615303565b9050613fdb565b5090949350505050565b600080806000805160206159f683398151915260036000805160206159f6833981519152866000805160206159f68339815191528889090908905060006140c3827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206159f683398151915261425f565b91959194509092505050565b6000610100825111156141585760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016106ec565b815160000361416957506000919050565b6000808360008151811061417f5761417f61529c565b0160200151600160f89190911c81901b92505b8451811015614256578481815181106141ad576141ad61529c565b0160200151600160f89190911c1b91508282116142425760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016106ec565b9181179161424f81615303565b9050614192565b50909392505050565b60008061426a614388565b6142726143a6565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082806142af57fe5b50826142fd5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016106ec565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806143576143c4565b81526020016143646143c4565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106fe57600080fd5b60006020828403121561440957600080fd5b8135613d75816143e2565b60006020828403121561442657600080fd5b5035919050565b600080600080600060a0868803121561444557600080fd5b8535614450816143e2565b94506020860135614460816143e2565b93506040860135614470816143e2565b92506060860135614480816143e2565b91506080860135614490816143e2565b809150509295509295909350565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156144d6576144d661449e565b60405290565b60405161010081016001600160401b03811182821017156144d6576144d661449e565b604051601f8201601f191681016001600160401b03811182821017156145275761452761449e565b604052919050565b60006040828403121561454157600080fd5b6145496144b4565b9050813581526020820135602082015292915050565b600082601f83011261457057600080fd5b6145786144b4565b80604084018581111561458a57600080fd5b845b818110156145a457803584526020938401930161458c565b509095945050505050565b6000608082840312156145c157600080fd5b6145c96144b4565b90506145d5838361455f565b81526145e4836040840161455f565b602082015292915050565b600080600080610120858703121561460657600080fd5b84359350614617866020870161452f565b925061462686606087016145af565b91506146358660e0870161452f565b905092959194509250565b63ffffffff811681146106fe57600080fd5b8035613f2281614640565b60006020828403121561466f57600080fd5b8135613d7581614640565b60006001600160401b038211156146935761469361449e565b5060051b60200190565b600080604083850312156146b057600080fd5b82356146bb816143e2565b91506020838101356001600160401b038111156146d757600080fd5b8401601f810186136146e857600080fd5b80356146fb6146f68261467a565b6144ff565b81815260059190911b8201830190838101908883111561471a57600080fd5b928401925b82841015614741578335614732816143e2565b8252928401929084019061471f565b80955050505050509250929050565b600081518084526020808501945080840160005b8381101561478057815187529582019590820190600101614764565b509495945050505050565b60208152600061479e6020830184614750565b9392505050565b6000806000606084860312156147ba57600080fd5b83356147c5816143e2565b92506020848101356001600160401b03808211156147e257600080fd5b818701915087601f8301126147f657600080fd5b8135818111156148085761480861449e565b61481a601f8201601f191685016144ff565b9150808252888482850101111561483057600080fd5b808484018584013760008482840101525080945050505061485360408501614652565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156148f2578385038a52825180518087529087019087870190845b818110156148dd57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614899565b50509a87019a9550509185019160010161487b565b509298975050505050505050565b60208152600061479e602083018461485c565b80151581146106fe57600080fd5b60006020828403121561493357600080fd5b8135613d7581614913565b600082601f83011261494f57600080fd5b8135602061495f6146f68361467a565b82815260059290921b8401810191818101908684111561497e57600080fd5b8286015b848110156149995780358352918301918301614982565b509695505050505050565b600080604083850312156149b757600080fd5b82356149c2816143e2565b915060208301356001600160401b038111156149dd57600080fd5b6149e98582860161493e565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614a345783516001600160a01b031683529284019291840191600101614a0f565b50909695505050505050565b60008083601f840112614a5257600080fd5b5081356001600160401b03811115614a6957600080fd5b602083019150836020828501011115614a8157600080fd5b9250929050565b60008060008060008060808789031215614aa157600080fd5b8635614aac816143e2565b95506020870135614abc81614640565b945060408701356001600160401b0380821115614ad857600080fd5b614ae48a838b01614a40565b90965094506060890135915080821115614afd57600080fd5b818901915089601f830112614b1157600080fd5b813581811115614b2057600080fd5b8a60208260051b8501011115614b3557600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561478057815163ffffffff1687529582019590820190600101614b5f565b600060208083528351608082850152614b9d60a0850182614b4b565b905081850151601f1980868403016040870152614bba8383614b4b565b92506040870151915080868403016060870152614bd78383614b4b565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614c2e5784878303018452614c1c828751614b4b565b95880195938801939150600101614c02565b509998505050505050505050565b60ff811681146106fe57600080fd5b600060208284031215614c5d57600080fd5b8135613d7581614c3c565b600080600060608486031215614c7d57600080fd5b8335614c88816143e2565b925060208401356001600160401b03811115614ca357600080fd5b614caf8682870161493e565b9250506040840135614cc081614640565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614a3457835183529284019291840191600101614ce7565b600082601f830112614d1457600080fd5b81356020614d246146f68361467a565b82815260059290921b84018101918181019086841115614d4357600080fd5b8286015b84811015614999578035614d5a81614640565b8352918301918301614d47565b600082601f830112614d7857600080fd5b81356020614d886146f68361467a565b82815260069290921b84018101918181019086841115614da757600080fd5b8286015b8481101561499957614dbd888261452f565b835291830191604001614dab565b600082601f830112614ddc57600080fd5b81356020614dec6146f68361467a565b82815260059290921b84018101918181019086841115614e0b57600080fd5b8286015b848110156149995780356001600160401b03811115614e2e5760008081fd5b614e3c8986838b0101614d03565b845250918301918301614e0f565b60006101808284031215614e5d57600080fd5b614e656144dc565b905081356001600160401b0380821115614e7e57600080fd5b614e8a85838601614d03565b83526020840135915080821115614ea057600080fd5b614eac85838601614d67565b60208401526040840135915080821115614ec557600080fd5b614ed185838601614d67565b6040840152614ee385606086016145af565b6060840152614ef58560e0860161452f565b6080840152610120840135915080821115614f0f57600080fd5b614f1b85838601614d03565b60a0840152610140840135915080821115614f3557600080fd5b614f4185838601614d03565b60c0840152610160840135915080821115614f5b57600080fd5b50614f6884828501614dcb565b60e08301525092915050565b600080600080600060808688031215614f8c57600080fd5b8535945060208601356001600160401b0380821115614faa57600080fd5b614fb689838a01614a40565b909650945060408801359150614fcb82614640565b90925060608701359080821115614fe157600080fd5b50614fee88828901614e4a565b9150509295509295909350565b600081518084526020808501945080840160005b838110156147805781516001600160601b03168752958201959082019060010161500f565b604081526000835160408084015261504f6080840182614ffb565b90506020850151603f1984830301606085015261506c8282614ffb565b925050508260208301529392505050565b6000806000806060858703121561509357600080fd5b843561509e81614640565b935060208501356150ae81614640565b925060408501356001600160401b038111156150c957600080fd5b6150d587828801614a40565b95989497509550505050565b6000806000606084860312156150f657600080fd5b83356001600160401b038082111561510d57600080fd5b908501906080828803121561512157600080fd5b9093506020850135908082111561513757600080fd5b908501906040828803121561514b57600080fd5b9092506040850135908082111561516157600080fd5b5061516e86828701614e4a565b9150509250925092565b60008060006060848603121561518d57600080fd5b8335615198816143e2565b9250602084013591506040840135614cc081614640565b8281526040602082015260006151c8604083018461485c565b949350505050565b6000602082840312156151e257600080fd5b8151613d75816143e2565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561524957600080fd5b8151613d7581614913565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826152cf57634e487b7160e01b600052601260045260246000fd5b500690565b6000602082840312156152e657600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600060018201615315576153156152ed565b5060010190565b6000602080838503121561532f57600080fd5b82516001600160401b0381111561534557600080fd5b8301601f8101851361535657600080fd5b80516153646146f68261467a565b81815260059190911b8201830190838101908783111561538357600080fd5b928401925b828410156153a157835182529284019290840190615388565b979650505050505050565b6000602082840312156153be57600080fd5b81516001600160601b0381168114613d7557600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561540257600080fd5b8260051b8085606085013791909101606001949350505050565b6000602080838503121561542f57600080fd5b82516001600160401b0381111561544557600080fd5b8301601f8101851361545657600080fd5b80516154646146f68261467a565b81815260059190911b8201830190838101908783111561548357600080fd5b928401925b828410156153a157835161549b81614640565b82529284019290840190615488565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006154f36040830184866154aa565b95945050505050565b60006020828403121561550e57600080fd5b81516001600160c01b0381168114613d7557600080fd5b60006020828403121561553757600080fd5b8151613d7581614640565b600060ff821660ff8103615558576155586152ed565b60010192915050565b6040815260006155756040830185876154aa565b905063ffffffff83166020830152949350505050565b63ffffffff831681526040602082015260006151c86040830184614750565b6000602082840312156155bc57600080fd5b8151613d7581614c3c565b81810381811115613d7857613d786152ed565b80820180821115613d7857613d786152ed565b6000602082840312156155ff57600080fd5b815167ffffffffffffffff1981168114613d7557600080fd5b6001600160601b03828116828216039080821115610c3e57610c3e6152ed565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561567357815185529382019390820190600101615657565b5092979650505050505050565b6000602080835263ffffffff808551168285015280828601511660408501525060408401516080606085015280518060a086015260005b818110156156d35782810184015186820160c0015283016156b7565b50600085820160c00152606086015163ffffffff811660808701529250601f01601f19169390930160c001949350505050565b6000808335601e1984360301811261571d57600080fd5b8301803591506001600160401b0382111561573757600080fd5b602001915036819003821315614a8157600080fd5b6000808335601e1984360301811261576357600080fd5b83016020810192503590506001600160401b0381111561578257600080fd5b803603821315614a8157600080fd5b60208152600082356157a281614640565b63ffffffff8082166020850152602085013591506157bf82614640565b80821660408501526157d4604086018661574c565b9250608060608601526157eb60a0860184836154aa565b92505060608501356157fc81614640565b166080939093019290925250919050565b63ffffffff818116838216019080821115610c3e57610c3e6152ed565b81835260006020808501808196508560051b810191508460005b8781101561587a57828403895261585b828861574c565b6158668682846154aa565b9a87019a9550505090840190600101615844565b5091979650505050505050565b6000813561589481614640565b63ffffffff168352602082013536839003601e190181126158b457600080fd5b82016020810190356001600160401b038111156158d057600080fd5b8060051b36038213156158e257600080fd5b604060208601526154f360408601828461582a565b60208152600061479e6020830184615887565b6001600160601b03818116838216028082169190828114613894576138946152ed565b8082028115828204841417613d7857613d786152ed565b6000808335601e1984360301811261595b57600080fd5b8301803591506001600160401b0382111561597557600080fd5b6020019150600581901b3603821315614a8157600080fd5b6020815260006151c860208301848661582a565b6060815260006159b46060830185615887565b905063ffffffff8351166020830152602083015160408301529392505050565b600061ffff8083168181036159eb576159eb6152ed565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220d1f56bff3785cfdd85a30fee7121cc17e198e65fea430bcf7ef2435ad719937c64736f6c63430008150033",
}

// ContractZrTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractZrTaskManagerMetaData.ABI instead.
var ContractZrTaskManagerABI = ContractZrTaskManagerMetaData.ABI

// ContractZrTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractZrTaskManagerMetaData.Bin instead.
var ContractZrTaskManagerBin = ContractZrTaskManagerMetaData.Bin

// DeployContractZrTaskManager deploys a new Ethereum contract, binding an instance of ContractZrTaskManager to it.
func DeployContractZrTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractZrTaskManager, error) {
	parsed, err := ContractZrTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractZrTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractZrTaskManager{ContractZrTaskManagerCaller: ContractZrTaskManagerCaller{contract: contract}, ContractZrTaskManagerTransactor: ContractZrTaskManagerTransactor{contract: contract}, ContractZrTaskManagerFilterer: ContractZrTaskManagerFilterer{contract: contract}}, nil
}

// ContractZrTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractZrTaskManager struct {
	ContractZrTaskManagerCaller     // Read-only binding to the contract
	ContractZrTaskManagerTransactor // Write-only binding to the contract
	ContractZrTaskManagerFilterer   // Log filterer for contract events
}

// ContractZrTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractZrTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractZrTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractZrTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractZrTaskManagerSession struct {
	Contract     *ContractZrTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractZrTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractZrTaskManagerCallerSession struct {
	Contract *ContractZrTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractZrTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractZrTaskManagerTransactorSession struct {
	Contract     *ContractZrTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractZrTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractZrTaskManagerRaw struct {
	Contract *ContractZrTaskManager // Generic contract binding to access the raw methods on
}

// ContractZrTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractZrTaskManagerCallerRaw struct {
	Contract *ContractZrTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractZrTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractZrTaskManagerTransactorRaw struct {
	Contract *ContractZrTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractZrTaskManager creates a new instance of ContractZrTaskManager, bound to a specific deployed contract.
func NewContractZrTaskManager(address common.Address, backend bind.ContractBackend) (*ContractZrTaskManager, error) {
	contract, err := bindContractZrTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManager{ContractZrTaskManagerCaller: ContractZrTaskManagerCaller{contract: contract}, ContractZrTaskManagerTransactor: ContractZrTaskManagerTransactor{contract: contract}, ContractZrTaskManagerFilterer: ContractZrTaskManagerFilterer{contract: contract}}, nil
}

// NewContractZrTaskManagerCaller creates a new read-only instance of ContractZrTaskManager, bound to a specific deployed contract.
func NewContractZrTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractZrTaskManagerCaller, error) {
	contract, err := bindContractZrTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerCaller{contract: contract}, nil
}

// NewContractZrTaskManagerTransactor creates a new write-only instance of ContractZrTaskManager, bound to a specific deployed contract.
func NewContractZrTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractZrTaskManagerTransactor, error) {
	contract, err := bindContractZrTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerTransactor{contract: contract}, nil
}

// NewContractZrTaskManagerFilterer creates a new log filterer instance of ContractZrTaskManager, bound to a specific deployed contract.
func NewContractZrTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractZrTaskManagerFilterer, error) {
	contract, err := bindContractZrTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerFilterer{contract: contract}, nil
}

// bindContractZrTaskManager binds a generic wrapper to an already deployed contract.
func bindContractZrTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractZrTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZrTaskManager *ContractZrTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZrTaskManager.Contract.ContractZrTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZrTaskManager *ContractZrTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.ContractZrTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZrTaskManager *ContractZrTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.ContractZrTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZrTaskManager *ContractZrTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZrTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZrTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZrTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZrTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZrTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractZrTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractZrTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractZrTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractZrTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Aggregator(&_ContractZrTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Aggregator(&_ContractZrTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractZrTaskManager.Contract.AllTaskHashes(&_ContractZrTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractZrTaskManager.Contract.AllTaskHashes(&_ContractZrTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractZrTaskManager.Contract.AllTaskResponses(&_ContractZrTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractZrTaskManager.Contract.AllTaskResponses(&_ContractZrTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZrTaskManager.Contract.BlsApkRegistry(&_ContractZrTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZrTaskManager.Contract.BlsApkRegistry(&_ContractZrTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractZrTaskManager *ContractZrTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZrTaskManager.Contract.CheckSignatures(&_ContractZrTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZrTaskManager.Contract.CheckSignatures(&_ContractZrTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Delegation(&_ContractZrTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Delegation(&_ContractZrTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Generator() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Generator(&_ContractZrTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Generator(&_ContractZrTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractZrTaskManager.Contract.GetBatchOperatorFromId(&_ContractZrTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractZrTaskManager.Contract.GetBatchOperatorFromId(&_ContractZrTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractZrTaskManager.Contract.GetBatchOperatorId(&_ContractZrTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractZrTaskManager.Contract.GetBatchOperatorId(&_ContractZrTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZrTaskManager *ContractZrTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZrTaskManager.Contract.GetCheckSignaturesIndices(&_ContractZrTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZrTaskManager.Contract.GetCheckSignaturesIndices(&_ContractZrTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZrTaskManager *ContractZrTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZrTaskManager.Contract.GetOperatorState(&_ContractZrTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZrTaskManager.Contract.GetOperatorState(&_ContractZrTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractZrTaskManager *ContractZrTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZrTaskManager.Contract.GetOperatorState0(&_ContractZrTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZrTaskManager.Contract.GetOperatorState0(&_ContractZrTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZrTaskManager *ContractZrTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZrTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZrTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZrTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZrTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZrTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractZrTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZrTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractZrTaskManager.CallOpts)
}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) LatestTaskId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "latestTaskId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) LatestTaskId() (uint32, error) {
	return _ContractZrTaskManager.Contract.LatestTaskId(&_ContractZrTaskManager.CallOpts)
}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) LatestTaskId() (uint32, error) {
	return _ContractZrTaskManager.Contract.LatestTaskId(&_ContractZrTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Owner() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Owner(&_ContractZrTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractZrTaskManager.Contract.Owner(&_ContractZrTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractZrTaskManager.Contract.Paused(&_ContractZrTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractZrTaskManager.Contract.Paused(&_ContractZrTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractZrTaskManager.Contract.Paused0(&_ContractZrTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractZrTaskManager.Contract.Paused0(&_ContractZrTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractZrTaskManager.Contract.PauserRegistry(&_ContractZrTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractZrTaskManager.Contract.PauserRegistry(&_ContractZrTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZrTaskManager.Contract.RegistryCoordinator(&_ContractZrTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZrTaskManager.Contract.RegistryCoordinator(&_ContractZrTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractZrTaskManager.Contract.StakeRegistry(&_ContractZrTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractZrTaskManager.Contract.StakeRegistry(&_ContractZrTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZrTaskManager.Contract.StaleStakesForbidden(&_ContractZrTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZrTaskManager.Contract.StaleStakesForbidden(&_ContractZrTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractZrTaskManager.Contract.TaskNumber(&_ContractZrTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractZrTaskManager.Contract.TaskNumber(&_ContractZrTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractZrTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractZrTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractZrTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractZrTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractZrTaskManager *ContractZrTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZrTaskManager.Contract.TrySignatureAndApkVerification(&_ContractZrTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZrTaskManager.Contract.TrySignatureAndApkVerification(&_ContractZrTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// ZrServiceManager is a free data retrieval call binding the contract method 0x87712db5.
//
// Solidity: function zrServiceManager() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCaller) ZrServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrTaskManager.contract.Call(opts, &out, "zrServiceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZrServiceManager is a free data retrieval call binding the contract method 0x87712db5.
//
// Solidity: function zrServiceManager() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerSession) ZrServiceManager() (common.Address, error) {
	return _ContractZrTaskManager.Contract.ZrServiceManager(&_ContractZrTaskManager.CallOpts)
}

// ZrServiceManager is a free data retrieval call binding the contract method 0x87712db5.
//
// Solidity: function zrServiceManager() view returns(address)
func (_ContractZrTaskManager *ContractZrTaskManagerCallerSession) ZrServiceManager() (common.Address, error) {
	return _ContractZrTaskManager.Contract.ZrServiceManager(&_ContractZrTaskManager.CallOpts)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "createNewTask", taskId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) CreateNewTask(taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.CreateNewTask(&_ContractZrTaskManager.TransactOpts, taskId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 taskId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) CreateNewTask(taskId uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.CreateNewTask(&_ContractZrTaskManager.TransactOpts, taskId, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _zrServiceManager) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _zrServiceManager common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator, _zrServiceManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _zrServiceManager) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _zrServiceManager common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.Initialize(&_ContractZrTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _zrServiceManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _zrServiceManager) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _zrServiceManager common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.Initialize(&_ContractZrTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _zrServiceManager)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.Pause(&_ContractZrTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.Pause(&_ContractZrTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.PauseAll(&_ContractZrTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.PauseAll(&_ContractZrTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.RenounceOwnership(&_ContractZrTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.RenounceOwnership(&_ContractZrTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IZRTaskManagerTask, taskResponse IZRTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) RespondToTask(task IZRTaskManagerTask, taskResponse IZRTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.RespondToTask(&_ContractZrTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcaf73aa0.
//
// Solidity: function respondToTask((uint32,uint32,bytes,uint32) task, (uint32,string[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) RespondToTask(task IZRTaskManagerTask, taskResponse IZRTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.RespondToTask(&_ContractZrTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.SetPauserRegistry(&_ContractZrTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.SetPauserRegistry(&_ContractZrTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.SetStaleStakesForbidden(&_ContractZrTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.SetStaleStakesForbidden(&_ContractZrTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.TransferOwnership(&_ContractZrTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.TransferOwnership(&_ContractZrTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.Unpause(&_ContractZrTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrTaskManager *ContractZrTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrTaskManager.Contract.Unpause(&_ContractZrTaskManager.TransactOpts, newPausedStatus)
}

// ContractZrTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerInitializedIterator struct {
	Event *ContractZrTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerInitialized)
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
		it.Event = new(ContractZrTaskManagerInitialized)
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
func (it *ContractZrTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerInitialized represents a Initialized event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractZrTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerInitializedIterator{contract: _ContractZrTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerInitialized)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractZrTaskManagerInitialized, error) {
	event := new(ContractZrTaskManagerInitialized)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerNewTaskCreatedIterator struct {
	Event *ContractZrTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerNewTaskCreated)
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
		it.Event = new(ContractZrTaskManagerNewTaskCreated)
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
func (it *ContractZrTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerNewTaskCreated struct {
	TaskId uint32
	Task   IZRTaskManagerTask
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xf31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskId []uint32) (*ContractZrTaskManagerNewTaskCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerNewTaskCreatedIterator{contract: _ContractZrTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xf31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca5.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,bytes,uint32) task)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerNewTaskCreated, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerNewTaskCreated)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractZrTaskManagerNewTaskCreated, error) {
	event := new(ContractZrTaskManagerNewTaskCreated)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerOwnershipTransferredIterator struct {
	Event *ContractZrTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractZrTaskManagerOwnershipTransferred)
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
func (it *ContractZrTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractZrTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerOwnershipTransferredIterator{contract: _ContractZrTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerOwnershipTransferred)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractZrTaskManagerOwnershipTransferred, error) {
	event := new(ContractZrTaskManagerOwnershipTransferred)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerPausedIterator struct {
	Event *ContractZrTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerPaused)
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
		it.Event = new(ContractZrTaskManagerPaused)
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
func (it *ContractZrTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerPaused represents a Paused event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractZrTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerPausedIterator{contract: _ContractZrTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerPaused)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParsePaused(log types.Log) (*ContractZrTaskManagerPaused, error) {
	event := new(ContractZrTaskManagerPaused)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerPauserRegistrySetIterator struct {
	Event *ContractZrTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractZrTaskManagerPauserRegistrySet)
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
func (it *ContractZrTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractZrTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerPauserRegistrySetIterator{contract: _ContractZrTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerPauserRegistrySet)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractZrTaskManagerPauserRegistrySet, error) {
	event := new(ContractZrTaskManagerPauserRegistrySet)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractZrTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractZrTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractZrTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractZrTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractZrTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractZrTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractZrTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractZrTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractZrTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractZrTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskChallengedSuccessfully struct {
	TaskId     uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskId []uint32, challenger []common.Address) (*ContractZrTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractZrTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerTaskChallengedSuccessfully, taskId []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerTaskChallengedSuccessfully)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractZrTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractZrTaskManagerTaskChallengedSuccessfully)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractZrTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractZrTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractZrTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskChallengedUnsuccessfully struct {
	TaskId     uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskId []uint32, challenger []common.Address) (*ContractZrTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractZrTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerTaskChallengedUnsuccessfully, taskId []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractZrTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractZrTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskCompletedIterator struct {
	Event *ContractZrTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerTaskCompleted)
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
		it.Event = new(ContractZrTaskManagerTaskCompleted)
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
func (it *ContractZrTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskCompleted struct {
	TaskId uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskId)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []uint32) (*ContractZrTaskManagerTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerTaskCompletedIterator{contract: _ContractZrTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskId)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerTaskCompleted, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerTaskCompleted)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractZrTaskManagerTaskCompleted, error) {
	event := new(ContractZrTaskManagerTaskCompleted)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskRespondedIterator struct {
	Event *ContractZrTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerTaskResponded)
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
		it.Event = new(ContractZrTaskManagerTaskResponded)
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
func (it *ContractZrTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerTaskResponded represents a TaskResponded event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerTaskResponded struct {
	TaskResponse         IZRTaskManagerTaskResponse
	TaskResponseMetadata IZRTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567.
//
// Solidity: event TaskResponded((uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractZrTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerTaskRespondedIterator{contract: _ContractZrTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567.
//
// Solidity: event TaskResponded((uint32,string[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerTaskResponded)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractZrTaskManagerTaskResponded, error) {
	event := new(ContractZrTaskManagerTaskResponded)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerUnpausedIterator struct {
	Event *ContractZrTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractZrTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrTaskManagerUnpaused)
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
		it.Event = new(ContractZrTaskManagerUnpaused)
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
func (it *ContractZrTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrTaskManagerUnpaused represents a Unpaused event raised by the ContractZrTaskManager contract.
type ContractZrTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractZrTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrTaskManagerUnpausedIterator{contract: _ContractZrTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractZrTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrTaskManagerUnpaused)
				if err := _ContractZrTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractZrTaskManager *ContractZrTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractZrTaskManagerUnpaused, error) {
	event := new(ContractZrTaskManagerUnpaused)
	if err := _ContractZrTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
