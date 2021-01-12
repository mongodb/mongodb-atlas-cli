// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: CloudProviderAccessRoleCreator,CloudProviderAccessRoleEnabler,CloudProviderAccessRoleLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockCloudProviderAccessRoleCreator is a mock of CloudProviderAccessRoleCreator interface
type MockCloudProviderAccessRoleCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleCreatorMockRecorder
}

// MockCloudProviderAccessRoleCreatorMockRecorder is the mock recorder for MockCloudProviderAccessRoleCreator
type MockCloudProviderAccessRoleCreatorMockRecorder struct {
	mock *MockCloudProviderAccessRoleCreator
}

// NewMockCloudProviderAccessRoleCreator creates a new mock instance
func NewMockCloudProviderAccessRoleCreator(ctrl *gomock.Controller) *MockCloudProviderAccessRoleCreator {
	mock := &MockCloudProviderAccessRoleCreator{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudProviderAccessRoleCreator) EXPECT() *MockCloudProviderAccessRoleCreatorMockRecorder {
	return m.recorder
}

// CreateCloudProviderAccessRole mocks base method
func (m *MockCloudProviderAccessRoleCreator) CreateCloudProviderAccessRole(arg0, arg1 string) (*mongodbatlas.AWSIAMRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCloudProviderAccessRole", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.AWSIAMRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCloudProviderAccessRole indicates an expected call of CreateCloudProviderAccessRole
func (mr *MockCloudProviderAccessRoleCreatorMockRecorder) CreateCloudProviderAccessRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCloudProviderAccessRole", reflect.TypeOf((*MockCloudProviderAccessRoleCreator)(nil).CreateCloudProviderAccessRole), arg0, arg1)
}

// MockCloudProviderAccessRoleEnabler is a mock of CloudProviderAccessRoleEnabler interface
type MockCloudProviderAccessRoleEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleEnablerMockRecorder
}

// MockCloudProviderAccessRoleEnablerMockRecorder is the mock recorder for MockCloudProviderAccessRoleEnabler
type MockCloudProviderAccessRoleEnablerMockRecorder struct {
	mock *MockCloudProviderAccessRoleEnabler
}

// NewMockCloudProviderAccessRoleEnabler creates a new mock instance
func NewMockCloudProviderAccessRoleEnabler(ctrl *gomock.Controller) *MockCloudProviderAccessRoleEnabler {
	mock := &MockCloudProviderAccessRoleEnabler{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudProviderAccessRoleEnabler) EXPECT() *MockCloudProviderAccessRoleEnablerMockRecorder {
	return m.recorder
}

// EnableCloudProviderAccessRole mocks base method
func (m *MockCloudProviderAccessRoleEnabler) EnableCloudProviderAccessRole(arg0, arg1 string, arg2 *mongodbatlas.CloudProviderAuthorizationRequest) (*mongodbatlas.AWSIAMRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableCloudProviderAccessRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.AWSIAMRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableCloudProviderAccessRole indicates an expected call of EnableCloudProviderAccessRole
func (mr *MockCloudProviderAccessRoleEnablerMockRecorder) EnableCloudProviderAccessRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCloudProviderAccessRole", reflect.TypeOf((*MockCloudProviderAccessRoleEnabler)(nil).EnableCloudProviderAccessRole), arg0, arg1, arg2)
}

// MockCloudProviderAccessRoleLister is a mock of CloudProviderAccessRoleLister interface
type MockCloudProviderAccessRoleLister struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderAccessRoleListerMockRecorder
}

// MockCloudProviderAccessRoleListerMockRecorder is the mock recorder for MockCloudProviderAccessRoleLister
type MockCloudProviderAccessRoleListerMockRecorder struct {
	mock *MockCloudProviderAccessRoleLister
}

// NewMockCloudProviderAccessRoleLister creates a new mock instance
func NewMockCloudProviderAccessRoleLister(ctrl *gomock.Controller) *MockCloudProviderAccessRoleLister {
	mock := &MockCloudProviderAccessRoleLister{ctrl: ctrl}
	mock.recorder = &MockCloudProviderAccessRoleListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudProviderAccessRoleLister) EXPECT() *MockCloudProviderAccessRoleListerMockRecorder {
	return m.recorder
}

// CloudProviderAccessRoles mocks base method
func (m *MockCloudProviderAccessRoleLister) CloudProviderAccessRoles(arg0 string) (*mongodbatlas.CloudProviderAccessRoles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(*mongodbatlas.CloudProviderAccessRoles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderAccessRoles indicates an expected call of CloudProviderAccessRoles
func (mr *MockCloudProviderAccessRoleListerMockRecorder) CloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderAccessRoles", reflect.TypeOf((*MockCloudProviderAccessRoleLister)(nil).CloudProviderAccessRoles), arg0)
}
