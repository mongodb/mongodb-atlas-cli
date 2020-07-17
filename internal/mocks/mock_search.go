// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: SearchIndexLister,SearchIndexCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockSearchIndexLister is a mock of SearchIndexLister interface
type MockSearchIndexLister struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexListerMockRecorder
}

// MockSearchIndexListerMockRecorder is the mock recorder for MockSearchIndexLister
type MockSearchIndexListerMockRecorder struct {
	mock *MockSearchIndexLister
}

// NewMockSearchIndexLister creates a new mock instance
func NewMockSearchIndexLister(ctrl *gomock.Controller) *MockSearchIndexLister {
	mock := &MockSearchIndexLister{ctrl: ctrl}
	mock.recorder = &MockSearchIndexListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchIndexLister) EXPECT() *MockSearchIndexListerMockRecorder {
	return m.recorder
}

// SearchIndexes mocks base method
func (m *MockSearchIndexLister) SearchIndexes(arg0, arg1, arg2, arg3 string, arg4 *mongodbatlas.ListOptions) ([]*mongodbatlas.SearchIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIndexes", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*mongodbatlas.SearchIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIndexes indicates an expected call of SearchIndexes
func (mr *MockSearchIndexListerMockRecorder) SearchIndexes(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIndexes", reflect.TypeOf((*MockSearchIndexLister)(nil).SearchIndexes), arg0, arg1, arg2, arg3, arg4)
}

// MockSearchIndexCreator is a mock of SearchIndexCreator interface
type MockSearchIndexCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexCreatorMockRecorder
}

// MockSearchIndexCreatorMockRecorder is the mock recorder for MockSearchIndexCreator
type MockSearchIndexCreatorMockRecorder struct {
	mock *MockSearchIndexCreator
}

// NewMockSearchIndexCreator creates a new mock instance
func NewMockSearchIndexCreator(ctrl *gomock.Controller) *MockSearchIndexCreator {
	mock := &MockSearchIndexCreator{ctrl: ctrl}
	mock.recorder = &MockSearchIndexCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchIndexCreator) EXPECT() *MockSearchIndexCreatorMockRecorder {
	return m.recorder
}

// CreateSearchIndexes mocks base method
func (m *MockSearchIndexCreator) CreateSearchIndexes(arg0, arg1 string, arg2 *mongodbatlas.SearchIndex) (*mongodbatlas.SearchIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSearchIndexes", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.SearchIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSearchIndexes indicates an expected call of CreateSearchIndexes
func (mr *MockSearchIndexCreatorMockRecorder) CreateSearchIndexes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSearchIndexes", reflect.TypeOf((*MockSearchIndexCreator)(nil).CreateSearchIndexes), arg0, arg1, arg2)
}
