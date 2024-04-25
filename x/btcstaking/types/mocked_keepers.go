// Code generated by MockGen. DO NOT EDIT.
// Source: x/btcstaking/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	big "math/big"
	reflect "reflect"

	types "github.com/babylonchain/babylon/types"
	types0 "github.com/babylonchain/babylon/x/btccheckpoint/types"
	types1 "github.com/babylonchain/babylon/x/btclightclient/types"
	types2 "github.com/babylonchain/babylon/x/epoching/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBTCLightClientKeeper is a mock of BTCLightClientKeeper interface.
type MockBTCLightClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBTCLightClientKeeperMockRecorder
}

// MockBTCLightClientKeeperMockRecorder is the mock recorder for MockBTCLightClientKeeper.
type MockBTCLightClientKeeperMockRecorder struct {
	mock *MockBTCLightClientKeeper
}

// NewMockBTCLightClientKeeper creates a new mock instance.
func NewMockBTCLightClientKeeper(ctrl *gomock.Controller) *MockBTCLightClientKeeper {
	mock := &MockBTCLightClientKeeper{ctrl: ctrl}
	mock.recorder = &MockBTCLightClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBTCLightClientKeeper) EXPECT() *MockBTCLightClientKeeperMockRecorder {
	return m.recorder
}

// GetBaseBTCHeader mocks base method.
func (m *MockBTCLightClientKeeper) GetBaseBTCHeader(ctx context.Context) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseBTCHeader", ctx)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetBaseBTCHeader indicates an expected call of GetBaseBTCHeader.
func (mr *MockBTCLightClientKeeperMockRecorder) GetBaseBTCHeader(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseBTCHeader", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetBaseBTCHeader), ctx)
}

// GetHeaderByHash mocks base method.
func (m *MockBTCLightClientKeeper) GetHeaderByHash(ctx context.Context, hash *types.BTCHeaderHashBytes) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByHash", ctx, hash)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash.
func (mr *MockBTCLightClientKeeperMockRecorder) GetHeaderByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetHeaderByHash), ctx, hash)
}

// GetTipInfo mocks base method.
func (m *MockBTCLightClientKeeper) GetTipInfo(ctx context.Context) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTipInfo", ctx)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetTipInfo indicates an expected call of GetTipInfo.
func (mr *MockBTCLightClientKeeperMockRecorder) GetTipInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTipInfo", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetTipInfo), ctx)
}

// MockBtcCheckpointKeeper is a mock of BtcCheckpointKeeper interface.
type MockBtcCheckpointKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBtcCheckpointKeeperMockRecorder
}

// MockBtcCheckpointKeeperMockRecorder is the mock recorder for MockBtcCheckpointKeeper.
type MockBtcCheckpointKeeperMockRecorder struct {
	mock *MockBtcCheckpointKeeper
}

// NewMockBtcCheckpointKeeper creates a new mock instance.
func NewMockBtcCheckpointKeeper(ctrl *gomock.Controller) *MockBtcCheckpointKeeper {
	mock := &MockBtcCheckpointKeeper{ctrl: ctrl}
	mock.recorder = &MockBtcCheckpointKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcCheckpointKeeper) EXPECT() *MockBtcCheckpointKeeperMockRecorder {
	return m.recorder
}

// GetParams mocks base method.
func (m *MockBtcCheckpointKeeper) GetParams(ctx context.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetParams), ctx)
}

// GetPowLimit mocks base method.
func (m *MockBtcCheckpointKeeper) GetPowLimit() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPowLimit")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetPowLimit indicates an expected call of GetPowLimit.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetPowLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPowLimit", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetPowLimit))
}

// MockCheckpointingKeeper is a mock of CheckpointingKeeper interface.
type MockCheckpointingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointingKeeperMockRecorder
}

// MockCheckpointingKeeperMockRecorder is the mock recorder for MockCheckpointingKeeper.
type MockCheckpointingKeeperMockRecorder struct {
	mock *MockCheckpointingKeeper
}

// NewMockCheckpointingKeeper creates a new mock instance.
func NewMockCheckpointingKeeper(ctrl *gomock.Controller) *MockCheckpointingKeeper {
	mock := &MockCheckpointingKeeper{ctrl: ctrl}
	mock.recorder = &MockCheckpointingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointingKeeper) EXPECT() *MockCheckpointingKeeperMockRecorder {
	return m.recorder
}

// GetEpoch mocks base method.
func (m *MockCheckpointingKeeper) GetEpoch(ctx context.Context) *types2.Epoch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpoch", ctx)
	ret0, _ := ret[0].(*types2.Epoch)
	return ret0
}

// GetEpoch indicates an expected call of GetEpoch.
func (mr *MockCheckpointingKeeperMockRecorder) GetEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpoch", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetEpoch), ctx)
}

// GetLastFinalizedEpoch mocks base method.
func (m *MockCheckpointingKeeper) GetLastFinalizedEpoch(ctx context.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastFinalizedEpoch", ctx)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetLastFinalizedEpoch indicates an expected call of GetLastFinalizedEpoch.
func (mr *MockCheckpointingKeeperMockRecorder) GetLastFinalizedEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastFinalizedEpoch", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetLastFinalizedEpoch), ctx)
}

// MockBTCStkConsumerKeeper is a mock of BTCStkConsumerKeeper interface.
type MockBTCStkConsumerKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBTCStkConsumerKeeperMockRecorder
}

// MockBTCStkConsumerKeeperMockRecorder is the mock recorder for MockBTCStkConsumerKeeper.
type MockBTCStkConsumerKeeperMockRecorder struct {
	mock *MockBTCStkConsumerKeeper
}

// NewMockBTCStkConsumerKeeper creates a new mock instance.
func NewMockBTCStkConsumerKeeper(ctrl *gomock.Controller) *MockBTCStkConsumerKeeper {
	mock := &MockBTCStkConsumerKeeper{ctrl: ctrl}
	mock.recorder = &MockBTCStkConsumerKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBTCStkConsumerKeeper) EXPECT() *MockBTCStkConsumerKeeperMockRecorder {
	return m.recorder
}

// GetConsumerFinalityProvider mocks base method.
func (m *MockBTCStkConsumerKeeper) GetConsumerFinalityProvider(ctx context.Context, chainID string, fpBTCPK *types.BIP340PubKey) (*FinalityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumerFinalityProvider", ctx, chainID, fpBTCPK)
	ret0, _ := ret[0].(*FinalityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumerFinalityProvider indicates an expected call of GetConsumerFinalityProvider.
func (mr *MockBTCStkConsumerKeeperMockRecorder) GetConsumerFinalityProvider(ctx, chainID, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumerFinalityProvider", reflect.TypeOf((*MockBTCStkConsumerKeeper)(nil).GetConsumerFinalityProvider), ctx, chainID, fpBTCPK)
}

// GetConsumerOfFinalityProvider mocks base method.
func (m *MockBTCStkConsumerKeeper) GetConsumerOfFinalityProvider(ctx context.Context, fpBTCPK *types.BIP340PubKey) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumerOfFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumerOfFinalityProvider indicates an expected call of GetConsumerOfFinalityProvider.
func (mr *MockBTCStkConsumerKeeperMockRecorder) GetConsumerOfFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumerOfFinalityProvider", reflect.TypeOf((*MockBTCStkConsumerKeeper)(nil).GetConsumerOfFinalityProvider), ctx, fpBTCPK)
}

// HasConsumerFinalityProvider mocks base method.
func (m *MockBTCStkConsumerKeeper) HasConsumerFinalityProvider(ctx context.Context, fpBTCPK *types.BIP340PubKey) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasConsumerFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasConsumerFinalityProvider indicates an expected call of HasConsumerFinalityProvider.
func (mr *MockBTCStkConsumerKeeperMockRecorder) HasConsumerFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasConsumerFinalityProvider", reflect.TypeOf((*MockBTCStkConsumerKeeper)(nil).HasConsumerFinalityProvider), ctx, fpBTCPK)
}

// IsConsumerRegistered mocks base method.
func (m *MockBTCStkConsumerKeeper) IsConsumerRegistered(ctx context.Context, chainID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConsumerRegistered", ctx, chainID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConsumerRegistered indicates an expected call of IsConsumerRegistered.
func (mr *MockBTCStkConsumerKeeperMockRecorder) IsConsumerRegistered(ctx, chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConsumerRegistered", reflect.TypeOf((*MockBTCStkConsumerKeeper)(nil).IsConsumerRegistered), ctx, chainID)
}

// SetConsumerFinalityProvider mocks base method.
func (m *MockBTCStkConsumerKeeper) SetConsumerFinalityProvider(ctx context.Context, fp *FinalityProvider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConsumerFinalityProvider", ctx, fp)
}

// SetConsumerFinalityProvider indicates an expected call of SetConsumerFinalityProvider.
func (mr *MockBTCStkConsumerKeeperMockRecorder) SetConsumerFinalityProvider(ctx, fp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConsumerFinalityProvider", reflect.TypeOf((*MockBTCStkConsumerKeeper)(nil).SetConsumerFinalityProvider), ctx, fp)
}
