package handler

import (
	"github.com/JairDavid/Probien-Backend/internal/app"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ICustomerHandler interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type CustomerHandler struct {
	app application.CustomerApp
}

func NewCustomerHandler(app application.CustomerApp) ICustomerHandler {
	return CustomerHandler{
		app: app,
	}
}

func (cu CustomerHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := cu.app.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &customer})
}

func (cu CustomerHandler) GetAll(c *gin.Context) {
	params := c.Request.URL.Query()
	customers, paginationResult, err := cu.app.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &customers, Previous: "localhost:9000/api/v1/customers/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/customers/?page=" + paginationResult["next"].(string)})

}

func (cu CustomerHandler) Create(c *gin.Context) {
	var customerDto dto.Customer
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&customerDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	customer, err := cu.app.Create(&customerDto, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &customer})
}

func (cu CustomerHandler) Update(c *gin.Context) {
	requestBodyWithId := map[string]interface{}{}
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.Bind(&requestBodyWithId); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	id, errID := requestBodyWithId["id"]
	if !errID {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: response.ErrorBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	customer, err := cu.app.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Updated, Data: &customer})
}
