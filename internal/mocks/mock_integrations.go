// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: IntegrationCreator,IntegrationLister,IntegrationDeleter,IntegrationDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

// MockIntegrationCreator is a mock of IntegrationCreator interface.
type MockIntegrationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockIntegrationCreatorMockRecorder
}

// MockIntegrationCreatorMockRecorder is the mock recorder for MockIntegrationCreator.
type MockIntegrationCreatorMockRecorder struct {
	mock *MockIntegrationCreator
}

// NewMockIntegrationCreator creates a new mock instance.
func NewMockIntegrationCreator(ctrl *gomock.Controller) *MockIntegrationCreator {
	mock := &MockIntegrationCreator{ctrl: ctrl}
	mock.recorder = &MockIntegrationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntegrationCreator) EXPECT() *MockIntegrationCreatorMockRecorder {
	return m.recorder
}

// CreateIntegration mocks base method.
func (m *MockIntegrationCreator) CreateIntegration(arg0, arg1 string, arg2 *admin.ThirdPartyIntegration) (*admin.PaginatedIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIntegration", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIntegration indicates an expected call of CreateIntegration.
func (mr *MockIntegrationCreatorMockRecorder) CreateIntegration(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIntegration", reflect.TypeOf((*MockIntegrationCreator)(nil).CreateIntegration), arg0, arg1, arg2)
}

// MockIntegrationLister is a mock of IntegrationLister interface.
type MockIntegrationLister struct {
	ctrl     *gomock.Controller
	recorder *MockIntegrationListerMockRecorder
}

// MockIntegrationListerMockRecorder is the mock recorder for MockIntegrationLister.
type MockIntegrationListerMockRecorder struct {
	mock *MockIntegrationLister
}

// NewMockIntegrationLister creates a new mock instance.
func NewMockIntegrationLister(ctrl *gomock.Controller) *MockIntegrationLister {
	mock := &MockIntegrationLister{ctrl: ctrl}
	mock.recorder = &MockIntegrationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntegrationLister) EXPECT() *MockIntegrationListerMockRecorder {
	return m.recorder
}

// Integrations mocks base method.
func (m *MockIntegrationLister) Integrations(arg0 string) (*admin.PaginatedIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Integrations", arg0)
	ret0, _ := ret[0].(*admin.PaginatedIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Integrations indicates an expected call of Integrations.
func (mr *MockIntegrationListerMockRecorder) Integrations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Integrations", reflect.TypeOf((*MockIntegrationLister)(nil).Integrations), arg0)
}

// MockIntegrationDeleter is a mock of IntegrationDeleter interface.
type MockIntegrationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockIntegrationDeleterMockRecorder
}

// MockIntegrationDeleterMockRecorder is the mock recorder for MockIntegrationDeleter.
type MockIntegrationDeleterMockRecorder struct {
	mock *MockIntegrationDeleter
}

// NewMockIntegrationDeleter creates a new mock instance.
func NewMockIntegrationDeleter(ctrl *gomock.Controller) *MockIntegrationDeleter {
	mock := &MockIntegrationDeleter{ctrl: ctrl}
	mock.recorder = &MockIntegrationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntegrationDeleter) EXPECT() *MockIntegrationDeleterMockRecorder {
	return m.recorder
}

// DeleteIntegration mocks base method.
func (m *MockIntegrationDeleter) DeleteIntegration(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIntegration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIntegration indicates an expected call of DeleteIntegration.
func (mr *MockIntegrationDeleterMockRecorder) DeleteIntegration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIntegration", reflect.TypeOf((*MockIntegrationDeleter)(nil).DeleteIntegration), arg0, arg1)
}

// MockIntegrationDescriber is a mock of IntegrationDescriber interface.
type MockIntegrationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockIntegrationDescriberMockRecorder
}

// MockIntegrationDescriberMockRecorder is the mock recorder for MockIntegrationDescriber.
type MockIntegrationDescriberMockRecorder struct {
	mock *MockIntegrationDescriber
}

// NewMockIntegrationDescriber creates a new mock instance.
func NewMockIntegrationDescriber(ctrl *gomock.Controller) *MockIntegrationDescriber {
	mock := &MockIntegrationDescriber{ctrl: ctrl}
	mock.recorder = &MockIntegrationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntegrationDescriber) EXPECT() *MockIntegrationDescriberMockRecorder {
	return m.recorder
}

// Integration mocks base method.
func (m *MockIntegrationDescriber) Integration(arg0, arg1 string) (*admin.ThirdPartyIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Integration", arg0, arg1)
	ret0, _ := ret[0].(*admin.ThirdPartyIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Integration indicates an expected call of Integration.
func (mr *MockIntegrationDescriberMockRecorder) Integration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Integration", reflect.TypeOf((*MockIntegrationDescriber)(nil).Integration), arg0, arg1)
}
