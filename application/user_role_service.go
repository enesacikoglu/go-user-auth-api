package application

import (
	"fmt"
	"go-user-auth-api/application/repository"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
)

type UserRoleService interface {
	CreateUserRole(request domain.UserRoleCreateRequest) error
	GetAll() ([]domain.UserRoleDto, error)
	DeleteByUserIdAndRoleId(userId int, roleId int) error
	GetUserRoleWithPermissionsByUserId(id int) (*domain.UserDto, error)
}

type UserRoleServiceImp struct {
	userRepository     repository.UserRepository
	roleRepository     repository.RoleRepository
	userRoleRepository repository.UserRoleRepository
}

func NewUserRoleServiceImp(userRepository repository.UserRepository,
	roleRepository repository.RoleRepository, userRoleRepository repository.UserRoleRepository) *UserRoleServiceImp {
	return &UserRoleServiceImp{userRepository: userRepository,
		roleRepository:     roleRepository,
		userRoleRepository: userRoleRepository}
}

func (userRoleService *UserRoleServiceImp) CreateUserRole(request domain.UserRoleCreateRequest) error {
	userRole := domain.UserRole{
		UserId:    request.UserId,
		RoleId:    request.RoleId,
		CreatedBy: request.CreatedBy,
	}

	user, err := userRoleService.userRepository.GetUserById(request.UserId)
	if err != nil && user == nil {
		return errors.NotFound(fmt.Sprintf("User Not Found With Given Id %d", request.UserId))
	}

	role, err := userRoleService.roleRepository.GetRoleById(request.RoleId)
	if err != nil && role == nil {
		return errors.NotFound(fmt.Sprintf("Role Not Found With Given Id %d", request.RoleId))
	}

	err = userRoleService.userRoleRepository.CreateUserRole(userRole)
	if err != nil {
		return err
	}
	return nil
}

func (userRoleService *UserRoleServiceImp) GetAll() ([]domain.UserRoleDto, error) {
	userRoles, err := userRoleService.userRoleRepository.FindAll()

	if err != nil {
		return nil, err
	}

	roleDtos := make([]domain.UserRoleDto, 0)
	for _, role := range userRoles {
		roleDto := domain.UserRoleDto{
			Id:          role.Id,
			UserId:      role.UserId,
			RoleId:      role.RoleId,
			CreatedBy:   role.CreatedBy,
			CreatedDate: role.CreatedDate,
		}
		roleDtos = append(roleDtos, roleDto)
	}

	return roleDtos, nil
}

func (userRoleService *UserRoleServiceImp) DeleteByUserIdAndRoleId(userId int, roleId int) error {
	user, err := userRoleService.userRepository.GetUserById(userId)
	if err != nil && user == nil {
		return errors.NotFound(fmt.Sprintf("User Not Found With Given Id %d", userId))
	}

	role, err := userRoleService.roleRepository.GetRoleById(roleId)
	if err != nil && role == nil {
		return errors.NotFound(fmt.Sprintf("Role Not Found With Given Id %d", roleId))
	}

	return userRoleService.userRoleRepository.DeleteByUserIdAndRoleId(userId, roleId)
}

func (userRoleService *UserRoleServiceImp) GetUserRoleWithPermissionsByUserId(id int) (*domain.UserDto, error) {
	userRolePermissions, err := userRoleService.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	if userRolePermissions == nil {
		return nil, errors.NotFound(fmt.Sprintf("Could not find user role with permission with given userId %d", id))
	}

	user := domain.UserDto{
		Id:           userRolePermissions[0].UserId,
		Email:        userRolePermissions[0].Email,
		Name:         userRolePermissions[0].UserName,
		Surname:      userRolePermissions[0].Surname,
		CreatedBy:    userRolePermissions[0].UserCreatedBy,
		ModifiedBy:   userRolePermissions[0].UserModifiedBy,
		CreatedDate:  userRolePermissions[0].UserCreatedDate,
		ModifiedDate: userRolePermissions[0].UserModifiedDate,
	}

	roleMap := make(map[int]domain.RoleDto)
	for i := 0; i < len(userRolePermissions); i++ {
		role := domain.RoleDto{
			Id:           userRolePermissions[i].RoleId,
			Name:         userRolePermissions[i].RoleName,
			CreatedBy:    userRolePermissions[i].RoleCreatedBy,
			ModifiedBy:   userRolePermissions[i].RoleModifiedBy,
			CreatedDate:  userRolePermissions[i].RoleCreatedDate,
			ModifiedDate: userRolePermissions[i].RoleModifiedDate,
		}
		roleMap[userRolePermissions[i].RoleId] = role
	}

	roles := make([]domain.RoleDto, 0)
	for _, role := range roleMap {
		roles = append(roles, role)
	}

	//Map Role Permissions
	rolePermMap := make(map[int][]domain.PermissionDto)
	for i := 0; i < len(userRolePermissions); i++ {
		permission := domain.PermissionDto{
			Id:            userRolePermissions[i].PermissionId,
			Name:          userRolePermissions[i].PermissionName,
			CreatedBy:     userRolePermissions[i].PermissionCreatedBy,
			ModifiedBy:    userRolePermissions[i].PermissionModifiedBy,
			CreatedDate:   userRolePermissions[i].PermissionCreatedDate,
			ModifiedDate:  userRolePermissions[i].PermissionModifiedDate,
			ApplicationId: userRolePermissions[i].ApplicationId,
		}

		for _, role := range roleMap {
			if role.Id == userRolePermissions[i].RoleId {
				perms := rolePermMap[userRolePermissions[i].RoleId]
				perms = append(perms, permission)
				rolePermMap[userRolePermissions[i].RoleId] = perms
			}
		}
	}

	user.Roles = roles
	for roleId, permissions := range rolePermMap {
		for i, role := range user.Roles {
			if role.Id == roleId {
				user.Roles[i].Permissions = permissions
			}
		}
	}

	return &user, nil
}
