// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../../lib/eigenlayer-middleware/src/libraries/BN254.sol";

interface IZRTaskManager {
    // EVENTS
    event NewTaskCreated(uint32 indexed taskId, Task task);

    event TaskResponded(
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event TaskCompleted(uint32 indexed taskId);

    event TaskChallengedSuccessfully(
        uint32 indexed taskId,
        address indexed challenger
    );

    event TaskChallengedUnsuccessfully(
        uint32 indexed taskId,
        address indexed challenger
    );
    
    event ValidatorAddressesStored(uint32 indexed taskId, string[] addresses);

    // STRUCTS
    struct Task {
        uint32 taskId;
        uint32 taskCreatedBlock;
        // task submitter decides on the criteria for a task to be completed
        // note that this does not mean the task was "correctly" answered (i.e. the ID was signed correctly)
        // - this is for the challenge logic to verify
        // task is completed (and contract will accept its TaskResponse) when each quorumNumbers specified here
        // are signed by at least quorumThresholdPercentage of the operators
        // note that we set the quorumThresholdPercentage to be the same for all quorumNumbers, but this could be changed
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
    }

    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct TaskResponse {
        // Can be obtained by the operator from the event NewTaskCreated.
        uint32 referenceTaskId;
        string[] activeSetZRChain;
    }

    // Extra information related to taskResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than TaskResponse
    // This metadata is needed by the challenger, so we emit it in the TaskResponded event
    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS
    function createNewTask(
        uint32 taskId,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external;

    function taskNumber() external view returns (uint32);

    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external;

    function getTaskResponseWindowBlock() external view returns (uint32);
    function getLatestActiveSet() external view returns (string[] memory);
}
