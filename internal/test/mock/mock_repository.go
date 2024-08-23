// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/br4tech/auth-nex/internal/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockICompanyRepository is a mock of ICompanyRepository interface.
type MockICompanyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICompanyRepositoryMockRecorder
}

// MockICompanyRepositoryMockRecorder is the mock recorder for MockICompanyRepository.
type MockICompanyRepositoryMockRecorder struct {
	mock *MockICompanyRepository
}

// NewMockICompanyRepository creates a new mock instance.
func NewMockICompanyRepository(ctrl *gomock.Controller) *MockICompanyRepository {
	mock := &MockICompanyRepository{ctrl: ctrl}
	mock.recorder = &MockICompanyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICompanyRepository) EXPECT() *MockICompanyRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICompanyRepository) Create(company *domain.Company) (*domain.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", company)
	ret0, _ := ret[0].(*domain.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICompanyRepositoryMockRecorder) Create(company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICompanyRepository)(nil).Create), company)
}

// Delete mocks base method.
func (m *MockICompanyRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICompanyRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICompanyRepository)(nil).Delete), id)
}

// FindById mocks base method.
func (m *MockICompanyRepository) FindById(id int) (*domain.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockICompanyRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockICompanyRepository)(nil).FindById), id)
}

// Update mocks base method.
func (m *MockICompanyRepository) Update(company *domain.Company) (*domain.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", company)
	ret0, _ := ret[0].(*domain.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockICompanyRepositoryMockRecorder) Update(company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockICompanyRepository)(nil).Update), company)
}

// MockIPermissionRepository is a mock of IPermissionRepository interface.
type MockIPermissionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPermissionRepositoryMockRecorder
}

// MockIPermissionRepositoryMockRecorder is the mock recorder for MockIPermissionRepository.
type MockIPermissionRepositoryMockRecorder struct {
	mock *MockIPermissionRepository
}

// NewMockIPermissionRepository creates a new mock instance.
func NewMockIPermissionRepository(ctrl *gomock.Controller) *MockIPermissionRepository {
	mock := &MockIPermissionRepository{ctrl: ctrl}
	mock.recorder = &MockIPermissionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPermissionRepository) EXPECT() *MockIPermissionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIPermissionRepository) Create(permission *domain.Permission) (*domain.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", permission)
	ret0, _ := ret[0].(*domain.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIPermissionRepositoryMockRecorder) Create(permission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIPermissionRepository)(nil).Create), permission)
}

// Delete mocks base method.
func (m *MockIPermissionRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIPermissionRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIPermissionRepository)(nil).Delete), id)
}

// FindById mocks base method.
func (m *MockIPermissionRepository) FindById(id int) (*domain.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIPermissionRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIPermissionRepository)(nil).FindById), id)
}

// Update mocks base method.
func (m *MockIPermissionRepository) Update(permission *domain.Permission) (*domain.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", permission)
	ret0, _ := ret[0].(*domain.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIPermissionRepositoryMockRecorder) Update(permission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPermissionRepository)(nil).Update), permission)
}

// MockIProfileRepository is a mock of IProfileRepository interface.
type MockIProfileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProfileRepositoryMockRecorder
}

// MockIProfileRepositoryMockRecorder is the mock recorder for MockIProfileRepository.
type MockIProfileRepositoryMockRecorder struct {
	mock *MockIProfileRepository
}

// NewMockIProfileRepository creates a new mock instance.
func NewMockIProfileRepository(ctrl *gomock.Controller) *MockIProfileRepository {
	mock := &MockIProfileRepository{ctrl: ctrl}
	mock.recorder = &MockIProfileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProfileRepository) EXPECT() *MockIProfileRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProfileRepository) Create(profile *domain.Profile) (*domain.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", profile)
	ret0, _ := ret[0].(*domain.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIProfileRepositoryMockRecorder) Create(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProfileRepository)(nil).Create), profile)
}

// Delete mocks base method.
func (m *MockIProfileRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIProfileRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIProfileRepository)(nil).Delete), id)
}

// FindById mocks base method.
func (m *MockIProfileRepository) FindById(id int) (*domain.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIProfileRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIProfileRepository)(nil).FindById), id)
}

// Upate mocks base method.
func (m *MockIProfileRepository) Upate(profile *domain.Profile) (*domain.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upate", profile)
	ret0, _ := ret[0].(*domain.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upate indicates an expected call of Upate.
func (mr *MockIProfileRepositoryMockRecorder) Upate(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upate", reflect.TypeOf((*MockIProfileRepository)(nil).Upate), profile)
}

// MockITenantRepository is a mock of ITenantRepository interface.
type MockITenantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITenantRepositoryMockRecorder
}

// MockITenantRepositoryMockRecorder is the mock recorder for MockITenantRepository.
type MockITenantRepositoryMockRecorder struct {
	mock *MockITenantRepository
}

// NewMockITenantRepository creates a new mock instance.
func NewMockITenantRepository(ctrl *gomock.Controller) *MockITenantRepository {
	mock := &MockITenantRepository{ctrl: ctrl}
	mock.recorder = &MockITenantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITenantRepository) EXPECT() *MockITenantRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockITenantRepository) Create(tenant *domain.Tenant) (*domain.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", tenant)
	ret0, _ := ret[0].(*domain.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockITenantRepositoryMockRecorder) Create(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITenantRepository)(nil).Create), tenant)
}

// Delete mocks base method.
func (m *MockITenantRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITenantRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITenantRepository)(nil).Delete), id)
}

// FindById mocks base method.
func (m *MockITenantRepository) FindById(id int) (*domain.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockITenantRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockITenantRepository)(nil).FindById), id)
}

// FindByName mocks base method.
func (m *MockITenantRepository) FindByName(name string) (*domain.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].(*domain.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockITenantRepositoryMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockITenantRepository)(nil).FindByName), name)
}

// Update mocks base method.
func (m *MockITenantRepository) Update(tenant *domain.Tenant) (*domain.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", tenant)
	ret0, _ := ret[0].(*domain.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockITenantRepositoryMockRecorder) Update(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITenantRepository)(nil).Update), tenant)
}

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepository) Create(user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepository)(nil).Create), user)
}

// Delete mocks base method.
func (m *MockIUserRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserRepository)(nil).Delete), id)
}

// FindBy mocks base method.
func (m *MockIUserRepository) FindBy(filter map[string]interface{}) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBy", filter)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBy indicates an expected call of FindBy.
func (mr *MockIUserRepositoryMockRecorder) FindBy(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBy", reflect.TypeOf((*MockIUserRepository)(nil).FindBy), filter)
}

// FindByEmail mocks base method.
func (m *MockIUserRepository) FindByEmail(email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockIUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockIUserRepository)(nil).FindByEmail), email)
}

// FindById mocks base method.
func (m *MockIUserRepository) FindById(id int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIUserRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIUserRepository)(nil).FindById), id)
}

// Update mocks base method.
func (m *MockIUserRepository) Update(user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepository)(nil).Update), user)
}
