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
	ethRpcClient, err := eth.NewClient("https://alien-fittest-dew.ethereum-holesky.quiknode.pro/83464d9ac98adf68666cc53b8f8743ad7309335e/") // add RPC URL here
	if err != nil {
		panic(err)
	}

	// contract address is our AVS' StakeRegistry proxy (not implementation)
	stakeReg, err := NewContractStakeRegistry(common.HexToAddress("0xf12a98238e2532e6911AdaFAdb2c617AD0e1F01f"), ethRpcClient)
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
	if _, err := stakeReg.RemoveStrategies(opts, 0, []*big.Int{big.NewInt(0)}); err != nil {
		panic(err)
	}

	fmt.Println("Successfully removed strategy/strategies")
}
