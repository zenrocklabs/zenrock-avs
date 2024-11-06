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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveSet\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestActiveSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidatorInTaskSet\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"validatorAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activeSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskManagerZR.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAddressesStored\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006112380380620061128339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615e22620002f06000396000818161027d015281816105770152612e7801526000818161054001526122f20152600081816103ee01526124d4015260008181610415015281816126aa015261286c01526000818161043c01528181610df801528181611fdc01528181612155015261238f0152615e226000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80636efb463611610125578063b98d0908116100ad578063f2fde38b1161007c578063f2fde38b14610562578063f5c9899d14610575578063f63c5bab1461059b578063f8c8765e146105a4578063fabc1cbc146105b757600080fd5b8063b98d0908146104fa578063caf73aa014610507578063cefdc1d41461051a578063df5cf7231461053b57600080fd5b806379e2c40f116100f457806379e2c40f1461049d5780637afa1eed146104b0578063886f1195146104c35780638da5cb5b146104d65780639d3a0f2d146104e757600080fd5b80636efb46361461045e578063715018a61461047f57806372d18e8d14610487578063767917c61461049557600080fd5b80634c71c165116101a85780635c155662116101775780635c155662146103b85780635c975abb146103d85780635df45946146103e957806368304835146104105780636d14a9871461043757600080fd5b80634c71c1651461034a5780634f739f741461035d578063595c6a671461037d5780635ac86ab71461038557600080fd5b8063245a7bfc116101ef578063245a7bfc146102b45780632d3de569146102e75780633563b0d1146102f75780633ec54dcf14610317578063416c7e5e1461033757600080fd5b806310d67a2f14610221578063136439dd14610236578063171f1d5b146102495780631ad4318914610278575b600080fd5b61023461022f366004614623565b6105ca565b005b610234610244366004614640565b610686565b61025c6102573660046147be565b6107c5565b6040805192151583529015156020830152015b60405180910390f35b61029f7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026f565b60c9546102cf9064010000000090046001600160a01b031681565b6040516001600160a01b03909116815260200161026f565b60c95461029f9063ffffffff1681565b61030a610305366004614883565b61094f565b60405161026f919061499e565b61032a6103253660046149b1565b610de5565b60405161026f9190614a2a565b610234610345366004614a9a565b610df6565b610234610358366004614b78565b610f6b565b61037061036b366004614c59565b611274565b60405161026f9190614d5d565b61023461199a565b6103a8610393366004614e27565b606654600160ff9092169190911b9081161490565b604051901515815260200161026f565b6103cb6103c6366004614e44565b611a61565b60405161026f9190614ef9565b60665460405190815260200161026f565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b61047161046c36600461514a565b611c29565b60405161026f92919061520a565b610234612b21565b60c95463ffffffff1661029f565b61032a612b35565b6103a86104ab366004615253565b612b4f565b60ca546102cf906001600160a01b031681565b6065546102cf906001600160a01b031681565b6033546001600160a01b03166102cf565b6102346104f53660046152b6565b612bfb565b6097546103a89060ff1681565b61023461051536600461531a565b612d47565b61052d6105283660046153a1565b6130e0565b60405161026f9291906153d8565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b610234610570366004614623565b613272565b7f000000000000000000000000000000000000000000000000000000000000000061029f565b61029f61271081565b6102346105b23660046153f9565b6132e8565b6102346105c5366004614640565b61344e565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561061d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106419190615455565b6001600160a01b0316336001600160a01b03161461067a5760405162461bcd60e51b815260040161067190615472565b60405180910390fd5b610683816135aa565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f291906154bc565b61070e5760405162461bcd60e51b8152600401610671906154d9565b606654818116146107875760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610671565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061080d5761080d615521565b60200201518951600160200201518a6020015160006002811061083257610832615521565b60200201518b6020015160016002811061084e5761084e615521565b602090810291909101518c518d8301516040516108ab9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108ce9190615537565b90506109416108e76108e088846136a1565b8690613738565b6108ef6137cc565b61093761092885610922604080518082018252600080825260209182015281518083019092526001825260029082015290565b906136a1565b6109318c61388c565b90613738565b886201d4c061391c565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610991573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b59190615455565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109f7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1b9190615455565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a819190615455565b9050600086516001600160401b03811115610a9e57610a9e614659565b604051908082528060200260200182016040528015610ad157816020015b6060815260200190600190039081610abc5790505b50905060005b8751811015610dd9576000888281518110610af457610af4615521565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b55573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b7d9190810190615559565b905080516001600160401b03811115610b9857610b98614659565b604051908082528060200260200182016040528015610be357816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bb65790505b50848481518110610bf657610bf6615521565b602002602001018190525060005b8151811015610dc3576040518060600160405280876001600160a01b03166347b314e8858581518110610c3957610c39615521565b60200260200101516040518263ffffffff1660e01b8152600401610c5f91815260200190565b602060405180830381865afa158015610c7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca09190615455565b6001600160a01b03168152602001838381518110610cc057610cc0615521565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610cee57610cee615521565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6e91906155e9565b6001600160601b0316815250858581518110610d8c57610d8c615521565b60200260200101518281518110610da557610da5615521565b60200260200101819052508080610dbb90615628565b915050610c04565b5050508080610dd190615628565b915050610ad7565b50979650505050505050565b6060610df082613b40565b92915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e789190615455565b6001600160a01b0316336001600160a01b031614610f245760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610671565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6000610f7a60208501856149b1565b63ffffffff8116600090815260cb60205260409020600181015491925090610fd45760405162461bcd60e51b815260206004820152600d60248201526c139bdd081c995cdc1bdb991959609a1b6044820152606401610671565b8484604051602001610fe7929190615772565b604051602081830303815290604052805190602001208160010154146110425760405162461bcd60e51b815260206004820152601060248201526f496e76616c696420726573706f6e736560801b6044820152606401610671565b600281015460ff161561108c5760405162461bcd60e51b8152602060048201526012602482015271105b1c9958591e4818da185b1b195b99d95960721b6044820152606401610671565b61271061109c60208601866149b1565b6110a691906157b0565b63ffffffff164311156110ef5760405162461bcd60e51b815260206004820152601160248201527010da185b1b195b99d948195e1c1a5c9959607a1b6044820152606401610671565b600083516001600160401b0381111561110a5761110a614659565b604051908082528060200260200182016040528015611133578160200160208202803683370190505b50905060005b84518110156111a55761117685828151811061115757611157615521565b6020026020010151805160009081526020918201519091526040902090565b82828151811061118857611188615521565b60209081029190910101528061119d81615628565b915050611139565b5084602001358760200160208101906111be91906149b1565b826040516020016111d09291906157d8565b60405160208183030381529060405280519060200120146112295760405162461bcd60e51b8152602060048201526013602482015272496e76616c6964206e6f6e2d7369676e65727360681b6044820152606401610671565b60028201805460ff19166001179055604051339063ffffffff8516907fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a350505050505050565b61129f6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113039190615455565b90506113306040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611360908b9089908990600401615813565b600060405180830381865afa15801561137d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526113a5919081019061585d565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906113d7908b908b908b906004016158eb565b600060405180830381865afa1580156113f4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261141c919081019061585d565b6040820152856001600160401b0381111561143957611439614659565b60405190808252806020026020018201604052801561146c57816020015b60608152602001906001900390816114575790505b50606082015260005b60ff81168711156118ab576000856001600160401b0381111561149a5761149a614659565b6040519080825280602002602001820160405280156114c3578160200160208202803683370190505b5083606001518360ff16815181106114dd576114dd615521565b602002602001018190525060005b868110156117ab5760008c6001600160a01b03166304ec63518a8a8581811061151657611516615521565b905060200201358e8860000151868151811061153457611534615521565b60200260200101516040518463ffffffff1660e01b81526004016115719392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561158e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b29190615914565b90506001600160c01b0381166116565760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610671565b8a8a8560ff1681811061166b5761166b615521565b6001600160c01b03841692013560f81c9190911c60019081161415905061179857856001600160a01b031663dd9846b98a8a858181106116ad576116ad615521565b905060200201358d8d8860ff168181106116c9576116c9615521565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561171f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611743919061593d565b85606001518560ff168151811061175c5761175c615521565b6020026020010151848151811061177557611775615521565b63ffffffff909216602092830291909101909101528261179481615628565b9350505b50806117a381615628565b9150506114eb565b506000816001600160401b038111156117c6576117c6614659565b6040519080825280602002602001820160405280156117ef578160200160208202803683370190505b50905060005b828110156118705784606001518460ff168151811061181657611816615521565b6020026020010151818151811061182f5761182f615521565b602002602001015182828151811061184957611849615521565b63ffffffff909216602092830291909101909101528061186881615628565b9150506117f5565b508084606001518460ff168151811061188b5761188b615521565b6020026020010181905250505080806118a39061595a565b915050611475565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119109190615455565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611943908b908b908e9060040161597a565b600060405180830381865afa158015611960573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611988919081019061585d565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156119e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a0691906154bc565b611a225760405162461bcd60e51b8152600401610671906154d9565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611a939291906159a4565b600060405180830381865afa158015611ab0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ad8919081019061585d565b9050600084516001600160401b03811115611af557611af5614659565b604051908082528060200260200182016040528015611b1e578160200160208202803683370190505b50905060005b8551811015611c1f57866001600160a01b03166304ec6351878381518110611b4e57611b4e615521565b602002602001015187868581518110611b6957611b69615521565b60200260200101516040518463ffffffff1660e01b8152600401611ba69392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611bc3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611be79190615914565b6001600160c01b0316828281518110611c0257611c02615521565b602090810291909101015280611c1781615628565b915050611b24565b5095945050505050565b6040805180820190915260608082526020820152600084611ca05760405162461bcd60e51b81526020600482015260376024820152600080516020615dcd83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610671565b60408301515185148015611cb8575060a08301515185145b8015611cc8575060c08301515185145b8015611cd8575060e08301515185145b611d425760405162461bcd60e51b81526020600482015260416024820152600080516020615dcd83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610671565b82515160208401515114611dba5760405162461bcd60e51b815260206004820152604460248201819052600080516020615dcd833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610671565b4363ffffffff168463ffffffff1610611e295760405162461bcd60e51b815260206004820152603c6024820152600080516020615dcd83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610671565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611e6a57611e6a614659565b604051908082528060200260200182016040528015611e93578160200160208202803683370190505b506020820152866001600160401b03811115611eb157611eb1614659565b604051908082528060200260200182016040528015611eda578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611f0e57611f0e614659565b604051908082528060200260200182016040528015611f37578160200160208202803683370190505b5081526020860151516001600160401b03811115611f5757611f57614659565b604051908082528060200260200182016040528015611f80578160200160208202803683370190505b50816020018190525060006120528a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612029573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061204d91906159f8565b613c8b565b905060005b8760200151518110156122ce5761207d8860200151828151811061115757611157615521565b8360200151828151811061209357612093615521565b602090810291909101015280156121535760208301516120b4600183615a15565b815181106120c4576120c4615521565b602002602001015160001c836020015182815181106120e5576120e5615521565b602002602001015160001c11612153576040805162461bcd60e51b8152602060048201526024810191909152600080516020615dcd83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610671565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061219857612198615521565b60200260200101518b8b6000015185815181106121b7576121b7615521565b60200260200101516040518463ffffffff1660e01b81526004016121f49392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612211573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122359190615914565b6001600160c01b03168360000151828151811061225457612254615521565b6020026020010181815250506122ba6108e061228e848660000151858151811061228057612280615521565b602002602001015116613d1c565b8a6020015184815181106122a4576122a4615521565b6020026020010151613d4790919063ffffffff16565b9450806122c681615628565b915050612057565b50506122d983613e2b565b60975490935060ff166000816122f0576000612372565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561234e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123729190615a2c565b905060005b8a8110156129f05782156124d2578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106123ce576123ce615521565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa15801561240e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124329190615a2c565b61243c9190615a45565b116124d25760405162461bcd60e51b81526020600482015260666024820152600080516020615dcd83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610671565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061251357612513615521565b9050013560f81c60f81b60f81c8c8c60a00151858151811061253757612537615521565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612593573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125b79190615a5d565b6001600160401b0319166125da8a60400151838151811061115757611157615521565b67ffffffffffffffff1916146126765760405162461bcd60e51b81526020600482015260616024820152600080516020615dcd83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610671565b6126a68960400151828151811061268f5761268f615521565b60200260200101518761373890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106126e9576126e9615521565b9050013560f81c60f81b60f81c8c8c60c00151858151811061270d5761270d615521565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612769573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061278d91906155e9565b856020015182815181106127a3576127a3615521565b6001600160601b039092166020928302919091018201528501518051829081106127cf576127cf615521565b6020026020010151856000015182815181106127ed576127ed615521565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156129db576128658660000151828151811061283757612837615521565b60200260200101518f8f8681811061285157612851615521565b600192013560f81c9290921c811614919050565b156129c9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106128ab576128ab615521565b9050013560f81c60f81b60f81c8e896020015185815181106128cf576128cf615521565b60200260200101518f60e0015188815181106128ed576128ed615521565b6020026020010151878151811061290657612906615521565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561296a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061298e91906155e9565b87518051859081106129a2576129a2615521565b602002602001018181516129b69190615a88565b6001600160601b03169052506001909101905b806129d381615628565b915050612811565b505080806129e890615628565b915050612377565b505050600080612a0a8c868a606001518b608001516107c5565b9150915081612a7b5760405162461bcd60e51b81526020600482015260436024820152600080516020615dcd83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610671565b80612adc5760405162461bcd60e51b81526020600482015260396024820152600080516020615dcd83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610671565b50506000878260200151604051602001612af79291906157d8565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612b29613ec6565b612b336000613f20565b565b60c954606090612b4a9063ffffffff16613b40565b905090565b63ffffffff8216600090815260cb60205260408120815b8160040154811015612bf05783604051602001612b839190615ab0565b60408051601f1981840301815282825280516020918201206000858152600387018352929092209192612bb7929101615b01565b604051602081830303815290604052805190602001201415612bde57600192505050610df0565b80612be881615628565b915050612b66565b506000949350505050565b60ca546001600160a01b03163314612c455760405162461bcd60e51b815260206004820152600d60248201526c2737ba1033b2b732b930ba37b960991b6044820152606401610671565b600060405180608001604052808663ffffffff1681526020014363ffffffff16815260200184848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525063ffffffff8616602091820152604051919250612cc291839101615b99565b60408051601f19818403018152828252805160209182012063ffffffff8916600081815260cb9093529290912055907ff31388a74ee5033e008e2f1789f853591ac3bb7c976ec29b94bdcc9fd46a7ca590612d1e908490615b99565b60405180910390a2505060c9805463ffffffff191663ffffffff94909416939093179092555050565b60c95464010000000090046001600160a01b03163314612d9a5760405162461bcd60e51b815260206004820152600e60248201526d2737ba1030b3b3b932b3b0ba37b960911b6044820152606401610671565b600060cb81612dac60208701876149b1565b63ffffffff1663ffffffff168152602001908152602001600020905083604051602001612dd99190615beb565b60405160208183030381529060405280519060200120816000015414612e305760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964207461736b60a01b6044820152606401610671565b600181015415612e765760405162461bcd60e51b8152602060048201526011602482015270105b1c9958591e481c995cdc1bdb991959607a1b6044820152606401610671565b7f0000000000000000000000000000000000000000000000000000000000000000612ea760408601602087016149b1565b612eb191906157b0565b63ffffffff16431115612ef15760405162461bcd60e51b8152602060048201526008602482015267546f6f206c61746560c01b6044820152606401610671565b600083604051602001612f049190615c67565b604051602081830303815290604052805190602001209050600080612f4983888060400190612f339190615c7a565b612f4360408c0160208d016149b1565b89611c29565b9150915060005b612f5d6040890189615c7a565b905081101561303257612f766080890160608a016149b1565b63ffffffff1683602001518281518110612f9257612f92615521565b6020026020010151612fa49190615cc0565b6001600160601b0316604384600001518381518110612fc557612fc5615521565b60200260200101516001600160601b0316612fe09190615cef565b10156130205760405162461bcd60e51b815260206004820152600f60248201526e10995b1bddc81d1a1c995cda1bdb19608a1b6044820152606401610671565b8061302a81615628565b915050612f50565b5061305561304360208901896149b1565b6130506020890189615d0e565b613f72565b60408051808201825263ffffffff43168152602080820184905291519091613081918991849101615d57565b60408051601f1981840301815290829052805160209091012060018701557fb3d0a4ff7bf22dfcf9cb480d1e535911448d11b47a0fa850ff82f39bf3ce3567906130ce9089908490615d57565b60405180910390a15050505050505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061311b5761311b615521565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061315790889086906004016159a4565b600060405180830381865afa158015613174573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261319c919081019061585d565b6000815181106131ae576131ae615521565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561321a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061323e9190615914565b6001600160c01b0316905060006132548261402e565b9050816132628a838a61094f565b9550955050505050935093915050565b61327a613ec6565b6001600160a01b0381166132df5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610671565b61068381613f20565b600054610100900460ff16158080156133085750600054600160ff909116105b806133225750303b158015613322575060005460ff166001145b6133855760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610671565b6000805460ff1916600117905580156133a8576000805461ff0019166101001790555b6133b38560006140fa565b6133bc84613f20565b60c98054640100000000600160c01b0319166401000000006001600160a01b03868116919091029190911790915560ca80546001600160a01b0319169184169190911790558015613447576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134c59190615455565b6001600160a01b0316336001600160a01b0316146134f55760405162461bcd60e51b815260040161067190615472565b6066541981196066541916146135735760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610671565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107ba565b6001600160a01b0381166136385760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610671565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526136bd61449b565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156136f0576136f2565bfe5b50806137305760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610671565b505092915050565b60408051808201909152600080825260208201526137546144b9565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156136f05750806137305760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610671565b6137d46144d7565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806138bc600080516020615dad83398151915286615537565b90505b6138c8816141e4565b9093509150600080516020615dad833981519152828309831415613902576040805180820190915290815260208101919091529392505050565b600080516020615dad8339815191526001820890506138bf565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061394e6144fc565b60005b6002811015613b13576000613967826006615cef565b905084826002811061397b5761397b615521565b6020020151518361398d836000615a45565b600c811061399d5761399d615521565b60200201528482600281106139b4576139b4615521565b602002015160200151838260016139cb9190615a45565b600c81106139db576139db615521565b60200201528382600281106139f2576139f2615521565b6020020151515183613a05836002615a45565b600c8110613a1557613a15615521565b6020020152838260028110613a2c57613a2c615521565b6020020151516001602002015183613a45836003615a45565b600c8110613a5557613a55615521565b6020020152838260028110613a6c57613a6c615521565b602002015160200151600060028110613a8757613a87615521565b602002015183613a98836004615a45565b600c8110613aa857613aa8615521565b6020020152838260028110613abf57613abf615521565b602002015160200151600160028110613ada57613ada615521565b602002015183613aeb836005615a45565b600c8110613afb57613afb615521565b60200201525080613b0b81615628565b915050613951565b50613b1c61451b565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b63ffffffff8116600090815260cb602052604081206004810154606092906001600160401b03811115613b7557613b75614659565b604051908082528060200260200182016040528015613ba857816020015b6060815260200190600190039081613b935790505b50905060005b8260040154811015613c8357600081815260038401602052604090208054613bd590615acc565b80601f0160208091040260200160405190810160405280929190818152602001828054613c0190615acc565b8015613c4e5780601f10613c2357610100808354040283529160200191613c4e565b820191906000526020600020905b815481529060010190602001808311613c3157829003601f168201915b5050505050828281518110613c6557613c65615521565b60200260200101819052508080613c7b90615628565b915050613bae565b509392505050565b600080613c9784614266565b9050808360ff166001901b11613d155760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610671565b9392505050565b6000805b8215610df057613d31600184615a15565b9092169180613d3f81615d8a565b915050613d20565b60408051808201909152600080825260208201526102008261ffff1610613da35760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610671565b8161ffff1660011415613db7575081610df0565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613e2057600161ffff871660ff83161c81161415613e0357613e008484613738565b93505b613e0d8384613738565b92506201fffe600192831b169101613dd3565b509195945050505050565b60408051808201909152600080825260208201528151158015613e5057506020820151155b15613e6e575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615dad8339815191528460200151613ea19190615537565b613eb990600080516020615dad833981519152615a15565b905292915050565b919050565b6033546001600160a01b03163314612b335760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610671565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b63ffffffff8316600090815260cb6020526040812082915b82811015613fe457848482818110613fa457613fa4615521565b9050602002810190613fb69190615c7a565b60008381526003850160205260409020613fd1929091614539565b5080613fdc81615628565b915050613f8a565b506004810182905560405182815263ffffffff8616907f5a04eb24e83275653ac5d71ef5205698d4ca126250977a126ab5b124d17f2c399060200160405180910390a25050505050565b606060008061403c84613d1c565b61ffff166001600160401b0381111561405757614057614659565b6040519080825280601f01601f191660200182016040528015614081576020820181803683370190505b5090506000805b825182108015614099575061010081105b156140f0576001811b9350858416156140e0578060f81b8383815181106140c2576140c2615521565b60200101906001600160f81b031916908160001a9053508160010191505b6140e981615628565b9050614088565b5090949350505050565b6065546001600160a01b031615801561411b57506001600160a01b03821615155b61419d5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610671565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26141e0826135aa565b5050565b60008080600080516020615dad8339815191526003600080516020615dad83398151915286600080516020615dad83398151915288890909089050600061425a827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615dad8339815191526143f3565b91959194509092505050565b6000610100825111156142ef5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610671565b81516142fd57506000919050565b6000808360008151811061431357614313615521565b0160200151600160f89190911c81901b92505b84518110156143ea5784818151811061434157614341615521565b0160200151600160f89190911c1b91508282116143d65760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610671565b918117916143e381615628565b9050614326565b50909392505050565b6000806143fe61451b565b6144066145bd565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156136f05750826144905760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610671565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806144ea6145db565b81526020016144f76145db565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b82805461454590615acc565b90600052602060002090601f01602090048101928261456757600085556145ad565b82601f106145805782800160ff198235161785556145ad565b828001600101855582156145ad579182015b828111156145ad578235825591602001919060010190614592565b506145b99291506145f9565b5090565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b5b808211156145b957600081556001016145fa565b6001600160a01b038116811461068357600080fd5b60006020828403121561463557600080fd5b8135613d158161460e565b60006020828403121561465257600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561469157614691614659565b60405290565b60405161010081016001600160401b038111828210171561469157614691614659565b604051601f8201601f191681016001600160401b03811182821017156146e2576146e2614659565b604052919050565b6000604082840312156146fc57600080fd5b61470461466f565b9050813581526020820135602082015292915050565b600082601f83011261472b57600080fd5b604051604081018181106001600160401b038211171561474d5761474d614659565b806040525080604084018581111561476457600080fd5b845b81811015613e20578035835260209283019201614766565b60006080828403121561479057600080fd5b61479861466f565b90506147a4838361471a565b81526147b3836040840161471a565b602082015292915050565b60008060008061012085870312156147d557600080fd5b843593506147e686602087016146ea565b92506147f5866060870161477e565b91506148048660e087016146ea565b905092959194509250565b60006001600160401b0383111561482857614828614659565b61483b601f8401601f19166020016146ba565b905082815283838301111561484f57600080fd5b828260208301376000602084830101529392505050565b63ffffffff8116811461068357600080fd5b8035613ec181614866565b60008060006060848603121561489857600080fd5b83356148a38161460e565b925060208401356001600160401b038111156148be57600080fd5b8401601f810186136148cf57600080fd5b6148de8682356020840161480f565b92505060408401356148ef81614866565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614990578385038a52825180518087529087019087870190845b8181101561497b57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614937565b50509a87019a95505091850191600101614919565b509298975050505050505050565b602081526000613d1560208301846148fa565b6000602082840312156149c357600080fd5b8135613d1581614866565b60005b838110156149e95781810151838201526020016149d1565b838111156149f8576000848401525b50505050565b60008151808452614a168160208601602086016149ce565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015614a7f57603f19888603018452614a6d8583516149fe565b94509285019290850190600101614a51565b5092979650505050505050565b801515811461068357600080fd5b600060208284031215614aac57600080fd5b8135613d1581614a8c565b600060808284031215614ac957600080fd5b50919050565b600060408284031215614ac957600080fd5b60006001600160401b03821115614afa57614afa614659565b5060051b60200190565b600082601f830112614b1557600080fd5b81356020614b2a614b2583614ae1565b6146ba565b82815260069290921b84018101918181019086841115614b4957600080fd5b8286015b84811015614b6d57614b5f88826146ea565b835291830191604001614b4d565b509695505050505050565b60008060008060a08587031215614b8e57600080fd5b84356001600160401b0380821115614ba557600080fd5b614bb188838901614ab7565b95506020870135915080821115614bc757600080fd5b614bd388838901614acf565b9450614be28860408901614acf565b93506080870135915080821115614bf857600080fd5b50614c0587828801614b04565b91505092959194509250565b60008083601f840112614c2357600080fd5b5081356001600160401b03811115614c3a57600080fd5b602083019150836020828501011115614c5257600080fd5b9250929050565b60008060008060008060808789031215614c7257600080fd5b8635614c7d8161460e565b95506020870135614c8d81614866565b945060408701356001600160401b0380821115614ca957600080fd5b614cb58a838b01614c11565b90965094506060890135915080821115614cce57600080fd5b818901915089601f830112614ce257600080fd5b813581811115614cf157600080fd5b8a60208260051b8501011115614d0657600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614d5257815163ffffffff1687529582019590820190600101614d30565b509495945050505050565b600060208083528351608082850152614d7960a0850182614d1c565b905081850151601f1980868403016040870152614d968383614d1c565b92506040870151915080868403016060870152614db38383614d1c565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614e0a5784878303018452614df8828751614d1c565b95880195938801939150600101614dde565b509998505050505050505050565b60ff8116811461068357600080fd5b600060208284031215614e3957600080fd5b8135613d1581614e18565b600080600060608486031215614e5957600080fd5b8335614e648161460e565b92506020848101356001600160401b03811115614e8057600080fd5b8501601f81018713614e9157600080fd5b8035614e9f614b2582614ae1565b81815260059190911b82018301908381019089831115614ebe57600080fd5b928401925b82841015614edc57833582529284019290840190614ec3565b8096505050505050614ef060408501614878565b90509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614f3157835183529284019291840191600101614f15565b50909695505050505050565b600082601f830112614f4e57600080fd5b81356020614f5e614b2583614ae1565b82815260059290921b84018101918181019086841115614f7d57600080fd5b8286015b84811015614b6d578035614f9481614866565b8352918301918301614f81565b600082601f830112614fb257600080fd5b81356020614fc2614b2583614ae1565b82815260059290921b84018101918181019086841115614fe157600080fd5b8286015b84811015614b6d5780356001600160401b038111156150045760008081fd5b6150128986838b0101614f3d565b845250918301918301614fe5565b6000610180828403121561503357600080fd5b61503b614697565b905081356001600160401b038082111561505457600080fd5b61506085838601614f3d565b8352602084013591508082111561507657600080fd5b61508285838601614b04565b6020840152604084013591508082111561509b57600080fd5b6150a785838601614b04565b60408401526150b9856060860161477e565b60608401526150cb8560e086016146ea565b60808401526101208401359150808211156150e557600080fd5b6150f185838601614f3d565b60a084015261014084013591508082111561510b57600080fd5b61511785838601614f3d565b60c084015261016084013591508082111561513157600080fd5b5061513e84828501614fa1565b60e08301525092915050565b60008060008060006080868803121561516257600080fd5b8535945060208601356001600160401b038082111561518057600080fd5b61518c89838a01614c11565b9096509450604088013591506151a182614866565b909250606087013590808211156151b757600080fd5b506151c488828901615020565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614d525781516001600160601b0316875295820195908201906001016151e5565b604081526000835160408084015261522560808401826151d1565b90506020850151603f1984830301606085015261524282826151d1565b925050508260208301529392505050565b6000806040838503121561526657600080fd5b823561527181614866565b915060208301356001600160401b0381111561528c57600080fd5b8301601f8101851361529d57600080fd5b6152ac8582356020840161480f565b9150509250929050565b600080600080606085870312156152cc57600080fd5b84356152d781614866565b935060208501356152e781614866565b925060408501356001600160401b0381111561530257600080fd5b61530e87828801614c11565b95989497509550505050565b60008060006060848603121561532f57600080fd5b83356001600160401b038082111561534657600080fd5b61535287838801614ab7565b9450602086013591508082111561536857600080fd5b61537487838801614acf565b9350604086013591508082111561538a57600080fd5b5061539786828701615020565b9150509250925092565b6000806000606084860312156153b657600080fd5b83356153c18161460e565b92506020840135915060408401356148ef81614866565b8281526040602082015260006153f160408301846148fa565b949350505050565b6000806000806080858703121561540f57600080fd5b843561541a8161460e565b9350602085013561542a8161460e565b9250604085013561543a8161460e565b9150606085013561544a8161460e565b939692955090935050565b60006020828403121561546757600080fd5b8151613d158161460e565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156154ce57600080fd5b8151613d1581614a8c565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261555457634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561556c57600080fd5b82516001600160401b0381111561558257600080fd5b8301601f8101851361559357600080fd5b80516155a1614b2582614ae1565b81815260059190911b820183019083810190878311156155c057600080fd5b928401925b828410156155de578351825292840192908401906155c5565b979650505050505050565b6000602082840312156155fb57600080fd5b81516001600160601b0381168114613d1557600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561563c5761563c615612565b5060010190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e1984360301811261568357600080fd5b83016020810192503590506001600160401b038111156156a257600080fd5b803603831315614c5257600080fd5b60006040830182356156c281614866565b63ffffffff16845260208381013536859003601e190181126156e357600080fd5b840181810190356001600160401b038111156156fe57600080fd5b8060051b80360387131561571157600080fd5b6040848901529381905260609387018401938290880160005b8381101561576457898703605f19018252615745838661566c565b615750898284615643565b98505050918501919085019060010161572a565b509498975050505050505050565b60608152600061578560608301856156b1565b9050823561579281614866565b63ffffffff8116602084015250602083013560408301529392505050565b600063ffffffff8083168185168083038211156157cf576157cf615612565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015614a7f578151855293820193908201906001016157f7565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561584057600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561587057600080fd5b82516001600160401b0381111561588657600080fd5b8301601f8101851361589757600080fd5b80516158a5614b2582614ae1565b81815260059190911b820183019083810190878311156158c457600080fd5b928401925b828410156155de5783516158dc81614866565b825292840192908401906158c9565b63ffffffff8416815260406020820152600061590b604083018486615643565b95945050505050565b60006020828403121561592657600080fd5b81516001600160c01b0381168114613d1557600080fd5b60006020828403121561594f57600080fd5b8151613d1581614866565b600060ff821660ff81141561597157615971615612565b60010192915050565b60408152600061598e604083018587615643565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156159eb578451835293830193918301916001016159cf565b5090979650505050505050565b600060208284031215615a0a57600080fd5b8151613d1581614e18565b600082821015615a2757615a27615612565b500390565b600060208284031215615a3e57600080fd5b5051919050565b60008219821115615a5857615a58615612565b500190565b600060208284031215615a6f57600080fd5b815167ffffffffffffffff1981168114613d1557600080fd5b60006001600160601b0383811690831681811015615aa857615aa8615612565b039392505050565b60008251615ac28184602087016149ce565b9190910192915050565b600181811c90821680615ae057607f821691505b60208210811415614ac957634e487b7160e01b600052602260045260246000fd5b600080835481600182811c915080831680615b1d57607f831692505b6020808410821415615b3d57634e487b7160e01b86526022600452602486fd5b818015615b515760018114615b6257615764565b60ff19861689528489019650615764565b60008a81526020902060005b86811015615b875781548b820152908501908301615b6e565b50505096909201979650505050505050565b60208152600063ffffffff808451166020840152806020850151166040840152604084015160806060850152615bd260a08501826149fe565b9050816060860151166080850152809250505092915050565b6020815260008235615bfc81614866565b63ffffffff808216602085015260208501359150615c1982614866565b8082166040850152615c2e604086018661566c565b925060806060860152615c4560a086018483615643565b9250506060850135615c5681614866565b166080939093019290925250919050565b602081526000613d1560208301846156b1565b6000808335601e19843603018112615c9157600080fd5b8301803591506001600160401b03821115615cab57600080fd5b602001915036819003821315614c5257600080fd5b60006001600160601b0380831681851681830481118215151615615ce657615ce6615612565b02949350505050565b6000816000190483118215151615615d0957615d09615612565b500290565b6000808335601e19843603018112615d2557600080fd5b8301803591506001600160401b03821115615d3f57600080fd5b6020019150600581901b3603821315614c5257600080fd5b606081526000615d6a60608301856156b1565b905063ffffffff8351166020830152602083015160408301529392505050565b600061ffff80831681811415615da257615da2615612565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212206f274708482c05e014828478f0a78440d7bf65337be34aa2b68fe53fa806409364736f6c634300080c0033",
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
