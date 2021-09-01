// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: ServiceVersionGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockServiceVersionGetter is a mock of ServiceVersionGetter interface.
type MockServiceVersionGetter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceVersionGetterMockRecorder
}

// MockServiceVersionGetterMockRecorder is the mock recorder for MockServiceVersionGetter.
type MockServiceVersionGetterMockRecorder struct {
	mock *MockServiceVersionGetter
}

// NewMockServiceVersionGetter creates a new mock instance.
func NewMockServiceVersionGetter(ctrl *gomock.Controller) *MockServiceVersionGetter {
	mock := &MockServiceVersionGetter{ctrl: ctrl}
	mock.recorder = &MockServiceVersionGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceVersionGetter) EXPECT() *MockServiceVersionGetterMockRecorder {
	return m.recorder
}

// GetServiceVersion mocks base method.
func (m *MockServiceVersionGetter) GetServiceVersion() (*mongodbatlas.ServiceVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceVersion")
	ret0, _ := ret[0].(*mongodbatlas.ServiceVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceVersion indicates an expected call of GetServiceVersion.
func (mr *MockServiceVersionGetterMockRecorder) GetServiceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceVersion", reflect.TypeOf((*MockServiceVersionGetter)(nil).GetServiceVersion))
}
