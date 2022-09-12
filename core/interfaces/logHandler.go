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

	v1.GET("/sessions", logRouter.getAllSessions)
	v1.GET("/sessions/:id", logRouter.getAllSessionsById)
	v1.GET("/movements", logRouter.getAllMovements)
	v1.GET("/movements/:id", logRouter.getAllMovementsById)
}

func (router *logRouter) getAllSessions(c *gin.Context) {
	sessions, paginationResult, err := router.logInteractor.GetAllSessions(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: sessions, Previous: "localhost:9000/probien/api/v1/logs/sessions/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/probien/api/v1/logs/sessions/?page=" + paginationResult["next"].(string)})
	}

}

func (router *logRouter) getAllSessionsById(c *gin.Context) {
	sessions, err := router.logInteractor.GetAllSessionsByEmployeeId(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &sessions})
	}
}

func (router *logRouter) getAllMovements(c *gin.Context) {
	movements, err := router.logInteractor.GetAllMovements(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &movements})
	}
}

func (router *logRouter) getAllMovementsById(c *gin.Context) {
	movements, err := router.logInteractor.GetAllMovementsByEmployeeId(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &movements})
	}
}
