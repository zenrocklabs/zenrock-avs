// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractZrRegistryCoordinator

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

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IRegistryCoordinatorOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// IRegistryCoordinatorOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// IRegistryCoordinatorOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// IRegistryCoordinatorQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// IStakeRegistryStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ContractZrRegistryCoordinatorMetaData contains all meta data concerning the ContractZrRegistryCoordinator contract.
var ContractZrRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"contractIZrServiceManager\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"},{\"name\":\"_socketRegistry\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEjectionTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithChurn\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"churnApproverSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIZrServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjectionCooldown\",\"inputs\":[{\"name\":\"_ejectionCooldown\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"socketRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101e06040523480156200001257600080fd5b50604051620060c3380380620060c3833981016040819052620000359162000240565b604080518082018252601681527f4156535265676973747279436f6f7264696e61746f7200000000000000000000602080830191825283518085018552600681526576302e302e3160d01b908201529151902060e08190527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a996101008190524660a081815285517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818701819052818801959095526060810193909352608080840192909252308382018190528651808503909201825260c09384019096528051940193909320909252919052610120526001600160a01b038086166101405280851661018052808416610160528083166101a05281166101c0526200015a62000165565b5050505050620002c0565b600054610100900460ff1615620001d25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000225576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200023d57600080fd5b50565b600080600080600060a086880312156200025957600080fd5b8551620002668162000227565b6020870151909550620002798162000227565b60408701519094506200028c8162000227565b60608701519093506200029f8162000227565b6080870151909250620002b28162000227565b809150509295509295909350565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c051615ce9620003da60003960008181610848015261221f0152600081816106c60152818161125a01528181611e1201528181612f0201528181613490015261377401526000818161061e01528181611d9f0152818161215901528181612e82015281816133e7015281816136f301526143180152600081816105e401528181610f2201528181611ddd015281816124b30152818161252d01528181612e040152818161336701526137f001526000818161052801528181612d5a01526132a301526000613d6501526000613db401526000613d8f01526000613ce801526000613d1201526000613d3c0152615ce96000f3fe608060405234801561001057600080fd5b50600436106102f05760003560e01c80635c975abb1161019d578063a96f783e116100e9578063d75b4c88116100a2578063f2fde38b1161007c578063f2fde38b1461086a578063f7013ef61461087d578063fabc1cbc14610890578063fd39105a146108a357600080fd5b8063d75b4c881461078d578063e65797ad146107a0578063ea32afae1461084357600080fd5b8063a96f783e1461070f578063ae42cdb214610718578063c391425e1461072b578063ca0de8821461074b578063ca4f2d9714610772578063d72d8dd61461078557600080fd5b806384ca5213116101565780638da5cb5b116101305780638da5cb5b1461069a5780639aa1653d146106a25780639e9923c2146106c15780639feab859146106e857600080fd5b806384ca52131461065b578063871ef0491461066e578063886f11951461068157600080fd5b80635c975abb146105d75780635df45946146105df5780636347c9001461060657806368304835146106195780636e3b17db14610640578063715018a61461065357600080fd5b8063249a0c421161025c5780633998fdd3116102155780635865c60c116101ef5780635865c60c1461057d578063595c6a671461059d5780635ac86ab7146105a55780635b0b829f146105c457600080fd5b80633998fdd3146105235780633c2a7f4c1461054a5780635140a5481461056a57600080fd5b8063249a0c42146104a457806328f61b31146104c4578063296bb064146104d757806329d1e0c3146104ea5780632b21ea6d146104fd5780632cdd1e861461051057600080fd5b806310d67a2f116102ae57806310d67a2f146103b9578063125e0584146103cc57806313542a4e146103ec578063136439dd146104155780631478851f146104285780631eb812da1461045b57600080fd5b8062cf2ab5146102f557806303fd34921461030a57806304ec63511461033d578063054310e6146103685780630cf4b767146103935780630d3f2134146103a6575b600080fd5b61030861030336600461498c565b6108df565b005b61032a6103183660046149cd565b60009081526098602052604090205490565b6040519081526020015b60405180910390f35b61035061034b3660046149f8565b6109f4565b6040516001600160c01b039091168152602001610334565b609d5461037b906001600160a01b031681565b6040516001600160a01b039091168152602001610334565b6103086103a1366004614b17565b610bcd565b6103086103b43660046149cd565b610c77565b6103086103c7366004614b7c565b610c84565b61032a6103da366004614b7c565b609f6020526000908152604090205481565b61032a6103fa366004614b7c565b6001600160a01b031660009081526099602052604090205490565b6103086104233660046149cd565b610d34565b61044b6104363660046149cd565b609a6020526000908152604090205460ff1681565b6040519015158152602001610334565b61046e610469366004614b99565b610e78565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b031690820152606001610334565b61032a6104b2366004614bd1565b609b6020526000908152604090205481565b609e5461037b906001600160a01b031681565b61037b6104e53660046149cd565b610f09565b6103086104f8366004614b7c565b610f95565b61030861050b366004614cf9565b610fa6565b61030861051e366004614b7c565b6110d0565b61037b7f000000000000000000000000000000000000000000000000000000000000000081565b61055d610558366004614b7c565b6110e1565b6040516103349190614e21565b610308610578366004614e38565b611160565b61059061058b366004614b7c565b6115ea565b6040516103349190614edb565b61030861165e565b61044b6105b3366004614bd1565b6001805460ff9092161b9081161490565b6103086105d2366004614f60565b61172a565b60015461032a565b61037b7f000000000000000000000000000000000000000000000000000000000000000081565b61037b6106143660046149cd565b61174b565b61037b7f000000000000000000000000000000000000000000000000000000000000000081565b61030861064e366004614f94565b611775565b610308611894565b61032a61066936600461504b565b6118a8565b61035061067c3660046149cd565b6118f2565b60005461037b906201000090046001600160a01b031681565b61037b6118fd565b6096546106af9060ff1681565b60405160ff9091168152602001610334565b61037b7f000000000000000000000000000000000000000000000000000000000000000081565b61032a7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b61032a60a05481565b610308610726366004615118565b611916565b61073e6107393660046151eb565b611ad2565b6040516103349190615290565b61032a7f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a81565b6103086107803660046152da565b611b8b565b609c5461032a565b61030861079b366004615324565b611bf1565b61080f6107ae366004614bd1565b60408051606080820183526000808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff908116918301919091529282015190921690820152606001610334565b61037b7f000000000000000000000000000000000000000000000000000000000000000081565b610308610878366004614b7c565b611c04565b61030861088b366004615413565b611c7a565b61030861089e3660046149cd565b611e88565b6108d26108b1366004614b7c565b6001600160a01b031660009081526099602052604090206001015460ff1690565b6040516103349190615477565b6001546002906004908116036109105760405162461bcd60e51b815260040161090790615485565b60405180910390fd5b60005b828110156109ee57600084848381811061092f5761092f6154bc565b90506020020160208101906109449190614b7c565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff16600281111561098f5761098f614ea3565b60028111156109a0576109a0614ea3565b905250805190915060006109b382611fe4565b905060006109c9826001600160c01b031661204f565b90506109d685858361211b565b505050505080806109e6906154e8565b915050610913565b50505050565b6000838152609860205260408120805482919084908110610a1757610a176154bc565b600091825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b03169183019190915290925085161015610b035760405162461bcd60e51b815260206004820152605a60248201527f526567436f6f72642e67657451756f72756d4269746d61704174426c6f636b4e60448201527f756d6265724279496e6465783a2071756f72756d4269746d617055706461746560648201527f2069732066726f6d20616674657220626c6f636b4e756d626572000000000000608482015260a401610907565b602081015163ffffffff161580610b295750806020015163ffffffff168463ffffffff16105b610bc15760405162461bcd60e51b815260206004820152605b60248201527f526567436f6f72642e67657451756f72756d4269746d61704174426c6f636b4e60448201527f756d6265724279496e6465783a2071756f72756d4269746d617055706461746560648201527f2069732066726f6d206265666f726520626c6f636b4e756d6265720000000000608482015260a401610907565b60400151949350505050565b60013360009081526099602052604090206001015460ff166002811115610bf657610bf6614ea3565b14610c5a5760405162461bcd60e51b815260206004820152602e60248201527f526567436f6f72642e757064617465536f636b65743a206f70657261746f722060448201526d1b9bdd081c9959da5cdd195c995960921b6064820152608401610907565b33600090815260996020526040902054610c749082612208565b50565b610c7f6122c4565b60a055565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cd7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cfb9190615501565b6001600160a01b0316336001600160a01b031614610d2b5760405162461bcd60e51b81526004016109079061551e565b610c7481612323565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610d81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da59190615568565b610dc15760405162461bcd60e51b81526004016109079061558a565b60015481811614610e3a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610907565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60408051606081018252600080825260208201819052918101919091526000838152609860205260409020805483908110610eb557610eb56154bc565b600091825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610f71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f039190615501565b610f9d6122c4565b610c7481612428565b6001805460009190811603610fcd5760405162461bcd60e51b815260040161090790615485565b838b146110425760405162461bcd60e51b815260206004820152603960248201527f526567436f6f72642e72656769737465724f70657261746f725769746843687560448201527f726e3a20696e707574206c656e677468206d69736d61746368000000000000006064820152608401610907565b600061104e3388612491565b90506110ae33828888808060200260200160405190810160405280939291908181526020016000905b828210156110a357611094604083028601368190038101906155d2565b81526020019060010190611077565b5050505050876125cd565b6110c181338f8f8f8f8f8f8e8e8d612744565b50505050505050505050505050565b6110d86122c4565b610c74816127e0565b6040805180820190915260008082526020820152610f0361115b7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de6846040516020016111409291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120612849565b612897565b6001546002906004908116036111885760405162461bcd60e51b815260040161090790615485565b8382146111eb5760405162461bcd60e51b81526020600482015260386024820152600080516020615c9483398151915260448201527f6d3a20696e707574206c656e677468206d69736d6174636800000000000000006064820152608401610907565b60005b828110156115e257600084848381811061120a5761120a6154bc565b919091013560f81c9150369050600088888581811061122b5761122b6154bc565b905060200281019061123d91906155ee565b6040516379a0849160e11b815260ff8616600482015291935091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f341092290602401602060405180830381865afa1580156112a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112cd9190615637565b63ffffffff16811461135b5760405162461bcd60e51b815260206004820152605a6024820152600080516020615c9483398151915260448201527f6d3a206e756d626572206f662075706461746564206f70657261746f7273206460648201527f6f6573206e6f74206d617463682071756f72756d20746f74616c000000000000608482015260a401610907565b6000805b8281101561158157600084848381811061137b5761137b6154bc565b90506020020160208101906113909190614b7c565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff1660028111156113db576113db614ea3565b60028111156113ec576113ec614ea3565b905250805190915060006113ff82611fe4565b905060016001600160c01b03821660ff8b161c8116146114755760405162461bcd60e51b81526020600482015260396024820152600080516020615c9483398151915260448201527f6d3a206f70657261746f72206e6f7420696e2071756f72756d000000000000006064820152608401610907565b856001600160a01b0316846001600160a01b0316116115105760405162461bcd60e51b815260206004820152605c6024820152600080516020615c9483398151915260448201527f6d3a206f70657261746f7273206172726179206d75737420626520736f72746560648201527f6420696e20617363656e64696e672061646472657373206f7264657200000000608482015260a401610907565b5061156b83838e8c8f611524826001615654565b9261153193929190615667565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061211b92505050565b5090925061157a9050816154e8565b905061135f565b5060ff84166000818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a250505050806115db906154e8565b90506111ee565b505050505050565b60408051808201909152600080825260208201526001600160a01b0382166000908152609960209081526040918290208251808401909352805483526001810154909183019060ff16600281111561164457611644614ea3565b600281111561165557611655614ea3565b90525092915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156116ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116cf9190615568565b6116eb5760405162461bcd60e51b81526004016109079061558a565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6117326122c4565b8161173c81612926565b6117468383612993565b505050565b609c818154811061175b57600080fd5b6000918252602090912001546001600160a01b0316905081565b61177d612a39565b6001600160a01b0383166000908152609f602090815260408083204290556099825280832080548251601f87018590048502810185019093528583529093909290916117ea9187908790819084018382808284376000920191909152505060965460ff169150612aab9050565b905060006117f783611fe4565b905060018085015460ff16600281111561181357611813614ea3565b14801561182857506001600160c01b03821615155b801561184657506118466001600160c01b0383811690831681161490565b1561188b5761188b8787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612b3592505050565b50505050505050565b61189c6122c4565b6118a66000612f76565b565b60006118e87f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a878787878760405160200161114096959493929190615691565b9695505050505050565b6000610f0382611fe4565b60006119116064546001600160a01b031690565b905090565b600180546000919081160361193d5760405162461bcd60e51b815260040161090790615485565b60006119493385612491565b905060006119d733838d8d8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508c8c8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250612fc8915050565b51905060005b8a811015611ac45760008c8c838181106119f9576119f96154bc565b919091013560f81c600081815260976020526040902054855191935063ffffffff169150849084908110611a2f57611a2f6154bc565b602002602001015163ffffffff161115611ab15760405162461bcd60e51b815260206004820152603960248201527f526567436f6f72642e72656769737465724f70657261746f723a206f7065726160448201527f746f7220636f756e742065786365656473206d6178696d756d000000000000006064820152608401610907565b5080611abc816154e8565b9150506119dd565b505050505050505050505050565b6060600082516001600160401b03811115611aef57611aef614a30565b604051908082528060200260200182016040528015611b18578160200160208202803683370190505b50905060005b8351811015611b8357611b4a85858381518110611b3d57611b3d6154bc565b602002602001015161351f565b828281518110611b5c57611b5c6154bc565b63ffffffff9092166020928302919091019091015280611b7b816154e8565b915050611b1e565b509392505050565b60018054600290811603611bb15760405162461bcd60e51b815260040161090790615485565b6117463384848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612b3592505050565b611bf96122c4565b611746838383613646565b611c0c6122c4565b6001600160a01b038116611c715760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610907565b610c7481612f76565b600054610100900460ff1615808015611c9a5750600054600160ff909116105b80611cb45750303b158015611cb4575060005460ff166001145b611d175760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610907565b6000805460ff191660011790558015611d3a576000805461ff0019166101001790555b611d4386612f76565b611d4d838361385d565b611d5685612428565b611d5f846127e0565b609c80546001818101835560008390527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c91820180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925584548084018655840180547f000000000000000000000000000000000000000000000000000000000000000084169083161790558454928301909455910180547f00000000000000000000000000000000000000000000000000000000000000009092169190921617905580156115e2576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611edb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eff9190615501565b6001600160a01b0316336001600160a01b031614611f2f5760405162461bcd60e51b81526004016109079061551e565b600154198119600154191614611fad5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610907565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610e6d565b6000818152609860205260408120548082036120035750600092915050565b600083815260986020526040902061201c600183615716565b8154811061202c5761202c6154bc565b600091825260209091200154600160401b90046001600160c01b03169392505050565b606060008061205d8461394d565b61ffff166001600160401b0381111561207857612078614a30565b6040519080825280601f01601f1916602001820160405280156120a2576020820181803683370190505b5090506000805b8251821080156120ba575061010081105b15612111576001811b935085841615612101578060f81b8383815181106120e3576120e36154bc565b60200101906001600160f81b031916908160001a9053508160010191505b61210a816154e8565b90506120a9565b5090949350505050565b60018260200151600281111561213357612133614ea3565b1461213d57505050565b81516040516333567f7f60e11b81526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906366acfefe906121929088908690889060040161576f565b6020604051808303816000875af11580156121b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121d5919061579f565b90506001600160c01b0381161561220157612201856121fc836001600160c01b031661204f565b612b35565b5050505050565b6040516378219b3f60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f043367e9061225690859085906004016157c8565b600060405180830381600087803b15801561227057600080fd5b505af1158015612284573d6000803e3d6000fd5b50505050817fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa826040516122b891906157e1565b60405180910390a25050565b336122cd6118fd565b6001600160a01b0316146118a65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610907565b6001600160a01b0381166123b15760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610907565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b6040516309aa152760e11b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa1580156124fc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061252091906157f4565b90506000819003610f03577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bf79ce588484612565876110e1565b6040518463ffffffff1660e01b81526004016125839392919061580d565b6020604051808303816000875af11580156125a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c691906157f4565b9392505050565b6020808201516000908152609a909152604090205460ff16156126685760405162461bcd60e51b815260206004820152604760248201527f526567436f6f72642e5f766572696679436875726e417070726f76657253696760448201527f6e61747572653a20636875726e417070726f7665722073616c7420616c726561606482015266191e481d5cd95960ca1b608482015260a401610907565b42816040015110156126f25760405162461bcd60e51b815260206004820152604760248201527f526567436f6f72642e5f766572696679436875726e417070726f76657253696760448201527f6e61747572653a20636875726e417070726f766572207369676e617475726520606482015266195e1c1a5c995960ca1b608482015260a401610907565b602080820180516000908152609a909252604091829020805460ff19166001179055609d549051918301516109ee926001600160a01b039092169161273d91889188918891906118a8565b8351613978565b60006127d08b8d8c8c8c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b9250612fc8915050565b9050611ac4818c8c8c8888613b32565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610f03612856613cdb565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b6040805180820190915260008082526020820152600080806128c7600080516020615c7483398151915286615893565b90505b6128d381613e02565b9093509150600080516020615c74833981519152828309830361290c576040805180820190915290815260208101919091529392505050565b600080516020615c748339815191526001820890506128ca565b60965460ff90811690821610610c745760405162461bcd60e51b815260206004820152602c60248201527f526567436f6f72642e71756f72756d4578697374733a2071756f72756d20646f60448201526b195cc81b9bdd08195e1a5cdd60a21b6064820152608401610907565b60ff8216600081815260976020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac906060016122b8565b609e546001600160a01b031633146118a65760405162461bcd60e51b815260206004820152602f60248201527f526567436f6f72642e6f6e6c79456a6563746f723a2063616c6c65722069732060448201526e3737ba103a34329032b532b1ba37b960891b6064820152608401610907565b600080612ab784613e84565b9050808360ff166001901b116125c65760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610907565b6001600160a01b0382166000908152609960205260409020805460018083015460ff166002811115612b6957612b69614ea3565b14612bdc5760405162461bcd60e51b815260206004820152603860248201527f526567436f6f72642e5f646572656769737465724f70657261746f723a206f7060448201527f657261746f72206973206e6f74207265676973746572656400000000000000006064820152608401610907565b609654600090612bf090859060ff16612aab565b90506000612bfd83611fe4565b90506001600160c01b038216612c6e5760405162461bcd60e51b815260206004820152603060248201527f526567436f6f72642e5f646572656769737465724f70657261746f723a20626960448201526f0746d61702063616e6e6f7420626520360841b6064820152608401610907565b612c856001600160c01b0383811690831681161490565b612d055760405162461bcd60e51b8152602060048201526044602482018190527f526567436f6f72642e5f646572656769737465724f70657261746f723a206f70908201527f657261746f72206973206e6f74207265676973746572656420666f722071756f60648201526372756d7360e01b608482015260a401610907565b6001600160c01b0382811619821616612d1e8482614014565b6001600160c01b038116612ded5760018501805460ff191660021790556040516351b27a6d60e11b81526001600160a01b0388811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b158015612d9e57600080fd5b505af1158015612db2573d6000803e3d6000fd5b50506040518692506001600160a01b038a1691507f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e490600090a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe590612e3b908a908a906004016158a7565b600060405180830381600087803b158015612e5557600080fd5b505af1158015612e69573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd9150612ebb9087908a906004016157c8565b600060405180830381600087803b158015612ed557600080fd5b505af1158015612ee9573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd9150612f3b9087908a906004016157c8565b600060405180830381600087803b158015612f5557600080fd5b505af1158015612f69573d6000803e3d6000fd5b5050505050505050505050565b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b612fec60405180606001604052806060815260200160608152602001606081525090565b600061303487878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff169150612aab9050565b9050600061304189611fe4565b90506001600160c01b0382166130b05760405162461bcd60e51b815260206004820152602e60248201527f526567436f6f72642e5f72656769737465724f70657261746f723a206269746d60448201526d061702063616e6e6f7420626520360941b6064820152608401610907565b8082166001600160c01b0316156131405760405162461bcd60e51b815260206004820152604860248201527f526567436f6f72642e5f72656769737465724f70657261746f723a206f70657260448201527f61746f7220616c7265616479207265676973746572656420666f7220736f6d656064820152672071756f72756d7360c01b608482015260a401610907565b60a0546001600160a01b038b166000908152609f60205260409020546001600160c01b03838116908516179142916131789190615654565b106131eb5760405162461bcd60e51b815260206004820152603a60248201527f526567436f6f72642e5f72656769737465724f70657261746f723a206f70657260448201527f61746f722063616e6e6f742072657265676973746572207965740000000000006064820152608401610907565b6131f58a82614014565b60016001600160a01b038c1660009081526099602052604090206001015460ff16600281111561322757613227614ea3565b1461335057604080518082019091528a815260208101600190526001600160a01b038c166000908152609960209081526040909120825181559082015160018083018054909160ff199091169083600281111561328657613286614ea3565b02179055505060405163f891899d60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915063f891899d906132dd908e908a908a906004016158cb565b600060405180830381600087803b1580156132f757600080fd5b505af115801561330b573d6000803e3d6000fd5b505050506133198a88612208565b6040518a906001600160a01b038d16907fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe90600090a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb27952906133a0908e908d908d90600401615956565b600060405180830381600087803b1580156133ba57600080fd5b505af11580156133ce573d6000803e3d6000fd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150613424908e908e908e908e9060040161597b565b6000604051808303816000875af1158015613443573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261346b9190810190615a12565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d906134c8908d908d908d90600401615a75565b6000604051808303816000875af11580156134e7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261350f9190810190615a8f565b8452505050979650505050505050565b600081815260986020526040812054815b818110156135b15760016135448284615716565b61354e9190615716565b92508463ffffffff16609860008681526020019081526020016000208463ffffffff1681548110613581576135816154bc565b60009182526020909120015463ffffffff161161359f575050610f03565b806135a9816154e8565b915050613530565b5060405162461bcd60e51b815260206004820152605e60248201527f526567436f6f72642e67657451756f72756d4269746d6170496e64657841744260448201527f6c6f636b4e756d6265723a206e6f206269746d61702075706461746520666f7560648201527f6e6420666f72206f70657261746f7220617420626c6f636b4e756d6265720000608482015260a401610907565b60965460ff1660c081106136af5760405162461bcd60e51b815260206004820152602a60248201527f526567436f6f72642e63726561746551756f72756d3a206d61782071756f72756044820152691b5cc81c995858da195960b21b6064820152608401610907565b6136ba816001615b28565b6096805460ff191660ff92909216919091179055806136d98186612993565b60405160016296b58960e01b031981526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ff694a779061372c90849088908890600401615b41565b600060405180830381600087803b15801561374657600080fd5b505af115801561375a573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b1580156137c257600080fd5b505af11580156137d6573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b15801561383e57600080fd5b505af1158015613852573d6000803e3d6000fd5b505050505050505050565b6000546201000090046001600160a01b031615801561388457506001600160a01b03821615155b6139065760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610907565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261394982612323565b5050565b6000805b8215610f0357613962600184615716565b909216918061397081615bba565b915050613951565b6001600160a01b0383163b15613a9257604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906139b890869086906004016157c8565b602060405180830381865afa1580156139d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139f99190615bdb565b6001600160e01b031916146117465760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a401610907565b826001600160a01b0316613aa683836141d5565b6001600160a01b0316146117465760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a401610907565b60005b8381101561188b576000858583818110613b5157613b516154bc565b919091013560f81c6000818152609760209081526040918290208251606081018452905463ffffffff811680835261ffff600160201b8304811694840194909452600160301b909104909216928101929092528b5180519395509193509185908110613bbf57613bbf6154bc565b602002602001015163ffffffff161115613cc657613c42828a604001518581518110613bed57613bed6154bc565b60200260200101518a8c602001518781518110613c0c57613c0c6154bc565b6020026020010151898989818110613c2657613c266154bc565b905060400201803603810190613c3c91906155d2565b866141f1565b6040805160018082528183019092526000916020820181803683370190505090508260f81b81600081518110613c7a57613c7a6154bc565b60200101906001600160f81b031916908160001a905350613cc4868686818110613ca657613ca66154bc565b9050604002016020016020810190613cbe9190614b7c565b82612b35565b505b50508080613cd3906154e8565b915050613b35565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015613d3457507f000000000000000000000000000000000000000000000000000000000000000046145b15613d5e57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b60008080600080516020615c748339815191526003600080516020615c7483398151915286600080516020615c74833981519152888909090890506000613e78827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615c748339815191526144db565b91959194509092505050565b600061010082511115613f0d5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610907565b8151600003613f1e57506000919050565b60008083600081518110613f3457613f346154bc565b0160200151600160f89190911c81901b92505b845181101561400b57848181518110613f6257613f626154bc565b0160200151600160f89190911c1b9150828211613ff75760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610907565b91811791614004816154e8565b9050613f47565b50909392505050565b600082815260986020526040812054908190036140bc576000838152609860209081526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055505050565b60008381526098602052604081206140d5600184615716565b815481106140e5576140e56154bc565b6000918252602090912001805490915063ffffffff4381169116036141275780546001600160401b0316600160401b6001600160c01b038516021781556109ee565b805463ffffffff438116600160201b81810267ffffffff0000000019909416939093178455600087815260986020908152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff199093169390941692909217179190911691909117905550505050565b60008060006141e48585614584565b91509150611b83816145f2565b6020808301516001600160a01b03808216600081815260999094526040909320549192908716036142775760405162461bcd60e51b815260206004820152602a60248201527f526567436f6f72642e5f76616c6964617465436875726e3a2063616e6e6f742060448201526931b43ab9371039b2b63360b11b6064820152608401610907565b8760ff16846000015160ff16146142f65760405162461bcd60e51b815260206004820152603c60248201527f526567436f6f72642e5f76616c6964617465436875726e3a2071756f72756d4e60448201527f756d626572206e6f74207468652073616d65206173207369676e6564000000006064820152608401610907565b604051635401ed2760e01b81526004810182905260ff891660248201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa158015614367573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061438b9190615c05565b905061439781856147a8565b6001600160601b0316866001600160601b0316116144315760405162461bcd60e51b815260206004820152604b60248201527f526567436f6f72642e5f76616c6964617465436875726e3a20696e636f6d696e60448201527f67206f70657261746f722068617320696e73756666696369656e74207374616b60648201526a32903337b91031b43ab93760a91b608482015260a401610907565b61443b88856147cc565b6001600160601b0316816001600160601b0316106138525760405162461bcd60e51b815260206004820152605160248201527f526567436f6f72642e5f76616c6964617465436875726e3a2063616e6e6f742060448201527f6b69636b206f70657261746f722077697468206d6f7265207468616e206b69636064820152706b424950734f66546f74616c5374616b6560781b608482015260a401610907565b6000806144e661490c565b6144ee61492a565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828061452b57fe5b50826145795760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610907565b505195945050505050565b60008082516041036145ba5760208301516040840151606085015160001a6145ae878285856147e6565b945094505050506145eb565b82516040036145e357602083015160408401516145d88683836148d3565b9350935050506145eb565b506000905060025b9250929050565b600081600481111561460657614606614ea3565b0361460e5750565b600181600481111561462257614622614ea3565b0361466f5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610907565b600281600481111561468357614683614ea3565b036146d05760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610907565b60038160048111156146e4576146e4614ea3565b0361473c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610907565b600481600481111561475057614750614ea3565b03610c745760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610907565b6020810151600090612710906147c29061ffff1685615c22565b6125c69190615c4d565b6040810151600090612710906147c29061ffff1685615c22565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561481d57506000905060036148ca565b8460ff16601b1415801561483557508460ff16601c14155b1561484657506000905060046148ca565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561489a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166148c3576000600192509250506148ca565b9150600090505b94509492505050565b6000806001600160ff1b038316816148f060ff86901c601b615654565b90506148fe878288856147e6565b935093505050935093915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60008083601f84011261495a57600080fd5b5081356001600160401b0381111561497157600080fd5b6020830191508360208260051b85010111156145eb57600080fd5b6000806020838503121561499f57600080fd5b82356001600160401b038111156149b557600080fd5b6149c185828601614948565b90969095509350505050565b6000602082840312156149df57600080fd5b5035919050565b63ffffffff81168114610c7457600080fd5b600080600060608486031215614a0d57600080fd5b833592506020840135614a1f816149e6565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614a6857614a68614a30565b60405290565b604080519081016001600160401b0381118282101715614a6857614a68614a30565b604051601f8201601f191681016001600160401b0381118282101715614ab857614ab8614a30565b604052919050565b60006001600160401b03831115614ad957614ad9614a30565b614aec601f8401601f1916602001614a90565b9050828152838383011115614b0057600080fd5b828260208301376000602084830101529392505050565b600060208284031215614b2957600080fd5b81356001600160401b03811115614b3f57600080fd5b8201601f81018413614b5057600080fd5b614b5f84823560208401614ac0565b949350505050565b6001600160a01b0381168114610c7457600080fd5b600060208284031215614b8e57600080fd5b81356125c681614b67565b60008060408385031215614bac57600080fd5b50508035926020909101359150565b803560ff81168114614bcc57600080fd5b919050565b600060208284031215614be357600080fd5b6125c682614bbb565b60008083601f840112614bfe57600080fd5b5081356001600160401b03811115614c1557600080fd5b6020830191508360208285010111156145eb57600080fd5b60006101008284031215614c4057600080fd5b50919050565b60008083601f840112614c5857600080fd5b5081356001600160401b03811115614c6f57600080fd5b6020830191508360208260061b85010111156145eb57600080fd5b600060608284031215614c9c57600080fd5b614ca4614a46565b905081356001600160401b03811115614cbc57600080fd5b8201601f81018413614ccd57600080fd5b614cdc84823560208401614ac0565b825250602082013560208201526040820135604082015292915050565b60008060008060008060008060008060006101c08c8e031215614d1b57600080fd5b6001600160401b03808d351115614d3157600080fd5b614d3e8e8e358f01614bec565b909c509a5060208d0135811015614d5457600080fd5b614d648e60208f01358f01614bec565b909a50985060408d0135811015614d7a57600080fd5b614d8a8e60408f01358f01614bec565b9098509650614d9c8e60608f01614c2d565b9550806101608e01351115614db057600080fd5b614dc18e6101608f01358f01614c46565b90955093506101808d0135811015614dd857600080fd5b614de98e6101808f01358f01614c8a565b9250806101a08e01351115614dfd57600080fd5b50614e0f8d6101a08e01358e01614c8a565b90509295989b509295989b9093969950565b815181526020808301519082015260408101610f03565b60008060008060408587031215614e4e57600080fd5b84356001600160401b0380821115614e6557600080fd5b614e7188838901614948565b90965094506020870135915080821115614e8a57600080fd5b50614e9787828801614bec565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b60038110614ed757634e487b7160e01b600052602160045260246000fd5b9052565b815181526020808301516040830191614ef690840182614eb9565b5092915050565b803561ffff81168114614bcc57600080fd5b600060608284031215614f2157600080fd5b614f29614a46565b90508135614f36816149e6565b8152614f4460208301614efd565b6020820152614f5560408301614efd565b604082015292915050565b60008060808385031215614f7357600080fd5b614f7c83614bbb565b9150614f8b8460208501614f0f565b90509250929050565b600080600060408486031215614fa957600080fd5b8335614fb481614b67565b925060208401356001600160401b03811115614fcf57600080fd5b614fdb86828701614bec565b9497909650939450505050565b60006001600160401b0382111561500157615001614a30565b5060051b60200190565b60006040828403121561501d57600080fd5b615025614a6e565b905061503082614bbb565b8152602082013561504081614b67565b602082015292915050565b600080600080600060a0868803121561506357600080fd5b853561506e81614b67565b945060208681013594506040808801356001600160401b0381111561509257600080fd5b8801601f81018a136150a357600080fd5b80356150b66150b182614fe8565b614a90565b81815260069190911b8201840190848101908c8311156150d557600080fd5b928501925b828410156150fb576150ec8d8561500b565b825292840192908501906150da565b999c989b5098996060810135995060800135979650505050505050565b600080600080600080600080610180898b03121561513557600080fd5b88356001600160401b038082111561514c57600080fd5b6151588c838d01614bec565b909a50985060208b013591508082111561517157600080fd5b61517d8c838d01614bec565b909850965060408b013591508082111561519657600080fd5b6151a28c838d01614bec565b90965094508491506151b78c60608d01614c2d565b93506101608b01359150808211156151ce57600080fd5b506151db8b828c01614c8a565b9150509295985092959890939650565b600080604083850312156151fe57600080fd5b8235615209816149e6565b91506020838101356001600160401b0381111561522557600080fd5b8401601f8101861361523657600080fd5b80356152446150b182614fe8565b81815260059190911b8201830190838101908883111561526357600080fd5b928401925b8284101561528157833582529284019290840190615268565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b818110156152ce57835163ffffffff16835292840192918401916001016152ac565b50909695505050505050565b600080602083850312156152ed57600080fd5b82356001600160401b0381111561530357600080fd5b6149c185828601614bec565b6001600160601b0381168114610c7457600080fd5b600080600060a0848603121561533957600080fd5b6153438585614f0f565b925060608401356153538161530f565b915060808401356001600160401b0381111561536e57600080fd5b8401601f8101861361537f57600080fd5b8035602061538f6150b183614fe8565b82815260069290921b830181019181810190898411156153ae57600080fd5b938201935b83851015615404576040858b0312156153cc5760008081fd5b6153d4614a6e565b85356153df81614b67565b8152858401356153ee8161530f565b81850152825260409490940193908201906153b3565b80955050505050509250925092565b600080600080600060a0868803121561542b57600080fd5b853561543681614b67565b9450602086013561544681614b67565b9350604086013561545681614b67565b9250606086013561546681614b67565b949793965091946080013592915050565b60208101610f038284614eb9565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016154fa576154fa6154d2565b5060010190565b60006020828403121561551357600080fd5b81516125c681614b67565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561557a57600080fd5b815180151581146125c657600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000604082840312156155e457600080fd5b6125c6838361500b565b6000808335601e1984360301811261560557600080fd5b8301803591506001600160401b0382111561561f57600080fd5b6020019150600581901b36038213156145eb57600080fd5b60006020828403121561564957600080fd5b81516125c6816149e6565b80820180821115610f0357610f036154d2565b6000808585111561567757600080fd5b8386111561568457600080fd5b5050820193919092039150565b600060c08201888352602060018060a01b03808a16828601526040898187015260c0606087015283895180865260e088019150848b01955060005b818110156156f6578651805160ff16845286015185168684015295850195918301916001016156cc565b505060808701989098525050505060a09091019190915250949350505050565b81810381811115610f0357610f036154d2565b6000815180845260005b8181101561574f57602081850181015186830182015201615733565b506000602082860101526020601f19601f83011685010191505092915050565b60018060a01b03841681528260208201526060604082015260006157966060830184615729565b95945050505050565b6000602082840312156157b157600080fd5b81516001600160c01b03811681146125c657600080fd5b828152604060208201526000614b5f6040830184615729565b6020815260006125c66020830184615729565b60006020828403121561580657600080fd5b5051919050565b6001600160a01b03841681526101608101615835602083018580358252602090810135910152565b61584f606083016040860180358252602090810135910152565b60406080850160a0840137604060c0850160e084013782516101208301526020830151610140830152614b5f565b634e487b7160e01b600052601260045260246000fd5b6000826158a2576158a261587d565b500690565b6001600160a01b0383168152604060208201819052600090614b5f90830184615729565b6001600160a01b03841681526060602082018190526000906158ef90830185615729565b828103604084015283516060825261590a6060830182615729565b905060208501516020830152604085015160408301528092505050949350505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0384168152604060208201819052600090615796908301848661592d565b60018060a01b03851681528360208201526060604082015260006118e860608301848661592d565b600082601f8301126159b457600080fd5b815160206159c46150b183614fe8565b82815260059290921b840181019181810190868411156159e357600080fd5b8286015b84811015615a075780516159fa8161530f565b83529183019183016159e7565b509695505050505050565b60008060408385031215615a2557600080fd5b82516001600160401b0380821115615a3c57600080fd5b615a48868387016159a3565b93506020850151915080821115615a5e57600080fd5b50615a6b858286016159a3565b9150509250929050565b83815260406020820152600061579660408301848661592d565b60006020808385031215615aa257600080fd5b82516001600160401b03811115615ab857600080fd5b8301601f81018513615ac957600080fd5b8051615ad76150b182614fe8565b81815260059190911b82018301908381019087831115615af657600080fd5b928401925b82841015615b1d578351615b0e816149e6565b82529284019290840190615afb565b979650505050505050565b60ff8181168382160190811115610f0357610f036154d2565b60006060820160ff8616835260206001600160601b03808716828601526040606081870152838751808652608088019150848901955060005b81811015615baa57865180516001600160a01b031684528601518516868401529585019591830191600101615b7a565b50909a9950505050505050505050565b600061ffff808316818103615bd157615bd16154d2565b6001019392505050565b600060208284031215615bed57600080fd5b81516001600160e01b0319811681146125c657600080fd5b600060208284031215615c1757600080fd5b81516125c68161530f565b6001600160601b03818116838216028082169190828114615c4557615c456154d2565b505092915050565b60006001600160601b0380841680615c6757615c6761587d565b9216919091049291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47526567436f6f72642e7570646174654f70657261746f7273466f7251756f7275a2646970667358221220788135bb8745ce48b89144ad463410191e0ac00f4355532f29105673fbace3fa64736f6c63430008150033",
}

// ContractZrRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractZrRegistryCoordinatorMetaData.ABI instead.
var ContractZrRegistryCoordinatorABI = ContractZrRegistryCoordinatorMetaData.ABI

// ContractZrRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractZrRegistryCoordinatorMetaData.Bin instead.
var ContractZrRegistryCoordinatorBin = ContractZrRegistryCoordinatorMetaData.Bin

// DeployContractZrRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractZrRegistryCoordinator to it.
func DeployContractZrRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _serviceManager common.Address, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address, _socketRegistry common.Address) (common.Address, *types.Transaction, *ContractZrRegistryCoordinator, error) {
	parsed, err := ContractZrRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractZrRegistryCoordinatorBin), backend, _serviceManager, _stakeRegistry, _blsApkRegistry, _indexRegistry, _socketRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractZrRegistryCoordinator{ContractZrRegistryCoordinatorCaller: ContractZrRegistryCoordinatorCaller{contract: contract}, ContractZrRegistryCoordinatorTransactor: ContractZrRegistryCoordinatorTransactor{contract: contract}, ContractZrRegistryCoordinatorFilterer: ContractZrRegistryCoordinatorFilterer{contract: contract}}, nil
}

// ContractZrRegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractZrRegistryCoordinator struct {
	ContractZrRegistryCoordinatorCaller     // Read-only binding to the contract
	ContractZrRegistryCoordinatorTransactor // Write-only binding to the contract
	ContractZrRegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractZrRegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractZrRegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrRegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractZrRegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrRegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractZrRegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZrRegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractZrRegistryCoordinatorSession struct {
	Contract     *ContractZrRegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractZrRegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractZrRegistryCoordinatorCallerSession struct {
	Contract *ContractZrRegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractZrRegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractZrRegistryCoordinatorTransactorSession struct {
	Contract     *ContractZrRegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractZrRegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractZrRegistryCoordinatorRaw struct {
	Contract *ContractZrRegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractZrRegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractZrRegistryCoordinatorCallerRaw struct {
	Contract *ContractZrRegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractZrRegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractZrRegistryCoordinatorTransactorRaw struct {
	Contract *ContractZrRegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractZrRegistryCoordinator creates a new instance of ContractZrRegistryCoordinator, bound to a specific deployed contract.
func NewContractZrRegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractZrRegistryCoordinator, error) {
	contract, err := bindContractZrRegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinator{ContractZrRegistryCoordinatorCaller: ContractZrRegistryCoordinatorCaller{contract: contract}, ContractZrRegistryCoordinatorTransactor: ContractZrRegistryCoordinatorTransactor{contract: contract}, ContractZrRegistryCoordinatorFilterer: ContractZrRegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractZrRegistryCoordinatorCaller creates a new read-only instance of ContractZrRegistryCoordinator, bound to a specific deployed contract.
func NewContractZrRegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractZrRegistryCoordinatorCaller, error) {
	contract, err := bindContractZrRegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractZrRegistryCoordinatorTransactor creates a new write-only instance of ContractZrRegistryCoordinator, bound to a specific deployed contract.
func NewContractZrRegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractZrRegistryCoordinatorTransactor, error) {
	contract, err := bindContractZrRegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractZrRegistryCoordinatorFilterer creates a new log filterer instance of ContractZrRegistryCoordinator, bound to a specific deployed contract.
func NewContractZrRegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractZrRegistryCoordinatorFilterer, error) {
	contract, err := bindContractZrRegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractZrRegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractZrRegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractZrRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZrRegistryCoordinator.Contract.ContractZrRegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.ContractZrRegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.ContractZrRegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZrRegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractZrRegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractZrRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractZrRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractZrRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.BlsApkRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.BlsApkRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractZrRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractZrRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.ChurnApprover(&_ContractZrRegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.ChurnApprover(&_ContractZrRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) EjectionCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "ejectionCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) EjectionCooldown() (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.EjectionCooldown(&_ContractZrRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) EjectionCooldown() (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.EjectionCooldown(&_ContractZrRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.Ejector(&_ContractZrRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.Ejector(&_ContractZrRegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractZrRegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractZrRegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorInfo)).(*IRegistryCoordinatorOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperator(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperator(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorFromId(&_ContractZrRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorFromId(&_ContractZrRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorId(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorId(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorSetParam)).(*IRegistryCoordinatorOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractZrRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractZrRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorStatus(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractZrRegistryCoordinator.Contract.GetOperatorStatus(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractZrRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractZrRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractZrRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractZrRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractZrRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractZrRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(IRegistryCoordinatorQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorQuorumBitmapUpdate)).(*IRegistryCoordinatorQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractZrRegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractZrRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractZrRegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.IndexRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.IndexRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractZrRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractZrRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "lastEjectionTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.NumRegistries(&_ContractZrRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.NumRegistries(&_ContractZrRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.Owner(&_ContractZrRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.Owner(&_ContractZrRegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractZrRegistryCoordinator.Contract.Paused(&_ContractZrRegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractZrRegistryCoordinator.Contract.Paused(&_ContractZrRegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.Paused0(&_ContractZrRegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.Paused0(&_ContractZrRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.PauserRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.PauserRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractZrRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractZrRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractZrRegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractZrRegistryCoordinator.Contract.QuorumCount(&_ContractZrRegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractZrRegistryCoordinator.Contract.QuorumCount(&_ContractZrRegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractZrRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.Registries(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.Registries(&_ContractZrRegistryCoordinator.CallOpts, arg0)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) ServiceManager() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.ServiceManager(&_ContractZrRegistryCoordinator.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) ServiceManager() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.ServiceManager(&_ContractZrRegistryCoordinator.CallOpts)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) SocketRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "socketRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) SocketRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.SocketRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) SocketRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.SocketRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrRegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.StakeRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractZrRegistryCoordinator.Contract.StakeRegistry(&_ContractZrRegistryCoordinator.CallOpts)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) CreateQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "createQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) CreateQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.CreateQuorum(&_ContractZrRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) CreateQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.CreateQuorum(&_ContractZrRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "deregisterOperator", quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) DeregisterOperator(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.DeregisterOperator(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) DeregisterOperator(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.DeregisterOperator(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.EjectOperator(&_ContractZrRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.EjectOperator(&_ContractZrRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.Initialize(&_ContractZrRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.Initialize(&_ContractZrRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.Pause(&_ContractZrRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.Pause(&_ContractZrRegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.PauseAll(&_ContractZrRegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.PauseAll(&_ContractZrRegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xae42cdb2.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, string validatorAddr, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, socket string, validatorAddr string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "registerOperator", quorumNumbers, socket, validatorAddr, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xae42cdb2.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, string validatorAddr, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) RegisterOperator(quorumNumbers []byte, socket string, validatorAddr string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.RegisterOperator(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumbers, socket, validatorAddr, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xae42cdb2.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, string validatorAddr, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) RegisterOperator(quorumNumbers []byte, socket string, validatorAddr string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.RegisterOperator(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumbers, socket, validatorAddr, params, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x2b21ea6d.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, string validatorAddr, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, socket string, validatorAddr string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "registerOperatorWithChurn", quorumNumbers, socket, validatorAddr, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x2b21ea6d.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, string validatorAddr, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, validatorAddr string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumbers, socket, validatorAddr, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x2b21ea6d.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, string validatorAddr, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, validatorAddr string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumbers, socket, validatorAddr, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.RenounceOwnership(&_ContractZrRegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.RenounceOwnership(&_ContractZrRegistryCoordinator.TransactOpts)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetChurnApprover(&_ContractZrRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetChurnApprover(&_ContractZrRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "setEjectionCooldown", _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractZrRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractZrRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetEjector(&_ContractZrRegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetEjector(&_ContractZrRegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractZrRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetPauserRegistry(&_ContractZrRegistryCoordinator.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.SetPauserRegistry(&_ContractZrRegistryCoordinator.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.TransferOwnership(&_ContractZrRegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.TransferOwnership(&_ContractZrRegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.Unpause(&_ContractZrRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.Unpause(&_ContractZrRegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.UpdateOperators(&_ContractZrRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.UpdateOperators(&_ContractZrRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractZrRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractZrRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.UpdateSocket(&_ContractZrRegistryCoordinator.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractZrRegistryCoordinator.Contract.UpdateSocket(&_ContractZrRegistryCoordinator.TransactOpts, socket)
}

// ContractZrRegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractZrRegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorChurnApproverUpdated)
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
		it.Event = new(ContractZrRegistryCoordinatorChurnApproverUpdated)
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
func (it *ContractZrRegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractZrRegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractZrRegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorChurnApproverUpdated)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
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

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractZrRegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractZrRegistryCoordinatorChurnApproverUpdated)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractZrRegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorEjectorUpdated)
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
		it.Event = new(ContractZrRegistryCoordinatorEjectorUpdated)
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
func (it *ContractZrRegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractZrRegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorEjectorUpdatedIterator{contract: _ContractZrRegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorEjectorUpdated)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
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

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractZrRegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractZrRegistryCoordinatorEjectorUpdated)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorInitializedIterator struct {
	Event *ContractZrRegistryCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorInitialized)
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
		it.Event = new(ContractZrRegistryCoordinatorInitialized)
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
func (it *ContractZrRegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorInitialized represents a Initialized event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractZrRegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorInitializedIterator{contract: _ContractZrRegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorInitialized)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractZrRegistryCoordinatorInitialized, error) {
	event := new(ContractZrRegistryCoordinatorInitialized)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractZrRegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorOperatorDeregistered)
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
		it.Event = new(ContractZrRegistryCoordinatorOperatorDeregistered)
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
func (it *ContractZrRegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractZrRegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractZrRegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorOperatorDeregistered)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractZrRegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractZrRegistryCoordinatorOperatorDeregistered)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractZrRegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorOperatorRegistered)
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
		it.Event = new(ContractZrRegistryCoordinatorOperatorRegistered)
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
func (it *ContractZrRegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractZrRegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorOperatorRegisteredIterator{contract: _ContractZrRegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorOperatorRegistered)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractZrRegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractZrRegistryCoordinatorOperatorRegistered)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractZrRegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorOperatorSetParamsUpdated)
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
		it.Event = new(ContractZrRegistryCoordinatorOperatorSetParamsUpdated)
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
func (it *ContractZrRegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams IRegistryCoordinatorOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractZrRegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractZrRegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
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

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractZrRegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractZrRegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorSocketUpdateIterator struct {
	Event *ContractZrRegistryCoordinatorOperatorSocketUpdate // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorOperatorSocketUpdate)
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
		it.Event = new(ContractZrRegistryCoordinatorOperatorSocketUpdate)
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
func (it *ContractZrRegistryCoordinatorOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractZrRegistryCoordinatorOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorOperatorSocketUpdateIterator{contract: _ContractZrRegistryCoordinator.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorOperatorSocketUpdate)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
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

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseOperatorSocketUpdate(log types.Log) (*ContractZrRegistryCoordinatorOperatorSocketUpdate, error) {
	event := new(ContractZrRegistryCoordinatorOperatorSocketUpdate)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOwnershipTransferredIterator struct {
	Event *ContractZrRegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorOwnershipTransferred)
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
		it.Event = new(ContractZrRegistryCoordinatorOwnershipTransferred)
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
func (it *ContractZrRegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractZrRegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorOwnershipTransferredIterator{contract: _ContractZrRegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorOwnershipTransferred)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractZrRegistryCoordinatorOwnershipTransferred, error) {
	event := new(ContractZrRegistryCoordinatorOwnershipTransferred)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorPausedIterator struct {
	Event *ContractZrRegistryCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorPaused)
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
		it.Event = new(ContractZrRegistryCoordinatorPaused)
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
func (it *ContractZrRegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorPaused represents a Paused event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractZrRegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorPausedIterator{contract: _ContractZrRegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorPaused)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParsePaused(log types.Log) (*ContractZrRegistryCoordinatorPaused, error) {
	event := new(ContractZrRegistryCoordinatorPaused)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorPauserRegistrySetIterator struct {
	Event *ContractZrRegistryCoordinatorPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorPauserRegistrySet)
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
		it.Event = new(ContractZrRegistryCoordinatorPauserRegistrySet)
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
func (it *ContractZrRegistryCoordinatorPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorPauserRegistrySet represents a PauserRegistrySet event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractZrRegistryCoordinatorPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorPauserRegistrySetIterator{contract: _ContractZrRegistryCoordinator.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorPauserRegistrySet)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParsePauserRegistrySet(log types.Log) (*ContractZrRegistryCoordinatorPauserRegistrySet, error) {
	event := new(ContractZrRegistryCoordinatorPauserRegistrySet)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractZrRegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorQuorumBlockNumberUpdated)
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
		it.Event = new(ContractZrRegistryCoordinatorQuorumBlockNumberUpdated)
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
func (it *ContractZrRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractZrRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractZrRegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
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

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractZrRegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractZrRegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZrRegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorUnpausedIterator struct {
	Event *ContractZrRegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractZrRegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZrRegistryCoordinatorUnpaused)
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
		it.Event = new(ContractZrRegistryCoordinatorUnpaused)
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
func (it *ContractZrRegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZrRegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZrRegistryCoordinatorUnpaused represents a Unpaused event raised by the ContractZrRegistryCoordinator contract.
type ContractZrRegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractZrRegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZrRegistryCoordinatorUnpausedIterator{contract: _ContractZrRegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractZrRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZrRegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZrRegistryCoordinatorUnpaused)
				if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractZrRegistryCoordinator *ContractZrRegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractZrRegistryCoordinatorUnpaused, error) {
	event := new(ContractZrRegistryCoordinatorUnpaused)
	if err := _ContractZrRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
