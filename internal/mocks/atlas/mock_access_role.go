// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: CloudProviderAccessRoleLister)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231001002/admin"
)

// MockCloudProviderAccessRoleLister is a mock of CloudProviderAccessRoleLister interface.
type MockCloudProviderAccessRoleLister struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleListerMockRecorder
}

// MockCloudProviderAccessRoleListerMockRecorder is the mock recorder for MockCloudProviderAccessRoleLister.
type MockCloudProviderAccessRoleListerMockRecorder struct {
	mock *MockCloudProviderAccessRoleLister
}

// NewMockCloudProviderAccessRoleLister creates a new mock instance.
func NewMockCloudProviderAccessRoleLister(ctrl *gomock.Controller) *MockCloudProviderAccessRoleLister {
	mock := &MockCloudProviderAccessRoleLister{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderAccessRoleLister) EXPECT() *MockCloudProviderAccessRoleListerMockRecorder {
	return m.recorder
}

// CloudProviderAccessRoles mocks base method.
func (m *MockCloudProviderAccessRoleLister) CloudProviderAccessRoles(arg0 string) (*admin.CloudProviderAccessRoles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(*admin.CloudProviderAccessRoles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderAccessRoles indicates an expected call of CloudProviderAccessRoles.
func (mr *MockCloudProviderAccessRoleListerMockRecorder) CloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderAccessRoles", reflect.TypeOf((*MockCloudProviderAccessRoleLister)(nil).CloudProviderAccessRoles), arg0)
}
