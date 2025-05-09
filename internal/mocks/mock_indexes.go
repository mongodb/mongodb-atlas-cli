// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: IndexCreator)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_indexes.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store IndexCreator
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312002/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockIndexCreator is a mock of IndexCreator interface.
type MockIndexCreator struct {
	ctrl     *gomock.Controller
	recorder *MockIndexCreatorMockRecorder
	isgomock struct{}
}

// MockIndexCreatorMockRecorder is the mock recorder for MockIndexCreator.
type MockIndexCreatorMockRecorder struct {
	mock *MockIndexCreator
}

// NewMockIndexCreator creates a new mock instance.
func NewMockIndexCreator(ctrl *gomock.Controller) *MockIndexCreator {
	mock := &MockIndexCreator{ctrl: ctrl}
	mock.recorder = &MockIndexCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexCreator) EXPECT() *MockIndexCreatorMockRecorder {
	return m.recorder
}

// CreateIndex mocks base method.
func (m *MockIndexCreator) CreateIndex(arg0, arg1 string, arg2 *admin.DatabaseRollingIndexRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIndex indicates an expected call of CreateIndex.
func (mr *MockIndexCreatorMockRecorder) CreateIndex(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockIndexCreator)(nil).CreateIndex), arg0, arg1, arg2)
}
