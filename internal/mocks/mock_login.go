// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/cli/auth (interfaces: Authenticator,LoginConfig)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v4"
	gomock "github.com/golang/mock/gomock"
	auth "go.mongodb.org/atlas/auth"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// PollToken mocks base method.
func (m *MockAuthenticator) PollToken(arg0 context.Context, arg1 *auth.DeviceCode) (*auth.Token, *mongodbatlas.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollToken", arg0, arg1)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(*mongodbatlas.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PollToken indicates an expected call of PollToken.
func (mr *MockAuthenticatorMockRecorder) PollToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollToken", reflect.TypeOf((*MockAuthenticator)(nil).PollToken), arg0, arg1)
}

// RequestCode mocks base method.
func (m *MockAuthenticator) RequestCode(arg0 context.Context) (*auth.DeviceCode, *mongodbatlas.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCode", arg0)
	ret0, _ := ret[0].(*auth.DeviceCode)
	ret1, _ := ret[1].(*mongodbatlas.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RequestCode indicates an expected call of RequestCode.
func (mr *MockAuthenticatorMockRecorder) RequestCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCode", reflect.TypeOf((*MockAuthenticator)(nil).RequestCode), arg0)
}

// MockLoginConfig is a mock of LoginConfig interface.
type MockLoginConfig struct {
	ctrl     *gomock.Controller
	recorder *MockLoginConfigMockRecorder
}

// MockLoginConfigMockRecorder is the mock recorder for MockLoginConfig.
type MockLoginConfigMockRecorder struct {
	mock *MockLoginConfig
}

// NewMockLoginConfig creates a new mock instance.
func NewMockLoginConfig(ctrl *gomock.Controller) *MockLoginConfig {
	mock := &MockLoginConfig{ctrl: ctrl}
	mock.recorder = &MockLoginConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginConfig) EXPECT() *MockLoginConfigMockRecorder {
	return m.recorder
}

// Access mocks base method.
func (m *MockLoginConfig) Access() (*jwt.Token, jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Access")
	ret0, _ := ret[0].(*jwt.Token)
	ret1, _ := ret[1].(jwt.RegisteredClaims)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Access indicates an expected call of Access.
func (mr *MockLoginConfigMockRecorder) Access() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Access", reflect.TypeOf((*MockLoginConfig)(nil).Access))
}

// Save mocks base method.
func (m *MockLoginConfig) Save() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockLoginConfigMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockLoginConfig)(nil).Save))
}

// Set mocks base method.
func (m *MockLoginConfig) Set(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockLoginConfigMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLoginConfig)(nil).Set), arg0, arg1)
}

// SetGlobal mocks base method.
func (m *MockLoginConfig) SetGlobal(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGlobal", arg0, arg1)
}

// SetGlobal indicates an expected call of SetGlobal.
func (mr *MockLoginConfigMockRecorder) SetGlobal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGlobal", reflect.TypeOf((*MockLoginConfig)(nil).SetGlobal), arg0, arg1)
}
