// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/store (interfaces: PrivateEndpointListerDeprecated,PrivateEndpointDescriberDeprecated,PrivateEndpointCreatorDeprecated,PrivateEndpointDeleterDeprecated,InterfaceEndpointCreatorDeprecated,InterfaceEndpointDescriberDeprecated,InterfaceEndpointDeleterDeprecated)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockPrivateEndpointListerDeprecated is a mock of PrivateEndpointListerDeprecated interface.
type MockPrivateEndpointListerDeprecated struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointListerDeprecatedMockRecorder
}

// MockPrivateEndpointListerDeprecatedMockRecorder is the mock recorder for MockPrivateEndpointListerDeprecated.
type MockPrivateEndpointListerDeprecatedMockRecorder struct {
	mock *MockPrivateEndpointListerDeprecated
}

// NewMockPrivateEndpointListerDeprecated creates a new mock instance.
func NewMockPrivateEndpointListerDeprecated(ctrl *gomock.Controller) *MockPrivateEndpointListerDeprecated {
	mock := &MockPrivateEndpointListerDeprecated{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointListerDeprecatedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointListerDeprecated) EXPECT() *MockPrivateEndpointListerDeprecatedMockRecorder {
	return m.recorder
}

// PrivateEndpointsDeprecated mocks base method.
func (m *MockPrivateEndpointListerDeprecated) PrivateEndpointsDeprecated(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.PrivateEndpointConnectionDeprecated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpointsDeprecated", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.PrivateEndpointConnectionDeprecated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpointsDeprecated indicates an expected call of PrivateEndpointsDeprecated.
func (mr *MockPrivateEndpointListerDeprecatedMockRecorder) PrivateEndpointsDeprecated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpointsDeprecated", reflect.TypeOf((*MockPrivateEndpointListerDeprecated)(nil).PrivateEndpointsDeprecated), arg0, arg1)
}

// MockPrivateEndpointDescriberDeprecated is a mock of PrivateEndpointDescriberDeprecated interface.
type MockPrivateEndpointDescriberDeprecated struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointDescriberDeprecatedMockRecorder
}

// MockPrivateEndpointDescriberDeprecatedMockRecorder is the mock recorder for MockPrivateEndpointDescriberDeprecated.
type MockPrivateEndpointDescriberDeprecatedMockRecorder struct {
	mock *MockPrivateEndpointDescriberDeprecated
}

// NewMockPrivateEndpointDescriberDeprecated creates a new mock instance.
func NewMockPrivateEndpointDescriberDeprecated(ctrl *gomock.Controller) *MockPrivateEndpointDescriberDeprecated {
	mock := &MockPrivateEndpointDescriberDeprecated{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointDescriberDeprecatedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointDescriberDeprecated) EXPECT() *MockPrivateEndpointDescriberDeprecatedMockRecorder {
	return m.recorder
}

// PrivateEndpointDeprecated mocks base method.
func (m *MockPrivateEndpointDescriberDeprecated) PrivateEndpointDeprecated(arg0, arg1 string) (*mongodbatlas.PrivateEndpointConnectionDeprecated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpointDeprecated", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.PrivateEndpointConnectionDeprecated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpointDeprecated indicates an expected call of PrivateEndpointDeprecated.
func (mr *MockPrivateEndpointDescriberDeprecatedMockRecorder) PrivateEndpointDeprecated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpointDeprecated", reflect.TypeOf((*MockPrivateEndpointDescriberDeprecated)(nil).PrivateEndpointDeprecated), arg0, arg1)
}

// MockPrivateEndpointCreatorDeprecated is a mock of PrivateEndpointCreatorDeprecated interface.
type MockPrivateEndpointCreatorDeprecated struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointCreatorDeprecatedMockRecorder
}

// MockPrivateEndpointCreatorDeprecatedMockRecorder is the mock recorder for MockPrivateEndpointCreatorDeprecated.
type MockPrivateEndpointCreatorDeprecatedMockRecorder struct {
	mock *MockPrivateEndpointCreatorDeprecated
}

// NewMockPrivateEndpointCreatorDeprecated creates a new mock instance.
func NewMockPrivateEndpointCreatorDeprecated(ctrl *gomock.Controller) *MockPrivateEndpointCreatorDeprecated {
	mock := &MockPrivateEndpointCreatorDeprecated{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointCreatorDeprecatedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointCreatorDeprecated) EXPECT() *MockPrivateEndpointCreatorDeprecatedMockRecorder {
	return m.recorder
}

// CreatePrivateEndpointDeprecated mocks base method.
func (m *MockPrivateEndpointCreatorDeprecated) CreatePrivateEndpointDeprecated(arg0 string, arg1 *mongodbatlas.PrivateEndpointConnectionDeprecated) (*mongodbatlas.PrivateEndpointConnectionDeprecated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePrivateEndpointDeprecated", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.PrivateEndpointConnectionDeprecated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePrivateEndpointDeprecated indicates an expected call of CreatePrivateEndpointDeprecated.
func (mr *MockPrivateEndpointCreatorDeprecatedMockRecorder) CreatePrivateEndpointDeprecated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrivateEndpointDeprecated", reflect.TypeOf((*MockPrivateEndpointCreatorDeprecated)(nil).CreatePrivateEndpointDeprecated), arg0, arg1)
}

// MockPrivateEndpointDeleterDeprecated is a mock of PrivateEndpointDeleterDeprecated interface.
type MockPrivateEndpointDeleterDeprecated struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointDeleterDeprecatedMockRecorder
}

// MockPrivateEndpointDeleterDeprecatedMockRecorder is the mock recorder for MockPrivateEndpointDeleterDeprecated.
type MockPrivateEndpointDeleterDeprecatedMockRecorder struct {
	mock *MockPrivateEndpointDeleterDeprecated
}

// NewMockPrivateEndpointDeleterDeprecated creates a new mock instance.
func NewMockPrivateEndpointDeleterDeprecated(ctrl *gomock.Controller) *MockPrivateEndpointDeleterDeprecated {
	mock := &MockPrivateEndpointDeleterDeprecated{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointDeleterDeprecatedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointDeleterDeprecated) EXPECT() *MockPrivateEndpointDeleterDeprecatedMockRecorder {
	return m.recorder
}

// DeletePrivateEndpointDeprecated mocks base method.
func (m *MockPrivateEndpointDeleterDeprecated) DeletePrivateEndpointDeprecated(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePrivateEndpointDeprecated", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePrivateEndpointDeprecated indicates an expected call of DeletePrivateEndpointDeprecated.
func (mr *MockPrivateEndpointDeleterDeprecatedMockRecorder) DeletePrivateEndpointDeprecated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePrivateEndpointDeprecated", reflect.TypeOf((*MockPrivateEndpointDeleterDeprecated)(nil).DeletePrivateEndpointDeprecated), arg0, arg1)
}

// MockInterfaceEndpointCreatorDeprecated is a mock of InterfaceEndpointCreatorDeprecated interface.
type MockInterfaceEndpointCreatorDeprecated struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointCreatorDeprecatedMockRecorder
}

// MockInterfaceEndpointCreatorDeprecatedMockRecorder is the mock recorder for MockInterfaceEndpointCreatorDeprecated.
type MockInterfaceEndpointCreatorDeprecatedMockRecorder struct {
	mock *MockInterfaceEndpointCreatorDeprecated
}

// NewMockInterfaceEndpointCreatorDeprecated creates a new mock instance.
func NewMockInterfaceEndpointCreatorDeprecated(ctrl *gomock.Controller) *MockInterfaceEndpointCreatorDeprecated {
	mock := &MockInterfaceEndpointCreatorDeprecated{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointCreatorDeprecatedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointCreatorDeprecated) EXPECT() *MockInterfaceEndpointCreatorDeprecatedMockRecorder {
	return m.recorder
}

// CreateInterfaceEndpointDeprecated mocks base method.
func (m *MockInterfaceEndpointCreatorDeprecated) CreateInterfaceEndpointDeprecated(arg0, arg1, arg2 string) (*mongodbatlas.InterfaceEndpointConnectionDeprecated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInterfaceEndpointDeprecated", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.InterfaceEndpointConnectionDeprecated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInterfaceEndpointDeprecated indicates an expected call of CreateInterfaceEndpointDeprecated.
func (mr *MockInterfaceEndpointCreatorDeprecatedMockRecorder) CreateInterfaceEndpointDeprecated(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInterfaceEndpointDeprecated", reflect.TypeOf((*MockInterfaceEndpointCreatorDeprecated)(nil).CreateInterfaceEndpointDeprecated), arg0, arg1, arg2)
}

// MockInterfaceEndpointDescriberDeprecated is a mock of InterfaceEndpointDescriberDeprecated interface.
type MockInterfaceEndpointDescriberDeprecated struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointDescriberDeprecatedMockRecorder
}

// MockInterfaceEndpointDescriberDeprecatedMockRecorder is the mock recorder for MockInterfaceEndpointDescriberDeprecated.
type MockInterfaceEndpointDescriberDeprecatedMockRecorder struct {
	mock *MockInterfaceEndpointDescriberDeprecated
}

// NewMockInterfaceEndpointDescriberDeprecated creates a new mock instance.
func NewMockInterfaceEndpointDescriberDeprecated(ctrl *gomock.Controller) *MockInterfaceEndpointDescriberDeprecated {
	mock := &MockInterfaceEndpointDescriberDeprecated{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointDescriberDeprecatedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointDescriberDeprecated) EXPECT() *MockInterfaceEndpointDescriberDeprecatedMockRecorder {
	return m.recorder
}

// InterfaceEndpointDeprecated mocks base method.
func (m *MockInterfaceEndpointDescriberDeprecated) InterfaceEndpointDeprecated(arg0, arg1, arg2 string) (*mongodbatlas.InterfaceEndpointConnectionDeprecated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterfaceEndpointDeprecated", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.InterfaceEndpointConnectionDeprecated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InterfaceEndpointDeprecated indicates an expected call of InterfaceEndpointDeprecated.
func (mr *MockInterfaceEndpointDescriberDeprecatedMockRecorder) InterfaceEndpointDeprecated(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterfaceEndpointDeprecated", reflect.TypeOf((*MockInterfaceEndpointDescriberDeprecated)(nil).InterfaceEndpointDeprecated), arg0, arg1, arg2)
}

// MockInterfaceEndpointDeleterDeprecated is a mock of InterfaceEndpointDeleterDeprecated interface.
type MockInterfaceEndpointDeleterDeprecated struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceEndpointDeleterDeprecatedMockRecorder
}

// MockInterfaceEndpointDeleterDeprecatedMockRecorder is the mock recorder for MockInterfaceEndpointDeleterDeprecated.
type MockInterfaceEndpointDeleterDeprecatedMockRecorder struct {
	mock *MockInterfaceEndpointDeleterDeprecated
}

// NewMockInterfaceEndpointDeleterDeprecated creates a new mock instance.
func NewMockInterfaceEndpointDeleterDeprecated(ctrl *gomock.Controller) *MockInterfaceEndpointDeleterDeprecated {
	mock := &MockInterfaceEndpointDeleterDeprecated{ctrl: ctrl}
	mock.recorder = &MockInterfaceEndpointDeleterDeprecatedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceEndpointDeleterDeprecated) EXPECT() *MockInterfaceEndpointDeleterDeprecatedMockRecorder {
	return m.recorder
}

// DeleteInterfaceEndpointDeprecated mocks base method.
func (m *MockInterfaceEndpointDeleterDeprecated) DeleteInterfaceEndpointDeprecated(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInterfaceEndpointDeprecated", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInterfaceEndpointDeprecated indicates an expected call of DeleteInterfaceEndpointDeprecated.
func (mr *MockInterfaceEndpointDeleterDeprecatedMockRecorder) DeleteInterfaceEndpointDeprecated(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInterfaceEndpointDeprecated", reflect.TypeOf((*MockInterfaceEndpointDeleterDeprecated)(nil).DeleteInterfaceEndpointDeprecated), arg0, arg1, arg2)
}
