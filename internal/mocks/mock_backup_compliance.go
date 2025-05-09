// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: CompliancePolicyDescriber,CompliancePolicyUpdater,CompliancePolicyEncryptionAtRestEnabler,CompliancePolicyEncryptionAtRestDisabler,CompliancePolicyEnabler,CompliancePolicyCopyProtectionEnabler,CompliancePolicyCopyProtectionDisabler,CompliancePolicyPointInTimeRestoresEnabler,CompliancePolicyOnDemandPolicyCreator,CompliancePolicyScheduledPolicyCreator,CompliancePolicyScheduledPolicyDeleter,CompliancePolicyScheduledPolicyUpdater)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_backup_compliance.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store CompliancePolicyDescriber,CompliancePolicyUpdater,CompliancePolicyEncryptionAtRestEnabler,CompliancePolicyEncryptionAtRestDisabler,CompliancePolicyEnabler,CompliancePolicyCopyProtectionEnabler,CompliancePolicyCopyProtectionDisabler,CompliancePolicyPointInTimeRestoresEnabler,CompliancePolicyOnDemandPolicyCreator,CompliancePolicyScheduledPolicyCreator,CompliancePolicyScheduledPolicyDeleter,CompliancePolicyScheduledPolicyUpdater
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312002/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockCompliancePolicyDescriber is a mock of CompliancePolicyDescriber interface.
type MockCompliancePolicyDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyDescriberMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyDescriber) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyDescriberMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyDescriber)(nil).DescribeCompliancePolicy), projectID)
}

// MockCompliancePolicyUpdater is a mock of CompliancePolicyUpdater interface.
type MockCompliancePolicyUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyUpdaterMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyUpdater) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyUpdaterMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyUpdater)(nil).DescribeCompliancePolicy), projectID)
}

// UpdateCompliancePolicy mocks base method.
func (m *MockCompliancePolicyUpdater) UpdateCompliancePolicy(projectID string, opts *admin.DataProtectionSettings20231001) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompliancePolicy", projectID, opts)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompliancePolicy indicates an expected call of UpdateCompliancePolicy.
func (mr *MockCompliancePolicyUpdaterMockRecorder) UpdateCompliancePolicy(projectID, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyUpdater)(nil).UpdateCompliancePolicy), projectID, opts)
}

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
func (mr *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestEnabler)(nil).DescribeCompliancePolicy), projectID)
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
func (mr *MockCompliancePolicyEncryptionAtRestEnablerMockRecorder) EnableEncryptionAtRest(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableEncryptionAtRest", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestEnabler)(nil).EnableEncryptionAtRest), projectID)
}

// MockCompliancePolicyEncryptionAtRestDisabler is a mock of CompliancePolicyEncryptionAtRestDisabler interface.
type MockCompliancePolicyEncryptionAtRestDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyEncryptionAtRestDisablerMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyEncryptionAtRestDisabler) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyEncryptionAtRestDisablerMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestDisabler)(nil).DescribeCompliancePolicy), projectID)
}

// DisableEncryptionAtRest mocks base method.
func (m *MockCompliancePolicyEncryptionAtRestDisabler) DisableEncryptionAtRest(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableEncryptionAtRest", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableEncryptionAtRest indicates an expected call of DisableEncryptionAtRest.
func (mr *MockCompliancePolicyEncryptionAtRestDisablerMockRecorder) DisableEncryptionAtRest(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableEncryptionAtRest", reflect.TypeOf((*MockCompliancePolicyEncryptionAtRestDisabler)(nil).DisableEncryptionAtRest), projectID)
}

// MockCompliancePolicyEnabler is a mock of CompliancePolicyEnabler interface.
type MockCompliancePolicyEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyEnablerMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyEnabler) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyEnablerMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEnabler)(nil).DescribeCompliancePolicy), projectID)
}

// EnableCompliancePolicy mocks base method.
func (m *MockCompliancePolicyEnabler) EnableCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableCompliancePolicy", projectID, authorizedEmail, authorizedFirstName, authorizedLastName)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableCompliancePolicy indicates an expected call of EnableCompliancePolicy.
func (mr *MockCompliancePolicyEnablerMockRecorder) EnableCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyEnabler)(nil).EnableCompliancePolicy), projectID, authorizedEmail, authorizedFirstName, authorizedLastName)
}

// MockCompliancePolicyCopyProtectionEnabler is a mock of CompliancePolicyCopyProtectionEnabler interface.
type MockCompliancePolicyCopyProtectionEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyCopyProtectionEnablerMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyCopyProtectionEnabler) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyCopyProtectionEnablerMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyCopyProtectionEnabler)(nil).DescribeCompliancePolicy), projectID)
}

// EnableCopyProtection mocks base method.
func (m *MockCompliancePolicyCopyProtectionEnabler) EnableCopyProtection(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableCopyProtection", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableCopyProtection indicates an expected call of EnableCopyProtection.
func (mr *MockCompliancePolicyCopyProtectionEnablerMockRecorder) EnableCopyProtection(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCopyProtection", reflect.TypeOf((*MockCompliancePolicyCopyProtectionEnabler)(nil).EnableCopyProtection), projectID)
}

// MockCompliancePolicyCopyProtectionDisabler is a mock of CompliancePolicyCopyProtectionDisabler interface.
type MockCompliancePolicyCopyProtectionDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyCopyProtectionDisablerMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyCopyProtectionDisabler) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyCopyProtectionDisablerMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyCopyProtectionDisabler)(nil).DescribeCompliancePolicy), projectID)
}

// DisableCopyProtection mocks base method.
func (m *MockCompliancePolicyCopyProtectionDisabler) DisableCopyProtection(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableCopyProtection", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableCopyProtection indicates an expected call of DisableCopyProtection.
func (mr *MockCompliancePolicyCopyProtectionDisablerMockRecorder) DisableCopyProtection(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableCopyProtection", reflect.TypeOf((*MockCompliancePolicyCopyProtectionDisabler)(nil).DisableCopyProtection), projectID)
}

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
func (mr *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyPointInTimeRestoresEnabler)(nil).DescribeCompliancePolicy), projectID)
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
func (mr *MockCompliancePolicyPointInTimeRestoresEnablerMockRecorder) EnablePointInTimeRestore(projectID, restoreWindowDays any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePointInTimeRestore", reflect.TypeOf((*MockCompliancePolicyPointInTimeRestoresEnabler)(nil).EnablePointInTimeRestore), projectID, restoreWindowDays)
}

// MockCompliancePolicyOnDemandPolicyCreator is a mock of CompliancePolicyOnDemandPolicyCreator interface.
type MockCompliancePolicyOnDemandPolicyCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyOnDemandPolicyCreatorMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyOnDemandPolicyCreator) CreateOnDemandPolicy(projectID string, policy *admin.BackupComplianceOnDemandPolicyItem) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOnDemandPolicy", projectID, policy)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOnDemandPolicy indicates an expected call of CreateOnDemandPolicy.
func (mr *MockCompliancePolicyOnDemandPolicyCreatorMockRecorder) CreateOnDemandPolicy(projectID, policy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOnDemandPolicy", reflect.TypeOf((*MockCompliancePolicyOnDemandPolicyCreator)(nil).CreateOnDemandPolicy), projectID, policy)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyOnDemandPolicyCreator) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyOnDemandPolicyCreatorMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyOnDemandPolicyCreator)(nil).DescribeCompliancePolicy), projectID)
}

// MockCompliancePolicyScheduledPolicyCreator is a mock of CompliancePolicyScheduledPolicyCreator interface.
type MockCompliancePolicyScheduledPolicyCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyScheduledPolicyCreatorMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyScheduledPolicyCreator) CreateScheduledPolicy(projectID string, policy *admin.BackupComplianceScheduledPolicyItem) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScheduledPolicy", projectID, policy)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScheduledPolicy indicates an expected call of CreateScheduledPolicy.
func (mr *MockCompliancePolicyScheduledPolicyCreatorMockRecorder) CreateScheduledPolicy(projectID, policy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduledPolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyCreator)(nil).CreateScheduledPolicy), projectID, policy)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyCreator) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyScheduledPolicyCreatorMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyCreator)(nil).DescribeCompliancePolicy), projectID)
}

// MockCompliancePolicyScheduledPolicyDeleter is a mock of CompliancePolicyScheduledPolicyDeleter interface.
type MockCompliancePolicyScheduledPolicyDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyScheduledPolicyDeleterMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyScheduledPolicyDeleter) DeleteScheduledPolicy(projectID, scheduledPolicyID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduledPolicy", projectID, scheduledPolicyID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScheduledPolicy indicates an expected call of DeleteScheduledPolicy.
func (mr *MockCompliancePolicyScheduledPolicyDeleterMockRecorder) DeleteScheduledPolicy(projectID, scheduledPolicyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduledPolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyDeleter)(nil).DeleteScheduledPolicy), projectID, scheduledPolicyID)
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyDeleter) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyScheduledPolicyDeleterMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyDeleter)(nil).DescribeCompliancePolicy), projectID)
}

// MockCompliancePolicyScheduledPolicyUpdater is a mock of CompliancePolicyScheduledPolicyUpdater interface.
type MockCompliancePolicyScheduledPolicyUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyScheduledPolicyUpdaterMockRecorder
	isgomock struct{}
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
func (m *MockCompliancePolicyScheduledPolicyUpdater) DescribeCompliancePolicy(projectID string) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", projectID)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyScheduledPolicyUpdaterMockRecorder) DescribeCompliancePolicy(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyUpdater)(nil).DescribeCompliancePolicy), projectID)
}

// UpdateScheduledPolicy mocks base method.
func (m *MockCompliancePolicyScheduledPolicyUpdater) UpdateScheduledPolicy(projectID string, policy *admin.BackupComplianceScheduledPolicyItem) (*admin.DataProtectionSettings20231001, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduledPolicy", projectID, policy)
	ret0, _ := ret[0].(*admin.DataProtectionSettings20231001)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduledPolicy indicates an expected call of UpdateScheduledPolicy.
func (mr *MockCompliancePolicyScheduledPolicyUpdaterMockRecorder) UpdateScheduledPolicy(projectID, policy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduledPolicy", reflect.TypeOf((*MockCompliancePolicyScheduledPolicyUpdater)(nil).UpdateScheduledPolicy), projectID, policy)
}
