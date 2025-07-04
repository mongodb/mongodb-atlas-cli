// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters (interfaces: ClusterDescriber)
//
// Generated by this command:
//
//	mockgen -typed -destination=describe_mock_test.go -package=clusters . ClusterDescriber
//

// Package clusters is a generated GoMock package.
package clusters

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20240530005/admin"
	admin0 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockClusterDescriber is a mock of ClusterDescriber interface.
type MockClusterDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDescriberMockRecorder
	isgomock struct{}
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
func (mr *MockClusterDescriberMockRecorder) AtlasCluster(arg0, arg1 any) *MockClusterDescriberAtlasClusterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockClusterDescriber)(nil).AtlasCluster), arg0, arg1)
	return &MockClusterDescriberAtlasClusterCall{Call: call}
}

// MockClusterDescriberAtlasClusterCall wrap *gomock.Call
type MockClusterDescriberAtlasClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClusterDescriberAtlasClusterCall) Return(arg0 *admin.AdvancedClusterDescription, arg1 error) *MockClusterDescriberAtlasClusterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClusterDescriberAtlasClusterCall) Do(f func(string, string) (*admin.AdvancedClusterDescription, error)) *MockClusterDescriberAtlasClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClusterDescriberAtlasClusterCall) DoAndReturn(f func(string, string) (*admin.AdvancedClusterDescription, error)) *MockClusterDescriberAtlasClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FlexCluster mocks base method.
func (m *MockClusterDescriber) FlexCluster(arg0, arg1 string) (*admin0.FlexClusterDescription20241113, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlexCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin0.FlexClusterDescription20241113)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FlexCluster indicates an expected call of FlexCluster.
func (mr *MockClusterDescriberMockRecorder) FlexCluster(arg0, arg1 any) *MockClusterDescriberFlexClusterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlexCluster", reflect.TypeOf((*MockClusterDescriber)(nil).FlexCluster), arg0, arg1)
	return &MockClusterDescriberFlexClusterCall{Call: call}
}

// MockClusterDescriberFlexClusterCall wrap *gomock.Call
type MockClusterDescriberFlexClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClusterDescriberFlexClusterCall) Return(arg0 *admin0.FlexClusterDescription20241113, arg1 error) *MockClusterDescriberFlexClusterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClusterDescriberFlexClusterCall) Do(f func(string, string) (*admin0.FlexClusterDescription20241113, error)) *MockClusterDescriberFlexClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClusterDescriberFlexClusterCall) DoAndReturn(f func(string, string) (*admin0.FlexClusterDescription20241113, error)) *MockClusterDescriberFlexClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LatestAtlasCluster mocks base method.
func (m *MockClusterDescriber) LatestAtlasCluster(arg0, arg1 string) (*admin0.ClusterDescription20240805, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestAtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin0.ClusterDescription20240805)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestAtlasCluster indicates an expected call of LatestAtlasCluster.
func (mr *MockClusterDescriberMockRecorder) LatestAtlasCluster(arg0, arg1 any) *MockClusterDescriberLatestAtlasClusterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestAtlasCluster", reflect.TypeOf((*MockClusterDescriber)(nil).LatestAtlasCluster), arg0, arg1)
	return &MockClusterDescriberLatestAtlasClusterCall{Call: call}
}

// MockClusterDescriberLatestAtlasClusterCall wrap *gomock.Call
type MockClusterDescriberLatestAtlasClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClusterDescriberLatestAtlasClusterCall) Return(arg0 *admin0.ClusterDescription20240805, arg1 error) *MockClusterDescriberLatestAtlasClusterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClusterDescriberLatestAtlasClusterCall) Do(f func(string, string) (*admin0.ClusterDescription20240805, error)) *MockClusterDescriberLatestAtlasClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClusterDescriberLatestAtlasClusterCall) DoAndReturn(f func(string, string) (*admin0.ClusterDescription20240805, error)) *MockClusterDescriberLatestAtlasClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
