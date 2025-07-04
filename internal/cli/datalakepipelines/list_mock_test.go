// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/datalakepipelines (interfaces: PipelinesLister)
//
// Generated by this command:
//
//	mockgen -typed -destination=list_mock_test.go -package=datalakepipelines . PipelinesLister
//

// Package datalakepipelines is a generated GoMock package.
package datalakepipelines

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockPipelinesLister is a mock of PipelinesLister interface.
type MockPipelinesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesListerMockRecorder
	isgomock struct{}
}

// MockPipelinesListerMockRecorder is the mock recorder for MockPipelinesLister.
type MockPipelinesListerMockRecorder struct {
	mock *MockPipelinesLister
}

// NewMockPipelinesLister creates a new mock instance.
func NewMockPipelinesLister(ctrl *gomock.Controller) *MockPipelinesLister {
	mock := &MockPipelinesLister{ctrl: ctrl}
	mock.recorder = &MockPipelinesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesLister) EXPECT() *MockPipelinesListerMockRecorder {
	return m.recorder
}

// Pipelines mocks base method.
func (m *MockPipelinesLister) Pipelines(arg0 string) ([]admin.DataLakeIngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipelines", arg0)
	ret0, _ := ret[0].([]admin.DataLakeIngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipelines indicates an expected call of Pipelines.
func (mr *MockPipelinesListerMockRecorder) Pipelines(arg0 any) *MockPipelinesListerPipelinesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipelines", reflect.TypeOf((*MockPipelinesLister)(nil).Pipelines), arg0)
	return &MockPipelinesListerPipelinesCall{Call: call}
}

// MockPipelinesListerPipelinesCall wrap *gomock.Call
type MockPipelinesListerPipelinesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesListerPipelinesCall) Return(arg0 []admin.DataLakeIngestionPipeline, arg1 error) *MockPipelinesListerPipelinesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesListerPipelinesCall) Do(f func(string) ([]admin.DataLakeIngestionPipeline, error)) *MockPipelinesListerPipelinesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesListerPipelinesCall) DoAndReturn(f func(string) ([]admin.DataLakeIngestionPipeline, error)) *MockPipelinesListerPipelinesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
