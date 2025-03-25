// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: OperatorGenericStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	admin "go.mongodb.org/atlas-sdk/v20240530005/admin"
	admin0 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

// MockOperatorGenericStore is a mock of OperatorGenericStore interface.
type MockOperatorGenericStore struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorGenericStoreMockRecorder
}

// MockOperatorGenericStoreMockRecorder is the mock recorder for MockOperatorGenericStore.
type MockOperatorGenericStoreMockRecorder struct {
	mock *MockOperatorGenericStore
}

// NewMockOperatorGenericStore creates a new mock instance.
func NewMockOperatorGenericStore(ctrl *gomock.Controller) *MockOperatorGenericStore {
	mock := &MockOperatorGenericStore{ctrl: ctrl}
	mock.recorder = &MockOperatorGenericStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorGenericStore) EXPECT() *MockOperatorGenericStoreMockRecorder {
	return m.recorder
}

// AlertConfigurations mocks base method.
func (m *MockOperatorGenericStore) AlertConfigurations(arg0 *admin0.ListAlertConfigurationsApiParams) (*admin0.PaginatedAlertConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertConfigurations", arg0)
	ret0, _ := ret[0].(*admin0.PaginatedAlertConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertConfigurations indicates an expected call of AlertConfigurations.
func (mr *MockOperatorGenericStoreMockRecorder) AlertConfigurations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertConfigurations", reflect.TypeOf((*MockOperatorGenericStore)(nil).AlertConfigurations), arg0)
}

// AssignProjectAPIKey mocks base method.
func (m *MockOperatorGenericStore) AssignProjectAPIKey(arg0, arg1 string, arg2 *admin0.UpdateAtlasProjectApiKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignProjectAPIKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignProjectAPIKey indicates an expected call of AssignProjectAPIKey.
func (mr *MockOperatorGenericStoreMockRecorder) AssignProjectAPIKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignProjectAPIKey", reflect.TypeOf((*MockOperatorGenericStore)(nil).AssignProjectAPIKey), arg0, arg1, arg2)
}

// AtlasCluster mocks base method.
func (m *MockOperatorGenericStore) AtlasCluster(arg0, arg1 string) (*admin.AdvancedClusterDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.AdvancedClusterDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster.
func (mr *MockOperatorGenericStoreMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockOperatorGenericStore)(nil).AtlasCluster), arg0, arg1)
}

// AtlasClusterConfigurationOptions mocks base method.
func (m *MockOperatorGenericStore) AtlasClusterConfigurationOptions(arg0, arg1 string) (*admin.ClusterDescriptionProcessArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasClusterConfigurationOptions", arg0, arg1)
	ret0, _ := ret[0].(*admin.ClusterDescriptionProcessArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasClusterConfigurationOptions indicates an expected call of AtlasClusterConfigurationOptions.
func (mr *MockOperatorGenericStoreMockRecorder) AtlasClusterConfigurationOptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasClusterConfigurationOptions", reflect.TypeOf((*MockOperatorGenericStore)(nil).AtlasClusterConfigurationOptions), arg0, arg1)
}

// Auditing mocks base method.
func (m *MockOperatorGenericStore) Auditing(arg0 string) (*admin0.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auditing", arg0)
	ret0, _ := ret[0].(*admin0.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auditing indicates an expected call of Auditing.
func (mr *MockOperatorGenericStoreMockRecorder) Auditing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auditing", reflect.TypeOf((*MockOperatorGenericStore)(nil).Auditing), arg0)
}

// CloudProviderAccessRoles mocks base method.
func (m *MockOperatorGenericStore) CloudProviderAccessRoles(arg0 string) (*admin0.CloudProviderAccessRoles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(*admin0.CloudProviderAccessRoles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderAccessRoles indicates an expected call of CloudProviderAccessRoles.
func (mr *MockOperatorGenericStoreMockRecorder) CloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderAccessRoles", reflect.TypeOf((*MockOperatorGenericStore)(nil).CloudProviderAccessRoles), arg0)
}

// CreateOrganizationAPIKey mocks base method.
func (m *MockOperatorGenericStore) CreateOrganizationAPIKey(arg0 string, arg1 *admin0.CreateAtlasOrganizationApiKey) (*admin0.ApiKeyUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganizationAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*admin0.ApiKeyUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganizationAPIKey indicates an expected call of CreateOrganizationAPIKey.
func (mr *MockOperatorGenericStoreMockRecorder) CreateOrganizationAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizationAPIKey", reflect.TypeOf((*MockOperatorGenericStore)(nil).CreateOrganizationAPIKey), arg0, arg1)
}

// CreateProject mocks base method.
func (m *MockOperatorGenericStore) CreateProject(arg0 *admin0.CreateProjectApiParams) (*admin0.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(*admin0.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockOperatorGenericStoreMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockOperatorGenericStore)(nil).CreateProject), arg0)
}

// CreateProjectAPIKey mocks base method.
func (m *MockOperatorGenericStore) CreateProjectAPIKey(arg0 string, arg1 *admin0.CreateAtlasProjectApiKey) (*admin0.ApiKeyUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*admin0.ApiKeyUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectAPIKey indicates an expected call of CreateProjectAPIKey.
func (mr *MockOperatorGenericStoreMockRecorder) CreateProjectAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectAPIKey", reflect.TypeOf((*MockOperatorGenericStore)(nil).CreateProjectAPIKey), arg0, arg1)
}

// DataFederation mocks base method.
func (m *MockOperatorGenericStore) DataFederation(arg0, arg1 string) (*admin0.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederation", arg0, arg1)
	ret0, _ := ret[0].(*admin0.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederation indicates an expected call of DataFederation.
func (mr *MockOperatorGenericStoreMockRecorder) DataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederation", reflect.TypeOf((*MockOperatorGenericStore)(nil).DataFederation), arg0, arg1)
}

// DataFederationList mocks base method.
func (m *MockOperatorGenericStore) DataFederationList(arg0 string) ([]admin0.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationList", arg0)
	ret0, _ := ret[0].([]admin0.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationList indicates an expected call of DataFederationList.
func (mr *MockOperatorGenericStoreMockRecorder) DataFederationList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationList", reflect.TypeOf((*MockOperatorGenericStore)(nil).DataFederationList), arg0)
}

// DatabaseRoles mocks base method.
func (m *MockOperatorGenericStore) DatabaseRoles(arg0 string) ([]admin0.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRoles", arg0)
	ret0, _ := ret[0].([]admin0.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRoles indicates an expected call of DatabaseRoles.
func (mr *MockOperatorGenericStoreMockRecorder) DatabaseRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRoles", reflect.TypeOf((*MockOperatorGenericStore)(nil).DatabaseRoles), arg0)
}

// DatabaseUsers mocks base method.
func (m *MockOperatorGenericStore) DatabaseUsers(arg0 string, arg1 *store.ListOptions) (*admin0.PaginatedApiAtlasDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin0.PaginatedApiAtlasDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUsers indicates an expected call of DatabaseUsers.
func (mr *MockOperatorGenericStoreMockRecorder) DatabaseUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUsers", reflect.TypeOf((*MockOperatorGenericStore)(nil).DatabaseUsers), arg0, arg1)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockOperatorGenericStore) DescribeCompliancePolicy(arg0 string) (*admin0.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin0.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockOperatorGenericStoreMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockOperatorGenericStore)(nil).DescribeCompliancePolicy), arg0)
}

// DescribeSchedule mocks base method.
func (m *MockOperatorGenericStore) DescribeSchedule(arg0, arg1 string) (*admin.DiskBackupSnapshotSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSchedule", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedule indicates an expected call of DescribeSchedule.
func (mr *MockOperatorGenericStoreMockRecorder) DescribeSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedule", reflect.TypeOf((*MockOperatorGenericStore)(nil).DescribeSchedule), arg0, arg1)
}

// EncryptionAtRest mocks base method.
func (m *MockOperatorGenericStore) EncryptionAtRest(arg0 string) (*admin0.EncryptionAtRest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptionAtRest", arg0)
	ret0, _ := ret[0].(*admin0.EncryptionAtRest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptionAtRest indicates an expected call of EncryptionAtRest.
func (mr *MockOperatorGenericStoreMockRecorder) EncryptionAtRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptionAtRest", reflect.TypeOf((*MockOperatorGenericStore)(nil).EncryptionAtRest), arg0)
}

// FederationSetting mocks base method.
func (m *MockOperatorGenericStore) FederationSetting(arg0 *admin0.GetFederationSettingsApiParams) (*admin0.OrgFederationSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederationSetting", arg0)
	ret0, _ := ret[0].(*admin0.OrgFederationSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FederationSetting indicates an expected call of FederationSetting.
func (mr *MockOperatorGenericStoreMockRecorder) FederationSetting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederationSetting", reflect.TypeOf((*MockOperatorGenericStore)(nil).FederationSetting), arg0)
}

// FlexCluster mocks base method.
func (m *MockOperatorGenericStore) FlexCluster(arg0, arg1 string) (*admin0.FlexClusterDescription20241113, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlexCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin0.FlexClusterDescription20241113)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FlexCluster indicates an expected call of FlexCluster.
func (mr *MockOperatorGenericStoreMockRecorder) FlexCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlexCluster", reflect.TypeOf((*MockOperatorGenericStore)(nil).FlexCluster), arg0, arg1)
}

// GetConnectedOrgConfig mocks base method.
func (m *MockOperatorGenericStore) GetConnectedOrgConfig(arg0 *admin0.GetConnectedOrgConfigApiParams) (*admin0.ConnectedOrgConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectedOrgConfig", arg0)
	ret0, _ := ret[0].(*admin0.ConnectedOrgConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectedOrgConfig indicates an expected call of GetConnectedOrgConfig.
func (mr *MockOperatorGenericStoreMockRecorder) GetConnectedOrgConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectedOrgConfig", reflect.TypeOf((*MockOperatorGenericStore)(nil).GetConnectedOrgConfig), arg0)
}

// GetOrgProjects mocks base method.
func (m *MockOperatorGenericStore) GetOrgProjects(arg0 string, arg1 *store.ListOptions) (*admin0.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgProjects", arg0, arg1)
	ret0, _ := ret[0].(*admin0.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgProjects indicates an expected call of GetOrgProjects.
func (mr *MockOperatorGenericStoreMockRecorder) GetOrgProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgProjects", reflect.TypeOf((*MockOperatorGenericStore)(nil).GetOrgProjects), arg0, arg1)
}

// GetServerlessInstance mocks base method.
func (m *MockOperatorGenericStore) GetServerlessInstance(arg0, arg1 string) (*admin.ServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*admin.ServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerlessInstance indicates an expected call of GetServerlessInstance.
func (mr *MockOperatorGenericStoreMockRecorder) GetServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerlessInstance", reflect.TypeOf((*MockOperatorGenericStore)(nil).GetServerlessInstance), arg0, arg1)
}

// GlobalCluster mocks base method.
func (m *MockOperatorGenericStore) GlobalCluster(arg0, arg1 string) (*admin.GeoSharding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.GeoSharding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalCluster indicates an expected call of GlobalCluster.
func (mr *MockOperatorGenericStoreMockRecorder) GlobalCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalCluster", reflect.TypeOf((*MockOperatorGenericStore)(nil).GlobalCluster), arg0, arg1)
}

// IdentityProvider mocks base method.
func (m *MockOperatorGenericStore) IdentityProvider(arg0 *admin0.GetIdentityProviderApiParams) (*admin0.FederationIdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityProvider", arg0)
	ret0, _ := ret[0].(*admin0.FederationIdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IdentityProvider indicates an expected call of IdentityProvider.
func (mr *MockOperatorGenericStoreMockRecorder) IdentityProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityProvider", reflect.TypeOf((*MockOperatorGenericStore)(nil).IdentityProvider), arg0)
}

// IdentityProviders mocks base method.
func (m *MockOperatorGenericStore) IdentityProviders(arg0 *admin0.ListIdentityProvidersApiParams) (*admin0.PaginatedFederationIdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityProviders", arg0)
	ret0, _ := ret[0].(*admin0.PaginatedFederationIdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IdentityProviders indicates an expected call of IdentityProviders.
func (mr *MockOperatorGenericStoreMockRecorder) IdentityProviders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityProviders", reflect.TypeOf((*MockOperatorGenericStore)(nil).IdentityProviders), arg0)
}

// Integrations mocks base method.
func (m *MockOperatorGenericStore) Integrations(arg0 string) (*admin0.PaginatedIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Integrations", arg0)
	ret0, _ := ret[0].(*admin0.PaginatedIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Integrations indicates an expected call of Integrations.
func (mr *MockOperatorGenericStoreMockRecorder) Integrations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Integrations", reflect.TypeOf((*MockOperatorGenericStore)(nil).Integrations), arg0)
}

// InterfaceEndpoint mocks base method.
func (m *MockOperatorGenericStore) InterfaceEndpoint(arg0, arg1, arg2, arg3 string) (*admin0.PrivateLinkEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterfaceEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin0.PrivateLinkEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InterfaceEndpoint indicates an expected call of InterfaceEndpoint.
func (mr *MockOperatorGenericStoreMockRecorder) InterfaceEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterfaceEndpoint", reflect.TypeOf((*MockOperatorGenericStore)(nil).InterfaceEndpoint), arg0, arg1, arg2, arg3)
}

// ListFlexClusters mocks base method.
func (m *MockOperatorGenericStore) ListFlexClusters(arg0 *admin0.ListFlexClustersApiParams) (*admin0.PaginatedFlexClusters20241113, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFlexClusters", arg0)
	ret0, _ := ret[0].(*admin0.PaginatedFlexClusters20241113)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFlexClusters indicates an expected call of ListFlexClusters.
func (mr *MockOperatorGenericStoreMockRecorder) ListFlexClusters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFlexClusters", reflect.TypeOf((*MockOperatorGenericStore)(nil).ListFlexClusters), arg0)
}

// MaintenanceWindow mocks base method.
func (m *MockOperatorGenericStore) MaintenanceWindow(arg0 string) (*admin0.GroupMaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaintenanceWindow", arg0)
	ret0, _ := ret[0].(*admin0.GroupMaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceWindow indicates an expected call of MaintenanceWindow.
func (mr *MockOperatorGenericStoreMockRecorder) MaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceWindow", reflect.TypeOf((*MockOperatorGenericStore)(nil).MaintenanceWindow), arg0)
}

// PeeringConnections mocks base method.
func (m *MockOperatorGenericStore) PeeringConnections(arg0 string, arg1 *store.ContainersListOptions) ([]admin0.BaseNetworkPeeringConnectionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnections", arg0, arg1)
	ret0, _ := ret[0].([]admin0.BaseNetworkPeeringConnectionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnections indicates an expected call of PeeringConnections.
func (mr *MockOperatorGenericStoreMockRecorder) PeeringConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnections", reflect.TypeOf((*MockOperatorGenericStore)(nil).PeeringConnections), arg0, arg1)
}

// PrivateEndpoints mocks base method.
func (m *MockOperatorGenericStore) PrivateEndpoints(arg0, arg1 string) ([]admin0.EndpointService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin0.EndpointService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoints indicates an expected call of PrivateEndpoints.
func (mr *MockOperatorGenericStoreMockRecorder) PrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoints", reflect.TypeOf((*MockOperatorGenericStore)(nil).PrivateEndpoints), arg0, arg1)
}

// Project mocks base method.
func (m *MockOperatorGenericStore) Project(arg0 string) (*admin0.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(*admin0.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockOperatorGenericStoreMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockOperatorGenericStore)(nil).Project), arg0)
}

// ProjectByName mocks base method.
func (m *MockOperatorGenericStore) ProjectByName(arg0 string) (*admin0.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectByName", arg0)
	ret0, _ := ret[0].(*admin0.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectByName indicates an expected call of ProjectByName.
func (mr *MockOperatorGenericStoreMockRecorder) ProjectByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectByName", reflect.TypeOf((*MockOperatorGenericStore)(nil).ProjectByName), arg0)
}

// ProjectClusters mocks base method.
func (m *MockOperatorGenericStore) ProjectClusters(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedAdvancedClusterDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAdvancedClusterDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters.
func (mr *MockOperatorGenericStoreMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockOperatorGenericStore)(nil).ProjectClusters), arg0, arg1)
}

// ProjectIPAccessLists mocks base method.
func (m *MockOperatorGenericStore) ProjectIPAccessLists(arg0 string, arg1 *store.ListOptions) (*admin0.PaginatedNetworkAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectIPAccessLists", arg0, arg1)
	ret0, _ := ret[0].(*admin0.PaginatedNetworkAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectIPAccessLists indicates an expected call of ProjectIPAccessLists.
func (mr *MockOperatorGenericStoreMockRecorder) ProjectIPAccessLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectIPAccessLists", reflect.TypeOf((*MockOperatorGenericStore)(nil).ProjectIPAccessLists), arg0, arg1)
}

// ProjectSettings mocks base method.
func (m *MockOperatorGenericStore) ProjectSettings(arg0 string) (*admin0.GroupSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*admin0.GroupSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectSettings indicates an expected call of ProjectSettings.
func (mr *MockOperatorGenericStoreMockRecorder) ProjectSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectSettings", reflect.TypeOf((*MockOperatorGenericStore)(nil).ProjectSettings), arg0)
}

// ProjectStreams mocks base method.
func (m *MockOperatorGenericStore) ProjectStreams(arg0 *admin0.ListStreamInstancesApiParams) (*admin0.PaginatedApiStreamsTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectStreams", arg0)
	ret0, _ := ret[0].(*admin0.PaginatedApiStreamsTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectStreams indicates an expected call of ProjectStreams.
func (mr *MockOperatorGenericStoreMockRecorder) ProjectStreams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectStreams", reflect.TypeOf((*MockOperatorGenericStore)(nil).ProjectStreams), arg0)
}

// ProjectTeams mocks base method.
func (m *MockOperatorGenericStore) ProjectTeams(arg0 string, arg1 *store.ListOptions) (*admin0.PaginatedTeamRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectTeams", arg0, arg1)
	ret0, _ := ret[0].(*admin0.PaginatedTeamRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectTeams indicates an expected call of ProjectTeams.
func (mr *MockOperatorGenericStoreMockRecorder) ProjectTeams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectTeams", reflect.TypeOf((*MockOperatorGenericStore)(nil).ProjectTeams), arg0, arg1)
}

// Projects mocks base method.
func (m *MockOperatorGenericStore) Projects(arg0 *store.ListOptions) (*admin0.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects", arg0)
	ret0, _ := ret[0].(*admin0.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockOperatorGenericStoreMockRecorder) Projects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockOperatorGenericStore)(nil).Projects), arg0)
}

// ServerlessInstances mocks base method.
func (m *MockOperatorGenericStore) ServerlessInstances(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstances", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstances indicates an expected call of ServerlessInstances.
func (mr *MockOperatorGenericStoreMockRecorder) ServerlessInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstances", reflect.TypeOf((*MockOperatorGenericStore)(nil).ServerlessInstances), arg0, arg1)
}

// ServerlessPrivateEndpoints mocks base method.
func (m *MockOperatorGenericStore) ServerlessPrivateEndpoints(arg0, arg1 string) ([]admin.ServerlessTenantEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessPrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.ServerlessTenantEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessPrivateEndpoints indicates an expected call of ServerlessPrivateEndpoints.
func (mr *MockOperatorGenericStoreMockRecorder) ServerlessPrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessPrivateEndpoints", reflect.TypeOf((*MockOperatorGenericStore)(nil).ServerlessPrivateEndpoints), arg0, arg1)
}

// StreamsConnections mocks base method.
func (m *MockOperatorGenericStore) StreamsConnections(arg0, arg1 string) (*admin0.PaginatedApiStreamsConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamsConnections", arg0, arg1)
	ret0, _ := ret[0].(*admin0.PaginatedApiStreamsConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamsConnections indicates an expected call of StreamsConnections.
func (mr *MockOperatorGenericStoreMockRecorder) StreamsConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamsConnections", reflect.TypeOf((*MockOperatorGenericStore)(nil).StreamsConnections), arg0, arg1)
}

// TeamByID mocks base method.
func (m *MockOperatorGenericStore) TeamByID(arg0, arg1 string) (*admin0.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByID", arg0, arg1)
	ret0, _ := ret[0].(*admin0.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByID indicates an expected call of TeamByID.
func (mr *MockOperatorGenericStoreMockRecorder) TeamByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByID", reflect.TypeOf((*MockOperatorGenericStore)(nil).TeamByID), arg0, arg1)
}

// TeamByName mocks base method.
func (m *MockOperatorGenericStore) TeamByName(arg0, arg1 string) (*admin0.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByName", arg0, arg1)
	ret0, _ := ret[0].(*admin0.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByName indicates an expected call of TeamByName.
func (mr *MockOperatorGenericStoreMockRecorder) TeamByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByName", reflect.TypeOf((*MockOperatorGenericStore)(nil).TeamByName), arg0, arg1)
}

// TeamUsers mocks base method.
func (m *MockOperatorGenericStore) TeamUsers(arg0, arg1 string) (*admin0.PaginatedOrgUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin0.PaginatedOrgUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamUsers indicates an expected call of TeamUsers.
func (mr *MockOperatorGenericStoreMockRecorder) TeamUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUsers", reflect.TypeOf((*MockOperatorGenericStore)(nil).TeamUsers), arg0, arg1)
}
