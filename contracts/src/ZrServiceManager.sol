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
import "./interfaces/IZRTaskManager.sol"; // Use the interface instead of the full contract

import "./libraries/ZrServiceManagerLib.sol";
import "./ReentrancyGuard.sol";

import {IZrRegistryCoordinator} from "./interfaces/IZrRegistryCoordinator.sol";
import {ZrServiceManagerBase} from "./ZrServiceManagerBase.sol";

contract ZrServiceManager is
    ReentrancyGuard,
    Pausable,
    OperatorStateRetriever,
    IZrServiceManager,
    ZrServiceManagerBase
{
    using BN254 for BN254.G1Point;

    uint8 public constant QUORUM_NUMBER = 0;

    /// @custom:storage-location erc7201:zenrock.storage.ZrServiceManager
    struct ZrServiceManagerStorage {
        mapping(bytes32 => ZrServiceManagerLib.Validator) _validators;
        mapping(address => bytes32) operatorToValidator;
        mapping(address => uint256) operatorToValidatorIndex;
        bytes32[] _allValidators;
        IZRTaskManager taskManager; // Use the interface
    }

    bytes32 private constant ZR_SERVICE_MANAGER_STORAGE_LOCATION =
        0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00;

    function _getZrServiceManagerStorage()
        private
        pure
        returns (ZrServiceManagerStorage storage $)
    {
        assembly {
            $.slot := ZR_SERVICE_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRewardsCoordinator _rewardsCoordinator,
        IZrRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry
    )
        ZrServiceManagerBase(
            _avsDirectory,
            _rewardsCoordinator,
            _registryCoordinator,
            _stakeRegistry
        )
    {}

    function initialize(
        IPauserRegistry _pauserRegistry,
        address _initialOwner,
        IZRTaskManager _taskManager // Accept the address of the deployed TaskManager
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        __ZrServiceManager_init(_initialOwner);

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        // Assign the provided TaskManager
        $.taskManager = _taskManager;
    }

    function __ZrServiceManager_init(
        address _initialOwner
    ) internal virtual onlyInitializing {
        __ServiceManagerBase_init(_initialOwner);
    }

    function registerOperatorToAVS(
        address operatorAddr,
        string memory validatorAddr,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public virtual override onlyRegistryCoordinator {
        _ensureValidatorRegistered(validatorAddr);
        _avsDirectory.registerOperatorToAVS(operatorAddr, operatorSignature);
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        _assignOperatorToValidator(operatorAddr, validatorHash);
    }

    function deregisterOperatorFromAVS(
        address oprAddr
    )
        public
        virtual
        override(IZrServiceManager, ZrServiceManagerBase)
        onlyRegistryCoordinator
    {
        _deRegisterOperator(oprAddr);
        _avsDirectory.deregisterOperatorFromAVS(oprAddr);
    }

    function ejectValidators(
        string[] memory inactiveValidatorSet
    ) public virtual override {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        require(
            msg.sender == address($.taskManager),
            "Only TaskManager can eject validators"
        );
        _ejectValidators(inactiveValidatorSet);
    }

    function updateOperators() public virtual override {
        address[] memory operators = getAllOperatorsAddresses();
        _registryCoordinator.updateOperators(operators);
    }

    function getValidator(
        bytes32 validatorHash
    ) external view override returns (ZrServiceManagerLib.Validator memory) {
        return _getValidator(validatorHash);
    }

    function getAllValidator()
        external
        view
        override
        returns (ZrServiceManagerLib.Validator[] memory)
    {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        ZrServiceManagerLib.Validator[]
            memory validators = new ZrServiceManagerLib.Validator[](
                $._allValidators.length
            );
        for (uint256 i = 0; i < $._allValidators.length; i++) {
            validators[i] = _getValidator($._allValidators[i]);
        }
        return validators;
    }

    function getAllOperatorsAddresses()
        public
        view
        override
        returns (address[] memory)
    {
        bytes32[] memory registeredOperatorIDs = IRegistryCoordinator(
            address(_registryCoordinator)
        ).indexRegistry().getOperatorListAtBlockNumber(
                QUORUM_NUMBER,
                uint32(block.number)
            );
        address[] memory operators = new address[](
            registeredOperatorIDs.length
        );
        for (uint256 i = 0; i < registeredOperatorIDs.length; i++) {
            operators[i] = _registryCoordinator.getOperatorFromId(
                registeredOperatorIDs[i]
            );
        }
        return operators;
    }

    function getEigenStake(
        address operator,
        uint8 quorumNumber
    ) public view override returns (uint96) {
        return
            IStakeRegistry(address(_stakeRegistry)).getCurrentStake(
                _registryCoordinator.getOperator(operator).operatorId,
                quorumNumber
            );
    }

    function getStakedBalanceForValidator(
        string memory validatorAddr
    ) external view override returns (uint96) {
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        ZrServiceManagerLib.Validator memory validator = $._validators[
            validatorHash
        ];
        uint96 totalStake = 0;
        for (uint256 i = 0; i < validator.operators.length; i++) {
            totalStake += getEigenStake(validator.operators[i], QUORUM_NUMBER);
        }
        return totalStake;
    }

    function getTaskNumber() external view override returns (uint32) {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        return $.taskManager.taskNumber();
    }

    function getTaskResponseWindowBlock()
        external
        view
        override
        returns (uint32)
    {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        return $.taskManager.getTaskResponseWindowBlock();
    }

    // ------------------------------------------------- INTERNAL ----------------------------------------------------------------------
    function _ejectValidators(
        string[] memory inactiveValidatorSet
    ) internal virtual {
        bytes32[] memory validatorHashes = new bytes32[](
            inactiveValidatorSet.length
        );

        for (uint256 i = 0; i < inactiveValidatorSet.length; i++) {
            bytes32 validatorHash = keccak256(
                abi.encodePacked(inactiveValidatorSet[i])
            );
            validatorHashes[i] = validatorHash;
            _deRegisterValidator(validatorHash);
        }

        emit ValidatorsEjected(validatorHashes);
    }

    function _encodeQuorumNumbers(
        uint8[] memory quorumNumbers
    ) internal pure virtual returns (bytes memory) {
        bytes memory encodedData = new bytes(quorumNumbers.length);
        for (uint256 i = 0; i < quorumNumbers.length; ++i) {
            encodedData[i] = bytes1(quorumNumbers[i]);
        }
        return encodedData;
    }

    function _ensureValidatorRegistered(
        string memory validatorAddr
    ) internal virtual {
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        if (_validatorExists(validatorHash) == false) {
            _registerValidator(validatorAddr);
        }
    }

    function _registerValidator(
        string memory validatorAddr
    ) internal virtual returns (ZrServiceManagerLib.Validator memory) {
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));

        require(
            _validatorExists(validatorHash) == false,
            "Validator already registered"
        );

        ZrServiceManagerLib.Validator memory validator = ZrServiceManagerLib
            .Validator({
                validatorAddr: validatorAddr,
                validatorHash: validatorHash,
                operators: new address[](0)
            });

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        $._validators[validatorHash] = validator;
        $._allValidators.push(validatorHash);

        emit ValidatorRegistered(validatorHash, validatorAddr);

        return validator;
    }

    function _assignOperatorToValidator(
        address oprAddr,
        bytes32 validatorHash
    ) internal virtual {
        require(oprAddr != address(0), "zero operator address");
        require(validatorHash != bytes32(0), "invalid validator address");
        require(_validatorExists(validatorHash), "unknown validator hash");

        uint96 stake = getEigenStake(oprAddr, QUORUM_NUMBER);
        require(stake > 0, "invalid stake");

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        $.operatorToValidator[oprAddr] = validatorHash;
        $.operatorToValidatorIndex[oprAddr] = $
            ._validators[validatorHash]
            .operators
            .length;
        $._validators[validatorHash].operators.push(oprAddr);
        emit OperatorAssigned(oprAddr, validatorHash, stake);
    }

    function _deRegisterOperator(address oprAddr) internal virtual {
        require(oprAddr != address(0), "zero operator address");

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        bytes32 validatorHash = $.operatorToValidator[oprAddr];
        require(validatorHash != bytes32(0), "invalid validator hash");

        uint256 index = $.operatorToValidatorIndex[oprAddr];

        ZrServiceManagerLib.Validator storage validator = $._validators[
            validatorHash
        ];

        require(index < validator.operators.length, "Index out of bounds");
        require(
            validator.operators[index] == oprAddr,
            "Operator address mismatch"
        );

        uint256 lastOperatorIndex = validator.operators.length - 1;
        address lastOperator = validator.operators[lastOperatorIndex];
        
        if (index != lastOperatorIndex) {
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

        ZrServiceManagerLib.Validator storage validator = $._validators[
            validatorHash
        ];

        if (validator.validatorHash != bytes32(0)) {
            for (uint256 i = 0; i < validator.operators.length; i++) {
                _ejectOperator(validator.operators[i]);
                emit OperatorDeregistered(
                    validator.operators[i],
                    validatorHash
                );
            }

            delete $._validators[validatorHash];
        }
    }

    function _ejectOperator(address operatorAddr) internal virtual {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        bytes32 validatorHash = $.operatorToValidator[operatorAddr];
        if (validatorHash != bytes32(0)) {
            uint8[] memory quorumNumbers = new uint8[](1);
            quorumNumbers[0] = QUORUM_NUMBER;
            _registryCoordinator.ejectOperator(
                operatorAddr,
                _encodeQuorumNumbers(quorumNumbers)
            );
        }
    }

    function _validatorExists(
        bytes32 validatorHash
    ) internal view virtual returns (bool) {
        return bytes(_getValidator(validatorHash).validatorAddr).length != 0;
    }

    function _getValidator(
        bytes32 validatorHash
    ) internal view virtual returns (ZrServiceManagerLib.Validator memory) {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        return $._validators[validatorHash];
    }
}
