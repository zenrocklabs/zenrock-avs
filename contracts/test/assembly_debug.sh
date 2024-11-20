#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
VALIDATOR_ADDR="zenvaloper126hek6zagmp3jqf97x7pq7c0j9jqs0ndvcepy6"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

STORAGE_LOCATION="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"
VALIDATOR_HASH=$(cast keccak256 $(cast --from-utf8 "$VALIDATOR_ADDR"))

echo "==============================================="
echo "Assembly Storage Debug"
echo "==============================================="

# Check the exact storage layout based on assembly
echo "Storage Layout Analysis"
echo "--------------------------------"

# 1. Check the struct slots sequentially from the storage location
for i in {0..5}; do
    SLOT=$(cast keccak256 abi:$(cast --concat-hex $STORAGE_LOCATION $(cast --to-hex $i)))
    echo "Struct slot $i from assembly location: $(cast storage $CONTRACT_ADDRESS $SLOT --rpc-url $RPC_URL)"
done

# 2. Check mapping slots based on assembly location
echo -e "\nMapping Analysis"
echo "--------------------------------"
# Hash of mapping position
MAPPING_POSITION=$(cast keccak256 abi:$(cast --concat-hex $STORAGE_LOCATION "0x0000000000000000000000000000000000000000000000000000000000000000"))
echo "Base mapping position: $MAPPING_POSITION"

# Check validator in mapping
VALIDATOR_POSITION=$(cast keccak256 abi:$(cast --concat-hex $VALIDATOR_HASH $MAPPING_POSITION))
echo "Validator mapping position: $VALIDATOR_POSITION"
echo "Value: $(cast storage $CONTRACT_ADDRESS $VALIDATOR_POSITION --rpc-url $RPC_URL)"

# 3. Check array length
ARRAY_LENGTH_POSITION=$(cast keccak256 abi:$(cast --concat-hex $STORAGE_LOCATION "0x0000000000000000000000000000000000000000000000000000000000000003"))
echo -e "\nArray Analysis"
echo "--------------------------------"
echo "Array length position: $ARRAY_LENGTH_POSITION"
echo "Length value: $(cast storage $CONTRACT_ADDRESS $ARRAY_LENGTH_POSITION --rpc-url $RPC_URL)"

# 4. Check packed storage
echo -e "\nPacked Storage Analysis"
echo "--------------------------------"
for i in {0..10}; do
    PACKED_SLOT=$(cast add-hex $STORAGE_LOCATION $(cast --to-hex $i))
    VAL=$(cast storage $CONTRACT_ADDRESS $PACKED_SLOT --rpc-url $RPC_URL)
    echo "Packed slot $i: $VAL"
done

echo -e "\n==============================================="