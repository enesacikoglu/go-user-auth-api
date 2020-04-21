package handler

import (
	"go-user-auth-api/application"
	"go-user-auth-api/domain"
)

type RoleCommandHandler struct {
	roleService application.RoleService
}

func NewRoleCommandHandler(roleService application.RoleService) *RoleCommandHandler {
	return &RoleCommandHandler{
		roleService: roleService,
	}
}

func (handler *RoleCommandHandler) CreateRoleCommandHandler(command interface{}) (interface{}, error) {
	createRoleCommand := command.(domain.CreateRoleCommand)
	id, err := handler.roleService.CreateRole(createRoleCommand)
	if err != nil {
		return nil, err
	}
	return domain.RoleCreatedEvent{Id: id}, nil
}

func (handler *RoleCommandHandler) DeleteRoleCommandHandler(command interface{}) (interface{}, error) {
	deleteRoleCommand := command.(domain.DeleteRoleCommand)
	err := handler.roleService.DeleteRoleById(deleteRoleCommand.Id)
	if err != nil {
		return nil, err
	}
	return domain.RoleDeletedEvent{Id: deleteRoleCommand.Id}, nil
}
