// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/customdbroles (interfaces: DatabaseRoleUpdater)
//
// Generated by this command:
//
//	mockgen -typed -destination=update_mock_test.go -package=customdbroles . DatabaseRoleUpdater
//

// Package customdbroles is a generated GoMock package.
package customdbroles

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabaseRoleUpdater is a mock of DatabaseRoleUpdater interface.
type MockDatabaseRoleUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRoleUpdaterMockRecorder
	isgomock struct{}
}

// MockDatabaseRoleUpdaterMockRecorder is the mock recorder for MockDatabaseRoleUpdater.
type MockDatabaseRoleUpdaterMockRecorder struct {
	mock *MockDatabaseRoleUpdater
}

// NewMockDatabaseRoleUpdater creates a new mock instance.
func NewMockDatabaseRoleUpdater(ctrl *gomock.Controller) *MockDatabaseRoleUpdater {
	mock := &MockDatabaseRoleUpdater{ctrl: ctrl}
	mock.recorder = &MockDatabaseRoleUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRoleUpdater) EXPECT() *MockDatabaseRoleUpdaterMockRecorder {
	return m.recorder
}

// DatabaseRole mocks base method.
func (m *MockDatabaseRoleUpdater) DatabaseRole(arg0, arg1 string) (*admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRole", arg0, arg1)
	ret0, _ := ret[0].(*admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRole indicates an expected call of DatabaseRole.
func (mr *MockDatabaseRoleUpdaterMockRecorder) DatabaseRole(arg0, arg1 any) *MockDatabaseRoleUpdaterDatabaseRoleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRole", reflect.TypeOf((*MockDatabaseRoleUpdater)(nil).DatabaseRole), arg0, arg1)
	return &MockDatabaseRoleUpdaterDatabaseRoleCall{Call: call}
}

// MockDatabaseRoleUpdaterDatabaseRoleCall wrap *gomock.Call
type MockDatabaseRoleUpdaterDatabaseRoleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDatabaseRoleUpdaterDatabaseRoleCall) Return(arg0 *admin.UserCustomDBRole, arg1 error) *MockDatabaseRoleUpdaterDatabaseRoleCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDatabaseRoleUpdaterDatabaseRoleCall) Do(f func(string, string) (*admin.UserCustomDBRole, error)) *MockDatabaseRoleUpdaterDatabaseRoleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDatabaseRoleUpdaterDatabaseRoleCall) DoAndReturn(f func(string, string) (*admin.UserCustomDBRole, error)) *MockDatabaseRoleUpdaterDatabaseRoleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateDatabaseRole mocks base method.
func (m *MockDatabaseRoleUpdater) UpdateDatabaseRole(arg0, arg1 string, arg2 *admin.UserCustomDBRole) (*admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDatabaseRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDatabaseRole indicates an expected call of UpdateDatabaseRole.
func (mr *MockDatabaseRoleUpdaterMockRecorder) UpdateDatabaseRole(arg0, arg1, arg2 any) *MockDatabaseRoleUpdaterUpdateDatabaseRoleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDatabaseRole", reflect.TypeOf((*MockDatabaseRoleUpdater)(nil).UpdateDatabaseRole), arg0, arg1, arg2)
	return &MockDatabaseRoleUpdaterUpdateDatabaseRoleCall{Call: call}
}

// MockDatabaseRoleUpdaterUpdateDatabaseRoleCall wrap *gomock.Call
type MockDatabaseRoleUpdaterUpdateDatabaseRoleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDatabaseRoleUpdaterUpdateDatabaseRoleCall) Return(arg0 *admin.UserCustomDBRole, arg1 error) *MockDatabaseRoleUpdaterUpdateDatabaseRoleCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDatabaseRoleUpdaterUpdateDatabaseRoleCall) Do(f func(string, string, *admin.UserCustomDBRole) (*admin.UserCustomDBRole, error)) *MockDatabaseRoleUpdaterUpdateDatabaseRoleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDatabaseRoleUpdaterUpdateDatabaseRoleCall) DoAndReturn(f func(string, string, *admin.UserCustomDBRole) (*admin.UserCustomDBRole, error)) *MockDatabaseRoleUpdaterUpdateDatabaseRoleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
