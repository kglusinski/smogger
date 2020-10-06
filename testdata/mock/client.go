// Code generated by MockGen. DO NOT EDIT.
// Source: smogger/internal/smogger (interfaces: ApiClient)

// Package mock_smogger is a generated GoMock package.
package mock_smogger

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApiClient is a mock of ApiClient interface
type MockApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientMockRecorder
}

// MockApiClientMockRecorder is the mock recorder for MockApiClient
type MockApiClientMockRecorder struct {
	mock *MockApiClient
}

// NewMockApiClient creates a new mock instance
func NewMockApiClient(ctrl *gomock.Controller) *MockApiClient {
	mock := &MockApiClient{ctrl: ctrl}
	mock.recorder = &MockApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiClient) EXPECT() *MockApiClientMockRecorder {
	return m.recorder
}

// Cities mocks base method
func (m *MockApiClient) Cities(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cities", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cities indicates an expected call of Cities
func (mr *MockApiClientMockRecorder) Cities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cities", reflect.TypeOf((*MockApiClient)(nil).Cities), arg0, arg1)
}
