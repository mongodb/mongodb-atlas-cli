// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: LDAPConfigurationVerifier)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockLDAPConfigurationVerifier is a mock of LDAPConfigurationVerifier interface
type MockLDAPConfigurationVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConfigurationVerifierMockRecorder
}

// MockLDAPConfigurationVerifierMockRecorder is the mock recorder for MockLDAPConfigurationVerifier
type MockLDAPConfigurationVerifierMockRecorder struct {
	mock *MockLDAPConfigurationVerifier
}

// NewMockLDAPConfigurationVerifier creates a new mock instance
func NewMockLDAPConfigurationVerifier(ctrl *gomock.Controller) *MockLDAPConfigurationVerifier {
	mock := &MockLDAPConfigurationVerifier{ctrl: ctrl}
	mock.recorder = &MockLDAPConfigurationVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLDAPConfigurationVerifier) EXPECT() *MockLDAPConfigurationVerifierMockRecorder {
	return m.recorder
}

// VerifyLDAPConfiguration mocks base method
func (m *MockLDAPConfigurationVerifier) VerifyLDAPConfiguration(arg0 string, arg1 *mongodbatlas.LDAP) (*mongodbatlas.LDAPConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLDAPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.LDAPConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyLDAPConfiguration indicates an expected call of VerifyLDAPConfiguration
func (mr *MockLDAPConfigurationVerifierMockRecorder) VerifyLDAPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLDAPConfiguration", reflect.TypeOf((*MockLDAPConfigurationVerifier)(nil).VerifyLDAPConfiguration), arg0, arg1)
}
