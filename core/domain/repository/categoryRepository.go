package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type ICategoryRepository interface {
	GetById(id int) (*domain.Category, error)
	GetAll(params url.Values) (*[]domain.Category, map[string]interface{}, error)
	Create(categoryDto *domain.Category, userSessionId int) (*domain.Category, error)
	Delete(id int, userSessionId int) (*domain.Category, error)
	Update(id int, categoryDto map[string]interface{}, userSessionId int) (*domain.Category, error)
}
