// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: ProjectIPAccessListDescriber,ProjectIPAccessListLister,ProjectIPAccessListCreator,ProjectIPAccessListDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	admin "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

// MockProjectIPAccessListDescriber is a mock of ProjectIPAccessListDescriber interface.
type MockProjectIPAccessListDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockProjectIPAccessListDescriberMockRecorder
}

// MockProjectIPAccessListDescriberMockRecorder is the mock recorder for MockProjectIPAccessListDescriber.
type MockProjectIPAccessListDescriberMockRecorder struct {
	mock *MockProjectIPAccessListDescriber
}

// NewMockProjectIPAccessListDescriber creates a new mock instance.
func NewMockProjectIPAccessListDescriber(ctrl *gomock.Controller) *MockProjectIPAccessListDescriber {
	mock := &MockProjectIPAccessListDescriber{ctrl: ctrl}
	mock.recorder = &MockProjectIPAccessListDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectIPAccessListDescriber) EXPECT() *MockProjectIPAccessListDescriberMockRecorder {
	return m.recorder
}

// IPAccessList mocks base method.
func (m *MockProjectIPAccessListDescriber) IPAccessList(arg0, arg1 string) (*admin.NetworkPermissionEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IPAccessList", arg0, arg1)
	ret0, _ := ret[0].(*admin.NetworkPermissionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IPAccessList indicates an expected call of IPAccessList.
func (mr *MockProjectIPAccessListDescriberMockRecorder) IPAccessList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IPAccessList", reflect.TypeOf((*MockProjectIPAccessListDescriber)(nil).IPAccessList), arg0, arg1)
}

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
func (m *MockProjectIPAccessListLister) ProjectIPAccessLists(arg0 string, arg1 *store.ListOptions) (*admin.PaginatedNetworkAccess, error) {
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

// MockProjectIPAccessListCreator is a mock of ProjectIPAccessListCreator interface.
type MockProjectIPAccessListCreator struct {
	ctrl     *gomock.Controller
	recorder *MockProjectIPAccessListCreatorMockRecorder
}

// MockProjectIPAccessListCreatorMockRecorder is the mock recorder for MockProjectIPAccessListCreator.
type MockProjectIPAccessListCreatorMockRecorder struct {
	mock *MockProjectIPAccessListCreator
}

// NewMockProjectIPAccessListCreator creates a new mock instance.
func NewMockProjectIPAccessListCreator(ctrl *gomock.Controller) *MockProjectIPAccessListCreator {
	mock := &MockProjectIPAccessListCreator{ctrl: ctrl}
	mock.recorder = &MockProjectIPAccessListCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectIPAccessListCreator) EXPECT() *MockProjectIPAccessListCreatorMockRecorder {
	return m.recorder
}

// CreateProjectIPAccessList mocks base method.
func (m *MockProjectIPAccessListCreator) CreateProjectIPAccessList(arg0 []*admin.NetworkPermissionEntry) (*admin.PaginatedNetworkAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectIPAccessList", arg0)
	ret0, _ := ret[0].(*admin.PaginatedNetworkAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectIPAccessList indicates an expected call of CreateProjectIPAccessList.
func (mr *MockProjectIPAccessListCreatorMockRecorder) CreateProjectIPAccessList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectIPAccessList", reflect.TypeOf((*MockProjectIPAccessListCreator)(nil).CreateProjectIPAccessList), arg0)
}

// MockProjectIPAccessListDeleter is a mock of ProjectIPAccessListDeleter interface.
type MockProjectIPAccessListDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockProjectIPAccessListDeleterMockRecorder
}

// MockProjectIPAccessListDeleterMockRecorder is the mock recorder for MockProjectIPAccessListDeleter.
type MockProjectIPAccessListDeleterMockRecorder struct {
	mock *MockProjectIPAccessListDeleter
}

// NewMockProjectIPAccessListDeleter creates a new mock instance.
func NewMockProjectIPAccessListDeleter(ctrl *gomock.Controller) *MockProjectIPAccessListDeleter {
	mock := &MockProjectIPAccessListDeleter{ctrl: ctrl}
	mock.recorder = &MockProjectIPAccessListDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectIPAccessListDeleter) EXPECT() *MockProjectIPAccessListDeleterMockRecorder {
	return m.recorder
}

// DeleteProjectIPAccessList mocks base method.
func (m *MockProjectIPAccessListDeleter) DeleteProjectIPAccessList(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectIPAccessList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProjectIPAccessList indicates an expected call of DeleteProjectIPAccessList.
func (mr *MockProjectIPAccessListDeleterMockRecorder) DeleteProjectIPAccessList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectIPAccessList", reflect.TypeOf((*MockProjectIPAccessListDeleter)(nil).DeleteProjectIPAccessList), arg0, arg1)
}
