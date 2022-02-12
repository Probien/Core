package category_domain

import "github.com/gin-gonic/gin"

type CategoryRepository interface {
	GetById(c *gin.Context) (*Category, error)
	GetAll() (*[]Category, error)
	Create(c *gin.Context) (*Category, error)
	Delete(c *gin.Context) (*Category, error)
	Update(c *gin.Context) (*Category, error)
}
