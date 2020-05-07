package application

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-user-auth-api/domain"
	repository "go-user-auth-api/mocks/application/repository"
	"testing"
	"time"
)

func TestApplicationsServiceImp_CreateApplication(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)

	defer controller.Finish()

	applicationCreateRequest := domain.ApplicationCreateRequest{
		Name:       "Pim UI",
		CreatedBy:  "enes.acikoglu@trendyol.com",
		ModifiedBy: "duygu.acikoglu@trendyol.com",
	}

	application := domain.Application{
		Name:       "Pim UI",
		CreatedBy:  "enes.acikoglu@trendyol.com",
		ModifiedBy: "duygu.acikoglu@trendyol.com",
	}

	applicationRepository := repository.NewMockApplicationRepository(controller)

	applicationRepository.EXPECT().CreateApplication(application).Return(nil)

	//Act
	applicationService := NewApplicationServiceImp(applicationRepository)
	err := applicationService.CreateApplication(applicationCreateRequest)

	//Assert
	assert.Nil(t, err)
}

func TestApplicationsServiceImp_GetApplicationById(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)

	defer controller.Finish()

	expectedApplication := &domain.Application{
		Id:           1,
		Name:         "PIM UI",
		CreatedBy:    "enes.acikoglu@trendyol.com",
		ModifiedBy:   "duygu.acikoglu@trendyol.com",
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}

	applicationRepository := repository.NewMockApplicationRepository(controller)

	applicationRepository.EXPECT().GetApplicationById(1).Return(expectedApplication, nil)

	//Act
	applicationService := NewApplicationServiceImp(applicationRepository)
	actualApplication, err := applicationService.GetApplicationById(1)

	//Assert
	assert.Nil(t, err)
	assert.EqualValues(t, expectedApplication, actualApplication)
}

func TestApplicationsServiceImp_GetApplicationById_ReturnNotFoundError_WhenApplicationNotFound(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)

	defer controller.Finish()

	applicationRepository := repository.NewMockApplicationRepository(controller)

	applicationRepository.EXPECT().GetApplicationById(1).
		Return(nil, errors.New("Application Not Found!"))

	//Act
	applicationService := NewApplicationServiceImp(applicationRepository)
	actualApplication, err := applicationService.GetApplicationById(1)

	//Assert
	assert.Error(t, err, "Application Not Found!")
	assert.Nil(t, actualApplication)
}

func TestApplicationsServiceImp_GetAll(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)

	defer controller.Finish()

	expectedApplications := []domain.Application{
		{
			Id:           1,
			Name:         "PIM UI",
			CreatedBy:    "enes.acikoglu@trendyol.com",
			ModifiedBy:   "duygu.acikoglu@trendyol.com",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		{
			Id:           2,
			Name:         "LOOKUP UI",
			CreatedBy:    "enes.acikoglu@trendyol.com",
			ModifiedBy:   "duygu.acikoglu@trendyol.com",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
	}

	applicationRepository := repository.NewMockApplicationRepository(controller)

	applicationRepository.EXPECT().FindAll().Return(expectedApplications, nil)

	//Act
	applicationService := NewApplicationServiceImp(applicationRepository)
	actualApplications, err := applicationService.GetAll()

	//Assert
	assert.Nil(t, err)
	assert.EqualValues(t, expectedApplications[0], actualApplications[0])
	assert.EqualValues(t, expectedApplications[1], actualApplications[1])
}

func TestApplicationsServiceImp_GetAll_ReturnEmptyArray_WhenApplicationsNotFound(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)

	defer controller.Finish()

	applicationRepository := repository.NewMockApplicationRepository(controller)

	expectedApplications := make([]domain.Application, 0)
	applicationRepository.EXPECT().FindAll().Return(expectedApplications, nil)

	//Act
	applicationService := NewApplicationServiceImp(applicationRepository)
	actualApplications, err := applicationService.GetAll()

	//Assert
	assert.Nil(t, err)
	assert.Empty(t, actualApplications)
}