package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type logRouter struct {
	logInteractor application.LogsInteractor
}

func LogHandler(v1 *gin.RouterGroup) {

	var logRouter logRouter
	logHandler := v1.Group("/logs")

	logHandler.GET("/sessions", logRouter.getAllSessions)
	logHandler.GET("/sessions/:id", logRouter.getAllSessionsById)
	logHandler.GET("/payments", logRouter.getAllPayments)
	logHandler.GET("/payments/:id", logRouter.getAllPaymentsById)
	logHandler.GET("/movements", logRouter.getAllMovements)
	logHandler.GET("/movements/:id", logRouter.getAllMovementsById)
}

func (router *logRouter) getAllSessions(c *gin.Context) {
	sessions, err := router.logInteractor.GetAllSessions()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &sessions})
	}

}

func (router *logRouter) getAllSessionsById(c *gin.Context) {
	sessions, err := router.logInteractor.GetAllSessionsByEmployeeId(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &sessions})
	}
}

func (router *logRouter) getAllPayments(c *gin.Context) {
	payments, err := router.logInteractor.GetAllPayments()

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &payments})
	}
}

func (router *logRouter) getAllPaymentsById(c *gin.Context) {
	payments, err := router.logInteractor.GetAllPaymentsByCustomerId(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &payments})
	}
}

func (router *logRouter) getAllMovements(c *gin.Context) {
	movements, err := router.logInteractor.GetAllMovements()

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &movements})
	}
}

func (router *logRouter) getAllMovementsById(c *gin.Context) {
	movements, err := router.logInteractor.GetAllMovementsByEmployeeId(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &movements})
	}
}
