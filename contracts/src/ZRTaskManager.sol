pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./interfaces/ZRTaskManagerI.sol";

contract ZRTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ZRTaskManagerI
{
    using BN254 for BN254.G1Point;

    /* CONSTANT */
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 10000;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    /* STORAGE */
    uint32 public latestTaskId;

    // mapping of task indices to all tasks hashes
    // when a task is created, task hash is stored here,
    // and responses need to pass the actual task,
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allTaskHashes;

    // mapping of task indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allTaskResponses;

    mapping(uint32 => bool) public taskSuccesfullyChallenged;

    address public aggregator;
    address public generator;

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    // onlyTaskGenerator is used to restrict createNewTask from only being called by a permissioned entity
    // in a real world scenario, this would be removed by instead making createNewTask a payable function
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
        int64 zrChainBlockHeight,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyTaskGenerator {
        Task memory newTask;
        newTask.taskId = taskId;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.zrChainBlockHeight = zrChainBlockHeight;
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;

        allTaskHashes[taskId] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(taskId, newTask);
        latestTaskId = taskId;
    }

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;

        // check that the task is valid, hasn't been responsed yet, and is being responsed in time
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

    function taskNumber() external view returns (uint32) {
        return latestTaskId;
    }

    // NOTE: this function enables a challenger to raise and resolve a challenge.
    // TODO: require challenger to pay a bond for raising a challenge
    // TODO(samlaf): should we check that quorumNumbers is same as the one recorded in the task?
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

        // get the list of hash of pubkeys of operators who weren't part of the task response submitted by the aggregator
        bytes32[] memory hashesOfPubkeysOfNonSigningOperators = new bytes32[](
            pubkeysOfNonSigningOperators.length
        );
        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            hashesOfPubkeysOfNonSigningOperators[i] = pubkeysOfNonSigningOperators[i].hashG1Point();
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
        taskSuccesfullyChallenged[referenceTaskId] = true;

        emit TaskChallengedSuccessfully(referenceTaskId, msg.sender);
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }
}