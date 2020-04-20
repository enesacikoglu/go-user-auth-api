package application

import (

	"fmt"
	"go-user-auth-api/application/repository"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
)

type RolePermissionService interface {
	CreateRolePermission(request domain.RolePermissionCreateRequest) error
	GetAll() ([]domain.RolePermissionDto, error)
	DeleteByRoleIdAndPermissionIdAndApplicationId(roleId int, permissionId int, applicationId int) error
}

type RolePermissionServiceImp struct {
	roleRepository           repository.RoleRepository
	permissionRepository     repository.PermissionRepository
	rolePermissionRepository repository.RolePermissionRepository
	applicationRepository    repository.ApplicationRepository
}

func NewRolePermissionServiceImp(roleRepository repository.RoleRepository,
	permissionRepository repository.PermissionRepository,
	rolePermissionRepository repository.RolePermissionRepository,
	applicationRepository repository.ApplicationRepository) *RolePermissionServiceImp {
	return &RolePermissionServiceImp{roleRepository: roleRepository,
		permissionRepository:     permissionRepository,
		rolePermissionRepository: rolePermissionRepository,
		applicationRepository:    applicationRepository}
}

func (rolePermissionService *RolePermissionServiceImp) CreateRolePermission(request domain.RolePermissionCreateRequest) error {
	rolePermission := domain.RolePermission{
		RoleId:        request.RoleId,
		PermissionId:  request.PermissionId,
		CreatedBy:     request.CreatedBy,
		ApplicationId: request.ApplicationId,
	}

	user, err := rolePermissionService.roleRepository.GetRoleById(request.RoleId)
	if err != nil && user == nil {
		return errors.NotFound(fmt.Sprintf("Role Not Found With Given Id %d", request.RoleId))
	}

	application, err := rolePermissionService.applicationRepository.GetApplicationById(request.ApplicationId)
	if err != nil && application == nil {
		return errors.NotFound(fmt.Sprintf("Application Not Found With Given Id %d", request.ApplicationId))
	}

	role, err := rolePermissionService.permissionRepository.GetPermissionById(request.PermissionId)
	if err != nil && role == nil {
		return errors.NotFound(fmt.Sprintf("Permission Not Found With Given Id %d", request.PermissionId))
	}

	err = rolePermissionService.rolePermissionRepository.CreateRolePermission(rolePermission)
	if err != nil {
		return err
	}
	return nil
}

func (rolePermissionService *RolePermissionServiceImp) GetAll() ([]domain.RolePermissionDto, error) {
	rolePermissions, err := rolePermissionService.rolePermissionRepository.FindAll()

	if err != nil {
		return nil, err
	}

	permissionDtos := make([]domain.RolePermissionDto, 0)
	for _, permission := range rolePermissions {
		permissionDto := domain.RolePermissionDto{
			Id:            permission.Id,
			RoleId:        permission.RoleId,
			PermissionId:  permission.PermissionId,
			CreatedBy:     permission.CreatedBy,
			CreatedDate:   permission.CreatedDate,
			ApplicationId: permission.ApplicationId,
		}
		permissionDtos = append(permissionDtos, permissionDto)
	}
	return permissionDtos, nil
}

func (rolePermissionService *RolePermissionServiceImp) DeleteByRoleIdAndPermissionIdAndApplicationId(roleId int, permissionId int, applicationId int) error {
	user, err := rolePermissionService.roleRepository.GetRoleById(roleId)
	if err != nil && user == nil {
		return errors.NotFound(fmt.Sprintf("Role Not Found With Given Id %d", roleId))
	}

	permission, err := rolePermissionService.permissionRepository.GetPermissionById(permissionId)
	if err != nil && permission == nil {
		return errors.NotFound(fmt.Sprintf("Permission Not Found With Given Id %d", permissionId))
	}

	application, err := rolePermissionService.applicationRepository.GetApplicationById(applicationId)
	if err != nil && application == nil {
		return errors.NotFound(fmt.Sprintf("Application Not Found With Given Id %d", applicationId))
	}

	return rolePermissionService.rolePermissionRepository.DeleteByRoleIdAndPermissionIdAndApplicationId(roleId, permissionId, applicationId)
}
