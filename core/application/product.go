package application

import (
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type ProductInteractor struct {
	repository repository.IProductRepository
}

func NewProductInteractor(repository repository.IProductRepository) ProductInteractor {
	return ProductInteractor{
		repository: repository,
	}
}

func (p *ProductInteractor) GetById(id int) (*domain.Product, error) {
	return p.repository.GetById(id)
}

func (p *ProductInteractor) GetAll(params url.Values) (*[]domain.Product, map[string]interface{}, error) {
	return p.repository.GetAll(params)
}

func (p *ProductInteractor) Create(productDto *domain.Product, userSessionId int) (*domain.Product, error) {
	return p.repository.Create(productDto, userSessionId)
}

func (p *ProductInteractor) Update(id int, productDto map[string]interface{}, userSessionId int) (*domain.Product, error) {
	return p.repository.Update(id, productDto, userSessionId)
}
