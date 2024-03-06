// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: ProjectLister,OrgProjectLister,ProjectCreator,ProjectDeleter,ProjectDescriber,ProjectUsersLister,ProjectUserDeleter,ProjectTeamLister,ProjectTeamAdder,ProjectTeamDeleter)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	atlas "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockProjectLister is a mock of ProjectLister interface.
type MockProjectLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectListerMockRecorder
}

// MockProjectListerMockRecorder is the mock recorder for MockProjectLister.
type MockProjectListerMockRecorder struct {
	mock *MockProjectLister
}

// NewMockProjectLister creates a new mock instance.
func NewMockProjectLister(ctrl *gomock.Controller) *MockProjectLister {
	mock := &MockProjectLister{ctrl: ctrl}
	mock.recorder = &MockProjectListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectLister) EXPECT() *MockProjectListerMockRecorder {
	return m.recorder
}

// Projects mocks base method.
func (m *MockProjectLister) Projects(arg0 *atlas.ListOptions) (*admin.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects", arg0)
	ret0, _ := ret[0].(*admin.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockProjectListerMockRecorder) Projects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockProjectLister)(nil).Projects), arg0)
}

// MockOrgProjectLister is a mock of OrgProjectLister interface.
type MockOrgProjectLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrgProjectListerMockRecorder
}

// MockOrgProjectListerMockRecorder is the mock recorder for MockOrgProjectLister.
type MockOrgProjectListerMockRecorder struct {
	mock *MockOrgProjectLister
}

// NewMockOrgProjectLister creates a new mock instance.
func NewMockOrgProjectLister(ctrl *gomock.Controller) *MockOrgProjectLister {
	mock := &MockOrgProjectLister{ctrl: ctrl}
	mock.recorder = &MockOrgProjectListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrgProjectLister) EXPECT() *MockOrgProjectListerMockRecorder {
	return m.recorder
}

// GetOrgProjects mocks base method.
func (m *MockOrgProjectLister) GetOrgProjects(arg0 string, arg1 *atlas.ListOptions) (*admin.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgProjects", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgProjects indicates an expected call of GetOrgProjects.
func (mr *MockOrgProjectListerMockRecorder) GetOrgProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgProjects", reflect.TypeOf((*MockOrgProjectLister)(nil).GetOrgProjects), arg0, arg1)
}

// Projects mocks base method.
func (m *MockOrgProjectLister) Projects(arg0 *atlas.ListOptions) (*admin.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects", arg0)
	ret0, _ := ret[0].(*admin.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockOrgProjectListerMockRecorder) Projects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockOrgProjectLister)(nil).Projects), arg0)
}

// MockProjectCreator is a mock of ProjectCreator interface.
type MockProjectCreator struct {
	ctrl     *gomock.Controller
	recorder *MockProjectCreatorMockRecorder
}

// MockProjectCreatorMockRecorder is the mock recorder for MockProjectCreator.
type MockProjectCreatorMockRecorder struct {
	mock *MockProjectCreator
}

// NewMockProjectCreator creates a new mock instance.
func NewMockProjectCreator(ctrl *gomock.Controller) *MockProjectCreator {
	mock := &MockProjectCreator{ctrl: ctrl}
	mock.recorder = &MockProjectCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectCreator) EXPECT() *MockProjectCreatorMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockProjectCreator) CreateProject(arg0 *admin.CreateProjectApiParams) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectCreatorMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectCreator)(nil).CreateProject), arg0)
}

// MockProjectDeleter is a mock of ProjectDeleter interface.
type MockProjectDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockProjectDeleterMockRecorder
}

// MockProjectDeleterMockRecorder is the mock recorder for MockProjectDeleter.
type MockProjectDeleterMockRecorder struct {
	mock *MockProjectDeleter
}

// NewMockProjectDeleter creates a new mock instance.
func NewMockProjectDeleter(ctrl *gomock.Controller) *MockProjectDeleter {
	mock := &MockProjectDeleter{ctrl: ctrl}
	mock.recorder = &MockProjectDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectDeleter) EXPECT() *MockProjectDeleterMockRecorder {
	return m.recorder
}

// DeleteProject mocks base method.
func (m *MockProjectDeleter) DeleteProject(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectDeleterMockRecorder) DeleteProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectDeleter)(nil).DeleteProject), arg0)
}

// MockProjectDescriber is a mock of ProjectDescriber interface.
type MockProjectDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockProjectDescriberMockRecorder
}

// MockProjectDescriberMockRecorder is the mock recorder for MockProjectDescriber.
type MockProjectDescriberMockRecorder struct {
	mock *MockProjectDescriber
}

// NewMockProjectDescriber creates a new mock instance.
func NewMockProjectDescriber(ctrl *gomock.Controller) *MockProjectDescriber {
	mock := &MockProjectDescriber{ctrl: ctrl}
	mock.recorder = &MockProjectDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectDescriber) EXPECT() *MockProjectDescriberMockRecorder {
	return m.recorder
}

// Project mocks base method.
func (m *MockProjectDescriber) Project(arg0 string) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockProjectDescriberMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockProjectDescriber)(nil).Project), arg0)
}

// ProjectByName mocks base method.
func (m *MockProjectDescriber) ProjectByName(arg0 string) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectByName", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectByName indicates an expected call of ProjectByName.
func (mr *MockProjectDescriberMockRecorder) ProjectByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectByName", reflect.TypeOf((*MockProjectDescriber)(nil).ProjectByName), arg0)
}

// MockProjectUsersLister is a mock of ProjectUsersLister interface.
type MockProjectUsersLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectUsersListerMockRecorder
}

// MockProjectUsersListerMockRecorder is the mock recorder for MockProjectUsersLister.
type MockProjectUsersListerMockRecorder struct {
	mock *MockProjectUsersLister
}

// NewMockProjectUsersLister creates a new mock instance.
func NewMockProjectUsersLister(ctrl *gomock.Controller) *MockProjectUsersLister {
	mock := &MockProjectUsersLister{ctrl: ctrl}
	mock.recorder = &MockProjectUsersListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectUsersLister) EXPECT() *MockProjectUsersListerMockRecorder {
	return m.recorder
}

// ProjectUsers mocks base method.
func (m *MockProjectUsersLister) ProjectUsers(arg0 string, arg1 *atlas.ListOptions) (*admin.PaginatedAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUsers indicates an expected call of ProjectUsers.
func (mr *MockProjectUsersListerMockRecorder) ProjectUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUsers", reflect.TypeOf((*MockProjectUsersLister)(nil).ProjectUsers), arg0, arg1)
}

// MockProjectUserDeleter is a mock of ProjectUserDeleter interface.
type MockProjectUserDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockProjectUserDeleterMockRecorder
}

// MockProjectUserDeleterMockRecorder is the mock recorder for MockProjectUserDeleter.
type MockProjectUserDeleterMockRecorder struct {
	mock *MockProjectUserDeleter
}

// NewMockProjectUserDeleter creates a new mock instance.
func NewMockProjectUserDeleter(ctrl *gomock.Controller) *MockProjectUserDeleter {
	mock := &MockProjectUserDeleter{ctrl: ctrl}
	mock.recorder = &MockProjectUserDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectUserDeleter) EXPECT() *MockProjectUserDeleterMockRecorder {
	return m.recorder
}

// DeleteUserFromProject mocks base method.
func (m *MockProjectUserDeleter) DeleteUserFromProject(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserFromProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserFromProject indicates an expected call of DeleteUserFromProject.
func (mr *MockProjectUserDeleterMockRecorder) DeleteUserFromProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserFromProject", reflect.TypeOf((*MockProjectUserDeleter)(nil).DeleteUserFromProject), arg0, arg1)
}

// MockProjectTeamLister is a mock of ProjectTeamLister interface.
type MockProjectTeamLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectTeamListerMockRecorder
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
func (m *MockProjectTeamLister) ProjectTeams(arg0 string) (*admin.PaginatedTeamRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectTeams", arg0)
	ret0, _ := ret[0].(*admin.PaginatedTeamRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectTeams indicates an expected call of ProjectTeams.
func (mr *MockProjectTeamListerMockRecorder) ProjectTeams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectTeams", reflect.TypeOf((*MockProjectTeamLister)(nil).ProjectTeams), arg0)
}

// MockProjectTeamAdder is a mock of ProjectTeamAdder interface.
type MockProjectTeamAdder struct {
	ctrl     *gomock.Controller
	recorder *MockProjectTeamAdderMockRecorder
}

// MockProjectTeamAdderMockRecorder is the mock recorder for MockProjectTeamAdder.
type MockProjectTeamAdderMockRecorder struct {
	mock *MockProjectTeamAdder
}

// NewMockProjectTeamAdder creates a new mock instance.
func NewMockProjectTeamAdder(ctrl *gomock.Controller) *MockProjectTeamAdder {
	mock := &MockProjectTeamAdder{ctrl: ctrl}
	mock.recorder = &MockProjectTeamAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectTeamAdder) EXPECT() *MockProjectTeamAdderMockRecorder {
	return m.recorder
}

// AddTeamsToProject mocks base method.
func (m *MockProjectTeamAdder) AddTeamsToProject(arg0 string, arg1 []admin.TeamRole) (*admin.PaginatedTeamRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTeamsToProject", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedTeamRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTeamsToProject indicates an expected call of AddTeamsToProject.
func (mr *MockProjectTeamAdderMockRecorder) AddTeamsToProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTeamsToProject", reflect.TypeOf((*MockProjectTeamAdder)(nil).AddTeamsToProject), arg0, arg1)
}

// MockProjectTeamDeleter is a mock of ProjectTeamDeleter interface.
type MockProjectTeamDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockProjectTeamDeleterMockRecorder
}

// MockProjectTeamDeleterMockRecorder is the mock recorder for MockProjectTeamDeleter.
type MockProjectTeamDeleterMockRecorder struct {
	mock *MockProjectTeamDeleter
}

// NewMockProjectTeamDeleter creates a new mock instance.
func NewMockProjectTeamDeleter(ctrl *gomock.Controller) *MockProjectTeamDeleter {
	mock := &MockProjectTeamDeleter{ctrl: ctrl}
	mock.recorder = &MockProjectTeamDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectTeamDeleter) EXPECT() *MockProjectTeamDeleterMockRecorder {
	return m.recorder
}

// DeleteTeamFromProject mocks base method.
func (m *MockProjectTeamDeleter) DeleteTeamFromProject(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeamFromProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeamFromProject indicates an expected call of DeleteTeamFromProject.
func (mr *MockProjectTeamDeleterMockRecorder) DeleteTeamFromProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeamFromProject", reflect.TypeOf((*MockProjectTeamDeleter)(nil).DeleteTeamFromProject), arg0, arg1)
}
