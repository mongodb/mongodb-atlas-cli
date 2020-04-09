// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/process_measurements.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockAtlasProcessMeasurementLister is a mock of AtlasProcessMeasurementLister interface
type MockAtlasProcessMeasurementLister struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasProcessMeasurementListerMockRecorder
}

// MockAtlasProcessMeasurementListerMockRecorder is the mock recorder for MockAtlasProcessMeasurementLister
type MockAtlasProcessMeasurementListerMockRecorder struct {
	mock *MockAtlasProcessMeasurementLister
}

// NewMockAtlasProcessMeasurementLister creates a new mock instance
func NewMockAtlasProcessMeasurementLister(ctrl *gomock.Controller) *MockAtlasProcessMeasurementLister {
	mock := &MockAtlasProcessMeasurementLister{ctrl: ctrl}
	mock.recorder = &MockAtlasProcessMeasurementListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAtlasProcessMeasurementLister) EXPECT() *MockAtlasProcessMeasurementListerMockRecorder {
	return m.recorder
}

// AtlasProcessMeasurements mocks base method
func (m *MockAtlasProcessMeasurementLister) AtlasProcessMeasurements(arg0, arg1 string, arg2 int, arg3 *mongodbatlas.ProcessMeasurementListOptions) (*mongodbatlas.ProcessMeasurements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasProcessMeasurements", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*mongodbatlas.ProcessMeasurements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasProcessMeasurements indicates an expected call of AtlasProcessMeasurements
func (mr *MockAtlasProcessMeasurementListerMockRecorder) AtlasProcessMeasurements(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasProcessMeasurements", reflect.TypeOf((*MockAtlasProcessMeasurementLister)(nil).AtlasProcessMeasurements), arg0, arg1, arg2, arg3)
}

// MockOpsManagerProcessMeasurementLister is a mock of OpsManagerProcessMeasurementLister interface
type MockOpsManagerProcessMeasurementLister struct {
	ctrl     *gomock.Controller
	recorder *MockOpsManagerProcessMeasurementListerMockRecorder
}

// MockOpsManagerProcessMeasurementListerMockRecorder is the mock recorder for MockOpsManagerProcessMeasurementLister
type MockOpsManagerProcessMeasurementListerMockRecorder struct {
	mock *MockOpsManagerProcessMeasurementLister
}

// NewMockOpsManagerProcessMeasurementLister creates a new mock instance
func NewMockOpsManagerProcessMeasurementLister(ctrl *gomock.Controller) *MockOpsManagerProcessMeasurementLister {
	mock := &MockOpsManagerProcessMeasurementLister{ctrl: ctrl}
	mock.recorder = &MockOpsManagerProcessMeasurementListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpsManagerProcessMeasurementLister) EXPECT() *MockOpsManagerProcessMeasurementListerMockRecorder {
	return m.recorder
}

// OpsManagerHostMeasurements mocks base method
func (m *MockOpsManagerProcessMeasurementLister) OpsManagerHostMeasurements(arg0, arg1 string, arg2 *mongodbatlas.ProcessMeasurementListOptions) (*mongodbatlas.ProcessMeasurements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpsManagerHostMeasurements", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.ProcessMeasurements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpsManagerHostMeasurements indicates an expected call of OpsManagerHostMeasurements
func (mr *MockOpsManagerProcessMeasurementListerMockRecorder) OpsManagerHostMeasurements(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpsManagerHostMeasurements", reflect.TypeOf((*MockOpsManagerProcessMeasurementLister)(nil).OpsManagerHostMeasurements), arg0, arg1, arg2)
}
