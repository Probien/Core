package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/employee/application"
	"github.com/gin-gonic/gin"
)

func EmployeeHandler(v1 *gin.RouterGroup) {

	interactor := application.EmployeeInteractor{}
	employeeHandlerV1 := *v1.Group("/employee")

	employeeHandlerV1.POST("/login", func(c *gin.Context) {
		employee, err := interactor.Login(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": employee})
		}
	})

	employeeHandlerV1.POST("/", func(c *gin.Context) {
		employee, err := interactor.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": employee})
	})

	employeeHandlerV1.GET("/", func(c *gin.Context) {

	})
}
