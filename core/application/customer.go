package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"

	"github.com/gin-gonic/gin"
)

type CustomerInteractor struct {
}

func (CI *CustomerInteractor) GetById(c *gin.Context) (*domain.Customer, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (CI *CustomerInteractor) GetAll(c *gin.Context) (*[]domain.Customer, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.GetAll(c)
}

func (CI *CustomerInteractor) Create(c *gin.Context) (*domain.Customer, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (CI *CustomerInteractor) Update(c *gin.Context) (*domain.Customer, error) {
	repository := postgres.NewCustomerRepositoryImpl(config.Database)
	return repository.Update(c)
}
