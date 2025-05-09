// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: PipelineDatasetDeleter)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_data_lake_pipelines_datasets.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store PipelineDatasetDeleter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPipelineDatasetDeleter is a mock of PipelineDatasetDeleter interface.
type MockPipelineDatasetDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineDatasetDeleterMockRecorder
	isgomock struct{}
}

// MockPipelineDatasetDeleterMockRecorder is the mock recorder for MockPipelineDatasetDeleter.
type MockPipelineDatasetDeleterMockRecorder struct {
	mock *MockPipelineDatasetDeleter
}

// NewMockPipelineDatasetDeleter creates a new mock instance.
func NewMockPipelineDatasetDeleter(ctrl *gomock.Controller) *MockPipelineDatasetDeleter {
	mock := &MockPipelineDatasetDeleter{ctrl: ctrl}
	mock.recorder = &MockPipelineDatasetDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineDatasetDeleter) EXPECT() *MockPipelineDatasetDeleterMockRecorder {
	return m.recorder
}

// DeletePipelineDataset mocks base method.
func (m *MockPipelineDatasetDeleter) DeletePipelineDataset(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipelineDataset", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipelineDataset indicates an expected call of DeletePipelineDataset.
func (mr *MockPipelineDatasetDeleterMockRecorder) DeletePipelineDataset(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipelineDataset", reflect.TypeOf((*MockPipelineDatasetDeleter)(nil).DeletePipelineDataset), arg0, arg1, arg2)
}
