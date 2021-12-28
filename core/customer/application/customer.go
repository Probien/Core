package application

import (
	"github.com/JairDavid/Probien-Backend/core/customer/domain"
	"github.com/gin-gonic/gin"
)

type CustomerInteractor struct {
	repository domain.CustomerRepository
}

func (CI *CustomerInteractor) GetById(c *gin.Context) (domain.Customer, error) {
	return CI.repository.GetById(c)
}

func (CI *CustomerInteractor) GetAll() ([]domain.Customer, error) {
	return CI.repository.GetAll()
}

func (CI *CustomerInteractor) Create(c *gin.Context) (domain.Customer, error) {
	return CI.repository.Create(c)
}

func (CI *CustomerInteractor) Update(c *gin.Context) (domain.Customer, error) {
	return CI.repository.Update(c)
}
