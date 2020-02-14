// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/database_users.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockDatabaseUserLister is a mock of DatabaseUserLister interface
type MockDatabaseUserLister struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserListerMockRecorder
}

// MockDatabaseUserListerMockRecorder is the mock recorder for MockDatabaseUserLister
type MockDatabaseUserListerMockRecorder struct {
	mock *MockDatabaseUserLister
}

// NewMockDatabaseUserLister creates a new mock instance
func NewMockDatabaseUserLister(ctrl *gomock.Controller) *MockDatabaseUserLister {
	mock := &MockDatabaseUserLister{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserLister) EXPECT() *MockDatabaseUserListerMockRecorder {
	return m.recorder
}

// ProjectDatabaseUser mocks base method
func (m *MockDatabaseUserLister) ProjectDatabaseUser(groupID string, opts *mongodbatlas.ListOptions) ([]mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectDatabaseUser", groupID, opts)
	ret0, _ := ret[0].([]mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectDatabaseUser indicates an expected call of ProjectDatabaseUser
func (mr *MockDatabaseUserListerMockRecorder) ProjectDatabaseUser(groupID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectDatabaseUser", reflect.TypeOf((*MockDatabaseUserLister)(nil).ProjectDatabaseUser), groupID, opts)
}

// MockDatabaseUserCreator is a mock of DatabaseUserCreator interface
type MockDatabaseUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserCreatorMockRecorder
}

// MockDatabaseUserCreatorMockRecorder is the mock recorder for MockDatabaseUserCreator
type MockDatabaseUserCreatorMockRecorder struct {
	mock *MockDatabaseUserCreator
}

// NewMockDatabaseUserCreator creates a new mock instance
func NewMockDatabaseUserCreator(ctrl *gomock.Controller) *MockDatabaseUserCreator {
	mock := &MockDatabaseUserCreator{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserCreator) EXPECT() *MockDatabaseUserCreatorMockRecorder {
	return m.recorder
}

// CreateDatabaseUser mocks base method
func (m *MockDatabaseUserCreator) CreateDatabaseUser(arg0 *mongodbatlas.DatabaseUser) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser
func (mr *MockDatabaseUserCreatorMockRecorder) CreateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockDatabaseUserCreator)(nil).CreateDatabaseUser), arg0)
}

// MockDatabaseUserDeleter is a mock of DatabaseUserDeleter interface
type MockDatabaseUserDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserDeleterMockRecorder
}

// MockDatabaseUserDeleterMockRecorder is the mock recorder for MockDatabaseUserDeleter
type MockDatabaseUserDeleterMockRecorder struct {
	mock *MockDatabaseUserDeleter
}

// NewMockDatabaseUserDeleter creates a new mock instance
func NewMockDatabaseUserDeleter(ctrl *gomock.Controller) *MockDatabaseUserDeleter {
	mock := &MockDatabaseUserDeleter{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserDeleter) EXPECT() *MockDatabaseUserDeleterMockRecorder {
	return m.recorder
}

// DeleteDatabaseUser mocks base method
func (m *MockDatabaseUserDeleter) DeleteDatabaseUser(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatabaseUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDatabaseUser indicates an expected call of DeleteDatabaseUser
func (mr *MockDatabaseUserDeleterMockRecorder) DeleteDatabaseUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatabaseUser", reflect.TypeOf((*MockDatabaseUserDeleter)(nil).DeleteDatabaseUser), arg0, arg1)
}

// MockDatabaseUserStore is a mock of DatabaseUserStore interface
type MockDatabaseUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserStoreMockRecorder
}

// MockDatabaseUserStoreMockRecorder is the mock recorder for MockDatabaseUserStore
type MockDatabaseUserStoreMockRecorder struct {
	mock *MockDatabaseUserStore
}

// NewMockDatabaseUserStore creates a new mock instance
func NewMockDatabaseUserStore(ctrl *gomock.Controller) *MockDatabaseUserStore {
	mock := &MockDatabaseUserStore{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseUserStore) EXPECT() *MockDatabaseUserStoreMockRecorder {
	return m.recorder
}

// CreateDatabaseUser mocks base method
func (m *MockDatabaseUserStore) CreateDatabaseUser(arg0 *mongodbatlas.DatabaseUser) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser
func (mr *MockDatabaseUserStoreMockRecorder) CreateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockDatabaseUserStore)(nil).CreateDatabaseUser), arg0)
}

// DeleteDatabaseUser mocks base method
func (m *MockDatabaseUserStore) DeleteDatabaseUser(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatabaseUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDatabaseUser indicates an expected call of DeleteDatabaseUser
func (mr *MockDatabaseUserStoreMockRecorder) DeleteDatabaseUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatabaseUser", reflect.TypeOf((*MockDatabaseUserStore)(nil).DeleteDatabaseUser), arg0, arg1)
}

// ProjectDatabaseUser mocks base method
func (m *MockDatabaseUserStore) ProjectDatabaseUser(groupID string, opts *mongodbatlas.ListOptions) ([]mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectDatabaseUser", groupID, opts)
	ret0, _ := ret[0].([]mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectDatabaseUser indicates an expected call of ProjectDatabaseUser
func (mr *MockDatabaseUserStoreMockRecorder) ProjectDatabaseUser(groupID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectDatabaseUser", reflect.TypeOf((*MockDatabaseUserStore)(nil).ProjectDatabaseUser), groupID, opts)
}
