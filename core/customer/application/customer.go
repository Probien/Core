package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/customer/domain"
	"github.com/JairDavid/Probien-Backend/core/customer/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

type CustomerInteractor struct {
}

func (CI *CustomerInteractor) GetById(c *gin.Context) (domain.Customer, error) {
	repository := persistance.NewCustomerRepositoryImpl(config.GetDBInstance())
	return repository.GetById(c)
}

func (CI *CustomerInteractor) GetAll() ([]domain.Customer, error) {
	repository := persistance.NewCustomerRepositoryImpl(config.GetDBInstance())
	return repository.GetAll()
}

func (CI *CustomerInteractor) Create(c *gin.Context) (domain.Customer, error) {
	repository := persistance.NewCustomerRepositoryImpl(config.GetDBInstance())
	return repository.Create(c)
}

func (CI *CustomerInteractor) Update(c *gin.Context) (domain.Customer, error) {
	repository := persistance.NewCustomerRepositoryImpl(config.GetDBInstance())
	return repository.Update(c)
}
