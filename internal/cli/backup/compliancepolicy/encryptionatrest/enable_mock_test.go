// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup/compliancepolicy/encryptionatrest (interfaces: CompliancePolicyEncryptionAtRestEnabler)
//
// Generated by this command:
//
//	mockgen -typed -destination=enable_mock_test.go -package=encryptionatrest . CompliancePolicyEncryptionAtRestEnabler
//

// Package encryptionatrest is a generated GoMock package.
package encryptionatrest

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockCompliancePolicyEncryptionAtRestEnabler is a mock of CompliancePolicyEncryptionAtRestEnabler interface.
type MockCompliancePolicyEncryptionAtRestEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder
	isgomock struct{}
}

// MockCompliancePolicyEncryptionAtRestEnablerMockRecorder is the mock recorder for MockCompliancePolicyEncryptionAtRestEnabler.
type MockCompliancePolicyEncryptionAtRestEnablerMockRecorder struct {
	mock *MockCompliancePolicyEncryptionAtRestEnabler
}

// NewMockCompliancePolicyEncryptionAtRestEnabler creates a new mock instance.
func NewMockCompliancePolicyEncryptionAtRestEnabler(ctrl *gomock.Controller) *MockCompliancePolicyEncryptionAtRestEnabler {
	mock := &MockCompliancePolicyEncryptionAtRestEnabler{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyEncryptionAtRestEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyEncryptionAtRestEnabler) EXPECT() *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyEncryptionAtRestEnabler) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder) DescribeCompliancePolicy(projectID any) *MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestEnabler)(nil).DescribeCompliancePolicy), projectID)
	return &MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall{Call: call}
}

// MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall wrap *gomock.Call
type MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall) Return(arg0 *admin.DataProtectionSettings20231001, arg1 error) *MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall) Do(f func(string) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall) DoAndReturn(f func(string) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyEncryptionAtRestEnablerDescribeCompliancePolicyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EnableEncryptionAtRest mocks base method.
func (m *MockCompliancePolicyEncryptionAtRestEnabler) EnableEncryptionAtRest(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableEncryptionAtRest", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableEncryptionAtRest indicates an expected call of EnableEncryptionAtRest.
func (mr *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder) EnableEncryptionAtRest(projectID any) *MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableEncryptionAtRest", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestEnabler)(nil).EnableEncryptionAtRest), projectID)
	return &MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall{Call: call}
}

// MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall wrap *gomock.Call
type MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall) Return(arg0 *admin.DataProtectionSettings20231001, arg1 error) *MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall) Do(f func(string) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall) DoAndReturn(f func(string) (*admin.DataProtectionSettings20231001, error)) *MockCompliancePolicyEncryptionAtRestEnablerEnableEncryptionAtRestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
