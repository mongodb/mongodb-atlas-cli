// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: OperatorProjectStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	admin "go.mongodb.org/atlas-sdk/v20240530003/admin"
)

// MockOperatorProjectStore is a mock of OperatorProjectStore interface.
type MockOperatorProjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorProjectStoreMockRecorder
}

// MockOperatorProjectStoreMockRecorder is the mock recorder for MockOperatorProjectStore.
type MockOperatorProjectStoreMockRecorder struct {
	mock *MockOperatorProjectStore
}

// NewMockOperatorProjectStore creates a new mock instance.
func NewMockOperatorProjectStore(ctrl *gomock.Controller) *MockOperatorProjectStore {
	mock := &MockOperatorProjectStore{ctrl: ctrl}
	mock.recorder = &MockOperatorProjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorProjectStore) EXPECT() *MockOperatorProjectStoreMockRecorder {
	return m.recorder
}

// AlertConfigurations mocks base method.
func (m *MockOperatorProjectStore) AlertConfigurations(arg0 *admin.ListAlertConfigurationsApiParams) (*admin.PaginatedAlertConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertConfigurations", arg0)
	ret0, _ := ret[0].(*admin.PaginatedAlertConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertConfigurations indicates an expected call of AlertConfigurations.
func (mr *MockOperatorProjectStoreMockRecorder) AlertConfigurations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertConfigurations", reflect.TypeOf((*MockOperatorProjectStore)(nil).AlertConfigurations), arg0)
}

// Auditing mocks base method.
func (m *MockOperatorProjectStore) Auditing(arg0 string) (*admin.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auditing", arg0)
	ret0, _ := ret[0].(*admin.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auditing indicates an expected call of Auditing.
func (mr *MockOperatorProjectStoreMockRecorder) Auditing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auditing", reflect.TypeOf((*MockOperatorProjectStore)(nil).Auditing), arg0)
}

// CloudProviderAccessRoles mocks base method.
func (m *MockOperatorProjectStore) CloudProviderAccessRoles(arg0 string) (*admin.CloudProviderAccessRoles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(*admin.CloudProviderAccessRoles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderAccessRoles indicates an expected call of CloudProviderAccessRoles.
func (mr *MockOperatorProjectStoreMockRecorder) CloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderAccessRoles", reflect.TypeOf((*MockOperatorProjectStore)(nil).CloudProviderAccessRoles), arg0)
}

// CreateProject mocks base method.
func (m *MockOperatorProjectStore) CreateProject(arg0 *admin.CreateProjectApiParams) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockOperatorProjectStoreMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockOperatorProjectStore)(nil).CreateProject), arg0)
}

// CreateProjectAPIKey mocks base method.
func (m *MockOperatorProjectStore) CreateProjectAPIKey(arg0 string, arg1 *admin.CreateAtlasProjectApiKey) (*admin.ApiKeyUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiKeyUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectAPIKey indicates an expected call of CreateProjectAPIKey.
func (mr *MockOperatorProjectStoreMockRecorder) CreateProjectAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectAPIKey", reflect.TypeOf((*MockOperatorProjectStore)(nil).CreateProjectAPIKey), arg0, arg1)
}

// DatabaseRoles mocks base method.
func (m *MockOperatorProjectStore) DatabaseRoles(arg0 string) ([]admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRoles", arg0)
	ret0, _ := ret[0].([]admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRoles indicates an expected call of DatabaseRoles.
func (mr *MockOperatorProjectStoreMockRecorder) DatabaseRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRoles", reflect.TypeOf((*MockOperatorProjectStore)(nil).DatabaseRoles), arg0)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockOperatorProjectStore) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockOperatorProjectStoreMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockOperatorProjectStore)(nil).DescribeCompliancePolicy), arg0)
}

// EncryptionAtRest mocks base method.
func (m *MockOperatorProjectStore) EncryptionAtRest(arg0 string) (*admin.EncryptionAtRest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptionAtRest", arg0)
	ret0, _ := ret[0].(*admin.EncryptionAtRest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptionAtRest indicates an expected call of EncryptionAtRest.
func (mr *MockOperatorProjectStoreMockRecorder) EncryptionAtRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptionAtRest", reflect.TypeOf((*MockOperatorProjectStore)(nil).EncryptionAtRest), arg0)
}

// GetOrgProjects mocks base method.
func (m *MockOperatorProjectStore) GetOrgProjects(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgProjects", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgProjects indicates an expected call of GetOrgProjects.
func (mr *MockOperatorProjectStoreMockRecorder) GetOrgProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgProjects", reflect.TypeOf((*MockOperatorProjectStore)(nil).GetOrgProjects), arg0, arg1)
}

// Integrations mocks base method.
func (m *MockOperatorProjectStore) Integrations(arg0 string) (*admin.PaginatedIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Integrations", arg0)
	ret0, _ := ret[0].(*admin.PaginatedIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Integrations indicates an expected call of Integrations.
func (mr *MockOperatorProjectStoreMockRecorder) Integrations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Integrations", reflect.TypeOf((*MockOperatorProjectStore)(nil).Integrations), arg0)
}

// MaintenanceWindow mocks base method.
func (m *MockOperatorProjectStore) MaintenanceWindow(arg0 string) (*admin.GroupMaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaintenanceWindow", arg0)
	ret0, _ := ret[0].(*admin.GroupMaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceWindow indicates an expected call of MaintenanceWindow.
func (mr *MockOperatorProjectStoreMockRecorder) MaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceWindow", reflect.TypeOf((*MockOperatorProjectStore)(nil).MaintenanceWindow), arg0)
}

// PeeringConnections mocks base method.
func (m *MockOperatorProjectStore) PeeringConnections(arg0 string, arg1 *store.ContainersListOptions) ([]admin.BaseNetworkPeeringConnectionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnections", arg0, arg1)
	ret0, _ := ret[0].([]admin.BaseNetworkPeeringConnectionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnections indicates an expected call of PeeringConnections.
func (mr *MockOperatorProjectStoreMockRecorder) PeeringConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnections", reflect.TypeOf((*MockOperatorProjectStore)(nil).PeeringConnections), arg0, arg1)
}

// PrivateEndpoints mocks base method.
func (m *MockOperatorProjectStore) PrivateEndpoints(arg0, arg1 string) ([]admin.EndpointService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.EndpointService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoints indicates an expected call of PrivateEndpoints.
func (mr *MockOperatorProjectStoreMockRecorder) PrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoints", reflect.TypeOf((*MockOperatorProjectStore)(nil).PrivateEndpoints), arg0, arg1)
}

// Project mocks base method.
func (m *MockOperatorProjectStore) Project(arg0 string) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockOperatorProjectStoreMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockOperatorProjectStore)(nil).Project), arg0)
}

// ProjectByName mocks base method.
func (m *MockOperatorProjectStore) ProjectByName(arg0 string) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectByName", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectByName indicates an expected call of ProjectByName.
func (mr *MockOperatorProjectStoreMockRecorder) ProjectByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectByName", reflect.TypeOf((*MockOperatorProjectStore)(nil).ProjectByName), arg0)
}

// ProjectIPAccessLists mocks base method.
func (m *MockOperatorProjectStore) ProjectIPAccessLists(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedNetworkAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectIPAccessLists", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedNetworkAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectIPAccessLists indicates an expected call of ProjectIPAccessLists.
func (mr *MockOperatorProjectStoreMockRecorder) ProjectIPAccessLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectIPAccessLists", reflect.TypeOf((*MockOperatorProjectStore)(nil).ProjectIPAccessLists), arg0, arg1)
}

// ProjectSettings mocks base method.
func (m *MockOperatorProjectStore) ProjectSettings(arg0 string) (*admin.GroupSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*admin.GroupSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectSettings indicates an expected call of ProjectSettings.
func (mr *MockOperatorProjectStoreMockRecorder) ProjectSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectSettings", reflect.TypeOf((*MockOperatorProjectStore)(nil).ProjectSettings), arg0)
}

// ProjectTeams mocks base method.
func (m *MockOperatorProjectStore) ProjectTeams(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedTeamRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectTeams", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedTeamRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectTeams indicates an expected call of ProjectTeams.
func (mr *MockOperatorProjectStoreMockRecorder) ProjectTeams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectTeams", reflect.TypeOf((*MockOperatorProjectStore)(nil).ProjectTeams), arg0, arg1)
}

// Projects mocks base method.
func (m *MockOperatorProjectStore) Projects(arg0 *store.ListOptions) (*admin.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects", arg0)
	ret0, _ := ret[0].(*admin.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockOperatorProjectStoreMockRecorder) Projects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockOperatorProjectStore)(nil).Projects), arg0)
}

// TeamByID mocks base method.
func (m *MockOperatorProjectStore) TeamByID(arg0, arg1 string) (*admin.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByID", arg0, arg1)
	ret0, _ := ret[0].(*admin.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByID indicates an expected call of TeamByID.
func (mr *MockOperatorProjectStoreMockRecorder) TeamByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByID", reflect.TypeOf((*MockOperatorProjectStore)(nil).TeamByID), arg0, arg1)
}

// TeamByName mocks base method.
func (m *MockOperatorProjectStore) TeamByName(arg0, arg1 string) (*admin.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByName", arg0, arg1)
	ret0, _ := ret[0].(*admin.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByName indicates an expected call of TeamByName.
func (mr *MockOperatorProjectStoreMockRecorder) TeamByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByName", reflect.TypeOf((*MockOperatorProjectStore)(nil).TeamByName), arg0, arg1)
}

// TeamUsers mocks base method.
func (m *MockOperatorProjectStore) TeamUsers(arg0, arg1 string) (*admin.PaginatedApiAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamUsers indicates an expected call of TeamUsers.
func (mr *MockOperatorProjectStoreMockRecorder) TeamUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUsers", reflect.TypeOf((*MockOperatorProjectStore)(nil).TeamUsers), arg0, arg1)
}
