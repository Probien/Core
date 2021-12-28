package domain

import "github.com/gin-gonic/gin"

type PawnOrderRepository interface {
	GetById(c *gin.Context) (PawnOrder, error)
	GetAll() ([]PawnOrder, error)
	Create(c *gin.Context) (PawnOrder, error)
	Update(c *gin.Context) (PawnOrder, error)
}
