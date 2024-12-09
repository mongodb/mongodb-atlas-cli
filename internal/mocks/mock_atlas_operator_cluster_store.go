// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: OperatorClusterStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	admin "go.mongodb.org/atlas-sdk/v20240530005/admin"
	admin0 "go.mongodb.org/atlas-sdk/v20241113002/admin"
)

// MockOperatorClusterStore is a mock of OperatorClusterStore interface.
type MockOperatorClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorClusterStoreMockRecorder
}

// MockOperatorClusterStoreMockRecorder is the mock recorder for MockOperatorClusterStore.
type MockOperatorClusterStoreMockRecorder struct {
	mock *MockOperatorClusterStore
}

// NewMockOperatorClusterStore creates a new mock instance.
func NewMockOperatorClusterStore(ctrl *gomock.Controller) *MockOperatorClusterStore {
	mock := &MockOperatorClusterStore{ctrl: ctrl}
	mock.recorder = &MockOperatorClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorClusterStore) EXPECT() *MockOperatorClusterStoreMockRecorder {
	return m.recorder
}

// AtlasCluster mocks base method.
func (m *MockOperatorClusterStore) AtlasCluster(arg0, arg1 string) (*admin.AdvancedClusterDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.AdvancedClusterDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster.
func (mr *MockOperatorClusterStoreMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockOperatorClusterStore)(nil).AtlasCluster), arg0, arg1)
}

// AtlasClusterConfigurationOptions mocks base method.
func (m *MockOperatorClusterStore) AtlasClusterConfigurationOptions(arg0, arg1 string) (*admin.ClusterDescriptionProcessArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasClusterConfigurationOptions", arg0, arg1)
	ret0, _ := ret[0].(*admin.ClusterDescriptionProcessArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasClusterConfigurationOptions indicates an expected call of AtlasClusterConfigurationOptions.
func (mr *MockOperatorClusterStoreMockRecorder) AtlasClusterConfigurationOptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasClusterConfigurationOptions", reflect.TypeOf((*MockOperatorClusterStore)(nil).AtlasClusterConfigurationOptions), arg0, arg1)
}

// DescribeSchedule mocks base method.
func (m *MockOperatorClusterStore) DescribeSchedule(arg0, arg1 string) (*admin.DiskBackupSnapshotSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSchedule", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedule indicates an expected call of DescribeSchedule.
func (mr *MockOperatorClusterStoreMockRecorder) DescribeSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedule", reflect.TypeOf((*MockOperatorClusterStore)(nil).DescribeSchedule), arg0, arg1)
}

// FlexCluster mocks base method.
func (m *MockOperatorClusterStore) FlexCluster(arg0, arg1 string) (*admin0.FlexClusterDescription20241113, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlexCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin0.FlexClusterDescription20241113)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FlexCluster indicates an expected call of FlexCluster.
func (mr *MockOperatorClusterStoreMockRecorder) FlexCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlexCluster", reflect.TypeOf((*MockOperatorClusterStore)(nil).FlexCluster), arg0, arg1)
}

// GetServerlessInstance mocks base method.
func (m *MockOperatorClusterStore) GetServerlessInstance(arg0, arg1 string) (*admin.ServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*admin.ServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerlessInstance indicates an expected call of GetServerlessInstance.
func (mr *MockOperatorClusterStoreMockRecorder) GetServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerlessInstance", reflect.TypeOf((*MockOperatorClusterStore)(nil).GetServerlessInstance), arg0, arg1)
}

// GlobalCluster mocks base method.
func (m *MockOperatorClusterStore) GlobalCluster(arg0, arg1 string) (*admin.GeoSharding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.GeoSharding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalCluster indicates an expected call of GlobalCluster.
func (mr *MockOperatorClusterStoreMockRecorder) GlobalCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalCluster", reflect.TypeOf((*MockOperatorClusterStore)(nil).GlobalCluster), arg0, arg1)
}

// ListFlexClusters mocks base method.
func (m *MockOperatorClusterStore) ListFlexClusters(arg0 *admin0.ListFlexClustersApiParams) (*admin0.PaginatedFlexClusters20241113, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFlexClusters", arg0)
	ret0, _ := ret[0].(*admin0.PaginatedFlexClusters20241113)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFlexClusters indicates an expected call of ListFlexClusters.
func (mr *MockOperatorClusterStoreMockRecorder) ListFlexClusters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFlexClusters", reflect.TypeOf((*MockOperatorClusterStore)(nil).ListFlexClusters), arg0)
}

// ProjectClusters mocks base method.
func (m *MockOperatorClusterStore) ProjectClusters(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedAdvancedClusterDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAdvancedClusterDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters.
func (mr *MockOperatorClusterStoreMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockOperatorClusterStore)(nil).ProjectClusters), arg0, arg1)
}

// ServerlessInstances mocks base method.
func (m *MockOperatorClusterStore) ServerlessInstances(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstances", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstances indicates an expected call of ServerlessInstances.
func (mr *MockOperatorClusterStoreMockRecorder) ServerlessInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstances", reflect.TypeOf((*MockOperatorClusterStore)(nil).ServerlessInstances), arg0, arg1)
}

// ServerlessPrivateEndpoints mocks base method.
func (m *MockOperatorClusterStore) ServerlessPrivateEndpoints(arg0, arg1 string) ([]admin.ServerlessTenantEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessPrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.ServerlessTenantEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessPrivateEndpoints indicates an expected call of ServerlessPrivateEndpoints.
func (mr *MockOperatorClusterStoreMockRecorder) ServerlessPrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessPrivateEndpoints", reflect.TypeOf((*MockOperatorClusterStore)(nil).ServerlessPrivateEndpoints), arg0, arg1)
}
