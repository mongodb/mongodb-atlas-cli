// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: AuditingDescriber)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231001002/admin"
)

// MockAuditingDescriber is a mock of AuditingDescriber interface.
type MockAuditingDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockAuditingDescriberMockRecorder
}

// MockAuditingDescriberMockRecorder is the mock recorder for MockAuditingDescriber.
type MockAuditingDescriberMockRecorder struct {
	mock *MockAuditingDescriber
}

// NewMockAuditingDescriber creates a new mock instance.
func NewMockAuditingDescriber(ctrl *gomock.Controller) *MockAuditingDescriber {
	mock := &MockAuditingDescriber{ctrl: ctrl}
	mock.recorder = &MockAuditingDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditingDescriber) EXPECT() *MockAuditingDescriberMockRecorder {
	return m.recorder
}

// Auditing mocks base method.
func (m *MockAuditingDescriber) Auditing(arg0 string) (*admin.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auditing", arg0)
	ret0, _ := ret[0].(*admin.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auditing indicates an expected call of Auditing.
func (mr *MockAuditingDescriberMockRecorder) Auditing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auditing", reflect.TypeOf((*MockAuditingDescriber)(nil).Auditing), arg0)
}
