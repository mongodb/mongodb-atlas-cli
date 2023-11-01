// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: AtlasOperatorDBUsersStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231001002/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockAtlasOperatorDBUsersStore is a mock of AtlasOperatorDBUsersStore interface.
type MockAtlasOperatorDBUsersStore struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasOperatorDBUsersStoreMockRecorder
}

// MockAtlasOperatorDBUsersStoreMockRecorder is the mock recorder for MockAtlasOperatorDBUsersStore.
type MockAtlasOperatorDBUsersStoreMockRecorder struct {
	mock *MockAtlasOperatorDBUsersStore
}

// NewMockAtlasOperatorDBUsersStore creates a new mock instance.
func NewMockAtlasOperatorDBUsersStore(ctrl *gomock.Controller) *MockAtlasOperatorDBUsersStore {
	mock := &MockAtlasOperatorDBUsersStore{ctrl: ctrl}
	mock.recorder = &MockAtlasOperatorDBUsersStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtlasOperatorDBUsersStore) EXPECT() *MockAtlasOperatorDBUsersStoreMockRecorder {
	return m.recorder
}

// DatabaseUsers mocks base method.
func (m *MockAtlasOperatorDBUsersStore) DatabaseUsers(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedApiAtlasDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUsers indicates an expected call of DatabaseUsers.
func (mr *MockAtlasOperatorDBUsersStoreMockRecorder) DatabaseUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUsers", reflect.TypeOf((*MockAtlasOperatorDBUsersStore)(nil).DatabaseUsers), arg0, arg1)
}
