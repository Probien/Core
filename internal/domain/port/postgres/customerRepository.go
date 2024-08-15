package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type ICustomerRepository interface {
	GetById(id int) (*dto.Customer, error)
	GetAll(params url.Values) (*[]dto.Customer, map[string]interface{}, error)
	Create(customerDto *dto.Customer, userSessionId int) (*dto.Customer, error)
	Update(id int, customerDto map[string]interface{}, userSessionId int) (*dto.Customer, error)
}
