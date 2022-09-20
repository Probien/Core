package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type ProductInteractor struct {
}

func (PI *ProductInteractor) GetById(id int) (*domain.Product, error) {
	repository := postgres.NewProductRepositoryImpl(config.Database)
	return repository.GetById(id)
}

func (PI *ProductInteractor) GetAll(params url.Values) (*[]domain.Product, map[string]interface{}, error) {
	repository := postgres.NewProductRepositoryImpl(config.Database)
	return repository.GetAll(params)
}

func (PI *ProductInteractor) Create(productDto *domain.Product, userSessionId int) (*domain.Product, error) {
	repository := postgres.NewProductRepositoryImpl(config.Database)
	return repository.Create(productDto, userSessionId)
}

func (PI *ProductInteractor) Update(id int, productDto map[string]interface{}, userSessionId int) (*domain.Product, error) {
	repository := postgres.NewProductRepositoryImpl(config.Database)
	return repository.Update(id, productDto, userSessionId)
}
