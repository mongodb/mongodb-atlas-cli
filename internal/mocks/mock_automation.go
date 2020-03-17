// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/automation.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	opsmngr "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	reflect "reflect"
)

// MockAutomationGetter is a mock of AutomationGetter interface
type MockAutomationGetter struct {
	ctrl     *gomock.Controller
	recorder *MockAutomationGetterMockRecorder
}

// MockAutomationGetterMockRecorder is the mock recorder for MockAutomationGetter
type MockAutomationGetterMockRecorder struct {
	mock *MockAutomationGetter
}

// NewMockAutomationGetter creates a new mock instance
func NewMockAutomationGetter(ctrl *gomock.Controller) *MockAutomationGetter {
	mock := &MockAutomationGetter{ctrl: ctrl}
	mock.recorder = &MockAutomationGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAutomationGetter) EXPECT() *MockAutomationGetterMockRecorder {
	return m.recorder
}

// GetAutomationConfig mocks base method
func (m *MockAutomationGetter) GetAutomationConfig(arg0 string) (*opsmngr.AutomationConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutomationConfig", arg0)
	ret0, _ := ret[0].(*opsmngr.AutomationConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutomationConfig indicates an expected call of GetAutomationConfig
func (mr *MockAutomationGetterMockRecorder) GetAutomationConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutomationConfig", reflect.TypeOf((*MockAutomationGetter)(nil).GetAutomationConfig), arg0)
}

// MockAutomationUpdater is a mock of AutomationUpdater interface
type MockAutomationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAutomationUpdaterMockRecorder
}

// MockAutomationUpdaterMockRecorder is the mock recorder for MockAutomationUpdater
type MockAutomationUpdaterMockRecorder struct {
	mock *MockAutomationUpdater
}

// NewMockAutomationUpdater creates a new mock instance
func NewMockAutomationUpdater(ctrl *gomock.Controller) *MockAutomationUpdater {
	mock := &MockAutomationUpdater{ctrl: ctrl}
	mock.recorder = &MockAutomationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAutomationUpdater) EXPECT() *MockAutomationUpdaterMockRecorder {
	return m.recorder
}

// UpdateAutomationConfig mocks base method
func (m *MockAutomationUpdater) UpdateAutomationConfig(arg0 string, arg1 *opsmngr.AutomationConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAutomationConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAutomationConfig indicates an expected call of UpdateAutomationConfig
func (mr *MockAutomationUpdaterMockRecorder) UpdateAutomationConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAutomationConfig", reflect.TypeOf((*MockAutomationUpdater)(nil).UpdateAutomationConfig), arg0, arg1)
}

// MockAutomationStatusGetter is a mock of AutomationStatusGetter interface
type MockAutomationStatusGetter struct {
	ctrl     *gomock.Controller
	recorder *MockAutomationStatusGetterMockRecorder
}

// MockAutomationStatusGetterMockRecorder is the mock recorder for MockAutomationStatusGetter
type MockAutomationStatusGetterMockRecorder struct {
	mock *MockAutomationStatusGetter
}

// NewMockAutomationStatusGetter creates a new mock instance
func NewMockAutomationStatusGetter(ctrl *gomock.Controller) *MockAutomationStatusGetter {
	mock := &MockAutomationStatusGetter{ctrl: ctrl}
	mock.recorder = &MockAutomationStatusGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAutomationStatusGetter) EXPECT() *MockAutomationStatusGetterMockRecorder {
	return m.recorder
}

// GetAutomationConfigStatus mocks base method
func (m *MockAutomationStatusGetter) GetAutomationConfigStatus(arg0 string) (*opsmngr.AutomationStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutomationConfigStatus", arg0)
	ret0, _ := ret[0].(*opsmngr.AutomationStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutomationConfigStatus indicates an expected call of GetAutomationConfigStatus
func (mr *MockAutomationStatusGetterMockRecorder) GetAutomationConfigStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutomationConfigStatus", reflect.TypeOf((*MockAutomationStatusGetter)(nil).GetAutomationConfigStatus), arg0)
}

// MockAllClusterLister is a mock of AllClusterLister interface
type MockAllClusterLister struct {
	ctrl     *gomock.Controller
	recorder *MockAllClusterListerMockRecorder
}

// MockAllClusterListerMockRecorder is the mock recorder for MockAllClusterLister
type MockAllClusterListerMockRecorder struct {
	mock *MockAllClusterLister
}

// NewMockAllClusterLister creates a new mock instance
func NewMockAllClusterLister(ctrl *gomock.Controller) *MockAllClusterLister {
	mock := &MockAllClusterLister{ctrl: ctrl}
	mock.recorder = &MockAllClusterListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAllClusterLister) EXPECT() *MockAllClusterListerMockRecorder {
	return m.recorder
}

// ListAllClustersProjects mocks base method
func (m *MockAllClusterLister) ListAllClustersProjects() (*opsmngr.AllClustersProjects, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllClustersProjects")
	ret0, _ := ret[0].(*opsmngr.AllClustersProjects)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllClustersProjects indicates an expected call of ListAllClustersProjects
func (mr *MockAllClusterListerMockRecorder) ListAllClustersProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllClustersProjects", reflect.TypeOf((*MockAllClusterLister)(nil).ListAllClustersProjects))
}

// MockAutomationStore is a mock of AutomationStore interface
type MockAutomationStore struct {
	ctrl     *gomock.Controller
	recorder *MockAutomationStoreMockRecorder
}

// MockAutomationStoreMockRecorder is the mock recorder for MockAutomationStore
type MockAutomationStoreMockRecorder struct {
	mock *MockAutomationStore
}

// NewMockAutomationStore creates a new mock instance
func NewMockAutomationStore(ctrl *gomock.Controller) *MockAutomationStore {
	mock := &MockAutomationStore{ctrl: ctrl}
	mock.recorder = &MockAutomationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAutomationStore) EXPECT() *MockAutomationStoreMockRecorder {
	return m.recorder
}

// GetAutomationConfig mocks base method
func (m *MockAutomationStore) GetAutomationConfig(arg0 string) (*opsmngr.AutomationConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutomationConfig", arg0)
	ret0, _ := ret[0].(*opsmngr.AutomationConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutomationConfig indicates an expected call of GetAutomationConfig
func (mr *MockAutomationStoreMockRecorder) GetAutomationConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutomationConfig", reflect.TypeOf((*MockAutomationStore)(nil).GetAutomationConfig), arg0)
}

// UpdateAutomationConfig mocks base method
func (m *MockAutomationStore) UpdateAutomationConfig(arg0 string, arg1 *opsmngr.AutomationConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAutomationConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAutomationConfig indicates an expected call of UpdateAutomationConfig
func (mr *MockAutomationStoreMockRecorder) UpdateAutomationConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAutomationConfig", reflect.TypeOf((*MockAutomationStore)(nil).UpdateAutomationConfig), arg0, arg1)
}

// GetAutomationConfigStatus mocks base method
func (m *MockAutomationStore) GetAutomationConfigStatus(arg0 string) (*opsmngr.AutomationStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutomationConfigStatus", arg0)
	ret0, _ := ret[0].(*opsmngr.AutomationStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutomationConfigStatus indicates an expected call of GetAutomationConfigStatus
func (mr *MockAutomationStoreMockRecorder) GetAutomationConfigStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutomationConfigStatus", reflect.TypeOf((*MockAutomationStore)(nil).GetAutomationConfigStatus), arg0)
}

// MockCloudManagerClustersLister is a mock of CloudManagerClustersLister interface
type MockCloudManagerClustersLister struct {
	ctrl     *gomock.Controller
	recorder *MockCloudManagerClustersListerMockRecorder
}

// MockCloudManagerClustersListerMockRecorder is the mock recorder for MockCloudManagerClustersLister
type MockCloudManagerClustersListerMockRecorder struct {
	mock *MockCloudManagerClustersLister
}

// NewMockCloudManagerClustersLister creates a new mock instance
func NewMockCloudManagerClustersLister(ctrl *gomock.Controller) *MockCloudManagerClustersLister {
	mock := &MockCloudManagerClustersLister{ctrl: ctrl}
	mock.recorder = &MockCloudManagerClustersListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudManagerClustersLister) EXPECT() *MockCloudManagerClustersListerMockRecorder {
	return m.recorder
}

// GetAutomationConfig mocks base method
func (m *MockCloudManagerClustersLister) GetAutomationConfig(arg0 string) (*opsmngr.AutomationConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutomationConfig", arg0)
	ret0, _ := ret[0].(*opsmngr.AutomationConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutomationConfig indicates an expected call of GetAutomationConfig
func (mr *MockCloudManagerClustersListerMockRecorder) GetAutomationConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutomationConfig", reflect.TypeOf((*MockCloudManagerClustersLister)(nil).GetAutomationConfig), arg0)
}

// ListAllClustersProjects mocks base method
func (m *MockCloudManagerClustersLister) ListAllClustersProjects() (*opsmngr.AllClustersProjects, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllClustersProjects")
	ret0, _ := ret[0].(*opsmngr.AllClustersProjects)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllClustersProjects indicates an expected call of ListAllClustersProjects
func (mr *MockCloudManagerClustersListerMockRecorder) ListAllClustersProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllClustersProjects", reflect.TypeOf((*MockCloudManagerClustersLister)(nil).ListAllClustersProjects))
}
