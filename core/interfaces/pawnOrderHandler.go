package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/gin-gonic/gin"
)

func PawnOrderHandler(v1 *gin.RouterGroup) {

	pawnOrderHandlerV1 := *v1.Group("/pawn-order")
	interactor := application.PawnOrderInteractor{}

	pawnOrderHandlerV1.POST("/", func(c *gin.Context) {
		pawnOrder, err := interactor.Create(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": &pawnOrder})
	})

	pawnOrderHandlerV1.GET("/", func(c *gin.Context) {
		pawnOrders, err := interactor.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": "something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &pawnOrders})
		}
	})

	pawnOrderHandlerV1.GET("/:id", func(c *gin.Context) {
		pawnOrder, err := interactor.GetById(c)

		if err != nil {
			c.JSON(http.StatusFound, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &pawnOrder})
		}
	})

	pawnOrderHandlerV1.PATCH("/", func(c *gin.Context) {
		pawnOrder, err := interactor.Update(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &pawnOrder})
		}
	})
}
