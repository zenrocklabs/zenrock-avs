package actions

import (
	"encoding/json"
	"log"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/urfave/cli"
	"github.com/zenrocklabs/zenrock-avs/core/config"
	"github.com/zenrocklabs/zenrock-avs/operator"
	"github.com/zenrocklabs/zenrock-avs/types"
)

func DeregisterOperatorWithAvs(ctx *cli.Context) error {

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

	if err = operator.RegistrationSetup(); err != nil {
		return err
	}

	if err = operator.DeregisterOperatorWithAvs(); err != nil {
		return err
	}

	return nil
}
