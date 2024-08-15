package handler

import (
	"github.com/JairDavid/Probien-Backend/internal/app"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ICategoryHandler interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type CategoryHandler struct {
	app application.CategoryApp
}

func NewCategoryHandler(app application.CategoryApp) ICategoryHandler {
	return CategoryHandler{
		app: app,
	}
}

func (ca CategoryHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := ca.app.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &category})
}

func (ca CategoryHandler) GetAll(c *gin.Context) {
	params := c.Request.URL.Query()
	categories, paginationResult, err := ca.app.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &categories, Previous: "localhost:9000/api/v1/categories/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/categories/?page=" + paginationResult["next"].(string)})
}

func (ca CategoryHandler) Create(c *gin.Context) {
	var categoryDto dto.Category
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&categoryDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	category, err := ca.app.Create(&categoryDto, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &category})
}

func (ca CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")
	category, err := ca.app.Delete(id, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Deleted, Data: &category})
}

func (ca CategoryHandler) Update(c *gin.Context) {
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

	category, err := ca.app.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Updated, Data: &category})
}
