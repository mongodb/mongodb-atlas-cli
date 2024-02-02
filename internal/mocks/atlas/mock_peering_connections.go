// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: PeeringConnectionLister)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	atlas "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	admin "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

// MockPeeringConnectionLister is a mock of PeeringConnectionLister interface.
type MockPeeringConnectionLister struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionListerMockRecorder
}

// MockPeeringConnectionListerMockRecorder is the mock recorder for MockPeeringConnectionLister.
type MockPeeringConnectionListerMockRecorder struct {
	mock *MockPeeringConnectionLister
}

// NewMockPeeringConnectionLister creates a new mock instance.
func NewMockPeeringConnectionLister(ctrl *gomock.Controller) *MockPeeringConnectionLister {
	mock := &MockPeeringConnectionLister{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeeringConnectionLister) EXPECT() *MockPeeringConnectionListerMockRecorder {
	return m.recorder
}

// PeeringConnections mocks base method.
func (m *MockPeeringConnectionLister) PeeringConnections(arg0 string, arg1 *atlas.ContainersListOptions) ([]admin.BaseNetworkPeeringConnectionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnections", arg0, arg1)
	ret0, _ := ret[0].([]admin.BaseNetworkPeeringConnectionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnections indicates an expected call of PeeringConnections.
func (mr *MockPeeringConnectionListerMockRecorder) PeeringConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnections", reflect.TypeOf((*MockPeeringConnectionLister)(nil).PeeringConnections), arg0, arg1)
}
