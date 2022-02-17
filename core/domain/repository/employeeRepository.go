package repository

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/gin-gonic/gin"
)

type EmployeeRepository interface {
	Login(c *gin.Context) (*domain.Employee, error)
	GetByEmail(c *gin.Context) (*domain.Employee, error)
	GetAll() (*[]domain.Employee, error)
	Create(c *gin.Context) (*domain.Employee, error)
	Update(c *gin.Context) (*domain.Employee, error)
}
