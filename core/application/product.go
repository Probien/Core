package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

type ProductInteractor struct {
}

func (PI *ProductInteractor) GetById(c *gin.Context) (*domain.Product, error) {
	repository := persistence.NewProductRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (PI *ProductInteractor) GetAll(c *gin.Context) (*[]domain.Product, error) {
	repository := persistence.NewProductRepositoryImpl(config.Database)
	return repository.GetAll(c)
}

func (PI *ProductInteractor) Create(c *gin.Context) (*domain.Product, error) {
	repository := persistence.NewProductRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (PI *ProductInteractor) Update(c *gin.Context) (*domain.Product, error) {
	repository := persistence.NewProductRepositoryImpl(config.Database)
	return repository.Update(c)
}
