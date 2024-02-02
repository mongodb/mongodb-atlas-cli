// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: ProjectIPAccessListLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	atlas "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	admin "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

// MockProjectIPAccessListLister is a mock of ProjectIPAccessListLister interface.
type MockProjectIPAccessListLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectIPAccessListListerMockRecorder
}

// MockProjectIPAccessListListerMockRecorder is the mock recorder for MockProjectIPAccessListLister.
type MockProjectIPAccessListListerMockRecorder struct {
	mock *MockProjectIPAccessListLister
}

// NewMockProjectIPAccessListLister creates a new mock instance.
func NewMockProjectIPAccessListLister(ctrl *gomock.Controller) *MockProjectIPAccessListLister {
	mock := &MockProjectIPAccessListLister{ctrl: ctrl}
	mock.recorder = &MockProjectIPAccessListListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectIPAccessListLister) EXPECT() *MockProjectIPAccessListListerMockRecorder {
	return m.recorder
}

// ProjectIPAccessLists mocks base method.
func (m *MockProjectIPAccessListLister) ProjectIPAccessLists(arg0 string, arg1 *atlas.ListOptions) (*admin.PaginatedNetworkAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectIPAccessLists", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedNetworkAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectIPAccessLists indicates an expected call of ProjectIPAccessLists.
func (mr *MockProjectIPAccessListListerMockRecorder) ProjectIPAccessLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectIPAccessLists", reflect.TypeOf((*MockProjectIPAccessListLister)(nil).ProjectIPAccessLists), arg0, arg1)
}
