package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type CategoryInteractor struct {
}

func (CI *CategoryInteractor) GetById(id int) (*domain.Category, error) {
	repository := postgres.NewCategoryRepositoryImpl(config.Database)
	return repository.GetById(id)
}

func (CI *CategoryInteractor) GetAll(params url.Values) (*[]domain.Category, map[string]interface{}, error) {
	repository := postgres.NewCategoryRepositoryImpl(config.Database)
	return repository.GetAll(params)
}

func (CI *CategoryInteractor) Create(categoryDto *domain.Category, userSessionId int) (*domain.Category, error) {
	repository := postgres.NewCategoryRepositoryImpl(config.Database)
	return repository.Create(categoryDto, userSessionId)
}

func (CI *CategoryInteractor) Delete(id int, userSessionId int) (*domain.Category, error) {
	repository := postgres.NewCategoryRepositoryImpl(config.Database)
	return repository.Delete(id, userSessionId)
}

func (CI *CategoryInteractor) Update(id int, categoryDto map[string]interface{}, userSessionId int) (*domain.Category, error) {
	repository := postgres.NewCategoryRepositoryImpl(config.Database)
	return repository.Update(id, categoryDto, userSessionId)
}
