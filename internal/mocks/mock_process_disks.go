// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: ProcessDisksLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231001002/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockProcessDisksLister is a mock of ProcessDisksLister interface.
type MockProcessDisksLister struct {
	ctrl     *gomock.Controller
	recorder *MockProcessDisksListerMockRecorder
}

// MockProcessDisksListerMockRecorder is the mock recorder for MockProcessDisksLister.
type MockProcessDisksListerMockRecorder struct {
	mock *MockProcessDisksLister
}

// NewMockProcessDisksLister creates a new mock instance.
func NewMockProcessDisksLister(ctrl *gomock.Controller) *MockProcessDisksLister {
	mock := &MockProcessDisksLister{ctrl: ctrl}
	mock.recorder = &MockProcessDisksListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessDisksLister) EXPECT() *MockProcessDisksListerMockRecorder {
	return m.recorder
}

// ProcessDisks mocks base method.
func (m *MockProcessDisksLister) ProcessDisks(arg0, arg1 string, arg2 int, arg3 *mongodbatlas.ListOptions) (*admin.PaginatedDiskPartition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDisks", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.PaginatedDiskPartition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessDisks indicates an expected call of ProcessDisks.
func (mr *MockProcessDisksListerMockRecorder) ProcessDisks(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDisks", reflect.TypeOf((*MockProcessDisksLister)(nil).ProcessDisks), arg0, arg1, arg2, arg3)
}
