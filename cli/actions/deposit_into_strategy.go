package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
	"github.com/zenrocklabs/zenrock-avs/core/config"
	"github.com/zenrocklabs/zenrock-avs/operator"
	"github.com/zenrocklabs/zenrock-avs/types"
)

func DepositIntoStrategy(ctx *cli.Context) error {

	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}

	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig, nil)
	if err != nil {
		return err
	}

	strategyAddrStr := ctx.String("strategy-addr")
	strategyAddr := common.HexToAddress(strategyAddrStr)
	amountStr := ctx.String("amount")
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		fmt.Println("Error converting amount to big.Int")
		return err
	}

	if err = operator.DepositIntoStrategy(strategyAddr, amount); err != nil {
		return err
	}

	return nil
}
