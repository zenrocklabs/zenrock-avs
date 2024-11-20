#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
VALIDATOR_ADDR="zenvaloper126hek6zagmp3jqf97x7pq7c0j9jqs0ndvcepy6"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

echo "==============================================="
echo "Storage Layout Debug"
echo "==============================================="

# Calculate validator hash
VALIDATOR_HASH=$(cast keccak256 $(cast --from-utf8 "$VALIDATOR_ADDR"))
echo "Validator hash: $VALIDATOR_HASH"

# ZrServiceManagerStorage location
STORAGE_BASE="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"
echo -e "\nStorage base location: $STORAGE_BASE"

# Calculate storage slot for _validators mapping
# mapping(bytes32 => ZrServiceManagerLib.Validator) _validators
VALIDATOR_MAPPING_SLOT=$(cast keccak256 abi:$VALIDATOR_HASH$STORAGE_BASE)
echo -e "\nValidator mapping slot: $VALIDATOR_MAPPING_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $VALIDATOR_MAPPING_SLOT --rpc-url $RPC_URL)"

# Calculate storage slot for _allValidators array
VALIDATORS_ARRAY_SLOT=$(cast keccak256 abi:$STORAGE_BASE)
echo -e "\nValidators array length slot: $VALIDATORS_ARRAY_SLOT"
echo "Length: $(cast storage $CONTRACT_ADDRESS $VALIDATORS_ARRAY_SLOT --rpc-url $RPC_URL)"

# Check operators array for first validator
VALIDATOR_OPERATORS_SLOT=$(cast keccak256 abi:$(cast --to-hex 2)$VALIDATOR_MAPPING_SLOT)
echo -e "\nValidator operators array slot: $VALIDATOR_OPERATORS_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $VALIDATOR_OPERATORS_SLOT --rpc-url $RPC_URL)"

# Check first operator's mapping to validator
OPERATOR="0xba68163d4edcca7500844b1c3f1f8e4f5d547066"
OPERATOR_TO_VALIDATOR_SLOT=$(cast keccak256 abi:$(cast --concat-hex $OPERATOR $(cast --to-hex 1))$STORAGE_BASE)
echo -e "\nOperator to validator mapping slot: $OPERATOR_TO_VALIDATOR_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $OPERATOR_TO_VALIDATOR_SLOT --rpc-url $RPC_URL)"

# Check the validator's address string storage
STRING_SLOT=$(cast keccak256 abi:$(cast --to-hex 0)$VALIDATOR_MAPPING_SLOT)
echo -e "\nValidator address string slot: $STRING_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $STRING_SLOT --rpc-url $RPC_URL)"

echo -e "\n==============================================="