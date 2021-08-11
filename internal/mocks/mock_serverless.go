// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: ServerlessLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockServerlessLister is a mock of ServerlessLister interface.
type MockServerlessLister struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessListerMockRecorder
}

// MockServerlessListerMockRecorder is the mock recorder for MockServerlessLister.
type MockServerlessListerMockRecorder struct {
	mock *MockServerlessLister
}

// NewMockServerlessLister creates a new mock instance.
func NewMockServerlessLister(ctrl *gomock.Controller) *MockServerlessLister {
	mock := &MockServerlessLister{ctrl: ctrl}
	mock.recorder = &MockServerlessListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessLister) EXPECT() *MockServerlessListerMockRecorder {
	return m.recorder
}

// ListServerlessClusters mocks base method.
func (m *MockServerlessLister) ListServerlessClusters(arg0 string, arg1 *mongodbatlas.ListOptions) (*mongodbatlas.ClustersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServerlessClusters", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.ClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServerlessClusters indicates an expected call of ListServerlessClusters.
func (mr *MockServerlessListerMockRecorder) ListServerlessClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServerlessClusters", reflect.TypeOf((*MockServerlessLister)(nil).ListServerlessClusters), arg0, arg1)
}
