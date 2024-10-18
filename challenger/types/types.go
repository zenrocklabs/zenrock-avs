package types

import (
	"errors"

	cstaskmanager "github.com/zenrocklabs/zenrock-avs/contracts/bindings/ZRTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.ZRTaskManagerITaskResponse
	TaskResponseMetadata      cstaskmanager.ZRTaskManagerITaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
