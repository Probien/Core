package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type EmployeeInteractor struct {
}

func (EI *EmployeeInteractor) Login(loginCredential auth.LoginCredential) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.Login(loginCredential)
}

func (EI *EmployeeInteractor) GetByEmail(email string) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.GetByEmail(email)
}

func (EI *EmployeeInteractor) GetAll(params url.Values) (*[]domain.Employee, map[string]interface{}, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.GetAll(params)
}

func (EI *EmployeeInteractor) Create(employeeDto *domain.Employee, userSessionId int) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.Create(employeeDto, userSessionId)
}

func (EI *EmployeeInteractor) Update(id int, employeeDto map[string]interface{}, userSessionId int) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.Update(id, employeeDto, userSessionId)
}
