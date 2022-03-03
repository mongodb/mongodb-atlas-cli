// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/latestrelease (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// LoadLatestVersion mocks base method.
func (m *MockStore) LoadLatestVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadLatestVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadLatestVersion indicates an expected call of LoadLatestVersion.
func (mr *MockStoreMockRecorder) LoadLatestVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadLatestVersion", reflect.TypeOf((*MockStore)(nil).LoadLatestVersion))
}

// SaveLatestVersion mocks base method.
func (m *MockStore) SaveLatestVersion(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLatestVersion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLatestVersion indicates an expected call of SaveLatestVersion.
func (mr *MockStoreMockRecorder) SaveLatestVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLatestVersion", reflect.TypeOf((*MockStore)(nil).SaveLatestVersion), arg0)
}
