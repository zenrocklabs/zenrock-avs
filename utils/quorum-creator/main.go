package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	rpcUrl := "YOUR_RPC_URL_HERE"

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RPC: %v", err))
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(fmt.Sprintf("Failed to get chain ID: %v", err))
	}

	fmt.Printf("Connected to network with chain ID: %s\n", chainID.String())

	ethRpcClient, err := eth.NewClient(rpcUrl)
	if err != nil {
		panic(err)
	}

	// ZrRegistryCoordinator contract address
	registryCoordinatorAddress := common.HexToAddress("0xc3D4Ac223Ed66f7D97AbbEb9cbE8B7D59A16BF17")
	stakeReg, err := NewContractZrRegistryCoordinator(registryCoordinatorAddress, ethRpcClient)
	if err != nil {
		panic(err)
	}

	privateKeyHex := "YOUR_PRIVATE_KEY_HERE"
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		panic(fmt.Sprintf("Failed to parse private key: %v", err))
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(fmt.Sprintf("Failed to create transaction options: %v", err))
	}

	// Strategy address
	strategyAddress := common.HexToAddress("0xA0119075188e7add0D885a14981B9EF300865D0c")

	tx, err := stakeReg.CreateQuorum(
		opts,
		IRegistryCoordinatorOperatorSetParam{
			MaxOperatorCount:        64,
			KickBIPsOfOperatorStake: 500,
			KickBIPsOfTotalStake:    500,
		},
		big.NewInt(0),
		[]IStakeRegistryStrategyParams{{
			Strategy:   strategyAddress,
			Multiplier: big.NewInt(1000000000000000000),
		}},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Transaction submitted! Hash: %s\n", tx.Hash().Hex())
	fmt.Println("Successfully created quorum")
}
