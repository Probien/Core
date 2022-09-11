package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type IPawnOrderRepository interface {
	GetById(c *gin.Context) (*domain.PawnOrder, error)
	GetByIdForUpdate(id uint) (*domain.PawnOrder, error)
	GetAll(c *gin.Context) (*[]domain.PawnOrder, error)
	Create(c *gin.Context) (*domain.PawnOrder, error)
	Update(c *gin.Context) (*domain.PawnOrder, error)
}
