package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/product/application"
	"github.com/gin-gonic/gin"
)

func ProductHandler(v1 *gin.RouterGroup) {
	productHandlerV1 := *v1.Group("/product")
	interactor := application.ProductInteractor{}

	productHandlerV1.POST("/", func(c *gin.Context) {
		product, err := interactor.Create(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": &product})
	})

	productHandlerV1.GET("/", func(c *gin.Context) {
		products, err := interactor.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": "something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &products})
		}
	})

	productHandlerV1.GET("/:id", func(c *gin.Context) {
		product, err := interactor.GetById(c)

		if err != nil {
			c.JSON(http.StatusFound, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &product})
		}
	})

	productHandlerV1.PATCH("/", func(c *gin.Context) {
		product, err := interactor.Update(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &product})
		}
	})

}
