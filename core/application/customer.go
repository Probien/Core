package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type CustomerInteractor struct {
}

func (CI *CustomerInteractor) GetById(id int) (*domain.Customer, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.GetById(id)
}

func (CI *CustomerInteractor) GetAll(params url.Values) (*[]domain.Customer, map[string]interface{}, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.GetAll(params)
}

func (CI *CustomerInteractor) Create(customerDto *domain.Customer, userSessionId int) (*domain.Customer, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.Create(customerDto, userSessionId)
}

func (CI *CustomerInteractor) Update(id int, customerDto map[string]interface{}, userSessionId int) (*domain.Customer, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.Update(id, customerDto, userSessionId)
}
