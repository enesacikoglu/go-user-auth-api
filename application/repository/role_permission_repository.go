package repository

import "go-user-auth-api/domain"

type RolePermissionRepository interface {
	CreateRolePermission(role domain.RolePermission) error
	FindAll() ([]domain.RolePermission, error)
	DeleteByRoleIdAndPermissionIdAndApplicationId(roleId int, permissionId int, applicationId int) error
}
