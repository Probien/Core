package domain

import "github.com/gin-gonic/gin"

type EmployeeRepository interface {
	GetByEmail(c *gin.Context) (Employee, error)
	GetAll() ([]Employee, error)
	Create(c *gin.Context) (Employee, error)
	Update(c *gin.Context) (Employee, error)
}
