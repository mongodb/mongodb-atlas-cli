// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: CompliancePolicyDescriber,CompliancePolicyUpdater,CompliancePolicyEncryptionAtRestEnabler,CompliancePolicyEncryptionAtRestDisabler,CompliancePolicyEnabler,CompliancePolicyCopyProtectionEnabler,CompliancePolicyCopyProtectionDisabler,CompliancePolicyPointInTimeRestoresEnabler,CompliancePolicyOnDemandPolicyCreator,CompliancePolicyScheduledPolicyCreator,CompliancePolicyScheduledPolicyDeleter,CompliancePolicyScheduledPolicyUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

// MockCompliancePolicyDescriber is a mock of CompliancePolicyDescriber interface.
type MockCompliancePolicyDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyDescriberMockRecorder
}

// MockCompliancePolicyDescriberMockRecorder is the mock recorder for MockCompliancePolicyDescriber.
type MockCompliancePolicyDescriberMockRecorder struct {
	mock *MockCompliancePolicyDescriber
}

// NewMockCompliancePolicyDescriber creates a new mock instance.
func NewMockCompliancePolicyDescriber(ctrl *gomock.Controller) *MockCompliancePolicyDescriber {
	mock := &MockCompliancePolicyDescriber{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyDescriber) EXPECT() *MockCompliancePolicyDescriberMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyDescriber) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyDescriberMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyDescriber)(nil).DescribeCompliancePolicy), arg0)
}

// MockCompliancePolicyUpdater is a mock of CompliancePolicyUpdater interface.
type MockCompliancePolicyUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyUpdaterMockRecorder
}

// MockCompliancePolicyUpdaterMockRecorder is the mock recorder for MockCompliancePolicyUpdater.
type MockCompliancePolicyUpdaterMockRecorder struct {
	mock *MockCompliancePolicyUpdater
}

// NewMockCompliancePolicyUpdater creates a new mock instance.
func NewMockCompliancePolicyUpdater(ctrl *gomock.Controller) *MockCompliancePolicyUpdater {
	mock := &MockCompliancePolicyUpdater{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyUpdater) EXPECT() *MockCompliancePolicyUpdaterMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyUpdater) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyUpdaterMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyUpdater)(nil).DescribeCompliancePolicy), arg0)
}

// UpdateCompliancePolicy mocks base method.
func (m *MockCompliancePolicyUpdater) UpdateCompliancePolicy(arg0 string, arg1 *admin.DataProtectionSettings20231001) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompliancePolicy", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompliancePolicy indicates an expected call of UpdateCompliancePolicy.
func (mr *MockCompliancePolicyUpdaterMockRecorder) UpdateCompliancePolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyUpdater)(nil).UpdateCompliancePolicy), arg0, arg1)
}

// MockCompliancePolicyEncryptionAtRestEnabler is a mock of CompliancePolicyEncryptionAtRestEnabler interface.
type MockCompliancePolicyEncryptionAtRestEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder
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
func (m *MockCompliancePolicyEncryptionAtRestEnabler) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestEnabler)(nil).DescribeCompliancePolicy), arg0)
}

// EnableEncryptionAtRest mocks base method.
func (m *MockCompliancePolicyEncryptionAtRestEnabler) EnableEncryptionAtRest(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableEncryptionAtRest", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableEncryptionAtRest indicates an expected call of EnableEncryptionAtRest.
func (mr *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder) EnableEncryptionAtRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableEncryptionAtRest", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestEnabler)(nil).EnableEncryptionAtRest), arg0)
}

// MockCompliancePolicyEncryptionAtRestDisabler is a mock of CompliancePolicyEncryptionAtRestDisabler interface.
type MockCompliancePolicyEncryptionAtRestDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyEncryptionAtRestDisablerMockRecorder
}

// MockCompliancePolicyEncryptionAtRestDisablerMockRecorder is the mock recorder for MockCompliancePolicyEncryptionAtRestDisabler.
type MockCompliancePolicyEncryptionAtRestDisablerMockRecorder struct {
	mock *MockCompliancePolicyEncryptionAtRestDisabler
}

// NewMockCompliancePolicyEncryptionAtRestDisabler creates a new mock instance.
func NewMockCompliancePolicyEncryptionAtRestDisabler(ctrl *gomock.Controller) *MockCompliancePolicyEncryptionAtRestDisabler {
	mock := &MockCompliancePolicyEncryptionAtRestDisabler{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyEncryptionAtRestDisablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyEncryptionAtRestDisabler) EXPECT() *MockCompliancePolicyEncryptionAtRestDisablerMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyEncryptionAtRestDisabler) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyEncryptionAtRestDisablerMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestDisabler)(nil).DescribeCompliancePolicy), arg0)
}

// DisableEncryptionAtRest mocks base method.
func (m *MockCompliancePolicyEncryptionAtRestDisabler) DisableEncryptionAtRest(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableEncryptionAtRest", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableEncryptionAtRest indicates an expected call of DisableEncryptionAtRest.
func (mr *MockCompliancePolicyEncryptionAtRestDisablerMockRecorder) DisableEncryptionAtRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableEncryptionAtRest", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestDisabler)(nil).DisableEncryptionAtRest), arg0)
}

// MockCompliancePolicyEnabler is a mock of CompliancePolicyEnabler interface.
type MockCompliancePolicyEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyEnablerMockRecorder
}

// MockCompliancePolicyEnablerMockRecorder is the mock recorder for MockCompliancePolicyEnabler.
type MockCompliancePolicyEnablerMockRecorder struct {
	mock *MockCompliancePolicyEnabler
}

// NewMockCompliancePolicyEnabler creates a new mock instance.
func NewMockCompliancePolicyEnabler(ctrl *gomock.Controller) *MockCompliancePolicyEnabler {
	mock := &MockCompliancePolicyEnabler{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyEnabler) EXPECT() *MockCompliancePolicyEnablerMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyEnabler) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyEnablerMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEnabler)(nil).DescribeCompliancePolicy), arg0)
}

// EnableCompliancePolicy mocks base method.
func (m *MockCompliancePolicyEnabler) EnableCompliancePolicy(arg0, arg1, arg2, arg3 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableCompliancePolicy", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableCompliancePolicy indicates an expected call of EnableCompliancePolicy.
func (mr *MockCompliancePolicyEnablerMockRecorder) EnableCompliancePolicy(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEnabler)(nil).EnableCompliancePolicy), arg0, arg1, arg2, arg3)
}

// MockCompliancePolicyCopyProtectionEnabler is a mock of CompliancePolicyCopyProtectionEnabler interface.
type MockCompliancePolicyCopyProtectionEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyCopyProtectionEnablerMockRecorder
}

// MockCompliancePolicyCopyProtectionEnablerMockRecorder is the mock recorder for MockCompliancePolicyCopyProtectionEnabler.
type MockCompliancePolicyCopyProtectionEnablerMockRecorder struct {
	mock *MockCompliancePolicyCopyProtectionEnabler
}

// NewMockCompliancePolicyCopyProtectionEnabler creates a new mock instance.
func NewMockCompliancePolicyCopyProtectionEnabler(ctrl *gomock.Controller) *MockCompliancePolicyCopyProtectionEnabler {
	mock := &MockCompliancePolicyCopyProtectionEnabler{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyCopyProtectionEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyCopyProtectionEnabler) EXPECT() *MockCompliancePolicyCopyProtectionEnablerMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyCopyProtectionEnabler) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyCopyProtectionEnablerMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyCopyProtectionEnabler)(nil).DescribeCompliancePolicy), arg0)
}

// EnableCopyProtection mocks base method.
func (m *MockCompliancePolicyCopyProtectionEnabler) EnableCopyProtection(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableCopyProtection", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableCopyProtection indicates an expected call of EnableCopyProtection.
func (mr *MockCompliancePolicyCopyProtectionEnablerMockRecorder) EnableCopyProtection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCopyProtection", reflect.TypeOf((*MockCompliancePolicyCopyProtectionEnabler)(nil).EnableCopyProtection), arg0)
}

// MockCompliancePolicyCopyProtectionDisabler is a mock of CompliancePolicyCopyProtectionDisabler interface.
type MockCompliancePolicyCopyProtectionDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyCopyProtectionDisablerMockRecorder
}

// MockCompliancePolicyCopyProtectionDisablerMockRecorder is the mock recorder for MockCompliancePolicyCopyProtectionDisabler.
type MockCompliancePolicyCopyProtectionDisablerMockRecorder struct {
	mock *MockCompliancePolicyCopyProtectionDisabler
}

// NewMockCompliancePolicyCopyProtectionDisabler creates a new mock instance.
func NewMockCompliancePolicyCopyProtectionDisabler(ctrl *gomock.Controller) *MockCompliancePolicyCopyProtectionDisabler {
	mock := &MockCompliancePolicyCopyProtectionDisabler{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyCopyProtectionDisablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyCopyProtectionDisabler) EXPECT() *MockCompliancePolicyCopyProtectionDisablerMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyCopyProtectionDisabler) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyCopyProtectionDisablerMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyCopyProtectionDisabler)(nil).DescribeCompliancePolicy), arg0)
}

// DisableCopyProtection mocks base method.
func (m *MockCompliancePolicyCopyProtectionDisabler) DisableCopyProtection(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableCopyProtection", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableCopyProtection indicates an expected call of DisableCopyProtection.
func (mr *MockCompliancePolicyCopyProtectionDisablerMockRecorder) DisableCopyProtection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableCopyProtection", reflect.TypeOf((*MockCompliancePolicyCopyProtectionDisabler)(nil).DisableCopyProtection), arg0)
}

// MockCompliancePolicyPointInTimeRestoresEnabler is a mock of CompliancePolicyPointInTimeRestoresEnabler interface.
type MockCompliancePolicyPointInTimeRestoresEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder
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
func (m *MockCompliancePolicyPointInTimeRestoresEnabler) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyPointInTimeRestoresEnabler)(nil).DescribeCompliancePolicy), arg0)
}

// EnablePointInTimeRestore mocks base method.
func (m *MockCompliancePolicyPointInTimeRestoresEnabler) EnablePointInTimeRestore(arg0 string, arg1 int) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePointInTimeRestore", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnablePointInTimeRestore indicates an expected call of EnablePointInTimeRestore.
func (mr *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder) EnablePointInTimeRestore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePointInTimeRestore", reflect.TypeOf((*MockCompliancePolicyPointInTimeRestoresEnabler)(nil).EnablePointInTimeRestore), arg0, arg1)
}

// MockCompliancePolicyOnDemandPolicyCreator is a mock of CompliancePolicyOnDemandPolicyCreator interface.
type MockCompliancePolicyOnDemandPolicyCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyOnDemandPolicyCreatorMockRecorder
}

// MockCompliancePolicyOnDemandPolicyCreatorMockRecorder is the mock recorder for MockCompliancePolicyOnDemandPolicyCreator.
type MockCompliancePolicyOnDemandPolicyCreatorMockRecorder struct {
	mock *MockCompliancePolicyOnDemandPolicyCreator
}

// NewMockCompliancePolicyOnDemandPolicyCreator creates a new mock instance.
func NewMockCompliancePolicyOnDemandPolicyCreator(ctrl *gomock.Controller) *MockCompliancePolicyOnDemandPolicyCreator {
	mock := &MockCompliancePolicyOnDemandPolicyCreator{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyOnDemandPolicyCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyOnDemandPolicyCreator) EXPECT() *MockCompliancePolicyOnDemandPolicyCreatorMockRecorder {
	return m.recorder
}

// CreateOnDemandPolicy mocks base method.
func (m *MockCompliancePolicyOnDemandPolicyCreator) CreateOnDemandPolicy(arg0 string, arg1 *admin.BackupComplianceOnDemandPolicyItem) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOnDemandPolicy", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOnDemandPolicy indicates an expected call of CreateOnDemandPolicy.
func (mr *MockCompliancePolicyOnDemandPolicyCreatorMockRecorder) CreateOnDemandPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOnDemandPolicy", reflect.TypeOf((*MockCompliancePolicyOnDemandPolicyCreator)(nil).CreateOnDemandPolicy), arg0, arg1)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyOnDemandPolicyCreator) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyOnDemandPolicyCreatorMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyOnDemandPolicyCreator)(nil).DescribeCompliancePolicy), arg0)
}

// MockCompliancePolicyScheduledPolicyCreator is a mock of CompliancePolicyScheduledPolicyCreator interface.
type MockCompliancePolicyScheduledPolicyCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyScheduledPolicyCreatorMockRecorder
}

// MockCompliancePolicyScheduledPolicyCreatorMockRecorder is the mock recorder for MockCompliancePolicyScheduledPolicyCreator.
type MockCompliancePolicyScheduledPolicyCreatorMockRecorder struct {
	mock *MockCompliancePolicyScheduledPolicyCreator
}

// NewMockCompliancePolicyScheduledPolicyCreator creates a new mock instance.
func NewMockCompliancePolicyScheduledPolicyCreator(ctrl *gomock.Controller) *MockCompliancePolicyScheduledPolicyCreator {
	mock := &MockCompliancePolicyScheduledPolicyCreator{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyScheduledPolicyCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyScheduledPolicyCreator) EXPECT() *MockCompliancePolicyScheduledPolicyCreatorMockRecorder {
	return m.recorder
}

// CreateScheduledPolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyCreator) CreateScheduledPolicy(arg0 string, arg1 *admin.BackupComplianceScheduledPolicyItem) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScheduledPolicy", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScheduledPolicy indicates an expected call of CreateScheduledPolicy.
func (mr *MockCompliancePolicyScheduledPolicyCreatorMockRecorder) CreateScheduledPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduledPolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyCreator)(nil).CreateScheduledPolicy), arg0, arg1)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyCreator) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyScheduledPolicyCreatorMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyCreator)(nil).DescribeCompliancePolicy), arg0)
}

// MockCompliancePolicyScheduledPolicyDeleter is a mock of CompliancePolicyScheduledPolicyDeleter interface.
type MockCompliancePolicyScheduledPolicyDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyScheduledPolicyDeleterMockRecorder
}

// MockCompliancePolicyScheduledPolicyDeleterMockRecorder is the mock recorder for MockCompliancePolicyScheduledPolicyDeleter.
type MockCompliancePolicyScheduledPolicyDeleterMockRecorder struct {
	mock *MockCompliancePolicyScheduledPolicyDeleter
}

// NewMockCompliancePolicyScheduledPolicyDeleter creates a new mock instance.
func NewMockCompliancePolicyScheduledPolicyDeleter(ctrl *gomock.Controller) *MockCompliancePolicyScheduledPolicyDeleter {
	mock := &MockCompliancePolicyScheduledPolicyDeleter{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyScheduledPolicyDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyScheduledPolicyDeleter) EXPECT() *MockCompliancePolicyScheduledPolicyDeleterMockRecorder {
	return m.recorder
}

// DeleteScheduledPolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyDeleter) DeleteScheduledPolicy(arg0, arg1 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduledPolicy", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScheduledPolicy indicates an expected call of DeleteScheduledPolicy.
func (mr *MockCompliancePolicyScheduledPolicyDeleterMockRecorder) DeleteScheduledPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduledPolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyDeleter)(nil).DeleteScheduledPolicy), arg0, arg1)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyDeleter) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyScheduledPolicyDeleterMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyDeleter)(nil).DescribeCompliancePolicy), arg0)
}

// MockCompliancePolicyScheduledPolicyUpdater is a mock of CompliancePolicyScheduledPolicyUpdater interface.
type MockCompliancePolicyScheduledPolicyUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyScheduledPolicyUpdaterMockRecorder
}

// MockCompliancePolicyScheduledPolicyUpdaterMockRecorder is the mock recorder for MockCompliancePolicyScheduledPolicyUpdater.
type MockCompliancePolicyScheduledPolicyUpdaterMockRecorder struct {
	mock *MockCompliancePolicyScheduledPolicyUpdater
}

// NewMockCompliancePolicyScheduledPolicyUpdater creates a new mock instance.
func NewMockCompliancePolicyScheduledPolicyUpdater(ctrl *gomock.Controller) *MockCompliancePolicyScheduledPolicyUpdater {
	mock := &MockCompliancePolicyScheduledPolicyUpdater{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyScheduledPolicyUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyScheduledPolicyUpdater) EXPECT() *MockCompliancePolicyScheduledPolicyUpdaterMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyUpdater) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyScheduledPolicyUpdaterMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyUpdater)(nil).DescribeCompliancePolicy), arg0)
}

// UpdateScheduledPolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyUpdater) UpdateScheduledPolicy(arg0 string, arg1 *admin.BackupComplianceScheduledPolicyItem) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduledPolicy", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduledPolicy indicates an expected call of UpdateScheduledPolicy.
func (mr *MockCompliancePolicyScheduledPolicyUpdaterMockRecorder) UpdateScheduledPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduledPolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyUpdater)(nil).UpdateScheduledPolicy), arg0, arg1)
}
