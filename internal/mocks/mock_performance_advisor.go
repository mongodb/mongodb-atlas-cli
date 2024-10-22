// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: PerformanceAdvisorNamespacesLister,PerformanceAdvisorSlowQueriesLister,PerformanceAdvisorIndexesLister,PerformanceAdvisorSlowOperationThresholdEnabler,PerformanceAdvisorSlowOperationThresholdDisabler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

// MockPerformanceAdvisorNamespacesLister is a mock of PerformanceAdvisorNamespacesLister interface.
type MockPerformanceAdvisorNamespacesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorNamespacesListerMockRecorder
}

// MockPerformanceAdvisorNamespacesListerMockRecorder is the mock recorder for MockPerformanceAdvisorNamespacesLister.
type MockPerformanceAdvisorNamespacesListerMockRecorder struct {
	mock *MockPerformanceAdvisorNamespacesLister
}

// NewMockPerformanceAdvisorNamespacesLister creates a new mock instance.
func NewMockPerformanceAdvisorNamespacesLister(ctrl *gomock.Controller) *MockPerformanceAdvisorNamespacesLister {
	mock := &MockPerformanceAdvisorNamespacesLister{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorNamespacesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorNamespacesLister) EXPECT() *MockPerformanceAdvisorNamespacesListerMockRecorder {
	return m.recorder
}

// PerformanceAdvisorNamespaces mocks base method.
func (m *MockPerformanceAdvisorNamespacesLister) PerformanceAdvisorNamespaces(arg0 *admin.ListSlowQueryNamespacesApiParams) (*admin.Namespaces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformanceAdvisorNamespaces", arg0)
	ret0, _ := ret[0].(*admin.Namespaces)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformanceAdvisorNamespaces indicates an expected call of PerformanceAdvisorNamespaces.
func (mr *MockPerformanceAdvisorNamespacesListerMockRecorder) PerformanceAdvisorNamespaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformanceAdvisorNamespaces", reflect.TypeOf((*MockPerformanceAdvisorNamespacesLister)(nil).PerformanceAdvisorNamespaces), arg0)
}

// MockPerformanceAdvisorSlowQueriesLister is a mock of PerformanceAdvisorSlowQueriesLister interface.
type MockPerformanceAdvisorSlowQueriesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorSlowQueriesListerMockRecorder
}

// MockPerformanceAdvisorSlowQueriesListerMockRecorder is the mock recorder for MockPerformanceAdvisorSlowQueriesLister.
type MockPerformanceAdvisorSlowQueriesListerMockRecorder struct {
	mock *MockPerformanceAdvisorSlowQueriesLister
}

// NewMockPerformanceAdvisorSlowQueriesLister creates a new mock instance.
func NewMockPerformanceAdvisorSlowQueriesLister(ctrl *gomock.Controller) *MockPerformanceAdvisorSlowQueriesLister {
	mock := &MockPerformanceAdvisorSlowQueriesLister{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorSlowQueriesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorSlowQueriesLister) EXPECT() *MockPerformanceAdvisorSlowQueriesListerMockRecorder {
	return m.recorder
}

// PerformanceAdvisorSlowQueries mocks base method.
func (m *MockPerformanceAdvisorSlowQueriesLister) PerformanceAdvisorSlowQueries(arg0 *admin.ListSlowQueriesApiParams) (*admin.PerformanceAdvisorSlowQueryList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformanceAdvisorSlowQueries", arg0)
	ret0, _ := ret[0].(*admin.PerformanceAdvisorSlowQueryList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformanceAdvisorSlowQueries indicates an expected call of PerformanceAdvisorSlowQueries.
func (mr *MockPerformanceAdvisorSlowQueriesListerMockRecorder) PerformanceAdvisorSlowQueries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformanceAdvisorSlowQueries", reflect.TypeOf((*MockPerformanceAdvisorSlowQueriesLister)(nil).PerformanceAdvisorSlowQueries), arg0)
}

// MockPerformanceAdvisorIndexesLister is a mock of PerformanceAdvisorIndexesLister interface.
type MockPerformanceAdvisorIndexesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorIndexesListerMockRecorder
}

// MockPerformanceAdvisorIndexesListerMockRecorder is the mock recorder for MockPerformanceAdvisorIndexesLister.
type MockPerformanceAdvisorIndexesListerMockRecorder struct {
	mock *MockPerformanceAdvisorIndexesLister
}

// NewMockPerformanceAdvisorIndexesLister creates a new mock instance.
func NewMockPerformanceAdvisorIndexesLister(ctrl *gomock.Controller) *MockPerformanceAdvisorIndexesLister {
	mock := &MockPerformanceAdvisorIndexesLister{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorIndexesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorIndexesLister) EXPECT() *MockPerformanceAdvisorIndexesListerMockRecorder {
	return m.recorder
}

// PerformanceAdvisorIndexes mocks base method.
func (m *MockPerformanceAdvisorIndexesLister) PerformanceAdvisorIndexes(arg0 *admin.ListSuggestedIndexesApiParams) (*admin.PerformanceAdvisorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformanceAdvisorIndexes", arg0)
	ret0, _ := ret[0].(*admin.PerformanceAdvisorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformanceAdvisorIndexes indicates an expected call of PerformanceAdvisorIndexes.
func (mr *MockPerformanceAdvisorIndexesListerMockRecorder) PerformanceAdvisorIndexes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformanceAdvisorIndexes", reflect.TypeOf((*MockPerformanceAdvisorIndexesLister)(nil).PerformanceAdvisorIndexes), arg0)
}

// MockPerformanceAdvisorSlowOperationThresholdEnabler is a mock of PerformanceAdvisorSlowOperationThresholdEnabler interface.
type MockPerformanceAdvisorSlowOperationThresholdEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorSlowOperationThresholdEnablerMockRecorder
}

// MockPerformanceAdvisorSlowOperationThresholdEnablerMockRecorder is the mock recorder for MockPerformanceAdvisorSlowOperationThresholdEnabler.
type MockPerformanceAdvisorSlowOperationThresholdEnablerMockRecorder struct {
	mock *MockPerformanceAdvisorSlowOperationThresholdEnabler
}

// NewMockPerformanceAdvisorSlowOperationThresholdEnabler creates a new mock instance.
func NewMockPerformanceAdvisorSlowOperationThresholdEnabler(ctrl *gomock.Controller) *MockPerformanceAdvisorSlowOperationThresholdEnabler {
	mock := &MockPerformanceAdvisorSlowOperationThresholdEnabler{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorSlowOperationThresholdEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorSlowOperationThresholdEnabler) EXPECT() *MockPerformanceAdvisorSlowOperationThresholdEnablerMockRecorder {
	return m.recorder
}

// EnablePerformanceAdvisorSlowOperationThreshold mocks base method.
func (m *MockPerformanceAdvisorSlowOperationThresholdEnabler) EnablePerformanceAdvisorSlowOperationThreshold(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePerformanceAdvisorSlowOperationThreshold", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnablePerformanceAdvisorSlowOperationThreshold indicates an expected call of EnablePerformanceAdvisorSlowOperationThreshold.
func (mr *MockPerformanceAdvisorSlowOperationThresholdEnablerMockRecorder) EnablePerformanceAdvisorSlowOperationThreshold(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePerformanceAdvisorSlowOperationThreshold", reflect.TypeOf((*MockPerformanceAdvisorSlowOperationThresholdEnabler)(nil).EnablePerformanceAdvisorSlowOperationThreshold), arg0)
}

// MockPerformanceAdvisorSlowOperationThresholdDisabler is a mock of PerformanceAdvisorSlowOperationThresholdDisabler interface.
type MockPerformanceAdvisorSlowOperationThresholdDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorSlowOperationThresholdDisablerMockRecorder
}

// MockPerformanceAdvisorSlowOperationThresholdDisablerMockRecorder is the mock recorder for MockPerformanceAdvisorSlowOperationThresholdDisabler.
type MockPerformanceAdvisorSlowOperationThresholdDisablerMockRecorder struct {
	mock *MockPerformanceAdvisorSlowOperationThresholdDisabler
}

// NewMockPerformanceAdvisorSlowOperationThresholdDisabler creates a new mock instance.
func NewMockPerformanceAdvisorSlowOperationThresholdDisabler(ctrl *gomock.Controller) *MockPerformanceAdvisorSlowOperationThresholdDisabler {
	mock := &MockPerformanceAdvisorSlowOperationThresholdDisabler{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorSlowOperationThresholdDisablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorSlowOperationThresholdDisabler) EXPECT() *MockPerformanceAdvisorSlowOperationThresholdDisablerMockRecorder {
	return m.recorder
}

// DisablePerformanceAdvisorSlowOperationThreshold mocks base method.
func (m *MockPerformanceAdvisorSlowOperationThresholdDisabler) DisablePerformanceAdvisorSlowOperationThreshold(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisablePerformanceAdvisorSlowOperationThreshold", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisablePerformanceAdvisorSlowOperationThreshold indicates an expected call of DisablePerformanceAdvisorSlowOperationThreshold.
func (mr *MockPerformanceAdvisorSlowOperationThresholdDisablerMockRecorder) DisablePerformanceAdvisorSlowOperationThreshold(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisablePerformanceAdvisorSlowOperationThreshold", reflect.TypeOf((*MockPerformanceAdvisorSlowOperationThresholdDisabler)(nil).DisablePerformanceAdvisorSlowOperationThreshold), arg0)
}
