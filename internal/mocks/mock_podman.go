// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	podman "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ContainerHealthStatus mocks base method.
func (m *MockClient) ContainerHealthStatus(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerHealthStatus", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerHealthStatus indicates an expected call of ContainerHealthStatus.
func (mr *MockClientMockRecorder) ContainerHealthStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerHealthStatus", reflect.TypeOf((*MockClient)(nil).ContainerHealthStatus), arg0, arg1)
}

// ContainerInspect mocks base method.
func (m *MockClient) ContainerInspect(arg0 context.Context, arg1 ...string) ([]*podman.InspectContainerData, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerInspect", varargs...)
	ret0, _ := ret[0].([]*podman.InspectContainerData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerInspect indicates an expected call of ContainerInspect.
func (mr *MockClientMockRecorder) ContainerInspect(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerInspect", reflect.TypeOf((*MockClient)(nil).ContainerInspect), varargs...)
}

// ContainerLogs mocks base method.
func (m *MockClient) ContainerLogs(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerLogs", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerLogs indicates an expected call of ContainerLogs.
func (mr *MockClientMockRecorder) ContainerLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerLogs", reflect.TypeOf((*MockClient)(nil).ContainerLogs), arg0, arg1)
}

// ContainerStatus mocks base method.
func (m *MockClient) ContainerStatus(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerStatus", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerStatus indicates an expected call of ContainerStatus.
func (mr *MockClientMockRecorder) ContainerStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStatus", reflect.TypeOf((*MockClient)(nil).ContainerStatus), arg0, arg1)
}

// ContainerUptime mocks base method.
func (m *MockClient) ContainerUptime(arg0 context.Context, arg1 string) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerUptime", arg0, arg1)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerUptime indicates an expected call of ContainerUptime.
func (mr *MockClientMockRecorder) ContainerUptime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerUptime", reflect.TypeOf((*MockClient)(nil).ContainerUptime), arg0, arg1)
}

// ImageHealthCheck mocks base method.
func (m *MockClient) ImageHealthCheck(arg0 context.Context, arg1 string) (*podman.Schema2HealthConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageHealthCheck", arg0, arg1)
	ret0, _ := ret[0].(*podman.Schema2HealthConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageHealthCheck indicates an expected call of ImageHealthCheck.
func (mr *MockClientMockRecorder) ImageHealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageHealthCheck", reflect.TypeOf((*MockClient)(nil).ImageHealthCheck), arg0, arg1)
}

// ListContainers mocks base method.
func (m *MockClient) ListContainers(arg0 context.Context, arg1 string) ([]*podman.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainers", arg0, arg1)
	ret0, _ := ret[0].([]*podman.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockClientMockRecorder) ListContainers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockClient)(nil).ListContainers), arg0, arg1)
}

// ListImages mocks base method.
func (m *MockClient) ListImages(arg0 context.Context, arg1 string) ([]*podman.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].([]*podman.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockClientMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockClient)(nil).ListImages), arg0, arg1)
}

// Logs mocks base method.
func (m *MockClient) Logs(arg0 context.Context) (map[string]interface{}, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockClientMockRecorder) Logs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockClient)(nil).Logs), arg0)
}

// PullImage mocks base method.
func (m *MockClient) PullImage(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullImage indicates an expected call of PullImage.
func (mr *MockClientMockRecorder) PullImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockClient)(nil).PullImage), arg0, arg1)
}

// Ready mocks base method.
func (m *MockClient) Ready(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ready indicates an expected call of Ready.
func (mr *MockClientMockRecorder) Ready(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockClient)(nil).Ready), arg0)
}

// RemoveContainers mocks base method.
func (m *MockClient) RemoveContainers(arg0 context.Context, arg1 ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveContainers", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveContainers indicates an expected call of RemoveContainers.
func (mr *MockClientMockRecorder) RemoveContainers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainers", reflect.TypeOf((*MockClient)(nil).RemoveContainers), varargs...)
}

// RunContainer mocks base method.
func (m *MockClient) RunContainer(arg0 context.Context, arg1 podman.RunContainerOpts) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunContainer", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunContainer indicates an expected call of RunContainer.
func (mr *MockClientMockRecorder) RunContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunContainer", reflect.TypeOf((*MockClient)(nil).RunContainer), arg0, arg1)
}

// RunHealthcheck mocks base method.
func (m *MockClient) RunHealthcheck(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunHealthcheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunHealthcheck indicates an expected call of RunHealthcheck.
func (mr *MockClientMockRecorder) RunHealthcheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunHealthcheck", reflect.TypeOf((*MockClient)(nil).RunHealthcheck), arg0, arg1)
}

// StartContainers mocks base method.
func (m *MockClient) StartContainers(arg0 context.Context, arg1 ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartContainers", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartContainers indicates an expected call of StartContainers.
func (mr *MockClientMockRecorder) StartContainers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainers", reflect.TypeOf((*MockClient)(nil).StartContainers), varargs...)
}

// StopContainers mocks base method.
func (m *MockClient) StopContainers(arg0 context.Context, arg1 ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopContainers", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopContainers indicates an expected call of StopContainers.
func (mr *MockClientMockRecorder) StopContainers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainers", reflect.TypeOf((*MockClient)(nil).StopContainers), varargs...)
}

// UnpauseContainers mocks base method.
func (m *MockClient) UnpauseContainers(arg0 context.Context, arg1 ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnpauseContainers", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpauseContainers indicates an expected call of UnpauseContainers.
func (mr *MockClientMockRecorder) UnpauseContainers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpauseContainers", reflect.TypeOf((*MockClient)(nil).UnpauseContainers), varargs...)
}

// Version mocks base method.
func (m *MockClient) Version(arg0 context.Context) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockClientMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockClient)(nil).Version), arg0)
}
