// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: TeamLister,TeamDescriber,TeamCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockTeamLister is a mock of TeamLister interface
type MockTeamLister struct {
	ctrl     *gomock.Controller
	recorder *MockTeamListerMockRecorder
}

// MockTeamListerMockRecorder is the mock recorder for MockTeamLister
type MockTeamListerMockRecorder struct {
	mock *MockTeamLister
}

// NewMockTeamLister creates a new mock instance
func NewMockTeamLister(ctrl *gomock.Controller) *MockTeamLister {
	mock := &MockTeamLister{ctrl: ctrl}
	mock.recorder = &MockTeamListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTeamLister) EXPECT() *MockTeamListerMockRecorder {
	return m.recorder
}

// Teams mocks base method
func (m *MockTeamLister) Teams(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teams", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teams indicates an expected call of Teams
func (mr *MockTeamListerMockRecorder) Teams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teams", reflect.TypeOf((*MockTeamLister)(nil).Teams), arg0, arg1)
}

// MockTeamDescriber is a mock of TeamDescriber interface
type MockTeamDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockTeamDescriberMockRecorder
}

// MockTeamDescriberMockRecorder is the mock recorder for MockTeamDescriber
type MockTeamDescriberMockRecorder struct {
	mock *MockTeamDescriber
}

// NewMockTeamDescriber creates a new mock instance
func NewMockTeamDescriber(ctrl *gomock.Controller) *MockTeamDescriber {
	mock := &MockTeamDescriber{ctrl: ctrl}
	mock.recorder = &MockTeamDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTeamDescriber) EXPECT() *MockTeamDescriberMockRecorder {
	return m.recorder
}

// TeamByID mocks base method
func (m *MockTeamDescriber) TeamByID(arg0, arg1 string) (*mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByID", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByID indicates an expected call of TeamByID
func (mr *MockTeamDescriberMockRecorder) TeamByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByID", reflect.TypeOf((*MockTeamDescriber)(nil).TeamByID), arg0, arg1)
}

// TeamByName mocks base method
func (m *MockTeamDescriber) TeamByName(arg0, arg1 string) (*mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByName", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByName indicates an expected call of TeamByName
func (mr *MockTeamDescriberMockRecorder) TeamByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByName", reflect.TypeOf((*MockTeamDescriber)(nil).TeamByName), arg0, arg1)
}

// MockTeamCreator is a mock of TeamCreator interface
type MockTeamCreator struct {
	ctrl     *gomock.Controller
	recorder *MockTeamCreatorMockRecorder
}

// MockTeamCreatorMockRecorder is the mock recorder for MockTeamCreator
type MockTeamCreatorMockRecorder struct {
	mock *MockTeamCreator
}

// NewMockTeamCreator creates a new mock instance
func NewMockTeamCreator(ctrl *gomock.Controller) *MockTeamCreator {
	mock := &MockTeamCreator{ctrl: ctrl}
	mock.recorder = &MockTeamCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTeamCreator) EXPECT() *MockTeamCreatorMockRecorder {
	return m.recorder
}

// CreateTeam mocks base method
func (m *MockTeamCreator) CreateTeam(arg0 string, arg1 *mongodbatlas.Team) (*mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeam indicates an expected call of CreateTeam
func (mr *MockTeamCreatorMockRecorder) CreateTeam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockTeamCreator)(nil).CreateTeam), arg0, arg1)
}
