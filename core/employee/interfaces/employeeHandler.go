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
		interactor.Login(c)
	})

	employeeHandlerV1.POST("/", func(c *gin.Context) {
		employee, err := interactor.Create(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": employee})
	})

	employeeHandlerV1.GET("/", func(c *gin.Context) {

	})
}
