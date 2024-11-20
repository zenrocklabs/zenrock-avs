// SPDX-License-Identifier: BUSL
pragma solidity 0.8.21;

import {BN254} from "../lib/eigenlayer-middleware/src/libraries/BN254.sol";
import {IRegistryCoordinator} from "../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BLSSignatureChecker, IBLSSignatureChecker} from "../lib/eigenlayer-middleware/src/BLSSignatureChecker.sol";
import "./libraries/ZrServiceManagerLib.sol";
import {Initializable} from "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "../lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";

contract ZrTaskManager is Initializable, OwnableUpgradeable, BLSSignatureChecker {
    using BN254 for BN254.G1Point;

    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 10000;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 67;

    /// @custom:storage-location erc7201:zenrock.storage.TaskManager
    struct TaskManagerStorage {
        uint32 TASK_RESPONSE_WINDOW_BLOCK;
        uint32 latestTaskId;
        mapping(uint32 => bytes32) allTaskHashes;
        mapping(uint32 => bytes32) allTaskResponses;
        mapping(uint32 => bool) taskSuccesfullyChallenged;
        address aggregator;
        address generator;
        IRegistryCoordinator registryCoordinator;
    }

    // keccak256(abi.encode(uint256(keccak256("zenrock.storage.taskmanager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant TASK_MANAGER_STORAGE_LOCATION = 
        0x9e5d0bf83ef884a66a66b2d439fd65f5546f8f4489c6a744f987ecb90e5d7100;

    constructor(address _registryCoordinator) BLSSignatureChecker(IRegistryCoordinator(_registryCoordinator)) {}

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

    function initialize(
        address _aggregator,
        address _generator,
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock,
        address initialOwner
    ) external initializer {
        TaskManagerStorage storage $ = _getTaskManagerStorage();
        $.aggregator = _aggregator;
        $.generator = _generator;
        $.registryCoordinator = _registryCoordinator;
        $.TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
        _transferOwnership(initialOwner);
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
        emit NewTaskCreated(taskId, newTask);
        $.latestTaskId = taskId;
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

        require(
            keccak256(abi.encode(task)) == $.allTaskHashes[task.taskId],
            "supplied task does not match the one recorded in the contract"
        );
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

        ZrServiceManagerLib.TaskResponseMetadata memory taskResponseMetadata = ZrServiceManagerLib
            .TaskResponseMetadata(uint32(block.number), hashOfNonSigners);
        $.allTaskResponses[task.taskId] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        emit TaskResponded(taskResponse, taskResponseMetadata);
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