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
	Create(employeeDto *domain.Employee, userSessionId int) (*domain.Employee, error)
	Update(id int, employeeDto map[string]interface{}, userSessionId int) (*domain.Employee, error)
}
