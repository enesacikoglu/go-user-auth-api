package repository

import "go-user-auth-api/domain"

type ApplicationRepository interface {
	CreateApplication(application domain.Application) error
	GetApplicationById(id int) (*domain.Application, error)
	FindAll() ([]domain.Application, error)
}
