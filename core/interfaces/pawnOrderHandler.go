package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func PawnOrderHandler(v1 *gin.RouterGroup) {

	pawnOrderHandlerV1 := *v1.Group("/pawn-orders")
	pawnOrderHandlerV1.Use(auth.JwtAuth(false))
	interactor := application.PawnOrderInteractor{}

	pawnOrderHandlerV1.POST("/", func(c *gin.Context) {
		pawnOrder, err := interactor.Create(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &pawnOrder})
		}
	})

	pawnOrderHandlerV1.GET("/", func(c *gin.Context) {
		pawnOrders, err := interactor.GetAll()

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &pawnOrders})
		}
	})

	pawnOrderHandlerV1.GET("/:id", func(c *gin.Context) {
		pawnOrder, err := interactor.GetById(c)

		if err != nil {
			c.JSON(
				http.StatusNotFound,
				common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &pawnOrder})
		}
	})

	pawnOrderHandlerV1.PATCH("/", func(c *gin.Context) {
		pawnOrder, err := interactor.Update(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.UPDATED, Data: &pawnOrder})
		}
	})
}
