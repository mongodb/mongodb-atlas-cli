// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: UserCreator,UserDescriber,UserLister,TeamUserLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20241113005/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockUserCreator is a mock of UserCreator interface.
type MockUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockUserCreatorMockRecorder
}

// MockUserCreatorMockRecorder is the mock recorder for MockUserCreator.
type MockUserCreatorMockRecorder struct {
	mock *MockUserCreator
}

// NewMockUserCreator creates a new mock instance.
func NewMockUserCreator(ctrl *gomock.Controller) *MockUserCreator {
	mock := &MockUserCreator{ctrl: ctrl}
	mock.recorder = &MockUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCreator) EXPECT() *MockUserCreatorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserCreator) CreateUser(arg0 *admin.CloudAppUser) (*admin.CloudAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*admin.CloudAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserCreatorMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserCreator)(nil).CreateUser), arg0)
}

// MockUserDescriber is a mock of UserDescriber interface.
type MockUserDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockUserDescriberMockRecorder
}

// MockUserDescriberMockRecorder is the mock recorder for MockUserDescriber.
type MockUserDescriberMockRecorder struct {
	mock *MockUserDescriber
}

// NewMockUserDescriber creates a new mock instance.
func NewMockUserDescriber(ctrl *gomock.Controller) *MockUserDescriber {
	mock := &MockUserDescriber{ctrl: ctrl}
	mock.recorder = &MockUserDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDescriber) EXPECT() *MockUserDescriberMockRecorder {
	return m.recorder
}

// UserByID mocks base method.
func (m *MockUserDescriber) UserByID(arg0 string) (*admin.CloudAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByID", arg0)
	ret0, _ := ret[0].(*admin.CloudAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByID indicates an expected call of UserByID.
func (mr *MockUserDescriberMockRecorder) UserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByID", reflect.TypeOf((*MockUserDescriber)(nil).UserByID), arg0)
}

// UserByName mocks base method.
func (m *MockUserDescriber) UserByName(arg0 string) (*admin.CloudAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByName", arg0)
	ret0, _ := ret[0].(*admin.CloudAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByName indicates an expected call of UserByName.
func (mr *MockUserDescriberMockRecorder) UserByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByName", reflect.TypeOf((*MockUserDescriber)(nil).UserByName), arg0)
}

// MockUserLister is a mock of UserLister interface.
type MockUserLister struct {
	ctrl     *gomock.Controller
	recorder *MockUserListerMockRecorder
}

// MockUserListerMockRecorder is the mock recorder for MockUserLister.
type MockUserListerMockRecorder struct {
	mock *MockUserLister
}

// NewMockUserLister creates a new mock instance.
func NewMockUserLister(ctrl *gomock.Controller) *MockUserLister {
	mock := &MockUserLister{ctrl: ctrl}
	mock.recorder = &MockUserListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserLister) EXPECT() *MockUserListerMockRecorder {
	return m.recorder
}

// OrganizationUsers mocks base method.
func (m *MockUserLister) OrganizationUsers(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationUsers indicates an expected call of OrganizationUsers.
func (mr *MockUserListerMockRecorder) OrganizationUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationUsers", reflect.TypeOf((*MockUserLister)(nil).OrganizationUsers), arg0, arg1)
}

// MockTeamUserLister is a mock of TeamUserLister interface.
type MockTeamUserLister struct {
	ctrl     *gomock.Controller
	recorder *MockTeamUserListerMockRecorder
}

// MockTeamUserListerMockRecorder is the mock recorder for MockTeamUserLister.
type MockTeamUserListerMockRecorder struct {
	mock *MockTeamUserLister
}

// NewMockTeamUserLister creates a new mock instance.
func NewMockTeamUserLister(ctrl *gomock.Controller) *MockTeamUserLister {
	mock := &MockTeamUserLister{ctrl: ctrl}
	mock.recorder = &MockTeamUserListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamUserLister) EXPECT() *MockTeamUserListerMockRecorder {
	return m.recorder
}

// TeamUsers mocks base method.
func (m *MockTeamUserLister) TeamUsers(arg0, arg1 string) (*admin.PaginatedAppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedAppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamUsers indicates an expected call of TeamUsers.
func (mr *MockTeamUserListerMockRecorder) TeamUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUsers", reflect.TypeOf((*MockTeamUserLister)(nil).TeamUsers), arg0, arg1)
}
