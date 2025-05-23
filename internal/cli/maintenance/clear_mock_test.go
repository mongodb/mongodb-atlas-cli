// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/maintenance (interfaces: Clearer)
//
// Generated by this command:
//
//	mockgen -typed -destination=clear_mock_test.go -package=maintenance . Clearer
//

// Package maintenance is a generated GoMock package.
package maintenance

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockClearer is a mock of Clearer interface.
type MockClearer struct {
	ctrl     *gomock.Controller
	recorder *MockClearerMockRecorder
	isgomock struct{}
}

// MockClearerMockRecorder is the mock recorder for MockClearer.
type MockClearerMockRecorder struct {
	mock *MockClearer
}

// NewMockClearer creates a new mock instance.
func NewMockClearer(ctrl *gomock.Controller) *MockClearer {
	mock := &MockClearer{ctrl: ctrl}
	mock.recorder = &MockClearerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClearer) EXPECT() *MockClearerMockRecorder {
	return m.recorder
}

// ClearMaintenanceWindow mocks base method.
func (m *MockClearer) ClearMaintenanceWindow(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearMaintenanceWindow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearMaintenanceWindow indicates an expected call of ClearMaintenanceWindow.
func (mr *MockClearerMockRecorder) ClearMaintenanceWindow(arg0 any) *MockClearerClearMaintenanceWindowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearMaintenanceWindow", reflect.TypeOf((*MockClearer)(nil).ClearMaintenanceWindow), arg0)
	return &MockClearerClearMaintenanceWindowCall{Call: call}
}

// MockClearerClearMaintenanceWindowCall wrap *gomock.Call
type MockClearerClearMaintenanceWindowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClearerClearMaintenanceWindowCall) Return(arg0 error) *MockClearerClearMaintenanceWindowCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClearerClearMaintenanceWindowCall) Do(f func(string) error) *MockClearerClearMaintenanceWindowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClearerClearMaintenanceWindowCall) DoAndReturn(f func(string) error) *MockClearerClearMaintenanceWindowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
