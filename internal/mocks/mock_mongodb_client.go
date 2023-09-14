// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/mongodbclient (interfaces: MongoDBClient,Database,SearchIndex)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbclient "github.com/mongodb/mongodb-atlas-cli/internal/mongodbclient"
	admin "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

// MockMongoDBClient is a mock of MongoDBClient interface.
type MockMongoDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockMongoDBClientMockRecorder
}

// MockMongoDBClientMockRecorder is the mock recorder for MockMongoDBClient.
type MockMongoDBClientMockRecorder struct {
	mock *MockMongoDBClient
}

// NewMockMongoDBClient creates a new mock instance.
func NewMockMongoDBClient(ctrl *gomock.Controller) *MockMongoDBClient {
	mock := &MockMongoDBClient{ctrl: ctrl}
	mock.recorder = &MockMongoDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMongoDBClient) EXPECT() *MockMongoDBClientMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockMongoDBClient) Connect(arg0 context.Context, arg1 string, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockMongoDBClientMockRecorder) Connect(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockMongoDBClient)(nil).Connect), arg0, arg1, arg2)
}

// Database mocks base method.
func (m *MockMongoDBClient) Database(arg0 string) mongodbclient.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Database", arg0)
	ret0, _ := ret[0].(mongodbclient.Database)
	return ret0
}

// Database indicates an expected call of Database.
func (mr *MockMongoDBClientMockRecorder) Database(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Database", reflect.TypeOf((*MockMongoDBClient)(nil).Database), arg0)
}

// Disconnect mocks base method.
func (m *MockMongoDBClient) Disconnect(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect", arg0)
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockMongoDBClientMockRecorder) Disconnect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockMongoDBClient)(nil).Disconnect), arg0)
}

// SearchIndex mocks base method.
func (m *MockMongoDBClient) SearchIndex(arg0 context.Context, arg1 string) (*mongodbclient.SearchIndexDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIndex", arg0, arg1)
	ret0, _ := ret[0].(*mongodbclient.SearchIndexDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIndex indicates an expected call of SearchIndex.
func (mr *MockMongoDBClientMockRecorder) SearchIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIndex", reflect.TypeOf((*MockMongoDBClient)(nil).SearchIndex), arg0, arg1)
}

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateSearchIndex mocks base method.
func (m *MockDatabase) CreateSearchIndex(arg0 context.Context, arg1 string, arg2 *admin.ClusterSearchIndex) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSearchIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSearchIndex indicates an expected call of CreateSearchIndex.
func (mr *MockDatabaseMockRecorder) CreateSearchIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSearchIndex", reflect.TypeOf((*MockDatabase)(nil).CreateSearchIndex), arg0, arg1, arg2)
}

// InitiateReplicaSet mocks base method.
func (m *MockDatabase) InitiateReplicaSet(arg0 context.Context, arg1, arg2 string, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiateReplicaSet", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitiateReplicaSet indicates an expected call of InitiateReplicaSet.
func (mr *MockDatabaseMockRecorder) InitiateReplicaSet(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiateReplicaSet", reflect.TypeOf((*MockDatabase)(nil).InitiateReplicaSet), arg0, arg1, arg2, arg3, arg4)
}

// InsertOne mocks base method.
func (m *MockDatabase) InsertOne(arg0 context.Context, arg1 string, arg2 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", arg0, arg1, arg2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockDatabaseMockRecorder) InsertOne(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockDatabase)(nil).InsertOne), arg0, arg1, arg2)
}

// RunCommand mocks base method.
func (m *MockDatabase) RunCommand(arg0 context.Context, arg1 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCommand", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCommand indicates an expected call of RunCommand.
func (mr *MockDatabaseMockRecorder) RunCommand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockDatabase)(nil).RunCommand), arg0, arg1)
}

// SearchIndex mocks base method.
func (m *MockDatabase) SearchIndex(arg0 context.Context, arg1 string) (*mongodbclient.SearchIndexDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIndex", arg0, arg1)
	ret0, _ := ret[0].(*mongodbclient.SearchIndexDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIndex indicates an expected call of SearchIndex.
func (mr *MockDatabaseMockRecorder) SearchIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIndex", reflect.TypeOf((*MockDatabase)(nil).SearchIndex), arg0, arg1)
}

// MockSearchIndex is a mock of SearchIndex interface.
type MockSearchIndex struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexMockRecorder
}

// MockSearchIndexMockRecorder is the mock recorder for MockSearchIndex.
type MockSearchIndexMockRecorder struct {
	mock *MockSearchIndex
}

// NewMockSearchIndex creates a new mock instance.
func NewMockSearchIndex(ctrl *gomock.Controller) *MockSearchIndex {
	mock := &MockSearchIndex{ctrl: ctrl}
	mock.recorder = &MockSearchIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchIndex) EXPECT() *MockSearchIndexMockRecorder {
	return m.recorder
}

// CreateSearchIndex mocks base method.
func (m *MockSearchIndex) CreateSearchIndex(arg0 context.Context, arg1 string, arg2 *admin.ClusterSearchIndex) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSearchIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSearchIndex indicates an expected call of CreateSearchIndex.
func (mr *MockSearchIndexMockRecorder) CreateSearchIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSearchIndex", reflect.TypeOf((*MockSearchIndex)(nil).CreateSearchIndex), arg0, arg1, arg2)
}

// SearchIndex mocks base method.
func (m *MockSearchIndex) SearchIndex(arg0 context.Context, arg1 string) (*mongodbclient.SearchIndexDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIndex", arg0, arg1)
	ret0, _ := ret[0].(*mongodbclient.SearchIndexDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIndex indicates an expected call of SearchIndex.
func (mr *MockSearchIndexMockRecorder) SearchIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIndex", reflect.TypeOf((*MockSearchIndex)(nil).SearchIndex), arg0, arg1)
}
