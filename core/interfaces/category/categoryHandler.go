package category_interface

import (
	"net/http"

	application "github.com/JairDavid/Probien-Backend/core/application/category"
	"github.com/gin-gonic/gin"
)

func CategoryHandler(v1 *gin.RouterGroup) {

	categoryHandlerV1 := *v1.Group("/category")
	interactor := application.CategoryInteractor{}

	categoryHandlerV1.POST("/", func(c *gin.Context) {
		category, err := interactor.Create(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": &category})
	})

	categoryHandlerV1.GET("/", func(c *gin.Context) {
		categories, err := interactor.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": "something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &categories})
		}
	})

	categoryHandlerV1.GET("/:id", func(c *gin.Context) {
		category, err := interactor.GetById(c)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &category})
		}
	})

	categoryHandlerV1.PATCH("/", func(c *gin.Context) {
		category, err := interactor.Update(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &category})
		}
	})

	categoryHandlerV1.DELETE("/", func(c *gin.Context) {
		category, err := interactor.Delete(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &category})
		}
	})
}
