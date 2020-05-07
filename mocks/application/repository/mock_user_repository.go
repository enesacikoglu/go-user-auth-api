// Code generated by MockGen. DO NOT EDIT.
// Source: application/repository/user_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	domain "go-user-auth-api/domain"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserRepository) CreateUser(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// GetUserByEmail mocks base method
func (m *MockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), email)
}

// GetUserById mocks base method
func (m *MockUserRepository) GetUserById(id int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockUserRepositoryMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepository)(nil).GetUserById), id)
}

// UpdateUserById mocks base method
func (m *MockUserRepository) UpdateUserById(id int, user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserById", id, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserById indicates an expected call of UpdateUserById
func (mr *MockUserRepositoryMockRecorder) UpdateUserById(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserById), id, user)
}

// DeleteUserById mocks base method
func (m *MockUserRepository) DeleteUserById(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById
func (mr *MockUserRepositoryMockRecorder) DeleteUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUserRepository)(nil).DeleteUserById), id)
}

// GetUser mocks base method
func (m *MockUserRepository) GetUser(id int) (domain.UserRolePermissions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(domain.UserRolePermissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserRepositoryMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), id)
}

// GetUserRolesById mocks base method
func (m *MockUserRepository) GetUserRolesById(id int) ([]domain.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRolesById", id)
	ret0, _ := ret[0].([]domain.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRolesById indicates an expected call of GetUserRolesById
func (mr *MockUserRepositoryMockRecorder) GetUserRolesById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRolesById", reflect.TypeOf((*MockUserRepository)(nil).GetUserRolesById), id)
}

// GetUserRolePermissionsByUserIdAndRoleId mocks base method
func (m *MockUserRepository) GetUserRolePermissionsByUserIdAndRoleId(id, roleId int) ([]domain.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRolePermissionsByUserIdAndRoleId", id, roleId)
	ret0, _ := ret[0].([]domain.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRolePermissionsByUserIdAndRoleId indicates an expected call of GetUserRolePermissionsByUserIdAndRoleId
func (mr *MockUserRepositoryMockRecorder) GetUserRolePermissionsByUserIdAndRoleId(id, roleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRolePermissionsByUserIdAndRoleId", reflect.TypeOf((*MockUserRepository)(nil).GetUserRolePermissionsByUserIdAndRoleId), id, roleId)
}

// FindAll mocks base method
func (m *MockUserRepository) FindAll() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockUserRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserRepository)(nil).FindAll))
}