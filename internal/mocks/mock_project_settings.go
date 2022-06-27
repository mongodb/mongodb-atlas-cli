// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: ProjectSettingsDescriber,ProjectSettingsUpdater,ProjectSettingsGetterUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockProjectSettingsDescriber is a mock of ProjectSettingsDescriber interface.
type MockProjectSettingsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockProjectSettingsDescriberMockRecorder
}

// MockProjectSettingsDescriberMockRecorder is the mock recorder for MockProjectSettingsDescriber.
type MockProjectSettingsDescriberMockRecorder struct {
	mock *MockProjectSettingsDescriber
}

// NewMockProjectSettingsDescriber creates a new mock instance.
func NewMockProjectSettingsDescriber(ctrl *gomock.Controller) *MockProjectSettingsDescriber {
	mock := &MockProjectSettingsDescriber{ctrl: ctrl}
	mock.recorder = &MockProjectSettingsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectSettingsDescriber) EXPECT() *MockProjectSettingsDescriberMockRecorder {
	return m.recorder
}

// ProjectSettings mocks base method.
func (m *MockProjectSettingsDescriber) ProjectSettings(arg0 string) (*mongodbatlas.ProjectSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*mongodbatlas.ProjectSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectSettings indicates an expected call of ProjectSettings.
func (mr *MockProjectSettingsDescriberMockRecorder) ProjectSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectSettings", reflect.TypeOf((*MockProjectSettingsDescriber)(nil).ProjectSettings), arg0)
}

// MockProjectSettingsUpdater is a mock of ProjectSettingsUpdater interface.
type MockProjectSettingsUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockProjectSettingsUpdaterMockRecorder
}

// MockProjectSettingsUpdaterMockRecorder is the mock recorder for MockProjectSettingsUpdater.
type MockProjectSettingsUpdaterMockRecorder struct {
	mock *MockProjectSettingsUpdater
}

// NewMockProjectSettingsUpdater creates a new mock instance.
func NewMockProjectSettingsUpdater(ctrl *gomock.Controller) *MockProjectSettingsUpdater {
	mock := &MockProjectSettingsUpdater{ctrl: ctrl}
	mock.recorder = &MockProjectSettingsUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectSettingsUpdater) EXPECT() *MockProjectSettingsUpdaterMockRecorder {
	return m.recorder
}

// UpdateProjectSettings mocks base method.
func (m *MockProjectSettingsUpdater) UpdateProjectSettings(arg0 string, arg1 *mongodbatlas.ProjectSettings) (*mongodbatlas.ProjectSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectSettings", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.ProjectSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectSettings indicates an expected call of UpdateProjectSettings.
func (mr *MockProjectSettingsUpdaterMockRecorder) UpdateProjectSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectSettings", reflect.TypeOf((*MockProjectSettingsUpdater)(nil).UpdateProjectSettings), arg0, arg1)
}

// MockProjectSettingsGetterUpdater is a mock of ProjectSettingsGetterUpdater interface.
type MockProjectSettingsGetterUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockProjectSettingsGetterUpdaterMockRecorder
}

// MockProjectSettingsGetterUpdaterMockRecorder is the mock recorder for MockProjectSettingsGetterUpdater.
type MockProjectSettingsGetterUpdaterMockRecorder struct {
	mock *MockProjectSettingsGetterUpdater
}

// NewMockProjectSettingsGetterUpdater creates a new mock instance.
func NewMockProjectSettingsGetterUpdater(ctrl *gomock.Controller) *MockProjectSettingsGetterUpdater {
	mock := &MockProjectSettingsGetterUpdater{ctrl: ctrl}
	mock.recorder = &MockProjectSettingsGetterUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectSettingsGetterUpdater) EXPECT() *MockProjectSettingsGetterUpdaterMockRecorder {
	return m.recorder
}

// ProjectSettings mocks base method.
func (m *MockProjectSettingsGetterUpdater) ProjectSettings(arg0 string) (*mongodbatlas.ProjectSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*mongodbatlas.ProjectSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectSettings indicates an expected call of ProjectSettings.
func (mr *MockProjectSettingsGetterUpdaterMockRecorder) ProjectSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectSettings", reflect.TypeOf((*MockProjectSettingsGetterUpdater)(nil).ProjectSettings), arg0)
}

// UpdateProjectSettings mocks base method.
func (m *MockProjectSettingsGetterUpdater) UpdateProjectSettings(arg0 string, arg1 *mongodbatlas.ProjectSettings) (*mongodbatlas.ProjectSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectSettings", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.ProjectSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectSettings indicates an expected call of UpdateProjectSettings.
func (mr *MockProjectSettingsGetterUpdaterMockRecorder) UpdateProjectSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectSettings", reflect.TypeOf((*MockProjectSettingsGetterUpdater)(nil).UpdateProjectSettings), arg0, arg1)
}
