#!/usr/bin/env bash

# Contract details
CONTRACT_ADDRESS="0x5F7286A313D6F9359CDa18A489286b011E654a20"
RPC_URL="https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/"

echo "==============================================="
echo "Debugging ZrServiceManager at $CONTRACT_ADDRESS"
echo "==============================================="

# 1. Get the storage slot for ZrServiceManagerStorage
STORAGE_SLOT="0xd4228801fe6fc3a8dd60908f229d9fce53db7f61e606c4aa422fc6791e331f00"

echo -e "\n1. Checking Basic Contract State"
echo "--------------------------------"
# Check if contract has code
CODE_SIZE=$(cast code $CONTRACT_ADDRESS --rpc-url $RPC_URL | wc -c)
echo "Contract code size: $CODE_SIZE bytes"

# Get owner
echo "Owner: $(cast call $CONTRACT_ADDRESS "owner()" --rpc-url $RPC_URL)"

echo -e "\n2. Checking Validators Array"
echo "--------------------------------"
# Get the length of _allValidators array
LENGTH_SLOT=$(cast keccak256 abi:$STORAGE_SLOT)
echo "Length slot: $LENGTH_SLOT"
LENGTH=$(cast storage $CONTRACT_ADDRESS $LENGTH_SLOT --rpc-url $RPC_URL)
echo "_allValidators length: $LENGTH"

# If length > 0, get the first few validator hashes
if [ "$LENGTH" != "0x0000000000000000000000000000000000000000000000000000000000000000" ]; then
    # Get the slot where the array data starts
    ARRAY_SLOT=$(cast compute-slot-addr $LENGTH_SLOT 0)
    echo "First validator hash stored at: $ARRAY_SLOT"
    VALIDATOR_HASH=$(cast storage $CONTRACT_ADDRESS $ARRAY_SLOT --rpc-url $RPC_URL)
    echo "First validator hash: $VALIDATOR_HASH"
    
    # Get the validator data
    if [ "$VALIDATOR_HASH" != "0x0000000000000000000000000000000000000000000000000000000000000000" ]; then
        VALIDATOR_SLOT=$(cast keccak256 abi:$VALIDATOR_HASH)
        echo "Validator data stored at: $VALIDATOR_SLOT"
        VALIDATOR_DATA=$(cast storage $CONTRACT_ADDRESS $VALIDATOR_SLOT --rpc-url $RPC_URL)
        echo "Validator data: $VALIDATOR_DATA"
    fi
fi

echo -e "\n3. Checking Direct Function Calls"
echo "--------------------------------"
echo "Calling getAllValidator():"
cast call $CONTRACT_ADDRESS "getAllValidator()" --rpc-url $RPC_URL

echo -e "\nCalling getAllOperatorsAddresses():"
cast call $CONTRACT_ADDRESS "getAllOperatorsAddresses()" --rpc-url $RPC_URL

echo -e "\n4. Recent Transaction Analysis"
echo "--------------------------------"
echo "Last 5 transactions to contract:"
for TX in $(cast recent-txs $CONTRACT_ADDRESS --rpc-url $RPC_URL | head -n 5); do
    echo -e "\nTransaction: $TX"
    echo "Status: $(cast receipt $TX --rpc-url $RPC_URL | grep status)"
    echo "Method: $(cast 4byte-decode $(cast tx $TX --rpc-url $RPC_URL | grep input | cut -d' ' -f2 | cut -c1-10))"
    echo "Gas used: $(cast receipt $TX --rpc-url $RPC_URL | grep gasUsed)"
done

echo -e "\n5. Check Registry Coordinator Info"
echo "--------------------------------"
# Get RegistryCoordinator address
RC_ADDR=$(cast call $CONTRACT_ADDRESS "registryCoordinator()" --rpc-url $RPC_URL 2>/dev/null || echo "Could not get RegistryCoordinator address")
echo "RegistryCoordinator Address: $RC_ADDR"

if [ "$RC_ADDR" != "Could not get RegistryCoordinator address" ]; then
    echo "Checking quorum count:"
    cast call $RC_ADDR "quorumCount()" --rpc-url $RPC_URL
fi

echo -e "\n6. Memory Layout Analysis"
echo "--------------------------------"
# Get raw storage slots to analyze memory layout
for i in {0..5}; do
    echo "Storage slot $i: $(cast storage $CONTRACT_ADDRESS $i --rpc-url $RPC_URL)"
done

echo -e "\n7. Events Analysis"
echo "--------------------------------"
# Get recent events (ValidatorRegistered events)
echo "Recent ValidatorRegistered events:"
cast logs $CONTRACT_ADDRESS --rpc-url $RPC_URL \
    -e "event ValidatorRegistered(bytes32 indexed validatorHash, string validatorAddr)" \
    --max-blocks 10000

echo -e "\nRecent OperatorAssigned events:"
cast logs $CONTRACT_ADDRESS --rpc-url $RPC_URL \
    -e "event OperatorAssigned(address indexed operatorAddr, bytes32 indexed validatorHash)" \
    --max-blocks 10000

echo -e "\n==============================================="
echo "Debug Complete"
echo "==============================================="