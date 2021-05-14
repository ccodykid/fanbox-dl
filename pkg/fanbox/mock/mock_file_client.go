// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/fanbox/file_client.go

// Package mock_fanbox is a generated GoMock package.
package mock_fanbox

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileClient is a mock of FileClient interface.
type MockFileClient struct {
	ctrl     *gomock.Controller
	recorder *MockFileClientMockRecorder
}

// MockFileClientMockRecorder is the mock recorder for MockFileClient.
type MockFileClientMockRecorder struct {
	mock *MockFileClient
}

// NewMockFileClient creates a new mock instance.
func NewMockFileClient(ctrl *gomock.Controller) *MockFileClient {
	mock := &MockFileClient{ctrl: ctrl}
	mock.recorder = &MockFileClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileClient) EXPECT() *MockFileClientMockRecorder {
	return m.recorder
}

// DoesExist mocks base method.
func (m *MockFileClient) DoesExist(name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesExist", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesExist indicates an expected call of DoesExist.
func (mr *MockFileClientMockRecorder) DoesExist(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesExist", reflect.TypeOf((*MockFileClient)(nil).DoesExist), name)
}

// Save mocks base method.
func (m *MockFileClient) Save(name string, reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", name, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockFileClientMockRecorder) Save(name, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFileClient)(nil).Save), name, reader)
}
