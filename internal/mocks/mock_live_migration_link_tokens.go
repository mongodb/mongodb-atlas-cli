// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: LinkTokenCreator,LinkTokenDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlasv2 "go.mongodb.org/atlas-sdk/admin"
)

// MockLinkTokenCreator is a mock of LinkTokenCreator interface.
type MockLinkTokenCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLinkTokenCreatorMockRecorder
}

// MockLinkTokenCreatorMockRecorder is the mock recorder for MockLinkTokenCreator.
type MockLinkTokenCreatorMockRecorder struct {
	mock *MockLinkTokenCreator
}

// NewMockLinkTokenCreator creates a new mock instance.
func NewMockLinkTokenCreator(ctrl *gomock.Controller) *MockLinkTokenCreator {
	mock := &MockLinkTokenCreator{ctrl: ctrl}
	mock.recorder = &MockLinkTokenCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkTokenCreator) EXPECT() *MockLinkTokenCreatorMockRecorder {
	return m.recorder
}

// CreateLinkToken mocks base method.
func (m *MockLinkTokenCreator) CreateLinkToken(arg0 string, arg1 *mongodbatlasv2.TargetOrgRequest) (*mongodbatlasv2.TargetOrg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLinkToken", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.TargetOrg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLinkToken indicates an expected call of CreateLinkToken.
func (mr *MockLinkTokenCreatorMockRecorder) CreateLinkToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLinkToken", reflect.TypeOf((*MockLinkTokenCreator)(nil).CreateLinkToken), arg0, arg1)
}

// MockLinkTokenDeleter is a mock of LinkTokenDeleter interface.
type MockLinkTokenDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockLinkTokenDeleterMockRecorder
}

// MockLinkTokenDeleterMockRecorder is the mock recorder for MockLinkTokenDeleter.
type MockLinkTokenDeleterMockRecorder struct {
	mock *MockLinkTokenDeleter
}

// NewMockLinkTokenDeleter creates a new mock instance.
func NewMockLinkTokenDeleter(ctrl *gomock.Controller) *MockLinkTokenDeleter {
	mock := &MockLinkTokenDeleter{ctrl: ctrl}
	mock.recorder = &MockLinkTokenDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkTokenDeleter) EXPECT() *MockLinkTokenDeleterMockRecorder {
	return m.recorder
}

// DeleteLinkToken mocks base method.
func (m *MockLinkTokenDeleter) DeleteLinkToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLinkToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLinkToken indicates an expected call of DeleteLinkToken.
func (mr *MockLinkTokenDeleterMockRecorder) DeleteLinkToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLinkToken", reflect.TypeOf((*MockLinkTokenDeleter)(nil).DeleteLinkToken), arg0)
}
