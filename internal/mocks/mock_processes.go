// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: ProcessLister,ProcessDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

// MockProcessLister is a mock of ProcessLister interface.
type MockProcessLister struct {
	ctrl     *gomock.Controller
	recorder *MockProcessListerMockRecorder
}

// MockProcessListerMockRecorder is the mock recorder for MockProcessLister.
type MockProcessListerMockRecorder struct {
	mock *MockProcessLister
}

// NewMockProcessLister creates a new mock instance.
func NewMockProcessLister(ctrl *gomock.Controller) *MockProcessLister {
	mock := &MockProcessLister{ctrl: ctrl}
	mock.recorder = &MockProcessListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessLister) EXPECT() *MockProcessListerMockRecorder {
	return m.recorder
}

// Processes mocks base method.
func (m *MockProcessLister) Processes(arg0 *admin.ListAtlasProcessesApiParams) (*admin.PaginatedHostViewAtlas, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Processes", arg0)
	ret0, _ := ret[0].(*admin.PaginatedHostViewAtlas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Processes indicates an expected call of Processes.
func (mr *MockProcessListerMockRecorder) Processes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Processes", reflect.TypeOf((*MockProcessLister)(nil).Processes), arg0)
}

// MockProcessDescriber is a mock of ProcessDescriber interface.
type MockProcessDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockProcessDescriberMockRecorder
}

// MockProcessDescriberMockRecorder is the mock recorder for MockProcessDescriber.
type MockProcessDescriberMockRecorder struct {
	mock *MockProcessDescriber
}

// NewMockProcessDescriber creates a new mock instance.
func NewMockProcessDescriber(ctrl *gomock.Controller) *MockProcessDescriber {
	mock := &MockProcessDescriber{ctrl: ctrl}
	mock.recorder = &MockProcessDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessDescriber) EXPECT() *MockProcessDescriberMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockProcessDescriber) Process(arg0 *admin.GetAtlasProcessApiParams) (*admin.ApiHostViewAtlas, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0)
	ret0, _ := ret[0].(*admin.ApiHostViewAtlas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Process indicates an expected call of Process.
func (mr *MockProcessDescriberMockRecorder) Process(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockProcessDescriber)(nil).Process), arg0)
}
