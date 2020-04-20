package application

import (
	"fmt"
	"go-user-auth-api/application/repository"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
)

type ApplicationsService interface {
	CreateApplication(request domain.ApplicationCreateRequest) error
	GetApplicationById(id int) (*domain.ApplicationDto, error)
	GetAll() ([]domain.ApplicationDto, error)
}

type ApplicationsServiceImp struct {
	repository repository.ApplicationRepository
}

func NewApplicationServiceImp(repository repository.ApplicationRepository) *ApplicationsServiceImp {
	return &ApplicationsServiceImp{repository: repository}
}

func (applicationsService *ApplicationsServiceImp) CreateApplication(request domain.ApplicationCreateRequest) error {
	application := domain.Application{
		Name:       request.Name,
		CreatedBy:  request.CreatedBy,
		ModifiedBy: request.ModifiedBy,
	}
	err := applicationsService.repository.CreateApplication(application)
	if err != nil {
		return err
	}
	return nil
}

func (applicationsService *ApplicationsServiceImp) GetApplicationById(id int) (*domain.ApplicationDto, error) {
	application, err := applicationsService.repository.GetApplicationById(id)
	if err != nil {
		return nil, errors.NotFound(fmt.Sprintf("Application not found with given id %d", id))
	}

	applicationDto := domain.ApplicationDto{
		Id:           application.Id,
		Name:         application.Name,
		CreatedBy:    application.CreatedBy,
		ModifiedBy:   application.ModifiedBy,
		CreatedDate:  application.CreatedDate,
		ModifiedDate: application.ModifiedDate,
	}

	return &applicationDto, nil
}

func (applicationsService *ApplicationsServiceImp) GetAll() ([]domain.ApplicationDto, error) {
	applications, err := applicationsService.repository.FindAll()
	if err != nil {
		return nil, err
	}

	applicationDtos := make([]domain.ApplicationDto, 0)
	for _, application := range applications {
		applicationDto := domain.ApplicationDto{
			Id:           application.Id,
			Name:         application.Name,
			CreatedBy:    application.CreatedBy,
			ModifiedBy:   application.ModifiedBy,
			CreatedDate:  application.CreatedDate,
			ModifiedDate: application.ModifiedDate,
		}
		applicationDtos = append(applicationDtos, applicationDto)
	}
	return applicationDtos, nil
}
