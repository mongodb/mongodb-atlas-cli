// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/organizations (interfaces: OrganizationCreator)
//
// Generated by this command:
//
//	mockgen -typed -destination=create_mock_test.go -package=organizations . OrganizationCreator
//

// Package organizations is a generated GoMock package.
package organizations

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockOrganizationCreator is a mock of OrganizationCreator interface.
type MockOrganizationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationCreatorMockRecorder
	isgomock struct{}
}

// MockOrganizationCreatorMockRecorder is the mock recorder for MockOrganizationCreator.
type MockOrganizationCreatorMockRecorder struct {
	mock *MockOrganizationCreator
}

// NewMockOrganizationCreator creates a new mock instance.
func NewMockOrganizationCreator(ctrl *gomock.Controller) *MockOrganizationCreator {
	mock := &MockOrganizationCreator{ctrl: ctrl}
	mock.recorder = &MockOrganizationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationCreator) EXPECT() *MockOrganizationCreatorMockRecorder {
	return m.recorder
}

// CreateAtlasOrganization mocks base method.
func (m *MockOrganizationCreator) CreateAtlasOrganization(arg0 *admin.CreateOrganizationRequest) (*admin.CreateOrganizationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAtlasOrganization", arg0)
	ret0, _ := ret[0].(*admin.CreateOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAtlasOrganization indicates an expected call of CreateAtlasOrganization.
func (mr *MockOrganizationCreatorMockRecorder) CreateAtlasOrganization(arg0 any) *MockOrganizationCreatorCreateAtlasOrganizationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAtlasOrganization", reflect.TypeOf((*MockOrganizationCreator)(nil).CreateAtlasOrganization), arg0)
	return &MockOrganizationCreatorCreateAtlasOrganizationCall{Call: call}
}

// MockOrganizationCreatorCreateAtlasOrganizationCall wrap *gomock.Call
type MockOrganizationCreatorCreateAtlasOrganizationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockOrganizationCreatorCreateAtlasOrganizationCall) Return(arg0 *admin.CreateOrganizationResponse, arg1 error) *MockOrganizationCreatorCreateAtlasOrganizationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockOrganizationCreatorCreateAtlasOrganizationCall) Do(f func(*admin.CreateOrganizationRequest) (*admin.CreateOrganizationResponse, error)) *MockOrganizationCreatorCreateAtlasOrganizationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockOrganizationCreatorCreateAtlasOrganizationCall) DoAndReturn(f func(*admin.CreateOrganizationRequest) (*admin.CreateOrganizationResponse, error)) *MockOrganizationCreatorCreateAtlasOrganizationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
