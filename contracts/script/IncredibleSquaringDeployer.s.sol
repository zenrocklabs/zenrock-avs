// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import {IRewardsCoordinator} from "eigenlayer-contracts/src/contracts/interfaces/IRewardsCoordinator.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {IBLSApkRegistry} from "@eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {IIndexRegistry} from "@eigenlayer-middleware/src/interfaces/IIndexRegistry.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {ISocketRegistry} from "@eigenlayer-middleware/src/interfaces/ISocketRegistry.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import {SocketRegistry} from "@eigenlayer-middleware/src/SocketRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {ZrServiceManager} from "../src/ZrServiceManager.sol";
import {ZrRegistryCoordinator} from "../src/ZrRegistryCoordinator.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract IncredibleSquaringDeployer is Script, Utils {
    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 67;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 300;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 0;
    
    address public constant AGGREGATOR_ADDR = 0x52CD7Cb69053c1F00360B8917809E601b78498Fc;
    address public constant TASK_GENERATOR_ADDR = 0x52CD7Cb69053c1F00360B8917809E601b78498Fc;
    address public constant REWARDS_INITIATOR = 0x52CD7Cb69053c1F00360B8917809E601b78498Fc;

    // Contract instances
    ProxyAdmin public incredibleSquaringProxyAdmin;
    PauserRegistry public incredibleSquaringPauserReg;

    ZrRegistryCoordinator public registryCoordinator;
    ZrRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    ISocketRegistry public socketRegistry;
    ISocketRegistry public socketRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    ZrServiceManager public incredibleSquaringServiceManager;
    ZrServiceManager public incredibleSquaringServiceManagerImplementation;

    function run() external {
        // Get Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput("eigenlayer_deployment_output");
        
        IStrategyManager strategyManager = IStrategyManager(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.strategyManager")
        );
        IDelegationManager delegationManager = IDelegationManager(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.delegation")
        );
        IAVSDirectory avsDirectory = IAVSDirectory(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.avsDirectory")
        );
        IRewardsCoordinator rewardsCoordinator = IRewardsCoordinator(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.rewardsCoordinator")
        );

        address incredibleSquaringCommunityMultisig = msg.sender;
        address incredibleSquaringPauser = msg.sender;

        vm.startBroadcast();
        
        // Use existing WETH strategy
        StrategyBase wethStrat = StrategyBase(0xFb83e1D133D0157775eC4F19Ff81478Df1103305);
        
        _deployIncredibleSquaringContracts(
            delegationManager,
            avsDirectory,
            rewardsCoordinator,
            wethStrat,
            incredibleSquaringCommunityMultisig,
            incredibleSquaringPauser
        );
        
        vm.stopBroadcast();
    }

    function _deployIncredibleSquaringContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        IRewardsCoordinator rewardsCoordinator,
        IStrategy strat,
        address incredibleSquaringCommunityMultisig,
        address credibleSquaringPauser
    ) internal {
        // Setup single strategy array
        IStrategy[1] memory deployedStrategyArray = [strat];
        uint numStrategies = deployedStrategyArray.length;

        // Deploy proxy admin
        incredibleSquaringProxyAdmin = new ProxyAdmin();

        // Deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = credibleSquaringPauser;
            pausers[1] = incredibleSquaringCommunityMultisig;
            incredibleSquaringPauserReg = new PauserRegistry(
                pausers,
                incredibleSquaringCommunityMultisig
            );
        }

        EmptyContract emptyContract = new EmptyContract();

        // Deploy proxy contracts with empty implementation initially
        incredibleSquaringServiceManager = ZrServiceManager(
            address(new TransparentUpgradeableProxy(
                address(emptyContract),
                address(incredibleSquaringProxyAdmin),
                ""
            ))
        );

        registryCoordinator = ZrRegistryCoordinator(
            address(new TransparentUpgradeableProxy(
                address(emptyContract),
                address(incredibleSquaringProxyAdmin),
                ""
            ))
        );

        blsApkRegistry = IBLSApkRegistry(
            address(new TransparentUpgradeableProxy(
                address(emptyContract),
                address(incredibleSquaringProxyAdmin),
                ""
            ))
        );

        indexRegistry = IIndexRegistry(
            address(new TransparentUpgradeableProxy(
                address(emptyContract),
                address(incredibleSquaringProxyAdmin),
                ""
            ))
        );

        stakeRegistry = IStakeRegistry(
            address(new TransparentUpgradeableProxy(
                address(emptyContract),
                address(incredibleSquaringProxyAdmin),
                ""
            ))
        );

        socketRegistry = ISocketRegistry(
            address(new TransparentUpgradeableProxy(
                address(emptyContract),
                address(incredibleSquaringProxyAdmin),
                ""
            ))
        );

        operatorStateRetriever = new OperatorStateRetriever();

        // Deploy implementations and upgrade proxies
        {
            socketRegistryImplementation = new SocketRegistry(registryCoordinator);
            
            incredibleSquaringProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(socketRegistry))),
                address(socketRegistryImplementation)
            );

            stakeRegistryImplementation = new StakeRegistry(
                registryCoordinator,
                delegationManager
            );

            incredibleSquaringProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation)
            );

            blsApkRegistryImplementation = new BLSApkRegistry(registryCoordinator);

            incredibleSquaringProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))),
                address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry(registryCoordinator);

            incredibleSquaringProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(indexRegistry))),
                address(indexRegistryImplementation)
            );
        }

        // Deploy and initialize RegistryCoordinator
        registryCoordinatorImplementation = new ZrRegistryCoordinator(
            incredibleSquaringServiceManager,
            IStakeRegistry(address(stakeRegistry)),
            IBLSApkRegistry(address(blsApkRegistry)),
            IIndexRegistry(address(indexRegistry)),
            ISocketRegistry(address(socketRegistry))
        );

        {
            uint numQuorums = 1;
            ZrRegistryCoordinator.OperatorSetParam[] memory quorumsOperatorSetParams = new ZrRegistryCoordinator.OperatorSetParam[](numQuorums);
            
            for (uint i = 0; i < numQuorums; i++) {
                quorumsOperatorSetParams[i] = ZrRegistryCoordinator.OperatorSetParam({
                    maxOperatorCount: 75,
                    kickBIPsOfOperatorStake: 15000,
                    kickBIPsOfTotalStake: 100
                });
            }

            uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
            IStakeRegistry.StrategyParams[][] memory quorumsStrategyParams = new IStakeRegistry.StrategyParams[][](numQuorums);
            
            for (uint i = 0; i < numQuorums; i++) {
                quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](numStrategies);
                for (uint j = 0; j < numStrategies; j++) {
                    quorumsStrategyParams[i][j] = IStakeRegistry.StrategyParams({
                        strategy: deployedStrategyArray[j],
                        multiplier: 1 ether
                    });
                }
            }

            incredibleSquaringProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(registryCoordinator))),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    ZrRegistryCoordinator.initialize.selector,
                    incredibleSquaringCommunityMultisig,
                    incredibleSquaringCommunityMultisig,
                    incredibleSquaringCommunityMultisig,
                    incredibleSquaringPauserReg,
                    0,
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
                )
            );
        }

        // Deploy and initialize ServiceManager
        incredibleSquaringServiceManagerImplementation = new ZrServiceManager(
            avsDirectory,
            rewardsCoordinator,
            ZrRegistryCoordinator(address(registryCoordinator)),
            IStakeRegistry(address(stakeRegistry))
        );

        incredibleSquaringProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(incredibleSquaringServiceManager))),
            address(incredibleSquaringServiceManagerImplementation),
            abi.encodeWithSelector(
                ZrServiceManager.initialize.selector,
                incredibleSquaringPauserReg,
                incredibleSquaringCommunityMultisig,
                AGGREGATOR_ADDR,
                TASK_GENERATOR_ADDR,
                REWARDS_INITIATOR,
                TASK_RESPONSE_WINDOW_BLOCK
            )
        );

        // Write deployment output
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";

        vm.serializeAddress(deployed_addresses, "incredibleSquaringServiceManager", address(incredibleSquaringServiceManager));
        vm.serializeAddress(deployed_addresses, "incredibleSquaringServiceManagerImplementation", address(incredibleSquaringServiceManagerImplementation));
        vm.serializeAddress(deployed_addresses, "registryCoordinator", address(registryCoordinator));
        vm.serializeAddress(deployed_addresses, "registryCoordinatorImplementation", address(registryCoordinatorImplementation));
        vm.serializeAddress(deployed_addresses, "stakeRegistry", address(stakeRegistry));
        vm.serializeAddress(deployed_addresses, "stakeRegistryImplementation", address(stakeRegistryImplementation));
        vm.serializeAddress(deployed_addresses, "socketRegistry", address(socketRegistry));
        vm.serializeAddress(deployed_addresses, "socketRegistryImplementation", address(socketRegistryImplementation));
        string memory deployed_addresses_output = vm.serializeAddress(deployed_addresses, "operatorStateRetriever", address(operatorStateRetriever));

        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        writeOutput(finalJson, "incredible_squaring_avs_deployment_output");
    }
}