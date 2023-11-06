package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type IEmployeeRepository interface {
	GetByEmail(email string) (*dto.Employee, error)
	GetAll(params url.Values) (*[]dto.Employee, map[string]interface{}, error)
	Create(employeeDto *dto.Employee, userSessionId int) (*dto.Employee, error)
	Update(id int, employeeDto map[string]interface{}, userSessionId int) (*dto.Employee, error)
}
