#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
VALIDATOR_ADDR="zenvaloper126hek6zagmp3jqf97x7pq7c0j9jqs0ndvcepy6"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

echo "==============================================="
echo "Hash Verification"
echo "==============================================="

# Calculate validator hash
VALIDATOR_HASH=$(cast keccak256 $(cast --from-utf8 "$VALIDATOR_ADDR"))
echo "Expected validator hash: $VALIDATOR_HASH"

# Verify string encoding
ENCODED_STRING_1="0x7a656e76616c6f70657231323668656b367a61676d70336a7166393778377071"
ENCODED_STRING_2="0x3763306a396a7173306e64766365707936"
echo -e "\nString encoding verification:"
echo "Part 1 decoded: $(cast --from-hex $ENCODED_STRING_1)"
echo "Part 2 decoded: $(cast --from-hex $ENCODED_STRING_2)"

# Check storage for validator struct at computed locations
STORAGE_BASE="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"

# Check the mapping slot itself
MAPPING_SLOT=$(cast keccak256 abi:$(cast --concat-hex $STORAGE_BASE "0x0000000000000000000000000000000000000000000000000000000000000000"))
echo -e "\nMapping base slot: $MAPPING_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $MAPPING_SLOT --rpc-url $RPC_URL)"

# Check the validator struct slot
VALIDATOR_SLOT=$(cast keccak256 abi:$(cast --concat-hex $VALIDATOR_HASH $MAPPING_SLOT))
echo -e "\nValidator struct slot: $VALIDATOR_SLOT"
echo "Value: $(cast storage $CONTRACT_ADDRESS $VALIDATOR_SLOT --rpc-url $RPC_URL)"

# Check alternative storage patterns
echo -e "\nChecking alternative storage patterns..."
for i in {0..3}; do
    ALT_SLOT=$(cast keccak256 abi:$(cast --concat-hex $VALIDATOR_HASH $(cast --to-hex $i)))
    echo "Pattern $i slot: $ALT_SLOT"
    echo "Value: $(cast storage $CONTRACT_ADDRESS $ALT_SLOT --rpc-url $RPC_URL)"
done

echo -e "\n==============================================="