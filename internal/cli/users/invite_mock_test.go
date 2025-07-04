// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/users (interfaces: UserCreator)
//
// Generated by this command:
//
//	mockgen -typed -destination=invite_mock_test.go -package=users . UserCreator
//

// Package users is a generated GoMock package.
package users

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockUserCreator is a mock of UserCreator interface.
type MockUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockUserCreatorMockRecorder
	isgomock struct{}
}

// MockUserCreatorMockRecorder is the mock recorder for MockUserCreator.
type MockUserCreatorMockRecorder struct {
	mock *MockUserCreator
}

// NewMockUserCreator creates a new mock instance.
func NewMockUserCreator(ctrl *gomock.Controller) *MockUserCreator {
	mock := &MockUserCreator{ctrl: ctrl}
	mock.recorder = &MockUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCreator) EXPECT() *MockUserCreatorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserCreator) CreateUser(user *admin.CloudAppUser) (*admin.CloudAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*admin.CloudAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserCreatorMockRecorder) CreateUser(user any) *MockUserCreatorCreateUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserCreator)(nil).CreateUser), user)
	return &MockUserCreatorCreateUserCall{Call: call}
}

// MockUserCreatorCreateUserCall wrap *gomock.Call
type MockUserCreatorCreateUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserCreatorCreateUserCall) Return(arg0 *admin.CloudAppUser, arg1 error) *MockUserCreatorCreateUserCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserCreatorCreateUserCall) Do(f func(*admin.CloudAppUser) (*admin.CloudAppUser, error)) *MockUserCreatorCreateUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserCreatorCreateUserCall) DoAndReturn(f func(*admin.CloudAppUser) (*admin.CloudAppUser, error)) *MockUserCreatorCreateUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
