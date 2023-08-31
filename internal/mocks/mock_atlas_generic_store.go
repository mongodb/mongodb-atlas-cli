// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: AtlasOperatorGenericStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201004/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockAtlasOperatorGenericStore is a mock of AtlasOperatorGenericStore interface.
type MockAtlasOperatorGenericStore struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasOperatorGenericStoreMockRecorder
}

// MockAtlasOperatorGenericStoreMockRecorder is the mock recorder for MockAtlasOperatorGenericStore.
type MockAtlasOperatorGenericStoreMockRecorder struct {
	mock *MockAtlasOperatorGenericStore
}

// NewMockAtlasOperatorGenericStore creates a new mock instance.
func NewMockAtlasOperatorGenericStore(ctrl *gomock.Controller) *MockAtlasOperatorGenericStore {
	mock := &MockAtlasOperatorGenericStore{ctrl: ctrl}
	mock.recorder = &MockAtlasOperatorGenericStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtlasOperatorGenericStore) EXPECT() *MockAtlasOperatorGenericStoreMockRecorder {
	return m.recorder
}

// AlertConfigurations mocks base method.
func (m *MockAtlasOperatorGenericStore) AlertConfigurations(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.AlertConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertConfigurations", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.AlertConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertConfigurations indicates an expected call of AlertConfigurations.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) AlertConfigurations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertConfigurations", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).AlertConfigurations), arg0, arg1)
}

// AssignProjectAPIKey mocks base method.
func (m *MockAtlasOperatorGenericStore) AssignProjectAPIKey(arg0, arg1 string, arg2 *mongodbatlas.AssignAPIKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignProjectAPIKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignProjectAPIKey indicates an expected call of AssignProjectAPIKey.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) AssignProjectAPIKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignProjectAPIKey", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).AssignProjectAPIKey), arg0, arg1, arg2)
}

// AtlasCluster mocks base method.
func (m *MockAtlasOperatorGenericStore) AtlasCluster(arg0, arg1 string) (*admin.AdvancedClusterDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.AdvancedClusterDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).AtlasCluster), arg0, arg1)
}

// AtlasClusterConfigurationOptions mocks base method.
func (m *MockAtlasOperatorGenericStore) AtlasClusterConfigurationOptions(arg0, arg1 string) (*admin.ClusterDescriptionProcessArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasClusterConfigurationOptions", arg0, arg1)
	ret0, _ := ret[0].(*admin.ClusterDescriptionProcessArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasClusterConfigurationOptions indicates an expected call of AtlasClusterConfigurationOptions.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) AtlasClusterConfigurationOptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasClusterConfigurationOptions", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).AtlasClusterConfigurationOptions), arg0, arg1)
}

// Auditing mocks base method.
func (m *MockAtlasOperatorGenericStore) Auditing(arg0 string) (*admin.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auditing", arg0)
	ret0, _ := ret[0].(*admin.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auditing indicates an expected call of Auditing.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) Auditing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auditing", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).Auditing), arg0)
}

// CloudProviderAccessRoles mocks base method.
func (m *MockAtlasOperatorGenericStore) CloudProviderAccessRoles(arg0 string) (*admin.CloudProviderAccessRoles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(*admin.CloudProviderAccessRoles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderAccessRoles indicates an expected call of CloudProviderAccessRoles.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) CloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderAccessRoles", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).CloudProviderAccessRoles), arg0)
}

// CreateOrganizationAPIKey mocks base method.
func (m *MockAtlasOperatorGenericStore) CreateOrganizationAPIKey(arg0 string, arg1 *mongodbatlas.APIKeyInput) (*mongodbatlas.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganizationAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganizationAPIKey indicates an expected call of CreateOrganizationAPIKey.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) CreateOrganizationAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizationAPIKey", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).CreateOrganizationAPIKey), arg0, arg1)
}

// CreateProject mocks base method.
func (m *MockAtlasOperatorGenericStore) CreateProject(arg0, arg1, arg2 string, arg3 *bool, arg4 *mongodbatlas.CreateProjectOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) CreateProject(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).CreateProject), arg0, arg1, arg2, arg3, arg4)
}

// CreateProjectAPIKey mocks base method.
func (m *MockAtlasOperatorGenericStore) CreateProjectAPIKey(arg0 string, arg1 *mongodbatlas.APIKeyInput) (*mongodbatlas.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectAPIKey indicates an expected call of CreateProjectAPIKey.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) CreateProjectAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectAPIKey", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).CreateProjectAPIKey), arg0, arg1)
}

// DataFederation mocks base method.
func (m *MockAtlasOperatorGenericStore) DataFederation(arg0, arg1 string) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederation", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederation indicates an expected call of DataFederation.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) DataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederation", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).DataFederation), arg0, arg1)
}

// DataFederationList mocks base method.
func (m *MockAtlasOperatorGenericStore) DataFederationList(arg0 string) ([]admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationList", arg0)
	ret0, _ := ret[0].([]admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationList indicates an expected call of DataFederationList.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) DataFederationList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationList", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).DataFederationList), arg0)
}

// DatabaseRoles mocks base method.
func (m *MockAtlasOperatorGenericStore) DatabaseRoles(arg0 string) ([]admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRoles", arg0)
	ret0, _ := ret[0].([]admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRoles indicates an expected call of DatabaseRoles.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) DatabaseRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRoles", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).DatabaseRoles), arg0)
}

// DatabaseUsers mocks base method.
func (m *MockAtlasOperatorGenericStore) DatabaseUsers(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedApiAtlasDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUsers indicates an expected call of DatabaseUsers.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) DatabaseUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUsers", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).DatabaseUsers), arg0, arg1)
}

// DescribeSchedule mocks base method.
func (m *MockAtlasOperatorGenericStore) DescribeSchedule(arg0, arg1 string) (*admin.DiskBackupSnapshotSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSchedule", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedule indicates an expected call of DescribeSchedule.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) DescribeSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedule", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).DescribeSchedule), arg0, arg1)
}

// EncryptionAtRest mocks base method.
func (m *MockAtlasOperatorGenericStore) EncryptionAtRest(arg0 string) (*admin.EncryptionAtRest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptionAtRest", arg0)
	ret0, _ := ret[0].(*admin.EncryptionAtRest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptionAtRest indicates an expected call of EncryptionAtRest.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) EncryptionAtRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptionAtRest", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).EncryptionAtRest), arg0)
}

// GetOrgProjects mocks base method.
func (m *MockAtlasOperatorGenericStore) GetOrgProjects(arg0 string, arg1 *mongodbatlas.ProjectsListOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgProjects", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgProjects indicates an expected call of GetOrgProjects.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) GetOrgProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgProjects", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).GetOrgProjects), arg0, arg1)
}

// GetServerlessInstance mocks base method.
func (m *MockAtlasOperatorGenericStore) GetServerlessInstance(arg0, arg1 string) (*admin.ServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*admin.ServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerlessInstance indicates an expected call of GetServerlessInstance.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) GetServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerlessInstance", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).GetServerlessInstance), arg0, arg1)
}

// GlobalCluster mocks base method.
func (m *MockAtlasOperatorGenericStore) GlobalCluster(arg0, arg1 string) (*admin.GeoSharding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.GeoSharding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalCluster indicates an expected call of GlobalCluster.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) GlobalCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalCluster", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).GlobalCluster), arg0, arg1)
}

// Integrations mocks base method.
func (m *MockAtlasOperatorGenericStore) Integrations(arg0 string) (*admin.PaginatedIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Integrations", arg0)
	ret0, _ := ret[0].(*admin.PaginatedIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Integrations indicates an expected call of Integrations.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) Integrations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Integrations", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).Integrations), arg0)
}

// MaintenanceWindow mocks base method.
func (m *MockAtlasOperatorGenericStore) MaintenanceWindow(arg0 string) (*admin.GroupMaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaintenanceWindow", arg0)
	ret0, _ := ret[0].(*admin.GroupMaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceWindow indicates an expected call of MaintenanceWindow.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) MaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceWindow", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).MaintenanceWindow), arg0)
}

// PeeringConnections mocks base method.
func (m *MockAtlasOperatorGenericStore) PeeringConnections(arg0 string, arg1 *mongodbatlas.ContainersListOptions) ([]admin.BaseNetworkPeeringConnectionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnections", arg0, arg1)
	ret0, _ := ret[0].([]admin.BaseNetworkPeeringConnectionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnections indicates an expected call of PeeringConnections.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) PeeringConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnections", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).PeeringConnections), arg0, arg1)
}

// PrivateEndpoints mocks base method.
func (m *MockAtlasOperatorGenericStore) PrivateEndpoints(arg0, arg1 string) ([]admin.EndpointService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.EndpointService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoints indicates an expected call of PrivateEndpoints.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) PrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoints", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).PrivateEndpoints), arg0, arg1)
}

// Project mocks base method.
func (m *MockAtlasOperatorGenericStore) Project(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).Project), arg0)
}

// ProjectByName mocks base method.
func (m *MockAtlasOperatorGenericStore) ProjectByName(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectByName", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectByName indicates an expected call of ProjectByName.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ProjectByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectByName", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ProjectByName), arg0)
}

// ProjectClusters mocks base method.
func (m *MockAtlasOperatorGenericStore) ProjectClusters(arg0 string, arg1 *mongodbatlas.ListOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ProjectClusters), arg0, arg1)
}

// ProjectIPAccessLists mocks base method.
func (m *MockAtlasOperatorGenericStore) ProjectIPAccessLists(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedNetworkAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectIPAccessLists", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedNetworkAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectIPAccessLists indicates an expected call of ProjectIPAccessLists.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ProjectIPAccessLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectIPAccessLists", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ProjectIPAccessLists), arg0, arg1)
}

// ProjectSettings mocks base method.
func (m *MockAtlasOperatorGenericStore) ProjectSettings(arg0 string) (*admin.GroupSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*admin.GroupSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectSettings indicates an expected call of ProjectSettings.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ProjectSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectSettings", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ProjectSettings), arg0)
}

// ProjectTeams mocks base method.
func (m *MockAtlasOperatorGenericStore) ProjectTeams(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectTeams", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectTeams indicates an expected call of ProjectTeams.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ProjectTeams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectTeams", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ProjectTeams), arg0)
}

// Projects mocks base method.
func (m *MockAtlasOperatorGenericStore) Projects(arg0 *mongodbatlas.ListOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) Projects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).Projects), arg0)
}

// ServerlessInstance mocks base method.
func (m *MockAtlasOperatorGenericStore) ServerlessInstance(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstance indicates an expected call of ServerlessInstance.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstance", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ServerlessInstance), arg0, arg1)
}

// ServerlessInstances mocks base method.
func (m *MockAtlasOperatorGenericStore) ServerlessInstances(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstances", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstances indicates an expected call of ServerlessInstances.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ServerlessInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstances", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ServerlessInstances), arg0, arg1)
}

// ServerlessPrivateEndpoints mocks base method.
func (m *MockAtlasOperatorGenericStore) ServerlessPrivateEndpoints(arg0, arg1 string) ([]admin.ServerlessTenantEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessPrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.ServerlessTenantEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessPrivateEndpoints indicates an expected call of ServerlessPrivateEndpoints.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ServerlessPrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessPrivateEndpoints", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ServerlessPrivateEndpoints), arg0, arg1)
}

// ServiceVersion mocks base method.
func (m *MockAtlasOperatorGenericStore) ServiceVersion() (*mongodbatlas.ServiceVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceVersion")
	ret0, _ := ret[0].(*mongodbatlas.ServiceVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceVersion indicates an expected call of ServiceVersion.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) ServiceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceVersion", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).ServiceVersion))
}

// TeamByID mocks base method.
func (m *MockAtlasOperatorGenericStore) TeamByID(arg0, arg1 string) (*mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByID", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByID indicates an expected call of TeamByID.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) TeamByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByID", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).TeamByID), arg0, arg1)
}

// TeamByName mocks base method.
func (m *MockAtlasOperatorGenericStore) TeamByName(arg0, arg1 string) (*mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByName", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByName indicates an expected call of TeamByName.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) TeamByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByName", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).TeamByName), arg0, arg1)
}

// TeamUsers mocks base method.
func (m *MockAtlasOperatorGenericStore) TeamUsers(arg0, arg1 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamUsers", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamUsers indicates an expected call of TeamUsers.
func (mr *MockAtlasOperatorGenericStoreMockRecorder) TeamUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUsers", reflect.TypeOf((*MockAtlasOperatorGenericStore)(nil).TeamUsers), arg0, arg1)
}
