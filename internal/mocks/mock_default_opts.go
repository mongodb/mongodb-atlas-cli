// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli (interfaces: ProjectOrgsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	admin "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

// MockProjectOrgsLister is a mock of ProjectOrgsLister interface.
type MockProjectOrgsLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectOrgsListerMockRecorder
}

// MockProjectOrgsListerMockRecorder is the mock recorder for MockProjectOrgsLister.
type MockProjectOrgsListerMockRecorder struct {
	mock *MockProjectOrgsLister
}

// NewMockProjectOrgsLister creates a new mock instance.
func NewMockProjectOrgsLister(ctrl *gomock.Controller) *MockProjectOrgsLister {
	mock := &MockProjectOrgsLister{ctrl: ctrl}
	mock.recorder = &MockProjectOrgsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectOrgsLister) EXPECT() *MockProjectOrgsListerMockRecorder {
	return m.recorder
}

// GetOrgProjects mocks base method.
func (m *MockProjectOrgsLister) GetOrgProjects(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgProjects", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgProjects indicates an expected call of GetOrgProjects.
func (mr *MockProjectOrgsListerMockRecorder) GetOrgProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgProjects", reflect.TypeOf((*MockProjectOrgsLister)(nil).GetOrgProjects), arg0, arg1)
}

// Organization mocks base method.
func (m *MockProjectOrgsLister) Organization(arg0 string) (*admin.AtlasOrganization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Organization", arg0)
	ret0, _ := ret[0].(*admin.AtlasOrganization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Organization indicates an expected call of Organization.
func (mr *MockProjectOrgsListerMockRecorder) Organization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Organization", reflect.TypeOf((*MockProjectOrgsLister)(nil).Organization), arg0)
}

// Organizations mocks base method.
func (m *MockProjectOrgsLister) Organizations(arg0 *admin.ListOrganizationsApiParams) (*admin.PaginatedOrganization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Organizations", arg0)
	ret0, _ := ret[0].(*admin.PaginatedOrganization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Organizations indicates an expected call of Organizations.
func (mr *MockProjectOrgsListerMockRecorder) Organizations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Organizations", reflect.TypeOf((*MockProjectOrgsLister)(nil).Organizations), arg0)
}

// Project mocks base method.
func (m *MockProjectOrgsLister) Project(arg0 string) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockProjectOrgsListerMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockProjectOrgsLister)(nil).Project), arg0)
}

// Projects mocks base method.
func (m *MockProjectOrgsLister) Projects(arg0 *store.ListOptions) (*admin.PaginatedAtlasGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects", arg0)
	ret0, _ := ret[0].(*admin.PaginatedAtlasGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockProjectOrgsListerMockRecorder) Projects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockProjectOrgsLister)(nil).Projects), arg0)
}
