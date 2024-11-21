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

contract ZrServiceManager is ReentrancyGuard, Pausable, OperatorStateRetriever, IZrServiceManager, ZrServiceManagerBase {
    using BN254 for BN254.G1Point;

    uint8 private constant QUORUM_NUMBER = 0;
    bytes32 private constant ZR_SERVICE_MANAGER_STORAGE_LOCATION = 0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00;

    error ZeroAddress();
    error InvalidValidatorHash();
    error ValidatorAlreadyRegistered();
    error UnknownValidatorHash();
    error IndexOutOfBounds();
    error OperatorAddressMismatch();
    error OnlyTaskManager();

    /// @custom:storage-location erc7201:zenrock.storage.ZrServiceManager
    struct ZrServiceManagerStorage {
        mapping(bytes32 => ZrServiceManagerLib.Validator) _validators;
        mapping(address => bytes32) operatorToValidator;
        mapping(address => uint256) operatorToValidatorIndex;
        bytes32[] _allValidators;
        IZRTaskManager taskManager;
    }

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
    ) ZrServiceManagerBase(_avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry) {}

    function initialize(
        IPauserRegistry _pauserRegistry,
        address _initialOwner,
        IZRTaskManager _taskManager
    ) external initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        __ZrServiceManager_init(_initialOwner);
        _getZrServiceManagerStorage().taskManager = _taskManager;
    }

    function __ZrServiceManager_init(address _initialOwner) internal virtual onlyInitializing {
        __ServiceManagerBase_init(_initialOwner);
    }

    function registerOperatorToAVS(
        address operatorAddr,
        string calldata validatorAddr,
        ISignatureUtils.SignatureWithSaltAndExpiry calldata operatorSignature
    ) external virtual override onlyRegistryCoordinator {
        _ensureValidatorRegistered(validatorAddr);
        _avsDirectory.registerOperatorToAVS(operatorAddr, operatorSignature);
        _assignOperatorToValidator(operatorAddr, keccak256(abi.encodePacked(validatorAddr)));
    }

    function deregisterOperatorFromAVS(
        address oprAddr
    ) public virtual override(IZrServiceManager, ZrServiceManagerBase) onlyRegistryCoordinator {
        _deRegisterOperator(oprAddr);
        _avsDirectory.deregisterOperatorFromAVS(oprAddr);
    }

    function ejectValidators(string[] calldata inactiveValidatorSet) external virtual override {
        if (msg.sender != address(_getZrServiceManagerStorage().taskManager)) revert OnlyTaskManager();
        _ejectValidators(inactiveValidatorSet);
    }

    function updateOperators() external virtual override {
        _registryCoordinator.updateOperators(getAllOperatorsAddresses());
    }

    function getValidator(bytes32 validatorHash) external view override returns (ZrServiceManagerLib.Validator memory) {
        return _getValidator(validatorHash);
    }

    function getAllValidator() external view override returns (ZrServiceManagerLib.Validator[] memory) {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        uint256 length = $._allValidators.length;
        ZrServiceManagerLib.Validator[] memory validators = new ZrServiceManagerLib.Validator[](length);
        
        unchecked {
            for (uint256 i; i < length; ++i) {
                validators[i] = _getValidator($._allValidators[i]);
            }
        }
        return validators;
    }

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

    function getEigenStake(address operator, uint8 quorumNumber) public view override returns (uint96) {
        return IStakeRegistry(address(_stakeRegistry)).getCurrentStake(
            _registryCoordinator.getOperator(operator).operatorId,
            quorumNumber
        );
    }

    function getStakedBalanceForValidator(string calldata validatorAddr) external view override returns (uint96) {
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        ZrServiceManagerLib.Validator memory validator = _getZrServiceManagerStorage()._validators[validatorHash];
        uint96 totalStake;
        
        unchecked {
            for (uint256 i; i < validator.operators.length; ++i) {
                totalStake += getEigenStake(validator.operators[i], QUORUM_NUMBER);
            }
        }
        return totalStake;
    }

    function getTaskNumber() external view override returns (uint32) {
        return _getZrServiceManagerStorage().taskManager.taskNumber();
    }

    function getTaskResponseWindowBlock() external view override returns (uint32) {
        return _getZrServiceManagerStorage().taskManager.getTaskResponseWindowBlock();
    }

    function taskManager() external view returns (IZRTaskManager) {
        return _getZrServiceManagerStorage().taskManager;
    }

    // Internal functions below

    function _ejectValidators(string[] memory inactiveValidatorSet) internal virtual {
        uint256 length = inactiveValidatorSet.length;
        bytes32[] memory validatorHashes = new bytes32[](length);
        
        unchecked {
            for (uint256 i; i < length; ++i) {
                bytes32 validatorHash = keccak256(abi.encodePacked(inactiveValidatorSet[i]));
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
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        if (!_validatorExists(validatorHash)) {
            _registerValidator(validatorAddr, validatorHash);
        }
    }

    function _registerValidator(string memory validatorAddr, bytes32 validatorHash) internal virtual returns (ZrServiceManagerLib.Validator memory) {
        if (_validatorExists(validatorHash)) revert ValidatorAlreadyRegistered();

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        ZrServiceManagerLib.Validator memory validator = ZrServiceManagerLib.Validator({
            validatorAddr: validatorAddr,
            validatorHash: validatorHash,
            operators: new address[](0)
        });

        $._validators[validatorHash] = validator;
        $._allValidators.push(validatorHash);

        emit ValidatorRegistered(validatorHash, validatorAddr);
        return validator;
    }

    function _assignOperatorToValidator(address oprAddr, bytes32 validatorHash) internal virtual {
        if (oprAddr == address(0)) revert ZeroAddress();
        if (validatorHash == bytes32(0)) revert InvalidValidatorHash();
        if (!_validatorExists(validatorHash)) revert UnknownValidatorHash();

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        $.operatorToValidator[oprAddr] = validatorHash;
        $.operatorToValidatorIndex[oprAddr] = $._validators[validatorHash].operators.length;
        $._validators[validatorHash].operators.push(oprAddr);
        
        emit OperatorAssigned(oprAddr, validatorHash);
    }

    function _deRegisterOperator(address oprAddr) internal virtual {
        if (oprAddr == address(0)) revert ZeroAddress();

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        bytes32 validatorHash = $.operatorToValidator[oprAddr];
        if (validatorHash == bytes32(0)) revert InvalidValidatorHash();

        uint256 index = $.operatorToValidatorIndex[oprAddr];
        ZrServiceManagerLib.Validator storage validator = $._validators[validatorHash];

        if (index >= validator.operators.length) revert IndexOutOfBounds();
        if (validator.operators[index] != oprAddr) revert OperatorAddressMismatch();

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
        ZrServiceManagerLib.Validator storage validator = $._validators[validatorHash];

        if (validator.validatorHash != bytes32(0)) {
            uint256 length = validator.operators.length;
            unchecked {
                for (uint256 i; i < length; ++i) {
                    address operator = validator.operators[i];
                    _ejectOperator(operator);
                    emit OperatorDeregistered(operator, validatorHash);
                }
            }
            delete $._validators[validatorHash];
        }
    }

    function _ejectOperator(address operatorAddr) internal virtual {
        bytes32 validatorHash = _getZrServiceManagerStorage().operatorToValidator[operatorAddr];
        if (validatorHash != bytes32(0)) {
            uint8[] memory quorumNumbers = new uint8[](1);
            quorumNumbers[0] = QUORUM_NUMBER;
            _registryCoordinator.ejectOperator(operatorAddr, _encodeQuorumNumbers(quorumNumbers));
        }
    }

    function _validatorExists(bytes32 validatorHash) internal view virtual returns (bool) {
        return bytes(_getValidator(validatorHash).validatorAddr).length != 0;
    }

    function _getValidator(bytes32 validatorHash) internal view virtual returns (ZrServiceManagerLib.Validator memory) {
        return _getZrServiceManagerStorage()._validators[validatorHash];
    }
}