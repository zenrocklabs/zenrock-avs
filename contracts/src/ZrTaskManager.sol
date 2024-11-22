// SPDX-License-Identifier: BUSL
pragma solidity 0.8.21;

import {BN254} from "../lib/eigenlayer-middleware/src/libraries/BN254.sol";
import {IRegistryCoordinator} from "../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BLSSignatureChecker, IBLSSignatureChecker} from "../lib/eigenlayer-middleware/src/BLSSignatureChecker.sol";
import "./libraries/ZrServiceManagerLib.sol";
import {Initializable} from "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "../lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import "./interfaces/IZrServiceManager.sol";

contract ZrTaskManager is OwnableUpgradeable, BLSSignatureChecker {
    using BN254 for BN254.G1Point;

    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 10000;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 67;

    // Storage structure
    struct TaskManagerStorage {
        uint32 TASK_RESPONSE_WINDOW_BLOCK;
        uint32 latestTaskId;
        mapping(uint32 => bytes32) allTaskHashes;
        mapping(uint32 => bytes32) allTaskResponses;
        mapping(uint32 => bool) taskSuccesfullyChallenged;
        mapping(uint32 => ZrServiceManagerLib.Task) tasks;  // Added: Store actual tasks
        uint32[] taskIds;  // Added: Array to track all task IDs
        address aggregator;
        address generator;
        IZrServiceManager zrServiceManager;
    }

    bytes32 private constant TASK_MANAGER_STORAGE_LOCATION =
        0x9e5d0bf83ef884a66a66b2d439fd65f5546f8f4489c6a744f987ecb90e5d7100;

    constructor(
        address _aggregator,
        address _generator,
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock,
        address initialOwner,
        IZrServiceManager _zrServiceManager
    ) BLSSignatureChecker(_registryCoordinator) {
        // Initialize OwnableUpgradeable
        _transferOwnership(initialOwner);

        TaskManagerStorage storage $ = _getTaskManagerStorage();
        $.aggregator = _aggregator;
        $.generator = _generator;
        $.TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
        $.zrServiceManager = _zrServiceManager;
    }

    function _getTaskManagerStorage() private pure returns (TaskManagerStorage storage $) {
        assembly {
            $.slot := TASK_MANAGER_STORAGE_LOCATION
        }
    }

    modifier onlyAggregator() {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        require(msg.sender == $.aggregator, "Aggregator must be the caller");
        _;
    }

    modifier onlyTaskGenerator() {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        require(msg.sender == $.generator, "Task generator must be the caller");
        _;
    }

    function createNewTask(
        uint32 taskId,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyTaskGenerator {
        ZrServiceManagerLib.Task memory newTask;
        newTask.taskId = taskId;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;
        
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        $.allTaskHashes[taskId] = keccak256(abi.encode(newTask));
        $.tasks[taskId] = newTask;  // Store the actual task
        $.taskIds.push(taskId);     // Track the task ID
        emit NewTaskCreated(taskId, newTask);
        $.latestTaskId = taskId;
    }

    // Added: Get a specific task by ID
    function getTask(uint32 taskId) external view returns (ZrServiceManagerLib.Task memory) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        require($.allTaskHashes[taskId] != bytes32(0), "Task does not exist");
        return $.tasks[taskId];
    }

    // Added: Get all task IDs
    function getAllTaskIds() external view returns (uint32[] memory) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        return $.taskIds;
    }

    // Added: Get task hash by ID
    function getTaskHash(uint32 taskId) external view returns (bytes32) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        return $.allTaskHashes[taskId];
    }

    // Added: Check if task was challenged successfully
    function isTaskChallenged(uint32 taskId) external view returns (bool) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        return $.taskSuccesfullyChallenged[taskId];
    }

    // Added: Get all task details for a specific ID
    function getTaskDetails(uint32 taskId) external view returns (
        ZrServiceManagerLib.Task memory task,
        bytes32 taskHash,
        bytes32 taskResponse,
        bool isChallenged
    ) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        require($.allTaskHashes[taskId] != bytes32(0), "Task does not exist");
        
        return (
            $.tasks[taskId],
            $.allTaskHashes[taskId],
            $.allTaskResponses[taskId],
            $.taskSuccesfullyChallenged[taskId]
        );
    }

    // Helper function to convert bytes32 to hex string
    function toHexString(bytes32 value) internal pure returns (string memory) {
        bytes memory buffer = new bytes(64);
        bytes16 HEX_DIGITS = "0123456789abcdef";
        
        for(uint256 i = 0; i < 32; i++) {
            buffer[i*2] = HEX_DIGITS[uint8(value[i] >> 4)];
            buffer[i*2+1] = HEX_DIGITS[uint8(value[i] & 0x0f)];
        }
        
        return string(buffer);
    }

    function respondToTask(
        ZrServiceManagerLib.Task calldata task,
        ZrServiceManagerLib.TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;
        TaskManagerStorage storage $ = _getTaskManagerStorage();

        _validateTaskHash($, task);

        require(
            $.allTaskResponses[task.taskId] == bytes32(0),
            "Aggregator has already responded to the task"
        );
        require(
            uint32(block.number) <= taskCreatedBlock + $.TASK_RESPONSE_WINDOW_BLOCK,
            "Aggregator has responded to the task too late"
        );

        bytes32 message = keccak256(abi.encode(taskResponse));

        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
                message,
                quorumNumbers,
                taskCreatedBlock,
                nonSignerStakesAndSignature
            );

        for (uint i = 0; i < quorumNumbers.length; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * _THRESHOLD_DENOMINATOR >=
                    quorumStakeTotals.totalStakeForQuorum[i] * uint8(quorumThresholdPercentage),
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }

        if (taskResponse.inactiveSetZRChain.length > 0) {
            $.zrServiceManager.ejectValidators(taskResponse.inactiveSetZRChain);
        }

        ZrServiceManagerLib.TaskResponseMetadata memory taskResponseMetadata = ZrServiceManagerLib
            .TaskResponseMetadata(uint32(block.number), hashOfNonSigners);
        $.allTaskResponses[task.taskId] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        emit TaskResponded(taskResponse, taskResponseMetadata);
    }

    function _validateTaskHash(TaskManagerStorage storage $, ZrServiceManagerLib.Task calldata task) internal view {
        bytes32 receivedHash = keccak256(abi.encode(task));
        bytes32 expectedHash = $.allTaskHashes[task.taskId];
        require(
            receivedHash == expectedHash,
            string(abi.encodePacked(
                "Task hash mismatch. Expected: ",
                toHexString(expectedHash),
                ", Received: ",
                toHexString(receivedHash)
            ))
        );
    }

    function raiseAndResolveChallenge(
        ZrServiceManagerLib.Task calldata task,
        ZrServiceManagerLib.TaskResponse calldata taskResponse,
        ZrServiceManagerLib.TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        uint32 referenceTaskId = taskResponse.referenceTaskId;
        TaskManagerStorage storage $ = _getTaskManagerStorage();

        require(
            $.allTaskResponses[referenceTaskId] != bytes32(0),
            "Task hasn't been responded to yet"
        );
        require(
            $.allTaskResponses[referenceTaskId] ==
                keccak256(abi.encode(taskResponse, taskResponseMetadata)),
            "Task response does not match the one recorded in the contract"
        );
        require(
            !$.taskSuccesfullyChallenged[referenceTaskId],
            "The response to this task has already been challenged successfully."
        );
        require(
            uint32(block.number) <=
                taskResponseMetadata.taskResponsedBlock +
                    TASK_CHALLENGE_WINDOW_BLOCK,
            "The challenge period for this task has already expired."
        );

        require(
            task.taskId == taskResponse.referenceTaskId,
            "Task ID mismatch"
        );

        bytes32[] memory hashesOfPubkeysOfNonSigningOperators = new bytes32[](
            pubkeysOfNonSigningOperators.length
        );
        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            hashesOfPubkeysOfNonSigningOperators[i] = pubkeysOfNonSigningOperators[i].hashG1Point();
        }

        bytes32 signatoryRecordHash = keccak256(
            abi.encodePacked(
                task.taskCreatedBlock,
                hashesOfPubkeysOfNonSigningOperators
            )
        );
        require(
            signatoryRecordHash == taskResponseMetadata.hashOfNonSigners,
            "The pubkeys of non-signing operators supplied by the challenger are not correct."
        );

        $.taskSuccesfullyChallenged[referenceTaskId] = true;

        emit TaskChallengedSuccessfully(referenceTaskId, msg.sender);
    }

    function getTaskNumber() external view returns (uint32) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        return $.latestTaskId;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        return $.TASK_RESPONSE_WINDOW_BLOCK;
    }

    function getTaskResponse(uint32 taskId) external view returns (bytes32) {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        return $.allTaskResponses[taskId];
    }

    event NewTaskCreated(uint32 indexed taskId, ZrServiceManagerLib.Task task);
    event TaskResponded(ZrServiceManagerLib.TaskResponse response, ZrServiceManagerLib.TaskResponseMetadata metadata);
    event TaskChallengedSuccessfully(uint32 indexed taskId, address challenger);
}