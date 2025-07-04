// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup/exports/buckets (interfaces: ExportBucketsDescriber)
//
// Generated by this command:
//
//	mockgen -typed -destination=describe_mock_test.go -package=buckets . ExportBucketsDescriber
//

// Package buckets is a generated GoMock package.
package buckets

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockExportBucketsDescriber is a mock of ExportBucketsDescriber interface.
type MockExportBucketsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockExportBucketsDescriberMockRecorder
	isgomock struct{}
}

// MockExportBucketsDescriberMockRecorder is the mock recorder for MockExportBucketsDescriber.
type MockExportBucketsDescriberMockRecorder struct {
	mock *MockExportBucketsDescriber
}

// NewMockExportBucketsDescriber creates a new mock instance.
func NewMockExportBucketsDescriber(ctrl *gomock.Controller) *MockExportBucketsDescriber {
	mock := &MockExportBucketsDescriber{ctrl: ctrl}
	mock.recorder = &MockExportBucketsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportBucketsDescriber) EXPECT() *MockExportBucketsDescriberMockRecorder {
	return m.recorder
}

// DescribeExportBucket mocks base method.
func (m *MockExportBucketsDescriber) DescribeExportBucket(arg0, arg1 string) (*admin.DiskBackupSnapshotExportBucketResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeExportBucket", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotExportBucketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeExportBucket indicates an expected call of DescribeExportBucket.
func (mr *MockExportBucketsDescriberMockRecorder) DescribeExportBucket(arg0, arg1 any) *MockExportBucketsDescriberDescribeExportBucketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeExportBucket", reflect.TypeOf((*MockExportBucketsDescriber)(nil).DescribeExportBucket), arg0, arg1)
	return &MockExportBucketsDescriberDescribeExportBucketCall{Call: call}
}

// MockExportBucketsDescriberDescribeExportBucketCall wrap *gomock.Call
type MockExportBucketsDescriberDescribeExportBucketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportBucketsDescriberDescribeExportBucketCall) Return(arg0 *admin.DiskBackupSnapshotExportBucketResponse, arg1 error) *MockExportBucketsDescriberDescribeExportBucketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportBucketsDescriberDescribeExportBucketCall) Do(f func(string, string) (*admin.DiskBackupSnapshotExportBucketResponse, error)) *MockExportBucketsDescriberDescribeExportBucketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportBucketsDescriberDescribeExportBucketCall) DoAndReturn(f func(string, string) (*admin.DiskBackupSnapshotExportBucketResponse, error)) *MockExportBucketsDescriberDescribeExportBucketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
