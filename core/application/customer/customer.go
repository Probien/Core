package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	customer_domain "github.com/JairDavid/Probien-Backend/core/domain/customer"
	customer_infra "github.com/JairDavid/Probien-Backend/core/infrastructure/customer"
	"github.com/gin-gonic/gin"
)

type CustomerInteractor struct {
}

func (CI *CustomerInteractor) GetById(c *gin.Context) (*customer_domain.Customer, error) {
	repository := customer_infra.NewCustomerRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (CI *CustomerInteractor) GetAll() (*[]customer_domain.Customer, error) {
	repository := customer_infra.NewCustomerRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (CI *CustomerInteractor) Create(c *gin.Context) (*customer_domain.Customer, error) {
	repository := customer_infra.NewCustomerRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (CI *CustomerInteractor) Update(c *gin.Context) (*customer_domain.Customer, error) {
	repository := customer_infra.NewCustomerRepositoryImpl(config.Database)
	return repository.Update(c)
}
