package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/pkg/domain"
)

type ICustomerRepository interface {
	GetById(id int) (*domain.Customer, error)
	GetAll(params url.Values) (*[]domain.Customer, map[string]interface{}, error)
	Create(customerDto *domain.Customer, userSessionId int) (*domain.Customer, error)
	Update(id int, customerDto map[string]interface{}, userSessionId int) (*domain.Customer, error)
}
