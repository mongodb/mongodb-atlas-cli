// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/datalakepipelines (interfaces: PipelinesDeleter)
//
// Generated by this command:
//
//	mockgen -typed -destination=delete_mock_test.go -package=datalakepipelines . PipelinesDeleter
//

// Package datalakepipelines is a generated GoMock package.
package datalakepipelines

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPipelinesDeleter is a mock of PipelinesDeleter interface.
type MockPipelinesDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesDeleterMockRecorder
	isgomock struct{}
}

// MockPipelinesDeleterMockRecorder is the mock recorder for MockPipelinesDeleter.
type MockPipelinesDeleterMockRecorder struct {
	mock *MockPipelinesDeleter
}

// NewMockPipelinesDeleter creates a new mock instance.
func NewMockPipelinesDeleter(ctrl *gomock.Controller) *MockPipelinesDeleter {
	mock := &MockPipelinesDeleter{ctrl: ctrl}
	mock.recorder = &MockPipelinesDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesDeleter) EXPECT() *MockPipelinesDeleterMockRecorder {
	return m.recorder
}

// DeletePipeline mocks base method.
func (m *MockPipelinesDeleter) DeletePipeline(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockPipelinesDeleterMockRecorder) DeletePipeline(arg0, arg1 any) *MockPipelinesDeleterDeletePipelineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockPipelinesDeleter)(nil).DeletePipeline), arg0, arg1)
	return &MockPipelinesDeleterDeletePipelineCall{Call: call}
}

// MockPipelinesDeleterDeletePipelineCall wrap *gomock.Call
type MockPipelinesDeleterDeletePipelineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesDeleterDeletePipelineCall) Return(arg0 error) *MockPipelinesDeleterDeletePipelineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesDeleterDeletePipelineCall) Do(f func(string, string) error) *MockPipelinesDeleterDeletePipelineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesDeleterDeletePipelineCall) DoAndReturn(f func(string, string) error) *MockPipelinesDeleterDeletePipelineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
