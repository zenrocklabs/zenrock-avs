// SPDX-License-Identifier: BUSL
// SPDX-FileCopyrightText: 2024 Zenrock labs Ltd.
pragma solidity ^0.8.21;

import {ZrServiceManagerLib} from "../libraries/ZrServiceManagerLib.sol";
import {ISignatureUtils} from "../../lib/eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

interface IZrServiceManager {
    // Events
    event ValidatorRegistered(
        bytes32 indexed validatorHash,
        string validatorAddr
    );
    event ValidatorDeregistered(bytes32 indexed validatorHash);
    event OperatorAssigned(
        address indexed operatorAddr,
        bytes32 indexed validatorHash
    );
    event OperatorDeregistered(
        address indexed operatorAddr,
        bytes32 indexed validatorHash
    );
    event ValidatorsEjected(bytes32[] validatorHashes);
    event NewTaskCreated(uint32 taskId, ZrServiceManagerLib.Task task);
    event TaskResponded(
        ZrServiceManagerLib.TaskResponse taskResponse,
        ZrServiceManagerLib.TaskResponseMetadata taskResponseMetadata
    );
    event TaskChallengedSuccessfully(uint32 taskId, address challenger);

    function registerOperatorToAVS(
        address operatorAddr,
        string memory validatorAddr,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    function deregisterOperatorFromAVS(address oprAddr) external;

    function ejectValidators(string[] memory inactiveValidatorSet) external;

    function updateOperators() external;

    function getValidator(
        bytes32 validatorHash
    ) external view returns (ZrServiceManagerLib.Validator memory);

    function getAllValidator()
        external
        view
        returns (ZrServiceManagerLib.Validator[] memory);

    function getAllOperatorsAddresses()
        external
        view
        returns (address[] memory);

    function getEigenStake(
        address operator,
        uint8 quorumNumber
    ) external view returns (uint96);

    function getStakedBalanceForValidator(
        string memory validatorAddress
    ) external view returns (uint96);
}
