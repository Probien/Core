package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func CustomerHandler(v1 *gin.RouterGroup) {

	customerHandlerV1 := *v1.Group("/customer")
	interactor := application.CustomerInteractor{}

	customerHandlerV1.POST("/", func(c *gin.Context) {
		customer, err := interactor.Create(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully created", Data: &customer})
		}
	})

	customerHandlerV1.GET("/", func(c *gin.Context) {
		customers, err := interactor.GetAll()

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.Response{Status: http.StatusInternalServerError, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &customers})
		}
	})

	customerHandlerV1.GET("/:id", func(c *gin.Context) {
		customer, err := interactor.GetById(c)

		if err != nil {
			c.JSON(
				http.StatusNotFound,
				common.Response{Status: http.StatusNotFound, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &customer})
		}
	})

	customerHandlerV1.PATCH("/", func(c *gin.Context) {
		customer, err := interactor.Update(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusAccepted, Message: "successfully updated", Data: &customer})
		}
	})

}
