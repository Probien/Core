package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type ICategoryRepository interface {
	GetById(id int) (*domain.Category, error)
	GetAll(params url.Values) (*[]domain.Category, map[string]interface{}, error)
	Create(categoryDto *domain.Category) (*domain.Category, error)
	Delete(id int) (*domain.Category, error)
	Update(categoryDto map[string]interface{}) (*domain.Category, error)
}
