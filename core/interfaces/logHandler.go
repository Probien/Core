package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func LogHandler(v1 *gin.RouterGroup) {
	logHandler := v1.Group("/logs")
	logHandler.Use(auth.RoutesAndAuthority(true))
	interactor := application.LogsInteractor{}

	logHandler.GET("/sessions", func(c *gin.Context) {
		sessions, err := interactor.GetAllSessions()
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully consulted", Data: &sessions})
		}

	})

	logHandler.GET("/sessions/:id", func(c *gin.Context) {
		sessions, err := interactor.GetAllSessionsByEmployeeId(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully consulted", Data: &sessions})
		}
	})

	logHandler.GET("/payments", func(c *gin.Context) {
		payments, err := interactor.GetAllPayments()

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully consulted", Data: &payments})
		}
	})

	logHandler.GET("/payments/:id", func(c *gin.Context) {
		payments, err := interactor.GetAllPaymentsByCustomerId(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully consulted", Data: &payments})
		}
	})

	logHandler.GET("/movements", func(c *gin.Context) {
		movements, err := interactor.GetAllMovements()

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully consulted", Data: &movements})
		}
	})

	logHandler.GET("/movements/:id", func(c *gin.Context) {
		movements, err := interactor.GetAllMovementsByEmployeeId(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully consulted", Data: &movements})
		}
	})
}
