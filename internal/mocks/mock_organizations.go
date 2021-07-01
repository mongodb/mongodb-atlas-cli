// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: OrganizationLister,OrganizationCreator,OrganizationDeleter,OrganizationDescriber,OrganizationInvitationLister,OrganizationInvitationDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockOrganizationLister is a mock of OrganizationLister interface
type MockOrganizationLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationListerMockRecorder
}

// MockOrganizationListerMockRecorder is the mock recorder for MockOrganizationLister
type MockOrganizationListerMockRecorder struct {
	mock *MockOrganizationLister
}

// NewMockOrganizationLister creates a new mock instance
func NewMockOrganizationLister(ctrl *gomock.Controller) *MockOrganizationLister {
	mock := &MockOrganizationLister{ctrl: ctrl}
	mock.recorder = &MockOrganizationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationLister) EXPECT() *MockOrganizationListerMockRecorder {
	return m.recorder
}

// Organizations mocks base method
func (m *MockOrganizationLister) Organizations(arg0 *mongodbatlas.OrganizationsListOptions) (*mongodbatlas.Organizations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Organizations", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Organizations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Organizations indicates an expected call of Organizations
func (mr *MockOrganizationListerMockRecorder) Organizations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Organizations", reflect.TypeOf((*MockOrganizationLister)(nil).Organizations), arg0)
}

// MockOrganizationCreator is a mock of OrganizationCreator interface
type MockOrganizationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationCreatorMockRecorder
}

// MockOrganizationCreatorMockRecorder is the mock recorder for MockOrganizationCreator
type MockOrganizationCreatorMockRecorder struct {
	mock *MockOrganizationCreator
}

// NewMockOrganizationCreator creates a new mock instance
func NewMockOrganizationCreator(ctrl *gomock.Controller) *MockOrganizationCreator {
	mock := &MockOrganizationCreator{ctrl: ctrl}
	mock.recorder = &MockOrganizationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationCreator) EXPECT() *MockOrganizationCreatorMockRecorder {
	return m.recorder
}

// CreateOrganization mocks base method
func (m *MockOrganizationCreator) CreateOrganization(arg0 string) (*mongodbatlas.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization
func (mr *MockOrganizationCreatorMockRecorder) CreateOrganization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockOrganizationCreator)(nil).CreateOrganization), arg0)
}

// MockOrganizationDeleter is a mock of OrganizationDeleter interface
type MockOrganizationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationDeleterMockRecorder
}

// MockOrganizationDeleterMockRecorder is the mock recorder for MockOrganizationDeleter
type MockOrganizationDeleterMockRecorder struct {
	mock *MockOrganizationDeleter
}

// NewMockOrganizationDeleter creates a new mock instance
func NewMockOrganizationDeleter(ctrl *gomock.Controller) *MockOrganizationDeleter {
	mock := &MockOrganizationDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationDeleter) EXPECT() *MockOrganizationDeleterMockRecorder {
	return m.recorder
}

// DeleteOrganization mocks base method
func (m *MockOrganizationDeleter) DeleteOrganization(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganization", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganization indicates an expected call of DeleteOrganization
func (mr *MockOrganizationDeleterMockRecorder) DeleteOrganization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganization", reflect.TypeOf((*MockOrganizationDeleter)(nil).DeleteOrganization), arg0)
}

// MockOrganizationDescriber is a mock of OrganizationDescriber interface
type MockOrganizationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationDescriberMockRecorder
}

// MockOrganizationDescriberMockRecorder is the mock recorder for MockOrganizationDescriber
type MockOrganizationDescriberMockRecorder struct {
	mock *MockOrganizationDescriber
}

// NewMockOrganizationDescriber creates a new mock instance
func NewMockOrganizationDescriber(ctrl *gomock.Controller) *MockOrganizationDescriber {
	mock := &MockOrganizationDescriber{ctrl: ctrl}
	mock.recorder = &MockOrganizationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationDescriber) EXPECT() *MockOrganizationDescriberMockRecorder {
	return m.recorder
}

// Organization mocks base method
func (m *MockOrganizationDescriber) Organization(arg0 string) (*mongodbatlas.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Organization", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Organization indicates an expected call of Organization
func (mr *MockOrganizationDescriberMockRecorder) Organization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Organization", reflect.TypeOf((*MockOrganizationDescriber)(nil).Organization), arg0)
}

// MockOrganizationInvitationLister is a mock of OrganizationInvitationLister interface
type MockOrganizationInvitationLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInvitationListerMockRecorder
}

// MockOrganizationInvitationListerMockRecorder is the mock recorder for MockOrganizationInvitationLister
type MockOrganizationInvitationListerMockRecorder struct {
	mock *MockOrganizationInvitationLister
}

// NewMockOrganizationInvitationLister creates a new mock instance
func NewMockOrganizationInvitationLister(ctrl *gomock.Controller) *MockOrganizationInvitationLister {
	mock := &MockOrganizationInvitationLister{ctrl: ctrl}
	mock.recorder = &MockOrganizationInvitationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationInvitationLister) EXPECT() *MockOrganizationInvitationListerMockRecorder {
	return m.recorder
}

// OrganizationInvitations mocks base method
func (m *MockOrganizationInvitationLister) OrganizationInvitations(arg0 string, arg1 *mongodbatlas.InvitationOptions) ([]*mongodbatlas.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationInvitations", arg0, arg1)
	ret0, _ := ret[0].([]*mongodbatlas.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationInvitations indicates an expected call of OrganizationInvitations
func (mr *MockOrganizationInvitationListerMockRecorder) OrganizationInvitations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationInvitations", reflect.TypeOf((*MockOrganizationInvitationLister)(nil).OrganizationInvitations), arg0, arg1)
}

// MockOrganizationInvitationDescriber is a mock of OrganizationInvitationDescriber interface
type MockOrganizationInvitationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInvitationDescriberMockRecorder
}

// MockOrganizationInvitationDescriberMockRecorder is the mock recorder for MockOrganizationInvitationDescriber
type MockOrganizationInvitationDescriberMockRecorder struct {
	mock *MockOrganizationInvitationDescriber
}

// NewMockOrganizationInvitationDescriber creates a new mock instance
func NewMockOrganizationInvitationDescriber(ctrl *gomock.Controller) *MockOrganizationInvitationDescriber {
	mock := &MockOrganizationInvitationDescriber{ctrl: ctrl}
	mock.recorder = &MockOrganizationInvitationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationInvitationDescriber) EXPECT() *MockOrganizationInvitationDescriberMockRecorder {
	return m.recorder
}

// OrganizationInvitation mocks base method
func (m *MockOrganizationInvitationDescriber) OrganizationInvitation(arg0, arg1 string) (*mongodbatlas.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationInvitation", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationInvitation indicates an expected call of OrganizationInvitation
func (mr *MockOrganizationInvitationDescriberMockRecorder) OrganizationInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationInvitation", reflect.TypeOf((*MockOrganizationInvitationDescriber)(nil).OrganizationInvitation), arg0, arg1)
}
