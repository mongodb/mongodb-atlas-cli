// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: ClusterLister,ClusterDescriber,ClusterConfigurationOptionsDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	atlas "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	admin "go.mongodb.org/atlas-sdk/v20231001002/admin"
)

// MockClusterLister is a mock of ClusterLister interface.
type MockClusterLister struct {
	ctrl     *gomock.Controller
	recorder *MockClusterListerMockRecorder
}

// MockClusterListerMockRecorder is the mock recorder for MockClusterLister.
type MockClusterListerMockRecorder struct {
	mock *MockClusterLister
}

// NewMockClusterLister creates a new mock instance.
func NewMockClusterLister(ctrl *gomock.Controller) *MockClusterLister {
	mock := &MockClusterLister{ctrl: ctrl}
	mock.recorder = &MockClusterListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterLister) EXPECT() *MockClusterListerMockRecorder {
	return m.recorder
}

// ProjectClusters mocks base method.
func (m *MockClusterLister) ProjectClusters(arg0 string, arg1 *atlas.ListOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters.
func (mr *MockClusterListerMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockClusterLister)(nil).ProjectClusters), arg0, arg1)
}

// MockClusterDescriber is a mock of ClusterDescriber interface.
type MockClusterDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDescriberMockRecorder
}

// MockClusterDescriberMockRecorder is the mock recorder for MockClusterDescriber.
type MockClusterDescriberMockRecorder struct {
	mock *MockClusterDescriber
}

// NewMockClusterDescriber creates a new mock instance.
func NewMockClusterDescriber(ctrl *gomock.Controller) *MockClusterDescriber {
	mock := &MockClusterDescriber{ctrl: ctrl}
	mock.recorder = &MockClusterDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterDescriber) EXPECT() *MockClusterDescriberMockRecorder {
	return m.recorder
}

// AtlasCluster mocks base method.
func (m *MockClusterDescriber) AtlasCluster(arg0, arg1 string) (*admin.AdvancedClusterDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.AdvancedClusterDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster.
func (mr *MockClusterDescriberMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockClusterDescriber)(nil).AtlasCluster), arg0, arg1)
}

// MockClusterConfigurationOptionsDescriber is a mock of ClusterConfigurationOptionsDescriber interface.
type MockClusterConfigurationOptionsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockClusterConfigurationOptionsDescriberMockRecorder
}

// MockClusterConfigurationOptionsDescriberMockRecorder is the mock recorder for MockClusterConfigurationOptionsDescriber.
type MockClusterConfigurationOptionsDescriberMockRecorder struct {
	mock *MockClusterConfigurationOptionsDescriber
}

// NewMockClusterConfigurationOptionsDescriber creates a new mock instance.
func NewMockClusterConfigurationOptionsDescriber(ctrl *gomock.Controller) *MockClusterConfigurationOptionsDescriber {
	mock := &MockClusterConfigurationOptionsDescriber{ctrl: ctrl}
	mock.recorder = &MockClusterConfigurationOptionsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterConfigurationOptionsDescriber) EXPECT() *MockClusterConfigurationOptionsDescriberMockRecorder {
	return m.recorder
}

// AtlasClusterConfigurationOptions mocks base method.
func (m *MockClusterConfigurationOptionsDescriber) AtlasClusterConfigurationOptions(arg0, arg1 string) (*admin.ClusterDescriptionProcessArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasClusterConfigurationOptions", arg0, arg1)
	ret0, _ := ret[0].(*admin.ClusterDescriptionProcessArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasClusterConfigurationOptions indicates an expected call of AtlasClusterConfigurationOptions.
func (mr *MockClusterConfigurationOptionsDescriberMockRecorder) AtlasClusterConfigurationOptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasClusterConfigurationOptions", reflect.TypeOf((*MockClusterConfigurationOptionsDescriber)(nil).AtlasClusterConfigurationOptions), arg0, arg1)
}
