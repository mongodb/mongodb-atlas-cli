// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/setup (interfaces: ProfileReader)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/mock_setup.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/setup ProfileReader
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProfileReader is a mock of ProfileReader interface.
type MockProfileReader struct {
	ctrl     *gomock.Controller
	recorder *MockProfileReaderMockRecorder
	isgomock struct{}
}

// MockProfileReaderMockRecorder is the mock recorder for MockProfileReader.
type MockProfileReaderMockRecorder struct {
	mock *MockProfileReader
}

// NewMockProfileReader creates a new mock instance.
func NewMockProfileReader(ctrl *gomock.Controller) *MockProfileReader {
	mock := &MockProfileReader{ctrl: ctrl}
	mock.recorder = &MockProfileReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileReader) EXPECT() *MockProfileReaderMockRecorder {
	return m.recorder
}

// OrgID mocks base method.
func (m *MockProfileReader) OrgID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrgID")
	ret0, _ := ret[0].(string)
	return ret0
}

// OrgID indicates an expected call of OrgID.
func (mr *MockProfileReaderMockRecorder) OrgID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrgID", reflect.TypeOf((*MockProfileReader)(nil).OrgID))
}

// ProjectID mocks base method.
func (m *MockProfileReader) ProjectID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProjectID indicates an expected call of ProjectID.
func (mr *MockProfileReaderMockRecorder) ProjectID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectID", reflect.TypeOf((*MockProfileReader)(nil).ProjectID))
}
