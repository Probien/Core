package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type IProductRepository interface {
	GetById(id int) (*domain.Product, error)
	GetAll(params url.Values) (*[]domain.Product, map[string]interface{}, error)
	Create(productDto *domain.Product, userSessionId int) (*domain.Product, error)
	Update(id int, productDto map[string]interface{}, userSessionId int) (*domain.Product, error)
}
