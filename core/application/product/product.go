package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	product_domain "github.com/JairDavid/Probien-Backend/core/domain/product"
	product_infra "github.com/JairDavid/Probien-Backend/core/infrastructure/product"
	"github.com/gin-gonic/gin"
)

type ProductInteractor struct {
}

func (PI *ProductInteractor) GetById(c *gin.Context) (*product_domain.Product, error) {
	repository := product_infra.NewProductRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (PI *ProductInteractor) GetAll() (*[]product_domain.Product, error) {
	repository := product_infra.NewProductRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (PI *ProductInteractor) Create(c *gin.Context) (*product_domain.Product, error) {
	repository := product_infra.NewProductRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (PI *ProductInteractor) Update(c *gin.Context) (*product_domain.Product, error) {
	repository := product_infra.NewProductRepositoryImpl(config.Database)
	return repository.Update(c)
}
