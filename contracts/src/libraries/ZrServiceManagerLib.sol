// SPDX-License-Identifier: BUSL
// SPDX-FileCopyrightText: 2024 Zenrock labs Ltd.
pragma solidity ^0.8.28;
library ZrServiceManagerLib {
    struct Validator {
        string validatorAddr;
        bytes32 validatorHash;
        address[] operators;
    }

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
        string[] inactiveSetZRChain;
    }

    // Extra information related to taskResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than TaskResponse
    // This metadata is needed by the challenger, so we emit it in the TaskResponded event
    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }
}
