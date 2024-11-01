// SPDX-License-Identifier: BSL
pragma solidity >=0.8.2 <0.9.0;

import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./interfaces/ITaskManagerZR.sol";

/**
 * @title ZRServiceManager
 * @notice This contract manages delegations from EigenLayer operators to validators on another blockchain.
 * @dev It implements a waiting period for undelegations and ensures that delegations don't exceed EigenLayer stakes.
 */
contract ZRServiceManager is
    ServiceManagerBase,
    ReentrancyGuard
{
    ITaskManagerZR public immutable taskManagerZR;
    uint8 private immutable QUORUM_NUMBER;
    uint256 public immutable UNDELEGATION_PERIOD;

    /**
     * @notice Struct to represent a delegation
     * @param balance Current active balance
     * @param pendingBalance Balance after pending changes
     * @param pendingBalanceActivationHeight Block number when pending balance becomes active
     */
    struct Delegation {
        uint256 balance;
        uint256 pendingBalance;
        uint256 pendingBalanceActivationHeight;
    }

    // Mapping of validator hash to operator (delegator) address to delegation details
    mapping(bytes32 => mapping(address => Delegation)) public delegations;
    // Mapping of operator address to their delegated validator hash
    mapping(address => bytes32) private _operatorValidator;
    // Mapping to store the original validator addresses
    mapping(bytes32 => string) public validatorAddresses;
    // Mapping of operator address to their index in _allOperators array
    mapping(address => uint256) private _operatorIndex;
    // Mapping to track if an operator exists in _allOperators
    mapping(address => bool) private _operatorExists;
    // Array of all operator addresses
    address[] private _allOperators;

    // Events
    event Delegated(address indexed operator, bytes32 validatorHash, uint256 amount);
    event UndelegationInitiated(address indexed operator, bytes32 validatorHash, uint256 amount, uint256 undelegationBlock);
    event UndelegationAnnulled(address indexed operator, bytes32 validatorHash, uint256 amount);
    event OperatorRemoved(address indexed operator, bytes32 validatorHash);

    /**
     * @notice Constructor to initialize the contract
     * @param _avsDirectory Address of the AVS Directory contract
     * @param _registryCoordinator Address of the Registry Coordinator contract
     * @param _stakeRegistry Address of the Stake Registry contract
     * @param _taskManagerZR Address of the ZRTaskManager contract
     * @param _undelegationPeriod Duration of the undelegation waiting period in blocks
     * @param _quorumNumber Used to identify the AVS quorum within the registries
     */
    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        ITaskManagerZR _taskManagerZR,
        uint256 _undelegationPeriod,
        uint8 _quorumNumber
    ) 
        ServiceManagerBase(
            _avsDirectory,
            IPaymentCoordinator(address(0)), // support for rewards will be added later
            _registryCoordinator,
            _stakeRegistry
        )
    {
        taskManagerZR = _taskManagerZR;
        UNDELEGATION_PERIOD = _undelegationPeriod;
        QUORUM_NUMBER = _quorumNumber;
    }

    /**
     * @notice Modifier to update the balance of a delegation if the waiting period has passed
     * @param validatorHash Hash of the validator address
     * @param operator Address of the operator
     */
    modifier updateBalance(bytes32 validatorHash, address operator) {
        Delegation storage delegation = delegations[validatorHash][operator];
        if (block.number >= delegation.pendingBalanceActivationHeight && delegation.pendingBalance != delegation.balance) {
            delegation.balance = delegation.pendingBalance;
            delegation.pendingBalanceActivationHeight = 0;

            // Remove operator if balance is zero
            if (delegation.balance == 0) {
                _removeOperator(operator, validatorHash);
            }
        }
        _;
    }

    /**
     * @notice Modifier to ensure the caller is delegated to the specified validator
     * @param validatorHash Hash of the validator address
     */
    modifier onlyDelegatedValidator(bytes32 validatorHash) {
        require(_operatorValidator[msg.sender] == validatorHash, "Not delegated to this validator");
        _;
    }
 
    /**
     * @notice Get the current balance of a delegation
     * @param validatorHash Hash of the validator address
     * @param operator Address of the operator
     * @return Current balance of the delegation
     */
    function getCurrentBalance(bytes32 validatorHash, address operator) public view returns (uint256) {
        Delegation storage delegation = delegations[validatorHash][operator];
        if (block.number >= delegation.pendingBalanceActivationHeight && delegation.pendingBalanceActivationHeight != 0) {
            return delegation.pendingBalance;
        }
        return delegation.balance;
    }

    /**
     * @notice Get the balance an EigenLayer operator has staked with the AVS
     * @param operator Address of the operator
     * @return Amount staked by the operator
     */
    function getEigenStake(address operator, uint8 quorumNumber) public view returns (uint96) {
        return IStakeRegistry(address(_stakeRegistry)).getCurrentStake(
            _registryCoordinator.getOperator(operator).operatorId,
            quorumNumber
        );
    }

    /**
     * @notice Delegate tokens to a zrChain validator
     * @param validatorAddress String address of the validator to delegate to
     * @param amount Amount of tokens to delegate
     */
    function delegate(string memory validatorAddress, uint256 amount) external nonReentrant updateBalance(keccak256(abi.encodePacked(validatorAddress)), msg.sender) {
        require(amount > 0, "Amount must be greater than zero");

        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddress));

        // Store the original validator address if not already stored
        if (bytes(validatorAddresses[validatorHash]).length == 0) {
            validatorAddresses[validatorHash] = validatorAddress;
        }

        // Ensure operator is only delegated to one validator
        if (_operatorValidator[msg.sender] != bytes32(0)) {
            require(_operatorValidator[msg.sender] == validatorHash,
                    "Already delegated to a different validator");
        }

        Delegation storage delegation = delegations[validatorHash][msg.sender];
        uint256 newTotalAmount = delegation.balance + amount;

        // Ensure total delegation doesn't exceed EigenLayer stake
        require(newTotalAmount <= getEigenStake(msg.sender, QUORUM_NUMBER), "Total delegation cannot exceed EigenLayer stake");

        _operatorValidator[msg.sender] = validatorHash;

        // Add operator to _allOperators if not already present
        if (!_operatorExists[msg.sender]) {
            _operatorIndex[msg.sender] = _allOperators.length;
            _allOperators.push(msg.sender);
            _operatorExists[msg.sender] = true; // Mark operator as existing
        }

        delegation.balance = newTotalAmount;

        // Reset pending balance if there was an ongoing undelegation
        if (delegation.pendingBalanceActivationHeight != 0) {
            delegation.pendingBalance = delegation.balance;
            delegation.pendingBalanceActivationHeight = 0;
        }

        emit Delegated(msg.sender, validatorHash, amount);
    }

    /**
     * @notice Initiate undelegation of tokens from a validator
     * @param validatorAddress String address of the validator to undelegate from
     * @param amount Amount of tokens to undelegate
     */
    function undelegate(string memory validatorAddress, uint256 amount) external nonReentrant updateBalance(keccak256(abi.encodePacked(validatorAddress)), msg.sender) {
        require(amount > 0, "Amount must be greater than zero");

        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddress));

        require(_operatorValidator[msg.sender] == validatorHash, "Not delegated to this validator");

        Delegation storage delegation = delegations[validatorHash][msg.sender];
        require(delegation.balance >= amount, "Insufficient balance to undelegate");

        // Ensure remaining delegation doesn't exceed EigenLayer stake
        require(delegation.balance - amount <= getEigenStake(msg.sender, QUORUM_NUMBER), "Undelegation would exceed EigenLayer stake");

        // Set pending balance to reflect the undelegation
        delegation.pendingBalance = delegation.balance - amount;
        delegation.pendingBalanceActivationHeight = block.number + UNDELEGATION_PERIOD;

        emit UndelegationInitiated(msg.sender, validatorHash, amount, delegation.pendingBalanceActivationHeight);
    }

    function completeUndelegate() external nonReentrant updateBalance(keccak256(abi.encodePacked(validatorAddress)), msg.sender){}

    /**
     * @notice Cancel a pending undelegation
     * @param validatorAddress String address of the validator
     * @param amount Amount of tokens to re-delegate
     */
    function annulUndelegation(string memory validatorAddress, uint256 amount) external nonReentrant updateBalance(keccak256(abi.encodePacked(validatorAddress)), msg.sender) {
        require(amount > 0, "Amount must be greater than zero");

        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddress));

        require(_operatorValidator[msg.sender] == validatorHash, "Not delegated to this validator");

        Delegation storage delegation = delegations[validatorHash][msg.sender];
        require(delegation.pendingBalanceActivationHeight > block.number, "No valid undelegation to cancel");
        require(delegation.pendingBalance < delegation.balance, "No pending undelegation to annul");
        uint256 undelegationAmount = delegation.balance - delegation.pendingBalance;
        require(amount <= undelegationAmount, "Cannot annul more than pending undelegation");

        // Ensure total delegation doesn't exceed EigenLayer stake
        require(delegation.pendingBalance + amount <= getEigenStake(msg.sender, QUORUM_NUMBER), "Annulment would exceed EigenLayer stake");

        // Adjust pending balance to reflect the annulled undelegation
        delegation.pendingBalance += amount;

        // Reset pending activation height if all undelegation is annulled
        if (delegation.pendingBalance == delegation.balance) {
            delegation.pendingBalanceActivationHeight = 0;
        }

        emit UndelegationAnnulled(msg.sender, validatorHash, amount);
    }

    function getCurrentValidatorSet() public view returns(string[] memory){
        return zrTaskManager.getLatestActiveSet();
    }

    // Helper functions used by the validator sidecar

    /**
     * @notice Get all operators who have delegated
     * @return Array of all operator addresses
     */
    function getAllOperators() external view returns (address[] memory) {
        return _allOperators;
    }

    /**
     * @notice Get delegation details for an operator
     * @param operator Address of the operator
     * @return Validator address and current balance
     */
    function getDelegationsForOperator(address operator) external view returns (string memory, uint256) {
        bytes32 validatorHash = _operatorValidator[operator];
        require(validatorHash != bytes32(0), "Operator not delegated");

        string memory validatorAddress = validatorAddresses[validatorHash];

        return (validatorAddress, getCurrentBalance(validatorHash, operator));
    }

    /**
     * @notice Internal function to remove an operator when their balance becomes zero
     * @param operator Address of the operator
     * @param validatorHash Hash of the validator address
     */
    function _removeOperator(address operator, bytes32 validatorHash) internal {
        uint256 indexToRemove = _operatorIndex[operator];
        address lastOperator = _allOperators[_allOperators.length - 1];

        // Swap and remove the operator
        _allOperators[indexToRemove] = lastOperator;
        _operatorIndex[lastOperator] = indexToRemove;
        _allOperators.pop();

        // Clean up mappings
        delete _operatorIndex[operator];
        delete _operatorExists[operator];
        delete _operatorValidator[operator];
        delete delegations[validatorHash][operator];

        emit OperatorRemoved(operator, validatorHash);
    }

    /**
     * @notice Get the original validator address from its hash
     * @param validatorHash Hash of the validator address
     * @return The original validator address string
     */
    function getValidatorAddress(bytes32 validatorHash) external view returns (string memory) {
        return validatorAddresses[validatorHash];
    }
}
