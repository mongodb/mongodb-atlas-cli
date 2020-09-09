// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: MaintenanceWindowUpdater,MaintenanceWindowClearer,MaintenanceWindowCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
	reflect "reflect"
)

// MockMaintenanceWindowUpdater is a mock of MaintenanceWindowUpdater interface
type MockMaintenanceWindowUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowUpdaterMockRecorder
}

// MockMaintenanceWindowUpdaterMockRecorder is the mock recorder for MockMaintenanceWindowUpdater
type MockMaintenanceWindowUpdaterMockRecorder struct {
	mock *MockMaintenanceWindowUpdater
}

// NewMockMaintenanceWindowUpdater creates a new mock instance
func NewMockMaintenanceWindowUpdater(ctrl *gomock.Controller) *MockMaintenanceWindowUpdater {
	mock := &MockMaintenanceWindowUpdater{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenanceWindowUpdater) EXPECT() *MockMaintenanceWindowUpdaterMockRecorder {
	return m.recorder
}

// UpdateMaintenanceWindow mocks base method
func (m *MockMaintenanceWindowUpdater) UpdateMaintenanceWindow(arg0 string, arg1 *mongodbatlas.MaintenanceWindow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMaintenanceWindow indicates an expected call of UpdateMaintenanceWindow
func (mr *MockMaintenanceWindowUpdaterMockRecorder) UpdateMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowUpdater)(nil).UpdateMaintenanceWindow), arg0, arg1)
}

// MockMaintenanceWindowClearer is a mock of MaintenanceWindowClearer interface
type MockMaintenanceWindowClearer struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowClearerMockRecorder
}

// MockMaintenanceWindowClearerMockRecorder is the mock recorder for MockMaintenanceWindowClearer
type MockMaintenanceWindowClearerMockRecorder struct {
	mock *MockMaintenanceWindowClearer
}

// NewMockMaintenanceWindowClearer creates a new mock instance
func NewMockMaintenanceWindowClearer(ctrl *gomock.Controller) *MockMaintenanceWindowClearer {
	mock := &MockMaintenanceWindowClearer{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowClearerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenanceWindowClearer) EXPECT() *MockMaintenanceWindowClearerMockRecorder {
	return m.recorder
}

// ClearMaintenanceWindow mocks base method
func (m *MockMaintenanceWindowClearer) ClearMaintenanceWindow(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearMaintenanceWindow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearMaintenanceWindow indicates an expected call of ClearMaintenanceWindow
func (mr *MockMaintenanceWindowClearerMockRecorder) ClearMaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowClearer)(nil).ClearMaintenanceWindow), arg0)
}

// MockMaintenanceWindowCreator is a mock of MaintenanceWindowCreator interface
type MockMaintenanceWindowCreator struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowCreatorMockRecorder
}

// MockMaintenanceWindowCreatorMockRecorder is the mock recorder for MockMaintenanceWindowCreator
type MockMaintenanceWindowCreatorMockRecorder struct {
	mock *MockMaintenanceWindowCreator
}

// NewMockMaintenanceWindowCreator creates a new mock instance
func NewMockMaintenanceWindowCreator(ctrl *gomock.Controller) *MockMaintenanceWindowCreator {
	mock := &MockMaintenanceWindowCreator{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenanceWindowCreator) EXPECT() *MockMaintenanceWindowCreatorMockRecorder {
	return m.recorder
}

// CreateMaintenanceWindow mocks base method
func (m *MockMaintenanceWindowCreator) CreateMaintenanceWindow(arg0 string, arg1 *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.MaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMaintenanceWindow indicates an expected call of CreateMaintenanceWindow
func (mr *MockMaintenanceWindowCreatorMockRecorder) CreateMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowCreator)(nil).CreateMaintenanceWindow), arg0, arg1)
}
