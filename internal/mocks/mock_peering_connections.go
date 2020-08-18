// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: PeeringConnectionLister,PeeringConnectionCreator,PeeringConnectionDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockPeeringConnectionLister is a mock of PeeringConnectionLister interface
type MockPeeringConnectionLister struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionListerMockRecorder
}

// MockPeeringConnectionListerMockRecorder is the mock recorder for MockPeeringConnectionLister
type MockPeeringConnectionListerMockRecorder struct {
	mock *MockPeeringConnectionLister
}

// NewMockPeeringConnectionLister creates a new mock instance
func NewMockPeeringConnectionLister(ctrl *gomock.Controller) *MockPeeringConnectionLister {
	mock := &MockPeeringConnectionLister{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeeringConnectionLister) EXPECT() *MockPeeringConnectionListerMockRecorder {
	return m.recorder
}

// PeeringConnections mocks base method
func (m *MockPeeringConnectionLister) PeeringConnections(arg0 string, arg1 *mongodbatlas.ContainersListOptions) ([]mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnections", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnections indicates an expected call of PeeringConnections
func (mr *MockPeeringConnectionListerMockRecorder) PeeringConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnections", reflect.TypeOf((*MockPeeringConnectionLister)(nil).PeeringConnections), arg0, arg1)
}

// MockPeeringConnectionCreator is a mock of PeeringConnectionCreator interface
type MockPeeringConnectionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionCreatorMockRecorder
}

// MockPeeringConnectionCreatorMockRecorder is the mock recorder for MockPeeringConnectionCreator
type MockPeeringConnectionCreatorMockRecorder struct {
	mock *MockPeeringConnectionCreator
}

// NewMockPeeringConnectionCreator creates a new mock instance
func NewMockPeeringConnectionCreator(ctrl *gomock.Controller) *MockPeeringConnectionCreator {
	mock := &MockPeeringConnectionCreator{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeeringConnectionCreator) EXPECT() *MockPeeringConnectionCreatorMockRecorder {
	return m.recorder
}

// AWSContainers mocks base method
func (m *MockPeeringConnectionCreator) AWSContainers(arg0 string) ([]mongodbatlas.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AWSContainers", arg0)
	ret0, _ := ret[0].([]mongodbatlas.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AWSContainers indicates an expected call of AWSContainers
func (mr *MockPeeringConnectionCreatorMockRecorder) AWSContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AWSContainers", reflect.TypeOf((*MockPeeringConnectionCreator)(nil).AWSContainers), arg0)
}

// AzureContainers mocks base method
func (m *MockPeeringConnectionCreator) AzureContainers(arg0 string) ([]mongodbatlas.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AzureContainers", arg0)
	ret0, _ := ret[0].([]mongodbatlas.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AzureContainers indicates an expected call of AzureContainers
func (mr *MockPeeringConnectionCreatorMockRecorder) AzureContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AzureContainers", reflect.TypeOf((*MockPeeringConnectionCreator)(nil).AzureContainers), arg0)
}

// CreateContainer mocks base method
func (m *MockPeeringConnectionCreator) CreateContainer(arg0 string, arg1 *mongodbatlas.Container) (*mongodbatlas.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockPeeringConnectionCreatorMockRecorder) CreateContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockPeeringConnectionCreator)(nil).CreateContainer), arg0, arg1)
}

// CreatePeeringConnection mocks base method
func (m *MockPeeringConnectionCreator) CreatePeeringConnection(arg0 string, arg1 *mongodbatlas.Peer) (*mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeeringConnection indicates an expected call of CreatePeeringConnection
func (mr *MockPeeringConnectionCreatorMockRecorder) CreatePeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeeringConnection", reflect.TypeOf((*MockPeeringConnectionCreator)(nil).CreatePeeringConnection), arg0, arg1)
}

// MockPeeringConnectionDeleter is a mock of PeeringConnectionDeleter interface
type MockPeeringConnectionDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionDeleterMockRecorder
}

// MockPeeringConnectionDeleterMockRecorder is the mock recorder for MockPeeringConnectionDeleter
type MockPeeringConnectionDeleterMockRecorder struct {
	mock *MockPeeringConnectionDeleter
}

// NewMockPeeringConnectionDeleter creates a new mock instance
func NewMockPeeringConnectionDeleter(ctrl *gomock.Controller) *MockPeeringConnectionDeleter {
	mock := &MockPeeringConnectionDeleter{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeeringConnectionDeleter) EXPECT() *MockPeeringConnectionDeleterMockRecorder {
	return m.recorder
}

// DeletePeeringConnection mocks base method
func (m *MockPeeringConnectionDeleter) DeletePeeringConnection(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeeringConnection indicates an expected call of DeletePeeringConnection
func (mr *MockPeeringConnectionDeleterMockRecorder) DeletePeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeeringConnection", reflect.TypeOf((*MockPeeringConnectionDeleter)(nil).DeletePeeringConnection), arg0, arg1)
}
