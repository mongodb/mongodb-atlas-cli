// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: ProjectSettingsDescriber,ProjectSettingsUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	mongodbatlasv2 "go.mongodb.org/atlas-sdk/admin"
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
func (m *MockProjectSettingsDescriber) ProjectSettings(arg0 string) (*mongodbatlasv2.GroupSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.GroupSettings)
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
func (m *MockProjectSettingsUpdater) UpdateProjectSettings(arg0 string, arg1 *mongodbatlas.ProjectSettings) (*mongodbatlasv2.GroupSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectSettings", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.GroupSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectSettings indicates an expected call of UpdateProjectSettings.
func (mr *MockProjectSettingsUpdaterMockRecorder) UpdateProjectSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectSettings", reflect.TypeOf((*MockProjectSettingsUpdater)(nil).UpdateProjectSettings), arg0, arg1)
}
