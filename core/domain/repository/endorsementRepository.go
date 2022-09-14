package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type IEndorsementRepository interface {
	GetById(c *gin.Context) (*domain.Endorsement, error)
	GetAll(c *gin.Context) (*[]domain.Endorsement, map[string]interface{}, error)
	Create(c *gin.Context) (*domain.Endorsement, error)
}
