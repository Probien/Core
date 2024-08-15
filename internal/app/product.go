package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type ProductApp struct {
	port port.IProductRepository
}

func NewProductApp(repository port.IProductRepository) ProductApp {
	return ProductApp{
		port: repository,
	}
}

func (p *ProductApp) GetById(id int) (*dto.Product, error) {
	return p.port.GetById(id)
}

func (p *ProductApp) GetAll(params url.Values) (*[]dto.Product, map[string]interface{}, error) {
	return p.port.GetAll(params)
}

func (p *ProductApp) Create(productDto *dto.Product, userSessionId int) (*dto.Product, error) {
	return p.port.Create(productDto, userSessionId)
}

func (p *ProductApp) Update(id int, productDto map[string]interface{}, userSessionId int) (*dto.Product, error) {
	return p.port.Update(id, productDto, userSessionId)
}
