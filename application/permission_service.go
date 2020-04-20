package application

import (
	"fmt"
	"go-user-auth-api/application/repository"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
)

type PermissionService interface {
	CreatePermission(request domain.PermissionCreateRequest) error
	GetPermissionById(id int) (*domain.PermissionDto, error)
	GetAll() ([]domain.PermissionDto, error)
	DeletePermissionById(id int) error
}

type PermissionServiceImp struct {
	repository repository.PermissionRepository
}

func NewPermissionServiceImp(repository repository.PermissionRepository) *PermissionServiceImp {
	return &PermissionServiceImp{repository: repository}
}

func (permissionService *PermissionServiceImp) CreatePermission(request domain.PermissionCreateRequest) error {
	permission := domain.Permission{
		Name:       request.Name,
		CreatedBy:  request.CreatedBy,
		ModifiedBy: request.ModifiedBy,
	}
	err := permissionService.repository.CreatePermission(permission)
	if err != nil {
		return err
	}
	return nil
}

func (permissionService *PermissionServiceImp) GetPermissionById(id int) (*domain.PermissionDto, error) {
	permission, err := permissionService.repository.GetPermissionById(id)
	if err != nil {
		return nil, errors.NotFound(fmt.Sprintf("Permission not found with given id %d", id))
	}
	permissionDto := domain.PermissionDto{
		Id:            permission.Id,
		Name:          permission.Name,
		CreatedBy:     permission.CreatedBy,
		ModifiedBy:    permission.ModifiedBy,
		CreatedDate:   permission.CreatedDate,
		ModifiedDate:  permission.ModifiedDate,
		ApplicationId: permission.ApplicationId,
	}

	return &permissionDto, nil
}

func (permissionService *PermissionServiceImp) DeletePermissionById(id int) error {
	return permissionService.repository.DeletePermissionById(id)
}

func (permissionService *PermissionServiceImp) GetAll() ([]domain.PermissionDto, error) {
	permissions, err := permissionService.repository.FindAll()
	if err != nil {
		return nil, err
	}

	permissionDtos := make([]domain.PermissionDto, 0)
	for _, permission := range permissions {
		permissionDto := domain.PermissionDto{
			Id:            permission.Id,
			Name:          permission.Name,
			CreatedBy:     permission.CreatedBy,
			ModifiedBy:    permission.ModifiedBy,
			CreatedDate:   permission.CreatedDate,
			ModifiedDate:  permission.ModifiedDate,
			ApplicationId: permission.ApplicationId,
		}
		permissionDtos = append(permissionDtos, permissionDto)
	}

	return permissionDtos, nil
}
