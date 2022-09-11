package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type ICustomerRepository interface {
	GetById(c *gin.Context) (*domain.Customer, error)
	GetAll(c *gin.Context) (*[]domain.Customer, error)
	Create(c *gin.Context) (*domain.Customer, error)
	Update(c *gin.Context) (*domain.Customer, error)
}
