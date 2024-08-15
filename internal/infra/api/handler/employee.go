package handler

import (
	"github.com/JairDavid/Probien-Backend/internal/app"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IEmployeeHandler interface {
	GetByEmail(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type EmployeeHandler struct {
	app application.EmployeeApp
}

func NewEmployeeHandler(app application.EmployeeApp) IEmployeeHandler {
	return EmployeeHandler{
		app: app,
	}
}

func (e EmployeeHandler) GetByEmail(c *gin.Context) {
	var requestEmailBody map[string]string

	if errBinding := c.ShouldBindJSON(&requestEmailBody); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	email, existEmail := requestEmailBody["email"]
	if !existEmail {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: response.ErrorBinding, Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := e.app.GetByEmail(email)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &employee})
}

func (e EmployeeHandler) GetAll(c *gin.Context) {
	params := c.Request.URL.Query()
	employees, paginationResult, err := e.app.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &employees, Previous: "localhost:9000/api/v1/employees/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/employees/?page=" + paginationResult["next"].(string)})
}

func (e EmployeeHandler) Create(c *gin.Context) {
	var employeeDto *dto.Employee
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&employeeDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := e.app.Create(employeeDto, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &employee})
}

func (e EmployeeHandler) Update(c *gin.Context) {
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

	employee, err := e.app.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Updated, Data: &employee})
}
