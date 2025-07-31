package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	ethRpcClient, err := eth.NewClient("") // add RPC URL here
	if err != nil {
		panic(err)
	}

	// contract address is our AVS' StakeRegistry proxy (not implementation)
	stakeReg, err := NewContractStakeRegistry(common.HexToAddress("0x68BD5a90EEcF77E96d93cf593d5e7b361793c8E2"), ethRpcClient)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("key2.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	opts, err := bind.NewTransactorWithChainID(file, "", big.NewInt(1))
	if err != nil {
		panic(err)
	}

	// contract address is EigenLayer's WETH StrategyBase proxy (not implementation)
	if _, err := stakeReg.AddStrategies(opts, 0, []IStakeRegistryStrategyParams{{
		Strategy:   common.HexToAddress("0xa5430Ca83713F877B77b54d5A24FD3D230DF854B"), // proxy
		Multiplier: big.NewInt(1000000000000000000),
	}}); err != nil {
		panic(err)
	}

	fmt.Println("Successfully added strategy/strategies")
}
