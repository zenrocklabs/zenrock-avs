// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractZRTaskManager

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

// ZRTaskManagerITask is an auto generated low-level Go binding around an user-defined struct.
type ZRTaskManagerITask struct {
	TaskId                    uint32
	TaskCreatedBlock          uint32
	ZrChainBlockHeight        int64
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// ZRTaskManagerITaskResponse is an auto generated low-level Go binding around an user-defined struct.
type ZRTaskManagerITaskResponse struct {
	ReferenceTaskId    uint32
	ZrChainBlockHeight int64
}

// ZRTaskManagerITaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type ZRTaskManagerITaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// ContractZRTaskManagerMetaData contains all meta data concerning the ContractZRTaskManager contract.
var ContractZRTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"zrChainBlockHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structZRTaskManagerI.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"zrChainBlockHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structZRTaskManagerI.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"zrChainBlockHeight\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structZRTaskManagerI.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structZRTaskManagerI.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"zrChainBlockHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structZRTaskManagerI.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"zrChainBlockHeight\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZRTaskManagerI.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"zrChainBlockHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZRTaskManagerI.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"zrChainBlockHeight\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZRTaskManagerI.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005ed338038062005ed38339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615bdc620002f7600039600081816102900152818161059c01526107e5015260008181610565015261249701526000818161042e0152818161267901526132c70152600081816104550152818161284f0152612a1101526000818161047c0152818161128701528181612162015281816122fa01526125340152615bdc6000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80636830483511610125578063b98d0908116100ad578063f2fde38b1161007c578063f2fde38b14610587578063f5c9899d1461059a578063f63c5bab146105c0578063f8c8765e146105c9578063fabc1cbc146105dc57600080fd5b8063b98d09081461051f578063ccbc7cc41461052c578063cefdc1d41461053f578063df5cf7231461056057600080fd5b806372d18e8d116100f457806372d18e8d146104c75780637afa1eed146104d55780637b9675d5146104e8578063886f1195146104fb5780638da5cb5b1461050e57600080fd5b806368304835146104505780636d14a987146104775780636efb46361461049e578063715018a6146104bf57600080fd5b80633563b0d1116101a85780635ac86ab7116101775780635ac86ab7146103ab5780635c155662146103de5780635c975abb146103fe5780635decc3f5146104065780635df459461461042957600080fd5b80633563b0d114610350578063416c7e5e146103705780634f739f7414610383578063595c6a67146103a357600080fd5b80631ad43189116101ef5780631ad431891461028b578063245a7bfc146102c75780632cb223d5146102f25780632d3de569146103205780632d89f6fc1461033057600080fd5b8063035b05601461022157806310d67a2f14610236578063136439dd14610249578063171f1d5b1461025c575b600080fd5b61023461022f366004614ac7565b6105ef565b005b610234610244366004614b50565b610a73565b610234610257366004614b6d565b610b26565b61026f61026a366004614b86565b610c65565b6040805192151583529015156020830152015b60405180910390f35b6102b27f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610282565b60cd546102da906001600160a01b031681565b6040516001600160a01b039091168152602001610282565b610312610300366004614bd7565b60cb6020526000908152604090205481565b604051908152602001610282565b60c9546102b29063ffffffff1681565b61031261033e366004614bd7565b60ca6020526000908152604090205481565b61036361035e366004614bf4565b610def565b6040516102829190614d4f565b61023461037e366004614d77565b611285565b610396610391366004614ddc565b6113fa565b6040516102829190614ee0565b610234611b20565b6103ce6103b9366004614faa565b606654600160ff9092169190911b9081161490565b6040519015158152602001610282565b6103f16103ec366004614fc7565b611be7565b6040516102829190615073565b606654610312565b6103ce610414366004614bd7565b60cc6020526000908152604090205460ff1681565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6104b16104ac3660046150b7565b611daf565b604051610282929190615177565b610234612cc6565b60c95463ffffffff166102b2565b60ce546102da906001600160a01b031681565b6102346104f63660046151d2565b612cda565b6065546102da906001600160a01b031681565b6033546001600160a01b03166102da565b6097546103ce9060ff1681565b61023461053a36600461524b565b612e43565b61055261054d3660046152d1565b6133fc565b604051610282929190615313565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b610234610595366004614b50565b61358e565b7f00000000000000000000000000000000000000000000000000000000000000006102b2565b6102b261271081565b6102346105d7366004615334565b613604565b6102346105ea366004614b6d565b613755565b60cd546001600160a01b0316331461064e5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064015b60405180910390fd5b60006106606040850160208601614bd7565b90503660006106726060870187615390565b9092509050600061068960a0880160808901614bd7565b905060ca600061069c60208a018a614bd7565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016106c891906153ff565b60405160208183030381529060405280519060200120146107515760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610645565b600060cb8161076360208b018b614bd7565b63ffffffff1663ffffffff16815260200190815260200160002054146107e05760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610645565b61080a7f0000000000000000000000000000000000000000000000000000000000000000856154da565b63ffffffff164363ffffffff16111561087b5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610645565b60008660405160200161088e919061552d565b6040516020818303038152906040528051906020012090506000806108b68387878a8c611daf565b9150915060005b858110156109b5578460ff16836020015182815181106108df576108df61553b565b60200260200101516108f19190615551565b6001600160601b03166064846000015183815181106109125761091261553b565b60200260200101516001600160601b031661092d9190615580565b10156109a3576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610645565b806109ad8161559f565b9150506108bd565b5060408051808201825263ffffffff431681526020808201849052915190916109e2918c918491016155ba565b6040516020818303038152906040528051906020012060cb60008d6000016020810190610a0f9190614bd7565b63ffffffff1663ffffffff168152602001908152602001600020819055507ff8a353d6881eae701031ace93bd11c2cf937b844f15a3646feef3f4840d681368a82604051610a5e9291906155ba565b60405180910390a15050505050505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ac6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aea91906155e6565b6001600160a01b0316336001600160a01b031614610b1a5760405162461bcd60e51b815260040161064590615603565b610b23816138b1565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610b6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b92919061564d565b610bae5760405162461bcd60e51b81526004016106459061566a565b60665481811614610c275760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610645565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610cad57610cad61553b565b60200201518951600160200201518a60200151600060028110610cd257610cd261553b565b60200201518b60200151600160028110610cee57610cee61553b565b602090810291909101518c518d830151604051610d4b9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d6e91906156b2565b9050610de1610d87610d8088846139a8565b8690613a3f565b610d8f613ad3565b610dd7610dc885610dc2604080518082018252600080825260209182015281518083019092526001825260029082015290565b906139a8565b610dd18c613b93565b90613a3f565b886201d4c0613c23565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5591906155e6565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ebb91906155e6565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610efd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2191906155e6565b9050600086516001600160401b03811115610f3e57610f3e6146a1565b604051908082528060200260200182016040528015610f7157816020015b6060815260200190600190039081610f5c5790505b50905060005b8751811015611279576000888281518110610f9457610f9461553b565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610ff5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261101d91908101906156d4565b905080516001600160401b03811115611038576110386146a1565b60405190808252806020026020018201604052801561108357816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816110565790505b508484815181106110965761109661553b565b602002602001018190525060005b8151811015611263576040518060600160405280876001600160a01b03166347b314e88585815181106110d9576110d961553b565b60200260200101516040518263ffffffff1660e01b81526004016110ff91815260200190565b602060405180830381865afa15801561111c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061114091906155e6565b6001600160a01b031681526020018383815181106111605761116061553b565b60200260200101518152602001896001600160a01b031663fa28c62785858151811061118e5761118e61553b565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa1580156111ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120e9190615764565b6001600160601b031681525085858151811061122c5761122c61553b565b602002602001015182815181106112455761124561553b565b6020026020010181905250808061125b9061559f565b9150506110a4565b50505080806112719061559f565b915050610f77565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130791906155e6565b6001600160a01b0316336001600160a01b0316146113b35760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610645565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6114256040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611465573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061148991906155e6565b90506114b66040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906114e6908b908990899060040161578d565b600060405180830381865afa158015611503573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261152b91908101906157d7565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061155d908b908b908b90600401615865565b600060405180830381865afa15801561157a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115a291908101906157d7565b6040820152856001600160401b038111156115bf576115bf6146a1565b6040519080825280602002602001820160405280156115f257816020015b60608152602001906001900390816115dd5790505b50606082015260005b60ff8116871115611a31576000856001600160401b03811115611620576116206146a1565b604051908082528060200260200182016040528015611649578160200160208202803683370190505b5083606001518360ff16815181106116635761166361553b565b602002602001018190525060005b868110156119315760008c6001600160a01b03166304ec63518a8a8581811061169c5761169c61553b565b905060200201358e886000015186815181106116ba576116ba61553b565b60200260200101516040518463ffffffff1660e01b81526004016116f79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611714573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611738919061588e565b90506001600160c01b0381166117dc5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610645565b8a8a8560ff168181106117f1576117f161553b565b6001600160c01b03841692013560f81c9190911c60019081161415905061191e57856001600160a01b031663dd9846b98a8a858181106118335761183361553b565b905060200201358d8d8860ff1681811061184f5761184f61553b565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156118a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118c991906158b7565b85606001518560ff16815181106118e2576118e261553b565b602002602001015184815181106118fb576118fb61553b565b63ffffffff909216602092830291909101909101528261191a8161559f565b9350505b50806119298161559f565b915050611671565b506000816001600160401b0381111561194c5761194c6146a1565b604051908082528060200260200182016040528015611975578160200160208202803683370190505b50905060005b828110156119f65784606001518460ff168151811061199c5761199c61553b565b602002602001015181815181106119b5576119b561553b565b60200260200101518282815181106119cf576119cf61553b565b63ffffffff90921660209283029190910190910152806119ee8161559f565b91505061197b565b508084606001518460ff1681518110611a1157611a1161553b565b602002602001018190525050508080611a29906158d4565b9150506115fb565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a9691906155e6565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611ac9908b908b908e906004016158f4565b600060405180830381865afa158015611ae6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b0e91908101906157d7565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611b68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8c919061564d565b611ba85760405162461bcd60e51b81526004016106459061566a565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611c1992919061591e565b600060405180830381865afa158015611c36573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c5e91908101906157d7565b9050600084516001600160401b03811115611c7b57611c7b6146a1565b604051908082528060200260200182016040528015611ca4578160200160208202803683370190505b50905060005b8551811015611da557866001600160a01b03166304ec6351878381518110611cd457611cd461553b565b602002602001015187868581518110611cef57611cef61553b565b60200260200101516040518463ffffffff1660e01b8152600401611d2c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611d49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d6d919061588e565b6001600160c01b0316828281518110611d8857611d8861553b565b602090810291909101015280611d9d8161559f565b915050611caa565b5095945050505050565b6040805180820190915260608082526020820152600084611e265760405162461bcd60e51b81526020600482015260376024820152600080516020615b8783398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610645565b60408301515185148015611e3e575060a08301515185145b8015611e4e575060c08301515185145b8015611e5e575060e08301515185145b611ec85760405162461bcd60e51b81526020600482015260416024820152600080516020615b8783398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610645565b82515160208401515114611f405760405162461bcd60e51b815260206004820152604460248201819052600080516020615b87833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610645565b4363ffffffff168463ffffffff1610611faf5760405162461bcd60e51b815260206004820152603c6024820152600080516020615b8783398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610645565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611ff057611ff06146a1565b604051908082528060200260200182016040528015612019578160200160208202803683370190505b506020820152866001600160401b03811115612037576120376146a1565b604051908082528060200260200182016040528015612060578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115612094576120946146a1565b6040519080825280602002602001820160405280156120bd578160200160208202803683370190505b5081526020860151516001600160401b038111156120dd576120dd6146a1565b604051908082528060200260200182016040528015612106578160200160208202803683370190505b50816020018190525060006121d88a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156121af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121d39190615972565b613e47565b905060005b87602001515181101561247357612222886020015182815181106122035761220361553b565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106122385761223861553b565b602090810291909101015280156122f857602083015161225960018361598f565b815181106122695761226961553b565b602002602001015160001c8360200151828151811061228a5761228a61553b565b602002602001015160001c116122f8576040805162461bcd60e51b8152602060048201526024810191909152600080516020615b8783398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610645565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061233d5761233d61553b565b60200260200101518b8b60000151858151811061235c5761235c61553b565b60200260200101516040518463ffffffff1660e01b81526004016123999392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156123b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123da919061588e565b6001600160c01b0316836000015182815181106123f9576123f961553b565b60200260200101818152505061245f610d8061243384866000015185815181106124255761242561553b565b602002602001015116613eda565b8a6020015184815181106124495761244961553b565b6020026020010151613f0590919063ffffffff16565b94508061246b8161559f565b9150506121dd565b505061247e83613fe9565b60975490935060ff16600081612495576000612517565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124f3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061251791906159a6565b905060005b8a811015612b95578215612677578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106125735761257361553b565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156125b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125d791906159a6565b6125e191906159bf565b116126775760405162461bcd60e51b81526020600482015260666024820152600080516020615b8783398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610645565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106126b8576126b861553b565b9050013560f81c60f81b60f81c8c8c60a0015185815181106126dc576126dc61553b565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612738573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061275c91906159d7565b6001600160401b03191661277f8a6040015183815181106122035761220361553b565b67ffffffffffffffff19161461281b5760405162461bcd60e51b81526020600482015260616024820152600080516020615b8783398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610645565b61284b896040015182815181106128345761283461553b565b602002602001015187613a3f90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061288e5761288e61553b565b9050013560f81c60f81b60f81c8c8c60c0015185815181106128b2576128b261553b565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561290e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129329190615764565b856020015182815181106129485761294861553b565b6001600160601b039092166020928302919091018201528501518051829081106129745761297461553b565b6020026020010151856000015182815181106129925761299261553b565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612b8057612a0a866000015182815181106129dc576129dc61553b565b60200260200101518f8f868181106129f6576129f661553b565b600192013560f81c9290921c811614919050565b15612b6e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612a5057612a5061553b565b9050013560f81c60f81b60f81c8e89602001518581518110612a7457612a7461553b565b60200260200101518f60e001518881518110612a9257612a9261553b565b60200260200101518781518110612aab57612aab61553b565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612b0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b339190615764565b8751805185908110612b4757612b4761553b565b60200260200101818151612b5b9190615a02565b6001600160601b03169052506001909101905b80612b788161559f565b9150506129b6565b50508080612b8d9061559f565b91505061251c565b505050600080612baf8c868a606001518b60800151610c65565b9150915081612c205760405162461bcd60e51b81526020600482015260436024820152600080516020615b8783398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610645565b80612c815760405162461bcd60e51b81526020600482015260396024820152600080516020615b8783398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610645565b50506000878260200151604051602001612c9c929190615a2a565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612cce614084565b612cd860006140de565b565b60ce546001600160a01b03163314612d3e5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610645565b6040805160a08101825260608082015263ffffffff8781168252438116602080840191909152600788900b8385015290861660808301528251601f850182900482028101820190935283835290919084908490819084018382808284376000920191909152505050506060820152604051612dbd908290602001615a72565b60408051601f19818403018152828252805160209182012063ffffffff8a16600081815260ca9093529290912055907f32c7ec7eda84a70313d6f8818bf0d1c5966ebdc7d653d15dbafdee4c91721b0590612e19908490615a72565b60405180910390a2505060c9805463ffffffff191663ffffffff9590951694909417909355505050565b6000612e526020850185614bd7565b63ffffffff8116600090815260cb6020526040902054909150612ec15760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610645565b8383604051602001612ed4929190615b0e565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb90935291205414612f735760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610645565b63ffffffff8116600090815260cc602052604090205460ff161561300b5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610645565b61271061301b6020850185614bd7565b61302591906154da565b63ffffffff164363ffffffff1611156130a65760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610645565b6130b36020850185614bd7565b63ffffffff166130c66020870187614bd7565b63ffffffff161461310c5760405162461bcd60e51b815260206004820152601060248201526f0a8c2e6d640928840dad2e6dac2e8c6d60831b6044820152606401610645565b600082516001600160401b03811115613127576131276146a1565b604051908082528060200260200182016040528015613150578160200160208202803683370190505b50905060005b83518110156131a3576131748482815181106122035761220361553b565b8282815181106131865761318661553b565b60209081029190910101528061319b8161559f565b915050613156565b5060006131b66040880160208901614bd7565b826040516020016131c8929190615a2a565b604051602081830303815290604052805190602001209050846020013581146132725760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610645565b600084516001600160401b0381111561328d5761328d6146a1565b6040519080825280602002602001820160405280156132b6578160200160208202803683370190505b50905060005b85518110156133a9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106133065761330661553b565b60200260200101516040518263ffffffff1660e01b815260040161332c91815260200190565b602060405180830381865afa158015613349573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061336d91906155e6565b82828151811061337f5761337f61553b565b6001600160a01b0390921660209283029190910190910152806133a18161559f565b9150506132bc565b5063ffffffff8416600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106134375761343761553b565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90613473908890869060040161591e565b600060405180830381865afa158015613490573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526134b891908101906157d7565b6000815181106134ca576134ca61553b565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613536573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061355a919061588e565b6001600160c01b03169050600061357082614130565b90508161357e8a838a610def565b9550955050505050935093915050565b613596614084565b6001600160a01b0381166135fb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610645565b610b23816140de565b600054610100900460ff16158080156136245750600054600160ff909116105b8061363e5750303b15801561363e575060005460ff166001145b6136a15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610645565b6000805460ff1916600117905580156136c4576000805461ff0019166101001790555b6136cf8560006141fc565b6136d8846140de565b60cd80546001600160a01b038086166001600160a01b03199283161790925560ce805492851692909116919091179055801561374e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137cc91906155e6565b6001600160a01b0316336001600160a01b0316146137fc5760405162461bcd60e51b815260040161064590615603565b60665419811960665419161461387a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610645565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610c5a565b6001600160a01b03811661393f5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610645565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526139c461459d565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156139f7576139f9565bfe5b5080613a375760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610645565b505092915050565b6040805180820190915260008082526020820152613a5b6145bb565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156139f7575080613a375760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610645565b613adb6145d9565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613bc3600080516020615b67833981519152866156b2565b90505b613bcf816142e6565b9093509150600080516020615b67833981519152828309831415613c09576040805180820190915290815260208101919091529392505050565b600080516020615b67833981519152600182089050613bc6565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613c556145fe565b60005b6002811015613e1a576000613c6e826006615580565b9050848260028110613c8257613c8261553b565b60200201515183613c948360006159bf565b600c8110613ca457613ca461553b565b6020020152848260028110613cbb57613cbb61553b565b60200201516020015183826001613cd291906159bf565b600c8110613ce257613ce261553b565b6020020152838260028110613cf957613cf961553b565b6020020151515183613d0c8360026159bf565b600c8110613d1c57613d1c61553b565b6020020152838260028110613d3357613d3361553b565b6020020151516001602002015183613d4c8360036159bf565b600c8110613d5c57613d5c61553b565b6020020152838260028110613d7357613d7361553b565b602002015160200151600060028110613d8e57613d8e61553b565b602002015183613d9f8360046159bf565b600c8110613daf57613daf61553b565b6020020152838260028110613dc657613dc661553b565b602002015160200151600160028110613de157613de161553b565b602002015183613df28360056159bf565b600c8110613e0257613e0261553b565b60200201525080613e128161559f565b915050613c58565b50613e2361461d565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613e5384614368565b9050808360ff166001901b11613ed15760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610645565b90505b92915050565b6000805b8215613ed457613eef60018461598f565b9092169180613efd81615b44565b915050613ede565b60408051808201909152600080825260208201526102008261ffff1610613f615760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610645565b8161ffff1660011415613f75575081613ed4565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613fde57600161ffff871660ff83161c81161415613fc157613fbe8484613a3f565b93505b613fcb8384613a3f565b92506201fffe600192831b169101613f91565b509195945050505050565b6040805180820190915260008082526020820152815115801561400e57506020820151155b1561402c575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615b67833981519152846020015161405f91906156b2565b61407790600080516020615b6783398151915261598f565b905292915050565b919050565b6033546001600160a01b03163314612cd85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610645565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606060008061413e84613eda565b61ffff166001600160401b03811115614159576141596146a1565b6040519080825280601f01601f191660200182016040528015614183576020820181803683370190505b5090506000805b82518210801561419b575061010081105b156141f2576001811b9350858416156141e2578060f81b8383815181106141c4576141c461553b565b60200101906001600160f81b031916908160001a9053508160010191505b6141eb8161559f565b905061418a565b5090949350505050565b6065546001600160a01b031615801561421d57506001600160a01b03821615155b61429f5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610645565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26142e2826138b1565b5050565b60008080600080516020615b678339815191526003600080516020615b6783398151915286600080516020615b6783398151915288890909089050600061435c827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615b678339815191526144f5565b91959194509092505050565b6000610100825111156143f15760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610645565b81516143ff57506000919050565b600080836000815181106144155761441561553b565b0160200151600160f89190911c81901b92505b84518110156144ec578481815181106144435761444361553b565b0160200151600160f89190911c1b91508282116144d85760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610645565b918117916144e58161559f565b9050614428565b50909392505050565b60008061450061461d565b61450861463b565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156139f75750826145925760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610645565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806145ec614659565b81526020016145f9614659565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b600060a0828403121561468957600080fd5b50919050565b60006040828403121561468957600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156146d9576146d96146a1565b60405290565b60405161010081016001600160401b03811182821017156146d9576146d96146a1565b604051601f8201601f191681016001600160401b038111828210171561472a5761472a6146a1565b604052919050565b60006001600160401b0382111561474b5761474b6146a1565b5060051b60200190565b63ffffffff81168114610b2357600080fd5b803561407f81614755565b600082601f83011261478357600080fd5b8135602061479861479383614732565b614702565b82815260059290921b840181019181810190868411156147b757600080fd5b8286015b848110156147db5780356147ce81614755565b83529183019183016147bb565b509695505050505050565b6000604082840312156147f857600080fd5b6148006146b7565b9050813581526020820135602082015292915050565b600082601f83011261482757600080fd5b8135602061483761479383614732565b82815260069290921b8401810191818101908684111561485657600080fd5b8286015b848110156147db5761486c88826147e6565b83529183019160400161485a565b600082601f83011261488b57600080fd5b604051604081018181106001600160401b03821117156148ad576148ad6146a1565b80604052508060408401858111156148c457600080fd5b845b81811015613fde5780358352602092830192016148c6565b6000608082840312156148f057600080fd5b6148f86146b7565b9050614904838361487a565b8152614913836040840161487a565b602082015292915050565b600082601f83011261492f57600080fd5b8135602061493f61479383614732565b82815260059290921b8401810191818101908684111561495e57600080fd5b8286015b848110156147db5780356001600160401b038111156149815760008081fd5b61498f8986838b0101614772565b845250918301918301614962565b600061018082840312156149b057600080fd5b6149b86146df565b905081356001600160401b03808211156149d157600080fd5b6149dd85838601614772565b835260208401359150808211156149f357600080fd5b6149ff85838601614816565b60208401526040840135915080821115614a1857600080fd5b614a2485838601614816565b6040840152614a3685606086016148de565b6060840152614a488560e086016147e6565b6080840152610120840135915080821115614a6257600080fd5b614a6e85838601614772565b60a0840152610140840135915080821115614a8857600080fd5b614a9485838601614772565b60c0840152610160840135915080821115614aae57600080fd5b50614abb8482850161491e565b60e08301525092915050565b600080600060808486031215614adc57600080fd5b83356001600160401b0380821115614af357600080fd5b614aff87838801614677565b9450614b0e876020880161468f565b93506060860135915080821115614b2457600080fd5b50614b318682870161499d565b9150509250925092565b6001600160a01b0381168114610b2357600080fd5b600060208284031215614b6257600080fd5b8135613ed181614b3b565b600060208284031215614b7f57600080fd5b5035919050565b6000806000806101208587031215614b9d57600080fd5b84359350614bae86602087016147e6565b9250614bbd86606087016148de565b9150614bcc8660e087016147e6565b905092959194509250565b600060208284031215614be957600080fd5b8135613ed181614755565b600080600060608486031215614c0957600080fd5b8335614c1481614b3b565b92506020848101356001600160401b0380821115614c3157600080fd5b818701915087601f830112614c4557600080fd5b813581811115614c5757614c576146a1565b614c69601f8201601f19168501614702565b91508082528884828501011115614c7f57600080fd5b8084840185840137600084828401015250809450505050614ca260408501614767565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614d41578385038a52825180518087529087019087870190845b81811015614d2c57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614ce8565b50509a87019a95505091850191600101614cca565b509298975050505050505050565b602081526000614d626020830184614cab565b9392505050565b8015158114610b2357600080fd5b600060208284031215614d8957600080fd5b8135613ed181614d69565b60008083601f840112614da657600080fd5b5081356001600160401b03811115614dbd57600080fd5b602083019150836020828501011115614dd557600080fd5b9250929050565b60008060008060008060808789031215614df557600080fd5b8635614e0081614b3b565b95506020870135614e1081614755565b945060408701356001600160401b0380821115614e2c57600080fd5b614e388a838b01614d94565b90965094506060890135915080821115614e5157600080fd5b818901915089601f830112614e6557600080fd5b813581811115614e7457600080fd5b8a60208260051b8501011115614e8957600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614ed557815163ffffffff1687529582019590820190600101614eb3565b509495945050505050565b600060208083528351608082850152614efc60a0850182614e9f565b905081850151601f1980868403016040870152614f198383614e9f565b92506040870151915080868403016060870152614f368383614e9f565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614f8d5784878303018452614f7b828751614e9f565b95880195938801939150600101614f61565b509998505050505050505050565b60ff81168114610b2357600080fd5b600060208284031215614fbc57600080fd5b8135613ed181614f9b565b600080600060608486031215614fdc57600080fd5b8335614fe781614b3b565b92506020848101356001600160401b0381111561500357600080fd5b8501601f8101871361501457600080fd5b803561502261479382614732565b81815260059190911b8201830190838101908983111561504157600080fd5b928401925b8284101561505f57833582529284019290840190615046565b8096505050505050614ca260408501614767565b6020808252825182820181905260009190848201906040850190845b818110156150ab5783518352928401929184019160010161508f565b50909695505050505050565b6000806000806000608086880312156150cf57600080fd5b8535945060208601356001600160401b03808211156150ed57600080fd5b6150f989838a01614d94565b90965094506040880135915061510e82614755565b9092506060870135908082111561512457600080fd5b506151318882890161499d565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614ed55781516001600160601b031687529582019590820190600101615152565b6040815260008351604080840152615192608084018261513e565b90506020850151603f198483030160608501526151af828261513e565b925050508260208301529392505050565b8035600781900b811461407f57600080fd5b6000806000806000608086880312156151ea57600080fd5b85356151f581614755565b9450615203602087016151c0565b9350604086013561521381614755565b925060608601356001600160401b0381111561522e57600080fd5b61523a88828901614d94565b969995985093965092949392505050565b60008060008060c0858703121561526157600080fd5b84356001600160401b038082111561527857600080fd5b61528488838901614677565b9550615293886020890161468f565b94506152a2886060890161468f565b935060a08701359150808211156152b857600080fd5b506152c587828801614816565b91505092959194509250565b6000806000606084860312156152e657600080fd5b83356152f181614b3b565b925060208401359150604084013561530881614755565b809150509250925092565b82815260406020820152600061532c6040830184614cab565b949350505050565b6000806000806080858703121561534a57600080fd5b843561535581614b3b565b9350602085013561536581614b3b565b9250604085013561537581614b3b565b9150606085013561538581614b3b565b939692955090935050565b6000808335601e198436030181126153a757600080fd5b8301803591506001600160401b038211156153c157600080fd5b602001915036819003821315614dd557600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526000823561541081614755565b63ffffffff80821660208501526020850135915061542d82614755565b80821660408501525050615443604084016151c0565b60070b60608301526060830135601e1984360301811261546257600080fd5b830180356001600160401b0381111561547a57600080fd5b80360385131561548957600080fd5b60a060808501526154a160c0850182602085016153d6565b9150506154b060808501614767565b63ffffffff811660a0850152509392505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168185168083038211156154f9576154f96154c4565b01949350505050565b803561550d81614755565b63ffffffff168252615521602082016151c0565b60070b60208301525050565b60408101613ed48284615502565b634e487b7160e01b600052603260045260246000fd5b60006001600160601b0380831681851681830481118215151615615577576155776154c4565b02949350505050565b600081600019048311821515161561559a5761559a6154c4565b500290565b60006000198214156155b3576155b36154c4565b5060010190565b608081016155c88285615502565b63ffffffff8351166040830152602083015160608301529392505050565b6000602082840312156155f857600080fd5b8151613ed181614b3b565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561565f57600080fd5b8151613ed181614d69565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000826156cf57634e487b7160e01b600052601260045260246000fd5b500690565b600060208083850312156156e757600080fd5b82516001600160401b038111156156fd57600080fd5b8301601f8101851361570e57600080fd5b805161571c61479382614732565b81815260059190911b8201830190838101908783111561573b57600080fd5b928401925b8284101561575957835182529284019290840190615740565b979650505050505050565b60006020828403121561577657600080fd5b81516001600160601b0381168114613ed157600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156157ba57600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156157ea57600080fd5b82516001600160401b0381111561580057600080fd5b8301601f8101851361581157600080fd5b805161581f61479382614732565b81815260059190911b8201830190838101908783111561583e57600080fd5b928401925b8284101561575957835161585681614755565b82529284019290840190615843565b63ffffffff841681526040602082015260006158856040830184866153d6565b95945050505050565b6000602082840312156158a057600080fd5b81516001600160c01b0381168114613ed157600080fd5b6000602082840312156158c957600080fd5b8151613ed181614755565b600060ff821660ff8114156158eb576158eb6154c4565b60010192915050565b6040815260006159086040830185876153d6565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561596557845183529383019391830191600101615949565b5090979650505050505050565b60006020828403121561598457600080fd5b8151613ed181614f9b565b6000828210156159a1576159a16154c4565b500390565b6000602082840312156159b857600080fd5b5051919050565b600082198211156159d2576159d26154c4565b500190565b6000602082840312156159e957600080fd5b815167ffffffffffffffff1981168114613ed157600080fd5b60006001600160601b0383811690831681811015615a2257615a226154c4565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615a6557815185529382019390820190600101615a49565b5092979650505050505050565b6000602080835263ffffffff8085511682850152808286015116604085015250604084015160070b6060840152606084015160a0608085015280518060c086015260005b81811015615ad25782810184015186820160e001528301615ab6565b81811115615ae457600060e083880101525b50608086015163ffffffff811660a08701529250601f01601f19169390930160e001949350505050565b60808101615b1c8285615502565b8235615b2781614755565b63ffffffff16604083015260209290920135606090910152919050565b600061ffff80831681811415615b5c57615b5c6154c4565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122060c53981c424ea9357130112014dffebfdfe3b49a1a4863743dc1a058f86622864736f6c634300080c0033",
}

// ContractZRTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractZRTaskManagerMetaData.ABI instead.
var ContractZRTaskManagerABI = ContractZRTaskManagerMetaData.ABI

// ContractZRTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractZRTaskManagerMetaData.Bin instead.
var ContractZRTaskManagerBin = ContractZRTaskManagerMetaData.Bin

// DeployContractZRTaskManager deploys a new Ethereum contract, binding an instance of ContractZRTaskManager to it.
func DeployContractZRTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractZRTaskManager, error) {
	parsed, err := ContractZRTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractZRTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractZRTaskManager{ContractZRTaskManagerCaller: ContractZRTaskManagerCaller{contract: contract}, ContractZRTaskManagerTransactor: ContractZRTaskManagerTransactor{contract: contract}, ContractZRTaskManagerFilterer: ContractZRTaskManagerFilterer{contract: contract}}, nil
}

// ContractZRTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractZRTaskManager struct {
	ContractZRTaskManagerCaller     // Read-only binding to the contract
	ContractZRTaskManagerTransactor // Write-only binding to the contract
	ContractZRTaskManagerFilterer   // Log filterer for contract events
}

// ContractZRTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractZRTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZRTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractZRTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZRTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractZRTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZRTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractZRTaskManagerSession struct {
	Contract     *ContractZRTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractZRTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractZRTaskManagerCallerSession struct {
	Contract *ContractZRTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractZRTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractZRTaskManagerTransactorSession struct {
	Contract     *ContractZRTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractZRTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractZRTaskManagerRaw struct {
	Contract *ContractZRTaskManager // Generic contract binding to access the raw methods on
}

// ContractZRTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractZRTaskManagerCallerRaw struct {
	Contract *ContractZRTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractZRTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractZRTaskManagerTransactorRaw struct {
	Contract *ContractZRTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractZRTaskManager creates a new instance of ContractZRTaskManager, bound to a specific deployed contract.
func NewContractZRTaskManager(address common.Address, backend bind.ContractBackend) (*ContractZRTaskManager, error) {
	contract, err := bindContractZRTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManager{ContractZRTaskManagerCaller: ContractZRTaskManagerCaller{contract: contract}, ContractZRTaskManagerTransactor: ContractZRTaskManagerTransactor{contract: contract}, ContractZRTaskManagerFilterer: ContractZRTaskManagerFilterer{contract: contract}}, nil
}

// NewContractZRTaskManagerCaller creates a new read-only instance of ContractZRTaskManager, bound to a specific deployed contract.
func NewContractZRTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractZRTaskManagerCaller, error) {
	contract, err := bindContractZRTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerCaller{contract: contract}, nil
}

// NewContractZRTaskManagerTransactor creates a new write-only instance of ContractZRTaskManager, bound to a specific deployed contract.
func NewContractZRTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractZRTaskManagerTransactor, error) {
	contract, err := bindContractZRTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerTransactor{contract: contract}, nil
}

// NewContractZRTaskManagerFilterer creates a new log filterer instance of ContractZRTaskManager, bound to a specific deployed contract.
func NewContractZRTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractZRTaskManagerFilterer, error) {
	contract, err := bindContractZRTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerFilterer{contract: contract}, nil
}

// bindContractZRTaskManager binds a generic wrapper to an already deployed contract.
func bindContractZRTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractZRTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZRTaskManager *ContractZRTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZRTaskManager.Contract.ContractZRTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZRTaskManager *ContractZRTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.ContractZRTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZRTaskManager *ContractZRTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.ContractZRTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZRTaskManager *ContractZRTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZRTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZRTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZRTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZRTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZRTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractZRTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractZRTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractZRTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractZRTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Aggregator(&_ContractZRTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Aggregator(&_ContractZRTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractZRTaskManager.Contract.AllTaskHashes(&_ContractZRTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractZRTaskManager.Contract.AllTaskHashes(&_ContractZRTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractZRTaskManager.Contract.AllTaskResponses(&_ContractZRTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractZRTaskManager.Contract.AllTaskResponses(&_ContractZRTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZRTaskManager.Contract.BlsApkRegistry(&_ContractZRTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZRTaskManager.Contract.BlsApkRegistry(&_ContractZRTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractZRTaskManager *ContractZRTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZRTaskManager.Contract.CheckSignatures(&_ContractZRTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZRTaskManager.Contract.CheckSignatures(&_ContractZRTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Delegation(&_ContractZRTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Delegation(&_ContractZRTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Generator() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Generator(&_ContractZRTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Generator(&_ContractZRTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZRTaskManager *ContractZRTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZRTaskManager.Contract.GetCheckSignaturesIndices(&_ContractZRTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZRTaskManager.Contract.GetCheckSignaturesIndices(&_ContractZRTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZRTaskManager *ContractZRTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZRTaskManager.Contract.GetOperatorState(&_ContractZRTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZRTaskManager.Contract.GetOperatorState(&_ContractZRTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractZRTaskManager *ContractZRTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZRTaskManager.Contract.GetOperatorState0(&_ContractZRTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZRTaskManager.Contract.GetOperatorState0(&_ContractZRTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZRTaskManager *ContractZRTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZRTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZRTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZRTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZRTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZRTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractZRTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZRTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractZRTaskManager.CallOpts)
}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) LatestTaskId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "latestTaskId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) LatestTaskId() (uint32, error) {
	return _ContractZRTaskManager.Contract.LatestTaskId(&_ContractZRTaskManager.CallOpts)
}

// LatestTaskId is a free data retrieval call binding the contract method 0x2d3de569.
//
// Solidity: function latestTaskId() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) LatestTaskId() (uint32, error) {
	return _ContractZRTaskManager.Contract.LatestTaskId(&_ContractZRTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Owner() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Owner(&_ContractZRTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractZRTaskManager.Contract.Owner(&_ContractZRTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractZRTaskManager.Contract.Paused(&_ContractZRTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractZRTaskManager.Contract.Paused(&_ContractZRTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractZRTaskManager.Contract.Paused0(&_ContractZRTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractZRTaskManager.Contract.Paused0(&_ContractZRTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractZRTaskManager.Contract.PauserRegistry(&_ContractZRTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractZRTaskManager.Contract.PauserRegistry(&_ContractZRTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZRTaskManager.Contract.RegistryCoordinator(&_ContractZRTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZRTaskManager.Contract.RegistryCoordinator(&_ContractZRTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractZRTaskManager.Contract.StakeRegistry(&_ContractZRTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractZRTaskManager.Contract.StakeRegistry(&_ContractZRTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZRTaskManager.Contract.StaleStakesForbidden(&_ContractZRTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZRTaskManager.Contract.StaleStakesForbidden(&_ContractZRTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractZRTaskManager.Contract.TaskNumber(&_ContractZRTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractZRTaskManager.Contract.TaskNumber(&_ContractZRTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractZRTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractZRTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractZRTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractZRTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZRTaskManager *ContractZRTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractZRTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractZRTaskManager *ContractZRTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZRTaskManager.Contract.TrySignatureAndApkVerification(&_ContractZRTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZRTaskManager *ContractZRTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZRTaskManager.Contract.TrySignatureAndApkVerification(&_ContractZRTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x7b9675d5.
//
// Solidity: function createNewTask(uint32 taskId, int64 zrChainBlockHeight, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, taskId uint32, zrChainBlockHeight int64, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "createNewTask", taskId, zrChainBlockHeight, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x7b9675d5.
//
// Solidity: function createNewTask(uint32 taskId, int64 zrChainBlockHeight, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) CreateNewTask(taskId uint32, zrChainBlockHeight int64, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.CreateNewTask(&_ContractZRTaskManager.TransactOpts, taskId, zrChainBlockHeight, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x7b9675d5.
//
// Solidity: function createNewTask(uint32 taskId, int64 zrChainBlockHeight, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) CreateNewTask(taskId uint32, zrChainBlockHeight int64, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.CreateNewTask(&_ContractZRTaskManager.TransactOpts, taskId, zrChainBlockHeight, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.Initialize(&_ContractZRTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.Initialize(&_ContractZRTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.Pause(&_ContractZRTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.Pause(&_ContractZRTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.PauseAll(&_ContractZRTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.PauseAll(&_ContractZRTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xccbc7cc4.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,int64,bytes,uint32) task, (uint32,int64) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task ZRTaskManagerITask, taskResponse ZRTaskManagerITaskResponse, taskResponseMetadata ZRTaskManagerITaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xccbc7cc4.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,int64,bytes,uint32) task, (uint32,int64) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) RaiseAndResolveChallenge(task ZRTaskManagerITask, taskResponse ZRTaskManagerITaskResponse, taskResponseMetadata ZRTaskManagerITaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.RaiseAndResolveChallenge(&_ContractZRTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xccbc7cc4.
//
// Solidity: function raiseAndResolveChallenge((uint32,uint32,int64,bytes,uint32) task, (uint32,int64) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) RaiseAndResolveChallenge(task ZRTaskManagerITask, taskResponse ZRTaskManagerITaskResponse, taskResponseMetadata ZRTaskManagerITaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.RaiseAndResolveChallenge(&_ContractZRTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.RenounceOwnership(&_ContractZRTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.RenounceOwnership(&_ContractZRTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x035b0560.
//
// Solidity: function respondToTask((uint32,uint32,int64,bytes,uint32) task, (uint32,int64) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task ZRTaskManagerITask, taskResponse ZRTaskManagerITaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x035b0560.
//
// Solidity: function respondToTask((uint32,uint32,int64,bytes,uint32) task, (uint32,int64) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) RespondToTask(task ZRTaskManagerITask, taskResponse ZRTaskManagerITaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.RespondToTask(&_ContractZRTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x035b0560.
//
// Solidity: function respondToTask((uint32,uint32,int64,bytes,uint32) task, (uint32,int64) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) RespondToTask(task ZRTaskManagerITask, taskResponse ZRTaskManagerITaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.RespondToTask(&_ContractZRTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.SetPauserRegistry(&_ContractZRTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.SetPauserRegistry(&_ContractZRTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.SetStaleStakesForbidden(&_ContractZRTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.SetStaleStakesForbidden(&_ContractZRTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.TransferOwnership(&_ContractZRTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.TransferOwnership(&_ContractZRTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZRTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.Unpause(&_ContractZRTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZRTaskManager *ContractZRTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZRTaskManager.Contract.Unpause(&_ContractZRTaskManager.TransactOpts, newPausedStatus)
}

// ContractZRTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerInitializedIterator struct {
	Event *ContractZRTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerInitialized)
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
		it.Event = new(ContractZRTaskManagerInitialized)
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
func (it *ContractZRTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerInitialized represents a Initialized event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractZRTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerInitializedIterator{contract: _ContractZRTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerInitialized)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractZRTaskManagerInitialized, error) {
	event := new(ContractZRTaskManagerInitialized)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerNewTaskCreatedIterator struct {
	Event *ContractZRTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerNewTaskCreated)
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
		it.Event = new(ContractZRTaskManagerNewTaskCreated)
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
func (it *ContractZRTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerNewTaskCreated struct {
	TaskId uint32
	Task   ZRTaskManagerITask
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x32c7ec7eda84a70313d6f8818bf0d1c5966ebdc7d653d15dbafdee4c91721b05.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,int64,bytes,uint32) task)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskId []uint32) (*ContractZRTaskManagerNewTaskCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerNewTaskCreatedIterator{contract: _ContractZRTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x32c7ec7eda84a70313d6f8818bf0d1c5966ebdc7d653d15dbafdee4c91721b05.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,int64,bytes,uint32) task)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerNewTaskCreated, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerNewTaskCreated)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x32c7ec7eda84a70313d6f8818bf0d1c5966ebdc7d653d15dbafdee4c91721b05.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (uint32,uint32,int64,bytes,uint32) task)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractZRTaskManagerNewTaskCreated, error) {
	event := new(ContractZRTaskManagerNewTaskCreated)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerOwnershipTransferredIterator struct {
	Event *ContractZRTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractZRTaskManagerOwnershipTransferred)
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
func (it *ContractZRTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractZRTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerOwnershipTransferredIterator{contract: _ContractZRTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerOwnershipTransferred)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractZRTaskManagerOwnershipTransferred, error) {
	event := new(ContractZRTaskManagerOwnershipTransferred)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerPausedIterator struct {
	Event *ContractZRTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerPaused)
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
		it.Event = new(ContractZRTaskManagerPaused)
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
func (it *ContractZRTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerPaused represents a Paused event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractZRTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerPausedIterator{contract: _ContractZRTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerPaused)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParsePaused(log types.Log) (*ContractZRTaskManagerPaused, error) {
	event := new(ContractZRTaskManagerPaused)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerPauserRegistrySetIterator struct {
	Event *ContractZRTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractZRTaskManagerPauserRegistrySet)
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
func (it *ContractZRTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractZRTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerPauserRegistrySetIterator{contract: _ContractZRTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerPauserRegistrySet)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractZRTaskManagerPauserRegistrySet, error) {
	event := new(ContractZRTaskManagerPauserRegistrySet)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractZRTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractZRTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractZRTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractZRTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractZRTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractZRTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractZRTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractZRTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractZRTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractZRTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskChallengedSuccessfully struct {
	TaskId     uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskId []uint32, challenger []common.Address) (*ContractZRTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractZRTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerTaskChallengedSuccessfully, taskId []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerTaskChallengedSuccessfully)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractZRTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractZRTaskManagerTaskChallengedSuccessfully)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractZRTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractZRTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractZRTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskChallengedUnsuccessfully struct {
	TaskId     uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskId []uint32, challenger []common.Address) (*ContractZRTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractZRTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskId, address indexed challenger)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerTaskChallengedUnsuccessfully, taskId []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIdRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractZRTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractZRTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskCompletedIterator struct {
	Event *ContractZRTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerTaskCompleted)
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
		it.Event = new(ContractZRTaskManagerTaskCompleted)
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
func (it *ContractZRTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskCompleted struct {
	TaskId uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskId)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []uint32) (*ContractZRTaskManagerTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerTaskCompletedIterator{contract: _ContractZRTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskId)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerTaskCompleted, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerTaskCompleted)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractZRTaskManagerTaskCompleted, error) {
	event := new(ContractZRTaskManagerTaskCompleted)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskRespondedIterator struct {
	Event *ContractZRTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerTaskResponded)
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
		it.Event = new(ContractZRTaskManagerTaskResponded)
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
func (it *ContractZRTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerTaskResponded represents a TaskResponded event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerTaskResponded struct {
	TaskResponse         ZRTaskManagerITaskResponse
	TaskResponseMetadata ZRTaskManagerITaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xf8a353d6881eae701031ace93bd11c2cf937b844f15a3646feef3f4840d68136.
//
// Solidity: event TaskResponded((uint32,int64) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractZRTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerTaskRespondedIterator{contract: _ContractZRTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xf8a353d6881eae701031ace93bd11c2cf937b844f15a3646feef3f4840d68136.
//
// Solidity: event TaskResponded((uint32,int64) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerTaskResponded)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0xf8a353d6881eae701031ace93bd11c2cf937b844f15a3646feef3f4840d68136.
//
// Solidity: event TaskResponded((uint32,int64) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractZRTaskManagerTaskResponded, error) {
	event := new(ContractZRTaskManagerTaskResponded)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZRTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerUnpausedIterator struct {
	Event *ContractZRTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractZRTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZRTaskManagerUnpaused)
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
		it.Event = new(ContractZRTaskManagerUnpaused)
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
func (it *ContractZRTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZRTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZRTaskManagerUnpaused represents a Unpaused event raised by the ContractZRTaskManager contract.
type ContractZRTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractZRTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZRTaskManagerUnpausedIterator{contract: _ContractZRTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractZRTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZRTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZRTaskManagerUnpaused)
				if err := _ContractZRTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractZRTaskManager *ContractZRTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractZRTaskManagerUnpaused, error) {
	event := new(ContractZRTaskManagerUnpaused)
	if err := _ContractZRTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
