package types

import (
	"errors"

	cstaskmanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/TaskManagerZR"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.ITaskManagerZRTaskResponse
	TaskResponseMetadata      cstaskmanager.ITaskManagerZRTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
