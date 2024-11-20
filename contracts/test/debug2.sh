#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

echo "==============================================="
echo "Focused Debug of ZrServiceManager"
echo "==============================================="

# Check operators one by one
for OPERATOR in "0xba68163d4edcca7500844b1c3f1f8e4f5d547066" "0x9e9c1a571d0d09737153417be4f3cf0cd45c9368" "0xa36951d284adda3343653bf6dcf0148c08b9ab2f"; do
    echo -e "\nChecking Operator: $OPERATOR"
    echo "--------------------------------"
    
    # Get operator to validator mapping
    SLOT=$(cast keccak256 abi:$(cast keccak256 abi:$OPERATOR$(cast --to-hex $(cast --to-dec 0x000000000000000000000000000000000000000000000000000000000000000001))))
    echo "Operator to validator mapping at slot: $SLOT"
    VALIDATOR_HASH=$(cast storage $CONTRACT_ADDRESS $SLOT --rpc-url $RPC_URL)
    echo "Mapped to validator hash: $VALIDATOR_HASH"

    if [ "$VALIDATOR_HASH" != "0x0000000000000000000000000000000000000000000000000000000000000000" ]; then
        # Get validator data slot
        VALIDATOR_SLOT=$(cast keccak256 abi:$VALIDATOR_HASH)
        echo "Validator data at slot: $VALIDATOR_SLOT"
        VALIDATOR_DATA=$(cast storage $CONTRACT_ADDRESS $VALIDATOR_SLOT --rpc-url $RPC_URL)
        echo "Validator data: $VALIDATOR_DATA"
    fi
done

# Check validator mapping storage
echo -e "\nChecking Validator Storage"
echo "--------------------------------"
VALIDATOR_MAPPING_SLOT="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"
echo "Base slot for _validators mapping: $VALIDATOR_MAPPING_SLOT"

# Try to find any validator data directly
for i in {0..5}; do
    PROBE_SLOT=$(cast keccak256 abi:$(cast --to-hex $(cast --to-dec $i))$VALIDATOR_MAPPING_SLOT)
    echo -e "\nProbing storage slot $i: $PROBE_SLOT"
    DATA=$(cast storage $CONTRACT_ADDRESS $PROBE_SLOT --rpc-url $RPC_URL)
    if [ "$DATA" != "0x0000000000000000000000000000000000000000000000000000000000000000" ]; then
        echo "Found non-zero data: $DATA"
    fi
done

# Try to get the block number of registrations
echo -e "\nChecking Recent Blocks for Events"
echo "--------------------------------"
cast logs $CONTRACT_ADDRESS \
    -e "event ValidatorRegistered(bytes32 indexed validatorHash, string validatorAddr)" \
    --to-block latest --from-block latest-1000 \
    --rpc-url $RPC_URL

echo -e "\nChecking Operator Assignments"
cast logs $CONTRACT_ADDRESS \
    -e "event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash)" \
    --to-block latest --from-block latest-1000 \
    --rpc-url $RPC_URL

echo -e "\n==============================================="
echo "Debug Complete"
echo "==============================================="