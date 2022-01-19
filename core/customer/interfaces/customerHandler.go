package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/customer/application"
	"github.com/gin-gonic/gin"
)

func CustomerHandler(v1 *gin.RouterGroup) {

	customerHandlerV1 := *v1.Group("/customer")
	interactor := application.CustomerInteractor{}

	customerHandlerV1.POST("/", func(c *gin.Context) {
		customer, err := interactor.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": customer})
	})

	customerHandlerV1.GET("/", func(c *gin.Context) {

	})

	customerHandlerV1.GET("/:id", func(c *gin.Context) {

	})

	customerHandlerV1.PATCH("/", func(c *gin.Context) {

	})

}
