package repository

import "go-user-auth-api/domain"

type RoleRepository interface {
	CreateRole(role domain.Role) error
	GetRoleById(id int) (*domain.Role, error)
	FindAll() ([]domain.Role, error)
	DeleteRoleById(id int) error
}
