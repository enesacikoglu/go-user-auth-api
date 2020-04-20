package repository

import "go-user-auth-api/domain"

type UserRoleRepository interface {
	CreateUserRole(role domain.UserRole) error
	FindAll() ([]domain.UserRole, error)
	DeleteByUserIdAndRoleId(userId int, roleId int) error
}
