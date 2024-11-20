#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
VALIDATOR_ADDR="zenvaloper126hek6zagmp3jqf97x7pq7c0j9jqs0ndvcepy6"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

STORAGE_LOCATION="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"
VALIDATOR_HASH=$(cast keccak256 $(cast --from-utf8 "$VALIDATOR_ADDR"))

echo "==============================================="
echo "Focused Storage Debug"
echo "==============================================="

# Check the struct slots directly
echo "Direct Struct Storage"
echo "--------------------------------"
echo "Storage location: $STORAGE_LOCATION"
echo "Direct value: $(cast storage $CONTRACT_ADDRESS $STORAGE_LOCATION --rpc-url $RPC_URL)"

# Check the validators mapping
echo -e "\nValidators Mapping (First Field)"
echo "--------------------------------"
VALIDATORS_SLOT=$(cast keccak256 abi:0x0000000000000000000000000000000000000000000000000000000000000000$STORAGE_LOCATION)
echo "Validators base slot: $VALIDATORS_SLOT"

# Get the specific validator slot
VALIDATOR_SLOT=$(cast keccak256 abi:$VALIDATOR_HASH$VALIDATORS_SLOT)
echo "Specific validator slot: $VALIDATOR_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $VALIDATOR_SLOT --rpc-url $RPC_URL)"

# Check operatorToValidator mapping (Second Field)
echo -e "\nOperatorToValidator Mapping"
echo "--------------------------------"
OPERATOR="0xba68163d4edcca7500844b1c3f1f8e4f5d547066"
OPERATOR_MAP_SLOT=$(cast keccak256 abi:0x0000000000000000000000000000000000000000000000000000000000000001$STORAGE_LOCATION)
OPERATOR_SLOT=$(cast keccak256 abi:$OPERATOR$OPERATOR_MAP_SLOT)
echo "Operator mapping slot: $OPERATOR_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $OPERATOR_SLOT --rpc-url $RPC_URL)"

# Check allValidators array length (Fourth Field)
echo -e "\nAllValidators Array"
echo "--------------------------------"
ALL_VALIDATORS_SLOT=$(cast keccak256 abi:0x0000000000000000000000000000000000000000000000000000000000000003$STORAGE_LOCATION)
echo "Array length slot: $ALL_VALIDATORS_SLOT"
echo "Length: $(cast storage $CONTRACT_ADDRESS $ALL_VALIDATORS_SLOT --rpc-url $RPC_URL)"

# Check the first few array elements if length > 0
for i in {0..2}; do
    ARRAY_ELEMENT_SLOT=$(cast keccak256 abi:$ALL_VALIDATORS_SLOT)
    ELEMENT_SLOT=$(cast keccak256 abi:$(cast --to-hex $i)$ARRAY_ELEMENT_SLOT)
    echo "Array element $i slot: $ELEMENT_SLOT"
    echo "Value: $(cast storage $CONTRACT_ADDRESS $ELEMENT_SLOT --rpc-url $RPC_URL)"
done

echo -e "\n==============================================="