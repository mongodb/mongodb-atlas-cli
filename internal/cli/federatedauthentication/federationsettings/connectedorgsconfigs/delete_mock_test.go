// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/federatedauthentication/federationsettings/connectedorgsconfigs (interfaces: ConnectedOrgConfigsDeleter)
//
// Generated by this command:
//
//	mockgen -typed -destination=delete_mock_test.go -package=connectedorgsconfigs . ConnectedOrgConfigsDeleter
//

// Package connectedorgsconfigs is a generated GoMock package.
package connectedorgsconfigs

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConnectedOrgConfigsDeleter is a mock of ConnectedOrgConfigsDeleter interface.
type MockConnectedOrgConfigsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockConnectedOrgConfigsDeleterMockRecorder
	isgomock struct{}
}

// MockConnectedOrgConfigsDeleterMockRecorder is the mock recorder for MockConnectedOrgConfigsDeleter.
type MockConnectedOrgConfigsDeleterMockRecorder struct {
	mock *MockConnectedOrgConfigsDeleter
}

// NewMockConnectedOrgConfigsDeleter creates a new mock instance.
func NewMockConnectedOrgConfigsDeleter(ctrl *gomock.Controller) *MockConnectedOrgConfigsDeleter {
	mock := &MockConnectedOrgConfigsDeleter{ctrl: ctrl}
	mock.recorder = &MockConnectedOrgConfigsDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectedOrgConfigsDeleter) EXPECT() *MockConnectedOrgConfigsDeleterMockRecorder {
	return m.recorder
}

// DeleteConnectedOrgConfig mocks base method.
func (m *MockConnectedOrgConfigsDeleter) DeleteConnectedOrgConfig(federationSettingsID, orgID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConnectedOrgConfig", federationSettingsID, orgID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConnectedOrgConfig indicates an expected call of DeleteConnectedOrgConfig.
func (mr *MockConnectedOrgConfigsDeleterMockRecorder) DeleteConnectedOrgConfig(federationSettingsID, orgID any) *MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConnectedOrgConfig", reflect.TypeOf((*MockConnectedOrgConfigsDeleter)(nil).DeleteConnectedOrgConfig), federationSettingsID, orgID)
	return &MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall{Call: call}
}

// MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall wrap *gomock.Call
type MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall) Return(arg0 error) *MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall) Do(f func(string, string) error) *MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall) DoAndReturn(f func(string, string) error) *MockConnectedOrgConfigsDeleterDeleteConnectedOrgConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
