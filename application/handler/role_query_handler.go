package handler

import (
	"go-user-auth-api/application"
)

type GetRoleByIdQuery struct {
	Id int
}

type GetAllRolesQuery struct {
}

type RoleQueryHandler struct {
	roleService application.RoleService
}

func NewRoleQueryHandler(roleService application.RoleService) *RoleQueryHandler {
	return &RoleQueryHandler{
		roleService: roleService,
	}
}

func (handler *RoleQueryHandler) GetRoleByIdQueryHandler(query interface{}) (interface{}, error) {
	roleByIdQuery := query.(GetRoleByIdQuery)
	role, err := handler.roleService.GetRoleById(roleByIdQuery.Id)
	if err != nil {
		return nil, err
	}
	return role, nil
}


func (handler *RoleQueryHandler) GetAllRolesQueryQueryHandler(query interface{}) (interface{}, error) {
	roles, err := handler.roleService.GetAll()
	if err != nil {
		return nil, err
	}
	return roles, nil
}