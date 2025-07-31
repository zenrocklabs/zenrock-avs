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

	if _, err := stakeReg.SetEjector(
		opts,
		common.HexToAddress("0x4ca852BD78D9B7295874A7D223023Bff011b7EB3"), // AVS ServiceManager proxy
	); err != nil {
		panic(err)
	}

	fmt.Println("Successfully set ejector")
}
