// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/datalakepipelines (interfaces: PipelinesResumer)
//
// Generated by this command:
//
//	mockgen -typed -destination=start_mock_test.go -package=datalakepipelines . PipelinesResumer
//

// Package datalakepipelines is a generated GoMock package.
package datalakepipelines

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockPipelinesResumer is a mock of PipelinesResumer interface.
type MockPipelinesResumer struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesResumerMockRecorder
	isgomock struct{}
}

// MockPipelinesResumerMockRecorder is the mock recorder for MockPipelinesResumer.
type MockPipelinesResumerMockRecorder struct {
	mock *MockPipelinesResumer
}

// NewMockPipelinesResumer creates a new mock instance.
func NewMockPipelinesResumer(ctrl *gomock.Controller) *MockPipelinesResumer {
	mock := &MockPipelinesResumer{ctrl: ctrl}
	mock.recorder = &MockPipelinesResumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesResumer) EXPECT() *MockPipelinesResumerMockRecorder {
	return m.recorder
}

// PipelineResume mocks base method.
func (m *MockPipelinesResumer) PipelineResume(arg0, arg1 string) (*admin.DataLakeIngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineResume", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataLakeIngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineResume indicates an expected call of PipelineResume.
func (mr *MockPipelinesResumerMockRecorder) PipelineResume(arg0, arg1 any) *MockPipelinesResumerPipelineResumeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineResume", reflect.TypeOf((*MockPipelinesResumer)(nil).PipelineResume), arg0, arg1)
	return &MockPipelinesResumerPipelineResumeCall{Call: call}
}

// MockPipelinesResumerPipelineResumeCall wrap *gomock.Call
type MockPipelinesResumerPipelineResumeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesResumerPipelineResumeCall) Return(arg0 *admin.DataLakeIngestionPipeline, arg1 error) *MockPipelinesResumerPipelineResumeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesResumerPipelineResumeCall) Do(f func(string, string) (*admin.DataLakeIngestionPipeline, error)) *MockPipelinesResumerPipelineResumeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesResumerPipelineResumeCall) DoAndReturn(f func(string, string) (*admin.DataLakeIngestionPipeline, error)) *MockPipelinesResumerPipelineResumeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
