package application

import (
	"fmt"
	"go-user-auth-api/application/repository"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
)

type UserService interface {
	CreateUser(request domain.UserCreateRequest) error
	GetUserByEmail(email string) (*domain.UserDto, error)
	GetUserById(id int) (*domain.UserDto, error)
	UpdateUserById(id int, request domain.UserUpdateRequest) error
	DeleteUserById(id int) error
	GetUserRolesById(id int) ([]domain.RoleDto, error)
	GetUserRolePermissionsByUserIdAndRoleId(id int, roleId int) ([]domain.PermissionDto, error)
	GetAll() ([]domain.UserDto, error)
}

type UserServiceImp struct {
	userRepository repository.UserRepository
}

func NewUserServiceImp(userRepository repository.UserRepository) *UserServiceImp {
	return &UserServiceImp{userRepository: userRepository}
}

func (userService *UserServiceImp) CreateUser(request domain.UserCreateRequest) error {
	user := domain.User{
		Email:      request.Email,
		Name:       request.Name,
		Surname:    request.Surname,
		CreatedBy:  request.CreatedBy,
		ModifiedBy: request.ModifiedBy,
	}
	err := userService.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserServiceImp) GetUserById(id int) (*domain.UserDto, error) {
	user, err := userService.userRepository.GetUserById(id)
	if err != nil {
		return nil, errors.NotFound(fmt.Sprintf("User not found with given id %d", id))
	}

	userDto := domain.UserDto{
		Id:           user.Id,
		Email:        user.Email,
		Name:         user.Name,
		Surname:      user.Surname,
		CreatedBy:    user.CreatedBy,
		ModifiedBy:   user.ModifiedBy,
		CreatedDate:  user.CreatedDate,
		ModifiedDate: user.ModifiedDate,
	}

	return &userDto, nil
}

func (userService *UserServiceImp) GetUserByEmail(email string) (*domain.UserDto, error) {
	user, err := userService.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.NotFound(fmt.Sprintf("User not found with given email %s", email))
	}

	userDto := domain.UserDto{
		Id:           user.Id,
		Email:        user.Email,
		Name:         user.Name,
		Surname:      user.Surname,
		CreatedBy:    user.CreatedBy,
		ModifiedBy:   user.ModifiedBy,
		CreatedDate:  user.CreatedDate,
		ModifiedDate: user.ModifiedDate,
	}

	return &userDto, nil
}

func (userService *UserServiceImp) UpdateUserById(id int, request domain.UserUpdateRequest) error {
	var existingUser, err = userService.userRepository.GetUserById(id)
	if err != nil {
		return errors.NotFound(fmt.Sprintf("User not found with given id %d", id))
	}

	user := domain.User{
		Email:       request.Email,
		Name:        request.Name,
		Surname:     request.Surname,
		ModifiedBy:  request.ModifiedBy,
		CreatedBy:   existingUser.CreatedBy,
		CreatedDate: existingUser.CreatedDate,
	}
	return userService.userRepository.UpdateUserById(id, user)
}

func (userService *UserServiceImp) DeleteUserById(id int) error {
	return userService.userRepository.DeleteUserById(id)
}

func (userService *UserServiceImp) GetUserRolesById(id int) ([]domain.RoleDto, error) {
	roles, err := userService.userRepository.GetUserRolesById(id)
	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return []domain.RoleDto{}, nil
	}

	roleResponse := make([]domain.RoleDto, 0)
	for _, role := range roles {
		roleDto := domain.RoleDto{
			Id:           role.Id,
			Name:         role.Name,
			CreatedBy:    role.CreatedBy,
			ModifiedBy:   role.ModifiedBy,
			CreatedDate:  role.CreatedDate,
			ModifiedDate: role.ModifiedDate,
		}
		roleResponse = append(roleResponse, roleDto)
	}

	return roleResponse, nil
}

func (userService *UserServiceImp) GetUserRolePermissionsByUserIdAndRoleId(id int, roleId int) ([]domain.PermissionDto, error) {
	permissions, err := userService.userRepository.GetUserRolePermissionsByUserIdAndRoleId(id, roleId)
	if err != nil {
		return nil, err
	}

	if len(permissions) == 0 {
		return []domain.PermissionDto{}, nil
	}

	permissionResponse := make([]domain.PermissionDto, 0)
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

		permissionResponse = append(permissionResponse, permissionDto)
	}

	return permissionResponse, nil
}

func (userService *UserServiceImp) GetAll() ([]domain.UserDto, error) {
	users, err := userService.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	userDtos := make([]domain.UserDto, 0)
	for _, user := range users {
		userDto := domain.UserDto{
			Id:           user.Id,
			Email:        user.Email,
			Name:         user.Name,
			Surname:      user.Surname,
			CreatedBy:    user.CreatedBy,
			ModifiedBy:   user.ModifiedBy,
			CreatedDate:  user.CreatedDate,
			ModifiedDate: user.ModifiedDate,
		}
		userDtos = append(userDtos, userDto)
	}

	return userDtos, nil
}
