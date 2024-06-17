// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container (interfaces: Engine)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	container "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// ContainerInspect mocks base method.
func (m *MockEngine) ContainerInspect(arg0 context.Context, arg1 ...string) ([]*container.ContainerInspectData, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerInspect", varargs...)
	ret0, _ := ret[0].([]*container.ContainerInspectData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerInspect indicates an expected call of ContainerInspect.
func (mr *MockEngineMockRecorder) ContainerInspect(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerInspect", reflect.TypeOf((*MockEngine)(nil).ContainerInspect), varargs...)
}

// ContainerList mocks base method.
func (m *MockEngine) ContainerList(arg0 context.Context, arg1 ...string) ([]container.Container, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerList", varargs...)
	ret0, _ := ret[0].([]container.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerList indicates an expected call of ContainerList.
func (mr *MockEngineMockRecorder) ContainerList(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerList", reflect.TypeOf((*MockEngine)(nil).ContainerList), varargs...)
}

// ContainerLogs mocks base method.
func (m *MockEngine) ContainerLogs(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerLogs", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerLogs indicates an expected call of ContainerLogs.
func (mr *MockEngineMockRecorder) ContainerLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerLogs", reflect.TypeOf((*MockEngine)(nil).ContainerLogs), arg0, arg1)
}

// ContainerRm mocks base method.
func (m *MockEngine) ContainerRm(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerRm", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerRm indicates an expected call of ContainerRm.
func (mr *MockEngineMockRecorder) ContainerRm(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerRm", reflect.TypeOf((*MockEngine)(nil).ContainerRm), varargs...)
}

// ContainerRun mocks base method.
func (m *MockEngine) ContainerRun(arg0 context.Context, arg1 string, arg2 *container.RunFlags) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerRun", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerRun indicates an expected call of ContainerRun.
func (mr *MockEngineMockRecorder) ContainerRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerRun", reflect.TypeOf((*MockEngine)(nil).ContainerRun), arg0, arg1, arg2)
}

// ContainerStart mocks base method.
func (m *MockEngine) ContainerStart(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerStart", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerStart indicates an expected call of ContainerStart.
func (mr *MockEngineMockRecorder) ContainerStart(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStart", reflect.TypeOf((*MockEngine)(nil).ContainerStart), varargs...)
}

// ContainerStop mocks base method.
func (m *MockEngine) ContainerStop(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerStop", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerStop indicates an expected call of ContainerStop.
func (mr *MockEngineMockRecorder) ContainerStop(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStop", reflect.TypeOf((*MockEngine)(nil).ContainerStop), varargs...)
}

// ContainerUnpause mocks base method.
func (m *MockEngine) ContainerUnpause(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerUnpause", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerUnpause indicates an expected call of ContainerUnpause.
func (mr *MockEngineMockRecorder) ContainerUnpause(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerUnpause", reflect.TypeOf((*MockEngine)(nil).ContainerUnpause), varargs...)
}

// ImageList mocks base method.
func (m *MockEngine) ImageList(arg0 context.Context, arg1 ...string) ([]container.Image, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImageList", varargs...)
	ret0, _ := ret[0].([]container.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageList indicates an expected call of ImageList.
func (mr *MockEngineMockRecorder) ImageList(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageList", reflect.TypeOf((*MockEngine)(nil).ImageList), varargs...)
}

// ImagePull mocks base method.
func (m *MockEngine) ImagePull(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagePull", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImagePull indicates an expected call of ImagePull.
func (mr *MockEngineMockRecorder) ImagePull(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagePull", reflect.TypeOf((*MockEngine)(nil).ImagePull), arg0, arg1)
}

// Ready mocks base method.
func (m *MockEngine) Ready(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ready indicates an expected call of Ready.
func (mr *MockEngineMockRecorder) Ready(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockEngine)(nil).Ready), arg0)
}
