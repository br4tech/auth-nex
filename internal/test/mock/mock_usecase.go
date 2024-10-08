// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/br4tech/auth-nex/internal/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockICreateTenantUseCase is a mock of ICreateTenantUseCase interface.
type MockICreateTenantUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockICreateTenantUseCaseMockRecorder
}

// MockICreateTenantUseCaseMockRecorder is the mock recorder for MockICreateTenantUseCase.
type MockICreateTenantUseCaseMockRecorder struct {
	mock *MockICreateTenantUseCase
}

// NewMockICreateTenantUseCase creates a new mock instance.
func NewMockICreateTenantUseCase(ctrl *gomock.Controller) *MockICreateTenantUseCase {
	mock := &MockICreateTenantUseCase{ctrl: ctrl}
	mock.recorder = &MockICreateTenantUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICreateTenantUseCase) EXPECT() *MockICreateTenantUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockICreateTenantUseCase) Execute(tenant *domain.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", tenant)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockICreateTenantUseCaseMockRecorder) Execute(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockICreateTenantUseCase)(nil).Execute), tenant)
}

// MockIGetTenantByIdUseCase is a mock of IGetTenantByIdUseCase interface.
type MockIGetTenantByIdUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIGetTenantByIdUseCaseMockRecorder
}

// MockIGetTenantByIdUseCaseMockRecorder is the mock recorder for MockIGetTenantByIdUseCase.
type MockIGetTenantByIdUseCaseMockRecorder struct {
	mock *MockIGetTenantByIdUseCase
}

// NewMockIGetTenantByIdUseCase creates a new mock instance.
func NewMockIGetTenantByIdUseCase(ctrl *gomock.Controller) *MockIGetTenantByIdUseCase {
	mock := &MockIGetTenantByIdUseCase{ctrl: ctrl}
	mock.recorder = &MockIGetTenantByIdUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGetTenantByIdUseCase) EXPECT() *MockIGetTenantByIdUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIGetTenantByIdUseCase) Execute(id int) (*domain.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", id)
	ret0, _ := ret[0].(*domain.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIGetTenantByIdUseCaseMockRecorder) Execute(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIGetTenantByIdUseCase)(nil).Execute), id)
}

// MockIGetTenantByNameUseCase is a mock of IGetTenantByNameUseCase interface.
type MockIGetTenantByNameUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIGetTenantByNameUseCaseMockRecorder
}

// MockIGetTenantByNameUseCaseMockRecorder is the mock recorder for MockIGetTenantByNameUseCase.
type MockIGetTenantByNameUseCaseMockRecorder struct {
	mock *MockIGetTenantByNameUseCase
}

// NewMockIGetTenantByNameUseCase creates a new mock instance.
func NewMockIGetTenantByNameUseCase(ctrl *gomock.Controller) *MockIGetTenantByNameUseCase {
	mock := &MockIGetTenantByNameUseCase{ctrl: ctrl}
	mock.recorder = &MockIGetTenantByNameUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGetTenantByNameUseCase) EXPECT() *MockIGetTenantByNameUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIGetTenantByNameUseCase) Execute(name string) (*domain.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", name)
	ret0, _ := ret[0].(*domain.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIGetTenantByNameUseCaseMockRecorder) Execute(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIGetTenantByNameUseCase)(nil).Execute), name)
}

// MockIUpdateTenantUseCase is a mock of IUpdateTenantUseCase interface.
type MockIUpdateTenantUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIUpdateTenantUseCaseMockRecorder
}

// MockIUpdateTenantUseCaseMockRecorder is the mock recorder for MockIUpdateTenantUseCase.
type MockIUpdateTenantUseCaseMockRecorder struct {
	mock *MockIUpdateTenantUseCase
}

// NewMockIUpdateTenantUseCase creates a new mock instance.
func NewMockIUpdateTenantUseCase(ctrl *gomock.Controller) *MockIUpdateTenantUseCase {
	mock := &MockIUpdateTenantUseCase{ctrl: ctrl}
	mock.recorder = &MockIUpdateTenantUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUpdateTenantUseCase) EXPECT() *MockIUpdateTenantUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIUpdateTenantUseCase) Execute(tenant *domain.Tenant) (*domain.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", tenant)
	ret0, _ := ret[0].(*domain.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIUpdateTenantUseCaseMockRecorder) Execute(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIUpdateTenantUseCase)(nil).Execute), tenant)
}

// MockIGenerateTokenUseCase is a mock of IGenerateTokenUseCase interface.
type MockIGenerateTokenUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIGenerateTokenUseCaseMockRecorder
}

// MockIGenerateTokenUseCaseMockRecorder is the mock recorder for MockIGenerateTokenUseCase.
type MockIGenerateTokenUseCaseMockRecorder struct {
	mock *MockIGenerateTokenUseCase
}

// NewMockIGenerateTokenUseCase creates a new mock instance.
func NewMockIGenerateTokenUseCase(ctrl *gomock.Controller) *MockIGenerateTokenUseCase {
	mock := &MockIGenerateTokenUseCase{ctrl: ctrl}
	mock.recorder = &MockIGenerateTokenUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGenerateTokenUseCase) EXPECT() *MockIGenerateTokenUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIGenerateTokenUseCase) Execute(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIGenerateTokenUseCaseMockRecorder) Execute(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIGenerateTokenUseCase)(nil).Execute), id)
}

// MockICreateUserUsecase is a mock of ICreateUserUsecase interface.
type MockICreateUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockICreateUserUsecaseMockRecorder
}

// MockICreateUserUsecaseMockRecorder is the mock recorder for MockICreateUserUsecase.
type MockICreateUserUsecaseMockRecorder struct {
	mock *MockICreateUserUsecase
}

// NewMockICreateUserUsecase creates a new mock instance.
func NewMockICreateUserUsecase(ctrl *gomock.Controller) *MockICreateUserUsecase {
	mock := &MockICreateUserUsecase{ctrl: ctrl}
	mock.recorder = &MockICreateUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICreateUserUsecase) EXPECT() *MockICreateUserUsecaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockICreateUserUsecase) Execute(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockICreateUserUsecaseMockRecorder) Execute(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockICreateUserUsecase)(nil).Execute), user)
}

// MockIDeleteUserUsecase is a mock of IDeleteUserUsecase interface.
type MockIDeleteUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIDeleteUserUsecaseMockRecorder
}

// MockIDeleteUserUsecaseMockRecorder is the mock recorder for MockIDeleteUserUsecase.
type MockIDeleteUserUsecaseMockRecorder struct {
	mock *MockIDeleteUserUsecase
}

// NewMockIDeleteUserUsecase creates a new mock instance.
func NewMockIDeleteUserUsecase(ctrl *gomock.Controller) *MockIDeleteUserUsecase {
	mock := &MockIDeleteUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIDeleteUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDeleteUserUsecase) EXPECT() *MockIDeleteUserUsecaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIDeleteUserUsecase) Execute(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockIDeleteUserUsecaseMockRecorder) Execute(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIDeleteUserUsecase)(nil).Execute), id)
}

// MockIGetUserByEmailUseCase is a mock of IGetUserByEmailUseCase interface.
type MockIGetUserByEmailUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIGetUserByEmailUseCaseMockRecorder
}

// MockIGetUserByEmailUseCaseMockRecorder is the mock recorder for MockIGetUserByEmailUseCase.
type MockIGetUserByEmailUseCaseMockRecorder struct {
	mock *MockIGetUserByEmailUseCase
}

// NewMockIGetUserByEmailUseCase creates a new mock instance.
func NewMockIGetUserByEmailUseCase(ctrl *gomock.Controller) *MockIGetUserByEmailUseCase {
	mock := &MockIGetUserByEmailUseCase{ctrl: ctrl}
	mock.recorder = &MockIGetUserByEmailUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGetUserByEmailUseCase) EXPECT() *MockIGetUserByEmailUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIGetUserByEmailUseCase) Execute(email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIGetUserByEmailUseCaseMockRecorder) Execute(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIGetUserByEmailUseCase)(nil).Execute), email)
}

// MockIUpdateUserUsecase is a mock of IUpdateUserUsecase interface.
type MockIUpdateUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUpdateUserUsecaseMockRecorder
}

// MockIUpdateUserUsecaseMockRecorder is the mock recorder for MockIUpdateUserUsecase.
type MockIUpdateUserUsecaseMockRecorder struct {
	mock *MockIUpdateUserUsecase
}

// NewMockIUpdateUserUsecase creates a new mock instance.
func NewMockIUpdateUserUsecase(ctrl *gomock.Controller) *MockIUpdateUserUsecase {
	mock := &MockIUpdateUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIUpdateUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUpdateUserUsecase) EXPECT() *MockIUpdateUserUsecaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIUpdateUserUsecase) Execute(user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIUpdateUserUsecaseMockRecorder) Execute(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIUpdateUserUsecase)(nil).Execute), user)
}
