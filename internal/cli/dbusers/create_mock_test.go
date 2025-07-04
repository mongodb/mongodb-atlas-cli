// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/dbusers (interfaces: DatabaseUserCreator)
//
// Generated by this command:
//
//	mockgen -typed -destination=create_mock_test.go -package=dbusers . DatabaseUserCreator
//

// Package dbusers is a generated GoMock package.
package dbusers

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabaseUserCreator is a mock of DatabaseUserCreator interface.
type MockDatabaseUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserCreatorMockRecorder
	isgomock struct{}
}

// MockDatabaseUserCreatorMockRecorder is the mock recorder for MockDatabaseUserCreator.
type MockDatabaseUserCreatorMockRecorder struct {
	mock *MockDatabaseUserCreator
}

// NewMockDatabaseUserCreator creates a new mock instance.
func NewMockDatabaseUserCreator(ctrl *gomock.Controller) *MockDatabaseUserCreator {
	mock := &MockDatabaseUserCreator{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseUserCreator) EXPECT() *MockDatabaseUserCreatorMockRecorder {
	return m.recorder
}

// CreateDatabaseUser mocks base method.
func (m *MockDatabaseUserCreator) CreateDatabaseUser(arg0 *admin.CloudDatabaseUser) (*admin.CloudDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0)
	ret0, _ := ret[0].(*admin.CloudDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser.
func (mr *MockDatabaseUserCreatorMockRecorder) CreateDatabaseUser(arg0 any) *MockDatabaseUserCreatorCreateDatabaseUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockDatabaseUserCreator)(nil).CreateDatabaseUser), arg0)
	return &MockDatabaseUserCreatorCreateDatabaseUserCall{Call: call}
}

// MockDatabaseUserCreatorCreateDatabaseUserCall wrap *gomock.Call
type MockDatabaseUserCreatorCreateDatabaseUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDatabaseUserCreatorCreateDatabaseUserCall) Return(arg0 *admin.CloudDatabaseUser, arg1 error) *MockDatabaseUserCreatorCreateDatabaseUserCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDatabaseUserCreatorCreateDatabaseUserCall) Do(f func(*admin.CloudDatabaseUser) (*admin.CloudDatabaseUser, error)) *MockDatabaseUserCreatorCreateDatabaseUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDatabaseUserCreatorCreateDatabaseUserCall) DoAndReturn(f func(*admin.CloudDatabaseUser) (*admin.CloudDatabaseUser, error)) *MockDatabaseUserCreatorCreateDatabaseUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
