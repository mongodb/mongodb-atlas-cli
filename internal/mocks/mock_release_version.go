// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version (interfaces: ReleaseVersionDescriber)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_release_version.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version ReleaseVersionDescriber
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	github "github.com/google/go-github/v61/github"
	version "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	gomock "go.uber.org/mock/gomock"
)

// MockReleaseVersionDescriber is a mock of ReleaseVersionDescriber interface.
type MockReleaseVersionDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockReleaseVersionDescriberMockRecorder
	isgomock struct{}
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
func (m *MockReleaseVersionDescriber) LatestWithCriteria(n int, matchCriteria version.Criteria) (*github.RepositoryRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestWithCriteria", n, matchCriteria)
	ret0, _ := ret[0].(*github.RepositoryRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestWithCriteria indicates an expected call of LatestWithCriteria.
func (mr *MockReleaseVersionDescriberMockRecorder) LatestWithCriteria(n, matchCriteria any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestWithCriteria", reflect.TypeOf((*MockReleaseVersionDescriber)(nil).LatestWithCriteria), n, matchCriteria)
}
