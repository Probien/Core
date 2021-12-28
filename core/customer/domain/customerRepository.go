package domain

import "github.com/gin-gonic/gin"

type CustomerRepository interface {
	GetById(c *gin.Context) (Customer, error)
	GetAll() ([]Customer, error)
	Create(c *gin.Context) (Customer, error)
	Update(c *gin.Context) (Customer, error)
}
