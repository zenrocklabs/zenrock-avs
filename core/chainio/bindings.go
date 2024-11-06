package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	cstaskmanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/TaskManagerZR"
	csservicemanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/ZRServiceManager"
)

type AvsManagersBindings struct {
	TaskManager    *cstaskmanager.ContractTaskManagerZR
	ServiceManager *csservicemanager.ContractZRServiceManager
	ethClient      eth.Client
	logger         logging.Logger
}

func NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethclient eth.Client, logger logging.Logger) (*AvsManagersBindings, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethclient)
	if err != nil {
		return nil, err
	}
	serviceManagerAddr, err := contractRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	contractServiceManager, err := csservicemanager.NewContractZRServiceManager(serviceManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}

	taskManagerAddr, err := contractServiceManager.TaskManagerZR(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}
	contractTaskManager, err := cstaskmanager.NewContractTaskManagerZR(taskManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IIncredibleSquaringTaskManager contract", "err", err)
		return nil, err
	}

	return &AvsManagersBindings{
		ServiceManager: contractServiceManager,
		TaskManager:    contractTaskManager,
		ethClient:      ethclient,
		logger:         logger,
	}, nil
}

// func (b *AvsManagersBindings) GetErc20Mock(tokenAddr common.Address) (*erc20mock.ContractERC20Mock, error) {
// 	contractErc20Mock, err := erc20mock.NewContractERC20Mock(tokenAddr, b.ethClient)
// 	if err != nil {
// 		b.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
// 		return nil, err
// 	}
// 	return contractErc20Mock, nil
// }
