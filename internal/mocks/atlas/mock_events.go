// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: OrganizationEventLister,ProjectEventLister,EventLister)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// MockOrganizationEventLister is a mock of OrganizationEventLister interface.
type MockOrganizationEventLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationEventListerMockRecorder
}

// MockOrganizationEventListerMockRecorder is the mock recorder for MockOrganizationEventLister.
type MockOrganizationEventListerMockRecorder struct {
	mock *MockOrganizationEventLister
}

// NewMockOrganizationEventLister creates a new mock instance.
func NewMockOrganizationEventLister(ctrl *gomock.Controller) *MockOrganizationEventLister {
	mock := &MockOrganizationEventLister{ctrl: ctrl}
	mock.recorder = &MockOrganizationEventListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationEventLister) EXPECT() *MockOrganizationEventListerMockRecorder {
	return m.recorder
}

// OrganizationEvents mocks base method.
func (m *MockOrganizationEventLister) OrganizationEvents(arg0 *mongodbatlasv2.ListOrganizationEventsApiParams) (*mongodbatlasv2.OrgPaginatedEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationEvents", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.OrgPaginatedEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationEvents indicates an expected call of OrganizationEvents.
func (mr *MockOrganizationEventListerMockRecorder) OrganizationEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationEvents", reflect.TypeOf((*MockOrganizationEventLister)(nil).OrganizationEvents), arg0)
}

// MockProjectEventLister is a mock of ProjectEventLister interface.
type MockProjectEventLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectEventListerMockRecorder
}

// MockProjectEventListerMockRecorder is the mock recorder for MockProjectEventLister.
type MockProjectEventListerMockRecorder struct {
	mock *MockProjectEventLister
}

// NewMockProjectEventLister creates a new mock instance.
func NewMockProjectEventLister(ctrl *gomock.Controller) *MockProjectEventLister {
	mock := &MockProjectEventLister{ctrl: ctrl}
	mock.recorder = &MockProjectEventListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectEventLister) EXPECT() *MockProjectEventListerMockRecorder {
	return m.recorder
}

// ProjectEvents mocks base method.
func (m *MockProjectEventLister) ProjectEvents(arg0 *mongodbatlasv2.ListProjectEventsApiParams) (*mongodbatlasv2.GroupPaginatedEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectEvents", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.GroupPaginatedEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectEvents indicates an expected call of ProjectEvents.
func (mr *MockProjectEventListerMockRecorder) ProjectEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectEvents", reflect.TypeOf((*MockProjectEventLister)(nil).ProjectEvents), arg0)
}

// MockEventLister is a mock of EventLister interface.
type MockEventLister struct {
	ctrl     *gomock.Controller
	recorder *MockEventListerMockRecorder
}

// MockEventListerMockRecorder is the mock recorder for MockEventLister.
type MockEventListerMockRecorder struct {
	mock *MockEventLister
}

// NewMockEventLister creates a new mock instance.
func NewMockEventLister(ctrl *gomock.Controller) *MockEventLister {
	mock := &MockEventLister{ctrl: ctrl}
	mock.recorder = &MockEventListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventLister) EXPECT() *MockEventListerMockRecorder {
	return m.recorder
}

// OrganizationEvents mocks base method.
func (m *MockEventLister) OrganizationEvents(arg0 *mongodbatlasv2.ListOrganizationEventsApiParams) (*mongodbatlasv2.OrgPaginatedEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationEvents", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.OrgPaginatedEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationEvents indicates an expected call of OrganizationEvents.
func (mr *MockEventListerMockRecorder) OrganizationEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationEvents", reflect.TypeOf((*MockEventLister)(nil).OrganizationEvents), arg0)
}

// ProjectEvents mocks base method.
func (m *MockEventLister) ProjectEvents(arg0 *mongodbatlasv2.ListProjectEventsApiParams) (*mongodbatlasv2.GroupPaginatedEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectEvents", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.GroupPaginatedEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectEvents indicates an expected call of ProjectEvents.
func (mr *MockEventListerMockRecorder) ProjectEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectEvents", reflect.TypeOf((*MockEventLister)(nil).ProjectEvents), arg0)
}
