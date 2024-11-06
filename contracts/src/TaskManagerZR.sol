pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./ITaskManagerZR.sol";

contract TaskManagerZR is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ITaskManagerZR
{
    using BN254 for BN254.G1Point;

    /* CONSTANTS */
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 10000;
    uint256 private constant _THRESHOLD_DENOMINATOR = 67;

    /* STORAGE */
    uint32 public latestTaskId;
    address public aggregator;
    address public generator;

    struct TaskData {
        bytes32 taskHash;
        bytes32 taskResponse;
        bool challenged;
        mapping(uint256 => string) validatorAddresses;
        uint256 validatorCount;
    }

    mapping(uint32 => TaskData) private taskData;

    /* EVENTS */
    event ValidatorAddressesStored(uint32 indexed taskId, uint256 count);

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Not aggregator");
        _;
    }

    modifier onlyTaskGenerator() {
        require(msg.sender == generator, "Not generator");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock
    ) BLSSignatureChecker(_registryCoordinator) {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator,
        address _generator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
        generator = _generator;
    }

    function createNewTask(
        uint32 taskId,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyTaskGenerator {
        Task memory newTask = Task({
            taskId: taskId,
            taskCreatedBlock: uint32(block.number),
            quorumThresholdPercentage: quorumThresholdPercentage,
            quorumNumbers: quorumNumbers
        });

        taskData[taskId].taskHash = keccak256(abi.encode(newTask));
        emit NewTaskCreated(taskId, newTask);
        latestTaskId = taskId;
    }

    function _storeValidatorAddresses(uint32 taskId, string[] calldata addresses) private {
        uint256 length = addresses.length;
        TaskData storage data = taskData[taskId];
        for(uint256 i = 0; i < length; i++) {
            data.validatorAddresses[i] = addresses[i];
        }
        data.validatorCount = length;
        emit ValidatorAddressesStored(taskId, length);
    }

    function _getValidatorAddresses(uint32 taskId) private view returns (string[] memory) {
        TaskData storage data = taskData[taskId];
        string[] memory addresses = new string[](data.validatorCount);
        for(uint256 i = 0; i < data.validatorCount; i++) {
            addresses[i] = data.validatorAddresses[i];
        }
        return addresses;
    }

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        TaskData storage data = taskData[task.taskId];
        
        require(data.taskHash == keccak256(abi.encode(task)), "Invalid task");
        require(data.taskResponse == bytes32(0), "Already responded");
        require(
            block.number <= task.taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "Too late"
        );

        bytes32 message = keccak256(abi.encode(taskResponse));
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
                message,
                task.quorumNumbers,
                task.taskCreatedBlock,
                nonSignerStakesAndSignature
            );

        for (uint i = 0; i < task.quorumNumbers.length; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * _THRESHOLD_DENOMINATOR >=
                quorumStakeTotals.totalStakeForQuorum[i] * task.quorumThresholdPercentage,
                "Below threshold"
            );
        }

        _storeValidatorAddresses(task.taskId, taskResponse.activeSetZRChain);

        TaskResponseMetadata memory metadata = TaskResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        data.taskResponse = keccak256(abi.encode(taskResponse, metadata));

        emit TaskResponded(taskResponse, metadata);
    }

    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        uint32 referenceTaskId = taskResponse.referenceTaskId;
        TaskData storage data = taskData[referenceTaskId];
        
        require(data.taskResponse != bytes32(0), "Not responded");
        require(
            data.taskResponse == keccak256(abi.encode(taskResponse, taskResponseMetadata)),
            "Invalid response"
        );
        require(!data.challenged, "Already challenged");
        require(
            block.number <= taskResponseMetadata.taskResponsedBlock + TASK_CHALLENGE_WINDOW_BLOCK,
            "Challenge expired"
        );

        bytes32[] memory hashes = new bytes32[](pubkeysOfNonSigningOperators.length);
        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            hashes[i] = pubkeysOfNonSigningOperators[i].hashG1Point();
        }

        require(
            keccak256(abi.encodePacked(task.taskCreatedBlock, hashes)) == 
            taskResponseMetadata.hashOfNonSigners,
            "Invalid non-signers"
        );

        data.challenged = true;
        emit TaskChallengedSuccessfully(referenceTaskId, msg.sender);
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskId;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }

    function getActiveSet(uint32 taskId) external view returns (string[] memory) {
        return _getValidatorAddresses(taskId);
    }

    function getLatestActiveSet() external view returns (string[] memory) {
        return _getValidatorAddresses(latestTaskId);
    }

    function isValidatorInTaskSet(uint32 taskId, string memory validatorAddress) external view returns (bool) {
        TaskData storage data = taskData[taskId];
        for (uint256 i = 0; i < data.validatorCount; i++) {
            if (keccak256(abi.encodePacked(data.validatorAddresses[i])) == 
                keccak256(abi.encodePacked(validatorAddress))) {
                return true;
            }
        }
        return false;
    }
}