// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: OnlineArchiveLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockOnlineArchiveLister is a mock of OnlineArchiveLister interface
type MockOnlineArchiveLister struct {
	ctrl     *gomock.Controller
	recorder *MockOnlineArchiveListerMockRecorder
}

// MockOnlineArchiveListerMockRecorder is the mock recorder for MockOnlineArchiveLister
type MockOnlineArchiveListerMockRecorder struct {
	mock *MockOnlineArchiveLister
}

// NewMockOnlineArchiveLister creates a new mock instance
func NewMockOnlineArchiveLister(ctrl *gomock.Controller) *MockOnlineArchiveLister {
	mock := &MockOnlineArchiveLister{ctrl: ctrl}
	mock.recorder = &MockOnlineArchiveListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnlineArchiveLister) EXPECT() *MockOnlineArchiveListerMockRecorder {
	return m.recorder
}

// OnlineArchives mocks base method
func (m *MockOnlineArchiveLister) OnlineArchives(arg0, arg1 string) ([]*mongodbatlas.OnlineArchive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlineArchives", arg0, arg1)
	ret0, _ := ret[0].([]*mongodbatlas.OnlineArchive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnlineArchives indicates an expected call of OnlineArchives
func (mr *MockOnlineArchiveListerMockRecorder) OnlineArchives(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlineArchives", reflect.TypeOf((*MockOnlineArchiveLister)(nil).OnlineArchives), arg0, arg1)
}
