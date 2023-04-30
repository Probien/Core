package application

import (
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/domain/repository"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/auth"
	"net/url"
)

type EmployeeInteractor struct {
	repository repository.IEmployeeRepository
}

func NewEmployeeInteractor(repository repository.IEmployeeRepository) EmployeeInteractor {
	return EmployeeInteractor{
		repository: repository,
	}
}

func (e *EmployeeInteractor) Login(loginCredential auth.LoginCredential) (*domain.Employee, error) {
	return e.repository.Login(loginCredential)
}

func (e *EmployeeInteractor) GetByEmail(email string) (*domain.Employee, error) {
	return e.repository.GetByEmail(email)
}

func (e *EmployeeInteractor) GetAll(params url.Values) (*[]domain.Employee, map[string]interface{}, error) {
	return e.repository.GetAll(params)
}

func (e *EmployeeInteractor) Create(employeeDto *domain.Employee, userSessionId int) (*domain.Employee, error) {
	return e.repository.Create(employeeDto, userSessionId)
}

func (e *EmployeeInteractor) Update(id int, employeeDto map[string]interface{}, userSessionId int) (*domain.Employee, error) {
	return e.repository.Update(id, employeeDto, userSessionId)
}
