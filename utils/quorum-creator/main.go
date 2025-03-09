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

	// contract address is our AVS' RegistryCoordinator proxy (not implementation)
	stakeReg, err := NewContractZrRegistryCoordinator(common.HexToAddress("0xFbFECE8f29f499c32206d8bFfA57da2b124790C7"), ethRpcClient)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("key.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	opts, err := bind.NewTransactorWithChainID(file, "", big.NewInt(1))
	if err != nil {
		panic(err)
	}

	// contract address is EigenLayer's WETH StrategyBase proxy (not implementation)
	if _, err := stakeReg.CreateQuorum(
		opts,
		IRegistryCoordinatorOperatorSetParam{
			MaxOperatorCount:        64,
			KickBIPsOfOperatorStake: 500,
			KickBIPsOfTotalStake:    500,
		},
		big.NewInt(0),
		[]IStakeRegistryStrategyParams{{
			Strategy:   common.HexToAddress("0xa5430Ca83713F877B77b54d5A24FD3D230DF854B"), // proxy
			Multiplier: big.NewInt(1000000000000000000),
		}},
	); err != nil {
		panic(err)
	}

	fmt.Println("Successfully created quorum")
}
