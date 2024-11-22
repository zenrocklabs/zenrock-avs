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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIZrRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"oprAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectValidators\",\"inputs\":[{\"name\":\"inactiveValidatorSet\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllOperatorsAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structZrServiceManagerLib.Validator[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEigenStake\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakedBalanceForValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Validator\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"contractIZRTaskManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIZRTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAssigned\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDeregistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorsEjected\",\"inputs\":[{\"name\":\"validatorHashes\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyValidatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IndexOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTaskManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownValidatorHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005bd738038062005bd7833981016040819052620000359162000160565b60016000556001600160a01b0380851660805280841660a05280831660c052811660e052838383836200006762000085565b505050506200007b6200008560201b60201c565b50505050620001c8565b603354610100900460ff1615620000f25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60335460ff908116101562000145576033805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015d57600080fd5b50565b600080600080608085870312156200017757600080fd5b8451620001848162000147565b6020860151909450620001978162000147565b6040860151909350620001aa8162000147565b6060860151909250620001bd8162000147565b939692955090935050565b60805160a05160c05160e05161592d620002aa60003960008181610a4c01528181610ba701528181610c3e0152818161291e01528181612af001528181612c730152612d12015260008181610877015281816109060152818161098601528181611f0401528181612057015281816121420152818161220e0152818161294801528181612a2b01528181612bce01528181612e680152818161313101526141d20152600081816132f3015281816133af015261349b015260008181610352015281816121960152818161229a015281816123190152612f60015261592d6000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c80639926ee7d1161011a578063d604902f116100ad578063f891899d1161007c578063f891899d146104e7578063fabc1cbc146104fa578063fb558ed51461050d578063fc299dee14610515578063fce36c7d1461052857600080fd5b8063d604902f1461048e578063da4dc49c146104b9578063e481af9d146104cc578063f2fde38b146104d457600080fd5b8063c0c53b8b116100e9578063c0c53b8b14610427578063c63e3c501461043a578063cefdc1d41461044d578063d5f20ff61461046e57600080fd5b80639926ee7d146103be578063a364f4da146103d1578063a50a640e146103e4578063a98fb3551461041457600080fd5b8063595c6a671161019d5780636b3aa72e1161016c5780636b3aa72e14610350578063715018a61461038a578063886f119514610392578063889e2c6e146103a55780638da5cb5b146103ad57600080fd5b8063595c6a67146102e45780635ac86ab7146102ec5780635c1556621461031f5780635c975abb1461033f57600080fd5b80633563b0d1116101d95780633563b0d11461027c5780634a91a2f81461029c5780634d2b57fe146102b15780634f739f74146102c457600080fd5b806310d67a2f1461020b578063136439dd1461022057806331b36bd91461023357806333cfb7b71461025c575b600080fd5b61021e6102193660046143ea565b61053b565b005b61021e61022e366004614407565b6105f7565b6102466102413660046144b1565b610736565b604051610253919061459f565b60405180910390f35b61026f61026a3660046143ea565b610852565b60405161025391906145b2565b61028f61028a366004614690565b610d21565b6040516102539190614797565b6102a46111b7565b6040516102539190614867565b61026f6102bf366004614924565b61141e565b6102d76102d23660046149ff565b611533565b6040516102539190614ac9565b61021e611c5b565b61030f6102fa366004614b93565b600254600160ff9092169190911b9081161490565b6040519015158152602001610253565b61033261032d366004614bb0565b611d22565b6040516102539190614bf7565b600254604051908152602001610253565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610253565b61021e611eea565b600154610372906001600160a01b031681565b61026f611efe565b6066546001600160a01b0316610372565b61021e6103cc366004614c2f565b612137565b61021e6103df3660046143ea565b612203565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f04546001600160a01b0316610372565b61021e610422366004614cda565b6122fa565b61021e610435366004614d16565b61234e565b61021e610448366004614d56565b61251a565b61046061045b366004614d97565b6125ad565b604051610253929190614dce565b61048161047c366004614407565b61273f565b6040516102539190614de7565b6104a161049c366004614dfa565b61276a565b6040516001600160601b039091168152602001610253565b6104a16104c7366004614e2f565b6128fc565b61026f612a25565b61021e6104e23660046143ea565b612de7565b61021e6104f5366004614e68565b612e5d565b61021e610508366004614407565b612fd3565b61021e61312f565b609854610372906001600160a01b031681565b61021e610536366004614d56565b6131af565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561058e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b29190614ee8565b6001600160a01b0316336001600160a01b0316146105eb5760405162461bcd60e51b81526004016105e290614f05565b60405180910390fd5b6105f4816134d2565b50565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561063f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106639190614f4f565b61067f5760405162461bcd60e51b81526004016105e290614f71565b600254818116146106f85760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016105e2565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b606081516001600160401b0381111561075157610751614420565b60405190808252806020026020018201604052801561077a578160200160208202803683370190505b50905060005b825181101561084b57836001600160a01b03166313542a4e8483815181106107aa576107aa614fb9565b60200260200101516040518263ffffffff1660e01b81526004016107dd91906001600160a01b0391909116815260200190565b602060405180830381865afa1580156107fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081e9190614fcf565b82828151811061083057610830614fb9565b602090810291909101015261084481614ffe565b9050610780565b5092915050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156108be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e29190614fcf565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561094d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109719190615017565b90506001600160c01b0381161580610a0b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a069190615040565b60ff16155b15610a2757505060408051600081526020810190915292915050565b6000610a3b826001600160c01b03166135c9565b90506000805b8251811015610b11577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610a8b57610a8b614fb9565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610acf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af39190614fcf565b610afd908361505d565b915080610b0981614ffe565b915050610a41565b506000816001600160401b03811115610b2c57610b2c614420565b604051908082528060200260200182016040528015610b55578160200160208202803683370190505b5090506000805b8451811015610d14576000858281518110610b7957610b79614fb9565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610bee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c129190614fcf565b905060005b81811015610cfe576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610c8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb09190615085565b60000151868681518110610cc657610cc6614fb9565b6001600160a01b039092166020928302919091019091015284610ce881614ffe565b9550508080610cf690614ffe565b915050610c17565b5050508080610d0c90614ffe565b915050610b5c565b5090979650505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d879190614ee8565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ded9190614ee8565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e539190614ee8565b9050600086516001600160401b03811115610e7057610e70614420565b604051908082528060200260200182016040528015610ea357816020015b6060815260200190600190039081610e8e5790505b50905060005b87518110156111ab576000888281518110610ec657610ec6614fb9565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610f27573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f4f91908101906150c6565b905080516001600160401b03811115610f6a57610f6a614420565b604051908082528060200260200182016040528015610fb557816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610f885790505b50848481518110610fc857610fc8614fb9565b602002602001018190525060005b8151811015611195576040518060600160405280876001600160a01b03166347b314e885858151811061100b5761100b614fb9565b60200260200101516040518263ffffffff1660e01b815260040161103191815260200190565b602060405180830381865afa15801561104e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110729190614ee8565b6001600160a01b0316815260200183838151811061109257611092614fb9565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106110c0576110c0614fb9565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa15801561111c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111409190615156565b6001600160601b031681525085858151811061115e5761115e614fb9565b6020026020010151828151811061117757611177614fb9565b6020026020010181905250808061118d90614ffe565b915050610fd6565b50505080806111a390614ffe565b915050610ea9565b50979650505050505050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f03546060906000805160206158d8833981519152906000816001600160401b0381111561120657611206614420565b60405190808252806020026020018201604052801561125357816020015b604080516060808201835280825260006020830152918101919091528152602001906001900390816112245790505b5090506000836003018054806020026020016040519081016040528092919081815260200182805480156112a657602002820191906000526020600020905b815481526020019060010190808311611292575b5050505050905060005b83811015611414578460000160008383815181106112d0576112d0614fb9565b6020026020010151815260200190815260200160002060405180606001604052908160008201805461130190615173565b80601f016020809104026020016040519081016040528092919081815260200182805461132d90615173565b801561137a5780601f1061134f5761010080835404028352916020019161137a565b820191906000526020600020905b81548152906001019060200180831161135d57829003601f168201915b5050505050815260200160018201548152602001600282018054806020026020016040519081016040528092919081815260200182805480156113e657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113c8575b50505050508152505083828151811061140157611401614fb9565b60209081029190910101526001016112b0565b5090949350505050565b606081516001600160401b0381111561143957611439614420565b604051908082528060200260200182016040528015611462578160200160208202803683370190505b50905060005b825181101561084b57836001600160a01b031663296bb06484838151811061149257611492614fb9565b60200260200101516040518263ffffffff1660e01b81526004016114b891815260200190565b602060405180830381865afa1580156114d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114f99190614ee8565b82828151811061150b5761150b614fb9565b6001600160a01b039092166020928302919091019091015261152c81614ffe565b9050611468565b61155e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561159e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115c29190614ee8565b90506115ef6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061161f908b90899089906004016151ad565b600060405180830381865afa15801561163c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261166491908101906151f4565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611696908b908b908b906004016152ab565b600060405180830381865afa1580156116b3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116db91908101906151f4565b6040820152856001600160401b038111156116f8576116f8614420565b60405190808252806020026020018201604052801561172b57816020015b60608152602001906001900390816117165790505b50606082015260005b60ff8116871115611b6c576000856001600160401b0381111561175957611759614420565b604051908082528060200260200182016040528015611782578160200160208202803683370190505b5083606001518360ff168151811061179c5761179c614fb9565b602002602001018190525060005b86811015611a6c5760008c6001600160a01b03166304ec63518a8a858181106117d5576117d5614fb9565b905060200201358e886000015186815181106117f3576117f3614fb9565b60200260200101516040518463ffffffff1660e01b81526004016118309392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561184d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118719190615017565b9050806001600160c01b03166000036119185760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a4016105e2565b8a8a8560ff1681811061192d5761192d614fb9565b60016001600160c01b038516919093013560f81c1c82169091039050611a5957856001600160a01b031663dd9846b98a8a8581811061196e5761196e614fb9565b905060200201358d8d8860ff1681811061198a5761198a614fb9565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156119e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a0491906152d4565b85606001518560ff1681518110611a1d57611a1d614fb9565b60200260200101518481518110611a3657611a36614fb9565b63ffffffff9092166020928302919091019091015282611a5581614ffe565b9350505b5080611a6481614ffe565b9150506117aa565b506000816001600160401b03811115611a8757611a87614420565b604051908082528060200260200182016040528015611ab0578160200160208202803683370190505b50905060005b82811015611b315784606001518460ff1681518110611ad757611ad7614fb9565b60200260200101518181518110611af057611af0614fb9565b6020026020010151828281518110611b0a57611b0a614fb9565b63ffffffff9092166020928302919091019091015280611b2981614ffe565b915050611ab6565b508084606001518460ff1681518110611b4c57611b4c614fb9565b602002602001018190525050508080611b64906152f1565b915050611734565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd19190614ee8565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611c04908b908b908e90600401615310565b600060405180830381865afa158015611c21573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c4991908101906151f4565b60208301525098975050505050505050565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611ca3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cc79190614f4f565b611ce35760405162461bcd60e51b81526004016105e290614f71565b600019600281905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611d5492919061533a565b600060405180830381865afa158015611d71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d9991908101906151f4565b9050600084516001600160401b03811115611db657611db6614420565b604051908082528060200260200182016040528015611ddf578160200160208202803683370190505b50905060005b8551811015611ee057866001600160a01b03166304ec6351878381518110611e0f57611e0f614fb9565b602002602001015187868581518110611e2a57611e2a614fb9565b60200260200101516040518463ffffffff1660e01b8152600401611e679392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611e84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea89190615017565b6001600160c01b0316828281518110611ec357611ec3614fb9565b602090810291909101015280611ed881614ffe565b915050611de5565b5095945050505050565b611ef261368b565b611efc60006136e5565b565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f849190614ee8565b604051638902624560e01b81526000600482015263ffffffff431660248201526001600160a01b039190911690638902624590604401600060405180830381865afa158015611fd7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611fff91908101906150c6565b80519091506000816001600160401b0381111561201e5761201e614420565b604051908082528060200260200182016040528015612047578160200160208202803683370190505b50905060005b8281101561212f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663296bb06485838151811061209657612096614fb9565b60200260200101516040518263ffffffff1660e01b81526004016120bc91815260200190565b602060405180830381865afa1580156120d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120fd9190614ee8565b82828151811061210f5761210f614fb9565b6001600160a01b039092166020928302919091019091015260010161204d565b509392505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461217f5760405162461bcd60e51b81526004016105e290615359565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d906121cd90859085906004016153d1565b600060405180830381600087803b1580156121e757600080fd5b505af11580156121fb573d6000803e3d6000fd5b505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461224b5760405162461bcd60e51b81526004016105e290615359565b6001600160a01b0381166122725760405163eb32d3bf60e01b815260040160405180910390fd5b61227b81613737565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b1580156122df57600080fd5b505af11580156122f3573d6000803e3d6000fd5b5050505050565b61230261368b565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906122c590849060040161541c565b603354610100900460ff161580801561236e5750603354600160ff909116105b806123885750303b158015612388575060335460ff166001145b6123eb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105e2565b6033805460ff19166001179055801561240e576033805461ff0019166101001790555b6001600160a01b0382166124355760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03831661245c5760405163d92e233d60e01b815260040160405180910390fd5b612467846000613994565b61247083613a7e565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0480546001600160a01b0319166001600160a01b03841617905560007fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f05558015612514576033805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f04546000805160206158d8833981519152906001600160a01b031633146125745760405163f06b6b4360e01b815260040160405180910390fd5b60008290036125965760405163339e1ffb60e01b815260040160405180910390fd5b6125a86125a3838561542f565b613aae565b505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106125e8576125e8614fb9565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612624908890869060040161533a565b600060405180830381865afa158015612641573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261266991908101906151f4565b60008151811061267b5761267b614fb9565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156126e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061270b9190615017565b6001600160c01b031690506000612721826135c9565b90508161272f8a838a610d21565b9550955050505050935093915050565b6040805160608082018352808252600060208301529181019190915261276482613b94565b92915050565b600080838360405161277d9291906154a2565b60405180910390209050600061279e6000805160206158d883398151915290565b60008381526020919091526040908190208151606081019092528054829082906127c790615173565b80601f01602080910402602001604051908101604052809291908181526020018280546127f390615173565b80156128405780601f1061281557610100808354040283529160200191612840565b820191906000526020600020905b81548152906001019060200180831161282357829003601f168201915b5050505050815260200160018201548152602001600282018054806020026020016040519081016040528092919081815260200182805480156128ac57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161288e575b50505050508152505090506000805b826040015151811015611ee0576128f0836040015182815181106128e1576128e1614fb9565b602002602001015160006128fc565b909101906001016128bb565b604051631619718360e21b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000821691635401ed27917f000000000000000000000000000000000000000000000000000000000000000090911690635865c60c906024016040805180830381865afa158015612990573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129b491906154b2565b5160405160e083901b6001600160e01b0319168152600481019190915260ff85166024820152604401602060405180830381865afa1580156129fa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a1e9190615156565b9392505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aab9190615040565b60ff16905080600003612acc57505060408051600081526020810190915290565b6000805b82811015612b8157604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015612b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b639190614fcf565b612b6d908361505d565b915080612b7981614ffe565b915050612ad0565b506000816001600160401b03811115612b9c57612b9c614420565b604051908082528060200260200182016040528015612bc5578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c4e9190615040565b60ff1681101561141457604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015612cc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ce69190614fcf565b905060005b81811015612dd2576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015612d60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d849190615085565b60000151858581518110612d9a57612d9a614fb9565b6001600160a01b039092166020928302919091019091015283612dbc81614ffe565b9450508080612dca90614ffe565b915050612ceb565b50508080612ddf90614ffe565b915050612bcc565b612def61368b565b6001600160a01b038116612e545760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105e2565b6105f4816136e5565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612ea55760405162461bcd60e51b81526004016105e290615359565b6001600160a01b038416612ecc5760405163eb32d3bf60e01b815260040160405180910390fd5b6000829003612eee5760405163713ce51160e01b815260040160405180910390fd5b60008383604051612f009291906154a2565b60405180910390209050612f4984848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613cdd92505050565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90612f9790889086906004016154e2565b600060405180830381600087803b158015612fb157600080fd5b505af1158015612fc5573d6000803e3d6000fd5b505050506122f38582613d1a565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613026573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061304a9190614ee8565b6001600160a01b0316336001600160a01b03161461307a5760405162461bcd60e51b81526004016105e290614f05565b6002541981196002541916146130f85760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016105e2565b600281905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161072b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031662cf2ab5613165611efe565b6040518263ffffffff1660e01b815260040161318191906145b2565b600060405180830381600087803b15801561319b57600080fd5b505af1158015612514573d6000803e3d6000fd5b60005b81811015613483578282828181106131cc576131cc614fb9565b90506020028101906131de919061556e565b6131ef9060408101906020016143ea565b6001600160a01b03166323b872dd333086868681811061321157613211614fb9565b9050602002810190613223919061556e565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af115801561327a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061329e9190614f4f565b5060008383838181106132b3576132b3614fb9565b90506020028101906132c5919061556e565b6132d69060408101906020016143ea565b604051636eb1769f60e11b81523060048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166024830152919091169063dd62ed3e90604401602060405180830381865afa158015613344573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133689190614fcf565b905083838381811061337c5761337c614fb9565b905060200281019061338e919061556e565b61339f9060408101906020016143ea565b6001600160a01b031663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000838787878181106133e1576133e1614fb9565b90506020028101906133f3919061556e565b60400135613401919061505d565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561344c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134709190614f4f565b50508061347c90614ffe565b90506131b2565b5060405163fce36c7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063fce36c7d906121cd90859085906004016155f4565b6001600160a01b0381166135605760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016105e2565b600154604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b60606000806135d784613e49565b61ffff166001600160401b038111156135f2576135f2614420565b6040519080825280601f01601f19166020018201604052801561361c576020820181803683370190505b5090506000805b825182108015613634575061010081105b15611414576001811b93508584161561367b578060f81b83838151811061365d5761365d614fb9565b60200101906001600160f81b031916908160001a9053508160010191505b61368481614ffe565b9050613623565b6066546001600160a01b03163314611efc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105e2565b606680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03811661375e5760405163eb32d3bf60e01b815260040160405180910390fd5b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0160205260409020546000805160206158d883398151915290806137c457604051631e7b7a8160e01b815260040160405180910390fd5b6001600160a01b0383166000908152600280840160209081526040808420548585529186905290922090810154821061381057604051634e23d03560e01b815260040160405180910390fd5b846001600160a01b031681600201838154811061382f5761382f614fb9565b6000918252602090912001546001600160a01b03161461386257604051634265bbb960e11b815260040160405180910390fd5b600281015460009061387690600190615702565b905080831461390457600082600201828154811061389657613896614fb9565b6000918252602090912001546002840180546001600160a01b0390921692508291869081106138c7576138c7614fb9565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394851617905592909116815260028701909152604090208390555b8160020180548061391757613917615715565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038816808352600188018252604080842084905560028901909252818320839055905186927f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a3505050505050565b6001546001600160a01b03161580156139b557506001600160a01b03821615155b613a375760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016105e2565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613a7a826134d2565b5050565b603354610100900460ff16613aa55760405162461bcd60e51b81526004016105e29061572b565b6105f481613e74565b80516000816001600160401b03811115613aca57613aca614420565b604051908082528060200260200182016040528015613af3578160200160208202803683370190505b50905060005b82811015613b57576000848281518110613b1557613b15614fb9565b602002602001015180519060200120905080838381518110613b3957613b39614fb9565b602002602001018181525050613b4e81613e9b565b50600101613af9565b507fb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a81604051613b87919061459f565b60405180910390a1505050565b604080516060808201835280825260006020830152918101919091526000805160206158d88339815191526000838152602091909152604090819020815160608101909252805482908290613be890615173565b80601f0160208091040260200160405190810160405280929190818152602001828054613c1490615173565b8015613c615780601f10613c3657610100808354040283529160200191613c61565b820191906000526020600020905b815481529060010190602001808311613c4457829003601f168201915b505050505081526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015613ccd57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613caf575b5050505050815250509050919050565b8051602082012060006000805160206158d8833981519152600083815260068201602052604090205490915060ff166125a8576125148383613fdf565b6001600160a01b038216613d415760405163eb32d3bf60e01b815260040160405180910390fd5b80613d5f57604051631e7b7a8160e01b815260040160405180910390fd5b60008181527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0660205260409020546000805160206158d88339815191529060ff16613dbd57604051634d709ad360e11b815260040160405180910390fd5b6001600160a01b0383166000818152600180840160209081526040808420879055868452858252808420600290810180548787529188018452828620829055878452938101845592845290832090910180546001600160a01b03191684179055518492917ffba1d8bd64cc55f778b2832392428433c24bd739c106b234ab0a0bee8a1a7f0391a3505050565b6000805b821561276457613e5e600184615702565b9092169180613e6c81615776565b915050613e4d565b603354610100900460ff16612e545760405162461bcd60e51b81526004016105e29061572b565b60008181527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0660205260409020546000805160206158d88339815191529060ff16613ee4575050565b6000828152602082905260408120600281015490915b81811015613f78576000836002018281548110613f1957613f19614fb9565b6000918252602090912001546001600160a01b03169050613f3981614137565b60405186906001600160a01b038316907f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e490600090a350600101613efa565b50600084815260208490526040812090613f928282614303565b6001820160009055600282016000613faa919061433d565b505060008481526006840160205260408120805460ff1916905560058401805491613fd483615797565b919050555050505050565b604080516060808201835280825260006020830181905292820152906000805160206158d8833981519152600084815260068201602052604090205490915060ff161561403f5760405163132e7efb60e31b815260040160405180910390fd5b6040805160608101825285815260208082018690528251600080825291810184529092820152600085815260208490526040902081519192508291819061408690826157f4565b506020828101516001830155604083015180516140a9926002850192019061435b565b50505060038201805460018181018355600092835260208084209092018790558683526006850190915260408220805460ff19169091179055600583018054916140f283614ffe565b9190505550837f1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea386604051614127919061541c565b60405180910390a2949350505050565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0160205260409020546000805160206158d88339815191529080156125a857604080516001808252818301909252600091602080830190803683370190505090506000816000815181106141bc576141bc614fb9565b602002602001019060ff16908160ff16815250507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e3b17db856142098461425e565b6040518363ffffffff1660e01b81526004016142269291906158b3565b600060405180830381600087803b15801561424057600080fd5b505af1158015614254573d6000803e3d6000fd5b5050505050505050565b80516060906000816001600160401b0381111561427d5761427d614420565b6040519080825280601f01601f1916602001820160405280156142a7576020820181803683370190505b50905060005b8281101561212f578481815181106142c7576142c7614fb9565b602002602001015160f81b8282815181106142e4576142e4614fb9565b60200101906001600160f81b031916908160001a9053506001016142ad565b50805461430f90615173565b6000825580601f1061431f575050565b601f0160209004906000526020600020908101906105f491906143c0565b50805460008255906000526020600020908101906105f491906143c0565b8280548282559060005260206000209081019282156143b0579160200282015b828111156143b057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061437b565b506143bc9291506143c0565b5090565b5b808211156143bc57600081556001016143c1565b6001600160a01b03811681146105f457600080fd5b6000602082840312156143fc57600080fd5b8135612a1e816143d5565b60006020828403121561441957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561445857614458614420565b60405290565b604051601f8201601f191681016001600160401b038111828210171561448657614486614420565b604052919050565b60006001600160401b038211156144a7576144a7614420565b5060051b60200190565b600080604083850312156144c457600080fd5b82356144cf816143d5565b91506020838101356001600160401b038111156144eb57600080fd5b8401601f810186136144fc57600080fd5b803561450f61450a8261448e565b61445e565b81815260059190911b8201830190838101908883111561452e57600080fd5b928401925b82841015614555578335614546816143d5565b82529284019290840190614533565b80955050505050509250929050565b600081518084526020808501945080840160005b8381101561459457815187529582019590820190600101614578565b509495945050505050565b602081526000612a1e6020830184614564565b6020808252825182820181905260009190848201906040850190845b818110156145f35783516001600160a01b0316835292840192918401916001016145ce565b50909695505050505050565b600082601f83011261461057600080fd5b81356001600160401b0381111561462957614629614420565b61463c601f8201601f191660200161445e565b81815284602083860101111561465157600080fd5b816020850160208301376000918101602001919091529392505050565b63ffffffff811681146105f457600080fd5b803561468b8161466e565b919050565b6000806000606084860312156146a557600080fd5b83356146b0816143d5565b925060208401356001600160401b038111156146cb57600080fd5b6146d7868287016145ff565b92505060408401356146e88161466e565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614789578385038a52825180518087529087019087870190845b8181101561477457835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614730565b50509a87019a95505091850191600101614712565b509298975050505050505050565b602081526000612a1e60208301846146f3565b6000815180845260005b818110156147d0576020818501810151868301820152016147b4565b506000602082860101526020601f19601f83011685010191505092915050565b600081516060845261480560608501826147aa565b9050602080840151818601526040840151858303604087015282815180855283850191508383019450600092505b8083101561485c5784516001600160a01b03168252938301936001929092019190830190614833565b509695505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156148bc57603f198886030184526148aa8583516147f0565b9450928501929085019060010161488e565b5092979650505050505050565b600082601f8301126148da57600080fd5b813560206148ea61450a8361448e565b82815260059290921b8401810191818101908684111561490957600080fd5b8286015b8481101561485c578035835291830191830161490d565b6000806040838503121561493757600080fd5b8235614942816143d5565b915060208301356001600160401b0381111561495d57600080fd5b614969858286016148c9565b9150509250929050565b60008083601f84011261498557600080fd5b5081356001600160401b0381111561499c57600080fd5b6020830191508360208285010111156149b457600080fd5b9250929050565b60008083601f8401126149cd57600080fd5b5081356001600160401b038111156149e457600080fd5b6020830191508360208260051b85010111156149b457600080fd5b60008060008060008060808789031215614a1857600080fd5b8635614a23816143d5565b95506020870135614a338161466e565b945060408701356001600160401b0380821115614a4f57600080fd5b614a5b8a838b01614973565b90965094506060890135915080821115614a7457600080fd5b50614a8189828a016149bb565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561459457815163ffffffff1687529582019590820190600101614aa7565b600060208083528351608082850152614ae560a0850182614a93565b905081850151601f1980868403016040870152614b028383614a93565b92506040870151915080868403016060870152614b1f8383614a93565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614b765784878303018452614b64828751614a93565b95880195938801939150600101614b4a565b509998505050505050505050565b60ff811681146105f457600080fd5b600060208284031215614ba557600080fd5b8135612a1e81614b84565b600080600060608486031215614bc557600080fd5b8335614bd0816143d5565b925060208401356001600160401b03811115614beb57600080fd5b6146d7868287016148c9565b6020808252825182820181905260009190848201906040850190845b818110156145f357835183529284019291840191600101614c13565b60008060408385031215614c4257600080fd5b8235614c4d816143d5565b915060208301356001600160401b0380821115614c6957600080fd5b9084019060608287031215614c7d57600080fd5b604051606081018181108382111715614c9857614c98614420565b604052823582811115614caa57600080fd5b614cb6888286016145ff565b82525060208301356020820152604083013560408201528093505050509250929050565b600060208284031215614cec57600080fd5b81356001600160401b03811115614d0257600080fd5b614d0e848285016145ff565b949350505050565b600080600060608486031215614d2b57600080fd5b8335614d36816143d5565b92506020840135614d46816143d5565b915060408401356146e8816143d5565b60008060208385031215614d6957600080fd5b82356001600160401b03811115614d7f57600080fd5b614d8b858286016149bb565b90969095509350505050565b600080600060608486031215614dac57600080fd5b8335614db7816143d5565b92506020840135915060408401356146e88161466e565b828152604060208201526000614d0e60408301846146f3565b602081526000612a1e60208301846147f0565b60008060208385031215614e0d57600080fd5b82356001600160401b03811115614e2357600080fd5b614d8b85828601614973565b60008060408385031215614e4257600080fd5b8235614e4d816143d5565b91506020830135614e5d81614b84565b809150509250929050565b60008060008060608587031215614e7e57600080fd5b8435614e89816143d5565b935060208501356001600160401b0380821115614ea557600080fd5b614eb188838901614973565b90955093506040870135915080821115614eca57600080fd5b50850160608188031215614edd57600080fd5b939692955090935050565b600060208284031215614efa57600080fd5b8151612a1e816143d5565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614f6157600080fd5b81518015158114612a1e57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600060208284031215614fe157600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60006001820161501057615010614fe8565b5060010190565b60006020828403121561502957600080fd5b81516001600160c01b0381168114612a1e57600080fd5b60006020828403121561505257600080fd5b8151612a1e81614b84565b8082018082111561276457612764614fe8565b6001600160601b03811681146105f457600080fd5b60006040828403121561509757600080fd5b61509f614436565b82516150aa816143d5565b815260208301516150ba81615070565b60208201529392505050565b600060208083850312156150d957600080fd5b82516001600160401b038111156150ef57600080fd5b8301601f8101851361510057600080fd5b805161510e61450a8261448e565b81815260059190911b8201830190838101908783111561512d57600080fd5b928401925b8284101561514b57835182529284019290840190615132565b979650505050505050565b60006020828403121561516857600080fd5b8151612a1e81615070565b600181811c9082168061518757607f821691505b6020821081036151a757634e487b7160e01b600052602260045260246000fd5b50919050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156151da57600080fd5b8260051b8085606085013791909101606001949350505050565b6000602080838503121561520757600080fd5b82516001600160401b0381111561521d57600080fd5b8301601f8101851361522e57600080fd5b805161523c61450a8261448e565b81815260059190911b8201830190838101908783111561525b57600080fd5b928401925b8284101561514b5783516152738161466e565b82529284019290840190615260565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006152cb604083018486615282565b95945050505050565b6000602082840312156152e657600080fd5b8151612a1e8161466e565b600060ff821660ff810361530757615307614fe8565b60010192915050565b604081526000615324604083018587615282565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201526000614d0e6040830184614564565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526153fb60a08401826147aa565b90506020840151606084015260408401516080840152809150509392505050565b602081526000612a1e60208301846147aa565b600061543d61450a8461448e565b80848252602080830192508560051b85013681111561545b57600080fd5b855b818110156154965780356001600160401b0381111561547c5760008081fd5b61548836828a016145ff565b86525093820193820161545d565b50919695505050505050565b8183823760009101908152919050565b6000604082840312156154c457600080fd5b6154cc614436565b825181526020830151600381106150ba57600080fd5b6001600160a01b0383168152604060208201526000823536849003601e1901811261550c57600080fd5b83016020810190356001600160401b0381111561552857600080fd5b80360382131561553757600080fd5b6060604085015261554c60a085018284615282565b9150506020840135606084015260408401356080840152809150509392505050565b60008235609e1983360301811261558457600080fd5b9190910192915050565b803561468b816143d5565b8183526000602080850194508260005b858110156145945781356155bc816143d5565b6001600160a01b03168752818301356155d481615070565b6001600160601b03168784015260409687019691909101906001016155a9565b60208082528181018390526000906040808401600586901b8501820187855b888110156156f457878303603f190184528135368b9003609e1901811261563957600080fd5b8a0160a0813536839003601e1901811261565257600080fd5b820188810190356001600160401b0381111561566d57600080fd5b8060061b360382131561567f57600080fd5b82875261568f8388018284615599565b9250505061569e88830161558e565b6001600160a01b031688860152818701358786015260606156c0818401614680565b63ffffffff169086015260806156d7838201614680565b63ffffffff16950194909452509285019290850190600101615613565b509098975050505050505050565b8181038181111561276457612764614fe8565b634e487b7160e01b600052603160045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600061ffff80831681810361578d5761578d614fe8565b6001019392505050565b6000816157a6576157a6614fe8565b506000190190565b601f8211156125a857600081815260208120601f850160051c810160208610156157d55750805b601f850160051c820191505b818110156121fb578281556001016157e1565b81516001600160401b0381111561580d5761580d614420565b6158218161581b8454615173565b846157ae565b602080601f831160018114615856576000841561583e5750858301515b600019600386901b1c1916600185901b1785556121fb565b600085815260208120601f198616915b8281101561588557888601518255948401946001909101908401615866565b50858210156158a35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b0383168152604060208201819052600090614d0e908301846147aa56fed4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00a264697066735822122036fc6755f92600f2212a697f04cc3ad9034d6f3fd1fc274267a8d157a9ad10db64736f6c63430008150033",
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
