// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/accesslogs (interfaces: Lister)
//
// Generated by this command:
//
//	mockgen -typed -destination=list_mock_test.go -package=accesslogs . Lister
//

// Package accesslogs is a generated GoMock package.
package accesslogs

import (
	reflect "reflect"

	store "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockLister is a mock of Lister interface.
type MockLister struct {
	ctrl     *gomock.Controller
	recorder *MockListerMockRecorder
	isgomock struct{}
}

// MockListerMockRecorder is the mock recorder for MockLister.
type MockListerMockRecorder struct {
	mock *MockLister
}

// NewMockLister creates a new mock instance.
func NewMockLister(ctrl *gomock.Controller) *MockLister {
	mock := &MockLister{ctrl: ctrl}
	mock.recorder = &MockListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLister) EXPECT() *MockListerMockRecorder {
	return m.recorder
}

// AccessLogsByClusterName mocks base method.
func (m *MockLister) AccessLogsByClusterName(arg0, arg1 string, arg2 *store.AccessLogOptions) (*admin.MongoDBAccessLogsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogsByClusterName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.MongoDBAccessLogsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessLogsByClusterName indicates an expected call of AccessLogsByClusterName.
func (mr *MockListerMockRecorder) AccessLogsByClusterName(arg0, arg1, arg2 any) *MockListerAccessLogsByClusterNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogsByClusterName", reflect.TypeOf((*MockLister)(nil).AccessLogsByClusterName), arg0, arg1, arg2)
	return &MockListerAccessLogsByClusterNameCall{Call: call}
}

// MockListerAccessLogsByClusterNameCall wrap *gomock.Call
type MockListerAccessLogsByClusterNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockListerAccessLogsByClusterNameCall) Return(arg0 *admin.MongoDBAccessLogsList, arg1 error) *MockListerAccessLogsByClusterNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockListerAccessLogsByClusterNameCall) Do(f func(string, string, *store.AccessLogOptions) (*admin.MongoDBAccessLogsList, error)) *MockListerAccessLogsByClusterNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockListerAccessLogsByClusterNameCall) DoAndReturn(f func(string, string, *store.AccessLogOptions) (*admin.MongoDBAccessLogsList, error)) *MockListerAccessLogsByClusterNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AccessLogsByHostname mocks base method.
func (m *MockLister) AccessLogsByHostname(arg0, arg1 string, arg2 *store.AccessLogOptions) (*admin.MongoDBAccessLogsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogsByHostname", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.MongoDBAccessLogsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessLogsByHostname indicates an expected call of AccessLogsByHostname.
func (mr *MockListerMockRecorder) AccessLogsByHostname(arg0, arg1, arg2 any) *MockListerAccessLogsByHostnameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogsByHostname", reflect.TypeOf((*MockLister)(nil).AccessLogsByHostname), arg0, arg1, arg2)
	return &MockListerAccessLogsByHostnameCall{Call: call}
}

// MockListerAccessLogsByHostnameCall wrap *gomock.Call
type MockListerAccessLogsByHostnameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockListerAccessLogsByHostnameCall) Return(arg0 *admin.MongoDBAccessLogsList, arg1 error) *MockListerAccessLogsByHostnameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockListerAccessLogsByHostnameCall) Do(f func(string, string, *store.AccessLogOptions) (*admin.MongoDBAccessLogsList, error)) *MockListerAccessLogsByHostnameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockListerAccessLogsByHostnameCall) DoAndReturn(f func(string, string, *store.AccessLogOptions) (*admin.MongoDBAccessLogsList, error)) *MockListerAccessLogsByHostnameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
