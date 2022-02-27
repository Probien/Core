package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func CategoryHandler(v1 *gin.RouterGroup) {

	categoryHandlerV1 := *v1.Group("/categories")
	interactor := application.CategoryInteractor{}

	categoryHandlerV1.POST("/", func(c *gin.Context) {
		category, err := interactor.Create(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully created", Data: &category})
		}
	})

	categoryHandlerV1.GET("/", func(c *gin.Context) {
		categories, err := interactor.GetAll()

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.Response{Status: http.StatusInternalServerError, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &categories})
		}
	})

	categoryHandlerV1.GET("/:id", func(c *gin.Context) {
		category, err := interactor.GetById(c)

		if err != nil {
			c.JSON(
				http.StatusNotFound,
				common.Response{Status: http.StatusNotFound, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &category})
		}
	})

	categoryHandlerV1.PATCH("/", func(c *gin.Context) {
		category, err := interactor.Update(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusAccepted, Message: "successfully updated", Data: &category})
		}
	})

	categoryHandlerV1.DELETE("/", func(c *gin.Context) {
		category, err := interactor.Delete(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusAccepted, Message: "successfully deleted", Data: &category})
		}
	})
}
