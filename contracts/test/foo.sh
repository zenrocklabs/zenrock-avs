#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
VALIDATOR_ADDR="zenvaloper126hek6zagmp3jqf97x7pq7c0j9jqs0ndvcepy6"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

STORAGE_LOCATION="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"
VALIDATOR_HASH="0x1671efc53137933f79925d2f24b54fd04ec145e282d58ff5627620c3f4141486"

echo "==============================================="
echo "String Storage Debug"
echo "==============================================="

# Check the validator struct storage
echo "Validator Struct Storage"
echo "--------------------------------"
# Calculate slot for validator struct in _validators mapping
VALIDATORS_SLOT=$(cast keccak256 abi:0x0000000000000000000000000000000000000000000000000000000000000000$STORAGE_LOCATION)
VALIDATOR_STRUCT_SLOT=$(cast keccak256 abi:$VALIDATOR_HASH$VALIDATORS_SLOT)
echo "Validator struct slot: $VALIDATOR_STRUCT_SLOT"

# Check string storage (string is stored at keccak256(slot))
STRING_LOCATION=$(cast keccak256 abi:$VALIDATOR_STRUCT_SLOT)
echo "String storage location: $STRING_LOCATION"
echo "String value at location: $(cast storage $CONTRACT_ADDRESS $STRING_LOCATION --rpc-url $RPC_URL)"

# Check next few slots for string continuation
for i in {1..3}; do
    NEXT_SLOT=$(cast keccak256 abi:$(cast --to-hex $i)$STRING_LOCATION)
    echo "String continuation slot $i: $NEXT_SLOT"
    echo "Value: $(cast storage $CONTRACT_ADDRESS $NEXT_SLOT --rpc-url $RPC_URL)"
done

echo -e "\nVerifying Event Data"
echo "--------------------------------"
# Split validator address into 32-byte chunks for comparison
ADDR_LEN=${#VALIDATOR_ADDR}
BYTES_PER_CHUNK=32
for ((i = 0; i < ADDR_LEN; i += BYTES_PER_CHUNK)); do
    CHUNK="${VALIDATOR_ADDR:$i:$BYTES_PER_CHUNK}"
    echo "Chunk $((i/BYTES_PER_CHUNK)): $CHUNK"
    # Convert to hex for comparison
    echo "Hex: $(cast --from-utf8 "$CHUNK" | cast --to-hex)"
done

echo -e "\n==============================================="