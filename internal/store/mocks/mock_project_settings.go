// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: ProjectSettingsDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115006/admin"
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
func (m *MockProjectSettingsDescriber) ProjectSettings(arg0 string) (*admin.GroupSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*admin.GroupSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectSettings indicates an expected call of ProjectSettings.
func (mr *MockProjectSettingsDescriberMockRecorder) ProjectSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectSettings", reflect.TypeOf((*MockProjectSettingsDescriber)(nil).ProjectSettings), arg0)
}
