package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type CustomerApp struct {
	port port.ICustomerRepository
}

func NewCustomerApp(repository port.ICustomerRepository) CustomerApp {
	return CustomerApp{
		port: repository,
	}
}

func (c *CustomerApp) GetById(id int) (*dto.Customer, error) {
	return c.port.GetById(id)
}

func (c *CustomerApp) GetAll(params url.Values) (*[]dto.Customer, map[string]interface{}, error) {
	return c.port.GetAll(params)
}

func (c *CustomerApp) Create(customerDto *dto.Customer, userSessionId int) (*dto.Customer, error) {
	return c.port.Create(customerDto, userSessionId)
}

func (c *CustomerApp) Update(id int, customerDto map[string]interface{}, userSessionId int) (*dto.Customer, error) {
	return c.port.Update(id, customerDto, userSessionId)
}
