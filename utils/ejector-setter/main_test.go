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

func TestEstimateGasForSetEjector(t *testing.T) {
	// Setup the Ethereum client
	ethRpcClient, err := eth.NewClient("") // add RPC URL here
	if err != nil {
		t.Fatalf("Failed to create Ethereum client: %v", err)
	}

	// Initialize the RegistryCoordinator contract
	stakeReg, err := NewContractZrRegistryCoordinator(common.HexToAddress("0xFbFECE8f29f499c32206d8bFfA57da2b124790C7"), ethRpcClient)
	if err != nil {
		t.Fatalf("Failed to initialize RegistryCoordinator contract: %v", err)
	}

	// Open the key file
	file, err := os.Open("key.json")
	if err != nil {
		t.Fatalf("Failed to open key file: %v", err)
	}
	defer file.Close()

	// Create transaction options
	opts, err := bind.NewTransactorWithChainID(file, "", big.NewInt(1))
	if err != nil {
		t.Fatalf("Failed to create transaction options: %v", err)
	}

	// Set NoSend to true to estimate gas without sending the transaction
	opts.NoSend = true
	// Set a specific gas limit to bypass funds check
	opts.GasLimit = 5000000
	// Ensure value is set to 0
	opts.Value = big.NewInt(0)

	// Define the ejector address (AVS ServiceManager proxy)
	ejectorAddress := common.HexToAddress("0x4ca852BD78D9B7295874A7D223023Bff011b7EB3")

	// Estimate gas for setting the ejector
	tx, err := stakeReg.SetEjector(opts, ejectorAddress)
	if err != nil {
		t.Fatalf("Failed to estimate gas for setting ejector: %v", err)
	}

	// Print the estimated gas
	fmt.Printf("Estimated gas for setting ejector: %d\n", tx.Gas())

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
