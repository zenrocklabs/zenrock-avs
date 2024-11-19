// SPDX-License-Identifier: BUSL
// SPDX-FileCopyrightText: 2024 Zenrock labs Ltd.
pragma solidity 0.8.21;

import "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";

import {IAVSDirectory} from "../lib/eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IRewardsCoordinator} from "../lib/eigenlayer-contracts/src/contracts/interfaces/IRewardsCoordinator.sol";
import "../lib/eigenlayer-contracts/src/contracts/permissions/Pausable.sol";

import {IStakeRegistry} from "../lib/eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {BN254} from "../lib/eigenlayer-middleware/src/libraries/BN254.sol";
import {BLSApkRegistry} from "../lib/eigenlayer-middleware/src/BLSApkRegistry.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "../lib/eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "../lib/eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {IServiceManager} from "../lib/eigenlayer-middleware/src/interfaces/IServiceManager.sol";

import "./interfaces/IZrServiceManager.sol";
import "./interfaces/IZRTaskManager.sol";

import "./libraries/ZrServiceManagerLib.sol";
import "./ReentrancyGuard.sol";

import {IZrRegistryCoordinator} from "./interfaces/IZrRegistryCoordinator.sol";
import {ZrServiceManagerBase} from "./ZrServiceManagerBase.sol";

contract ZrServiceManager is
    ReentrancyGuard,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    IZrServiceManager,
    ZrServiceManagerBase
{
    using BN254 for BN254.G1Point;

    uint8 public constant QUORUM_NUMBER = 0;
    uint256 public constant UNDELEGATION_PERIOD = 151200;

    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 10000;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 67;

    /// @custom:storage-location erc7201:zenrock.storage.ZrServiceManager
    struct ZrServiceManagerStorage {
        mapping(bytes32 => ZrServiceManagerLib.Validator) _validators;
        mapping(address => bytes32) operatorToValidator;
        mapping(address => uint256) operatorToValidatorIndex;
        bytes32[] _allValidators;
        //----------------------- TASK MANAGER -----------------------//
        uint32 TASK_RESPONSE_WINDOW_BLOCK;
        /* STORAGE */
        uint32 latestTaskId;
        // mapping of task indices to all tasks hashes
        // when a task is created, task hash is stored here,
        // and responses need to pass the actual task,
        // which is hashed onchain and checked against this mapping
        mapping(uint32 => bytes32) allTaskHashes;
        // mapping of task indices to hash of abi.encode(taskResponse, taskResponseMetadata)
        mapping(uint32 => bytes32) allTaskResponses;
        mapping(uint32 => bool) taskSuccesfullyChallenged;
        address aggregator;
        address generator;
    }

    // keccak256(abi.encode(uint256(keccak256("zenrock.storage.zrservicemanager")) - 1)) & ~bytes32(uint256(0xff));
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

    /* MODIFIERS */
    modifier onlyAggregator() {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        require(msg.sender == $.aggregator, "Aggregator must be the caller");
        _;
    }

    // onlyTaskGenerator is used to restrict createNewTask from only being called by a permissioned entity
    // in a real world scenario, this would be removed by instead making createNewTask a payable function
    modifier onlyTaskGenerator() {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        require(msg.sender == $.generator, "Task generator must be the caller");
        _;
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
        BLSSignatureChecker(_registryCoordinator)
    {}

    function initialize(
        IPauserRegistry _pauserRegistry,
        address _initialOwner,
        address _aggregator,
        address _generator,
        // address _rewardsInitiator,
        uint32 _taskResponseWindowBlock
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        __ZrServiceManager_init(_initialOwner); //, _rewardsInitiator);
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        $.aggregator = _aggregator;
        $.generator = _generator;
        $.TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function __ZrServiceManager_init(
        address _initialOwner
        // address _rewardsInitiator
    ) internal virtual onlyInitializing {
        __ServiceManagerBase_init(_initialOwner); //, _rewardsInitiator);
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
    ) public virtual override onlyAggregator {
        _ejectValidators(inactiveValidatorSet);
    }

    function updateOperators() public virtual override {
        address[] memory operators = getAllOperatorsAddresses();
        _registryCoordinator.updateOperators(operators);
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
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        $.allTaskHashes[taskId] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(taskId, newTask);
        $.latestTaskId = taskId;
        updateOperators();
    }

    function respondToTask(
        ZrServiceManagerLib.Task calldata task,
        ZrServiceManagerLib.TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        // check that the task is valid, hasn't been responsed yet, and is being responsed in time
        require(
            keccak256(abi.encode(task)) == $.allTaskHashes[task.taskId],
            "supplied task does not match the one recorded in the contract"
        );
        require(
            $.allTaskResponses[task.taskId] == bytes32(0),
            "Aggregator has already responded to the task"
        );
        require(
            uint32(block.number) <=
                taskCreatedBlock + $.TASK_RESPONSE_WINDOW_BLOCK,
            "Aggregator has responded to the task too late"
        );

        /* CHECKING SIGNATURES & WHETHER THRESHOLD IS MET OR NOT */
        // calculate message which operators signed
        bytes32 message = keccak256(abi.encode(taskResponse));

        // check the aggregated BLS signature
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
                message,
                quorumNumbers,
                taskCreatedBlock,
                nonSignerStakesAndSignature
            );

        // check that signatories own at least a threshold percentage of each quourm
        for (uint i = 0; i < quorumNumbers.length; i++) {
            // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying
            // signed stake > total stake
            require(
                quorumStakeTotals.signedStakeForQuorum[i] *
                    _THRESHOLD_DENOMINATOR >=
                    quorumStakeTotals.totalStakeForQuorum[i] *
                        uint8(quorumThresholdPercentage),
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }

        if (taskResponse.inactiveSetZRChain.length > 0) {
            _ejectValidators(taskResponse.inactiveSetZRChain);
        }

        ZrServiceManagerLib.TaskResponseMetadata
            memory taskResponseMetadata = ZrServiceManagerLib
                .TaskResponseMetadata(uint32(block.number), hashOfNonSigners);
        // updating the storage with task response
        $.allTaskResponses[task.taskId] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        emit TaskResponded(taskResponse, taskResponseMetadata);
    }

    // NOTE: this function enables a challenger to raise and resolve a challenge.
    // TODO: require challenger to pay a bond for raising a challenge
    // TODO(samlaf): should we check that quorumNumbers is same as the one recorded in the task?
    function raiseAndResolveChallenge(
        ZrServiceManagerLib.Task calldata task,
        ZrServiceManagerLib.TaskResponse calldata taskResponse,
        ZrServiceManagerLib.TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        uint32 referenceTaskId = taskResponse.referenceTaskId;
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

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
            $.taskSuccesfullyChallenged[referenceTaskId] == false,
            "The response to this task has already been challenged successfully."
        );
        require(
            uint32(block.number) <=
                taskResponseMetadata.taskResponsedBlock +
                    TASK_CHALLENGE_WINDOW_BLOCK,
            "The challenge period for this task has already expired."
        );

        // Verify that the taskId matches
        require(
            task.taskId == taskResponse.referenceTaskId,
            "Task ID mismatch"
        );

        // get the list of hash of pubkeys of operators who weren't part of the task response submitted by the aggregator
        bytes32[] memory hashesOfPubkeysOfNonSigningOperators = new bytes32[](
            pubkeysOfNonSigningOperators.length
        );
        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            hashesOfPubkeysOfNonSigningOperators[
                i
            ] = pubkeysOfNonSigningOperators[i].hashG1Point();
        }

        // verify whether the pubkeys of "claimed" non-signers supplied by challenger are actually non-signers as recorded before
        // when the aggregator responded to the task
        // currently inlined, as the MiddlewareUtils.computeSignatoryRecordHash function was removed from BLSSignatureChecker
        // in this PR: https://github.com/Layr-Labs/eigenlayer-contracts/commit/c836178bf57adaedff37262dff1def18310f3dce#diff-8ab29af002b60fc80e3d6564e37419017c804ae4e788f4c5ff468ce2249b4386L155-L158
        // TODO(samlaf): contracts team will add this function back in the BLSSignatureChecker, which we should use to prevent potential bugs from code duplication
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

        // get the address of operators who didn't sign
        address[] memory addresssOfNonSigningOperators = new address[](
            pubkeysOfNonSigningOperators.length
        );
        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            addresssOfNonSigningOperators[i] = BLSApkRegistry(
                address(blsApkRegistry)
            ).pubkeyHashToOperator(hashesOfPubkeysOfNonSigningOperators[i]);
        }

        // @dev the below code is commented out for the upcoming M2 release
        //      in which there will be no slashing. The slasher is also being redesigned
        //      so its interface may very well change.
        // ==========================================
        // // get the list of all operators who were active when the task was initialized
        // Operator[][] memory allOperatorInfo = getOperatorState(
        //     IRegistryCoordinator(address(registryCoordinator)),
        //     task.quorumNumbers,
        //     task.taskCreatedBlock
        // );
        // // freeze the operators who signed adversarially
        // for (uint i = 0; i < allOperatorInfo.length; i++) {
        //     // first for loop iterate over quorums

        //     for (uint j = 0; j < allOperatorInfo[i].length; j++) {
        //         // second for loop iterate over operators active in the quorum when the task was initialized

        //         // get the operator address
        //         bytes32 operatorID = allOperatorInfo[i][j].operatorId;
        //         address operatorAddress = BLSPubkeyRegistry(
        //             address(blsPubkeyRegistry)
        //         ).pubkeyCompendium().pubkeyHashToOperator(operatorID);

        //         // check if the operator has already NOT been frozen
        //         if (
        //             IServiceManager(
        //                 address(
        //                     BLSRegistryCoordinatorWithIndices(
        //                         address(registryCoordinator)
        //                     ).serviceManager()
        //                 )
        //             ).slasher().isFrozen(operatorAddress) == false
        //         ) {
        //             // check whether the operator was a signer for the task
        //             bool wasSigningOperator = true;
        //             for (
        //                 uint k = 0;
        //                 k < addresssOfNonSigningOperators.length;
        //                 k++
        //             ) {
        //                 if (
        //                     operatorAddress == addresssOfNonSigningOperators[k]
        //                 ) {
        //                     // if the operator was a non-signer, then we set the flag to false
        //                     wasSigningOperator == false;
        //                     break;
        //                 }
        //             }

        //             if (wasSigningOperator == true) {
        //                 BLSRegistryCoordinatorWithIndices(
        //                     address(registryCoordinator)
        //                 ).serviceManager().freezeOperator(operatorAddress);
        //             }
        //         }
        //     }
        // }

        // the task response has been challenged successfully
        $.taskSuccesfullyChallenged[referenceTaskId] = true;

        emit TaskChallengedSuccessfully(referenceTaskId, msg.sender);
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

    /**
     * @notice Get the balance an EigenLayer operator has staked with the AVS
     * @param operator Address of the operator
     * @return Amount staked by the operator
     */
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
        return $.latestTaskId;
    }
    function getTaskResponseWindowBlock()
        external
        view
        override
        returns (uint32)
    {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        return $.TASK_RESPONSE_WINDOW_BLOCK;
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
    /**
     * @notice Ensures that a validator is registered. If not, registers it.
     * @param validatorAddr The address of the validator in string format.
     */
    function _ensureValidatorRegistered(
        string memory validatorAddr
    ) internal virtual {
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        // If validator is not registered, register it
        if (_validatorExists(validatorHash) == false) {
            _registerValidator(validatorAddr);
        }
    }

    function _registerValidator(
        string memory validatorAddr
    ) internal virtual returns (ZrServiceManagerLib.Validator memory) {
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));

        // Ensure the validator isn't already registered
        require(
            _validatorExists(validatorHash) == false,
            "Validator already registered"
        );

        // Initialize the Validator struct with the operator included directly in the array
        ZrServiceManagerLib.Validator memory validator = ZrServiceManagerLib
            .Validator({
                validatorAddr: validatorAddr,
                validatorHash: validatorHash,
                operators: new address[](0)
            });

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        // Store the validator in the mapping
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
        require(_validatorExists(validatorHash), "unkown validator hash");

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

    /**
     * @notice Deregisters an operator by removing it from `_allOperators` and undelegating it from its validator.
     * @param oprAddr The address of the operator to deregister.
     */
    function _deRegisterOperator(address oprAddr) internal virtual {
        // Check that the operator address is valid
        require(oprAddr != address(0), "zero operator address");

        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        bytes32 validatorHash = $.operatorToValidator[oprAddr];
        require(validatorHash != bytes32(0), "invalid validator hash");

        uint256 index = $.operatorToValidatorIndex[oprAddr];

        ZrServiceManagerLib.Validator storage validator = $._validators[
            validatorHash
        ];

        // Verify the index is within bounds
        require(index < validator.operators.length, "Index out of bounds");

        // Ensure that the operator at the given index matches the provided operator address
        require(
            validator.operators[index] == oprAddr,
            "Operator address mismatch"
        );

        // Retrieve the last operator index from the array
        uint256 lastOperatorIndex = validator.operators.length - 1;
        // Retrieve the last operator in the array to replace the removed one
        address lastOperator = validator.operators[
            validator.operators.length - 1
        ];
        if (index != lastOperatorIndex) {
            validator.operators[index] = lastOperator;
        }
        validator.operators.pop();

        emit ValidatorDeregistered(validatorHash);
    }

    function _deRegisterValidator(bytes32 validatorHash) internal virtual {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();

        // Retrieve the validator struct
        ZrServiceManagerLib.Validator storage validator = $._validators[
            validatorHash
        ];

        // Ensure the validator exists
        if (validator.validatorHash != bytes32(0)) {
            // Iterate through the operators associated with this validator
            for (uint256 i = 0; i < validator.operators.length; i++) {
                _ejectOperator(validator.operators[i]);
                emit OperatorDeregistered(
                    validator.operators[i],
                    validatorHash
                );
            }

            // Remove the validator from the mapping
            delete $._validators[validatorHash];
        }
    }

    function _ejectOperator(address operatorAddr) internal virtual {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        bytes32 validatorHash = $.operatorToValidator[operatorAddr];
        // Ensure the operator's validatorHash matches before deregistration
        if (validatorHash != bytes32(0)) {
            uint8[] memory quorumNumbers = new uint8[](1);
            quorumNumbers[0] = QUORUM_NUMBER;
            _registryCoordinator.ejectOperator(
                operatorAddr,
                _encodeQuorumNumbers(quorumNumbers)
            );
        }
    }

    /**
     * @notice Checks if a validator exists in the `_validators` mapping.
     * @param validatorHash The hash of the validator address.
     * @return bool indicating whether the validator exists.
     */
    function _validatorExists(
        bytes32 validatorHash
    ) internal view virtual returns (bool) {
        return bytes(_getValidator(validatorHash).validatorAddr).length != 0;
    }

    /**
     * @notice Retrieves a `Validator` struct from `_validators` mapping.
     * @dev Fetches the validator struct associated with a specific hash. This function is used
     * internally to access all details of a validator in the mapping.
     * @param validatorHash The unique hash representing the validator's address.
     * @return A `Validator` struct containing the details of the requested validator.
     */
    function _getValidator(
        bytes32 validatorHash
    ) internal view virtual returns (ZrServiceManagerLib.Validator memory) {
        ZrServiceManagerStorage storage $ = _getZrServiceManagerStorage();
        return $._validators[validatorHash];
    }
}
