// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: ClusterLister,AtlasClusterDescriber,OpsManagerClusterDescriber,ClusterCreator,ClusterDeleter,ClusterUpdater,AtlasClusterGetterUpdater,ClusterPauser,ClusterStarter,AtlasClusterQuickStarter,SampleDataAdder,SampleDataStatusDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
)

// MockClusterLister is a mock of ClusterLister interface
type MockClusterLister struct {
	ctrl     *gomock.Controller
	recorder *MockClusterListerMockRecorder
}

// MockClusterListerMockRecorder is the mock recorder for MockClusterLister
type MockClusterListerMockRecorder struct {
	mock *MockClusterLister
}

// NewMockClusterLister creates a new mock instance
func NewMockClusterLister(ctrl *gomock.Controller) *MockClusterLister {
	mock := &MockClusterLister{ctrl: ctrl}
	mock.recorder = &MockClusterListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterLister) EXPECT() *MockClusterListerMockRecorder {
	return m.recorder
}

// ProjectClusters mocks base method
func (m *MockClusterLister) ProjectClusters(arg0 string, arg1 *mongodbatlas.ListOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters
func (mr *MockClusterListerMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockClusterLister)(nil).ProjectClusters), arg0, arg1)
}

// MockAtlasClusterDescriber is a mock of AtlasClusterDescriber interface
type MockAtlasClusterDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasClusterDescriberMockRecorder
}

// MockAtlasClusterDescriberMockRecorder is the mock recorder for MockAtlasClusterDescriber
type MockAtlasClusterDescriberMockRecorder struct {
	mock *MockAtlasClusterDescriber
}

// NewMockAtlasClusterDescriber creates a new mock instance
func NewMockAtlasClusterDescriber(ctrl *gomock.Controller) *MockAtlasClusterDescriber {
	mock := &MockAtlasClusterDescriber{ctrl: ctrl}
	mock.recorder = &MockAtlasClusterDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAtlasClusterDescriber) EXPECT() *MockAtlasClusterDescriberMockRecorder {
	return m.recorder
}

// AtlasCluster mocks base method
func (m *MockAtlasClusterDescriber) AtlasCluster(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster
func (mr *MockAtlasClusterDescriberMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockAtlasClusterDescriber)(nil).AtlasCluster), arg0, arg1)
}

// MockOpsManagerClusterDescriber is a mock of OpsManagerClusterDescriber interface
type MockOpsManagerClusterDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockOpsManagerClusterDescriberMockRecorder
}

// MockOpsManagerClusterDescriberMockRecorder is the mock recorder for MockOpsManagerClusterDescriber
type MockOpsManagerClusterDescriberMockRecorder struct {
	mock *MockOpsManagerClusterDescriber
}

// NewMockOpsManagerClusterDescriber creates a new mock instance
func NewMockOpsManagerClusterDescriber(ctrl *gomock.Controller) *MockOpsManagerClusterDescriber {
	mock := &MockOpsManagerClusterDescriber{ctrl: ctrl}
	mock.recorder = &MockOpsManagerClusterDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpsManagerClusterDescriber) EXPECT() *MockOpsManagerClusterDescriberMockRecorder {
	return m.recorder
}

// OpsManagerCluster mocks base method
func (m *MockOpsManagerClusterDescriber) OpsManagerCluster(arg0, arg1 string) (*opsmngr.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpsManagerCluster", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpsManagerCluster indicates an expected call of OpsManagerCluster
func (mr *MockOpsManagerClusterDescriberMockRecorder) OpsManagerCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpsManagerCluster", reflect.TypeOf((*MockOpsManagerClusterDescriber)(nil).OpsManagerCluster), arg0, arg1)
}

// MockClusterCreator is a mock of ClusterCreator interface
type MockClusterCreator struct {
	ctrl     *gomock.Controller
	recorder *MockClusterCreatorMockRecorder
}

// MockClusterCreatorMockRecorder is the mock recorder for MockClusterCreator
type MockClusterCreatorMockRecorder struct {
	mock *MockClusterCreator
}

// NewMockClusterCreator creates a new mock instance
func NewMockClusterCreator(ctrl *gomock.Controller) *MockClusterCreator {
	mock := &MockClusterCreator{ctrl: ctrl}
	mock.recorder = &MockClusterCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterCreator) EXPECT() *MockClusterCreatorMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockClusterCreator) CreateCluster(arg0 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockClusterCreatorMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockClusterCreator)(nil).CreateCluster), arg0)
}

// MockClusterDeleter is a mock of ClusterDeleter interface
type MockClusterDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDeleterMockRecorder
}

// MockClusterDeleterMockRecorder is the mock recorder for MockClusterDeleter
type MockClusterDeleterMockRecorder struct {
	mock *MockClusterDeleter
}

// NewMockClusterDeleter creates a new mock instance
func NewMockClusterDeleter(ctrl *gomock.Controller) *MockClusterDeleter {
	mock := &MockClusterDeleter{ctrl: ctrl}
	mock.recorder = &MockClusterDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterDeleter) EXPECT() *MockClusterDeleterMockRecorder {
	return m.recorder
}

// DeleteCluster mocks base method
func (m *MockClusterDeleter) DeleteCluster(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockClusterDeleterMockRecorder) DeleteCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterDeleter)(nil).DeleteCluster), arg0, arg1)
}

// MockClusterUpdater is a mock of ClusterUpdater interface
type MockClusterUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockClusterUpdaterMockRecorder
}

// MockClusterUpdaterMockRecorder is the mock recorder for MockClusterUpdater
type MockClusterUpdaterMockRecorder struct {
	mock *MockClusterUpdater
}

// NewMockClusterUpdater creates a new mock instance
func NewMockClusterUpdater(ctrl *gomock.Controller) *MockClusterUpdater {
	mock := &MockClusterUpdater{ctrl: ctrl}
	mock.recorder = &MockClusterUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterUpdater) EXPECT() *MockClusterUpdaterMockRecorder {
	return m.recorder
}

// UpdateCluster mocks base method
func (m *MockClusterUpdater) UpdateCluster(arg0, arg1 string, arg2 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockClusterUpdaterMockRecorder) UpdateCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockClusterUpdater)(nil).UpdateCluster), arg0, arg1, arg2)
}

// MockAtlasClusterGetterUpdater is a mock of AtlasClusterGetterUpdater interface
type MockAtlasClusterGetterUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasClusterGetterUpdaterMockRecorder
}

// MockAtlasClusterGetterUpdaterMockRecorder is the mock recorder for MockAtlasClusterGetterUpdater
type MockAtlasClusterGetterUpdaterMockRecorder struct {
	mock *MockAtlasClusterGetterUpdater
}

// NewMockAtlasClusterGetterUpdater creates a new mock instance
func NewMockAtlasClusterGetterUpdater(ctrl *gomock.Controller) *MockAtlasClusterGetterUpdater {
	mock := &MockAtlasClusterGetterUpdater{ctrl: ctrl}
	mock.recorder = &MockAtlasClusterGetterUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAtlasClusterGetterUpdater) EXPECT() *MockAtlasClusterGetterUpdaterMockRecorder {
	return m.recorder
}

// AtlasCluster mocks base method
func (m *MockAtlasClusterGetterUpdater) AtlasCluster(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster
func (mr *MockAtlasClusterGetterUpdaterMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockAtlasClusterGetterUpdater)(nil).AtlasCluster), arg0, arg1)
}

// UpdateCluster mocks base method
func (m *MockAtlasClusterGetterUpdater) UpdateCluster(arg0, arg1 string, arg2 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockAtlasClusterGetterUpdaterMockRecorder) UpdateCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockAtlasClusterGetterUpdater)(nil).UpdateCluster), arg0, arg1, arg2)
}

// MockClusterPauser is a mock of ClusterPauser interface
type MockClusterPauser struct {
	ctrl     *gomock.Controller
	recorder *MockClusterPauserMockRecorder
}

// MockClusterPauserMockRecorder is the mock recorder for MockClusterPauser
type MockClusterPauserMockRecorder struct {
	mock *MockClusterPauser
}

// NewMockClusterPauser creates a new mock instance
func NewMockClusterPauser(ctrl *gomock.Controller) *MockClusterPauser {
	mock := &MockClusterPauser{ctrl: ctrl}
	mock.recorder = &MockClusterPauserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterPauser) EXPECT() *MockClusterPauserMockRecorder {
	return m.recorder
}

// PauseCluster mocks base method
func (m *MockClusterPauser) PauseCluster(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseCluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PauseCluster indicates an expected call of PauseCluster
func (mr *MockClusterPauserMockRecorder) PauseCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseCluster", reflect.TypeOf((*MockClusterPauser)(nil).PauseCluster), arg0, arg1)
}

// MockClusterStarter is a mock of ClusterStarter interface
type MockClusterStarter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStarterMockRecorder
}

// MockClusterStarterMockRecorder is the mock recorder for MockClusterStarter
type MockClusterStarterMockRecorder struct {
	mock *MockClusterStarter
}

// NewMockClusterStarter creates a new mock instance
func NewMockClusterStarter(ctrl *gomock.Controller) *MockClusterStarter {
	mock := &MockClusterStarter{ctrl: ctrl}
	mock.recorder = &MockClusterStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterStarter) EXPECT() *MockClusterStarterMockRecorder {
	return m.recorder
}

// StartCluster mocks base method
func (m *MockClusterStarter) StartCluster(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCluster indicates an expected call of StartCluster
func (mr *MockClusterStarterMockRecorder) StartCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCluster", reflect.TypeOf((*MockClusterStarter)(nil).StartCluster), arg0, arg1)
}

// MockAtlasClusterQuickStarter is a mock of AtlasClusterQuickStarter interface
type MockAtlasClusterQuickStarter struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasClusterQuickStarterMockRecorder
}

// MockAtlasClusterQuickStarterMockRecorder is the mock recorder for MockAtlasClusterQuickStarter
type MockAtlasClusterQuickStarterMockRecorder struct {
	mock *MockAtlasClusterQuickStarter
}

// NewMockAtlasClusterQuickStarter creates a new mock instance
func NewMockAtlasClusterQuickStarter(ctrl *gomock.Controller) *MockAtlasClusterQuickStarter {
	mock := &MockAtlasClusterQuickStarter{ctrl: ctrl}
	mock.recorder = &MockAtlasClusterQuickStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAtlasClusterQuickStarter) EXPECT() *MockAtlasClusterQuickStarterMockRecorder {
	return m.recorder
}

// AddSampleData mocks base method
func (m *MockAtlasClusterQuickStarter) AddSampleData(arg0, arg1 string) (*mongodbatlas.SampleDatasetJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSampleData", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.SampleDatasetJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSampleData indicates an expected call of AddSampleData
func (mr *MockAtlasClusterQuickStarterMockRecorder) AddSampleData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSampleData", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).AddSampleData), arg0, arg1)
}

// AtlasCluster mocks base method
func (m *MockAtlasClusterQuickStarter) AtlasCluster(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster
func (mr *MockAtlasClusterQuickStarterMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).AtlasCluster), arg0, arg1)
}

// CloudProviderRegions mocks base method
func (m *MockAtlasClusterQuickStarter) CloudProviderRegions(arg0, arg1 string, arg2 []*string) (*mongodbatlas.CloudProviders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderRegions", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.CloudProviders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderRegions indicates an expected call of CloudProviderRegions
func (mr *MockAtlasClusterQuickStarterMockRecorder) CloudProviderRegions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderRegions", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).CloudProviderRegions), arg0, arg1, arg2)
}

// CreateCluster mocks base method
func (m *MockAtlasClusterQuickStarter) CreateCluster(arg0 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockAtlasClusterQuickStarterMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).CreateCluster), arg0)
}

// CreateDatabaseUser mocks base method
func (m *MockAtlasClusterQuickStarter) CreateDatabaseUser(arg0 *mongodbatlas.DatabaseUser) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser
func (mr *MockAtlasClusterQuickStarterMockRecorder) CreateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).CreateDatabaseUser), arg0)
}

// CreateProjectIPAccessList mocks base method
func (m *MockAtlasClusterQuickStarter) CreateProjectIPAccessList(arg0 []*mongodbatlas.ProjectIPAccessList) (*mongodbatlas.ProjectIPAccessLists, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectIPAccessList", arg0)
	ret0, _ := ret[0].(*mongodbatlas.ProjectIPAccessLists)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectIPAccessList indicates an expected call of CreateProjectIPAccessList
func (mr *MockAtlasClusterQuickStarterMockRecorder) CreateProjectIPAccessList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectIPAccessList", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).CreateProjectIPAccessList), arg0)
}

// DatabaseUser mocks base method
func (m *MockAtlasClusterQuickStarter) DatabaseUser(arg0, arg1, arg2 string) (*mongodbatlas.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUser indicates an expected call of DatabaseUser
func (mr *MockAtlasClusterQuickStarterMockRecorder) DatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUser", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).DatabaseUser), arg0, arg1, arg2)
}

// ProjectClusters mocks base method
func (m *MockAtlasClusterQuickStarter) ProjectClusters(arg0 string, arg1 *mongodbatlas.ListOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters
func (mr *MockAtlasClusterQuickStarterMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).ProjectClusters), arg0, arg1)
}

// SampleDataStatus mocks base method
func (m *MockAtlasClusterQuickStarter) SampleDataStatus(arg0, arg1 string) (*mongodbatlas.SampleDatasetJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SampleDataStatus", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.SampleDatasetJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SampleDataStatus indicates an expected call of SampleDataStatus
func (mr *MockAtlasClusterQuickStarterMockRecorder) SampleDataStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleDataStatus", reflect.TypeOf((*MockAtlasClusterQuickStarter)(nil).SampleDataStatus), arg0, arg1)
}

// MockSampleDataAdder is a mock of SampleDataAdder interface
type MockSampleDataAdder struct {
	ctrl     *gomock.Controller
	recorder *MockSampleDataAdderMockRecorder
}

// MockSampleDataAdderMockRecorder is the mock recorder for MockSampleDataAdder
type MockSampleDataAdderMockRecorder struct {
	mock *MockSampleDataAdder
}

// NewMockSampleDataAdder creates a new mock instance
func NewMockSampleDataAdder(ctrl *gomock.Controller) *MockSampleDataAdder {
	mock := &MockSampleDataAdder{ctrl: ctrl}
	mock.recorder = &MockSampleDataAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSampleDataAdder) EXPECT() *MockSampleDataAdderMockRecorder {
	return m.recorder
}

// AddSampleData mocks base method
func (m *MockSampleDataAdder) AddSampleData(arg0, arg1 string) (*mongodbatlas.SampleDatasetJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSampleData", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.SampleDatasetJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSampleData indicates an expected call of AddSampleData
func (mr *MockSampleDataAdderMockRecorder) AddSampleData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSampleData", reflect.TypeOf((*MockSampleDataAdder)(nil).AddSampleData), arg0, arg1)
}

// MockSampleDataStatusDescriber is a mock of SampleDataStatusDescriber interface
type MockSampleDataStatusDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockSampleDataStatusDescriberMockRecorder
}

// MockSampleDataStatusDescriberMockRecorder is the mock recorder for MockSampleDataStatusDescriber
type MockSampleDataStatusDescriberMockRecorder struct {
	mock *MockSampleDataStatusDescriber
}

// NewMockSampleDataStatusDescriber creates a new mock instance
func NewMockSampleDataStatusDescriber(ctrl *gomock.Controller) *MockSampleDataStatusDescriber {
	mock := &MockSampleDataStatusDescriber{ctrl: ctrl}
	mock.recorder = &MockSampleDataStatusDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSampleDataStatusDescriber) EXPECT() *MockSampleDataStatusDescriberMockRecorder {
	return m.recorder
}

// SampleDataStatus mocks base method
func (m *MockSampleDataStatusDescriber) SampleDataStatus(arg0, arg1 string) (*mongodbatlas.SampleDatasetJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SampleDataStatus", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.SampleDatasetJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SampleDataStatus indicates an expected call of SampleDataStatus
func (mr *MockSampleDataStatusDescriberMockRecorder) SampleDataStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleDataStatus", reflect.TypeOf((*MockSampleDataStatusDescriber)(nil).SampleDataStatus), arg0, arg1)
}
