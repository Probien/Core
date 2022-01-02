package domain

import "github.com/gin-gonic/gin"

type ProductRepository interface {
	GetById(c *gin.Context) (Product, error)
	GetAll() ([]Product, error)
	Create(c *gin.Context) ([]Product, error)
	Update(c *gin.Context) (Product, error)
}
