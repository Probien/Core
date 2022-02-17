package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/gin-gonic/gin"
)

func EmployeeHandler(v1 *gin.RouterGroup) {

	interactor := application.EmployeeInteractor{}
	employeeHandlerV1 := *v1.Group("/employee")

	employeeHandlerV1.POST("/login", func(c *gin.Context) {
		tokenizer := make(chan string, 1)
		employee, err := interactor.Login(c)

		interactor.GenerateToken(employee, tokenizer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &employee, "token": <-tokenizer})
		}
	})

	employeeHandlerV1.POST("/", func(c *gin.Context) {
		employee, err := interactor.Create(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": &employee})
	})

	employeeHandlerV1.GET("/byEmail/", func(c *gin.Context) {
		employee, err := interactor.GetByEmail(c)

		if err != nil {
			c.JSON(http.StatusFound, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &employee})
		}
	})

	employeeHandlerV1.GET("/", func(c *gin.Context) {
		employees, err := interactor.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": "something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &employees})
		}
	})

	employeeHandlerV1.PATCH("/", func(c *gin.Context) {
		employee, err := interactor.Update(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &employee})
		}
	})
}
