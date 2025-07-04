// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options (interfaces: DeploymentTelemetry)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/mock_deployment_opts_telemetry.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options DeploymentTelemetry
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDeploymentTelemetry is a mock of DeploymentTelemetry interface.
type MockDeploymentTelemetry struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentTelemetryMockRecorder
	isgomock struct{}
}

// MockDeploymentTelemetryMockRecorder is the mock recorder for MockDeploymentTelemetry.
type MockDeploymentTelemetryMockRecorder struct {
	mock *MockDeploymentTelemetry
}

// NewMockDeploymentTelemetry creates a new mock instance.
func NewMockDeploymentTelemetry(ctrl *gomock.Controller) *MockDeploymentTelemetry {
	mock := &MockDeploymentTelemetry{ctrl: ctrl}
	mock.recorder = &MockDeploymentTelemetryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentTelemetry) EXPECT() *MockDeploymentTelemetryMockRecorder {
	return m.recorder
}

// AppendClusterWideScalingMode mocks base method.
func (m *MockDeploymentTelemetry) AppendClusterWideScalingMode() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendClusterWideScalingMode")
}

// AppendClusterWideScalingMode indicates an expected call of AppendClusterWideScalingMode.
func (mr *MockDeploymentTelemetryMockRecorder) AppendClusterWideScalingMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendClusterWideScalingMode", reflect.TypeOf((*MockDeploymentTelemetry)(nil).AppendClusterWideScalingMode))
}

// AppendDeploymentType mocks base method.
func (m *MockDeploymentTelemetry) AppendDeploymentType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendDeploymentType")
}

// AppendDeploymentType indicates an expected call of AppendDeploymentType.
func (mr *MockDeploymentTelemetryMockRecorder) AppendDeploymentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendDeploymentType", reflect.TypeOf((*MockDeploymentTelemetry)(nil).AppendDeploymentType))
}

// AppendDeploymentUUID mocks base method.
func (m *MockDeploymentTelemetry) AppendDeploymentUUID() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendDeploymentUUID")
}

// AppendDeploymentUUID indicates an expected call of AppendDeploymentUUID.
func (mr *MockDeploymentTelemetryMockRecorder) AppendDeploymentUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendDeploymentUUID", reflect.TypeOf((*MockDeploymentTelemetry)(nil).AppendDeploymentUUID))
}

// AppendIndependentShardScalingMode mocks base method.
func (m *MockDeploymentTelemetry) AppendIndependentShardScalingMode() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendIndependentShardScalingMode")
}

// AppendIndependentShardScalingMode indicates an expected call of AppendIndependentShardScalingMode.
func (mr *MockDeploymentTelemetryMockRecorder) AppendIndependentShardScalingMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendIndependentShardScalingMode", reflect.TypeOf((*MockDeploymentTelemetry)(nil).AppendIndependentShardScalingMode))
}
