// Code generated by MockGen. DO NOT EDIT.
// Source: internal/server/usecase/interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	models "github.com/Genry72/GophKeeper/internal/server/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIusers is a mock of Iusers interface.
type MockIusers struct {
	ctrl     *gomock.Controller
	recorder *MockIusersMockRecorder
}

// MockIusersMockRecorder is the mock recorder for MockIusers.
type MockIusersMockRecorder struct {
	mock *MockIusers
}

// NewMockIusers creates a new mock instance.
func NewMockIusers(ctrl *gomock.Controller) *MockIusers {
	mock := &MockIusers{ctrl: ctrl}
	mock.recorder = &MockIusersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIusers) EXPECT() *MockIusersMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockIusers) AuthUser(ctx context.Context, login, pass string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUser", ctx, login, pass)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockIusersMockRecorder) AuthUser(ctx, login, pass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockIusers)(nil).AuthUser), ctx, login, pass)
}

// RegisterUser mocks base method.
func (m *MockIusers) RegisterUser(ctx context.Context, login, pass string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, login, pass)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockIusersMockRecorder) RegisterUser(ctx, login, pass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIusers)(nil).RegisterUser), ctx, login, pass)
}

// MockISecrets is a mock of ISecrets interface.
type MockISecrets struct {
	ctrl     *gomock.Controller
	recorder *MockISecretsMockRecorder
}

// MockISecretsMockRecorder is the mock recorder for MockISecrets.
type MockISecretsMockRecorder struct {
	mock *MockISecrets
}

// NewMockISecrets creates a new mock instance.
func NewMockISecrets(ctrl *gomock.Controller) *MockISecrets {
	mock := &MockISecrets{ctrl: ctrl}
	mock.recorder = &MockISecretsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISecrets) EXPECT() *MockISecretsMockRecorder {
	return m.recorder
}

// AddSecret mocks base method.
func (m *MockISecrets) AddSecret(ctx context.Context, secretTypeID models.SecretTypeID, secretName string, secretContent []byte) (models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecret", ctx, secretTypeID, secretName, secretContent)
	ret0, _ := ret[0].(models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSecret indicates an expected call of AddSecret.
func (mr *MockISecretsMockRecorder) AddSecret(ctx, secretTypeID, secretName, secretContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecret", reflect.TypeOf((*MockISecrets)(nil).AddSecret), ctx, secretTypeID, secretName, secretContent)
}

// DeleteSecret mocks base method.
func (m *MockISecrets) DeleteSecret(ctx context.Context, secretID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", ctx, secretID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockISecretsMockRecorder) DeleteSecret(ctx, secretID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockISecrets)(nil).DeleteSecret), ctx, secretID)
}

// EditSecret mocks base method.
func (m *MockISecrets) EditSecret(ctx context.Context, secretName string, secretID int64, secretContent []byte) (models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditSecret", ctx, secretName, secretID, secretContent)
	ret0, _ := ret[0].(models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditSecret indicates an expected call of EditSecret.
func (mr *MockISecretsMockRecorder) EditSecret(ctx, secretName, secretID, secretContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditSecret", reflect.TypeOf((*MockISecrets)(nil).EditSecret), ctx, secretName, secretID, secretContent)
}

// GetSecretByID mocks base method.
func (m *MockISecrets) GetSecretByID(ctx context.Context, secretID int64) (models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretByID", ctx, secretID)
	ret0, _ := ret[0].(models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretByID indicates an expected call of GetSecretByID.
func (mr *MockISecretsMockRecorder) GetSecretByID(ctx, secretID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretByID", reflect.TypeOf((*MockISecrets)(nil).GetSecretByID), ctx, secretID)
}

// GetSecretTypes mocks base method.
func (m *MockISecrets) GetSecretTypes(ctx context.Context) ([]models.SecretType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretTypes", ctx)
	ret0, _ := ret[0].([]models.SecretType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretTypes indicates an expected call of GetSecretTypes.
func (mr *MockISecretsMockRecorder) GetSecretTypes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretTypes", reflect.TypeOf((*MockISecrets)(nil).GetSecretTypes), ctx)
}

// GetSecretsBySecretTypeID mocks base method.
func (m *MockISecrets) GetSecretsBySecretTypeID(ctx context.Context, secretTypeID models.SecretTypeID) ([]models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretsBySecretTypeID", ctx, secretTypeID)
	ret0, _ := ret[0].([]models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretsBySecretTypeID indicates an expected call of GetSecretsBySecretTypeID.
func (mr *MockISecretsMockRecorder) GetSecretsBySecretTypeID(ctx, secretTypeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretsBySecretTypeID", reflect.TypeOf((*MockISecrets)(nil).GetSecretsBySecretTypeID), ctx, secretTypeID)
}
