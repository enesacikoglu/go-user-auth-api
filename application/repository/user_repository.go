package repository

import "go-user-auth-api/domain"

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserById(id int) (*domain.User, error)
	UpdateUserById(id int, user domain.User) error
	DeleteUserById(id int) error
	GetUser(id int) (domain.UserRolePermissions, error)
	GetUserRolesById(id int) ([]domain.Role, error)
	GetUserRolePermissionsByUserIdAndRoleId(id int, roleId int) ([]domain.Permission, error)
	FindAll() ([]domain.User, error)
}
