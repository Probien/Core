package interfaces

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/persistence/postgres"
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/pkg/application"
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/interfaces/response"
	"github.com/gin-gonic/gin"
)

type endorsementRouter struct {
	endorsementInteractor application.EndorsementInteractor
}

func NewEndorsementHandler() *endorsementRouter {
	//dependency injection
	return &endorsementRouter{
		endorsementInteractor: application.NewEndorsementInteractor(postgres.NewEndorsementRepositoryImpl(config.GetConnection())),
	}
}

func (e *endorsementRouter) SetupRoutesAndFilter(v1 *gin.RouterGroup) {
	er := v1.Group("/").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE"))
	er.POST("endorsements", e.createEndorsement)
	er.GET("endorsements", e.getAllEndorsements)
	er.GET("endorsements/:id", e.getEndorsementById)
}

func (e *endorsementRouter) createEndorsement(c *gin.Context) {
	var endorsementDto domain.Endorsement
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&endorsementDto); errBinding != nil || endorsementDto.PawnOrderID == 0 || endorsementDto.EmployeeID == 0 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	endorsement, err := e.endorsementInteractor.Create(&endorsementDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &endorsement})
	}
}

func (e *endorsementRouter) getAllEndorsements(c *gin.Context) {
	params := c.Request.URL.Query()
	endorsements, paginationResult, err := e.endorsementInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &endorsements, Previous: "localhost:9000/api/v1/endorsements/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/endorsements/?page=" + paginationResult["next"].(string)})
	}
}

func (e *endorsementRouter) getEndorsementById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	endorsement, err := e.endorsementInteractor.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &endorsement})
	}
}
