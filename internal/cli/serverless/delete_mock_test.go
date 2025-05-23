// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/serverless (interfaces: Deleter)
//
// Generated by this command:
//
//	mockgen -typed -destination=delete_mock_test.go -package=serverless . Deleter
//

// Package serverless is a generated GoMock package.
package serverless

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDeleter is a mock of Deleter interface.
type MockDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDeleterMockRecorder
	isgomock struct{}
}

// MockDeleterMockRecorder is the mock recorder for MockDeleter.
type MockDeleterMockRecorder struct {
	mock *MockDeleter
}

// NewMockDeleter creates a new mock instance.
func NewMockDeleter(ctrl *gomock.Controller) *MockDeleter {
	mock := &MockDeleter{ctrl: ctrl}
	mock.recorder = &MockDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleter) EXPECT() *MockDeleterMockRecorder {
	return m.recorder
}

// DeleteServerlessInstance mocks base method.
func (m *MockDeleter) DeleteServerlessInstance(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServerlessInstance indicates an expected call of DeleteServerlessInstance.
func (mr *MockDeleterMockRecorder) DeleteServerlessInstance(arg0, arg1 any) *MockDeleterDeleteServerlessInstanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServerlessInstance", reflect.TypeOf((*MockDeleter)(nil).DeleteServerlessInstance), arg0, arg1)
	return &MockDeleterDeleteServerlessInstanceCall{Call: call}
}

// MockDeleterDeleteServerlessInstanceCall wrap *gomock.Call
type MockDeleterDeleteServerlessInstanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDeleterDeleteServerlessInstanceCall) Return(arg0 error) *MockDeleterDeleteServerlessInstanceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDeleterDeleteServerlessInstanceCall) Do(f func(string, string) error) *MockDeleterDeleteServerlessInstanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDeleterDeleteServerlessInstanceCall) DoAndReturn(f func(string, string) error) *MockDeleterDeleteServerlessInstanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
