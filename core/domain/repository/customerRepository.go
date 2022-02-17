package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type CustomerRepository interface {
	GetById(c *gin.Context) (*domain.Customer, error)
	GetAll() (*[]domain.Customer, error)
	Create(c *gin.Context) (*domain.Customer, error)
	Update(c *gin.Context) (*domain.Customer, error)
}
