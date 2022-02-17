package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type PawnOrderRepository interface {
	GetById(c *gin.Context) (*domain.PawnOrder, error)
	GetAll() (*[]domain.PawnOrder, error)
	Create(c *gin.Context) (*domain.PawnOrder, error)
	Update(c *gin.Context) (*domain.PawnOrder, error)
}
