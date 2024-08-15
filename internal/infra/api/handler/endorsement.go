package handler

import (
	"github.com/JairDavid/Probien-Backend/internal/app"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IEndorsementHandler interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

type EndorsementHandler struct {
	app application.EndorsementApp
}

func NewEndorsementHandler(app application.EndorsementApp) IEndorsementHandler {
	return EndorsementHandler{
		app: app,
	}
}

func (e EndorsementHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	endorsement, err := e.app.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &endorsement})
}

func (e EndorsementHandler) GetAll(c *gin.Context) {
	params := c.Request.URL.Query()
	endorsements, paginationResult, err := e.app.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &endorsements, Previous: "localhost:9000/api/v1/endorsements/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/endorsements/?page=" + paginationResult["next"].(string)})
}

func (e EndorsementHandler) Create(c *gin.Context) {
	var endorsementDto dto.Endorsement
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&endorsementDto); errBinding != nil || endorsementDto.PawnOrderID == 0 || endorsementDto.EmployeeID == 0 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	endorsement, err := e.app.Create(&endorsementDto, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &endorsement})
}
