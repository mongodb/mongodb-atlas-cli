// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: AccessLogsListerByClusterName,AccessLogsListerByHostname,AccessLogsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20241113004/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockAccessLogsListerByClusterName is a mock of AccessLogsListerByClusterName interface.
type MockAccessLogsListerByClusterName struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogsListerByClusterNameMockRecorder
}

// MockAccessLogsListerByClusterNameMockRecorder is the mock recorder for MockAccessLogsListerByClusterName.
type MockAccessLogsListerByClusterNameMockRecorder struct {
	mock *MockAccessLogsListerByClusterName
}

// NewMockAccessLogsListerByClusterName creates a new mock instance.
func NewMockAccessLogsListerByClusterName(ctrl *gomock.Controller) *MockAccessLogsListerByClusterName {
	mock := &MockAccessLogsListerByClusterName{ctrl: ctrl}
	mock.recorder = &MockAccessLogsListerByClusterNameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogsListerByClusterName) EXPECT() *MockAccessLogsListerByClusterNameMockRecorder {
	return m.recorder
}

// AccessLogsByClusterName mocks base method.
func (m *MockAccessLogsListerByClusterName) AccessLogsByClusterName(arg0, arg1 string, arg2 *mongodbatlas.AccessLogOptions) (*admin.MongoDBAccessLogsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogsByClusterName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.MongoDBAccessLogsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessLogsByClusterName indicates an expected call of AccessLogsByClusterName.
func (mr *MockAccessLogsListerByClusterNameMockRecorder) AccessLogsByClusterName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogsByClusterName", reflect.TypeOf((*MockAccessLogsListerByClusterName)(nil).AccessLogsByClusterName), arg0, arg1, arg2)
}

// MockAccessLogsListerByHostname is a mock of AccessLogsListerByHostname interface.
type MockAccessLogsListerByHostname struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogsListerByHostnameMockRecorder
}

// MockAccessLogsListerByHostnameMockRecorder is the mock recorder for MockAccessLogsListerByHostname.
type MockAccessLogsListerByHostnameMockRecorder struct {
	mock *MockAccessLogsListerByHostname
}

// NewMockAccessLogsListerByHostname creates a new mock instance.
func NewMockAccessLogsListerByHostname(ctrl *gomock.Controller) *MockAccessLogsListerByHostname {
	mock := &MockAccessLogsListerByHostname{ctrl: ctrl}
	mock.recorder = &MockAccessLogsListerByHostnameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogsListerByHostname) EXPECT() *MockAccessLogsListerByHostnameMockRecorder {
	return m.recorder
}

// AccessLogsByHostname mocks base method.
func (m *MockAccessLogsListerByHostname) AccessLogsByHostname(arg0, arg1 string, arg2 *mongodbatlas.AccessLogOptions) (*admin.MongoDBAccessLogsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogsByHostname", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.MongoDBAccessLogsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessLogsByHostname indicates an expected call of AccessLogsByHostname.
func (mr *MockAccessLogsListerByHostnameMockRecorder) AccessLogsByHostname(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogsByHostname", reflect.TypeOf((*MockAccessLogsListerByHostname)(nil).AccessLogsByHostname), arg0, arg1, arg2)
}

// MockAccessLogsLister is a mock of AccessLogsLister interface.
type MockAccessLogsLister struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogsListerMockRecorder
}

// MockAccessLogsListerMockRecorder is the mock recorder for MockAccessLogsLister.
type MockAccessLogsListerMockRecorder struct {
	mock *MockAccessLogsLister
}

// NewMockAccessLogsLister creates a new mock instance.
func NewMockAccessLogsLister(ctrl *gomock.Controller) *MockAccessLogsLister {
	mock := &MockAccessLogsLister{ctrl: ctrl}
	mock.recorder = &MockAccessLogsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogsLister) EXPECT() *MockAccessLogsListerMockRecorder {
	return m.recorder
}

// AccessLogsByClusterName mocks base method.
func (m *MockAccessLogsLister) AccessLogsByClusterName(arg0, arg1 string, arg2 *mongodbatlas.AccessLogOptions) (*admin.MongoDBAccessLogsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogsByClusterName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.MongoDBAccessLogsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessLogsByClusterName indicates an expected call of AccessLogsByClusterName.
func (mr *MockAccessLogsListerMockRecorder) AccessLogsByClusterName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogsByClusterName", reflect.TypeOf((*MockAccessLogsLister)(nil).AccessLogsByClusterName), arg0, arg1, arg2)
}

// AccessLogsByHostname mocks base method.
func (m *MockAccessLogsLister) AccessLogsByHostname(arg0, arg1 string, arg2 *mongodbatlas.AccessLogOptions) (*admin.MongoDBAccessLogsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogsByHostname", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.MongoDBAccessLogsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessLogsByHostname indicates an expected call of AccessLogsByHostname.
func (mr *MockAccessLogsListerMockRecorder) AccessLogsByHostname(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogsByHostname", reflect.TypeOf((*MockAccessLogsLister)(nil).AccessLogsByHostname), arg0, arg1, arg2)
}
