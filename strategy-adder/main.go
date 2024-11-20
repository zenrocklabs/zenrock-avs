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
	stakeReg, err := NewContractStakeRegistry(common.HexToAddress("0xDf7851e0e793C4B2806A6EA82c162C5198e51Fb3"), ethRpcClient)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("key.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	opts, err := bind.NewTransactorWithChainID(file, "", big.NewInt(17000))
	if err != nil {
		panic(err)
	}

	// contract address is EigenLayer's WETH StrategyBase proxy (not implementation)
	if _, err := stakeReg.AddStrategies(opts, 0, []IStakeRegistryStrategyParams{{
		Strategy:   common.HexToAddress("0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9"), // proxy
		Multiplier: big.NewInt(1000000000000000000),
	}}); err != nil {
		panic(err)
	}

	fmt.Println("Successfully added strategy/strategies")
}
