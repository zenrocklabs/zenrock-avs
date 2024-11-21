// SPDX-License-Identifier: BUSL
pragma solidity 0.8.21;

import {BN254} from "../lib/eigenlayer-middleware/src/libraries/BN254.sol";
import {IRegistryCoordinator} from "../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BLSSignatureChecker, IBLSSignatureChecker} from "../lib/eigenlayer-middleware/src/BLSSignatureChecker.sol";
import "./libraries/ZrServiceManagerLib.sol";
import {Initializable} from "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "../lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import "./interfaces/IZrServiceManager.sol";

/**
 * @title ZrTaskManager
 * @notice Manages task creation, responses, and challenges for the ZR protocol
 * @dev Implements BLS signature checking and task lifecycle management
 */
contract ZrTaskManager is Initializable, OwnableUpgradeable, BLSSignatureChecker {
    using BN254 for BN254.G1Point;

    // Constants
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 10000;
    uint256 private constant THRESHOLD_DENOMINATOR = 67;
    bytes32 private constant TASK_MANAGER_STORAGE_LOCATION = 0x9e5d0bf83ef884a66a66b2d439fd65f5546f8f4489c6a744f987ecb90e5d7100;

    // Custom errors for gas efficiency
    error OnlyAggregator();
    error OnlyTaskGenerator();
    error TaskAlreadyResponded();
    error TaskResponseTooLate();
    error TaskNotResponded();
    error TaskResponseMismatch();
    error TaskAlreadyChallenged();
    error ChallengePeriodExpired();
    error TaskIdMismatch();
    error ThresholdNotMet();
    error InvalidNonSigners();
    error InvalidTask();
    error ZeroAddress();
    error InvalidQuorumNumbers();

    // Events moved to top for better organization
    event NewTaskCreated(uint32 indexed taskId, ZrServiceManagerLib.Task task);
    event TaskResponded(ZrServiceManagerLib.TaskResponse response, ZrServiceManagerLib.TaskResponseMetadata metadata);
    event TaskChallengedSuccessfully(uint32 indexed taskId, address indexed challenger);

    // Storage structure
    struct TaskManagerStorage {
        uint32 taskResponseWindowBlock;
        uint32 latestTaskId;
        mapping(uint32 => bytes32) allTaskHashes;
        mapping(uint32 => bytes32) allTaskResponses;
        mapping(uint32 => bool) taskSuccessfullyChallenged;
        address aggregator;
        address generator;
        IZrServiceManager zrServiceManager;
        
        // New fields for optimization
        mapping(uint32 => uint32) taskCreationBlocks;
        mapping(bytes32 => bool) usedResponseHashes;
    }

    /**
     * @dev Returns storage pointer to TaskManagerStorage
     */
    function _getTaskManagerStorage() private pure returns (TaskManagerStorage storage $) {
        assembly {
            $.slot := TASK_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(
        IRegistryCoordinator _registryCoordinator
    ) BLSSignatureChecker(_registryCoordinator) {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract
     */
    function initialize(
        address _aggregator,
        address _generator,
        uint32 _taskResponseWindowBlock,
        IZrServiceManager _zrServiceManager
    ) external initializer {
        if (_aggregator == address(0)) revert ZeroAddress();
        if (_generator == address(0)) revert ZeroAddress();
        if (address(_zrServiceManager) == address(0)) revert ZeroAddress();

        __Ownable_init();
        
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        $.aggregator = _aggregator;
        $.generator = _generator;
        $.taskResponseWindowBlock = _taskResponseWindowBlock;
        $.zrServiceManager = _zrServiceManager;
    }

    /**
     * @dev Modifier that checks if sender is the aggregator
     */
    modifier onlyAggregator() {
        if (msg.sender != _getTaskManagerStorage().aggregator) revert OnlyAggregator();
        _;
    }

    /**
     * @dev Modifier that checks if sender is the task generator
     */
    modifier onlyTaskGenerator() {
        if (msg.sender != _getTaskManagerStorage().generator) revert OnlyTaskGenerator();
        _;
    }

    /**
     * @notice Creates a new task
     * @param taskId Task identifier
     * @param quorumThresholdPercentage Required percentage of quorum signatures
     * @param quorumNumbers Array of quorum numbers
     */
    function createNewTask(
        uint32 taskId,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyTaskGenerator {
        if (quorumNumbers.length == 0) revert InvalidQuorumNumbers();
        
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        
        // Create task struct inline for gas efficiency
        bytes32 taskHash = keccak256(abi.encode(
            ZrServiceManagerLib.Task({
                taskId: taskId,
                taskCreatedBlock: uint32(block.number),
                quorumThresholdPercentage: quorumThresholdPercentage,
                quorumNumbers: quorumNumbers
            })
        ));
        
        $.allTaskHashes[taskId] = taskHash;
        $.taskCreationBlocks[taskId] = uint32(block.number);
        $.latestTaskId = taskId;

        emit NewTaskCreated(
            taskId,
            ZrServiceManagerLib.Task({
                taskId: taskId,
                taskCreatedBlock: uint32(block.number),
                quorumThresholdPercentage: quorumThresholdPercentage,
                quorumNumbers: quorumNumbers
            })
        );
    }

    /**
     * @notice Responds to a task with signed data
     * @param task Task details
     * @param taskResponse Response data
     * @param nonSignerStakesAndSignature Non-signer stakes and signature data
     */
    function respondToTask(
        ZrServiceManagerLib.Task calldata task,
        ZrServiceManagerLib.TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        TaskManagerStorage storage $ = _getTaskManagerStorage();

        // Validate task and timing
        if (keccak256(abi.encode(task)) != $.allTaskHashes[task.taskId]) revert InvalidTask();
        if ($.allTaskResponses[task.taskId] != bytes32(0)) revert TaskAlreadyResponded();
        if (uint32(block.number) > task.taskCreatedBlock + $.taskResponseWindowBlock) revert TaskResponseTooLate();

        // Check signatures and stakes
        bytes32 message = keccak256(abi.encode(taskResponse));
        (QuorumStakeTotals memory quorumStakeTotals, bytes32 hashOfNonSigners) = 
            checkSignatures(
                message,
                task.quorumNumbers,
                task.taskCreatedBlock,
                nonSignerStakesAndSignature
            );

        // Verify threshold requirements
        for (uint256 i = 0; i < task.quorumNumbers.length;) {
            if (
                quorumStakeTotals.signedStakeForQuorum[i] * THRESHOLD_DENOMINATOR <
                quorumStakeTotals.totalStakeForQuorum[i] * uint8(task.quorumThresholdPercentage)
            ) revert ThresholdNotMet();
            unchecked { ++i; }
        }

        // Handle inactive validators
        if (taskResponse.inactiveSetZRChain.length > 0) {
            $.zrServiceManager.ejectValidators(taskResponse.inactiveSetZRChain);
        }

        // Store response
        ZrServiceManagerLib.TaskResponseMetadata memory metadata = 
            ZrServiceManagerLib.TaskResponseMetadata({
                taskResponsedBlock: uint32(block.number),
                hashOfNonSigners: hashOfNonSigners
            });
            
        bytes32 responseHash = keccak256(abi.encode(taskResponse, metadata));
        $.allTaskResponses[task.taskId] = responseHash;
        $.usedResponseHashes[responseHash] = true;

        emit TaskResponded(taskResponse, metadata);
    }

    /**
     * @notice Raises and resolves a challenge for a task response
     * @param task Task details
     * @param taskResponse Response data
     * @param taskResponseMetadata Response metadata
     * @param pubkeysOfNonSigningOperators Public keys of non-signing operators
     */
    function raiseAndResolveChallenge(
        ZrServiceManagerLib.Task calldata task,
        ZrServiceManagerLib.TaskResponse calldata taskResponse,
        ZrServiceManagerLib.TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        uint32 referenceTaskId = taskResponse.referenceTaskId;

        // Validate challenge conditions
        if ($.allTaskResponses[referenceTaskId] == bytes32(0)) revert TaskNotResponded();
        
        bytes32 responseHash = keccak256(abi.encode(taskResponse, taskResponseMetadata));
        if ($.allTaskResponses[referenceTaskId] != responseHash) revert TaskResponseMismatch();
        if ($.taskSuccessfullyChallenged[referenceTaskId]) revert TaskAlreadyChallenged();
        if (uint32(block.number) > taskResponseMetadata.taskResponsedBlock + TASK_CHALLENGE_WINDOW_BLOCK) {
            revert ChallengePeriodExpired();
        }
        if (task.taskId != taskResponse.referenceTaskId) revert TaskIdMismatch();

        // Verify non-signing operators
        bytes32[] memory hashesOfPubkeys = new bytes32[](pubkeysOfNonSigningOperators.length);
        for (uint256 i = 0; i < pubkeysOfNonSigningOperators.length;) {
            hashesOfPubkeys[i] = pubkeysOfNonSigningOperators[i].hashG1Point();
            unchecked { ++i; }
        }

        bytes32 signatoryRecordHash = keccak256(
            abi.encodePacked(task.taskCreatedBlock, hashesOfPubkeys)
        );
        
        if (signatoryRecordHash != taskResponseMetadata.hashOfNonSigners) revert InvalidNonSigners();

        // Mark task as successfully challenged
        $.taskSuccessfullyChallenged[referenceTaskId] = true;

        emit TaskChallengedSuccessfully(referenceTaskId, msg.sender);
    }

    // View functions

    function getTaskNumber() external view returns (uint32) {
        return _getTaskManagerStorage().latestTaskId;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return _getTaskManagerStorage().taskResponseWindowBlock;
    }

    function getTaskResponse(uint32 taskId) external view returns (bytes32) {
        return _getTaskManagerStorage().allTaskResponses[taskId];
    }

    function getTaskCreationBlock(uint32 taskId) external view returns (uint32) {
        return _getTaskManagerStorage().taskCreationBlocks[taskId];
    }
}