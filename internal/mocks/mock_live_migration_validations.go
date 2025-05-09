// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: LiveMigrationValidationsCreator,LiveMigrationCutoverCreator,LiveMigrationValidationsDescriber)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_live_migration_validations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store LiveMigrationValidationsCreator,LiveMigrationCutoverCreator,LiveMigrationValidationsDescriber
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20250312002/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockLiveMigrationValidationsCreator is a mock of LiveMigrationValidationsCreator interface.
type MockLiveMigrationValidationsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMigrationValidationsCreatorMockRecorder
	isgomock struct{}
}

// MockLiveMigrationValidationsCreatorMockRecorder is the mock recorder for MockLiveMigrationValidationsCreator.
type MockLiveMigrationValidationsCreatorMockRecorder struct {
	mock *MockLiveMigrationValidationsCreator
}

// NewMockLiveMigrationValidationsCreator creates a new mock instance.
func NewMockLiveMigrationValidationsCreator(ctrl *gomock.Controller) *MockLiveMigrationValidationsCreator {
	mock := &MockLiveMigrationValidationsCreator{ctrl: ctrl}
	mock.recorder = &MockLiveMigrationValidationsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiveMigrationValidationsCreator) EXPECT() *MockLiveMigrationValidationsCreatorMockRecorder {
	return m.recorder
}

// CreateValidation mocks base method.
func (m *MockLiveMigrationValidationsCreator) CreateValidation(arg0 string, arg1 *admin.LiveMigrationRequest20240530) (*admin.LiveImportValidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateValidation", arg0, arg1)
	ret0, _ := ret[0].(*admin.LiveImportValidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateValidation indicates an expected call of CreateValidation.
func (mr *MockLiveMigrationValidationsCreatorMockRecorder) CreateValidation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateValidation", reflect.TypeOf((*MockLiveMigrationValidationsCreator)(nil).CreateValidation), arg0, arg1)
}

// MockLiveMigrationCutoverCreator is a mock of LiveMigrationCutoverCreator interface.
type MockLiveMigrationCutoverCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMigrationCutoverCreatorMockRecorder
	isgomock struct{}
}

// MockLiveMigrationCutoverCreatorMockRecorder is the mock recorder for MockLiveMigrationCutoverCreator.
type MockLiveMigrationCutoverCreatorMockRecorder struct {
	mock *MockLiveMigrationCutoverCreator
}

// NewMockLiveMigrationCutoverCreator creates a new mock instance.
func NewMockLiveMigrationCutoverCreator(ctrl *gomock.Controller) *MockLiveMigrationCutoverCreator {
	mock := &MockLiveMigrationCutoverCreator{ctrl: ctrl}
	mock.recorder = &MockLiveMigrationCutoverCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiveMigrationCutoverCreator) EXPECT() *MockLiveMigrationCutoverCreatorMockRecorder {
	return m.recorder
}

// CreateLiveMigrationCutover mocks base method.
func (m *MockLiveMigrationCutoverCreator) CreateLiveMigrationCutover(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLiveMigrationCutover", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLiveMigrationCutover indicates an expected call of CreateLiveMigrationCutover.
func (mr *MockLiveMigrationCutoverCreatorMockRecorder) CreateLiveMigrationCutover(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLiveMigrationCutover", reflect.TypeOf((*MockLiveMigrationCutoverCreator)(nil).CreateLiveMigrationCutover), arg0, arg1)
}

// MockLiveMigrationValidationsDescriber is a mock of LiveMigrationValidationsDescriber interface.
type MockLiveMigrationValidationsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMigrationValidationsDescriberMockRecorder
	isgomock struct{}
}

// MockLiveMigrationValidationsDescriberMockRecorder is the mock recorder for MockLiveMigrationValidationsDescriber.
type MockLiveMigrationValidationsDescriberMockRecorder struct {
	mock *MockLiveMigrationValidationsDescriber
}

// NewMockLiveMigrationValidationsDescriber creates a new mock instance.
func NewMockLiveMigrationValidationsDescriber(ctrl *gomock.Controller) *MockLiveMigrationValidationsDescriber {
	mock := &MockLiveMigrationValidationsDescriber{ctrl: ctrl}
	mock.recorder = &MockLiveMigrationValidationsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiveMigrationValidationsDescriber) EXPECT() *MockLiveMigrationValidationsDescriberMockRecorder {
	return m.recorder
}

// GetValidationStatus mocks base method.
func (m *MockLiveMigrationValidationsDescriber) GetValidationStatus(arg0, arg1 string) (*admin.LiveImportValidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidationStatus", arg0, arg1)
	ret0, _ := ret[0].(*admin.LiveImportValidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidationStatus indicates an expected call of GetValidationStatus.
func (mr *MockLiveMigrationValidationsDescriberMockRecorder) GetValidationStatus(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidationStatus", reflect.TypeOf((*MockLiveMigrationValidationsDescriber)(nil).GetValidationStatus), arg0, arg1)
}
