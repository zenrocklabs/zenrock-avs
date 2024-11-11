// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zenrocklabs/zenrock-avs/core/chainio (interfaces: AvsReaderer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avs_reader.go -package=mocks github.com/zenrocklabs/zenrock-avs/core/chainio AvsReaderer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	contractOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	types "github.com/Layr-Labs/eigensdk-go/types"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	contractTaskManagerZR "github.com/zenrocklabs/zenrock-avs/contracts/bindings/TaskManagerZR"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsReaderer is a mock of AvsReaderer interface.
type MockAvsReaderer struct {
	ctrl     *gomock.Controller
	recorder *MockAvsReadererMockRecorder
}

// MockAvsReadererMockRecorder is the mock recorder for MockAvsReaderer.
type MockAvsReadererMockRecorder struct {
	mock *MockAvsReaderer
}

// NewMockAvsReaderer creates a new mock instance.
func NewMockAvsReaderer(ctrl *gomock.Controller) *MockAvsReaderer {
	mock := &MockAvsReaderer{ctrl: ctrl}
	mock.recorder = &MockAvsReadererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsReaderer) EXPECT() *MockAvsReadererMockRecorder {
	return m.recorder
}

// CheckSignatures mocks base method.
func (m *MockAvsReaderer) CheckSignatures(arg0 context.Context, arg1 [32]byte, arg2 []byte, arg3 uint32, arg4 contractTaskManagerZR.IBLSSignatureCheckerNonSignerStakesAndSignature) (contractTaskManagerZR.IBLSSignatureCheckerQuorumStakeTotals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSignatures", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(contractTaskManagerZR.IBLSSignatureCheckerQuorumStakeTotals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSignatures indicates an expected call of CheckSignatures.
func (mr *MockAvsReadererMockRecorder) CheckSignatures(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSignatures", reflect.TypeOf((*MockAvsReaderer)(nil).CheckSignatures), arg0, arg1, arg2, arg3, arg4)
}

// GetCheckSignaturesIndices mocks base method.
func (m *MockAvsReaderer) GetCheckSignaturesIndices(arg0 *bind.CallOpts, arg1 uint32, arg2 types.QuorumNums, arg3 []types.Bytes32) (contractOperatorStateRetriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSignaturesIndices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractOperatorStateRetriever.OperatorStateRetrieverCheckSignaturesIndices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSignaturesIndices indicates an expected call of GetCheckSignaturesIndices.
func (mr *MockAvsReadererMockRecorder) GetCheckSignaturesIndices(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSignaturesIndices", reflect.TypeOf((*MockAvsReaderer)(nil).GetCheckSignaturesIndices), arg0, arg1, arg2, arg3)
}

// GetLatestTaskNumber mocks base method.
func (m *MockAvsReaderer) GetLatestTaskNumber(arg0 context.Context) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestTaskNumber", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestTaskNumber indicates an expected call of GetLatestTaskNumber.
func (mr *MockAvsReadererMockRecorder) GetLatestTaskNumber(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestTaskNumber", reflect.TypeOf((*MockAvsReaderer)(nil).GetLatestTaskNumber), arg0)
}

// GetOperatorAddrsInQuorumsAtCurrentBlock mocks base method.
func (m *MockAvsReaderer) GetOperatorAddrsInQuorumsAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.QuorumNums) ([][]common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorAddrsInQuorumsAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([][]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorAddrsInQuorumsAtCurrentBlock indicates an expected call of GetOperatorAddrsInQuorumsAtCurrentBlock.
func (mr *MockAvsReadererMockRecorder) GetOperatorAddrsInQuorumsAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorAddrsInQuorumsAtCurrentBlock", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorAddrsInQuorumsAtCurrentBlock), arg0, arg1)
}

// GetOperatorFromId mocks base method.
func (m *MockAvsReaderer) GetOperatorFromId(arg0 *bind.CallOpts, arg1 types.Bytes32) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorFromId", arg0, arg1)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorFromId indicates an expected call of GetOperatorFromId.
func (mr *MockAvsReadererMockRecorder) GetOperatorFromId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorFromId", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorFromId), arg0, arg1)
}

// GetOperatorId mocks base method.
func (m *MockAvsReaderer) GetOperatorId(arg0 *bind.CallOpts, arg1 common.Address) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorId", arg0, arg1)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorId indicates an expected call of GetOperatorId.
func (mr *MockAvsReadererMockRecorder) GetOperatorId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorId", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorId), arg0, arg1)
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock mocks base method.
func (m *MockAvsReaderer) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.Bytes32) (map[types.QuorumNum]*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].(map[types.QuorumNum]*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock indicates an expected call of GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock.
func (mr *MockAvsReadererMockRecorder) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsAtBlock mocks base method.
func (m *MockAvsReaderer) GetOperatorsStakeInQuorumsAtBlock(arg0 *bind.CallOpts, arg1 types.QuorumNums, arg2 uint32) ([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtBlock indicates an expected call of GetOperatorsStakeInQuorumsAtBlock.
func (mr *MockAvsReadererMockRecorder) GetOperatorsStakeInQuorumsAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtBlock", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorsStakeInQuorumsAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsAtCurrentBlock mocks base method.
func (m *MockAvsReaderer) GetOperatorsStakeInQuorumsAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.QuorumNums) ([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtCurrentBlock indicates an expected call of GetOperatorsStakeInQuorumsAtCurrentBlock.
func (mr *MockAvsReadererMockRecorder) GetOperatorsStakeInQuorumsAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtCurrentBlock", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorsStakeInQuorumsAtCurrentBlock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock mocks base method.
func (m *MockAvsReaderer) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0 *bind.CallOpts, arg1 types.Bytes32, arg2 uint32) (types.QuorumNums, [][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.QuorumNums)
	ret1, _ := ret[1].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtBlock.
func (mr *MockAvsReadererMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorsStakeInQuorumsOfOperatorAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock mocks base method.
func (m *MockAvsReaderer) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.Bytes32) (types.QuorumNums, [][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].(types.QuorumNums)
	ret1, _ := ret[1].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock.
func (mr *MockAvsReadererMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock), arg0, arg1)
}

// GetQuorumCount mocks base method.
func (m *MockAvsReaderer) GetQuorumCount(arg0 *bind.CallOpts) (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuorumCount", arg0)
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuorumCount indicates an expected call of GetQuorumCount.
func (mr *MockAvsReadererMockRecorder) GetQuorumCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuorumCount", reflect.TypeOf((*MockAvsReaderer)(nil).GetQuorumCount), arg0)
}

// IsOperatorRegistered mocks base method.
func (m *MockAvsReaderer) IsOperatorRegistered(arg0 *bind.CallOpts, arg1 common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOperatorRegistered", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOperatorRegistered indicates an expected call of IsOperatorRegistered.
func (mr *MockAvsReadererMockRecorder) IsOperatorRegistered(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOperatorRegistered", reflect.TypeOf((*MockAvsReaderer)(nil).IsOperatorRegistered), arg0, arg1)
}

// QueryExistingRegisteredOperatorPubKeys mocks base method.
func (m *MockAvsReaderer) QueryExistingRegisteredOperatorPubKeys(arg0 context.Context, arg1, arg2 *big.Int) ([]common.Address, []types.OperatorPubkeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryExistingRegisteredOperatorPubKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].([]types.OperatorPubkeys)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryExistingRegisteredOperatorPubKeys indicates an expected call of QueryExistingRegisteredOperatorPubKeys.
func (mr *MockAvsReadererMockRecorder) QueryExistingRegisteredOperatorPubKeys(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryExistingRegisteredOperatorPubKeys", reflect.TypeOf((*MockAvsReaderer)(nil).QueryExistingRegisteredOperatorPubKeys), arg0, arg1, arg2)
}

// QueryExistingRegisteredOperatorSockets mocks base method.
func (m *MockAvsReaderer) QueryExistingRegisteredOperatorSockets(arg0 context.Context, arg1, arg2 *big.Int) (map[types.Bytes32]types.Socket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryExistingRegisteredOperatorSockets", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[types.Bytes32]types.Socket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryExistingRegisteredOperatorSockets indicates an expected call of QueryExistingRegisteredOperatorSockets.
func (mr *MockAvsReadererMockRecorder) QueryExistingRegisteredOperatorSockets(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryExistingRegisteredOperatorSockets", reflect.TypeOf((*MockAvsReaderer)(nil).QueryExistingRegisteredOperatorSockets), arg0, arg1, arg2)
}
