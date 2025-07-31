package types

import (
	"errors"

	cstaskmanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/ZrTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.ZrServiceManagerLibTaskResponse
	TaskResponseMetadata      cstaskmanager.ZrServiceManagerLibTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
