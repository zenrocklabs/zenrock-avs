package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/urfave/cli"
	"github.com/zenrocklabs/zenrock-avs/types"
)

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
		EnvVar:   "CONFIG",
	}
	EcdsaKeyPasswordFlag = cli.StringFlag{
		Name:     "ecdsa-key-password",
		Required: false,
		Usage:    "Password to decrypt the ecdsa key",
		EnvVar:   "ECDSA_KEY_PASSWORD",
	}
	BlsKeyPasswordFlag = cli.StringFlag{
		Name:     "bls-key-password",
		Required: false,
		Usage:    "Password to decrypt the bls key",
		EnvVar:   "BLS_KEY_PASSWORD",
	}
	OperationFlag = cli.StringFlag{
		Name:     "operation-type",
		Required: true,
		Usage:    "Supported operations: opt-in, deposit",
		EnvVar:   "OPERATION_TYPE",
	}
	StrategyAddrFlag = cli.StringFlag{
		Name:     "strategy-addr",
		Required: false,
		Usage:    "Strategy address for deposit mock tokens, only used for deposit action",
		EnvVar:   "STRATEGY_ADDR",
	}
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		ConfigFileFlag,
		EcdsaKeyPasswordFlag,
		BlsKeyPasswordFlag,
		OperationFlag,
		StrategyAddrFlag,
	}
	app.Name = "credible-squaring-plugin"
	app.Usage = "Credible Squaring Plugin"
	app.Description = "This is used to run one time operations like avs opt-in/opt-out"
	app.Action = plugin
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func plugin(ctx *cli.Context) {
	goCtx := context.Background()

	operationType := ctx.GlobalString(OperationFlag.Name)
	configPath := ctx.GlobalString(ConfigFileFlag.Name)

	avsConfig := types.NodeConfig{}
	err := utils.ReadYamlConfig(configPath, &avsConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(avsConfig)

	ecdsaKeyPassword := ctx.GlobalString(EcdsaKeyPasswordFlag.Name)

	buildClientConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 avsConfig.EthRpcUrl,
		EthWsUrl:                   avsConfig.EthWsUrl,
		RegistryCoordinatorAddr:    avsConfig.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: avsConfig.OperatorStateRetrieverAddress,
		AvsName:                    "incredible-squaring",
		PromMetricsIpPortAddress:   avsConfig.EigenMetricsIpPortAddress,
	}
	logger, _ := logging.NewZapLogger(logging.Development)

	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		avsConfig.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	clients, err := sdkclients.BuildAll(buildClientConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		fmt.Println(err)
		return
	}

	if operationType == "opt-in" {
		blsKeyPassword := ctx.GlobalString(BlsKeyPasswordFlag.Name)

		blsKeypair, err := bls.ReadPrivateKeyFromFile(avsConfig.BlsPrivateKeyStorePath, blsKeyPassword)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Register with registry coordination
		quorumNumbers := sdktypes.QuorumNums{0}
		socket := "Not Needed"
		sigValidForSeconds := int64(1_000_000)
		operatorToAvsRegistrationSigSalt := [32]byte{}
		if _, err := rand.Read(operatorToAvsRegistrationSigSalt[:]); err != nil {
			logger.Errorf("Failed to generate random salt: %v", err)
			return
		}
		operatorToAvsRegistrationSigExpiry := big.NewInt(int64(time.Now().Unix()) + sigValidForSeconds)
		logger.Infof("Registering with registry coordination with quorum numbers %v and socket %s", quorumNumbers, socket)
		r, err := clients.AvsRegistryChainWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
			goCtx,
			operatorEcdsaPrivateKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
			blsKeypair, avsConfig.OperatorValidatorAddress, quorumNumbers, socket,
		)
		if err != nil {
			logger.Errorf("Error assembling CreateNewTask tx: %v", err)
			return
		}
		logger.Infof("Registered with registry coordination successfully with tx hash %s", r.TxHash.Hex())
	} else if operationType == "opt-out" {
		fmt.Println("Opting out of slashing - unimplemented")
	} else {
		fmt.Println("Invalid operation type")
	}
}

func pubKeyG1ToBN254G1Point(p *bls.G1Point) regcoord.BN254G1Point {
	return regcoord.BN254G1Point{
		X: p.X.BigInt(new(big.Int)),
		Y: p.Y.BigInt(new(big.Int)),
	}
}
