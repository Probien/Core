package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type BranchOfficeRepository interface {
	GetAll() (*[]domain.BranchOffice, error)
	GetById(c *gin.Context) (*domain.BranchOffice, error)
	Create(c *gin.Context) (*domain.BranchOffice, error)
	Update(c *gin.Context) (*domain.BranchOffice, error)
}
