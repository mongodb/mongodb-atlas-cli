// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: ProjectInvitationLister,ProjectInvitationDescriber,ProjectInviter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockProjectInvitationLister is a mock of ProjectInvitationLister interface
type MockProjectInvitationLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectInvitationListerMockRecorder
}

// MockProjectInvitationListerMockRecorder is the mock recorder for MockProjectInvitationLister
type MockProjectInvitationListerMockRecorder struct {
	mock *MockProjectInvitationLister
}

// NewMockProjectInvitationLister creates a new mock instance
func NewMockProjectInvitationLister(ctrl *gomock.Controller) *MockProjectInvitationLister {
	mock := &MockProjectInvitationLister{ctrl: ctrl}
	mock.recorder = &MockProjectInvitationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectInvitationLister) EXPECT() *MockProjectInvitationListerMockRecorder {
	return m.recorder
}

// ProjectInvitations mocks base method
func (m *MockProjectInvitationLister) ProjectInvitations(arg0 string, arg1 *mongodbatlas.InvitationOptions) ([]*mongodbatlas.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectInvitations", arg0, arg1)
	ret0, _ := ret[0].([]*mongodbatlas.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectInvitations indicates an expected call of ProjectInvitations
func (mr *MockProjectInvitationListerMockRecorder) ProjectInvitations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectInvitations", reflect.TypeOf((*MockProjectInvitationLister)(nil).ProjectInvitations), arg0, arg1)
}

// MockProjectInvitationDescriber is a mock of ProjectInvitationDescriber interface
type MockProjectInvitationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockProjectInvitationDescriberMockRecorder
}

// MockProjectInvitationDescriberMockRecorder is the mock recorder for MockProjectInvitationDescriber
type MockProjectInvitationDescriberMockRecorder struct {
	mock *MockProjectInvitationDescriber
}

// NewMockProjectInvitationDescriber creates a new mock instance
func NewMockProjectInvitationDescriber(ctrl *gomock.Controller) *MockProjectInvitationDescriber {
	mock := &MockProjectInvitationDescriber{ctrl: ctrl}
	mock.recorder = &MockProjectInvitationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectInvitationDescriber) EXPECT() *MockProjectInvitationDescriberMockRecorder {
	return m.recorder
}

// ProjectInvitation mocks base method
func (m *MockProjectInvitationDescriber) ProjectInvitation(arg0, arg1 string) (*mongodbatlas.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectInvitation", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectInvitation indicates an expected call of ProjectInvitation
func (mr *MockProjectInvitationDescriberMockRecorder) ProjectInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectInvitation", reflect.TypeOf((*MockProjectInvitationDescriber)(nil).ProjectInvitation), arg0, arg1)
}

// MockProjectInviter is a mock of ProjectInviter interface
type MockProjectInviter struct {
	ctrl     *gomock.Controller
	recorder *MockProjectInviterMockRecorder
}

// MockProjectInviterMockRecorder is the mock recorder for MockProjectInviter
type MockProjectInviterMockRecorder struct {
	mock *MockProjectInviter
}

// NewMockProjectInviter creates a new mock instance
func NewMockProjectInviter(ctrl *gomock.Controller) *MockProjectInviter {
	mock := &MockProjectInviter{ctrl: ctrl}
	mock.recorder = &MockProjectInviterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectInviter) EXPECT() *MockProjectInviterMockRecorder {
	return m.recorder
}

// InviteUserToProject mocks base method
func (m *MockProjectInviter) InviteUserToProject(arg0 string, arg1 *mongodbatlas.Invitation) (*mongodbatlas.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteUserToProject", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteUserToProject indicates an expected call of InviteUserToProject
func (mr *MockProjectInviterMockRecorder) InviteUserToProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUserToProject", reflect.TypeOf((*MockProjectInviter)(nil).InviteUserToProject), arg0, arg1)
}
