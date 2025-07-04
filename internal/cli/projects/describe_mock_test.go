// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/projects (interfaces: ProjectDescriber)
//
// Generated by this command:
//
//	mockgen -typed -destination=describe_mock_test.go -package=projects . ProjectDescriber
//

// Package projects is a generated GoMock package.
package projects

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectDescriber is a mock of ProjectDescriber interface.
type MockProjectDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockProjectDescriberMockRecorder
	isgomock struct{}
}

// MockProjectDescriberMockRecorder is the mock recorder for MockProjectDescriber.
type MockProjectDescriberMockRecorder struct {
	mock *MockProjectDescriber
}

// NewMockProjectDescriber creates a new mock instance.
func NewMockProjectDescriber(ctrl *gomock.Controller) *MockProjectDescriber {
	mock := &MockProjectDescriber{ctrl: ctrl}
	mock.recorder = &MockProjectDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectDescriber) EXPECT() *MockProjectDescriberMockRecorder {
	return m.recorder
}

// Project mocks base method.
func (m *MockProjectDescriber) Project(arg0 string) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockProjectDescriberMockRecorder) Project(arg0 any) *MockProjectDescriberProjectCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockProjectDescriber)(nil).Project), arg0)
	return &MockProjectDescriberProjectCall{Call: call}
}

// MockProjectDescriberProjectCall wrap *gomock.Call
type MockProjectDescriberProjectCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectDescriberProjectCall) Return(arg0 *admin.Group, arg1 error) *MockProjectDescriberProjectCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectDescriberProjectCall) Do(f func(string) (*admin.Group, error)) *MockProjectDescriberProjectCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectDescriberProjectCall) DoAndReturn(f func(string) (*admin.Group, error)) *MockProjectDescriberProjectCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
