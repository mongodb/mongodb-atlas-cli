// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/logs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	opsmngr "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	io "io"
	reflect "reflect"
)

// MockLogsDownloader is a mock of LogsDownloader interface
type MockLogsDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockLogsDownloaderMockRecorder
}

// MockLogsDownloaderMockRecorder is the mock recorder for MockLogsDownloader
type MockLogsDownloaderMockRecorder struct {
	mock *MockLogsDownloader
}

// NewMockLogsDownloader creates a new mock instance
func NewMockLogsDownloader(ctrl *gomock.Controller) *MockLogsDownloader {
	mock := &MockLogsDownloader{ctrl: ctrl}
	mock.recorder = &MockLogsDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogsDownloader) EXPECT() *MockLogsDownloaderMockRecorder {
	return m.recorder
}

// DownloadLog mocks base method
func (m *MockLogsDownloader) DownloadLog(arg0, arg1, arg2 string, arg3 io.Writer, arg4 *mongodbatlas.DateRangetOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadLog", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadLog indicates an expected call of DownloadLog
func (mr *MockLogsDownloaderMockRecorder) DownloadLog(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadLog", reflect.TypeOf((*MockLogsDownloader)(nil).DownloadLog), arg0, arg1, arg2, arg3, arg4)
}

// MockLogCollector is a mock of LogCollector interface
type MockLogCollector struct {
	ctrl     *gomock.Controller
	recorder *MockLogCollectorMockRecorder
}

// MockLogCollectorMockRecorder is the mock recorder for MockLogCollector
type MockLogCollectorMockRecorder struct {
	mock *MockLogCollector
}

// NewMockLogCollector creates a new mock instance
func NewMockLogCollector(ctrl *gomock.Controller) *MockLogCollector {
	mock := &MockLogCollector{ctrl: ctrl}
	mock.recorder = &MockLogCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogCollector) EXPECT() *MockLogCollectorMockRecorder {
	return m.recorder
}

// Collect mocks base method
func (m *MockLogCollector) Collect(arg0 string, arg1 *opsmngr.LogCollectionJob) (*opsmngr.LogCollectionJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.LogCollectionJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect
func (mr *MockLogCollectorMockRecorder) Collect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockLogCollector)(nil).Collect), arg0, arg1)
}

// MockLogsLister is a mock of LogsLister interface
type MockLogsLister struct {
	ctrl     *gomock.Controller
	recorder *MockLogsListerMockRecorder
}

// MockLogsListerMockRecorder is the mock recorder for MockLogsLister
type MockLogsListerMockRecorder struct {
	mock *MockLogsLister
}

// NewMockLogsLister creates a new mock instance
func NewMockLogsLister(ctrl *gomock.Controller) *MockLogsLister {
	mock := &MockLogsLister{ctrl: ctrl}
	mock.recorder = &MockLogsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogsLister) EXPECT() *MockLogsListerMockRecorder {
	return m.recorder
}

// ListLogJobs mocks base method
func (m *MockLogsLister) ListLogJobs(arg0 string, arg1 *opsmngr.LogListOptions) (*opsmngr.LogCollectionJobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLogJobs", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.LogCollectionJobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLogJobs indicates an expected call of ListLogJobs
func (mr *MockLogsListerMockRecorder) ListLogJobs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLogJobs", reflect.TypeOf((*MockLogsLister)(nil).ListLogJobs), arg0, arg1)
}
