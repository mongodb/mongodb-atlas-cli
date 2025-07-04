// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters (interfaces: ClusterStarter)
//
// Generated by this command:
//
//	mockgen -typed -destination=start_mock_test.go -package=clusters . ClusterStarter
//

// Package clusters is a generated GoMock package.
package clusters

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20240530005/admin"
	admin0 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockClusterStarter is a mock of ClusterStarter interface.
type MockClusterStarter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStarterMockRecorder
	isgomock struct{}
}

// MockClusterStarterMockRecorder is the mock recorder for MockClusterStarter.
type MockClusterStarterMockRecorder struct {
	mock *MockClusterStarter
}

// NewMockClusterStarter creates a new mock instance.
func NewMockClusterStarter(ctrl *gomock.Controller) *MockClusterStarter {
	mock := &MockClusterStarter{ctrl: ctrl}
	mock.recorder = &MockClusterStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterStarter) EXPECT() *MockClusterStarterMockRecorder {
	return m.recorder
}

// StartCluster mocks base method.
func (m *MockClusterStarter) StartCluster(arg0, arg1 string) (*admin.AdvancedClusterDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.AdvancedClusterDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCluster indicates an expected call of StartCluster.
func (mr *MockClusterStarterMockRecorder) StartCluster(arg0, arg1 any) *MockClusterStarterStartClusterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCluster", reflect.TypeOf((*MockClusterStarter)(nil).StartCluster), arg0, arg1)
	return &MockClusterStarterStartClusterCall{Call: call}
}

// MockClusterStarterStartClusterCall wrap *gomock.Call
type MockClusterStarterStartClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClusterStarterStartClusterCall) Return(arg0 *admin.AdvancedClusterDescription, arg1 error) *MockClusterStarterStartClusterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClusterStarterStartClusterCall) Do(f func(string, string) (*admin.AdvancedClusterDescription, error)) *MockClusterStarterStartClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClusterStarterStartClusterCall) DoAndReturn(f func(string, string) (*admin.AdvancedClusterDescription, error)) *MockClusterStarterStartClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StartClusterLatest mocks base method.
func (m *MockClusterStarter) StartClusterLatest(arg0, arg1 string) (*admin0.ClusterDescription20240805, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartClusterLatest", arg0, arg1)
	ret0, _ := ret[0].(*admin0.ClusterDescription20240805)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartClusterLatest indicates an expected call of StartClusterLatest.
func (mr *MockClusterStarterMockRecorder) StartClusterLatest(arg0, arg1 any) *MockClusterStarterStartClusterLatestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartClusterLatest", reflect.TypeOf((*MockClusterStarter)(nil).StartClusterLatest), arg0, arg1)
	return &MockClusterStarterStartClusterLatestCall{Call: call}
}

// MockClusterStarterStartClusterLatestCall wrap *gomock.Call
type MockClusterStarterStartClusterLatestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClusterStarterStartClusterLatestCall) Return(arg0 *admin0.ClusterDescription20240805, arg1 error) *MockClusterStarterStartClusterLatestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClusterStarterStartClusterLatestCall) Do(f func(string, string) (*admin0.ClusterDescription20240805, error)) *MockClusterStarterStartClusterLatestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClusterStarterStartClusterLatestCall) DoAndReturn(f func(string, string) (*admin0.ClusterDescription20240805, error)) *MockClusterStarterStartClusterLatestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
