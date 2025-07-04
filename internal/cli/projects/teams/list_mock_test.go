// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/projects/teams (interfaces: ProjectTeamLister)
//
// Generated by this command:
//
//	mockgen -typed -destination=list_mock_test.go -package=teams . ProjectTeamLister
//

// Package teams is a generated GoMock package.
package teams

import (
	reflect "reflect"

	store "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectTeamLister is a mock of ProjectTeamLister interface.
type MockProjectTeamLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectTeamListerMockRecorder
	isgomock struct{}
}

// MockProjectTeamListerMockRecorder is the mock recorder for MockProjectTeamLister.
type MockProjectTeamListerMockRecorder struct {
	mock *MockProjectTeamLister
}

// NewMockProjectTeamLister creates a new mock instance.
func NewMockProjectTeamLister(ctrl *gomock.Controller) *MockProjectTeamLister {
	mock := &MockProjectTeamLister{ctrl: ctrl}
	mock.recorder = &MockProjectTeamListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectTeamLister) EXPECT() *MockProjectTeamListerMockRecorder {
	return m.recorder
}

// ProjectTeams mocks base method.
func (m *MockProjectTeamLister) ProjectTeams(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedTeamRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectTeams", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedTeamRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectTeams indicates an expected call of ProjectTeams.
func (mr *MockProjectTeamListerMockRecorder) ProjectTeams(arg0, arg1 any) *MockProjectTeamListerProjectTeamsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectTeams", reflect.TypeOf((*MockProjectTeamLister)(nil).ProjectTeams), arg0, arg1)
	return &MockProjectTeamListerProjectTeamsCall{Call: call}
}

// MockProjectTeamListerProjectTeamsCall wrap *gomock.Call
type MockProjectTeamListerProjectTeamsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectTeamListerProjectTeamsCall) Return(arg0 *admin.PaginatedTeamRole, arg1 error) *MockProjectTeamListerProjectTeamsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectTeamListerProjectTeamsCall) Do(f func(string, *store.ListOptions) (*admin.PaginatedTeamRole, error)) *MockProjectTeamListerProjectTeamsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectTeamListerProjectTeamsCall) DoAndReturn(f func(string, *store.ListOptions) (*admin.PaginatedTeamRole, error)) *MockProjectTeamListerProjectTeamsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
