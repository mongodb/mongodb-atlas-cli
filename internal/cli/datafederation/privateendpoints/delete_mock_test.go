// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/datafederation/privateendpoints (interfaces: DataFederationPrivateEndpointDeleter)
//
// Generated by this command:
//
//	mockgen -typed -destination=delete_mock_test.go -package=privateendpoints . DataFederationPrivateEndpointDeleter
//

// Package privateendpoints is a generated GoMock package.
package privateendpoints

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDataFederationPrivateEndpointDeleter is a mock of DataFederationPrivateEndpointDeleter interface.
type MockDataFederationPrivateEndpointDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationPrivateEndpointDeleterMockRecorder
	isgomock struct{}
}

// MockDataFederationPrivateEndpointDeleterMockRecorder is the mock recorder for MockDataFederationPrivateEndpointDeleter.
type MockDataFederationPrivateEndpointDeleterMockRecorder struct {
	mock *MockDataFederationPrivateEndpointDeleter
}

// NewMockDataFederationPrivateEndpointDeleter creates a new mock instance.
func NewMockDataFederationPrivateEndpointDeleter(ctrl *gomock.Controller) *MockDataFederationPrivateEndpointDeleter {
	mock := &MockDataFederationPrivateEndpointDeleter{ctrl: ctrl}
	mock.recorder = &MockDataFederationPrivateEndpointDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationPrivateEndpointDeleter) EXPECT() *MockDataFederationPrivateEndpointDeleterMockRecorder {
	return m.recorder
}

// DeleteDataFederationPrivateEndpoint mocks base method.
func (m *MockDataFederationPrivateEndpointDeleter) DeleteDataFederationPrivateEndpoint(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataFederationPrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataFederationPrivateEndpoint indicates an expected call of DeleteDataFederationPrivateEndpoint.
func (mr *MockDataFederationPrivateEndpointDeleterMockRecorder) DeleteDataFederationPrivateEndpoint(arg0, arg1 any) *MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataFederationPrivateEndpoint", reflect.TypeOf((*MockDataFederationPrivateEndpointDeleter)(nil).DeleteDataFederationPrivateEndpoint), arg0, arg1)
	return &MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall{Call: call}
}

// MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall wrap *gomock.Call
type MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall) Return(arg0 error) *MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall) Do(f func(string, string) error) *MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall) DoAndReturn(f func(string, string) error) *MockDataFederationPrivateEndpointDeleterDeleteDataFederationPrivateEndpointCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
