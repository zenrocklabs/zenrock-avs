package main

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func TestEstimateGasForAddStrategies(t *testing.T) {
	// Setup the Ethereum client
	ethRpcClient, err := eth.NewClient("https://rpc.ankr.com/eth") // add RPC URL here
	if err != nil {
		t.Fatalf("Failed to create Ethereum client: %v", err)
	}

	// Initialize the StakeRegistry contract
	stakeReg, err := NewContractStakeRegistry(common.HexToAddress("0x68BD5a90EEcF77E96d93cf593d5e7b361793c8E2"), ethRpcClient)
	if err != nil {
		t.Fatalf("Failed to initialize StakeRegistry contract: %v", err)
	}

	// Open the key file
	file, err := os.Open("key.json")
	if err != nil {
		t.Fatalf("Failed to open key file: %v", err)
	}
	defer file.Close()

	// Create transaction options
	opts, err := bind.NewTransactorWithChainID(file, "", big.NewInt(17000))
	if err != nil {
		t.Fatalf("Failed to create transaction options: %v", err)
	}

	// Set NoSend to true to estimate gas without sending the transaction
	opts.NoSend = true
	// Set a specific gas limit to bypass funds check
	opts.GasLimit = 5000000
	// Ensure value is set to 0
	opts.Value = big.NewInt(0)

	// Define the strategy parameters
	strategyParams := []IStakeRegistryStrategyParams{{
		Strategy:   common.HexToAddress("0xa5430Ca83713F877B77b54d5A24FD3D230DF854B"),
		Multiplier: big.NewInt(1000000000000000000),
	}}

	// Estimate gas for adding strategies
	tx, err := stakeReg.AddStrategies(opts, 0, strategyParams)
	if err != nil {
		t.Fatalf("Failed to estimate gas for adding strategies: %v", err)
	}

	// Print the estimated gas
	fmt.Printf("Estimated gas for adding strategies: %d\n", tx.Gas())

	// You can also print other transaction details
	fmt.Printf("Transaction data size: %d bytes\n", len(tx.Data()))

	// Convert wei to ETH (1 ETH = 10^18 wei)
	weiPerEth := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// Gas price in wei and ETH
	gasPriceWei := new(big.Float).SetInt(tx.GasPrice())
	gasPriceEth := new(big.Float).Quo(gasPriceWei, weiPerEth)
	fmt.Printf("Gas price: %s wei (%.18f ETH)\n", tx.GasPrice().String(), gasPriceEth)

	// Estimated cost in wei and ETH
	costWei := new(big.Float).SetInt(new(big.Int).Mul(big.NewInt(int64(tx.Gas())), tx.GasPrice()))
	costEth := new(big.Float).Quo(costWei, weiPerEth)
	fmt.Printf("Estimated cost: %s wei (%.18f ETH)\n", new(big.Int).Mul(big.NewInt(int64(tx.Gas())), tx.GasPrice()).String(), costEth)
}
