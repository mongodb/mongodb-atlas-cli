// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/dbusers/certs (interfaces: DBUserCertificateCreator)
//
// Generated by this command:
//
//	mockgen -typed -destination=create_mock_test.go -package=certs . DBUserCertificateCreator
//

// Package certs is a generated GoMock package.
package certs

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDBUserCertificateCreator is a mock of DBUserCertificateCreator interface.
type MockDBUserCertificateCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDBUserCertificateCreatorMockRecorder
	isgomock struct{}
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
func (mr *MockDBUserCertificateCreatorMockRecorder) CreateDBUserCertificate(arg0, arg1, arg2 any) *MockDBUserCertificateCreatorCreateDBUserCertificateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDBUserCertificate", reflect.TypeOf((*MockDBUserCertificateCreator)(nil).CreateDBUserCertificate), arg0, arg1, arg2)
	return &MockDBUserCertificateCreatorCreateDBUserCertificateCall{Call: call}
}

// MockDBUserCertificateCreatorCreateDBUserCertificateCall wrap *gomock.Call
type MockDBUserCertificateCreatorCreateDBUserCertificateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDBUserCertificateCreatorCreateDBUserCertificateCall) Return(arg0 string, arg1 error) *MockDBUserCertificateCreatorCreateDBUserCertificateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDBUserCertificateCreatorCreateDBUserCertificateCall) Do(f func(string, string, int) (string, error)) *MockDBUserCertificateCreatorCreateDBUserCertificateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDBUserCertificateCreatorCreateDBUserCertificateCall) DoAndReturn(f func(string, string, int) (string, error)) *MockDBUserCertificateCreatorCreateDBUserCertificateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
