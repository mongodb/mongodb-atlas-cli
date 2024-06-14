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
