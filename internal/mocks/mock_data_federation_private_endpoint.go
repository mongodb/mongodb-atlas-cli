// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: DataFederationPrivateEndpointLister,DataFederationPrivateEndpointDescriber,DataFederationPrivateEndpointCreator,DataFederationPrivateEndpointDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20240530003/admin"
)

// MockDataFederationPrivateEndpointLister is a mock of DataFederationPrivateEndpointLister interface.
type MockDataFederationPrivateEndpointLister struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationPrivateEndpointListerMockRecorder
}

// MockDataFederationPrivateEndpointListerMockRecorder is the mock recorder for MockDataFederationPrivateEndpointLister.
type MockDataFederationPrivateEndpointListerMockRecorder struct {
	mock *MockDataFederationPrivateEndpointLister
}

// NewMockDataFederationPrivateEndpointLister creates a new mock instance.
func NewMockDataFederationPrivateEndpointLister(ctrl *gomock.Controller) *MockDataFederationPrivateEndpointLister {
	mock := &MockDataFederationPrivateEndpointLister{ctrl: ctrl}
	mock.recorder = &MockDataFederationPrivateEndpointListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationPrivateEndpointLister) EXPECT() *MockDataFederationPrivateEndpointListerMockRecorder {
	return m.recorder
}

// DataFederationPrivateEndpoints mocks base method.
func (m *MockDataFederationPrivateEndpointLister) DataFederationPrivateEndpoints(arg0 string) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationPrivateEndpoints", arg0)
	ret0, _ := ret[0].(*admin.PaginatedPrivateNetworkEndpointIdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationPrivateEndpoints indicates an expected call of DataFederationPrivateEndpoints.
func (mr *MockDataFederationPrivateEndpointListerMockRecorder) DataFederationPrivateEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationPrivateEndpoints", reflect.TypeOf((*MockDataFederationPrivateEndpointLister)(nil).DataFederationPrivateEndpoints), arg0)
}

// MockDataFederationPrivateEndpointDescriber is a mock of DataFederationPrivateEndpointDescriber interface.
type MockDataFederationPrivateEndpointDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationPrivateEndpointDescriberMockRecorder
}

// MockDataFederationPrivateEndpointDescriberMockRecorder is the mock recorder for MockDataFederationPrivateEndpointDescriber.
type MockDataFederationPrivateEndpointDescriberMockRecorder struct {
	mock *MockDataFederationPrivateEndpointDescriber
}

// NewMockDataFederationPrivateEndpointDescriber creates a new mock instance.
func NewMockDataFederationPrivateEndpointDescriber(ctrl *gomock.Controller) *MockDataFederationPrivateEndpointDescriber {
	mock := &MockDataFederationPrivateEndpointDescriber{ctrl: ctrl}
	mock.recorder = &MockDataFederationPrivateEndpointDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationPrivateEndpointDescriber) EXPECT() *MockDataFederationPrivateEndpointDescriberMockRecorder {
	return m.recorder
}

// DataFederationPrivateEndpoint mocks base method.
func (m *MockDataFederationPrivateEndpointDescriber) DataFederationPrivateEndpoint(arg0, arg1 string) (*admin.PrivateNetworkEndpointIdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationPrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*admin.PrivateNetworkEndpointIdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationPrivateEndpoint indicates an expected call of DataFederationPrivateEndpoint.
func (mr *MockDataFederationPrivateEndpointDescriberMockRecorder) DataFederationPrivateEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationPrivateEndpoint", reflect.TypeOf((*MockDataFederationPrivateEndpointDescriber)(nil).DataFederationPrivateEndpoint), arg0, arg1)
}

// MockDataFederationPrivateEndpointCreator is a mock of DataFederationPrivateEndpointCreator interface.
type MockDataFederationPrivateEndpointCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationPrivateEndpointCreatorMockRecorder
}

// MockDataFederationPrivateEndpointCreatorMockRecorder is the mock recorder for MockDataFederationPrivateEndpointCreator.
type MockDataFederationPrivateEndpointCreatorMockRecorder struct {
	mock *MockDataFederationPrivateEndpointCreator
}

// NewMockDataFederationPrivateEndpointCreator creates a new mock instance.
func NewMockDataFederationPrivateEndpointCreator(ctrl *gomock.Controller) *MockDataFederationPrivateEndpointCreator {
	mock := &MockDataFederationPrivateEndpointCreator{ctrl: ctrl}
	mock.recorder = &MockDataFederationPrivateEndpointCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationPrivateEndpointCreator) EXPECT() *MockDataFederationPrivateEndpointCreatorMockRecorder {
	return m.recorder
}

// CreateDataFederationPrivateEndpoint mocks base method.
func (m *MockDataFederationPrivateEndpointCreator) CreateDataFederationPrivateEndpoint(arg0 string, arg1 *admin.PrivateNetworkEndpointIdEntry) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataFederationPrivateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedPrivateNetworkEndpointIdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataFederationPrivateEndpoint indicates an expected call of CreateDataFederationPrivateEndpoint.
func (mr *MockDataFederationPrivateEndpointCreatorMockRecorder) CreateDataFederationPrivateEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataFederationPrivateEndpoint", reflect.TypeOf((*MockDataFederationPrivateEndpointCreator)(nil).CreateDataFederationPrivateEndpoint), arg0, arg1)
}

// MockDataFederationPrivateEndpointDeleter is a mock of DataFederationPrivateEndpointDeleter interface.
type MockDataFederationPrivateEndpointDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationPrivateEndpointDeleterMockRecorder
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
func (mr *MockDataFederationPrivateEndpointDeleterMockRecorder) DeleteDataFederationPrivateEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataFederationPrivateEndpoint", reflect.TypeOf((*MockDataFederationPrivateEndpointDeleter)(nil).DeleteDataFederationPrivateEndpoint), arg0, arg1)
}
