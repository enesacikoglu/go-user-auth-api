package application

import (
	"go-user-auth-api/application/repository"
	"go-user-auth-api/domain"
)

type UserPermissionService interface {
	GetUserPermissionsByEmailAndAppId(email string, applicationId int) ([]domain.PermissionDto, error)
}

type UserPermissionServiceImp struct {
	userPermissionRepository repository.UserPermissionRepository
}

func NewUserPermissionServiceImp(userPermissionRepository repository.UserPermissionRepository) *UserPermissionServiceImp {
	return &UserPermissionServiceImp{userPermissionRepository: userPermissionRepository}
}

func (userPermissionServiceImp *UserPermissionServiceImp) GetUserPermissionsByEmailAndAppId(email string, applicationId int) ([]domain.PermissionDto, error) {
	permissions, err := userPermissionServiceImp.userPermissionRepository.GetUserPermissionByEmailAndApplicationId(email,applicationId)
	if err != nil {
		return nil, err
	}

	if len(permissions) == 0 {
		return []domain.PermissionDto{}, nil
	}

	permissionResponse := make([]domain.PermissionDto, 0)
	for _, permission := range permissions {
		permissionDto := domain.PermissionDto{
			Id:             permission.Id,
			Name:           permission.Name,
			CreatedBy:      permission.CreatedBy,
			ModifiedBy:     permission.ModifiedBy,
			CreatedDate:    permission.CreatedDate,
			ModifiedDate:   permission.ModifiedDate,
			ApplicationId:  permission.ApplicationId,
		}

		permissionResponse = append(permissionResponse, permissionDto)
	}

	return permissionResponse, nil
}
