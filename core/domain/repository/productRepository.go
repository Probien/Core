package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type IProductRepository interface {
	GetById(id int) (*domain.Product, error)
	GetAll(params url.Values) (*[]domain.Product, map[string]interface{}, error)
	Create(productDto *domain.Product) (*domain.Product, error)
	Update(productDto map[string]interface{}) (*domain.Product, error)
}
