// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/aws/interfaces (interfaces: InterfaceEndpointDeleter)
//
// Generated by this command:
//
//	mockgen -typed -destination=delete_mock_test.go -package=interfaces . InterfaceEndpointDeleter
//

// Package interfaces is a generated GoMock package.
package interfaces

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInterfaceEndpointDeleter is a mock of InterfaceEndpointDeleter interface.
type MockInterfaceEndpointDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointDeleterMockRecorder
	isgomock struct{}
}

// MockInterfaceEndpointDeleterMockRecorder is the mock recorder for MockInterfaceEndpointDeleter.
type MockInterfaceEndpointDeleterMockRecorder struct {
	mock *MockInterfaceEndpointDeleter
}

// NewMockInterfaceEndpointDeleter creates a new mock instance.
func NewMockInterfaceEndpointDeleter(ctrl *gomock.Controller) *MockInterfaceEndpointDeleter {
	mock := &MockInterfaceEndpointDeleter{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointDeleter) EXPECT() *MockInterfaceEndpointDeleterMockRecorder {
	return m.recorder
}

// DeleteInterfaceEndpoint mocks base method.
func (m *MockInterfaceEndpointDeleter) DeleteInterfaceEndpoint(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInterfaceEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInterfaceEndpoint indicates an expected call of DeleteInterfaceEndpoint.
func (mr *MockInterfaceEndpointDeleterMockRecorder) DeleteInterfaceEndpoint(arg0, arg1, arg2, arg3 any) *MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInterfaceEndpoint", reflect.TypeOf((*MockInterfaceEndpointDeleter)(nil).DeleteInterfaceEndpoint), arg0, arg1, arg2, arg3)
	return &MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall{Call: call}
}

// MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall wrap *gomock.Call
type MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall) Return(arg0 error) *MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall) Do(f func(string, string, string, string) error) *MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall) DoAndReturn(f func(string, string, string, string) error) *MockInterfaceEndpointDeleterDeleteInterfaceEndpointCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
