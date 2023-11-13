package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type EmployeeApp struct {
	port port.IEmployeeRepository
}

func NewEmployeeApp(repository port.IEmployeeRepository) EmployeeApp {
	return EmployeeApp{
		port: repository,
	}
}

func (e *EmployeeApp) GetByEmail(email string) (*dto.Employee, error) {
	return e.port.GetByEmail(email)
}

func (e *EmployeeApp) GetAll(params url.Values) (*[]dto.Employee, map[string]interface{}, error) {
	return e.port.GetAll(params)
}

func (e *EmployeeApp) Create(employeeDto *dto.Employee, userSessionId int) (*dto.Employee, error) {
	return e.port.Create(employeeDto, userSessionId)
}

func (e *EmployeeApp) Update(id int, employeeDto map[string]interface{}, userSessionId int) (*dto.Employee, error) {
	return e.port.Update(id, employeeDto, userSessionId)
}
