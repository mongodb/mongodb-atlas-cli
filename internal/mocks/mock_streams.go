// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: StreamsLister,StreamsDescriber,StreamsCreator,StreamsDeleter,StreamsUpdater,StreamsDownloader,ConnectionCreator,ConnectionDeleter,ConnectionUpdater,StreamsConnectionDescriber,StreamsConnectionLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115010/admin"
)

// MockStreamsLister is a mock of StreamsLister interface.
type MockStreamsLister struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsListerMockRecorder
}

// MockStreamsListerMockRecorder is the mock recorder for MockStreamsLister.
type MockStreamsListerMockRecorder struct {
	mock *MockStreamsLister
}

// NewMockStreamsLister creates a new mock instance.
func NewMockStreamsLister(ctrl *gomock.Controller) *MockStreamsLister {
	mock := &MockStreamsLister{ctrl: ctrl}
	mock.recorder = &MockStreamsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsLister) EXPECT() *MockStreamsListerMockRecorder {
	return m.recorder
}

// ProjectStreams mocks base method.
func (m *MockStreamsLister) ProjectStreams(arg0 *admin.ListStreamInstancesApiParams) (*admin.PaginatedApiStreamsTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectStreams", arg0)
	ret0, _ := ret[0].(*admin.PaginatedApiStreamsTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectStreams indicates an expected call of ProjectStreams.
func (mr *MockStreamsListerMockRecorder) ProjectStreams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectStreams", reflect.TypeOf((*MockStreamsLister)(nil).ProjectStreams), arg0)
}

// MockStreamsDescriber is a mock of StreamsDescriber interface.
type MockStreamsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsDescriberMockRecorder
}

// MockStreamsDescriberMockRecorder is the mock recorder for MockStreamsDescriber.
type MockStreamsDescriberMockRecorder struct {
	mock *MockStreamsDescriber
}

// NewMockStreamsDescriber creates a new mock instance.
func NewMockStreamsDescriber(ctrl *gomock.Controller) *MockStreamsDescriber {
	mock := &MockStreamsDescriber{ctrl: ctrl}
	mock.recorder = &MockStreamsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsDescriber) EXPECT() *MockStreamsDescriberMockRecorder {
	return m.recorder
}

// AtlasStream mocks base method.
func (m *MockStreamsDescriber) AtlasStream(arg0, arg1 string) (*admin.StreamsTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasStream", arg0, arg1)
	ret0, _ := ret[0].(*admin.StreamsTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasStream indicates an expected call of AtlasStream.
func (mr *MockStreamsDescriberMockRecorder) AtlasStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasStream", reflect.TypeOf((*MockStreamsDescriber)(nil).AtlasStream), arg0, arg1)
}

// MockStreamsCreator is a mock of StreamsCreator interface.
type MockStreamsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsCreatorMockRecorder
}

// MockStreamsCreatorMockRecorder is the mock recorder for MockStreamsCreator.
type MockStreamsCreatorMockRecorder struct {
	mock *MockStreamsCreator
}

// NewMockStreamsCreator creates a new mock instance.
func NewMockStreamsCreator(ctrl *gomock.Controller) *MockStreamsCreator {
	mock := &MockStreamsCreator{ctrl: ctrl}
	mock.recorder = &MockStreamsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsCreator) EXPECT() *MockStreamsCreatorMockRecorder {
	return m.recorder
}

// CreateStream mocks base method.
func (m *MockStreamsCreator) CreateStream(arg0 string, arg1 *admin.StreamsTenant) (*admin.StreamsTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStream", arg0, arg1)
	ret0, _ := ret[0].(*admin.StreamsTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStream indicates an expected call of CreateStream.
func (mr *MockStreamsCreatorMockRecorder) CreateStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStream", reflect.TypeOf((*MockStreamsCreator)(nil).CreateStream), arg0, arg1)
}

// MockStreamsDeleter is a mock of StreamsDeleter interface.
type MockStreamsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsDeleterMockRecorder
}

// MockStreamsDeleterMockRecorder is the mock recorder for MockStreamsDeleter.
type MockStreamsDeleterMockRecorder struct {
	mock *MockStreamsDeleter
}

// NewMockStreamsDeleter creates a new mock instance.
func NewMockStreamsDeleter(ctrl *gomock.Controller) *MockStreamsDeleter {
	mock := &MockStreamsDeleter{ctrl: ctrl}
	mock.recorder = &MockStreamsDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsDeleter) EXPECT() *MockStreamsDeleterMockRecorder {
	return m.recorder
}

// DeleteStream mocks base method.
func (m *MockStreamsDeleter) DeleteStream(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockStreamsDeleterMockRecorder) DeleteStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamsDeleter)(nil).DeleteStream), arg0, arg1)
}

// MockStreamsUpdater is a mock of StreamsUpdater interface.
type MockStreamsUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsUpdaterMockRecorder
}

// MockStreamsUpdaterMockRecorder is the mock recorder for MockStreamsUpdater.
type MockStreamsUpdaterMockRecorder struct {
	mock *MockStreamsUpdater
}

// NewMockStreamsUpdater creates a new mock instance.
func NewMockStreamsUpdater(ctrl *gomock.Controller) *MockStreamsUpdater {
	mock := &MockStreamsUpdater{ctrl: ctrl}
	mock.recorder = &MockStreamsUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsUpdater) EXPECT() *MockStreamsUpdaterMockRecorder {
	return m.recorder
}

// UpdateStream mocks base method.
func (m *MockStreamsUpdater) UpdateStream(arg0, arg1 string, arg2 *admin.StreamsDataProcessRegion) (*admin.StreamsTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.StreamsTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStream indicates an expected call of UpdateStream.
func (mr *MockStreamsUpdaterMockRecorder) UpdateStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStream", reflect.TypeOf((*MockStreamsUpdater)(nil).UpdateStream), arg0, arg1, arg2)
}

// MockStreamsDownloader is a mock of StreamsDownloader interface.
type MockStreamsDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsDownloaderMockRecorder
}

// MockStreamsDownloaderMockRecorder is the mock recorder for MockStreamsDownloader.
type MockStreamsDownloaderMockRecorder struct {
	mock *MockStreamsDownloader
}

// NewMockStreamsDownloader creates a new mock instance.
func NewMockStreamsDownloader(ctrl *gomock.Controller) *MockStreamsDownloader {
	mock := &MockStreamsDownloader{ctrl: ctrl}
	mock.recorder = &MockStreamsDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsDownloader) EXPECT() *MockStreamsDownloaderMockRecorder {
	return m.recorder
}

// DownloadAuditLog mocks base method.
func (m *MockStreamsDownloader) DownloadAuditLog(arg0 *admin.DownloadStreamTenantAuditLogsApiParams) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadAuditLog", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadAuditLog indicates an expected call of DownloadAuditLog.
func (mr *MockStreamsDownloaderMockRecorder) DownloadAuditLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAuditLog", reflect.TypeOf((*MockStreamsDownloader)(nil).DownloadAuditLog), arg0)
}

// MockConnectionCreator is a mock of ConnectionCreator interface.
type MockConnectionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionCreatorMockRecorder
}

// MockConnectionCreatorMockRecorder is the mock recorder for MockConnectionCreator.
type MockConnectionCreatorMockRecorder struct {
	mock *MockConnectionCreator
}

// NewMockConnectionCreator creates a new mock instance.
func NewMockConnectionCreator(ctrl *gomock.Controller) *MockConnectionCreator {
	mock := &MockConnectionCreator{ctrl: ctrl}
	mock.recorder = &MockConnectionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionCreator) EXPECT() *MockConnectionCreatorMockRecorder {
	return m.recorder
}

// CreateConnection mocks base method.
func (m *MockConnectionCreator) CreateConnection(arg0, arg1 string, arg2 *admin.StreamsConnection) (*admin.StreamsConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConnection", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.StreamsConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConnection indicates an expected call of CreateConnection.
func (mr *MockConnectionCreatorMockRecorder) CreateConnection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConnection", reflect.TypeOf((*MockConnectionCreator)(nil).CreateConnection), arg0, arg1, arg2)
}

// MockConnectionDeleter is a mock of ConnectionDeleter interface.
type MockConnectionDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionDeleterMockRecorder
}

// MockConnectionDeleterMockRecorder is the mock recorder for MockConnectionDeleter.
type MockConnectionDeleterMockRecorder struct {
	mock *MockConnectionDeleter
}

// NewMockConnectionDeleter creates a new mock instance.
func NewMockConnectionDeleter(ctrl *gomock.Controller) *MockConnectionDeleter {
	mock := &MockConnectionDeleter{ctrl: ctrl}
	mock.recorder = &MockConnectionDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionDeleter) EXPECT() *MockConnectionDeleterMockRecorder {
	return m.recorder
}

// DeleteConnection mocks base method.
func (m *MockConnectionDeleter) DeleteConnection(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConnection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConnection indicates an expected call of DeleteConnection.
func (mr *MockConnectionDeleterMockRecorder) DeleteConnection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConnection", reflect.TypeOf((*MockConnectionDeleter)(nil).DeleteConnection), arg0, arg1, arg2)
}

// MockConnectionUpdater is a mock of ConnectionUpdater interface.
type MockConnectionUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionUpdaterMockRecorder
}

// MockConnectionUpdaterMockRecorder is the mock recorder for MockConnectionUpdater.
type MockConnectionUpdaterMockRecorder struct {
	mock *MockConnectionUpdater
}

// NewMockConnectionUpdater creates a new mock instance.
func NewMockConnectionUpdater(ctrl *gomock.Controller) *MockConnectionUpdater {
	mock := &MockConnectionUpdater{ctrl: ctrl}
	mock.recorder = &MockConnectionUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionUpdater) EXPECT() *MockConnectionUpdaterMockRecorder {
	return m.recorder
}

// UpdateConnection mocks base method.
func (m *MockConnectionUpdater) UpdateConnection(arg0, arg1, arg2 string, arg3 *admin.StreamsConnection) (*admin.StreamsConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConnection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.StreamsConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConnection indicates an expected call of UpdateConnection.
func (mr *MockConnectionUpdaterMockRecorder) UpdateConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnection", reflect.TypeOf((*MockConnectionUpdater)(nil).UpdateConnection), arg0, arg1, arg2, arg3)
}

// MockStreamsConnectionDescriber is a mock of StreamsConnectionDescriber interface.
type MockStreamsConnectionDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsConnectionDescriberMockRecorder
}

// MockStreamsConnectionDescriberMockRecorder is the mock recorder for MockStreamsConnectionDescriber.
type MockStreamsConnectionDescriberMockRecorder struct {
	mock *MockStreamsConnectionDescriber
}

// NewMockStreamsConnectionDescriber creates a new mock instance.
func NewMockStreamsConnectionDescriber(ctrl *gomock.Controller) *MockStreamsConnectionDescriber {
	mock := &MockStreamsConnectionDescriber{ctrl: ctrl}
	mock.recorder = &MockStreamsConnectionDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsConnectionDescriber) EXPECT() *MockStreamsConnectionDescriberMockRecorder {
	return m.recorder
}

// StreamConnection mocks base method.
func (m *MockStreamsConnectionDescriber) StreamConnection(arg0, arg1, arg2 string) (*admin.StreamsConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamConnection", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.StreamsConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamConnection indicates an expected call of StreamConnection.
func (mr *MockStreamsConnectionDescriberMockRecorder) StreamConnection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamConnection", reflect.TypeOf((*MockStreamsConnectionDescriber)(nil).StreamConnection), arg0, arg1, arg2)
}

// MockStreamsConnectionLister is a mock of StreamsConnectionLister interface.
type MockStreamsConnectionLister struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsConnectionListerMockRecorder
}

// MockStreamsConnectionListerMockRecorder is the mock recorder for MockStreamsConnectionLister.
type MockStreamsConnectionListerMockRecorder struct {
	mock *MockStreamsConnectionLister
}

// NewMockStreamsConnectionLister creates a new mock instance.
func NewMockStreamsConnectionLister(ctrl *gomock.Controller) *MockStreamsConnectionLister {
	mock := &MockStreamsConnectionLister{ctrl: ctrl}
	mock.recorder = &MockStreamsConnectionListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsConnectionLister) EXPECT() *MockStreamsConnectionListerMockRecorder {
	return m.recorder
}

// StreamsConnections mocks base method.
func (m *MockStreamsConnectionLister) StreamsConnections(arg0, arg1 string) (*admin.PaginatedApiStreamsConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamsConnections", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedApiStreamsConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamsConnections indicates an expected call of StreamsConnections.
func (mr *MockStreamsConnectionListerMockRecorder) StreamsConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamsConnections", reflect.TypeOf((*MockStreamsConnectionLister)(nil).StreamsConnections), arg0, arg1)
}
