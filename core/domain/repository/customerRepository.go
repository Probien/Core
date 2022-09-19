package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type ICustomerRepository interface {
	GetById(id int) (*domain.Customer, error)
	GetAll(params url.Values) (*[]domain.Customer, map[string]interface{}, error)
	Create(customerDto *domain.Customer) (*domain.Customer, error)
	Update(customerDto map[string]interface{}) (*domain.Customer, error)
}
