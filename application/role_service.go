package application

import (
	"fmt"
	"go-user-auth-api/application/repository"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
)

type RoleService interface {
	CreateRole(request domain.CreateRoleCommand) (int, error)
	GetRoleById(id int) (*domain.RoleDto, error)
	GetAll() ([]domain.RoleDto, error)
	DeleteRoleById(id int) error
}

type RoleServiceImp struct {
	repository repository.RoleRepository
}

func NewRoleServiceImp(repository repository.RoleRepository) *RoleServiceImp {
	return &RoleServiceImp{repository: repository}
}

func (roleService *RoleServiceImp) CreateRole(request domain.CreateRoleCommand) (int, error) {
	role := domain.Role{
		Name:       request.Name,
		CreatedBy:  request.CreatedBy,
		ModifiedBy: request.ModifiedBy,
	}
	id, err := roleService.repository.CreateRole(role)
	if err != nil && id != 0 {
		return id, err
	}
	return id, nil
}

func (roleService *RoleServiceImp) GetRoleById(id int) (*domain.RoleDto, error) {
	role, err := roleService.repository.GetRoleById(id)
	if err != nil {
		return nil, errors.NotFound(fmt.Sprintf("Role not found with given id %d", id))
	}

	roleDto := domain.RoleDto{
		Id:           role.Id,
		Name:         role.Name,
		CreatedBy:    role.CreatedBy,
		ModifiedBy:   role.ModifiedBy,
		CreatedDate:  role.CreatedDate,
		ModifiedDate: role.ModifiedDate,
	}

	return &roleDto, nil
}

func (roleService *RoleServiceImp) DeleteRoleById(id int) error {
	return roleService.repository.DeleteRoleById(id)
}

func (roleService *RoleServiceImp) GetAll() ([]domain.RoleDto, error) {
	roles, err := roleService.repository.FindAll()
	if err != nil {
		return nil, err
	}

	roleDtos := make([]domain.RoleDto, 0)
	for _, role := range roles {
		roleDto := domain.RoleDto{
			Id:           role.Id,
			Name:         role.Name,
			CreatedBy:    role.CreatedBy,
			ModifiedBy:   role.ModifiedBy,
			CreatedDate:  role.CreatedDate,
			ModifiedDate: role.ModifiedDate,
		}
		roleDtos = append(roleDtos, roleDto)
	}

	return roleDtos, nil
}
