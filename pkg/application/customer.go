package application

import (
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/domain/repository"
	"net/url"
)

type CustomerInteractor struct {
	repository repository.ICustomerRepository
}

func NewCustomerInteractor(repository repository.ICustomerRepository) CustomerInteractor {
	return CustomerInteractor{
		repository: repository,
	}
}

func (c *CustomerInteractor) GetById(id int) (*domain.Customer, error) {
	return c.repository.GetById(id)
}

func (c *CustomerInteractor) GetAll(params url.Values) (*[]domain.Customer, map[string]interface{}, error) {
	return c.repository.GetAll(params)
}

func (c *CustomerInteractor) Create(customerDto *domain.Customer, userSessionId int) (*domain.Customer, error) {
	return c.repository.Create(customerDto, userSessionId)
}

func (c *CustomerInteractor) Update(id int, customerDto map[string]interface{}, userSessionId int) (*domain.Customer, error) {
	return c.repository.Update(id, customerDto, userSessionId)
}
