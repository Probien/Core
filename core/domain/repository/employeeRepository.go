package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
)

type IEmployeeRepository interface {
	Login(loginCredential auth.LoginCredential) (*domain.Employee, error)
	GetByEmail(email string) (*domain.Employee, error)
	GetAll(params url.Values) (*[]domain.Employee, map[string]interface{}, error)
	Create(employeeDto *domain.Employee) (*domain.Employee, error)
	Update(employeeDto map[string]interface{}) (*domain.Employee, error)
}
