// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: LiveMigrationValidationsCreator,LiveMigrationCutoverCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockLiveMigrationValidationsCreator is a mock of LiveMigrationValidationsCreator interface.
type MockLiveMigrationValidationsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMigrationValidationsCreatorMockRecorder
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
func (m *MockLiveMigrationValidationsCreator) CreateValidation(arg0 string, arg1 *mongodbatlas.LiveMigration) (*mongodbatlas.Validation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateValidation", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Validation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateValidation indicates an expected call of CreateValidation.
func (mr *MockLiveMigrationValidationsCreatorMockRecorder) CreateValidation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateValidation", reflect.TypeOf((*MockLiveMigrationValidationsCreator)(nil).CreateValidation), arg0, arg1)
}

// MockLiveMigrationCutoverCreator is a mock of LiveMigrationCutoverCreator interface.
type MockLiveMigrationCutoverCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMigrationCutoverCreatorMockRecorder
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
func (m *MockLiveMigrationCutoverCreator) CreateLiveMigrationCutover(arg0, arg1 string) (*mongodbatlas.Validation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLiveMigrationCutover", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Validation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLiveMigrationCutover indicates an expected call of CreateLiveMigrationCutover.
func (mr *MockLiveMigrationCutoverCreatorMockRecorder) CreateLiveMigrationCutover(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLiveMigrationCutover", reflect.TypeOf((*MockLiveMigrationCutoverCreator)(nil).CreateLiveMigrationCutover), arg0, arg1)
}
