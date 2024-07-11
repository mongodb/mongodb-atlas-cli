// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: IdentityProviderLister,IdentityProviderDescriber,IdentityProviderCreator,IdentityProviderDeleter,IdentityProviderUpdater,IdentityProviderJwkRevoker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

// MockIdentityProviderLister is a mock of IdentityProviderLister interface.
type MockIdentityProviderLister struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityProviderListerMockRecorder
}

// MockIdentityProviderListerMockRecorder is the mock recorder for MockIdentityProviderLister.
type MockIdentityProviderListerMockRecorder struct {
	mock *MockIdentityProviderLister
}

// NewMockIdentityProviderLister creates a new mock instance.
func NewMockIdentityProviderLister(ctrl *gomock.Controller) *MockIdentityProviderLister {
	mock := &MockIdentityProviderLister{ctrl: ctrl}
	mock.recorder = &MockIdentityProviderListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityProviderLister) EXPECT() *MockIdentityProviderListerMockRecorder {
	return m.recorder
}

// IdentityProviders mocks base method.
func (m *MockIdentityProviderLister) IdentityProviders(arg0 *admin.ListIdentityProvidersApiParams) (*admin.PaginatedFederationIdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityProviders", arg0)
	ret0, _ := ret[0].(*admin.PaginatedFederationIdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IdentityProviders indicates an expected call of IdentityProviders.
func (mr *MockIdentityProviderListerMockRecorder) IdentityProviders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityProviders", reflect.TypeOf((*MockIdentityProviderLister)(nil).IdentityProviders), arg0)
}

// MockIdentityProviderDescriber is a mock of IdentityProviderDescriber interface.
type MockIdentityProviderDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityProviderDescriberMockRecorder
}

// MockIdentityProviderDescriberMockRecorder is the mock recorder for MockIdentityProviderDescriber.
type MockIdentityProviderDescriberMockRecorder struct {
	mock *MockIdentityProviderDescriber
}

// NewMockIdentityProviderDescriber creates a new mock instance.
func NewMockIdentityProviderDescriber(ctrl *gomock.Controller) *MockIdentityProviderDescriber {
	mock := &MockIdentityProviderDescriber{ctrl: ctrl}
	mock.recorder = &MockIdentityProviderDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityProviderDescriber) EXPECT() *MockIdentityProviderDescriberMockRecorder {
	return m.recorder
}

// IdentityProvider mocks base method.
func (m *MockIdentityProviderDescriber) IdentityProvider(arg0 *admin.GetIdentityProviderApiParams) (*admin.FederationIdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityProvider", arg0)
	ret0, _ := ret[0].(*admin.FederationIdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IdentityProvider indicates an expected call of IdentityProvider.
func (mr *MockIdentityProviderDescriberMockRecorder) IdentityProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityProvider", reflect.TypeOf((*MockIdentityProviderDescriber)(nil).IdentityProvider), arg0)
}

// MockIdentityProviderCreator is a mock of IdentityProviderCreator interface.
type MockIdentityProviderCreator struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityProviderCreatorMockRecorder
}

// MockIdentityProviderCreatorMockRecorder is the mock recorder for MockIdentityProviderCreator.
type MockIdentityProviderCreatorMockRecorder struct {
	mock *MockIdentityProviderCreator
}

// NewMockIdentityProviderCreator creates a new mock instance.
func NewMockIdentityProviderCreator(ctrl *gomock.Controller) *MockIdentityProviderCreator {
	mock := &MockIdentityProviderCreator{ctrl: ctrl}
	mock.recorder = &MockIdentityProviderCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityProviderCreator) EXPECT() *MockIdentityProviderCreatorMockRecorder {
	return m.recorder
}

// CreateIdentityProvider mocks base method.
func (m *MockIdentityProviderCreator) CreateIdentityProvider(arg0 *admin.CreateIdentityProviderApiParams) (*admin.FederationOidcIdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIdentityProvider", arg0)
	ret0, _ := ret[0].(*admin.FederationOidcIdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdentityProvider indicates an expected call of CreateIdentityProvider.
func (mr *MockIdentityProviderCreatorMockRecorder) CreateIdentityProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdentityProvider", reflect.TypeOf((*MockIdentityProviderCreator)(nil).CreateIdentityProvider), arg0)
}

// MockIdentityProviderDeleter is a mock of IdentityProviderDeleter interface.
type MockIdentityProviderDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityProviderDeleterMockRecorder
}

// MockIdentityProviderDeleterMockRecorder is the mock recorder for MockIdentityProviderDeleter.
type MockIdentityProviderDeleterMockRecorder struct {
	mock *MockIdentityProviderDeleter
}

// NewMockIdentityProviderDeleter creates a new mock instance.
func NewMockIdentityProviderDeleter(ctrl *gomock.Controller) *MockIdentityProviderDeleter {
	mock := &MockIdentityProviderDeleter{ctrl: ctrl}
	mock.recorder = &MockIdentityProviderDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityProviderDeleter) EXPECT() *MockIdentityProviderDeleterMockRecorder {
	return m.recorder
}

// DeleteIdentityProvider mocks base method.
func (m *MockIdentityProviderDeleter) DeleteIdentityProvider(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIdentityProvider", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIdentityProvider indicates an expected call of DeleteIdentityProvider.
func (mr *MockIdentityProviderDeleterMockRecorder) DeleteIdentityProvider(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIdentityProvider", reflect.TypeOf((*MockIdentityProviderDeleter)(nil).DeleteIdentityProvider), arg0, arg1)
}

// MockIdentityProviderUpdater is a mock of IdentityProviderUpdater interface.
type MockIdentityProviderUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityProviderUpdaterMockRecorder
}

// MockIdentityProviderUpdaterMockRecorder is the mock recorder for MockIdentityProviderUpdater.
type MockIdentityProviderUpdaterMockRecorder struct {
	mock *MockIdentityProviderUpdater
}

// NewMockIdentityProviderUpdater creates a new mock instance.
func NewMockIdentityProviderUpdater(ctrl *gomock.Controller) *MockIdentityProviderUpdater {
	mock := &MockIdentityProviderUpdater{ctrl: ctrl}
	mock.recorder = &MockIdentityProviderUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityProviderUpdater) EXPECT() *MockIdentityProviderUpdaterMockRecorder {
	return m.recorder
}

// UpdateIdentityProvider mocks base method.
func (m *MockIdentityProviderUpdater) UpdateIdentityProvider(arg0 *admin.UpdateIdentityProviderApiParams) (*admin.FederationIdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIdentityProvider", arg0)
	ret0, _ := ret[0].(*admin.FederationIdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIdentityProvider indicates an expected call of UpdateIdentityProvider.
func (mr *MockIdentityProviderUpdaterMockRecorder) UpdateIdentityProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIdentityProvider", reflect.TypeOf((*MockIdentityProviderUpdater)(nil).UpdateIdentityProvider), arg0)
}

// MockIdentityProviderJwkRevoker is a mock of IdentityProviderJwkRevoker interface.
type MockIdentityProviderJwkRevoker struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityProviderJwkRevokerMockRecorder
}

// MockIdentityProviderJwkRevokerMockRecorder is the mock recorder for MockIdentityProviderJwkRevoker.
type MockIdentityProviderJwkRevokerMockRecorder struct {
	mock *MockIdentityProviderJwkRevoker
}

// NewMockIdentityProviderJwkRevoker creates a new mock instance.
func NewMockIdentityProviderJwkRevoker(ctrl *gomock.Controller) *MockIdentityProviderJwkRevoker {
	mock := &MockIdentityProviderJwkRevoker{ctrl: ctrl}
	mock.recorder = &MockIdentityProviderJwkRevokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityProviderJwkRevoker) EXPECT() *MockIdentityProviderJwkRevokerMockRecorder {
	return m.recorder
}

// RevokeJwksFromIdentityProvider mocks base method.
func (m *MockIdentityProviderJwkRevoker) RevokeJwksFromIdentityProvider(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeJwksFromIdentityProvider", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeJwksFromIdentityProvider indicates an expected call of RevokeJwksFromIdentityProvider.
func (mr *MockIdentityProviderJwkRevokerMockRecorder) RevokeJwksFromIdentityProvider(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeJwksFromIdentityProvider", reflect.TypeOf((*MockIdentityProviderJwkRevoker)(nil).RevokeJwksFromIdentityProvider), arg0, arg1)
}
