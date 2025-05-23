// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/federatedauthentication/federationsettings/identityprovider (interfaces: JwkRevoker)
//
// Generated by this command:
//
//	mockgen -typed -destination=revokejwk_mock_test.go -package=identityprovider . JwkRevoker
//

// Package identityprovider is a generated GoMock package.
package identityprovider

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockJwkRevoker is a mock of JwkRevoker interface.
type MockJwkRevoker struct {
	ctrl     *gomock.Controller
	recorder *MockJwkRevokerMockRecorder
	isgomock struct{}
}

// MockJwkRevokerMockRecorder is the mock recorder for MockJwkRevoker.
type MockJwkRevokerMockRecorder struct {
	mock *MockJwkRevoker
}

// NewMockJwkRevoker creates a new mock instance.
func NewMockJwkRevoker(ctrl *gomock.Controller) *MockJwkRevoker {
	mock := &MockJwkRevoker{ctrl: ctrl}
	mock.recorder = &MockJwkRevokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwkRevoker) EXPECT() *MockJwkRevokerMockRecorder {
	return m.recorder
}

// RevokeJwksFromIdentityProvider mocks base method.
func (m *MockJwkRevoker) RevokeJwksFromIdentityProvider(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeJwksFromIdentityProvider", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeJwksFromIdentityProvider indicates an expected call of RevokeJwksFromIdentityProvider.
func (mr *MockJwkRevokerMockRecorder) RevokeJwksFromIdentityProvider(arg0, arg1 any) *MockJwkRevokerRevokeJwksFromIdentityProviderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeJwksFromIdentityProvider", reflect.TypeOf((*MockJwkRevoker)(nil).RevokeJwksFromIdentityProvider), arg0, arg1)
	return &MockJwkRevokerRevokeJwksFromIdentityProviderCall{Call: call}
}

// MockJwkRevokerRevokeJwksFromIdentityProviderCall wrap *gomock.Call
type MockJwkRevokerRevokeJwksFromIdentityProviderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJwkRevokerRevokeJwksFromIdentityProviderCall) Return(arg0 error) *MockJwkRevokerRevokeJwksFromIdentityProviderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJwkRevokerRevokeJwksFromIdentityProviderCall) Do(f func(string, string) error) *MockJwkRevokerRevokeJwksFromIdentityProviderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJwkRevokerRevokeJwksFromIdentityProviderCall) DoAndReturn(f func(string, string) error) *MockJwkRevokerRevokeJwksFromIdentityProviderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
