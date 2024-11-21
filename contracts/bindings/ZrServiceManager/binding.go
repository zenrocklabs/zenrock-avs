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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIZrRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"oprAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectValidators\",\"inputs\":[{\"name\":\"inactiveValidatorSet\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllOperatorsAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structZrServiceManagerLib.Validator[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEigenStake\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakedBalanceForValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Validator\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"contractIZRTaskManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIZRTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAssigned\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDeregistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorsEjected\",\"inputs\":[{\"name\":\"validatorHashes\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IndexOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTaskManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownValidatorHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005a0638038062005a06833981016040819052620000359162000150565b60016000556001600160a01b0380851660805280841660a05280831660c052811660e052838383836200006762000075565b5050505050505050620001b8565b603354610100900460ff1615620000e25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60335460ff908116101562000135576033805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014d57600080fd5b50565b600080600080608085870312156200016757600080fd5b8451620001748162000137565b6020860151909450620001878162000137565b60408601519093506200019a8162000137565b6060860151909250620001ad8162000137565b939692955090935050565b60805160a05160c05160e05161576c6200029a60003960008181610a7501528181610bd001528181610c670152818161277b0152818161294d01528181612ad00152612b6f0152600081816108a00152818161092f015281816109af01528181611e4501528181611f980152818161207b01528181612147015281816127a50152818161288801528181612a2b01528181612d2601528181612fb20152613fed01526000818161317401528181613230015261331c015260008181610368015281816120cf015281816121ac0152818161222b0152612db9015261576c6000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80639926ee7d11610125578063da4dc49c116100ad578063f891899d1161007c578063f891899d14610510578063fabc1cbc14610523578063fb558ed514610536578063fc299dee1461053e578063fce36c7d1461055157600080fd5b8063da4dc49c146104da578063e481af9d146104ed578063f2fde38b146104f5578063f5c9899d1461050857600080fd5b8063c0c53b8b116100f4578063c0c53b8b14610448578063c63e3c501461045b578063cefdc1d41461046e578063d5f20ff61461048f578063d604902f146104af57600080fd5b80639926ee7d146103f1578063a364f4da14610404578063a50a640e14610417578063a98fb3551461043557600080fd5b80635ac86ab7116101a8578063715018a611610177578063715018a6146103a05780637b654c5d146103a8578063886f1195146103c5578063889e2c6e146103d85780638da5cb5b146103e057600080fd5b80635ac86ab7146103025780635c155662146103355780635c975abb146103555780636b3aa72e1461036657600080fd5b80633563b0d1116101ef5780633563b0d1146102925780634a91a2f8146102b25780634d2b57fe146102c75780634f739f74146102da578063595c6a67146102fa57600080fd5b806310d67a2f14610221578063136439dd1461023657806331b36bd91461024957806333cfb7b714610272575b600080fd5b61023461022f366004614204565b610564565b005b610234610244366004614221565b610620565b61025c6102573660046142cb565b61075f565b60405161026991906143b9565b60405180910390f35b610285610280366004614204565b61087b565b60405161026991906143cc565b6102a56102a03660046144aa565b610d4a565b60405161026991906145b1565b6102ba6111e0565b604051610269919061468b565b6102856102d5366004614748565b6112df565b6102ed6102e8366004614823565b6113f4565b60405161026991906148ed565b610234611b1c565b6103256103103660046149b7565b600254600160ff9092169190911b9081161490565b6040519015158152602001610269565b6103486103433660046149d4565b611be3565b6040516102699190614a1b565b600254604051908152602001610269565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610269565b610234611dab565b6103b0611dbf565b60405163ffffffff9091168152602001610269565b600154610388906001600160a01b031681565b610285611e3f565b6066546001600160a01b0316610388565b6102346103ff366004614a53565b612070565b610234610412366004614204565b61213c565b6000805160206156f7833981519152546001600160a01b0316610388565b610234610443366004614afe565b61220c565b610234610456366004614b3a565b612260565b610234610469366004614b7a565b6123a8565b61048161047c366004614bbb565b6123f9565b604051610269929190614bf2565b6104a261049d366004614221565b61258b565b6040516102699190614c0b565b6104c26104bd366004614c1e565b6125b6565b6040516001600160601b039091168152602001610269565b6104c26104e8366004614c53565b612759565b610285612882565b610234610503366004614204565b612c4e565b6103b0612cc4565b61023461051e366004614c8c565b612d1b565b610234610531366004614221565b612e54565b610234612fb0565b609854610388906001600160a01b031681565b61023461055f366004614b7a565b613030565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105db9190614d0c565b6001600160a01b0316336001600160a01b0316146106145760405162461bcd60e51b815260040161060b90614d29565b60405180910390fd5b61061d81613353565b50565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610668573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068c9190614d73565b6106a85760405162461bcd60e51b815260040161060b90614d95565b600254818116146107215760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161060b565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b606081516001600160401b0381111561077a5761077a61423a565b6040519080825280602002602001820160405280156107a3578160200160208202803683370190505b50905060005b825181101561087457836001600160a01b03166313542a4e8483815181106107d3576107d3614ddd565b60200260200101516040518263ffffffff1660e01b815260040161080691906001600160a01b0391909116815260200190565b602060405180830381865afa158015610823573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108479190614df3565b82828151811061085957610859614ddd565b602090810291909101015261086d81614e22565b90506107a9565b5092915050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156108e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090b9190614df3565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610976573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099a9190614e3b565b90506001600160c01b0381161580610a3457507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2f9190614e64565b60ff16155b15610a5057505060408051600081526020810190915292915050565b6000610a64826001600160c01b031661344a565b90506000805b8251811015610b3a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610ab457610ab4614ddd565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610af8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b1c9190614df3565b610b269083614e81565b915080610b3281614e22565b915050610a6a565b506000816001600160401b03811115610b5557610b5561423a565b604051908082528060200260200182016040528015610b7e578160200160208202803683370190505b5090506000805b8451811015610d3d576000858281518110610ba257610ba2614ddd565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610c17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3b9190614df3565b905060005b81811015610d27576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610cb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd99190614ea9565b60000151868681518110610cef57610cef614ddd565b6001600160a01b039092166020928302919091019091015284610d1181614e22565b9550508080610d1f90614e22565b915050610c40565b5050508080610d3590614e22565b915050610b85565b5090979650505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db09190614d0c565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610df2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e169190614d0c565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7c9190614d0c565b9050600086516001600160401b03811115610e9957610e9961423a565b604051908082528060200260200182016040528015610ecc57816020015b6060815260200190600190039081610eb75790505b50905060005b87518110156111d4576000888281518110610eef57610eef614ddd565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610f50573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f789190810190614eea565b905080516001600160401b03811115610f9357610f9361423a565b604051908082528060200260200182016040528015610fde57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610fb15790505b50848481518110610ff157610ff1614ddd565b602002602001018190525060005b81518110156111be576040518060600160405280876001600160a01b03166347b314e885858151811061103457611034614ddd565b60200260200101516040518263ffffffff1660e01b815260040161105a91815260200190565b602060405180830381865afa158015611077573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109b9190614d0c565b6001600160a01b031681526020018383815181106110bb576110bb614ddd565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106110e9576110e9614ddd565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611145573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111699190614f7a565b6001600160601b031681525085858151811061118757611187614ddd565b602002602001015182815181106111a0576111a0614ddd565b602002602001018190525080806111b690614e22565b915050610fff565b50505080806111cc90614e22565b915050610ed2565b50979650505050505050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0354606090600080516020615717833981519152906000816001600160401b0381111561122f5761122f61423a565b60405190808252806020026020018201604052801561127c57816020015b6040805160608082018352808252600060208301529181019190915281526020019060019003908161124d5790505b50905060005b828110156112d7576112b28460030182815481106112a2576112a2614ddd565b906000526020600020015461350c565b8282815181106112c4576112c4614ddd565b6020908102919091010152600101611282565b509392505050565b606081516001600160401b038111156112fa576112fa61423a565b604051908082528060200260200182016040528015611323578160200160208202803683370190505b50905060005b825181101561087457836001600160a01b031663296bb06484838151811061135357611353614ddd565b60200260200101516040518263ffffffff1660e01b815260040161137991815260200190565b602060405180830381865afa158015611396573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ba9190614d0c565b8282815181106113cc576113cc614ddd565b6001600160a01b03909216602092830291909101909101526113ed81614e22565b9050611329565b61141f6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561145f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114839190614d0c565b90506114b06040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906114e0908b9089908990600401614f97565b600060405180830381865afa1580156114fd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115259190810190614fde565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611557908b908b908b90600401615095565b600060405180830381865afa158015611574573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261159c9190810190614fde565b6040820152856001600160401b038111156115b9576115b961423a565b6040519080825280602002602001820160405280156115ec57816020015b60608152602001906001900390816115d75790505b50606082015260005b60ff8116871115611a2d576000856001600160401b0381111561161a5761161a61423a565b604051908082528060200260200182016040528015611643578160200160208202803683370190505b5083606001518360ff168151811061165d5761165d614ddd565b602002602001018190525060005b8681101561192d5760008c6001600160a01b03166304ec63518a8a8581811061169657611696614ddd565b905060200201358e886000015186815181106116b4576116b4614ddd565b60200260200101516040518463ffffffff1660e01b81526004016116f19392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561170e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117329190614e3b565b9050806001600160c01b03166000036117d95760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40161060b565b8a8a8560ff168181106117ee576117ee614ddd565b60016001600160c01b038516919093013560f81c1c8216909103905061191a57856001600160a01b031663dd9846b98a8a8581811061182f5761182f614ddd565b905060200201358d8d8860ff1681811061184b5761184b614ddd565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156118a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118c591906150be565b85606001518560ff16815181106118de576118de614ddd565b602002602001015184815181106118f7576118f7614ddd565b63ffffffff909216602092830291909101909101528261191681614e22565b9350505b508061192581614e22565b91505061166b565b506000816001600160401b038111156119485761194861423a565b604051908082528060200260200182016040528015611971578160200160208202803683370190505b50905060005b828110156119f25784606001518460ff168151811061199857611998614ddd565b602002602001015181815181106119b1576119b1614ddd565b60200260200101518282815181106119cb576119cb614ddd565b63ffffffff90921660209283029190910190910152806119ea81614e22565b915050611977565b508084606001518460ff1681518110611a0d57611a0d614ddd565b602002602001018190525050508080611a25906150db565b9150506115f5565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a929190614d0c565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611ac5908b908b908e906004016150fa565b600060405180830381865afa158015611ae2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b0a9190810190614fde565b60208301525098975050505050505050565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611b64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b889190614d73565b611ba45760405162461bcd60e51b815260040161060b90614d95565b600019600281905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611c15929190615124565b600060405180830381865afa158015611c32573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c5a9190810190614fde565b9050600084516001600160401b03811115611c7757611c7761423a565b604051908082528060200260200182016040528015611ca0578160200160208202803683370190505b50905060005b8551811015611da157866001600160a01b03166304ec6351878381518110611cd057611cd0614ddd565b602002602001015187868581518110611ceb57611ceb614ddd565b60200260200101516040518463ffffffff1660e01b8152600401611d289392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d699190614e3b565b6001600160c01b0316828281518110611d8457611d84614ddd565b602090810291909101015280611d9981614e22565b915050611ca6565b5095945050505050565b611db3613655565b611dbd60006136af565b565b6000805160206156f783398151915254604080516372d18e8d60e01b815290516000926001600160a01b0316916372d18e8d9160048281019260209291908290030181865afa158015611e16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e3a91906150be565b905090565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ea1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ec59190614d0c565b604051638902624560e01b81526000600482015263ffffffff431660248201526001600160a01b039190911690638902624590604401600060405180830381865afa158015611f18573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f409190810190614eea565b80519091506000816001600160401b03811115611f5f57611f5f61423a565b604051908082528060200260200182016040528015611f88578160200160208202803683370190505b50905060005b828110156112d7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663296bb064858381518110611fd757611fd7614ddd565b60200260200101516040518263ffffffff1660e01b8152600401611ffd91815260200190565b602060405180830381865afa15801561201a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203e9190614d0c565b82828151811061205057612050614ddd565b6001600160a01b0390921660209283029190910190910152600101611f8e565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146120b85760405162461bcd60e51b815260040161060b90615143565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061210690859085906004016151bb565b600060405180830381600087803b15801561212057600080fd5b505af1158015612134573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146121845760405162461bcd60e51b815260040161060b90615143565b61218d81613701565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b1580156121f157600080fd5b505af1158015612205573d6000803e3d6000fd5b5050505050565b612214613655565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906121d7908490600401615206565b603354610100900460ff16158080156122805750603354600160ff909116105b8061229a5750303b15801561229a575060335460ff166001145b6122fd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161060b565b6033805460ff191660011790558015612320576033805461ff0019166101001790555b61232b84600061395e565b61233483613a44565b6000805160206156f783398151915280546001600160a01b0319166001600160a01b03841617905580156123a2576033805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b600080516020615717833981519152600401546001600160a01b031633146123e35760405163f06b6b4360e01b815260040160405180910390fd5b6123f56123f08284615219565b613a74565b5050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061243457612434614ddd565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906124709088908690600401615124565b600060405180830381865afa15801561248d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526124b59190810190614fde565b6000815181106124c7576124c7614ddd565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612533573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125579190614e3b565b6001600160c01b03169050600061256d8261344a565b90508161257b8a838a610d4a565b9550955050505050935093915050565b604080516060808201835280825260006020830152918101919091526125b08261350c565b92915050565b60008083836040516020016125cc92919061528c565b60405160208183030381529060405280519060200120905060006125fb60008051602061571783398151915290565b60008381526020919091526040908190208151606081019092528054829082906126249061529c565b80601f01602080910402602001604051908101604052809291908181526020018280546126509061529c565b801561269d5780601f106126725761010080835404028352916020019161269d565b820191906000526020600020905b81548152906001019060200180831161268057829003601f168201915b50505050508152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561270957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116126eb575b50505050508152505090506000805b826040015151811015611da15761274d8360400151828151811061273e5761273e614ddd565b60200260200101516000612759565b90910190600101612718565b604051631619718360e21b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000821691635401ed27917f000000000000000000000000000000000000000000000000000000000000000090911690635865c60c906024016040805180830381865afa1580156127ed573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061281191906152d6565b5160405160e083901b6001600160e01b0319168152600481019190915260ff85166024820152604401602060405180830381865afa158015612857573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061287b9190614f7a565b9392505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129089190614e64565b60ff1690508060000361292957505060408051600081526020810190915290565b6000805b828110156129de57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa15801561299c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129c09190614df3565b6129ca9083614e81565b9150806129d681614e22565b91505061292d565b506000816001600160401b038111156129f9576129f961423a565b604051908082528060200260200182016040528015612a22578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aab9190614e64565b60ff16811015612c4457604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015612b1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b439190614df3565b905060005b81811015612c2f576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015612bbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612be19190614ea9565b60000151858581518110612bf757612bf7614ddd565b6001600160a01b039092166020928302919091019091015283612c1981614e22565b9450508080612c2790614e22565b915050612b48565b50508080612c3c90614e22565b915050612a29565b5090949350505050565b612c56613655565b6001600160a01b038116612cbb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161060b565b61061d816136af565b6000805160206156f7833981519152546040805163f5c9899d60e01b815290516000926001600160a01b03169163f5c9899d9160048281019260209291908290030181865afa158015611e16573d6000803e3d6000fd5b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612d635760405162461bcd60e51b815260040161060b90615143565b612da283838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613b7992505050565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90612df09087908590600401615306565b600060405180830381600087803b158015612e0a57600080fd5b505af1158015612e1e573d6000803e3d6000fd5b505050506123a2848484604051602001612e3992919061528c565b60405160208183030381529060405280519060200120613bc0565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ea7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ecb9190614d0c565b6001600160a01b0316336001600160a01b031614612efb5760405162461bcd60e51b815260040161060b90614d29565b600254198119600254191614612f795760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161060b565b600281905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610754565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031662cf2ab5612fe6611e3f565b6040518263ffffffff1660e01b815260040161300291906143cc565b600060405180830381600087803b15801561301c57600080fd5b505af11580156123a2573d6000803e3d6000fd5b60005b818110156133045782828281811061304d5761304d614ddd565b905060200281019061305f9190615392565b613070906040810190602001614204565b6001600160a01b03166323b872dd333086868681811061309257613092614ddd565b90506020028101906130a49190615392565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156130fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061311f9190614d73565b50600083838381811061313457613134614ddd565b90506020028101906131469190615392565b613157906040810190602001614204565b604051636eb1769f60e11b81523060048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166024830152919091169063dd62ed3e90604401602060405180830381865afa1580156131c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131e99190614df3565b90508383838181106131fd576131fd614ddd565b905060200281019061320f9190615392565b613220906040810190602001614204565b6001600160a01b031663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008387878781811061326257613262614ddd565b90506020028101906132749190615392565b604001356132829190614e81565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156132cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132f19190614d73565b5050806132fd90614e22565b9050613033565b5060405163fce36c7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063fce36c7d906121069085908590600401615418565b6001600160a01b0381166133e15760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161060b565b600154604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b606060008061345884613d03565b61ffff166001600160401b038111156134735761347361423a565b6040519080825280601f01601f19166020018201604052801561349d576020820181803683370190505b5090506000805b8251821080156134b5575061010081105b15612c44576001811b9350858416156134fc578060f81b8383815181106134de576134de614ddd565b60200101906001600160f81b031916908160001a9053508160010191505b61350581614e22565b90506134a4565b6040805160608082018352808252600060208301529181019190915260008051602061571783398151915260008381526020919091526040908190208151606081019092528054829082906135609061529c565b80601f016020809104026020016040519081016040528092919081815260200182805461358c9061529c565b80156135d95780601f106135ae576101008083540402835291602001916135d9565b820191906000526020600020905b8154815290600101906020018083116135bc57829003601f168201915b50505050508152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561364557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613627575b5050505050815250509050919050565b6066546001600160a01b03163314611dbd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161060b565b606680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0381166137285760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f016020526040902054600080516020615717833981519152908061378e57604051631e7b7a8160e01b815260040160405180910390fd5b6001600160a01b038316600090815260028084016020908152604080842054858552918690529092209081015482106137da57604051634e23d03560e01b815260040160405180910390fd5b846001600160a01b03168160020183815481106137f9576137f9614ddd565b6000918252602090912001546001600160a01b03161461382c57604051634265bbb960e11b815260040160405180910390fd5b600281015460009061384090600190615526565b90508083146138ce57600082600201828154811061386057613860614ddd565b6000918252602090912001546002840180546001600160a01b03909216925082918690811061389157613891614ddd565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394851617905592909116815260028701909152604090208390555b816002018054806138e1576138e1615539565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038816808352600188018252604080842084905560028901909252818320839055905186927f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a3505050505050565b6001546001600160a01b031615801561397f57506001600160a01b03821615155b613a015760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161060b565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26123f582613353565b603354610100900460ff16613a6b5760405162461bcd60e51b815260040161060b9061554f565b61061d81613d2e565b80516000816001600160401b03811115613a9057613a9061423a565b604051908082528060200260200182016040528015613ab9578160200160208202803683370190505b50905060005b82811015613b3c576000848281518110613adb57613adb614ddd565b6020026020010151604051602001613af3919061559a565b60405160208183030381529060405280519060200120905080838381518110613b1e57613b1e614ddd565b602002602001018181525050613b3381613d55565b50600101613abf565b507fb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a81604051613b6c91906143b9565b60405180910390a1505050565b600081604051602001613b8c919061559a565b604051602081830303815290604052805190602001209050613bad81613e35565b6123f557613bbb8282613e4a565b505050565b6001600160a01b038216613be75760405163d92e233d60e01b815260040160405180910390fd5b80613c0557604051631e7b7a8160e01b815260040160405180910390fd5b613c0e81613e35565b613c2b57604051634d709ad360e11b815260040160405180910390fd5b6001600160a01b03821660008181527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f016020908152604080832085905584835260008051602061571783398151915280835281842060020180548686527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f02855283862081905582855260018101825590855292842090920180546001600160a01b03191685179055519092849290917ffba1d8bd64cc55f778b2832392428433c24bd739c106b234ab0a0bee8a1a7f039190a3505050565b6000805b82156125b057613d18600184615526565b9092169180613d26816155ac565b915050613d07565b603354610100900460ff16612cbb5760405162461bcd60e51b815260040161060b9061554f565b600081815260008051602061571783398151915260208190526040909120600181015415613bbb57600281015460005b81811015613e03576000836002018281548110613da457613da4614ddd565b6000918252602090912001546001600160a01b03169050613dc481613f62565b60405186906001600160a01b038316907f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e490600090a350600101613d85565b50600084815260208490526040812090613e1d828261411d565b60018201600090556002820160006121349190614157565b6000613e408261350c565b5151151592915050565b60408051606080820183528082526000602083015291810191909152613e6f82613e35565b15613e8d5760405163132e7efb60e31b815260040160405180910390fd5b60408051606081018252848152602080820185905282516000808252818301855283850191909152858152600080516020615717833981519152918290529290922081518291908190613ee09082615613565b50602082810151600183015560408301518051613f039260028501920190614175565b505050600382018054600181018255600091825260209091200184905560405184907f1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea390613f52908890615206565b60405180910390a2949350505050565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f01602052604090205480156123f55760408051600180825281830190925260009160208083019080368337019050509050600081600081518110613fd757613fd7614ddd565b602002602001019060ff16908160ff16815250507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e3b17db8461402484614078565b6040518363ffffffff1660e01b81526004016140419291906156d2565b600060405180830381600087803b15801561405b57600080fd5b505af115801561406f573d6000803e3d6000fd5b50505050505050565b80516060906000816001600160401b038111156140975761409761423a565b6040519080825280601f01601f1916602001820160405280156140c1576020820181803683370190505b50905060005b828110156112d7578481815181106140e1576140e1614ddd565b602002602001015160f81b8282815181106140fe576140fe614ddd565b60200101906001600160f81b031916908160001a9053506001016140c7565b5080546141299061529c565b6000825580601f10614139575050565b601f01602090049060005260206000209081019061061d91906141da565b508054600082559060005260206000209081019061061d91906141da565b8280548282559060005260206000209081019282156141ca579160200282015b828111156141ca57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614195565b506141d69291506141da565b5090565b5b808211156141d657600081556001016141db565b6001600160a01b038116811461061d57600080fd5b60006020828403121561421657600080fd5b813561287b816141ef565b60006020828403121561423357600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156142725761427261423a565b60405290565b604051601f8201601f191681016001600160401b03811182821017156142a0576142a061423a565b604052919050565b60006001600160401b038211156142c1576142c161423a565b5060051b60200190565b600080604083850312156142de57600080fd5b82356142e9816141ef565b91506020838101356001600160401b0381111561430557600080fd5b8401601f8101861361431657600080fd5b8035614329614324826142a8565b614278565b81815260059190911b8201830190838101908883111561434857600080fd5b928401925b8284101561436f578335614360816141ef565b8252928401929084019061434d565b80955050505050509250929050565b600081518084526020808501945080840160005b838110156143ae57815187529582019590820190600101614392565b509495945050505050565b60208152600061287b602083018461437e565b6020808252825182820181905260009190848201906040850190845b8181101561440d5783516001600160a01b0316835292840192918401916001016143e8565b50909695505050505050565b600082601f83011261442a57600080fd5b81356001600160401b038111156144435761444361423a565b614456601f8201601f1916602001614278565b81815284602083860101111561446b57600080fd5b816020850160208301376000918101602001919091529392505050565b63ffffffff8116811461061d57600080fd5b80356144a581614488565b919050565b6000806000606084860312156144bf57600080fd5b83356144ca816141ef565b925060208401356001600160401b038111156144e557600080fd5b6144f186828701614419565b925050604084013561450281614488565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156145a3578385038a52825180518087529087019087870190845b8181101561458e57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b0316908401529289019260609092019160010161454a565b50509a87019a9550509185019160010161452c565b509298975050505050505050565b60208152600061287b602083018461450d565b60005b838110156145df5781810151838201526020016145c7565b50506000910152565b600081518084526146008160208601602086016145c4565b601f01601f19169290920160200192915050565b600081516060845261462960608501826145e8565b9050602080840151818601526040840151858303604087015282815180855283850191508383019450600092505b808310156146805784516001600160a01b03168252938301936001929092019190830190614657565b509695505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156146e057603f198886030184526146ce858351614614565b945092850192908501906001016146b2565b5092979650505050505050565b600082601f8301126146fe57600080fd5b8135602061470e614324836142a8565b82815260059290921b8401810191818101908684111561472d57600080fd5b8286015b848110156146805780358352918301918301614731565b6000806040838503121561475b57600080fd5b8235614766816141ef565b915060208301356001600160401b0381111561478157600080fd5b61478d858286016146ed565b9150509250929050565b60008083601f8401126147a957600080fd5b5081356001600160401b038111156147c057600080fd5b6020830191508360208285010111156147d857600080fd5b9250929050565b60008083601f8401126147f157600080fd5b5081356001600160401b0381111561480857600080fd5b6020830191508360208260051b85010111156147d857600080fd5b6000806000806000806080878903121561483c57600080fd5b8635614847816141ef565b9550602087013561485781614488565b945060408701356001600160401b038082111561487357600080fd5b61487f8a838b01614797565b9096509450606089013591508082111561489857600080fd5b506148a589828a016147df565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b838110156143ae57815163ffffffff16875295820195908201906001016148cb565b60006020808352835160808285015261490960a08501826148b7565b905081850151601f198086840301604087015261492683836148b7565b9250604087015191508086840301606087015261494383836148b7565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561499a57848783030184526149888287516148b7565b9588019593880193915060010161496e565b509998505050505050505050565b60ff8116811461061d57600080fd5b6000602082840312156149c957600080fd5b813561287b816149a8565b6000806000606084860312156149e957600080fd5b83356149f4816141ef565b925060208401356001600160401b03811115614a0f57600080fd5b6144f1868287016146ed565b6020808252825182820181905260009190848201906040850190845b8181101561440d57835183529284019291840191600101614a37565b60008060408385031215614a6657600080fd5b8235614a71816141ef565b915060208301356001600160401b0380821115614a8d57600080fd5b9084019060608287031215614aa157600080fd5b604051606081018181108382111715614abc57614abc61423a565b604052823582811115614ace57600080fd5b614ada88828601614419565b82525060208301356020820152604083013560408201528093505050509250929050565b600060208284031215614b1057600080fd5b81356001600160401b03811115614b2657600080fd5b614b3284828501614419565b949350505050565b600080600060608486031215614b4f57600080fd5b8335614b5a816141ef565b92506020840135614b6a816141ef565b91506040840135614502816141ef565b60008060208385031215614b8d57600080fd5b82356001600160401b03811115614ba357600080fd5b614baf858286016147df565b90969095509350505050565b600080600060608486031215614bd057600080fd5b8335614bdb816141ef565b925060208401359150604084013561450281614488565b828152604060208201526000614b32604083018461450d565b60208152600061287b6020830184614614565b60008060208385031215614c3157600080fd5b82356001600160401b03811115614c4757600080fd5b614baf85828601614797565b60008060408385031215614c6657600080fd5b8235614c71816141ef565b91506020830135614c81816149a8565b809150509250929050565b60008060008060608587031215614ca257600080fd5b8435614cad816141ef565b935060208501356001600160401b0380821115614cc957600080fd5b614cd588838901614797565b90955093506040870135915080821115614cee57600080fd5b50850160608188031215614d0157600080fd5b939692955090935050565b600060208284031215614d1e57600080fd5b815161287b816141ef565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614d8557600080fd5b8151801515811461287b57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600060208284031215614e0557600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600060018201614e3457614e34614e0c565b5060010190565b600060208284031215614e4d57600080fd5b81516001600160c01b038116811461287b57600080fd5b600060208284031215614e7657600080fd5b815161287b816149a8565b808201808211156125b0576125b0614e0c565b6001600160601b038116811461061d57600080fd5b600060408284031215614ebb57600080fd5b614ec3614250565b8251614ece816141ef565b81526020830151614ede81614e94565b60208201529392505050565b60006020808385031215614efd57600080fd5b82516001600160401b03811115614f1357600080fd5b8301601f81018513614f2457600080fd5b8051614f32614324826142a8565b81815260059190911b82018301908381019087831115614f5157600080fd5b928401925b82841015614f6f57835182529284019290840190614f56565b979650505050505050565b600060208284031215614f8c57600080fd5b815161287b81614e94565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115614fc457600080fd5b8260051b8085606085013791909101606001949350505050565b60006020808385031215614ff157600080fd5b82516001600160401b0381111561500757600080fd5b8301601f8101851361501857600080fd5b8051615026614324826142a8565b81815260059190911b8201830190838101908783111561504557600080fd5b928401925b82841015614f6f57835161505d81614488565b8252928401929084019061504a565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006150b560408301848661506c565b95945050505050565b6000602082840312156150d057600080fd5b815161287b81614488565b600060ff821660ff81036150f1576150f1614e0c565b60010192915050565b60408152600061510e60408301858761506c565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201526000614b32604083018461437e565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526151e560a08401826145e8565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061287b60208301846145e8565b6000615227614324846142a8565b80848252602080830192508560051b85013681111561524557600080fd5b855b818110156152805780356001600160401b038111156152665760008081fd5b61527236828a01614419565b865250938201938201615247565b50919695505050505050565b8183823760009101908152919050565b600181811c908216806152b057607f821691505b6020821081036152d057634e487b7160e01b600052602260045260246000fd5b50919050565b6000604082840312156152e857600080fd5b6152f0614250565b82518152602083015160038110614ede57600080fd5b6001600160a01b0383168152604060208201526000823536849003601e1901811261533057600080fd5b83016020810190356001600160401b0381111561534c57600080fd5b80360382131561535b57600080fd5b6060604085015261537060a08501828461506c565b9150506020840135606084015260408401356080840152809150509392505050565b60008235609e198336030181126153a857600080fd5b9190910192915050565b80356144a5816141ef565b8183526000602080850194508260005b858110156143ae5781356153e0816141ef565b6001600160a01b03168752818301356153f881614e94565b6001600160601b03168784015260409687019691909101906001016153cd565b60208082528181018390526000906040808401600586901b8501820187855b8881101561551857878303603f190184528135368b9003609e1901811261545d57600080fd5b8a0160a0813536839003601e1901811261547657600080fd5b820188810190356001600160401b0381111561549157600080fd5b8060061b36038213156154a357600080fd5b8287526154b383880182846153bd565b925050506154c28883016153b2565b6001600160a01b031688860152818701358786015260606154e481840161449a565b63ffffffff169086015260806154fb83820161449a565b63ffffffff16950194909452509285019290850190600101615437565b509098975050505050505050565b818103818111156125b0576125b0614e0c565b634e487b7160e01b600052603160045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600082516153a88184602087016145c4565b600061ffff8083168181036155c3576155c3614e0c565b6001019392505050565b601f821115613bbb57600081815260208120601f850160051c810160208610156155f45750805b601f850160051c820191505b8181101561213457828155600101615600565b81516001600160401b0381111561562c5761562c61423a565b6156408161563a845461529c565b846155cd565b602080601f831160018114615675576000841561565d5750858301515b600019600386901b1c1916600185901b178555612134565b600085815260208120601f198616915b828110156156a457888601518255948401946001909101908401615685565b50858210156156c25787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b0383168152604060208201819052600090614b32908301846145e856fed4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f04d4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00a26469706673582212202635a716871621f15c282a660ede6b99197fa733bb5daa061f15c3f30c90585464736f6c63430008150033",
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

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZrServiceManager.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerSession) TaskManager() (common.Address, error) {
	return _ContractZrServiceManager.Contract.TaskManager(&_ContractZrServiceManager.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractZrServiceManager *ContractZrServiceManagerCallerSession) TaskManager() (common.Address, error) {
	return _ContractZrServiceManager.Contract.TaskManager(&_ContractZrServiceManager.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _taskManager) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, _initialOwner common.Address, _taskManager common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.contract.Transact(opts, "initialize", _pauserRegistry, _initialOwner, _taskManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _taskManager) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerSession) Initialize(_pauserRegistry common.Address, _initialOwner common.Address, _taskManager common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Initialize(&_ContractZrServiceManager.TransactOpts, _pauserRegistry, _initialOwner, _taskManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address _initialOwner, address _taskManager) returns()
func (_ContractZrServiceManager *ContractZrServiceManagerTransactorSession) Initialize(_pauserRegistry common.Address, _initialOwner common.Address, _taskManager common.Address) (*types.Transaction, error) {
	return _ContractZrServiceManager.Contract.Initialize(&_ContractZrServiceManager.TransactOpts, _pauserRegistry, _initialOwner, _taskManager)
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
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAssigned is a free log retrieval operation binding the contract event 0xfba1d8bd64cc55f778b2832392428433c24bd739c106b234ab0a0bee8a1a7f03.
//
// Solidity: event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash)
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

// WatchOperatorAssigned is a free log subscription operation binding the contract event 0xfba1d8bd64cc55f778b2832392428433c24bd739c106b234ab0a0bee8a1a7f03.
//
// Solidity: event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash)
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

// ParseOperatorAssigned is a log parse operation binding the contract event 0xfba1d8bd64cc55f778b2832392428433c24bd739c106b234ab0a0bee8a1a7f03.
//
// Solidity: event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash)
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
