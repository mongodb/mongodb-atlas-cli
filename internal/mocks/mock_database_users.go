// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: DatabaseUserLister,DatabaseUserCreator,DatabaseUserDeleter,DatabaseUserUpdater,DatabaseUserDescriber,DBUserCertificateLister,DBUserCertificateCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201008/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockDatabaseUserLister is a mock of DatabaseUserLister interface.
type MockDatabaseUserLister struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserListerMockRecorder
}

// MockDatabaseUserListerMockRecorder is the mock recorder for MockDatabaseUserLister.
type MockDatabaseUserListerMockRecorder struct {
	mock *MockDatabaseUserLister
}

// NewMockDatabaseUserLister creates a new mock instance.
func NewMockDatabaseUserLister(ctrl *gomock.Controller) *MockDatabaseUserLister {
	mock := &MockDatabaseUserLister{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseUserLister) EXPECT() *MockDatabaseUserListerMockRecorder {
	return m.recorder
}

// DatabaseUsers mocks base method.
func (m *MockDatabaseUserLister) DatabaseUsers(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedApiAtlasDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUsers", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUsers indicates an expected call of DatabaseUsers.
func (mr *MockDatabaseUserListerMockRecorder) DatabaseUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUsers", reflect.TypeOf((*MockDatabaseUserLister)(nil).DatabaseUsers), arg0, arg1)
}

// MockDatabaseUserCreator is a mock of DatabaseUserCreator interface.
type MockDatabaseUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserCreatorMockRecorder
}

// MockDatabaseUserCreatorMockRecorder is the mock recorder for MockDatabaseUserCreator.
type MockDatabaseUserCreatorMockRecorder struct {
	mock *MockDatabaseUserCreator
}

// NewMockDatabaseUserCreator creates a new mock instance.
func NewMockDatabaseUserCreator(ctrl *gomock.Controller) *MockDatabaseUserCreator {
	mock := &MockDatabaseUserCreator{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseUserCreator) EXPECT() *MockDatabaseUserCreatorMockRecorder {
	return m.recorder
}

// CreateDatabaseUser mocks base method.
func (m *MockDatabaseUserCreator) CreateDatabaseUser(arg0 *admin.CloudDatabaseUser) (*admin.CloudDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseUser", arg0)
	ret0, _ := ret[0].(*admin.CloudDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseUser indicates an expected call of CreateDatabaseUser.
func (mr *MockDatabaseUserCreatorMockRecorder) CreateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseUser", reflect.TypeOf((*MockDatabaseUserCreator)(nil).CreateDatabaseUser), arg0)
}

// MockDatabaseUserDeleter is a mock of DatabaseUserDeleter interface.
type MockDatabaseUserDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserDeleterMockRecorder
}

// MockDatabaseUserDeleterMockRecorder is the mock recorder for MockDatabaseUserDeleter.
type MockDatabaseUserDeleterMockRecorder struct {
	mock *MockDatabaseUserDeleter
}

// NewMockDatabaseUserDeleter creates a new mock instance.
func NewMockDatabaseUserDeleter(ctrl *gomock.Controller) *MockDatabaseUserDeleter {
	mock := &MockDatabaseUserDeleter{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseUserDeleter) EXPECT() *MockDatabaseUserDeleterMockRecorder {
	return m.recorder
}

// DeleteDatabaseUser mocks base method.
func (m *MockDatabaseUserDeleter) DeleteDatabaseUser(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDatabaseUser indicates an expected call of DeleteDatabaseUser.
func (mr *MockDatabaseUserDeleterMockRecorder) DeleteDatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatabaseUser", reflect.TypeOf((*MockDatabaseUserDeleter)(nil).DeleteDatabaseUser), arg0, arg1, arg2)
}

// MockDatabaseUserUpdater is a mock of DatabaseUserUpdater interface.
type MockDatabaseUserUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserUpdaterMockRecorder
}

// MockDatabaseUserUpdaterMockRecorder is the mock recorder for MockDatabaseUserUpdater.
type MockDatabaseUserUpdaterMockRecorder struct {
	mock *MockDatabaseUserUpdater
}

// NewMockDatabaseUserUpdater creates a new mock instance.
func NewMockDatabaseUserUpdater(ctrl *gomock.Controller) *MockDatabaseUserUpdater {
	mock := &MockDatabaseUserUpdater{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseUserUpdater) EXPECT() *MockDatabaseUserUpdaterMockRecorder {
	return m.recorder
}

// UpdateDatabaseUser mocks base method.
func (m *MockDatabaseUserUpdater) UpdateDatabaseUser(arg0 *admin.UpdateDatabaseUserApiParams) (*admin.CloudDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDatabaseUser", arg0)
	ret0, _ := ret[0].(*admin.CloudDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDatabaseUser indicates an expected call of UpdateDatabaseUser.
func (mr *MockDatabaseUserUpdaterMockRecorder) UpdateDatabaseUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDatabaseUser", reflect.TypeOf((*MockDatabaseUserUpdater)(nil).UpdateDatabaseUser), arg0)
}

// MockDatabaseUserDescriber is a mock of DatabaseUserDescriber interface.
type MockDatabaseUserDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseUserDescriberMockRecorder
}

// MockDatabaseUserDescriberMockRecorder is the mock recorder for MockDatabaseUserDescriber.
type MockDatabaseUserDescriberMockRecorder struct {
	mock *MockDatabaseUserDescriber
}

// NewMockDatabaseUserDescriber creates a new mock instance.
func NewMockDatabaseUserDescriber(ctrl *gomock.Controller) *MockDatabaseUserDescriber {
	mock := &MockDatabaseUserDescriber{ctrl: ctrl}
	mock.recorder = &MockDatabaseUserDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseUserDescriber) EXPECT() *MockDatabaseUserDescriberMockRecorder {
	return m.recorder
}

// DatabaseUser mocks base method.
func (m *MockDatabaseUserDescriber) DatabaseUser(arg0, arg1, arg2 string) (*admin.CloudDatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.CloudDatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseUser indicates an expected call of DatabaseUser.
func (mr *MockDatabaseUserDescriberMockRecorder) DatabaseUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseUser", reflect.TypeOf((*MockDatabaseUserDescriber)(nil).DatabaseUser), arg0, arg1, arg2)
}

// MockDBUserCertificateLister is a mock of DBUserCertificateLister interface.
type MockDBUserCertificateLister struct {
	ctrl     *gomock.Controller
	recorder *MockDBUserCertificateListerMockRecorder
}

// MockDBUserCertificateListerMockRecorder is the mock recorder for MockDBUserCertificateLister.
type MockDBUserCertificateListerMockRecorder struct {
	mock *MockDBUserCertificateLister
}

// NewMockDBUserCertificateLister creates a new mock instance.
func NewMockDBUserCertificateLister(ctrl *gomock.Controller) *MockDBUserCertificateLister {
	mock := &MockDBUserCertificateLister{ctrl: ctrl}
	mock.recorder = &MockDBUserCertificateListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBUserCertificateLister) EXPECT() *MockDBUserCertificateListerMockRecorder {
	return m.recorder
}

// DBUserCertificates mocks base method.
func (m *MockDBUserCertificateLister) DBUserCertificates(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*admin.PaginatedUserCert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBUserCertificates", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedUserCert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBUserCertificates indicates an expected call of DBUserCertificates.
func (mr *MockDBUserCertificateListerMockRecorder) DBUserCertificates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBUserCertificates", reflect.TypeOf((*MockDBUserCertificateLister)(nil).DBUserCertificates), arg0, arg1, arg2)
}

// MockDBUserCertificateCreator is a mock of DBUserCertificateCreator interface.
type MockDBUserCertificateCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDBUserCertificateCreatorMockRecorder
}

// MockDBUserCertificateCreatorMockRecorder is the mock recorder for MockDBUserCertificateCreator.
type MockDBUserCertificateCreatorMockRecorder struct {
	mock *MockDBUserCertificateCreator
}

// NewMockDBUserCertificateCreator creates a new mock instance.
func NewMockDBUserCertificateCreator(ctrl *gomock.Controller) *MockDBUserCertificateCreator {
	mock := &MockDBUserCertificateCreator{ctrl: ctrl}
	mock.recorder = &MockDBUserCertificateCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBUserCertificateCreator) EXPECT() *MockDBUserCertificateCreatorMockRecorder {
	return m.recorder
}

// CreateDBUserCertificate mocks base method.
func (m *MockDBUserCertificateCreator) CreateDBUserCertificate(arg0, arg1 string, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDBUserCertificate", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDBUserCertificate indicates an expected call of CreateDBUserCertificate.
func (mr *MockDBUserCertificateCreatorMockRecorder) CreateDBUserCertificate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDBUserCertificate", reflect.TypeOf((*MockDBUserCertificateCreator)(nil).CreateDBUserCertificate), arg0, arg1, arg2)
}
