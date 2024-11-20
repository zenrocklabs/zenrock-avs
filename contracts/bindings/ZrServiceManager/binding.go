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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIZrRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"QUORUM_NUMBER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"oprAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectValidators\",\"inputs\":[{\"name\":\"inactiveValidatorSet\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllOperatorsAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structZrServiceManagerLib.Validator[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEigenStake\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakedBalanceForValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structZrServiceManagerLib.Validator\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"contractIZRTaskManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.Task\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAssigned\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponse\",\"components\":[{\"name\":\"referenceTaskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inactiveSetZRChain\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structZrServiceManagerLib.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDeregistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"validatorHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorAddr\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorsEjected\",\"inputs\":[{\"name\":\"validatorHashes\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005be538038062005be5833981016040819052620000359162000150565b60016000556001600160a01b0380851660805280841660a05280831660c052811660e052838383836200006762000075565b5050505050505050620001b8565b603354610100900460ff1615620000e25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60335460ff908116101562000135576033805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014d57600080fd5b50565b600080600080608085870312156200016757600080fd5b8451620001748162000137565b6020860151909450620001878162000137565b60408601519093506200019a8162000137565b6060860151909250620001ad8162000137565b939692955090935050565b60805160a05160c05160e05161594b6200029a60003960008181610a6d01528181610bc801528181610c5f015281816127f3015281816129c501528181612b480152612be701526000818161089801528181610927015281816109a701528181611e5801528181611faa01528181612097015281816121630152818161281d0152818161290001528181612aa301528181612db10152818161302b01526142440152600081816131a30152818161325f015261334b01526000818161037e015281816120eb015281816121c8015281816122470152612e0e015261594b6000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80638da5cb5b11610125578063da4dc49c116100ad578063f891899d1161007c578063f891899d14610508578063fabc1cbc1461051b578063fb558ed51461052e578063fc299dee14610536578063fce36c7d1461054957600080fd5b8063da4dc49c146104d2578063e481af9d146104e5578063f2fde38b146104ed578063f5c9899d1461050057600080fd5b8063c0c53b8b116100f4578063c0c53b8b14610440578063c63e3c5014610453578063cefdc1d414610466578063d5f20ff614610487578063d604902f146104a757600080fd5b80638da5cb5b146103f65780639926ee7d14610407578063a364f4da1461041a578063a98fb3551461042d57600080fd5b8063595c6a67116101a85780636b3aa72e116101775780636b3aa72e1461037c578063715018a6146103b65780637b654c5d146103be578063886f1195146103db578063889e2c6e146103ee57600080fd5b8063595c6a67146103105780635ac86ab7146103185780635c1556621461034b5780635c975abb1461036b57600080fd5b806333cfb7b7116101ef57806333cfb7b7146102885780633563b0d1146102a85780634a91a2f8146102c85780634d2b57fe146102dd5780634f739f74146102f057600080fd5b806310d67a2f14610221578063136439dd146102365780632dda9fc61461024957806331b36bd914610268575b600080fd5b61023461022f366004614463565b61055c565b005b610234610244366004614480565b610618565b610251600081565b60405160ff90911681526020015b60405180910390f35b61027b61027636600461452a565b610757565b60405161025f9190614618565b61029b610296366004614463565b610873565b60405161025f919061462b565b6102bb6102b6366004614709565b610d42565b60405161025f9190614810565b6102d06111d8565b60405161025f91906148ea565b61029b6102eb3660046149a7565b6112de565b6103036102fe366004614a41565b6113f3565b60405161025f9190614b39565b610234611b1b565b61033b610326366004614c03565b600254600160ff9092169190911b9081161490565b604051901515815260200161025f565b61035e610359366004614c20565b611be2565b60405161025f9190614c67565b60025460405190815260200161025f565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b03909116815260200161025f565b610234611daa565b6103c6611dbe565b60405163ffffffff909116815260200161025f565b60015461039e906001600160a01b031681565b61029b611e52565b6066546001600160a01b031661039e565b610234610415366004614d17565b61208c565b610234610428366004614463565b612158565b61023461043b366004614d5c565b612228565b61023461044e366004614d98565b61227c565b610234610461366004614dd8565b6123c4565b610479610474366004614e88565b612456565b60405161025f929190614ebf565b61049a610495366004614480565b6125e8565b60405161025f9190614ed8565b6104ba6104b5366004614d5c565b612613565b6040516001600160601b03909116815260200161025f565b6104ba6104e0366004614eeb565b6127d1565b61029b6128fa565b6102346104fb366004614463565b612cc6565b6103c6612d3c565b610234610516366004614f24565b612da6565b610234610529366004614480565b612eac565b610234613008565b60985461039e906001600160a01b031681565b610234610557366004614f99565b61305f565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d39190614fda565b6001600160a01b0316336001600160a01b03161461060c5760405162461bcd60e51b815260040161060390614ff7565b60405180910390fd5b61061581613382565b50565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610660573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106849190615041565b6106a05760405162461bcd60e51b815260040161060390615063565b600254818116146107195760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610603565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b606081516001600160401b0381111561077257610772614499565b60405190808252806020026020018201604052801561079b578160200160208202803683370190505b50905060005b825181101561086c57836001600160a01b03166313542a4e8483815181106107cb576107cb6150ab565b60200260200101516040518263ffffffff1660e01b81526004016107fe91906001600160a01b0391909116815260200190565b602060405180830381865afa15801561081b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083f91906150c1565b828281518110610851576108516150ab565b6020908102919091010152610865816150f0565b90506107a1565b5092915050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156108df573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090391906150c1565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561096e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109929190615109565b90506001600160c01b0381161580610a2c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a279190615132565b60ff16155b15610a4857505060408051600081526020810190915292915050565b6000610a5c826001600160c01b0316613479565b90506000805b8251811015610b32577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610aac57610aac6150ab565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610af0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b1491906150c1565b610b1e908361514f565b915080610b2a816150f0565b915050610a62565b506000816001600160401b03811115610b4d57610b4d614499565b604051908082528060200260200182016040528015610b76578160200160208202803683370190505b5090506000805b8451811015610d35576000858281518110610b9a57610b9a6150ab565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610c0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3391906150c1565b905060005b81811015610d1f576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610cad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd19190615177565b60000151868681518110610ce757610ce76150ab565b6001600160a01b039092166020928302919091019091015284610d09816150f0565b9550508080610d17906150f0565b915050610c38565b5050508080610d2d906150f0565b915050610b7d565b5090979650505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da89190614fda565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0e9190614fda565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e749190614fda565b9050600086516001600160401b03811115610e9157610e91614499565b604051908082528060200260200182016040528015610ec457816020015b6060815260200190600190039081610eaf5790505b50905060005b87518110156111cc576000888281518110610ee757610ee76150ab565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610f48573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f7091908101906151b8565b905080516001600160401b03811115610f8b57610f8b614499565b604051908082528060200260200182016040528015610fd657816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610fa95790505b50848481518110610fe957610fe96150ab565b602002602001018190525060005b81518110156111b6576040518060600160405280876001600160a01b03166347b314e885858151811061102c5761102c6150ab565b60200260200101516040518263ffffffff1660e01b815260040161105291815260200190565b602060405180830381865afa15801561106f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110939190614fda565b6001600160a01b031681526020018383815181106110b3576110b36150ab565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106110e1576110e16150ab565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa15801561113d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111619190615248565b6001600160601b031681525085858151811061117f5761117f6150ab565b60200260200101518281518110611198576111986150ab565b602002602001018190525080806111ae906150f0565b915050610ff7565b50505080806111c4906150f0565b915050610eca565b50979650505050505050565b7fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f03546060906000805160206158f6833981519152906000906001600160401b0381111561122757611227614499565b60405190808252806020026020018201604052801561127457816020015b604080516060808201835280825260006020830152918101919091528152602001906001900390816112455790505b50905060005b600383015481101561086c576112ae83600301828154811061129e5761129e6150ab565b906000526020600020015461353b565b8282815181106112c0576112c06150ab565b602002602001018190525080806112d6906150f0565b91505061127a565b606081516001600160401b038111156112f9576112f9614499565b604051908082528060200260200182016040528015611322578160200160208202803683370190505b50905060005b825181101561086c57836001600160a01b031663296bb064848381518110611352576113526150ab565b60200260200101516040518263ffffffff1660e01b815260040161137891815260200190565b602060405180830381865afa158015611395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113b99190614fda565b8282815181106113cb576113cb6150ab565b6001600160a01b03909216602092830291909101909101526113ec816150f0565b9050611328565b61141e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561145e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114829190614fda565b90506114af6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906114df908b9089908990600401615265565b600060405180830381865afa1580156114fc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261152491908101906152ac565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611556908b908b908b90600401615363565b600060405180830381865afa158015611573573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261159b91908101906152ac565b6040820152856001600160401b038111156115b8576115b8614499565b6040519080825280602002602001820160405280156115eb57816020015b60608152602001906001900390816115d65790505b50606082015260005b60ff8116871115611a2c576000856001600160401b0381111561161957611619614499565b604051908082528060200260200182016040528015611642578160200160208202803683370190505b5083606001518360ff168151811061165c5761165c6150ab565b602002602001018190525060005b8681101561192c5760008c6001600160a01b03166304ec63518a8a85818110611695576116956150ab565b905060200201358e886000015186815181106116b3576116b36150ab565b60200260200101516040518463ffffffff1660e01b81526004016116f09392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561170d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117319190615109565b9050806001600160c01b03166000036117d85760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610603565b8a8a8560ff168181106117ed576117ed6150ab565b60016001600160c01b038516919093013560f81c1c8216909103905061191957856001600160a01b031663dd9846b98a8a8581811061182e5761182e6150ab565b905060200201358d8d8860ff1681811061184a5761184a6150ab565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156118a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118c4919061538c565b85606001518560ff16815181106118dd576118dd6150ab565b602002602001015184815181106118f6576118f66150ab565b63ffffffff9092166020928302919091019091015282611915816150f0565b9350505b5080611924816150f0565b91505061166a565b506000816001600160401b0381111561194757611947614499565b604051908082528060200260200182016040528015611970578160200160208202803683370190505b50905060005b828110156119f15784606001518460ff1681518110611997576119976150ab565b602002602001015181815181106119b0576119b06150ab565b60200260200101518282815181106119ca576119ca6150ab565b63ffffffff90921660209283029190910190910152806119e9816150f0565b915050611976565b508084606001518460ff1681518110611a0c57611a0c6150ab565b602002602001018190525050508080611a24906153a9565b9150506115f4565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a919190614fda565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611ac4908b908b908e906004016153c8565b600060405180830381865afa158015611ae1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b0991908101906152ac565b60208301525098975050505050505050565b60015460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611b63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b879190615041565b611ba35760405162461bcd60e51b815260040161060390615063565b600019600281905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611c149291906153f2565b600060405180830381865afa158015611c31573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c5991908101906152ac565b9050600084516001600160401b03811115611c7657611c76614499565b604051908082528060200260200182016040528015611c9f578160200160208202803683370190505b50905060005b8551811015611da057866001600160a01b03166304ec6351878381518110611ccf57611ccf6150ab565b602002602001015187868581518110611cea57611cea6150ab565b60200260200101516040518463ffffffff1660e01b8152600401611d279392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611d44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d689190615109565b6001600160c01b0316828281518110611d8357611d836150ab565b602090810291909101015280611d98816150f0565b915050611ca5565b5095945050505050565b611db2613684565b611dbc60006136de565b565b6000805160206158d683398151915254604080516372d18e8d60e01b815290516000926000805160206158f6833981519152926001600160a01b03909116916372d18e8d916004818101926020929091908290030181865afa158015611e28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e4c919061538c565b91505090565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015611eb4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ed89190614fda565b604051638902624560e01b81526000600482015263ffffffff431660248201526001600160a01b039190911690638902624590604401600060405180830381865afa158015611f2b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f5391908101906151b8565b9050600081516001600160401b03811115611f7057611f70614499565b604051908082528060200260200182016040528015611f99578160200160208202803683370190505b50905060005b825181101561086c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663296bb064848381518110611fe957611fe96150ab565b60200260200101516040518263ffffffff1660e01b815260040161200f91815260200190565b602060405180830381865afa15801561202c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120509190614fda565b828281518110612062576120626150ab565b6001600160a01b039092166020928302919091019091015280612084816150f0565b915050611f9f565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146120d45760405162461bcd60e51b815260040161060390615411565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d906121229085908590600401615489565b600060405180830381600087803b15801561213c57600080fd5b505af1158015612150573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146121a05760405162461bcd60e51b815260040161060390615411565b6121a981613730565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b15801561220d57600080fd5b505af1158015612221573d6000803e3d6000fd5b5050505050565b612230613684565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906121f39084906004016154d4565b603354610100900460ff161580801561229c5750603354600160ff909116105b806122b65750303b1580156122b6575060335460ff166001145b6123195760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610603565b6033805460ff19166001179055801561233c576033805461ff0019166101001790555b612347846000613a2c565b61235083613b12565b6000805160206158d683398151915280546001600160a01b0319166001600160a01b03841617905580156123be576033805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6000805160206158d6833981519152546000805160206158f6833981519152906001600160a01b031633146124495760405162461bcd60e51b815260206004820152602560248201527f4f6e6c79205461736b4d616e616765722063616e20656a6563742076616c696460448201526461746f727360d81b6064820152608401610603565b61245282613b42565b5050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612491576124916150ab565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906124cd90889086906004016153f2565b600060405180830381865afa1580156124ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261251291908101906152ac565b600081518110612524576125246150ab565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612590573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125b49190615109565b6001600160c01b0316905060006125ca82613479565b9050816125d88a838a610d42565b9550955050505050935093915050565b6040805160608082018352808252600060208301529181019190915261260d8261353b565b92915050565b6000808260405160200161262791906154e7565b60405160208183030381529060405280519060200120905060006126566000805160206158f683398151915290565b9050600081600001600084815260200190815260200160002060405180606001604052908160008201805461268a90615503565b80601f01602080910402602001604051908101604052809291908181526020018280546126b690615503565b80156127035780601f106126d857610100808354040283529160200191612703565b820191906000526020600020905b8154815290600101906020018083116126e657829003601f168201915b50505050508152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561276f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612751575b50505050508152505090506000805b826040015151811015611da0576127b3836040015182815181106127a4576127a46150ab565b602002602001015160006127d1565b6127bd908361553d565b9150806127c9816150f0565b91505061277e565b604051631619718360e21b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000821691635401ed27917f000000000000000000000000000000000000000000000000000000000000000090911690635865c60c906024016040805180830381865afa158015612865573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612889919061555d565b5160405160e083901b6001600160e01b0319168152600481019190915260ff85166024820152604401602060405180830381865afa1580156128cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f39190615248565b9392505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561295c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129809190615132565b60ff169050806000036129a157505060408051600081526020810190915290565b6000805b82811015612a5657604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015612a14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a3891906150c1565b612a42908361514f565b915080612a4e816150f0565b9150506129a5565b506000816001600160401b03811115612a7157612a71614499565b604051908082528060200260200182016040528015612a9a578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612aff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b239190615132565b60ff16811015612cbc57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015612b97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bbb91906150c1565b905060005b81811015612ca7576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015612c35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c599190615177565b60000151858581518110612c6f57612c6f6150ab565b6001600160a01b039092166020928302919091019091015283612c91816150f0565b9450508080612c9f906150f0565b915050612bc0565b50508080612cb4906150f0565b915050612aa1565b5090949350505050565b612cce613684565b6001600160a01b038116612d335760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610603565b610615816136de565b6000805160206158d6833981519152546040805163f5c9899d60e01b815290516000926000805160206158f6833981519152926001600160a01b039091169163f5c9899d916004818101926020929091908290030181865afa158015611e28573d6000803e3d6000fd5b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612dee5760405162461bcd60e51b815260040161060390615411565b612df782613c50565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90612e459086908590600401615489565b600060405180830381600087803b158015612e5f57600080fd5b505af1158015612e73573d6000803e3d6000fd5b50505050600082604051602001612e8a91906154e7565b6040516020818303038152906040528051906020012090506123be8482613c9b565b600160009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612eff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f239190614fda565b6001600160a01b0316336001600160a01b031614612f535760405162461bcd60e51b815260040161060390614ff7565b600254198119600254191614612fd15760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610603565b600281905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161074c565b6000613012611e52565b60405162cf2ab560e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062cf2ab5906121f390849060040161462b565b60005b818110156133335782828281811061307c5761307c6150ab565b905060200281019061308e919061558d565b61309f906040810190602001614463565b6001600160a01b03166323b872dd33308686868181106130c1576130c16150ab565b90506020028101906130d3919061558d565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af115801561312a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061314e9190615041565b506000838383818110613163576131636150ab565b9050602002810190613175919061558d565b613186906040810190602001614463565b604051636eb1769f60e11b81523060048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166024830152919091169063dd62ed3e90604401602060405180830381865afa1580156131f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061321891906150c1565b905083838381811061322c5761322c6150ab565b905060200281019061323e919061558d565b61324f906040810190602001614463565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000083878787818110613291576132916150ab565b90506020028101906132a3919061558d565b604001356132b1919061514f565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156132fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133209190615041565b50508061332c906150f0565b9050613062565b5060405163fce36c7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063fce36c7d906121229085908590600401615609565b6001600160a01b0381166134105760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610603565b600154604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b606060008061348784613ec9565b61ffff166001600160401b038111156134a2576134a2614499565b6040519080825280601f01601f1916602001820160405280156134cc576020820181803683370190505b5090506000805b8251821080156134e4575061010081105b15612cbc576001811b93508584161561352b578060f81b83838151811061350d5761350d6150ab565b60200101906001600160f81b031916908160001a9053508160010191505b613534816150f0565b90506134d3565b6040805160608082018352808252600060208084018290528385018390528582526000805160206158f6833981519152908190529084902084519283019094528354929390928290829061358e90615503565b80601f01602080910402602001604051908101604052809291908181526020018280546135ba90615503565b80156136075780601f106135dc57610100808354040283529160200191613607565b820191906000526020600020905b8154815290600101906020018083116135ea57829003601f168201915b50505050508152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561367357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613655575b505050505081525050915050919050565b6066546001600160a01b03163314611dbc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610603565b606680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03811661377e5760405162461bcd60e51b81526020600482015260156024820152747a65726f206f70657261746f72206164647265737360581b6044820152606401610603565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0160205260409020546000805160206158f6833981519152908061380c5760405162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840ecc2d8d2c8c2e8dee440d0c2e6d60531b6044820152606401610603565b6001600160a01b0383166000908152600280840160209081526040808420548585529186905290922090810154821061387d5760405162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f6620626f756e647360681b6044820152606401610603565b846001600160a01b031681600201838154811061389c5761389c6150ab565b6000918252602090912001546001600160a01b0316146138fe5760405162461bcd60e51b815260206004820152601960248201527f4f70657261746f722061646472657373206d69736d61746368000000000000006044820152606401610603565b600281015460009061391290600190615717565b9050600082600201828154811061392b5761392b6150ab565b6000918252602090912001546001600160a01b0316905083821461399b578083600201858154811061395f5761395f6150ab565b600091825260208083209190910180546001600160a01b0319166001600160a01b03948516179055918316815260028801909152604090208490555b826002018054806139ae576139ae61572a565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038916808352600189018252604080842084905560028a01909252818320839055905187927f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a350505050505050565b6001546001600160a01b0316158015613a4d57506001600160a01b03821615155b613acf5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610603565b600281905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261245282613382565b603354610100900460ff16613b395760405162461bcd60e51b815260040161060390615740565b61061581613ef4565b600081516001600160401b03811115613b5d57613b5d614499565b604051908082528060200260200182016040528015613b86578160200160208202803683370190505b50905060005b8251811015613c14576000838281518110613ba957613ba96150ab565b6020026020010151604051602001613bc191906154e7565b60405160208183030381529060405280519060200120905080838381518110613bec57613bec6150ab565b602002602001018181525050613c0181613f1b565b5080613c0c816150f0565b915050613b8c565b507fb926bfb67574e846e1a181d79d54df4604ec6583a3ecef28f073ac6a1667550a81604051613c449190614618565b60405180910390a15050565b600081604051602001613c6391906154e7565b604051602081830303815290604052805190602001209050613c848161401b565b151560000361245257613c9682614030565b505050565b6001600160a01b038216613ce95760405162461bcd60e51b81526020600482015260156024820152747a65726f206f70657261746f72206164647265737360581b6044820152606401610603565b80613d365760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642076616c696461746f722061646472657373000000000000006044820152606401610603565b613d3f8161401b565b613d845760405162461bcd60e51b81526020600482015260166024820152750eadcd6dcdeeedc40ecc2d8d2c8c2e8dee440d0c2e6d60531b6044820152606401610603565b6000613d918360006127d1565b90506000816001600160601b031611613ddc5760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964207374616b6560981b6044820152606401610603565b6001600160a01b03831660008181527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f01602090815260408083208690558583526000805160206158f683398151915280835281842060020180548686527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0285528386208190558285526001810182559085529383902090930180546001600160a01b0319168517905580516001600160601b038616815290519293869390927f5ea87e0bfaffd2b7f2901f087737e2b175f237068c6123ca9ea5240076bb5668928290030190a350505050565b6000805b821561260d57613ede600184615717565b9092169180613eec8161578b565b915050613ecd565b603354610100900460ff16612d335760405162461bcd60e51b815260040161060390615740565b60008181526000805160206158f683398151915260208190526040909120600181015415613c965760005b6002820154811015613fe957613f84826002018281548110613f6a57613f6a6150ab565b6000918252602090912001546001600160a01b03166141a9565b83826002018281548110613f9a57613f9a6150ab565b60009182526020822001546040516001600160a01b03909116917f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a380613fe1816150f0565b915050613f46565b50600083815260208390526040812090614003828261437c565b600182016000905560028201600061222191906143b6565b60006140268261353b565b5151151592915050565b6040805160608082018352808252600060208301529181019190915260008260405160200161405f91906154e7565b6040516020818303038152906040528051906020012090506140808161401b565b156140cd5760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f7220616c72656164792072656769737465726564000000006044820152606401610603565b6040805160608101825284815260208082018490528251600080825291810184529092820152905060006000805160206158f6833981519152600084815260208290526040902083519192508391819061412790826157f2565b5060208281015160018301556040830151805161414a92600285019201906143d4565b505050600381018054600181018255600091825260209091200183905560405183907f1068d9ac50f7bcef2033b28273aaa313038dae4dd6993ca260688c805149bea3906141999088906154d4565b60405180910390a2509392505050565b6001600160a01b03811660009081527fd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f0160205260409020546000805160206158f6833981519152908015613c96576040805160018082528183019092526000916020808301908036833701905050905060008160008151811061422e5761422e6150ab565b602002602001019060ff16908160ff16815250507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e3b17db8561427b846142d0565b6040518363ffffffff1660e01b81526004016142989291906158b1565b600060405180830381600087803b1580156142b257600080fd5b505af11580156142c6573d6000803e3d6000fd5b5050505050505050565b6060600082516001600160401b038111156142ed576142ed614499565b6040519080825280601f01601f191660200182016040528015614317576020820181803683370190505b50905060005b835181101561086c57838181518110614338576143386150ab565b602002602001015160f81b828281518110614355576143556150ab565b60200101906001600160f81b031916908160001a905350614375816150f0565b905061431d565b50805461438890615503565b6000825580601f10614398575050565b601f0160209004906000526020600020908101906106159190614439565b50805460008255906000526020600020908101906106159190614439565b828054828255906000526020600020908101928215614429579160200282015b8281111561442957825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906143f4565b50614435929150614439565b5090565b5b80821115614435576000815560010161443a565b6001600160a01b038116811461061557600080fd5b60006020828403121561447557600080fd5b81356128f38161444e565b60006020828403121561449257600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156144d1576144d1614499565b60405290565b604051601f8201601f191681016001600160401b03811182821017156144ff576144ff614499565b604052919050565b60006001600160401b0382111561452057614520614499565b5060051b60200190565b6000806040838503121561453d57600080fd5b82356145488161444e565b91506020838101356001600160401b0381111561456457600080fd5b8401601f8101861361457557600080fd5b803561458861458382614507565b6144d7565b81815260059190911b820183019083810190888311156145a757600080fd5b928401925b828410156145ce5783356145bf8161444e565b825292840192908401906145ac565b80955050505050509250929050565b600081518084526020808501945080840160005b8381101561460d578151875295820195908201906001016145f1565b509495945050505050565b6020815260006128f360208301846145dd565b6020808252825182820181905260009190848201906040850190845b8181101561466c5783516001600160a01b031683529284019291840191600101614647565b50909695505050505050565b600082601f83011261468957600080fd5b81356001600160401b038111156146a2576146a2614499565b6146b5601f8201601f19166020016144d7565b8181528460208386010111156146ca57600080fd5b816020850160208301376000918101602001919091529392505050565b63ffffffff8116811461061557600080fd5b8035614704816146e7565b919050565b60008060006060848603121561471e57600080fd5b83356147298161444e565b925060208401356001600160401b0381111561474457600080fd5b61475086828701614678565b9250506040840135614761816146e7565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614802578385038a52825180518087529087019087870190845b818110156147ed57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016147a9565b50509a87019a9550509185019160010161478b565b509298975050505050505050565b6020815260006128f3602083018461476c565b60005b8381101561483e578181015183820152602001614826565b50506000910152565b6000815180845261485f816020860160208601614823565b601f01601f19169290920160200192915050565b60008151606084526148886060850182614847565b9050602080840151818601526040840151858303604087015282815180855283850191508383019450600092505b808310156148df5784516001600160a01b031682529383019360019290920191908301906148b6565b509695505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561493f57603f1988860301845261492d858351614873565b94509285019290850190600101614911565b5092979650505050505050565b600082601f83011261495d57600080fd5b8135602061496d61458383614507565b82815260059290921b8401810191818101908684111561498c57600080fd5b8286015b848110156148df5780358352918301918301614990565b600080604083850312156149ba57600080fd5b82356149c58161444e565b915060208301356001600160401b038111156149e057600080fd5b6149ec8582860161494c565b9150509250929050565b60008083601f840112614a0857600080fd5b5081356001600160401b03811115614a1f57600080fd5b6020830191508360208260051b8501011115614a3a57600080fd5b9250929050565b60008060008060008060808789031215614a5a57600080fd5b8635614a658161444e565b95506020870135614a75816146e7565b945060408701356001600160401b0380821115614a9157600080fd5b818901915089601f830112614aa557600080fd5b813581811115614ab457600080fd5b8a6020828501011115614ac657600080fd5b602083019650809550506060890135915080821115614ae457600080fd5b50614af189828a016149f6565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561460d57815163ffffffff1687529582019590820190600101614b17565b600060208083528351608082850152614b5560a0850182614b03565b905081850151601f1980868403016040870152614b728383614b03565b92506040870151915080868403016060870152614b8f8383614b03565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614be65784878303018452614bd4828751614b03565b95880195938801939150600101614bba565b509998505050505050505050565b60ff8116811461061557600080fd5b600060208284031215614c1557600080fd5b81356128f381614bf4565b600080600060608486031215614c3557600080fd5b8335614c408161444e565b925060208401356001600160401b03811115614c5b57600080fd5b6147508682870161494c565b6020808252825182820181905260009190848201906040850190845b8181101561466c57835183529284019291840191600101614c83565b600060608284031215614cb157600080fd5b604051606081016001600160401b038282108183111715614cd457614cd4614499565b816040528293508435915080821115614cec57600080fd5b50614cf985828601614678565b82525060208301356020820152604083013560408201525092915050565b60008060408385031215614d2a57600080fd5b8235614d358161444e565b915060208301356001600160401b03811115614d5057600080fd5b6149ec85828601614c9f565b600060208284031215614d6e57600080fd5b81356001600160401b03811115614d8457600080fd5b614d9084828501614678565b949350505050565b600080600060608486031215614dad57600080fd5b8335614db88161444e565b92506020840135614dc88161444e565b915060408401356147618161444e565b60006020808385031215614deb57600080fd5b82356001600160401b0380821115614e0257600080fd5b818501915085601f830112614e1657600080fd5b8135614e2461458382614507565b81815260059190911b83018401908481019088831115614e4357600080fd5b8585015b83811015614e7b57803585811115614e5f5760008081fd5b614e6d8b89838a0101614678565b845250918601918601614e47565b5098975050505050505050565b600080600060608486031215614e9d57600080fd5b8335614ea88161444e565b9250602084013591506040840135614761816146e7565b828152604060208201526000614d90604083018461476c565b6020815260006128f36020830184614873565b60008060408385031215614efe57600080fd5b8235614f098161444e565b91506020830135614f1981614bf4565b809150509250929050565b600080600060608486031215614f3957600080fd5b8335614f448161444e565b925060208401356001600160401b0380821115614f6057600080fd5b614f6c87838801614678565b93506040860135915080821115614f8257600080fd5b50614f8f86828701614c9f565b9150509250925092565b60008060208385031215614fac57600080fd5b82356001600160401b03811115614fc257600080fd5b614fce858286016149f6565b90969095509350505050565b600060208284031215614fec57600080fd5b81516128f38161444e565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561505357600080fd5b815180151581146128f357600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156150d357600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600060018201615102576151026150da565b5060010190565b60006020828403121561511b57600080fd5b81516001600160c01b03811681146128f357600080fd5b60006020828403121561514457600080fd5b81516128f381614bf4565b8082018082111561260d5761260d6150da565b6001600160601b038116811461061557600080fd5b60006040828403121561518957600080fd5b6151916144af565b825161519c8161444e565b815260208301516151ac81615162565b60208201529392505050565b600060208083850312156151cb57600080fd5b82516001600160401b038111156151e157600080fd5b8301601f810185136151f257600080fd5b805161520061458382614507565b81815260059190911b8201830190838101908783111561521f57600080fd5b928401925b8284101561523d57835182529284019290840190615224565b979650505050505050565b60006020828403121561525a57600080fd5b81516128f381615162565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561529257600080fd5b8260051b8085606085013791909101606001949350505050565b600060208083850312156152bf57600080fd5b82516001600160401b038111156152d557600080fd5b8301601f810185136152e657600080fd5b80516152f461458382614507565b81815260059190911b8201830190838101908783111561531357600080fd5b928401925b8284101561523d57835161532b816146e7565b82529284019290840190615318565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff8416815260406020820152600061538360408301848661533a565b95945050505050565b60006020828403121561539e57600080fd5b81516128f3816146e7565b600060ff821660ff81036153bf576153bf6150da565b60010192915050565b6040815260006153dc60408301858761533a565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201526000614d9060408301846145dd565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526154b360a0840182614847565b90506020840151606084015260408401516080840152809150509392505050565b6020815260006128f36020830184614847565b600082516154f9818460208701614823565b9190910192915050565b600181811c9082168061551757607f821691505b60208210810361553757634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160601b0381811683821601908082111561086c5761086c6150da565b60006040828403121561556f57600080fd5b6155776144af565b825181526020830151600381106151ac57600080fd5b60008235609e198336030181126154f957600080fd5b80356147048161444e565b8183526000602080850194508260005b8581101561460d5781356155d18161444e565b6001600160a01b03168752818301356155e981615162565b6001600160601b03168784015260409687019691909101906001016155be565b60208082528181018390526000906040808401600586901b8501820187855b8881101561570957878303603f190184528135368b9003609e1901811261564e57600080fd5b8a0160a0813536839003601e1901811261566757600080fd5b820188810190356001600160401b0381111561568257600080fd5b8060061b360382131561569457600080fd5b8287526156a483880182846155ae565b925050506156b38883016155a3565b6001600160a01b031688860152818701358786015260606156d58184016146f9565b63ffffffff169086015260806156ec8382016146f9565b63ffffffff16950194909452509285019290850190600101615628565b509098975050505050505050565b8181038181111561260d5761260d6150da565b634e487b7160e01b600052603160045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600061ffff8083168181036157a2576157a26150da565b6001019392505050565b601f821115613c9657600081815260208120601f850160051c810160208610156157d35750805b601f850160051c820191505b81811015612150578281556001016157df565b81516001600160401b0381111561580b5761580b614499565b61581f816158198454615503565b846157ac565b602080601f831160018114615854576000841561583c5750858301515b600019600386901b1c1916600185901b178555612150565b600085815260208120601f198616915b8281101561588357888601518255948401946001909101908401615864565b50858210156158a15787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b0383168152604060208201819052600090614d909083018461484756fed4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f04d4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00a2646970667358221220d35e75cc59a45f6bd069b571eca63d4bb266ab27082faf8a3bc39370eddbbfca64736f6c63430008150033",
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
