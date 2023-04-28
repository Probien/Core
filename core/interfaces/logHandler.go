package interfaces

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/response"
	"github.com/gin-gonic/gin"
)

type logRouter struct {
	logInteractor application.LogsInteractor
}

func LogHandler() *logRouter {
	//dependency injection
	return &logRouter{
		logInteractor: application.NewLogsInteractor(postgres.NewLogsRepositoryImp(config.GetConnection())),
	}
}

func (l *logRouter) SetupRoutesAndFilter(v1 *gin.RouterGroup) {
	v1.Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER"))
	v1.GET("/logs/sessions", l.getAllSessions)
	v1.GET("/logs/sessions/:id", l.getAllSessionsById)
	v1.GET("/logs/movements", l.getAllMovements)
	v1.GET("/logs/movements/:id", l.getAllMovementsById)
}

func (l *logRouter) getAllSessions(c *gin.Context) {
	params := c.Request.URL.Query()
	sessions, paginationResult, err := l.logInteractor.GetAllSessions(params)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &sessions, Previous: "localhost:9000/api/v1/logs/sessions/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/sessions/?page=" + paginationResult["next"].(string)})
	}
}

func (l *logRouter) getAllSessionsById(c *gin.Context) {
	params := c.Request.URL.Query()
	id, _ := strconv.Atoi(c.Param("id"))
	sessions, paginationResult, err := l.logInteractor.GetAllSessionsByEmployeeId(id, params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &sessions, Previous: "localhost:9000/api/v1/logs/sessions/" + c.Param("id") + "/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/sessions/" + c.Param("id") + "/?page=" + paginationResult["next"].(string)})
	}
}

func (l *logRouter) getAllMovements(c *gin.Context) {
	params := c.Request.URL.Query()
	movements, paginationResult, err := l.logInteractor.GetAllMovements(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &movements, Previous: "localhost:9000/api/v1/logs/movements/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/movements/?page=" + paginationResult["next"].(string)})
	}
}

func (l *logRouter) getAllMovementsById(c *gin.Context) {
	params := c.Request.URL.Query()
	id, _ := strconv.Atoi(c.Param("id"))
	movements, paginationResult, err := l.logInteractor.GetAllMovementsByEmployeeId(id, params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &movements, Previous: "localhost:9000/api/v1/logs/movements/" + c.Param("id") + "/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/logs/movements/" + c.Param("id") + "/?page=" + paginationResult["next"].(string)})
	}
}
