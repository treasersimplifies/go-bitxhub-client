// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	crypto "github.com/meshplus/bitxhub-kit/crypto"
	types "github.com/meshplus/bitxhub-kit/types"
	pb "github.com/meshplus/bitxhub-model/pb"
	go_bitxhub_client "github.com/meshplus/go-bitxhub-client"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Stop mocks base method
func (m *MockClient) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop))
}

// SetPrivateKey mocks base method
func (m *MockClient) SetPrivateKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPrivateKey", arg0)
}

// SetPrivateKey indicates an expected call of SetPrivateKey
func (mr *MockClientMockRecorder) SetPrivateKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrivateKey", reflect.TypeOf((*MockClient)(nil).SetPrivateKey), arg0)
}

// SendTransaction mocks base method
func (m *MockClient) SendTransaction(tx *pb.Transaction) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", tx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockClientMockRecorder) SendTransaction(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockClient)(nil).SendTransaction), tx)
}

// SendTransactionWithReceipt mocks base method
func (m *MockClient) SendTransactionWithReceipt(tx *pb.Transaction) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransactionWithReceipt", tx)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransactionWithReceipt indicates an expected call of SendTransactionWithReceipt
func (mr *MockClientMockRecorder) SendTransactionWithReceipt(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransactionWithReceipt", reflect.TypeOf((*MockClient)(nil).SendTransactionWithReceipt), tx)
}

// GetReceipt mocks base method
func (m *MockClient) GetReceipt(hash string) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipt", hash)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt
func (mr *MockClientMockRecorder) GetReceipt(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockClient)(nil).GetReceipt), hash)
}

// GetTransaction mocks base method
func (m *MockClient) GetTransaction(hash string) (*pb.GetTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", hash)
	ret0, _ := ret[0].(*pb.GetTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockClientMockRecorder) GetTransaction(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockClient)(nil).GetTransaction), hash)
}

// GetChainMeta mocks base method
func (m *MockClient) GetChainMeta() (*pb.ChainMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainMeta")
	ret0, _ := ret[0].(*pb.ChainMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainMeta indicates an expected call of GetChainMeta
func (mr *MockClientMockRecorder) GetChainMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainMeta", reflect.TypeOf((*MockClient)(nil).GetChainMeta))
}

// SyncMerkleWrapper mocks base method
func (m *MockClient) SyncMerkleWrapper(ctx context.Context, id string, num uint64) (chan *pb.MerkleWrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncMerkleWrapper", ctx, id, num)
	ret0, _ := ret[0].(chan *pb.MerkleWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncMerkleWrapper indicates an expected call of SyncMerkleWrapper
func (mr *MockClientMockRecorder) SyncMerkleWrapper(ctx, id, num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncMerkleWrapper", reflect.TypeOf((*MockClient)(nil).SyncMerkleWrapper), ctx, id, num)
}

// GetMerkleWrapper mocks base method
func (m *MockClient) GetMerkleWrapper(ctx context.Context, pid string, begin, end uint64, ch chan<- *pb.MerkleWrapper) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMerkleWrapper", ctx, pid, begin, end, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMerkleWrapper indicates an expected call of GetMerkleWrapper
func (mr *MockClientMockRecorder) GetMerkleWrapper(ctx, pid, begin, end, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMerkleWrapper", reflect.TypeOf((*MockClient)(nil).GetMerkleWrapper), ctx, pid, begin, end, ch)
}

// Subscribe mocks base method
func (m *MockClient) Subscribe(arg0 context.Context, arg1 go_bitxhub_client.SubscriptionType) (<-chan interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(<-chan interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockClientMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockClient)(nil).Subscribe), arg0, arg1)
}

// DeployContract mocks base method
func (m *MockClient) DeployContract(contract []byte) (types.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployContract", contract)
	ret0, _ := ret[0].(types.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployContract indicates an expected call of DeployContract
func (mr *MockClientMockRecorder) DeployContract(contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployContract", reflect.TypeOf((*MockClient)(nil).DeployContract), contract)
}

// InvokeContract mocks base method
func (m *MockClient) InvokeContract(vmType pb.TransactionData_VMType, address types.Address, method string, args ...*pb.Arg) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{vmType, address, method}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeContract", varargs...)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeContract indicates an expected call of InvokeContract
func (mr *MockClientMockRecorder) InvokeContract(vmType, address, method interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{vmType, address, method}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeContract", reflect.TypeOf((*MockClient)(nil).InvokeContract), varargs...)
}

// InvokeBVMContract mocks base method
func (m *MockClient) InvokeBVMContract(address types.Address, method string, args ...*pb.Arg) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{address, method}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeBVMContract", varargs...)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeBVMContract indicates an expected call of InvokeBVMContract
func (mr *MockClientMockRecorder) InvokeBVMContract(address, method interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{address, method}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeBVMContract", reflect.TypeOf((*MockClient)(nil).InvokeBVMContract), varargs...)
}

// InvokeXVMContract mocks base method
func (m *MockClient) InvokeXVMContract(address types.Address, method string, args ...*pb.Arg) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{address, method}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeXVMContract", varargs...)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeXVMContract indicates an expected call of InvokeXVMContract
func (mr *MockClientMockRecorder) InvokeXVMContract(address, method interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{address, method}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeXVMContract", reflect.TypeOf((*MockClient)(nil).InvokeXVMContract), varargs...)
}