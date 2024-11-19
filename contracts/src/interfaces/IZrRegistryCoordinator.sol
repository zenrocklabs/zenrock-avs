// SPDX-License-Identifier: BUSL
// SPDX-FileCopyrightText: 2024 Zenrock labs Ltd.

pragma solidity ^0.8.21;

import {IRegistryCoordinator} from "../../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";

interface IZrRegistryCoordinator is IRegistryCoordinator {
    function updateOperators(address[] calldata operators) external;
}
