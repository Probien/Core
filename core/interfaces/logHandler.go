package interfaces

import (
	"net/http"
	"strconv"

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
	params := c.Request.URL.Query()
	sessions, paginationResult, err := router.logInteractor.GetAllSessions(params)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &sessions, Previous: "localhost:9000/api/v1/logs/sessions/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/sessions/?page=" + paginationResult["next"].(string)})
	}
}

func (router *logRouter) getAllSessionsById(c *gin.Context) {
	params := c.Request.URL.Query()
	id, _ := strconv.Atoi(c.Param("id"))
	sessions, paginationResult, err := router.logInteractor.GetAllSessionsByEmployeeId(id, params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &sessions, Previous: "localhost:9000/api/v1/logs/sessions/" + c.Param("id") + "/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/sessions/" + c.Param("id") + "/?page=" + paginationResult["next"].(string)})
	}
}

func (router *logRouter) getAllMovements(c *gin.Context) {
	params := c.Request.URL.Query()
	movements, paginationResult, err := router.logInteractor.GetAllMovements(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &movements, Previous: "localhost:9000/api/v1/logs/movements/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/movements/?page=" + paginationResult["next"].(string)})
	}
}

func (router *logRouter) getAllMovementsById(c *gin.Context) {
	params := c.Request.URL.Query()
	id, _ := strconv.Atoi(c.Param("id"))
	movements, paginationResult, err := router.logInteractor.GetAllMovementsByEmployeeId(id, params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &movements, Previous: "localhost:9000/api/v1/logs/movements/" + c.Param("id") + "/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/movements/" + c.Param("id") + "/?page=" + paginationResult["next"].(string)})
	}
}
