// SPDX-License-Identifier: BUSL
pragma solidity 0.8.21;

import "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {IAVSDirectory} from "../lib/eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IRewardsCoordinator} from "../lib/eigenlayer-contracts/src/contracts/interfaces/IRewardsCoordinator.sol";
import "../lib/eigenlayer-contracts/src/contracts/permissions/Pausable.sol";
import {IStakeRegistry} from "../lib/eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IRegistryCoordinator} from "../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {OperatorStateRetriever} from "../lib/eigenlayer-middleware/src/OperatorStateRetriever.sol";
import {IServiceManager} from "../lib/eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import "./interfaces/IZrServiceManager.sol";
import "./interfaces/IZRTaskManager.sol";
import "./libraries/ZrServiceManagerLib.sol";
import "./ReentrancyGuard.sol";
import {IZrRegistryCoordinator} from "./interfaces/IZrRegistryCoordinator.sol";
import {ZrServiceManagerBase} from "./ZrServiceManagerBase.sol";

/**
 * @title ZrServiceManager
 * @notice Manages the registration and coordination of validators and operators
 * @dev Proxy-compatible implementation with gas optimizations
 */
contract ZrServiceManager is ReentrancyGuard, Pausable, OperatorStateRetriever, IZrServiceManager, ZrServiceManagerBase {
    using BN254 for BN254.G1Point;

    // Constants - can't use immutable with proxy
    uint8 private constant QUORUM_NUMBER = 0;
    bytes32 private constant ZR_SERVICE_MANAGER_STORAGE_LOCATION = 0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00;

    // Custom errors
    error ZeroAddress();
    error InvalidValidatorHash();
    error ValidatorAlreadyRegistered();
    error UnknownValidatorHash();
    error IndexOutOfBounds();
    error OperatorAddressMismatch();
    error OnlyTaskManager();
    error EmptyValidatorSet();
    error InvalidValidatorAddress();
    error InvalidOperatorAddress();

    /// @custom:storage-location erc7201:zenrock.storage.ZrServiceManager
    struct ZrServiceManagerStorage {
        // Packed structs for better storage efficiency
        mapping(bytes32 => ZrServiceManagerLib.Validator) _validators;
        mapping(address => bytes32) operatorToValidator;
        mapping(address => uint256) operatorToValidatorIndex;
        bytes32[] _allValidators;
        IZRTaskManager taskManager;
        
        // Cache frequently accessed values
        uint256 validatorCount;
        mapping(bytes32 => bool) validatorExists;
    }

    /**
     * @dev Returns storage pointer to ZrServiceManagerStorage
     */
    function _getZrServiceManagerStorage() private pure returns (ZrServiceManagerStorage storage $) {
        assembly {
            $.slot := ZR_SERVICE_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRewardsCoordinator _rewardsCoordinator,
        IZrRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry
    ) ZrServiceManagerBase(_avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry) {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract
     * @param _pauserRegistry The pauser registry contract
     * @param _initialOwner The initial owner address
     * @param _taskManager The task manager contract
     */
    function initialize(
        IPauserRegistry _pauserRegistry,
        address _initialOwner,
        IZRTaskManager _taskManager
    ) external initializer {
        if (address(_taskManager) == address(0)) revert ZeroAddress();
        if (_initialOwner == address(0)) revert ZeroAddress();
        
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        __ZrServiceManager_init(_initialOwner);
        
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        $.taskManager = _taskManager;
        $.validatorCount = 0;
    }

    function __ZrServiceManager_init(address _initialOwner) internal virtual onlyInitializing {
        __ServiceManagerBase_init(_initialOwner);
    }

    /**
     * @notice Registers an operator to AVS
     * @param operatorAddr The operator address
     * @param validatorAddr The validator address
     * @param operatorSignature The operator's signature
     */
    function registerOperatorToAVS(
        address operatorAddr,
        string calldata validatorAddr,
        ISignatureUtils.SignatureWithSaltAndExpiry calldata operatorSignature
    ) external virtual override onlyRegistryCoordinator {
        if (operatorAddr == address(0)) revert InvalidOperatorAddress();
        if (bytes(validatorAddr).length == 0) revert InvalidValidatorAddress();

        bytes32 validatorHash = keccak256(bytes(validatorAddr));
        _ensureValidatorRegistered(validatorAddr);
        _avsDirectory.registerOperatorToAVS(operatorAddr, operatorSignature);
        _assignOperatorToValidator(operatorAddr, validatorHash);
    }

    /**
     * @notice Deregisters an operator from AVS
     * @param oprAddr The operator address
     */
    function deregisterOperatorFromAVS(
        address oprAddr
    ) public virtual override(IZrServiceManager, ZrServiceManagerBase) onlyRegistryCoordinator {
        if (oprAddr == address(0)) revert InvalidOperatorAddress();
        
        _deRegisterOperator(oprAddr);
        _avsDirectory.deregisterOperatorFromAVS(oprAddr);
    }

    /**
     * @notice Ejects validators from the system
     * @param inactiveValidatorSet Array of validator addresses to eject
     */
    function ejectValidators(string[] calldata inactiveValidatorSet) external virtual override {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        if (msg.sender != address($.taskManager)) revert OnlyTaskManager();
        if (inactiveValidatorSet.length == 0) revert EmptyValidatorSet();
        
        _ejectValidators(inactiveValidatorSet);
    }

    /**
     * @notice Updates operators in the registry coordinator
     */
    function updateOperators() external virtual override {
        _registryCoordinator.updateOperators(getAllOperatorsAddresses());
    }

    /**
     * @notice Gets validator information
     * @param validatorHash The validator's hash
     */
    function getValidator(bytes32 validatorHash) external view override returns (ZrServiceManagerLib.Validator memory) {
        return _getValidator(validatorHash);
    }

    /**
     * @notice Gets all validators
     */
    function getAllValidator() external view override returns (ZrServiceManagerLib.Validator[] memory) {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        uint256 length = $._allValidators.length;
        ZrServiceManagerLib.Validator[] memory validators = new ZrServiceManagerLib.Validator[](length);
        
        // Cache storage reads
        bytes32[] memory validatorHashes = $._allValidators;
        
        unchecked {
            for (uint256 i; i < length; ++i) {
                validators[i] = $._validators[validatorHashes[i]];
            }
        }
        return validators;
    }

    /**
     * @notice Gets all operator addresses
     */
    function getAllOperatorsAddresses() public view override returns (address[] memory) {
        bytes32[] memory registeredOperatorIDs = IRegistryCoordinator(address(_registryCoordinator))
            .indexRegistry()
            .getOperatorListAtBlockNumber(QUORUM_NUMBER, uint32(block.number));
            
        uint256 length = registeredOperatorIDs.length;
        address[] memory operators = new address[](length);
        
        unchecked {
            for (uint256 i; i < length; ++i) {
                operators[i] = _registryCoordinator.getOperatorFromId(registeredOperatorIDs[i]);
            }
        }
        return operators;
    }

    /**
     * @notice Gets the EigenLayer stake for an operator
     * @param operator The operator address
     * @param quorumNumber The quorum number
     */
    function getEigenStake(address operator, uint8 quorumNumber) public view override returns (uint96) {
        return IStakeRegistry(address(_stakeRegistry)).getCurrentStake(
            _registryCoordinator.getOperator(operator).operatorId,
            quorumNumber
        );
    }

    /**
     * @notice Gets the staked balance for a validator
     * @param validatorAddr The validator address
     */
    function getStakedBalanceForValidator(string calldata validatorAddr) external view override returns (uint96) {
        bytes32 validatorHash = keccak256(bytes(validatorAddr));
        ZrServiceManagerLib.Validator memory validator = _getZrServiceManagerStorage()._validators[validatorHash];
        uint96 totalStake;
        
        unchecked {
            for (uint256 i; i < validator.operators.length; ++i) {
                totalStake += getEigenStake(validator.operators[i], QUORUM_NUMBER);
            }
        }
        return totalStake;
    }

    /**
     * @notice Gets the task manager
     */
    function taskManager() external view returns (IZRTaskManager) {
        return _getZrServiceManagerStorage().taskManager;
    }

    // Internal functions

    function _ejectValidators(string[] memory inactiveValidatorSet) internal virtual {
        uint256 length = inactiveValidatorSet.length;
        bytes32[] memory validatorHashes = new bytes32[](length);
        
        unchecked {
            for (uint256 i; i < length; ++i) {
                bytes32 validatorHash = keccak256(bytes(inactiveValidatorSet[i]));
                validatorHashes[i] = validatorHash;
                _deRegisterValidator(validatorHash);
            }
        }
        
        emit ValidatorsEjected(validatorHashes);
    }

    function _encodeQuorumNumbers(uint8[] memory quorumNumbers) internal pure virtual returns (bytes memory) {
        uint256 length = quorumNumbers.length;
        bytes memory encodedData = new bytes(length);
        
        unchecked {
            for (uint256 i; i < length; ++i) {
                encodedData[i] = bytes1(quorumNumbers[i]);
            }
        }
        return encodedData;
    }

    function _ensureValidatorRegistered(string memory validatorAddr) internal virtual {
        bytes32 validatorHash = keccak256(bytes(validatorAddr));
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        
        if (!$.validatorExists[validatorHash]) {
            _registerValidator(validatorAddr, validatorHash);
        }
    }

    function _registerValidator(string memory validatorAddr, bytes32 validatorHash) internal virtual returns (ZrServiceManagerLib.Validator memory) {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        
        if ($.validatorExists[validatorHash]) revert ValidatorAlreadyRegistered();

        ZrServiceManagerLib.Validator memory validator = ZrServiceManagerLib.Validator({
            validatorAddr: validatorAddr,
            validatorHash: validatorHash,
            operators: new address[](0)
        });

        $._validators[validatorHash] = validator;
        $._allValidators.push(validatorHash);
        $.validatorExists[validatorHash] = true;
        $.validatorCount++;

        emit ValidatorRegistered(validatorHash, validatorAddr);
        return validator;
    }

    function _assignOperatorToValidator(address oprAddr, bytes32 validatorHash) internal virtual {
        if (oprAddr == address(0)) revert InvalidOperatorAddress();
        if (validatorHash == bytes32(0)) revert InvalidValidatorHash();
        
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        if (!$.validatorExists[validatorHash]) revert UnknownValidatorHash();

        $.operatorToValidator[oprAddr] = validatorHash;
        $.operatorToValidatorIndex[oprAddr] = $._validators[validatorHash].operators.length;
        $._validators[validatorHash].operators.push(oprAddr);
        
        emit OperatorAssigned(oprAddr, validatorHash);
    }

    function _deRegisterOperator(address oprAddr) internal virtual {
        if (oprAddr == address(0)) revert InvalidOperatorAddress();

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        bytes32 validatorHash = $.operatorToValidator[oprAddr];
        if (validatorHash == bytes32(0)) revert InvalidValidatorHash();

        uint256 index = $.operatorToValidatorIndex[oprAddr];
        ZrServiceManagerLib.Validator storage validator = $._validators[validatorHash];

        if (index >= validator.operators.length) revert IndexOutOfBounds();
        if (validator.operators[index] != oprAddr) revert OperatorAddressMismatch();

        // Replace with last element and pop
        uint256 lastIndex = validator.operators.length - 1;
        if (index != lastIndex) {
            address lastOperator = validator.operators[lastIndex];
            validator.operators[index] = lastOperator;
            $.operatorToValidatorIndex[lastOperator] = index;
        }
        validator.operators.pop();

        delete $.operatorToValidator[oprAddr];
        delete $.operatorToValidatorIndex[oprAddr];

        emit OperatorDeregistered(oprAddr, validatorHash);
    }

    function _deRegisterValidator(bytes32 validatorHash) internal virtual {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        
        if (!$.validatorExists[validatorHash]) return;

        ZrServiceManagerLib.Validator storage validator = $._validators[validatorHash];
        uint256 length = validator.operators.length;
        
        unchecked {
            for (uint256 i; i < length; ++i) {
                address operator = validator.operators[i];
                _ejectOperator(operator);
                emit OperatorDeregistered(operator, validatorHash);
            }
        }
        
        delete $._validators[validatorHash];
        delete $.validatorExists[validatorHash];
        $.validatorCount--;
    }

    function _ejectOperator(address operatorAddr) internal virtual {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        bytes32 validatorHash = $.operatorToValidator[operatorAddr];
        
        if (validatorHash != bytes32(0)) {
            uint8[] memory quorumNumbers = new uint8[](1);
            quorumNumbers[0] = QUORUM_NUMBER;
            _registryCoordinator.ejectOperator(operatorAddr, _encodeQuorumNumbers(quorumNumbers));
        }
    }

    function _validatorExists(bytes32 validatorHash) internal view virtual returns (bool) {
        return _getZrServiceManagerStorage().validatorExists[validatorHash];
    }

    function _getValidator(bytes32 validatorHash) internal view virtual returns (ZrServiceManagerLib.Validator memory) {
        return _getZrServiceManagerStorage()._validators[validatorHash];
    }
}