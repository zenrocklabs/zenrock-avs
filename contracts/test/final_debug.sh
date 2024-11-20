#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
VALIDATOR_ADDR="zenvaloper126hek6zagmp3jqf97x7pq7c0j9jqs0ndvcepy6"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

echo "==============================================="
echo "Storage Layout Deep Dive"
echo "==============================================="

# Storage slot constants
STORAGE_BASE="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"
VALIDATOR_HASH=$(cast keccak256 $(cast --from-utf8 "$VALIDATOR_ADDR"))

echo "Checking key storage slots..."
echo "--------------------------------"

# 1. Check raw storage slots around the base
echo "Base slot (raw): $(cast storage $CONTRACT_ADDRESS $STORAGE_BASE --rpc-url $RPC_URL)"

# 2. Check the first few slots where _validators mapping might be
for i in {0..3}; do
    SLOT=$(cast keccak256 abi:$(cast --concat-hex $STORAGE_BASE $(cast --to-hex $i)))
    echo "Potential _validators slot $i: $(cast storage $CONTRACT_ADDRESS $SLOT --rpc-url $RPC_URL)"
done

# 3. Check if validator might be stored at a different base
for i in {0..3}; do
    BASE=$(cast --to-hex $i)
    SLOT=$(cast keccak256 abi:$(cast --concat-hex $VALIDATOR_HASH $BASE))
    echo "Alternative validator slot $i: $(cast storage $CONTRACT_ADDRESS $SLOT --rpc-url $RPC_URL)"
done

# 4. Check operators array slots
OPERATORS=("0xba68163d4edcca7500844b1c3f1f8e4f5d547066" "0x9e9c1a571d0d09737153417be4f3cf0cd45c9368" "0xa36951d284adda3343653bf6dcf0148c08b9ab2f")

echo -e "\nChecking operator mappings..."
echo "--------------------------------"

for OPERATOR in "${OPERATORS[@]}"; do
    # Check both possible storage patterns
    SLOT1=$(cast keccak256 abi:$(cast --concat-hex $OPERATOR $STORAGE_BASE))
    SLOT2=$(cast keccak256 abi:$(cast --concat-hex $STORAGE_BASE $OPERATOR))
    
    echo "Operator $OPERATOR"
    echo "Pattern 1: $(cast storage $CONTRACT_ADDRESS $SLOT1 --rpc-url $RPC_URL)"
    echo "Pattern 2: $(cast storage $CONTRACT_ADDRESS $SLOT2 --rpc-url $RPC_URL)"
done

# 5. Look for any non-zero storage in relevant range
echo -e "\nScanning for non-zero storage..."
echo "--------------------------------"
for i in {0..10}; do
    VAL=$(cast storage $CONTRACT_ADDRESS $(cast --to-hex $i) --rpc-url $RPC_URL)
    if [ "$VAL" != "0x0000000000000000000000000000000000000000000000000000000000000000" ]; then
        echo "Found non-zero at slot $i: $VAL"
    fi
done

echo -e "\n==============================================="