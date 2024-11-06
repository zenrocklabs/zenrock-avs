pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./ITaskManagerZR.sol";

contract TaskManagerZR is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ITaskManagerZR
{
    using BN254 for BN254.G1Point;

    /* CONSTANT */
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 10000;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 67;

    /* STORAGE */
    uint32 public latestTaskId;

    // mapping of task indices to all tasks hashes
    mapping(uint32 => bytes32) public allTaskHashes;

    // mapping of task indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allTaskResponses;

    mapping(uint32 => mapping(uint256 => string)) public taskValidatorAddressesByIndex;
    mapping(uint32 => uint256) public taskValidatorCount;

    mapping(uint32 => bool) public taskSuccesfullyChallenged;

    address public aggregator;
    address public generator;

    /* EVENTS */
    event ValidatorAddressesStored(uint32 indexed taskId, uint256 count);

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    modifier onlyTaskGenerator() {
        require(msg.sender == generator, "Task generator must be the caller");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock
    ) BLSSignatureChecker(_registryCoordinator) {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator,
        address _generator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
        generator = _generator;
    }

    function createNewTask(
        uint32 taskId,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyTaskGenerator {
        Task memory newTask;
        newTask.taskId = taskId;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;

        allTaskHashes[taskId] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(taskId, newTask);
        latestTaskId = taskId;
    }

    // Helper function to store validator addresses
    function storeValidatorAddresses(uint32 taskId, string[] calldata addresses) internal {
        uint256 length = addresses.length;
        for(uint256 i = 0; i < length; i++) {
            taskValidatorAddressesByIndex[taskId][i] = addresses[i];
        }
        taskValidatorCount[taskId] = length;
        emit ValidatorAddressesStored(taskId, length);
    }

    // Helper function to retrieve validator addresses
    function getValidatorAddresses(uint32 taskId) internal view returns (string[] memory) {
        uint256 length = taskValidatorCount[taskId];
        string[] memory addresses = new string[](length);
        for(uint256 i = 0; i < length; i++) {
            addresses[i] = taskValidatorAddressesByIndex[taskId][i];
        }
        return addresses;
    }

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;

        // check that the task is valid, hasn't been responded yet, and is being responded in time
        require(
            keccak256(abi.encode(task)) == allTaskHashes[task.taskId],
            "supplied task does not match the one recorded in the contract"
        );
        require(
            allTaskResponses[task.taskId] == bytes32(0),
            "Aggregator has already responded to the task"
        );
        require(
            uint32(block.number) <= taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
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

        // check that signatories own at least a threshold percentage of each quorum
        for (uint i = 0; i < quorumNumbers.length; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] *
                    _THRESHOLD_DENOMINATOR >=
                    quorumStakeTotals.totalStakeForQuorum[i] *
                        uint8(quorumThresholdPercentage),
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }

        // Store the validator addresses using the new storage pattern
        storeValidatorAddresses(task.taskId, taskResponse.activeSetZRChain);

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        // updating the storage with task response
        allTaskResponses[task.taskId] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        emit TaskResponded(taskResponse, taskResponseMetadata);
    }

    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        uint32 referenceTaskId = taskResponse.referenceTaskId;
        
        require(
            allTaskResponses[referenceTaskId] != bytes32(0),
            "Task hasn't been responded to yet"
        );
        require(
            allTaskResponses[referenceTaskId] ==
                keccak256(abi.encode(taskResponse, taskResponseMetadata)),
            "Task response does not match the one recorded in the contract"
        );
        require(
            taskSuccesfullyChallenged[referenceTaskId] == false,
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

        // get the list of hash of pubkeys of operators who weren't part of the task response
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

        // get the address of operators who didn't sign
        address[] memory addresssOfNonSigningOperators = new address[](
            pubkeysOfNonSigningOperators.length
        );
        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            addresssOfNonSigningOperators[i] = BLSApkRegistry(
                address(blsApkRegistry)
            ).pubkeyHashToOperator(hashesOfPubkeysOfNonSigningOperators[i]);
        }

        // the task response has been challenged successfully
        taskSuccesfullyChallenged[referenceTaskId] = true;

        emit TaskChallengedSuccessfully(referenceTaskId, msg.sender);
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskId;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }

    function getActiveSet(uint32 taskId) external view returns (string[] memory) {
        return getValidatorAddresses(taskId);
    }

    function getLatestActiveSet() external view returns (string[] memory) {
        return getValidatorAddresses(latestTaskId);
    }

    function isValidatorInTaskSet(uint32 taskId, string memory validatorAddress) external view returns (bool) {
        uint256 count = taskValidatorCount[taskId];
        for (uint256 i = 0; i < count; i++) {
            if (keccak256(abi.encodePacked(taskValidatorAddressesByIndex[taskId][i])) == 
                keccak256(abi.encodePacked(validatorAddress))) {
                return true;
            }
        }
        return false;
    }
}