// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: OperatorDBUsersStore)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	atlas "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

// MockOperatorDBUsersStore is a mock of OperatorDBUsersStore interface.
type MockOperatorDBUsersStore struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorDBUsersStoreMockRecorder
}

// MockOperatorDBUsersStoreMockRecorder is the mock recorder for MockOperatorDBUsersStore.
type MockOperatorDBUsersStoreMockRecorder struct {
	mock *MockOperatorDBUsersStore
}

// NewMockOperatorDBUsersStore creates a new mock instance.
func NewMockOperatorDBUsersStore(ctrl *gomock.Controller) *MockOperatorDBUsersStore {
	mock := &MockOperatorDBUsersStore{ctrl: ctrl}
	mock.recorder = &MockOperatorDBUsersStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorDBUsersStore) EXPECT() *MockOperatorDBUsersStoreMockRecorder {
	return m.recorder
}

// DatabaseUsers mocks base method.
func (m *MockOperatorDBUsersStore) DatabaseUsers(arg0 string, arg1 *atlas.ListOptions) (*admin.PaginatedApiAtlasDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUsers indicates an expected call of DatabaseUsers.
func (mr *MockOperatorDBUsersStoreMockRecorder) DatabaseUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUsers", reflect.TypeOf((*MockOperatorDBUsersStore)(nil).DatabaseUsers), arg0, arg1)
}
