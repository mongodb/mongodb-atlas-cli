// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: ConnectedOrgConfigsUpdater,ConnectedOrgConfigsDescriber,ConnectedOrgConfigsDeleter,ConnectedOrgConfigsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20241113001/admin"
)

// MockConnectedOrgConfigsUpdater is a mock of ConnectedOrgConfigsUpdater interface.
type MockConnectedOrgConfigsUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockConnectedOrgConfigsUpdaterMockRecorder
}

// MockConnectedOrgConfigsUpdaterMockRecorder is the mock recorder for MockConnectedOrgConfigsUpdater.
type MockConnectedOrgConfigsUpdaterMockRecorder struct {
	mock *MockConnectedOrgConfigsUpdater
}

// NewMockConnectedOrgConfigsUpdater creates a new mock instance.
func NewMockConnectedOrgConfigsUpdater(ctrl *gomock.Controller) *MockConnectedOrgConfigsUpdater {
	mock := &MockConnectedOrgConfigsUpdater{ctrl: ctrl}
	mock.recorder = &MockConnectedOrgConfigsUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectedOrgConfigsUpdater) EXPECT() *MockConnectedOrgConfigsUpdaterMockRecorder {
	return m.recorder
}

// UpdateConnectedOrgConfig mocks base method.
func (m *MockConnectedOrgConfigsUpdater) UpdateConnectedOrgConfig(arg0 *admin.UpdateConnectedOrgConfigApiParams) (*admin.ConnectedOrgConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConnectedOrgConfig", arg0)
	ret0, _ := ret[0].(*admin.ConnectedOrgConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConnectedOrgConfig indicates an expected call of UpdateConnectedOrgConfig.
func (mr *MockConnectedOrgConfigsUpdaterMockRecorder) UpdateConnectedOrgConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnectedOrgConfig", reflect.TypeOf((*MockConnectedOrgConfigsUpdater)(nil).UpdateConnectedOrgConfig), arg0)
}

// MockConnectedOrgConfigsDescriber is a mock of ConnectedOrgConfigsDescriber interface.
type MockConnectedOrgConfigsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockConnectedOrgConfigsDescriberMockRecorder
}

// MockConnectedOrgConfigsDescriberMockRecorder is the mock recorder for MockConnectedOrgConfigsDescriber.
type MockConnectedOrgConfigsDescriberMockRecorder struct {
	mock *MockConnectedOrgConfigsDescriber
}

// NewMockConnectedOrgConfigsDescriber creates a new mock instance.
func NewMockConnectedOrgConfigsDescriber(ctrl *gomock.Controller) *MockConnectedOrgConfigsDescriber {
	mock := &MockConnectedOrgConfigsDescriber{ctrl: ctrl}
	mock.recorder = &MockConnectedOrgConfigsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectedOrgConfigsDescriber) EXPECT() *MockConnectedOrgConfigsDescriberMockRecorder {
	return m.recorder
}

// GetConnectedOrgConfig mocks base method.
func (m *MockConnectedOrgConfigsDescriber) GetConnectedOrgConfig(arg0 *admin.GetConnectedOrgConfigApiParams) (*admin.ConnectedOrgConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectedOrgConfig", arg0)
	ret0, _ := ret[0].(*admin.ConnectedOrgConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectedOrgConfig indicates an expected call of GetConnectedOrgConfig.
func (mr *MockConnectedOrgConfigsDescriberMockRecorder) GetConnectedOrgConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectedOrgConfig", reflect.TypeOf((*MockConnectedOrgConfigsDescriber)(nil).GetConnectedOrgConfig), arg0)
}

// MockConnectedOrgConfigsDeleter is a mock of ConnectedOrgConfigsDeleter interface.
type MockConnectedOrgConfigsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockConnectedOrgConfigsDeleterMockRecorder
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
func (m *MockConnectedOrgConfigsDeleter) DeleteConnectedOrgConfig(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConnectedOrgConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConnectedOrgConfig indicates an expected call of DeleteConnectedOrgConfig.
func (mr *MockConnectedOrgConfigsDeleterMockRecorder) DeleteConnectedOrgConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConnectedOrgConfig", reflect.TypeOf((*MockConnectedOrgConfigsDeleter)(nil).DeleteConnectedOrgConfig), arg0, arg1)
}

// MockConnectedOrgConfigsLister is a mock of ConnectedOrgConfigsLister interface.
type MockConnectedOrgConfigsLister struct {
	ctrl     *gomock.Controller
	recorder *MockConnectedOrgConfigsListerMockRecorder
}

// MockConnectedOrgConfigsListerMockRecorder is the mock recorder for MockConnectedOrgConfigsLister.
type MockConnectedOrgConfigsListerMockRecorder struct {
	mock *MockConnectedOrgConfigsLister
}

// NewMockConnectedOrgConfigsLister creates a new mock instance.
func NewMockConnectedOrgConfigsLister(ctrl *gomock.Controller) *MockConnectedOrgConfigsLister {
	mock := &MockConnectedOrgConfigsLister{ctrl: ctrl}
	mock.recorder = &MockConnectedOrgConfigsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectedOrgConfigsLister) EXPECT() *MockConnectedOrgConfigsListerMockRecorder {
	return m.recorder
}

// ListConnectedOrgConfigs mocks base method.
func (m *MockConnectedOrgConfigsLister) ListConnectedOrgConfigs(arg0 *admin.ListConnectedOrgConfigsApiParams) (*admin.PaginatedConnectedOrgConfigs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnectedOrgConfigs", arg0)
	ret0, _ := ret[0].(*admin.PaginatedConnectedOrgConfigs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectedOrgConfigs indicates an expected call of ListConnectedOrgConfigs.
func (mr *MockConnectedOrgConfigsListerMockRecorder) ListConnectedOrgConfigs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectedOrgConfigs", reflect.TypeOf((*MockConnectedOrgConfigsLister)(nil).ListConnectedOrgConfigs), arg0)
}
