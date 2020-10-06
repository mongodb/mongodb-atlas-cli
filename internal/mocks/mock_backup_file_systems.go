// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: FileSystemsLister,FileSystemsDescriber,FileSystemsCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
	reflect "reflect"
)

// MockFileSystemsLister is a mock of FileSystemsLister interface
type MockFileSystemsLister struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemsListerMockRecorder
}

// MockFileSystemsListerMockRecorder is the mock recorder for MockFileSystemsLister
type MockFileSystemsListerMockRecorder struct {
	mock *MockFileSystemsLister
}

// NewMockFileSystemsLister creates a new mock instance
func NewMockFileSystemsLister(ctrl *gomock.Controller) *MockFileSystemsLister {
	mock := &MockFileSystemsLister{ctrl: ctrl}
	mock.recorder = &MockFileSystemsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileSystemsLister) EXPECT() *MockFileSystemsListerMockRecorder {
	return m.recorder
}

// ListFileSystems mocks base method
func (m *MockFileSystemsLister) ListFileSystems(arg0 *mongodbatlas.ListOptions) (*opsmngr.FileSystemStoreConfigurations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFileSystems", arg0)
	ret0, _ := ret[0].(*opsmngr.FileSystemStoreConfigurations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFileSystems indicates an expected call of ListFileSystems
func (mr *MockFileSystemsListerMockRecorder) ListFileSystems(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFileSystems", reflect.TypeOf((*MockFileSystemsLister)(nil).ListFileSystems), arg0)
}

// MockFileSystemsDescriber is a mock of FileSystemsDescriber interface
type MockFileSystemsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemsDescriberMockRecorder
}

// MockFileSystemsDescriberMockRecorder is the mock recorder for MockFileSystemsDescriber
type MockFileSystemsDescriberMockRecorder struct {
	mock *MockFileSystemsDescriber
}

// NewMockFileSystemsDescriber creates a new mock instance
func NewMockFileSystemsDescriber(ctrl *gomock.Controller) *MockFileSystemsDescriber {
	mock := &MockFileSystemsDescriber{ctrl: ctrl}
	mock.recorder = &MockFileSystemsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileSystemsDescriber) EXPECT() *MockFileSystemsDescriberMockRecorder {
	return m.recorder
}

// DescribeFileSystem mocks base method
func (m *MockFileSystemsDescriber) DescribeFileSystem(arg0 string) (*opsmngr.FileSystemStoreConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFileSystem", arg0)
	ret0, _ := ret[0].(*opsmngr.FileSystemStoreConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystem indicates an expected call of DescribeFileSystem
func (mr *MockFileSystemsDescriberMockRecorder) DescribeFileSystem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystem", reflect.TypeOf((*MockFileSystemsDescriber)(nil).DescribeFileSystem), arg0)
}

// MockFileSystemsCreator is a mock of FileSystemsCreator interface
type MockFileSystemsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemsCreatorMockRecorder
}

// MockFileSystemsCreatorMockRecorder is the mock recorder for MockFileSystemsCreator
type MockFileSystemsCreatorMockRecorder struct {
	mock *MockFileSystemsCreator
}

// NewMockFileSystemsCreator creates a new mock instance
func NewMockFileSystemsCreator(ctrl *gomock.Controller) *MockFileSystemsCreator {
	mock := &MockFileSystemsCreator{ctrl: ctrl}
	mock.recorder = &MockFileSystemsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileSystemsCreator) EXPECT() *MockFileSystemsCreatorMockRecorder {
	return m.recorder
}

// CreateFileSystems mocks base method
func (m *MockFileSystemsCreator) CreateFileSystems(arg0 *opsmngr.FileSystemStoreConfiguration) (*opsmngr.FileSystemStoreConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFileSystems", arg0)
	ret0, _ := ret[0].(*opsmngr.FileSystemStoreConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFileSystems indicates an expected call of CreateFileSystems
func (mr *MockFileSystemsCreatorMockRecorder) CreateFileSystems(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFileSystems", reflect.TypeOf((*MockFileSystemsCreator)(nil).CreateFileSystems), arg0)
}
