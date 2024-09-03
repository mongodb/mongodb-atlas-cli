// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: ProjectAPIKeyLister,ProjectAPIKeyCreator,OrganizationAPIKeyLister,OrganizationAPIKeyDescriber,OrganizationAPIKeyUpdater,OrganizationAPIKeyCreator,OrganizationAPIKeyDeleter,ProjectAPIKeyDeleter,ProjectAPIKeyAssigner)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20240805003/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockProjectAPIKeyLister is a mock of ProjectAPIKeyLister interface.
type MockProjectAPIKeyLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectAPIKeyListerMockRecorder
}

// MockProjectAPIKeyListerMockRecorder is the mock recorder for MockProjectAPIKeyLister.
type MockProjectAPIKeyListerMockRecorder struct {
	mock *MockProjectAPIKeyLister
}

// NewMockProjectAPIKeyLister creates a new mock instance.
func NewMockProjectAPIKeyLister(ctrl *gomock.Controller) *MockProjectAPIKeyLister {
	mock := &MockProjectAPIKeyLister{ctrl: ctrl}
	mock.recorder = &MockProjectAPIKeyListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectAPIKeyLister) EXPECT() *MockProjectAPIKeyListerMockRecorder {
	return m.recorder
}

// ProjectAPIKeys mocks base method.
func (m *MockProjectAPIKeyLister) ProjectAPIKeys(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedApiApiUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectAPIKeys", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiApiUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectAPIKeys indicates an expected call of ProjectAPIKeys.
func (mr *MockProjectAPIKeyListerMockRecorder) ProjectAPIKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectAPIKeys", reflect.TypeOf((*MockProjectAPIKeyLister)(nil).ProjectAPIKeys), arg0, arg1)
}

// MockProjectAPIKeyCreator is a mock of ProjectAPIKeyCreator interface.
type MockProjectAPIKeyCreator struct {
	ctrl     *gomock.Controller
	recorder *MockProjectAPIKeyCreatorMockRecorder
}

// MockProjectAPIKeyCreatorMockRecorder is the mock recorder for MockProjectAPIKeyCreator.
type MockProjectAPIKeyCreatorMockRecorder struct {
	mock *MockProjectAPIKeyCreator
}

// NewMockProjectAPIKeyCreator creates a new mock instance.
func NewMockProjectAPIKeyCreator(ctrl *gomock.Controller) *MockProjectAPIKeyCreator {
	mock := &MockProjectAPIKeyCreator{ctrl: ctrl}
	mock.recorder = &MockProjectAPIKeyCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectAPIKeyCreator) EXPECT() *MockProjectAPIKeyCreatorMockRecorder {
	return m.recorder
}

// CreateProjectAPIKey mocks base method.
func (m *MockProjectAPIKeyCreator) CreateProjectAPIKey(arg0 string, arg1 *admin.CreateAtlasProjectApiKey) (*admin.ApiKeyUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiKeyUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectAPIKey indicates an expected call of CreateProjectAPIKey.
func (mr *MockProjectAPIKeyCreatorMockRecorder) CreateProjectAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectAPIKey", reflect.TypeOf((*MockProjectAPIKeyCreator)(nil).CreateProjectAPIKey), arg0, arg1)
}

// MockOrganizationAPIKeyLister is a mock of OrganizationAPIKeyLister interface.
type MockOrganizationAPIKeyLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyListerMockRecorder
}

// MockOrganizationAPIKeyListerMockRecorder is the mock recorder for MockOrganizationAPIKeyLister.
type MockOrganizationAPIKeyListerMockRecorder struct {
	mock *MockOrganizationAPIKeyLister
}

// NewMockOrganizationAPIKeyLister creates a new mock instance.
func NewMockOrganizationAPIKeyLister(ctrl *gomock.Controller) *MockOrganizationAPIKeyLister {
	mock := &MockOrganizationAPIKeyLister{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyLister) EXPECT() *MockOrganizationAPIKeyListerMockRecorder {
	return m.recorder
}

// OrganizationAPIKeys mocks base method.
func (m *MockOrganizationAPIKeyLister) OrganizationAPIKeys(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedApiApiUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationAPIKeys", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiApiUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationAPIKeys indicates an expected call of OrganizationAPIKeys.
func (mr *MockOrganizationAPIKeyListerMockRecorder) OrganizationAPIKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationAPIKeys", reflect.TypeOf((*MockOrganizationAPIKeyLister)(nil).OrganizationAPIKeys), arg0, arg1)
}

// MockOrganizationAPIKeyDescriber is a mock of OrganizationAPIKeyDescriber interface.
type MockOrganizationAPIKeyDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyDescriberMockRecorder
}

// MockOrganizationAPIKeyDescriberMockRecorder is the mock recorder for MockOrganizationAPIKeyDescriber.
type MockOrganizationAPIKeyDescriberMockRecorder struct {
	mock *MockOrganizationAPIKeyDescriber
}

// NewMockOrganizationAPIKeyDescriber creates a new mock instance.
func NewMockOrganizationAPIKeyDescriber(ctrl *gomock.Controller) *MockOrganizationAPIKeyDescriber {
	mock := &MockOrganizationAPIKeyDescriber{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyDescriber) EXPECT() *MockOrganizationAPIKeyDescriberMockRecorder {
	return m.recorder
}

// OrganizationAPIKey mocks base method.
func (m *MockOrganizationAPIKeyDescriber) OrganizationAPIKey(arg0, arg1 string) (*admin.ApiKeyUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiKeyUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationAPIKey indicates an expected call of OrganizationAPIKey.
func (mr *MockOrganizationAPIKeyDescriberMockRecorder) OrganizationAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationAPIKey", reflect.TypeOf((*MockOrganizationAPIKeyDescriber)(nil).OrganizationAPIKey), arg0, arg1)
}

// MockOrganizationAPIKeyUpdater is a mock of OrganizationAPIKeyUpdater interface.
type MockOrganizationAPIKeyUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyUpdaterMockRecorder
}

// MockOrganizationAPIKeyUpdaterMockRecorder is the mock recorder for MockOrganizationAPIKeyUpdater.
type MockOrganizationAPIKeyUpdaterMockRecorder struct {
	mock *MockOrganizationAPIKeyUpdater
}

// NewMockOrganizationAPIKeyUpdater creates a new mock instance.
func NewMockOrganizationAPIKeyUpdater(ctrl *gomock.Controller) *MockOrganizationAPIKeyUpdater {
	mock := &MockOrganizationAPIKeyUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyUpdater) EXPECT() *MockOrganizationAPIKeyUpdaterMockRecorder {
	return m.recorder
}

// UpdateOrganizationAPIKey mocks base method.
func (m *MockOrganizationAPIKeyUpdater) UpdateOrganizationAPIKey(arg0, arg1 string, arg2 *admin.UpdateAtlasOrganizationApiKey) (*admin.ApiKeyUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrganizationAPIKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ApiKeyUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganizationAPIKey indicates an expected call of UpdateOrganizationAPIKey.
func (mr *MockOrganizationAPIKeyUpdaterMockRecorder) UpdateOrganizationAPIKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganizationAPIKey", reflect.TypeOf((*MockOrganizationAPIKeyUpdater)(nil).UpdateOrganizationAPIKey), arg0, arg1, arg2)
}

// MockOrganizationAPIKeyCreator is a mock of OrganizationAPIKeyCreator interface.
type MockOrganizationAPIKeyCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyCreatorMockRecorder
}

// MockOrganizationAPIKeyCreatorMockRecorder is the mock recorder for MockOrganizationAPIKeyCreator.
type MockOrganizationAPIKeyCreatorMockRecorder struct {
	mock *MockOrganizationAPIKeyCreator
}

// NewMockOrganizationAPIKeyCreator creates a new mock instance.
func NewMockOrganizationAPIKeyCreator(ctrl *gomock.Controller) *MockOrganizationAPIKeyCreator {
	mock := &MockOrganizationAPIKeyCreator{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyCreator) EXPECT() *MockOrganizationAPIKeyCreatorMockRecorder {
	return m.recorder
}

// CreateOrganizationAPIKey mocks base method.
func (m *MockOrganizationAPIKeyCreator) CreateOrganizationAPIKey(arg0 string, arg1 *admin.CreateAtlasOrganizationApiKey) (*admin.ApiKeyUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganizationAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiKeyUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganizationAPIKey indicates an expected call of CreateOrganizationAPIKey.
func (mr *MockOrganizationAPIKeyCreatorMockRecorder) CreateOrganizationAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizationAPIKey", reflect.TypeOf((*MockOrganizationAPIKeyCreator)(nil).CreateOrganizationAPIKey), arg0, arg1)
}

// MockOrganizationAPIKeyDeleter is a mock of OrganizationAPIKeyDeleter interface.
type MockOrganizationAPIKeyDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyDeleterMockRecorder
}

// MockOrganizationAPIKeyDeleterMockRecorder is the mock recorder for MockOrganizationAPIKeyDeleter.
type MockOrganizationAPIKeyDeleterMockRecorder struct {
	mock *MockOrganizationAPIKeyDeleter
}

// NewMockOrganizationAPIKeyDeleter creates a new mock instance.
func NewMockOrganizationAPIKeyDeleter(ctrl *gomock.Controller) *MockOrganizationAPIKeyDeleter {
	mock := &MockOrganizationAPIKeyDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyDeleter) EXPECT() *MockOrganizationAPIKeyDeleterMockRecorder {
	return m.recorder
}

// DeleteOrganizationAPIKey mocks base method.
func (m *MockOrganizationAPIKeyDeleter) DeleteOrganizationAPIKey(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganizationAPIKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganizationAPIKey indicates an expected call of DeleteOrganizationAPIKey.
func (mr *MockOrganizationAPIKeyDeleterMockRecorder) DeleteOrganizationAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganizationAPIKey", reflect.TypeOf((*MockOrganizationAPIKeyDeleter)(nil).DeleteOrganizationAPIKey), arg0, arg1)
}

// MockProjectAPIKeyDeleter is a mock of ProjectAPIKeyDeleter interface.
type MockProjectAPIKeyDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockProjectAPIKeyDeleterMockRecorder
}

// MockProjectAPIKeyDeleterMockRecorder is the mock recorder for MockProjectAPIKeyDeleter.
type MockProjectAPIKeyDeleterMockRecorder struct {
	mock *MockProjectAPIKeyDeleter
}

// NewMockProjectAPIKeyDeleter creates a new mock instance.
func NewMockProjectAPIKeyDeleter(ctrl *gomock.Controller) *MockProjectAPIKeyDeleter {
	mock := &MockProjectAPIKeyDeleter{ctrl: ctrl}
	mock.recorder = &MockProjectAPIKeyDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectAPIKeyDeleter) EXPECT() *MockProjectAPIKeyDeleterMockRecorder {
	return m.recorder
}

// DeleteProjectAPIKey mocks base method.
func (m *MockProjectAPIKeyDeleter) DeleteProjectAPIKey(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectAPIKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProjectAPIKey indicates an expected call of DeleteProjectAPIKey.
func (mr *MockProjectAPIKeyDeleterMockRecorder) DeleteProjectAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectAPIKey", reflect.TypeOf((*MockProjectAPIKeyDeleter)(nil).DeleteProjectAPIKey), arg0, arg1)
}

// MockProjectAPIKeyAssigner is a mock of ProjectAPIKeyAssigner interface.
type MockProjectAPIKeyAssigner struct {
	ctrl     *gomock.Controller
	recorder *MockProjectAPIKeyAssignerMockRecorder
}

// MockProjectAPIKeyAssignerMockRecorder is the mock recorder for MockProjectAPIKeyAssigner.
type MockProjectAPIKeyAssignerMockRecorder struct {
	mock *MockProjectAPIKeyAssigner
}

// NewMockProjectAPIKeyAssigner creates a new mock instance.
func NewMockProjectAPIKeyAssigner(ctrl *gomock.Controller) *MockProjectAPIKeyAssigner {
	mock := &MockProjectAPIKeyAssigner{ctrl: ctrl}
	mock.recorder = &MockProjectAPIKeyAssignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectAPIKeyAssigner) EXPECT() *MockProjectAPIKeyAssignerMockRecorder {
	return m.recorder
}

// AssignProjectAPIKey mocks base method.
func (m *MockProjectAPIKeyAssigner) AssignProjectAPIKey(arg0, arg1 string, arg2 *admin.UpdateAtlasProjectApiKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignProjectAPIKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignProjectAPIKey indicates an expected call of AssignProjectAPIKey.
func (mr *MockProjectAPIKeyAssignerMockRecorder) AssignProjectAPIKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignProjectAPIKey", reflect.TypeOf((*MockProjectAPIKeyAssigner)(nil).AssignProjectAPIKey), arg0, arg1, arg2)
}
