package aggregator

import (
	"context"
	"errors"
	"net/http"
	"net/rpc"

	cstaskmanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/ZrTaskManager"
	"github.com/zenrocklabs/zenrock-avs/core"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer(ctx context.Context) error {

	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

// Update the SignedTaskResponse struct to reflect the new task structure
type SignedTaskResponse struct {
	TaskResponse cstaskmanager.IZRTaskManagerTaskResponse
	BlsSignature bls.Signature
	OperatorId   types.OperatorId
}

// Update the ProcessSignedTaskResponse function to handle the new task structure
func (agg *Aggregator) ProcessSignedTaskResponse(signedTaskResponse *SignedTaskResponse, reply *bool) error {
	agg.logger.Infof("Received signed task response: %#v", signedTaskResponse)
	taskId := signedTaskResponse.TaskResponse.ReferenceTaskId
	taskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)
	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}
	agg.taskResponsesMu.Lock()
	if _, ok := agg.taskResponses[taskId]; !ok {
		agg.taskResponses[taskId] = make(map[sdktypes.TaskResponseDigest]cstaskmanager.IZRTaskManagerTaskResponse)
	}
	if _, ok := agg.taskResponses[taskId][taskResponseDigest]; !ok {
		agg.taskResponses[taskId][taskResponseDigest] = signedTaskResponse.TaskResponse
	}
	agg.taskResponsesMu.Unlock()

	err = agg.blsAggregationService.ProcessNewSignature(
		context.Background(), taskId, taskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId,
	)
	return err
}
