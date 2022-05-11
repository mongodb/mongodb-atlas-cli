// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: AgentLister,AgentUpgrader,AgentAPIKeyLister,AgentAPIKeyCreator,AgentAPIKeyDeleter,AgentGlobalVersionsLister,AgentProjectVersionsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
)

// MockAgentLister is a mock of AgentLister interface.
type MockAgentLister struct {
	ctrl     *gomock.Controller
	recorder *MockAgentListerMockRecorder
}

// MockAgentListerMockRecorder is the mock recorder for MockAgentLister.
type MockAgentListerMockRecorder struct {
	mock *MockAgentLister
}

// NewMockAgentLister creates a new mock instance.
func NewMockAgentLister(ctrl *gomock.Controller) *MockAgentLister {
	mock := &MockAgentLister{ctrl: ctrl}
	mock.recorder = &MockAgentListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentLister) EXPECT() *MockAgentListerMockRecorder {
	return m.recorder
}

// Agents mocks base method.
func (m *MockAgentLister) Agents(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*opsmngr.Agents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Agents", arg0, arg1, arg2)
	ret0, _ := ret[0].(*opsmngr.Agents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Agents indicates an expected call of Agents.
func (mr *MockAgentListerMockRecorder) Agents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Agents", reflect.TypeOf((*MockAgentLister)(nil).Agents), arg0, arg1, arg2)
}

// MockAgentUpgrader is a mock of AgentUpgrader interface.
type MockAgentUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockAgentUpgraderMockRecorder
}

// MockAgentUpgraderMockRecorder is the mock recorder for MockAgentUpgrader.
type MockAgentUpgraderMockRecorder struct {
	mock *MockAgentUpgrader
}

// NewMockAgentUpgrader creates a new mock instance.
func NewMockAgentUpgrader(ctrl *gomock.Controller) *MockAgentUpgrader {
	mock := &MockAgentUpgrader{ctrl: ctrl}
	mock.recorder = &MockAgentUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentUpgrader) EXPECT() *MockAgentUpgraderMockRecorder {
	return m.recorder
}

// UpgradeAgent mocks base method.
func (m *MockAgentUpgrader) UpgradeAgent(arg0 string) (*opsmngr.AutomationConfigAgent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeAgent", arg0)
	ret0, _ := ret[0].(*opsmngr.AutomationConfigAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeAgent indicates an expected call of UpgradeAgent.
func (mr *MockAgentUpgraderMockRecorder) UpgradeAgent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeAgent", reflect.TypeOf((*MockAgentUpgrader)(nil).UpgradeAgent), arg0)
}

// MockAgentAPIKeyLister is a mock of AgentAPIKeyLister interface.
type MockAgentAPIKeyLister struct {
	ctrl     *gomock.Controller
	recorder *MockAgentAPIKeyListerMockRecorder
}

// MockAgentAPIKeyListerMockRecorder is the mock recorder for MockAgentAPIKeyLister.
type MockAgentAPIKeyListerMockRecorder struct {
	mock *MockAgentAPIKeyLister
}

// NewMockAgentAPIKeyLister creates a new mock instance.
func NewMockAgentAPIKeyLister(ctrl *gomock.Controller) *MockAgentAPIKeyLister {
	mock := &MockAgentAPIKeyLister{ctrl: ctrl}
	mock.recorder = &MockAgentAPIKeyListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentAPIKeyLister) EXPECT() *MockAgentAPIKeyListerMockRecorder {
	return m.recorder
}

// AgentAPIKeys mocks base method.
func (m *MockAgentAPIKeyLister) AgentAPIKeys(arg0 string) ([]*opsmngr.AgentAPIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentAPIKeys", arg0)
	ret0, _ := ret[0].([]*opsmngr.AgentAPIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentAPIKeys indicates an expected call of AgentAPIKeys.
func (mr *MockAgentAPIKeyListerMockRecorder) AgentAPIKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentAPIKeys", reflect.TypeOf((*MockAgentAPIKeyLister)(nil).AgentAPIKeys), arg0)
}

// MockAgentAPIKeyCreator is a mock of AgentAPIKeyCreator interface.
type MockAgentAPIKeyCreator struct {
	ctrl     *gomock.Controller
	recorder *MockAgentAPIKeyCreatorMockRecorder
}

// MockAgentAPIKeyCreatorMockRecorder is the mock recorder for MockAgentAPIKeyCreator.
type MockAgentAPIKeyCreatorMockRecorder struct {
	mock *MockAgentAPIKeyCreator
}

// NewMockAgentAPIKeyCreator creates a new mock instance.
func NewMockAgentAPIKeyCreator(ctrl *gomock.Controller) *MockAgentAPIKeyCreator {
	mock := &MockAgentAPIKeyCreator{ctrl: ctrl}
	mock.recorder = &MockAgentAPIKeyCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentAPIKeyCreator) EXPECT() *MockAgentAPIKeyCreatorMockRecorder {
	return m.recorder
}

// CreateAgentAPIKey mocks base method.
func (m *MockAgentAPIKeyCreator) CreateAgentAPIKey(arg0 string, arg1 *opsmngr.AgentAPIKeysRequest) (*opsmngr.AgentAPIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAgentAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.AgentAPIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAgentAPIKey indicates an expected call of CreateAgentAPIKey.
func (mr *MockAgentAPIKeyCreatorMockRecorder) CreateAgentAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgentAPIKey", reflect.TypeOf((*MockAgentAPIKeyCreator)(nil).CreateAgentAPIKey), arg0, arg1)
}

// MockAgentAPIKeyDeleter is a mock of AgentAPIKeyDeleter interface.
type MockAgentAPIKeyDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockAgentAPIKeyDeleterMockRecorder
}

// MockAgentAPIKeyDeleterMockRecorder is the mock recorder for MockAgentAPIKeyDeleter.
type MockAgentAPIKeyDeleterMockRecorder struct {
	mock *MockAgentAPIKeyDeleter
}

// NewMockAgentAPIKeyDeleter creates a new mock instance.
func NewMockAgentAPIKeyDeleter(ctrl *gomock.Controller) *MockAgentAPIKeyDeleter {
	mock := &MockAgentAPIKeyDeleter{ctrl: ctrl}
	mock.recorder = &MockAgentAPIKeyDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentAPIKeyDeleter) EXPECT() *MockAgentAPIKeyDeleterMockRecorder {
	return m.recorder
}

// DeleteAgentAPIKey mocks base method.
func (m *MockAgentAPIKeyDeleter) DeleteAgentAPIKey(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgentAPIKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgentAPIKey indicates an expected call of DeleteAgentAPIKey.
func (mr *MockAgentAPIKeyDeleterMockRecorder) DeleteAgentAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentAPIKey", reflect.TypeOf((*MockAgentAPIKeyDeleter)(nil).DeleteAgentAPIKey), arg0, arg1)
}

// MockAgentGlobalVersionsLister is a mock of AgentGlobalVersionsLister interface.
type MockAgentGlobalVersionsLister struct {
	ctrl     *gomock.Controller
	recorder *MockAgentGlobalVersionsListerMockRecorder
}

// MockAgentGlobalVersionsListerMockRecorder is the mock recorder for MockAgentGlobalVersionsLister.
type MockAgentGlobalVersionsListerMockRecorder struct {
	mock *MockAgentGlobalVersionsLister
}

// NewMockAgentGlobalVersionsLister creates a new mock instance.
func NewMockAgentGlobalVersionsLister(ctrl *gomock.Controller) *MockAgentGlobalVersionsLister {
	mock := &MockAgentGlobalVersionsLister{ctrl: ctrl}
	mock.recorder = &MockAgentGlobalVersionsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentGlobalVersionsLister) EXPECT() *MockAgentGlobalVersionsListerMockRecorder {
	return m.recorder
}

// AgentGlobalVersions mocks base method.
func (m *MockAgentGlobalVersionsLister) AgentGlobalVersions() (*opsmngr.SoftwareVersions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentGlobalVersions")
	ret0, _ := ret[0].(*opsmngr.SoftwareVersions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentGlobalVersions indicates an expected call of AgentGlobalVersions.
func (mr *MockAgentGlobalVersionsListerMockRecorder) AgentGlobalVersions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentGlobalVersions", reflect.TypeOf((*MockAgentGlobalVersionsLister)(nil).AgentGlobalVersions))
}

// MockAgentProjectVersionsLister is a mock of AgentProjectVersionsLister interface.
type MockAgentProjectVersionsLister struct {
	ctrl     *gomock.Controller
	recorder *MockAgentProjectVersionsListerMockRecorder
}

// MockAgentProjectVersionsListerMockRecorder is the mock recorder for MockAgentProjectVersionsLister.
type MockAgentProjectVersionsListerMockRecorder struct {
	mock *MockAgentProjectVersionsLister
}

// NewMockAgentProjectVersionsLister creates a new mock instance.
func NewMockAgentProjectVersionsLister(ctrl *gomock.Controller) *MockAgentProjectVersionsLister {
	mock := &MockAgentProjectVersionsLister{ctrl: ctrl}
	mock.recorder = &MockAgentProjectVersionsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentProjectVersionsLister) EXPECT() *MockAgentProjectVersionsListerMockRecorder {
	return m.recorder
}

// AgentProjectVersions mocks base method.
func (m *MockAgentProjectVersionsLister) AgentProjectVersions(arg0 string) (*opsmngr.AgentVersions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentProjectVersions", arg0)
	ret0, _ := ret[0].(*opsmngr.AgentVersions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentProjectVersions indicates an expected call of AgentProjectVersions.
func (mr *MockAgentProjectVersionsListerMockRecorder) AgentProjectVersions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentProjectVersions", reflect.TypeOf((*MockAgentProjectVersionsLister)(nil).AgentProjectVersions), arg0)
}
