// SPDX-License-Identifier: BUSL
pragma solidity 0.8.21;

import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "./TestZrServiceManager.sol";

contract ZrServiceManageTest is BLSMockAVSDeployer {
    TestZrServiceManager public sm;
    TestZrServiceManager public smImplementation;
    
    address public constant OWNER = address(uint160(uint256(keccak256(abi.encodePacked("owner")))));
    address public constant OPERATOR = address(uint160(uint256(keccak256(abi.encodePacked("operator")))));
    address public constant OPERATOR2 = address(uint160(uint256(keccak256(abi.encodePacked("operator2")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        smImplementation = new TestZrServiceManager(
            avsDirectory,
            rewardsCoordinator,
            IZrRegistryCoordinator(address(registryCoordinator)),
            stakeRegistry
        );

        sm = TestZrServiceManager(
            address(
                new TransparentUpgradeableProxy(
                    address(smImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        sm.initialize.selector,
                        pauserRegistry,
                        OWNER,
                        IZRTaskManager(address(0))
                    )
                )
            )
        );
    }

    function test_ValidatorRegistrationAndRetrieval() public {
        // Setup
        string memory validatorAddr = "validator1";
        
        // Create dummy signature
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature;
        operatorSignature.signature = new bytes(65);
        operatorSignature.salt = bytes32(0);
        operatorSignature.expiry = block.timestamp + 1 days;
        
        // Register operator
        cheats.prank(address(registryCoordinator));
        sm.registerOperatorToAVS(OPERATOR, validatorAddr, operatorSignature);
        
        // Check individual validator retrieval
        bytes32 validatorHash = keccak256(abi.encodePacked(validatorAddr));
        ZrServiceManagerLib.Validator memory validator = sm.getValidator(validatorHash);
        
        assertEq(validator.validatorAddr, validatorAddr, "Individual validator address mismatch");
        assertEq(validator.validatorHash, validatorHash, "Individual validator hash mismatch");
        assertEq(validator.operators.length, 1, "Should have one operator");
        assertEq(validator.operators[0], OPERATOR, "Individual validator operator mismatch");

        // Check getAllValidator
        ZrServiceManagerLib.Validator[] memory allValidators = sm.getAllValidator();
        assertEq(allValidators.length, 1, "Should have one validator in getAllValidator");
        assertEq(allValidators[0].validatorAddr, validatorAddr, "All validators address mismatch");
        assertEq(allValidators[0].validatorHash, validatorHash, "All validators hash mismatch");
        assertEq(allValidators[0].operators.length, 1, "All validators operator count mismatch");
        assertEq(allValidators[0].operators[0], OPERATOR, "All validators operator mismatch");

        // Check getAllOperatorsAddresses
        address[] memory operators = sm.getAllOperatorsAddresses();
        assertEq(operators.length, 1, "Should have one operator address");
        assertEq(operators[0], OPERATOR, "Operator address mismatch in getAllOperatorsAddresses");

        // Register second operator to same validator
        cheats.prank(address(registryCoordinator));
        sm.registerOperatorToAVS(OPERATOR2, validatorAddr, operatorSignature);

        // Check updated validator
        validator = sm.getValidator(validatorHash);
        assertEq(validator.operators.length, 2, "Should have two operators");
        assertTrue(
            (validator.operators[0] == OPERATOR && validator.operators[1] == OPERATOR2) ||
            (validator.operators[0] == OPERATOR2 && validator.operators[1] == OPERATOR),
            "Both operators should be present"
        );
    }

    function test_ValidatorStorageAndRetrieval() public {
        string memory validatorAddr = "validator1";
        
        // Create dummy signature
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature;
        operatorSignature.signature = new bytes(65);
        operatorSignature.salt = bytes32(0);
        operatorSignature.expiry = block.timestamp + 1 days;
        
        // Register operator
        cheats.prank(address(registryCoordinator));
        sm.registerOperatorToAVS(OPERATOR, validatorAddr, operatorSignature);
        
        // Check storage directly
        bytes32[] memory validatorHashes = sm.debugGetAllValidatorHashes();
        assertEq(validatorHashes.length, 1, "Should have one validator hash stored");
        
        bytes32 expectedHash = keccak256(abi.encodePacked(validatorAddr));
        assertEq(validatorHashes[0], expectedHash, "Stored validator hash mismatch");
        
        // Verify the validator can be retrieved by hash
        ZrServiceManagerLib.Validator memory validator = sm.getValidator(expectedHash);
        assertEq(validator.validatorAddr, validatorAddr, "Retrieved validator address mismatch");
    }
}