// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: X509CertificateDescriber,X509CertificateSaver,X509CertificateStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockX509CertificateDescriber is a mock of X509CertificateDescriber interface
type MockX509CertificateDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockX509CertificateDescriberMockRecorder
}

// MockX509CertificateDescriberMockRecorder is the mock recorder for MockX509CertificateDescriber
type MockX509CertificateDescriberMockRecorder struct {
	mock *MockX509CertificateDescriber
}

// NewMockX509CertificateDescriber creates a new mock instance
func NewMockX509CertificateDescriber(ctrl *gomock.Controller) *MockX509CertificateDescriber {
	mock := &MockX509CertificateDescriber{ctrl: ctrl}
	mock.recorder = &MockX509CertificateDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockX509CertificateDescriber) EXPECT() *MockX509CertificateDescriberMockRecorder {
	return m.recorder
}

// X509Configuration mocks base method
func (m *MockX509CertificateDescriber) X509Configuration(arg0 string) (*mongodbatlas.CustomerX509, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "X509Configuration", arg0)
	ret0, _ := ret[0].(*mongodbatlas.CustomerX509)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// X509Configuration indicates an expected call of X509Configuration
func (mr *MockX509CertificateDescriberMockRecorder) X509Configuration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "X509Configuration", reflect.TypeOf((*MockX509CertificateDescriber)(nil).X509Configuration), arg0)
}

// MockX509CertificateSaver is a mock of X509CertificateSaver interface
type MockX509CertificateSaver struct {
	ctrl     *gomock.Controller
	recorder *MockX509CertificateSaverMockRecorder
}

// MockX509CertificateSaverMockRecorder is the mock recorder for MockX509CertificateSaver
type MockX509CertificateSaverMockRecorder struct {
	mock *MockX509CertificateSaver
}

// NewMockX509CertificateSaver creates a new mock instance
func NewMockX509CertificateSaver(ctrl *gomock.Controller) *MockX509CertificateSaver {
	mock := &MockX509CertificateSaver{ctrl: ctrl}
	mock.recorder = &MockX509CertificateSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockX509CertificateSaver) EXPECT() *MockX509CertificateSaverMockRecorder {
	return m.recorder
}

// SaveX509Configuration mocks base method
func (m *MockX509CertificateSaver) SaveX509Configuration(arg0, arg1 string) (*mongodbatlas.CustomerX509, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveX509Configuration", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.CustomerX509)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveX509Configuration indicates an expected call of SaveX509Configuration
func (mr *MockX509CertificateSaverMockRecorder) SaveX509Configuration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveX509Configuration", reflect.TypeOf((*MockX509CertificateSaver)(nil).SaveX509Configuration), arg0, arg1)
}

// MockX509CertificateStore is a mock of X509CertificateStore interface
type MockX509CertificateStore struct {
	ctrl     *gomock.Controller
	recorder *MockX509CertificateStoreMockRecorder
}

// MockX509CertificateStoreMockRecorder is the mock recorder for MockX509CertificateStore
type MockX509CertificateStoreMockRecorder struct {
	mock *MockX509CertificateStore
}

// NewMockX509CertificateStore creates a new mock instance
func NewMockX509CertificateStore(ctrl *gomock.Controller) *MockX509CertificateStore {
	mock := &MockX509CertificateStore{ctrl: ctrl}
	mock.recorder = &MockX509CertificateStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockX509CertificateStore) EXPECT() *MockX509CertificateStoreMockRecorder {
	return m.recorder
}

// DisableX509Configuration mocks base method
func (m *MockX509CertificateStore) DisableX509Configuration(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableX509Configuration", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableX509Configuration indicates an expected call of DisableX509Configuration
func (mr *MockX509CertificateStoreMockRecorder) DisableX509Configuration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableX509Configuration", reflect.TypeOf((*MockX509CertificateStore)(nil).DisableX509Configuration), arg0)
}

// GetUserCertificates mocks base method
func (m *MockX509CertificateStore) GetUserCertificates(arg0, arg1 string) ([]mongodbatlas.UserCertificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCertificates", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.UserCertificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCertificates indicates an expected call of GetUserCertificates
func (mr *MockX509CertificateStoreMockRecorder) GetUserCertificates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCertificates", reflect.TypeOf((*MockX509CertificateStore)(nil).GetUserCertificates), arg0, arg1)
}

// SaveX509Configuration mocks base method
func (m *MockX509CertificateStore) SaveX509Configuration(arg0, arg1 string) (*mongodbatlas.CustomerX509, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveX509Configuration", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.CustomerX509)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveX509Configuration indicates an expected call of SaveX509Configuration
func (mr *MockX509CertificateStoreMockRecorder) SaveX509Configuration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveX509Configuration", reflect.TypeOf((*MockX509CertificateStore)(nil).SaveX509Configuration), arg0, arg1)
}

// X509Configuration mocks base method
func (m *MockX509CertificateStore) X509Configuration(arg0 string) (*mongodbatlas.CustomerX509, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "X509Configuration", arg0)
	ret0, _ := ret[0].(*mongodbatlas.CustomerX509)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// X509Configuration indicates an expected call of X509Configuration
func (mr *MockX509CertificateStoreMockRecorder) X509Configuration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "X509Configuration", reflect.TypeOf((*MockX509CertificateStore)(nil).X509Configuration), arg0)
}
