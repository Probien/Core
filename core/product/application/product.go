package application

import (
	"github.com/JairDavid/Probien-Backend/core/product/domain"
	"github.com/gin-gonic/gin"
)

type ProductInteractor struct {
	repository domain.ProductRepository
}

func (PI *ProductInteractor) GetById(c *gin.Context) (domain.Product, error) {
	return PI.repository.GetById(c)
}

func (PI *ProductInteractor) GetAll() ([]domain.Product, error) {
	return PI.repository.GetAll()
}

func (PI *ProductInteractor) Create(c *gin.Context) ([]domain.Product, error) {
	return PI.repository.Create(c)
}

func (PI *ProductInteractor) Update(c *gin.Context) (domain.Product, error) {
	return PI.repository.Update(c)
}
