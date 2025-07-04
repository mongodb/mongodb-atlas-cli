// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/datalake/aws (interfaces: DataLakePrivateEndpointDescriber)
//
// Generated by this command:
//
//	mockgen -typed -destination=describe_mock_test.go -package=aws . DataLakePrivateEndpointDescriber
//

// Package aws is a generated GoMock package.
package aws

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockDataLakePrivateEndpointDescriber is a mock of DataLakePrivateEndpointDescriber interface.
type MockDataLakePrivateEndpointDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakePrivateEndpointDescriberMockRecorder
	isgomock struct{}
}

// MockDataLakePrivateEndpointDescriberMockRecorder is the mock recorder for MockDataLakePrivateEndpointDescriber.
type MockDataLakePrivateEndpointDescriberMockRecorder struct {
	mock *MockDataLakePrivateEndpointDescriber
}

// NewMockDataLakePrivateEndpointDescriber creates a new mock instance.
func NewMockDataLakePrivateEndpointDescriber(ctrl *gomock.Controller) *MockDataLakePrivateEndpointDescriber {
	mock := &MockDataLakePrivateEndpointDescriber{ctrl: ctrl}
	mock.recorder = &MockDataLakePrivateEndpointDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakePrivateEndpointDescriber) EXPECT() *MockDataLakePrivateEndpointDescriberMockRecorder {
	return m.recorder
}

// DataLakePrivateEndpoint mocks base method.
func (m *MockDataLakePrivateEndpointDescriber) DataLakePrivateEndpoint(arg0, arg1 string) (*admin.PrivateNetworkEndpointIdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLakePrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*admin.PrivateNetworkEndpointIdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLakePrivateEndpoint indicates an expected call of DataLakePrivateEndpoint.
func (mr *MockDataLakePrivateEndpointDescriberMockRecorder) DataLakePrivateEndpoint(arg0, arg1 any) *MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLakePrivateEndpoint", reflect.TypeOf((*MockDataLakePrivateEndpointDescriber)(nil).DataLakePrivateEndpoint), arg0, arg1)
	return &MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall{Call: call}
}

// MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall wrap *gomock.Call
type MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall) Return(arg0 *admin.PrivateNetworkEndpointIdEntry, arg1 error) *MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall) Do(f func(string, string) (*admin.PrivateNetworkEndpointIdEntry, error)) *MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall) DoAndReturn(f func(string, string) (*admin.PrivateNetworkEndpointIdEntry, error)) *MockDataLakePrivateEndpointDescriberDataLakePrivateEndpointCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
