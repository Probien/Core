package application

import (
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/domain/repository"
	"net/url"
)

type CategoryInteractor struct {
	repository repository.ICategoryRepository
}

func NewCategoryInteractor(repository repository.ICategoryRepository) CategoryInteractor {
	return CategoryInteractor{
		repository: repository,
	}
}

func (c *CategoryInteractor) GetById(id int) (*domain.Category, error) {
	return c.repository.GetById(id)
}

func (c *CategoryInteractor) GetAll(params url.Values) (*[]domain.Category, map[string]interface{}, error) {
	return c.repository.GetAll(params)
}

func (c *CategoryInteractor) Create(categoryDto *domain.Category, userSessionId int) (*domain.Category, error) {
	return c.repository.Create(categoryDto, userSessionId)
}

func (c *CategoryInteractor) Delete(id int, userSessionId int) (*domain.Category, error) {
	return c.repository.Delete(id, userSessionId)
}

func (c *CategoryInteractor) Update(id int, categoryDto map[string]interface{}, userSessionId int) (*domain.Category, error) {
	return c.repository.Update(id, categoryDto, userSessionId)
}
