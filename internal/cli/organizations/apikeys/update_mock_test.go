// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/organizations/apikeys (interfaces: OrganizationAPIKeyUpdater)
//
// Generated by this command:
//
//	mockgen -typed -destination=update_mock_test.go -package=apikeys . OrganizationAPIKeyUpdater
//

// Package apikeys is a generated GoMock package.
package apikeys

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockOrganizationAPIKeyUpdater is a mock of OrganizationAPIKeyUpdater interface.
type MockOrganizationAPIKeyUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyUpdaterMockRecorder
	isgomock struct{}
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
func (mr *MockOrganizationAPIKeyUpdaterMockRecorder) UpdateOrganizationAPIKey(arg0, arg1, arg2 any) *MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganizationAPIKey", reflect.TypeOf((*MockOrganizationAPIKeyUpdater)(nil).UpdateOrganizationAPIKey), arg0, arg1, arg2)
	return &MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall{Call: call}
}

// MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall wrap *gomock.Call
type MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall) Return(arg0 *admin.ApiKeyUserDetails, arg1 error) *MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall) Do(f func(string, string, *admin.UpdateAtlasOrganizationApiKey) (*admin.ApiKeyUserDetails, error)) *MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall) DoAndReturn(f func(string, string, *admin.UpdateAtlasOrganizationApiKey) (*admin.ApiKeyUserDetails, error)) *MockOrganizationAPIKeyUpdaterUpdateOrganizationAPIKeyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
