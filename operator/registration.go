package operator

// OUTDATED
// This file contains cli functions for registering an operator with the AVS and printing status
// However, all of this functionality has been moved to the plugin/ package
// we are just waiting for eigenlayer-cli to be open sourced so we can completely get rid of this registration functionality in the operator

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
)

func (o *Operator) registerOperatorOnStartup(
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	tokenStrategyAddr common.Address,
) {
	if err := o.RegisterOperatorWithEigenlayer(); err != nil {
		// This error might only be that the operator was already registered with eigenlayer, so we don't want to fatal
		o.logger.Debug("Error registering operator with eigenlayer", "err", err)
	} else {
		o.logger.Infof("Registered operator with eigenlayer")
	}

	// TODO: shouldn't hardcode value here
	amount := big.NewInt(10000000000000000)
	if err := o.DepositIntoStrategy(tokenStrategyAddr, amount); err != nil {
		o.logger.Debug("Error depositing into strategy", "err", err)
	} else {
		o.logger.Infof("Deposited %s into strategy %s", amount, tokenStrategyAddr)
	}

	if err := o.RegisterOperatorWithAvs(operatorEcdsaPrivateKey); err != nil {
		o.logger.Debug("Error registering operator with avs", "err", err)
	} else {
		o.logger.Infof("Registered operator with avs")
	}

	// if err := o.DelegateServiceManager(o.config.OperatorValidatorAddress, amount); err != nil {
	// 	o.logger.Debug("Error delegating via service manager", "err", err)
	// } else {
	// 	o.logger.Infof("Delegated AVS tokens via ZRServiceManager contract")
	// }
}

// func (o *Operator) DelegateServiceManager(validatorAddr string, amount *big.Int) error {
// 	serviceManager, err := avs.NewContractZrServiceManager(common.HexToAddress(o.config.ServiceManagerAddress), o.ethClient)
// 	if err != nil {
// 		o.logger.Fatal("Error creating service manager contract interface", "err", err)
// 		return err
// 	}

// 	txOpts, err := o.avsWriter.TxMgr.GetNoSendTxOpts()
// 	if err != nil {
// 		o.logger.Errorf("Error getting txOpts")
// 		return err
// 	}

// 	tx, err := serviceManager.Delegate(txOpts, validatorAddr, amount)
// 	if err != nil {
// 		o.logger.Debug("Error creating Delegate tx")
// 		return err
// 	}

// 	if _, err = o.avsWriter.TxMgr.Send(context.Background(), tx); err != nil {
// 		return errors.New("failed to send tx with err: " + err.Error())
// 	}

// 	return nil
// }

func (o *Operator) RegisterOperatorWithEigenlayer() error {
	op := eigenSdkTypes.Operator{
		Address:                 o.operatorAddr.String(),
		EarningsReceiverAddress: o.operatorAddr.String(),
	}
	if _, err := o.eigenlayerWriter.RegisterAsOperator(context.Background(), op); err != nil {
		o.logger.Debug("Error registering operator with eigenlayer", "err", err)
		return err
	}
	return nil
}

func (o *Operator) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {
	if _, err := o.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount); err != nil {
		o.logger.Debug("Error depositing into strategy", "err", err)
		return err
	}
	return nil
}

// Registration specific functions
func (o *Operator) RegisterOperatorWithAvs(
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
) error {
	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(0)}
	socket := "Not Needed"
	operatorToAvsRegistrationSigSalt := [32]byte{}
	if _, err := rand.Read(operatorToAvsRegistrationSigSalt[:]); err != nil {
		o.logger.Errorf("Failed to generate random salt: %v", err)
		return err
	}
	curBlockNum, err := o.ethClient.BlockNumber(context.Background())
	if err != nil {
		o.logger.Errorf("Unable to get current block number")
		return err
	}
	curBlock, err := o.ethClient.BlockByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
	if err != nil {
		o.logger.Errorf("Unable to get current block")
		return err
	}
	sigValidForSeconds := int64(1_000_000)
	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time()) + sigValidForSeconds)
	_, err = o.avsWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		context.Background(),
		operatorEcdsaKeyPair, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
		o.blsKeypair, quorumNumbers, socket,
	)
	if err != nil {
		o.logger.Debug("Unable to register operator with avs registry coordinator")
		return err
	}
	o.logger.Infof("Registered operator with avs registry coordinator.")

	return nil
}

// PRINTING STATUS OF OPERATOR: 1
// operator address: 0xa0ee7a142d267c1f36714e4a8f75612f20a79720
// dummy token balance: 0
// delegated shares in dummyTokenStrat: 200
// operator pubkey hash in AVS pubkey compendium (0 if not registered): 0x4b7b8243d970ff1c90a7c775c008baad825893ec6e806dfa5d3663dc093ed17f
// operator is opted in to eigenlayer: true
// operator is opted in to playgroundAVS (aka can be slashed): true
// operator status in AVS registry: REGISTERED
//
//	operatorId: 0x4b7b8243d970ff1c90a7c775c008baad825893ec6e806dfa5d3663dc093ed17f
//	middlewareTimesLen (# of stake updates): 0
//
// operator is frozen: false
type OperatorStatus struct {
	EcdsaAddress string
	// pubkey compendium related
	PubkeysRegistered bool
	G1Pubkey          string
	G2Pubkey          string
	// avs related
	RegisteredWithAvs bool
	OperatorId        string
}

func (o *Operator) PrintOperatorStatus() error {
	fmt.Println("Printing operator status")
	operatorId, err := o.avsReader.GetOperatorId(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		return err
	}
	pubkeysRegistered := operatorId != [32]byte{}
	registeredWithAvs := o.operatorId != [32]byte{}
	operatorStatus := OperatorStatus{
		EcdsaAddress:      o.operatorAddr.String(),
		PubkeysRegistered: pubkeysRegistered,
		G1Pubkey:          o.blsKeypair.GetPubKeyG1().String(),
		G2Pubkey:          o.blsKeypair.GetPubKeyG2().String(),
		RegisteredWithAvs: registeredWithAvs,
		OperatorId:        hex.EncodeToString(o.operatorId[:]),
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(operatorStatusJson))
	return nil
}
