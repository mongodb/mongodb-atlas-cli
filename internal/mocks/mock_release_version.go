// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/version (interfaces: ReleaseVersionDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v50/github"
	version "github.com/mongodb/mongodb-atlas-cli/internal/version"
)

// MockReleaseVersionDescriber is a mock of ReleaseVersionDescriber interface.
type MockReleaseVersionDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockReleaseVersionDescriberMockRecorder
}

// MockReleaseVersionDescriberMockRecorder is the mock recorder for MockReleaseVersionDescriber.
type MockReleaseVersionDescriberMockRecorder struct {
	mock *MockReleaseVersionDescriber
}

// NewMockReleaseVersionDescriber creates a new mock instance.
func NewMockReleaseVersionDescriber(ctrl *gomock.Controller) *MockReleaseVersionDescriber {
	mock := &MockReleaseVersionDescriber{ctrl: ctrl}
	mock.recorder = &MockReleaseVersionDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReleaseVersionDescriber) EXPECT() *MockReleaseVersionDescriberMockRecorder {
	return m.recorder
}

// LatestWithCriteria mocks base method.
func (m *MockReleaseVersionDescriber) LatestWithCriteria(arg0 int, arg1 version.Criteria) (*github.RepositoryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestWithCriteria", arg0, arg1)
	ret0, _ := ret[0].(*github.RepositoryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestWithCriteria indicates an expected call of LatestWithCriteria.
func (mr *MockReleaseVersionDescriberMockRecorder) LatestWithCriteria(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestWithCriteria", reflect.TypeOf((*MockReleaseVersionDescriber)(nil).LatestWithCriteria), arg0, arg1)
}
