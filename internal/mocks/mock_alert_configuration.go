// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/alert_configuration.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockAlertConfigurationLister is a mock of AlertConfigurationLister interface
type MockAlertConfigurationLister struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationListerMockRecorder
}

// MockAlertConfigurationListerMockRecorder is the mock recorder for MockAlertConfigurationLister
type MockAlertConfigurationListerMockRecorder struct {
	mock *MockAlertConfigurationLister
}

// NewMockAlertConfigurationLister creates a new mock instance
func NewMockAlertConfigurationLister(ctrl *gomock.Controller) *MockAlertConfigurationLister {
	mock := &MockAlertConfigurationLister{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlertConfigurationLister) EXPECT() *MockAlertConfigurationListerMockRecorder {
	return m.recorder
}

// ProjectAlertConfiguration mocks base method
func (m *MockAlertConfigurationLister) ProjectAlertConfiguration(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.AlertConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectAlertConfiguration", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.AlertConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectAlertConfiguration indicates an expected call of ProjectAlertConfiguration
func (mr *MockAlertConfigurationListerMockRecorder) ProjectAlertConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectAlertConfiguration", reflect.TypeOf((*MockAlertConfigurationLister)(nil).ProjectAlertConfiguration), arg0, arg1)
}

// MockAlertConfigurationStore is a mock of AlertConfigurationStore interface
type MockAlertConfigurationStore struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationStoreMockRecorder
}

// MockAlertConfigurationStoreMockRecorder is the mock recorder for MockAlertConfigurationStore
type MockAlertConfigurationStoreMockRecorder struct {
	mock *MockAlertConfigurationStore
}

// NewMockAlertConfigurationStore creates a new mock instance
func NewMockAlertConfigurationStore(ctrl *gomock.Controller) *MockAlertConfigurationStore {
	mock := &MockAlertConfigurationStore{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlertConfigurationStore) EXPECT() *MockAlertConfigurationStoreMockRecorder {
	return m.recorder
}

// ProjectAlertConfiguration mocks base method
func (m *MockAlertConfigurationStore) ProjectAlertConfiguration(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.AlertConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectAlertConfiguration", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.AlertConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectAlertConfiguration indicates an expected call of ProjectAlertConfiguration
func (mr *MockAlertConfigurationStoreMockRecorder) ProjectAlertConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectAlertConfiguration", reflect.TypeOf((*MockAlertConfigurationStore)(nil).ProjectAlertConfiguration), arg0, arg1)
}
