package repository

import "go-user-auth-api/domain"

type PermissionRepository interface {
	CreatePermission(permission domain.Permission) error
	GetPermissionById(id int) (*domain.Permission, error)
	FindAll() ([]domain.Permission, error)
	DeletePermissionById(id int) error
}
