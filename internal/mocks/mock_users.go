// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: UserCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	store "github.com/mongodb/mongocli/internal/store"
	reflect "reflect"
)

// MockUserCreator is a mock of UserCreator interface
type MockUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockUserCreatorMockRecorder
}

// MockUserCreatorMockRecorder is the mock recorder for MockUserCreator
type MockUserCreatorMockRecorder struct {
	mock *MockUserCreator
}

// NewMockUserCreator creates a new mock instance
func NewMockUserCreator(ctrl *gomock.Controller) *MockUserCreator {
	mock := &MockUserCreator{ctrl: ctrl}
	mock.recorder = &MockUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserCreator) EXPECT() *MockUserCreatorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserCreator) CreateUser(arg0 *store.UserRequest) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserCreatorMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserCreator)(nil).CreateUser), arg0)
}
