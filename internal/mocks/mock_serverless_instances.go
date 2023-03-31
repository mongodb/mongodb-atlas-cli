// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: ServerlessInstanceLister,ServerlessInstanceDescriber,ServerlessInstanceDeleter,ServerlessInstanceCreator,ServerlessInstanceUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockServerlessInstanceLister is a mock of ServerlessInstanceLister interface.
type MockServerlessInstanceLister struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessInstanceListerMockRecorder
}

// MockServerlessInstanceListerMockRecorder is the mock recorder for MockServerlessInstanceLister.
type MockServerlessInstanceListerMockRecorder struct {
	mock *MockServerlessInstanceLister
}

// NewMockServerlessInstanceLister creates a new mock instance.
func NewMockServerlessInstanceLister(ctrl *gomock.Controller) *MockServerlessInstanceLister {
	mock := &MockServerlessInstanceLister{ctrl: ctrl}
	mock.recorder = &MockServerlessInstanceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessInstanceLister) EXPECT() *MockServerlessInstanceListerMockRecorder {
	return m.recorder
}

// ServerlessInstances mocks base method.
func (m *MockServerlessInstanceLister) ServerlessInstances(arg0 string, arg1 *mongodbatlas.ListOptions) (*mongodbatlas.ClustersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstances", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.ClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstances indicates an expected call of ServerlessInstances.
func (mr *MockServerlessInstanceListerMockRecorder) ServerlessInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstances", reflect.TypeOf((*MockServerlessInstanceLister)(nil).ServerlessInstances), arg0, arg1)
}

// MockServerlessInstanceDescriber is a mock of ServerlessInstanceDescriber interface.
type MockServerlessInstanceDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessInstanceDescriberMockRecorder
}

// MockServerlessInstanceDescriberMockRecorder is the mock recorder for MockServerlessInstanceDescriber.
type MockServerlessInstanceDescriberMockRecorder struct {
	mock *MockServerlessInstanceDescriber
}

// NewMockServerlessInstanceDescriber creates a new mock instance.
func NewMockServerlessInstanceDescriber(ctrl *gomock.Controller) *MockServerlessInstanceDescriber {
	mock := &MockServerlessInstanceDescriber{ctrl: ctrl}
	mock.recorder = &MockServerlessInstanceDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessInstanceDescriber) EXPECT() *MockServerlessInstanceDescriberMockRecorder {
	return m.recorder
}

// ServerlessInstance mocks base method.
func (m *MockServerlessInstanceDescriber) ServerlessInstance(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstance indicates an expected call of ServerlessInstance.
func (mr *MockServerlessInstanceDescriberMockRecorder) ServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstance", reflect.TypeOf((*MockServerlessInstanceDescriber)(nil).ServerlessInstance), arg0, arg1)
}

// MockServerlessInstanceDeleter is a mock of ServerlessInstanceDeleter interface.
type MockServerlessInstanceDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessInstanceDeleterMockRecorder
}

// MockServerlessInstanceDeleterMockRecorder is the mock recorder for MockServerlessInstanceDeleter.
type MockServerlessInstanceDeleterMockRecorder struct {
	mock *MockServerlessInstanceDeleter
}

// NewMockServerlessInstanceDeleter creates a new mock instance.
func NewMockServerlessInstanceDeleter(ctrl *gomock.Controller) *MockServerlessInstanceDeleter {
	mock := &MockServerlessInstanceDeleter{ctrl: ctrl}
	mock.recorder = &MockServerlessInstanceDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessInstanceDeleter) EXPECT() *MockServerlessInstanceDeleterMockRecorder {
	return m.recorder
}

// DeleteServerlessInstance mocks base method.
func (m *MockServerlessInstanceDeleter) DeleteServerlessInstance(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServerlessInstance indicates an expected call of DeleteServerlessInstance.
func (mr *MockServerlessInstanceDeleterMockRecorder) DeleteServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServerlessInstance", reflect.TypeOf((*MockServerlessInstanceDeleter)(nil).DeleteServerlessInstance), arg0, arg1)
}

// MockServerlessInstanceCreator is a mock of ServerlessInstanceCreator interface.
type MockServerlessInstanceCreator struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessInstanceCreatorMockRecorder
}

// MockServerlessInstanceCreatorMockRecorder is the mock recorder for MockServerlessInstanceCreator.
type MockServerlessInstanceCreatorMockRecorder struct {
	mock *MockServerlessInstanceCreator
}

// NewMockServerlessInstanceCreator creates a new mock instance.
func NewMockServerlessInstanceCreator(ctrl *gomock.Controller) *MockServerlessInstanceCreator {
	mock := &MockServerlessInstanceCreator{ctrl: ctrl}
	mock.recorder = &MockServerlessInstanceCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessInstanceCreator) EXPECT() *MockServerlessInstanceCreatorMockRecorder {
	return m.recorder
}

// CreateServerlessInstance mocks base method.
func (m *MockServerlessInstanceCreator) CreateServerlessInstance(arg0 string, arg1 *mongodbatlas.ServerlessCreateRequestParams) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServerlessInstance indicates an expected call of CreateServerlessInstance.
func (mr *MockServerlessInstanceCreatorMockRecorder) CreateServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServerlessInstance", reflect.TypeOf((*MockServerlessInstanceCreator)(nil).CreateServerlessInstance), arg0, arg1)
}

// MockServerlessInstanceUpdater is a mock of ServerlessInstanceUpdater interface.
type MockServerlessInstanceUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessInstanceUpdaterMockRecorder
}

// MockServerlessInstanceUpdaterMockRecorder is the mock recorder for MockServerlessInstanceUpdater.
type MockServerlessInstanceUpdaterMockRecorder struct {
	mock *MockServerlessInstanceUpdater
}

// NewMockServerlessInstanceUpdater creates a new mock instance.
func NewMockServerlessInstanceUpdater(ctrl *gomock.Controller) *MockServerlessInstanceUpdater {
	mock := &MockServerlessInstanceUpdater{ctrl: ctrl}
	mock.recorder = &MockServerlessInstanceUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessInstanceUpdater) EXPECT() *MockServerlessInstanceUpdaterMockRecorder {
	return m.recorder
}

// UpdateServerlessInstance mocks base method.
func (m *MockServerlessInstanceUpdater) UpdateServerlessInstance(arg0, arg1 string, arg2 *mongodbatlas.ServerlessUpdateRequestParams) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServerlessInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServerlessInstance indicates an expected call of UpdateServerlessInstance.
func (mr *MockServerlessInstanceUpdaterMockRecorder) UpdateServerlessInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServerlessInstance", reflect.TypeOf((*MockServerlessInstanceUpdater)(nil).UpdateServerlessInstance), arg0, arg1, arg2)
}
