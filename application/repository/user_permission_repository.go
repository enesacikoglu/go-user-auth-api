package repository

import "go-user-auth-api/domain"

type UserPermissionRepository interface {
	GetUserPermissionByEmailAndApplicationId(email string, applicationId int) ([]domain.Permission, error)
}
