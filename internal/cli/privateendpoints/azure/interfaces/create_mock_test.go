// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/azure/interfaces (interfaces: InterfaceEndpointCreator)
//
// Generated by this command:
//
//	mockgen -typed -destination=create_mock_test.go -package=interfaces . InterfaceEndpointCreator
//

// Package interfaces is a generated GoMock package.
package interfaces

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockInterfaceEndpointCreator is a mock of InterfaceEndpointCreator interface.
type MockInterfaceEndpointCreator struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointCreatorMockRecorder
	isgomock struct{}
}

// MockInterfaceEndpointCreatorMockRecorder is the mock recorder for MockInterfaceEndpointCreator.
type MockInterfaceEndpointCreatorMockRecorder struct {
	mock *MockInterfaceEndpointCreator
}

// NewMockInterfaceEndpointCreator creates a new mock instance.
func NewMockInterfaceEndpointCreator(ctrl *gomock.Controller) *MockInterfaceEndpointCreator {
	mock := &MockInterfaceEndpointCreator{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointCreator) EXPECT() *MockInterfaceEndpointCreatorMockRecorder {
	return m.recorder
}

// CreateInterfaceEndpoint mocks base method.
func (m *MockInterfaceEndpointCreator) CreateInterfaceEndpoint(arg0, arg1, arg2 string, arg3 *admin.CreateEndpointRequest) (*admin.PrivateLinkEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInterfaceEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.PrivateLinkEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInterfaceEndpoint indicates an expected call of CreateInterfaceEndpoint.
func (mr *MockInterfaceEndpointCreatorMockRecorder) CreateInterfaceEndpoint(arg0, arg1, arg2, arg3 any) *MockInterfaceEndpointCreatorCreateInterfaceEndpointCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInterfaceEndpoint", reflect.TypeOf((*MockInterfaceEndpointCreator)(nil).CreateInterfaceEndpoint), arg0, arg1, arg2, arg3)
	return &MockInterfaceEndpointCreatorCreateInterfaceEndpointCall{Call: call}
}

// MockInterfaceEndpointCreatorCreateInterfaceEndpointCall wrap *gomock.Call
type MockInterfaceEndpointCreatorCreateInterfaceEndpointCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceEndpointCreatorCreateInterfaceEndpointCall) Return(arg0 *admin.PrivateLinkEndpoint, arg1 error) *MockInterfaceEndpointCreatorCreateInterfaceEndpointCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceEndpointCreatorCreateInterfaceEndpointCall) Do(f func(string, string, string, *admin.CreateEndpointRequest) (*admin.PrivateLinkEndpoint, error)) *MockInterfaceEndpointCreatorCreateInterfaceEndpointCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceEndpointCreatorCreateInterfaceEndpointCall) DoAndReturn(f func(string, string, string, *admin.CreateEndpointRequest) (*admin.PrivateLinkEndpoint, error)) *MockInterfaceEndpointCreatorCreateInterfaceEndpointCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
