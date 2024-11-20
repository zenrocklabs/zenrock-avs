#!/usr/bin/env bash

CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
REGISTRY_COORDINATOR="0xC74d923C096529Bb1334D84B68a5E1b4Ffdb88c5"
VALIDATOR_ADDR="zenvaloper126hek6zagmp3jqf97x7pq7c0j9jqs0ndvcepy6"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

echo "==============================================="
echo "Validator Registration Debug"
echo "==============================================="

# Calculate validator hash
VALIDATOR_HASH=$(cast keccak256 $(cast --from-utf8 "$VALIDATOR_ADDR"))
echo "Expected validator hash for $VALIDATOR_ADDR:"
echo $VALIDATOR_HASH

# Get validator data directly
echo -e "\nChecking Validator Mapping"
echo "--------------------------------"
STORAGE_SLOT="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"
VALIDATOR_SLOT=$(cast keccak256 abi:$VALIDATOR_HASH$STORAGE_SLOT)
echo "Validator slot: $VALIDATOR_SLOT"
VALIDATOR_DATA=$(cast storage $CONTRACT_ADDRESS $VALIDATOR_SLOT --rpc-url $RPC_URL)
echo "Validator data: $VALIDATOR_DATA"

# Check transaction
echo -e "\nAnalyzing Registration Transaction"
echo "--------------------------------"
TX="0xc9c58ffebd006b42cd771c6f99044202b028de3d960b023517c445469796aaf9"
echo "Transaction trace:"
cast trace $TX --rpc-url $RPC_URL

# Check all validators array
echo -e "\nChecking All Validators Array"
echo "--------------------------------"
VALIDATORS_LENGTH_SLOT=$(cast keccak256 abi:$STORAGE_SLOT)
echo "Validators length slot: $VALIDATORS_LENGTH_SLOT"
LENGTH=$(cast storage $CONTRACT_ADDRESS $VALIDATORS_LENGTH_SLOT --rpc-url $RPC_URL)
echo "Validators length: $LENGTH"

# Check recent blocks for events
echo -e "\nChecking Events (last 100 blocks)"
echo "--------------------------------"
CURRENT_BLOCK=$(cast block-number --rpc-url $RPC_URL)
FROM_BLOCK=$((CURRENT_BLOCK - 100))

echo "Validator Registration Events:"
cast logs $CONTRACT_ADDRESS \
    -e "event ValidatorRegistered(bytes32 indexed validatorHash, string validatorAddr)" \
    --from-block $FROM_BLOCK \
    --rpc-url $RPC_URL

echo -e "\nOperator Assignment Events:"
cast logs $CONTRACT_ADDRESS \
    -e "event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash)" \
    --from-block $FROM_BLOCK \
    --rpc-url $RPC_URL

# Check specific storage layout
echo -e "\nDetailed Storage Analysis"
echo "--------------------------------"
# Get storage around the validator mapping
for i in {0..2}; do
    # Calculate slot for validator mapping
    SLOT=$(cast keccak256 abi:$(cast --to-hex $i)$STORAGE_SLOT)
    echo "Storage at calculated slot $i: $(cast storage $CONTRACT_ADDRESS $SLOT --rpc-url $RPC_URL)"
done

echo -e "\n==============================================="
echo "Debug Complete"
echo "==============================================="