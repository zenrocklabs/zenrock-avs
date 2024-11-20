// TestZrServiceManager.sol
pragma solidity 0.8.21;

import "../src/ZrServiceManager.sol";

contract TestZrServiceManager is ZrServiceManager {
    constructor(
        IAVSDirectory _avsDirectory,
        IRewardsCoordinator _rewardsCoordinator,
        IZrRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry
    ) ZrServiceManager(
        _avsDirectory,
        _rewardsCoordinator,
        _registryCoordinator,
        _stakeRegistry
    ) {}

    function registerOperatorToAVS(
        address operatorAddr,
        string memory validatorAddr,
        ISignatureUtils.SignatureWithSaltAndExpiry memory // Keep param but ignore it
    ) public override onlyRegistryCoordinator {
        _ensureValidatorRegistered(validatorAddr);
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        _assignOperatorToValidator(operatorAddr, validatorHash);
    }
}