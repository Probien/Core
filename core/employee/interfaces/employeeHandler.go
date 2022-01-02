package interfaces

import "github.com/gin-gonic/gin"

func EmployeeHandler(v1 *gin.RouterGroup) {

	employeeHandlerV1 := *v1.Group("/employee")

	employeeHandlerV1.GET("/", func(c *gin.Context) {

	})
}
