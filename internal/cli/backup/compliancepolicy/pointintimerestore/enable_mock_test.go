// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup/compliancepolicy/pointintimerestore (interfaces: CompliancePolicyPointInTimeRestoresEnabler)
//
// Generated by this command:
//
//	mockgen -typed -destination=enable_mock_test.go -package=pointintimerestore . CompliancePolicyPointInTimeRestoresEnabler
//

// Package pointintimerestore is a generated GoMock package.
package pointintimerestore

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockCompliancePolicyPointInTimeRestoresEnabler is a mock of CompliancePolicyPointInTimeRestoresEnabler interface.
type MockCompliancePolicyPointInTimeRestoresEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder
	isgomock struct{}
}

// MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder is the mock recorder for MockCompliancePolicyPointInTimeRestoresEnabler.
type MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder struct {
	mock *MockCompliancePolicyPointInTimeRestoresEnabler
}

// NewMockCompliancePolicyPointInTimeRestoresEnabler creates a new mock instance.
func NewMockCompliancePolicyPointInTimeRestoresEnabler(ctrl *gomock.Controller) *MockCompliancePolicyPointInTimeRestoresEnabler {
	mock := &MockCompliancePolicyPointInTimeRestoresEnabler{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyPointInTimeRestoresEnabler) EXPECT() *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyPointInTimeRestoresEnabler) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder) DescribeCompliancePolicy(projectID any) *MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyPointInTimeRestoresEnabler)(nil).DescribeCompliancePolicy), projectID)
	return &MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall{Call: call}
}

// MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall wrap *gomock.Call
type MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall) Return(arg0 *admin.DataProtectionSettings20231001, arg1 error) *MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall) Do(f func(string) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall) DoAndReturn(f func(string) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyPointInTimeRestoresEnablerDescribeCompliancePolicyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EnablePointInTimeRestore mocks base method.
func (m *MockCompliancePolicyPointInTimeRestoresEnabler) EnablePointInTimeRestore(projectID string, restoreWindowDays int) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePointInTimeRestore", projectID, restoreWindowDays)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnablePointInTimeRestore indicates an expected call of EnablePointInTimeRestore.
func (mr *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder) EnablePointInTimeRestore(projectID, restoreWindowDays any) *MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePointInTimeRestore", reflect.TypeOf((*MockCompliancePolicyPointInTimeRestoresEnabler)(nil).EnablePointInTimeRestore), projectID, restoreWindowDays)
	return &MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall{Call: call}
}

// MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall wrap *gomock.Call
type MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall) Return(arg0 *admin.DataProtectionSettings20231001, arg1 error) *MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall) Do(f func(string, int) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall) DoAndReturn(f func(string, int) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyPointInTimeRestoresEnablerEnablePointInTimeRestoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
