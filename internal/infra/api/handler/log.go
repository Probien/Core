package handler

import (
	"github.com/JairDavid/Probien-Backend/internal/app"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ILogHandler interface {
	GetAllSessions(c *gin.Context)
	GetAllSessionsByEmployeeId(c *gin.Context)
	GetAllMovements(c *gin.Context)
	GetAllMovementsByEmployeeId(c *gin.Context)
}

type LogHandler struct {
	app application.LogApp
}

func NewLogHandler(app application.LogApp) ILogHandler {
	return LogHandler{
		app: app,
	}
}

func (l LogHandler) GetAllSessions(c *gin.Context) {
	params := c.Request.URL.Query()
	sessions, paginationResult, err := l.app.GetAllSessions(params)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &sessions, Previous: "localhost:9000/api/v1/logs/sessions/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/sessions/?page=" + paginationResult["next"].(string)})
}

func (l LogHandler) GetAllSessionsByEmployeeId(c *gin.Context) {
	params := c.Request.URL.Query()
	id, _ := strconv.Atoi(c.Param("id"))
	sessions, paginationResult, err := l.app.GetAllSessionsByEmployeeId(id, params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &sessions, Previous: "localhost:9000/api/v1/logs/sessions/" + c.Param("id") + "/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/sessions/" + c.Param("id") + "/?page=" + paginationResult["next"].(string)})
}

func (l LogHandler) GetAllMovements(c *gin.Context) {
	params := c.Request.URL.Query()
	movements, paginationResult, err := l.app.GetAllMovements(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &movements, Previous: "localhost:9000/api/v1/logs/movements/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/movements/?page=" + paginationResult["next"].(string)})
}

func (l LogHandler) GetAllMovementsByEmployeeId(c *gin.Context) {
	params := c.Request.URL.Query()
	id, _ := strconv.Atoi(c.Param("id"))
	movements, paginationResult, err := l.app.GetAllMovementsByEmployeeId(id, params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &movements, Previous: "localhost:9000/api/v1/logs/movements/" + c.Param("id") + "/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/movements/" + c.Param("id") + "/?page=" + paginationResult["next"].(string)})
}
