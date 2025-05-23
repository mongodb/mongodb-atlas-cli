// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/streams/instance (interfaces: StreamsDeleter)
//
// Generated by this command:
//
//	mockgen -typed -destination=delete_mock_test.go -package=instance . StreamsDeleter
//

// Package instance is a generated GoMock package.
package instance

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStreamsDeleter is a mock of StreamsDeleter interface.
type MockStreamsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsDeleterMockRecorder
	isgomock struct{}
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
func (mr *MockStreamsDeleterMockRecorder) DeleteStream(arg0, arg1 any) *MockStreamsDeleterDeleteStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamsDeleter)(nil).DeleteStream), arg0, arg1)
	return &MockStreamsDeleterDeleteStreamCall{Call: call}
}

// MockStreamsDeleterDeleteStreamCall wrap *gomock.Call
type MockStreamsDeleterDeleteStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamsDeleterDeleteStreamCall) Return(arg0 error) *MockStreamsDeleterDeleteStreamCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamsDeleterDeleteStreamCall) Do(f func(string, string) error) *MockStreamsDeleterDeleteStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamsDeleterDeleteStreamCall) DoAndReturn(f func(string, string) error) *MockStreamsDeleterDeleteStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
