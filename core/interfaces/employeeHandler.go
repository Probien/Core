package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type employeeRouter struct {
	employeeInteractor application.EmployeeInteractor
}

func EmployeeHandler(v1 *gin.RouterGroup) {
	var employeeRouter employeeRouter

	v1.POST("/", employeeRouter.createEmployee)
	v1.GET("/", employeeRouter.getAllEmployees)
	v1.GET("/byEmail/", employeeRouter.getEmployeeByEmail)
	v1.PATCH("/", employeeRouter.updateEmployee)
}

func (router *employeeRouter) createEmployee(c *gin.Context) {
	var employeeDto *domain.Employee
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&employeeDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := router.employeeInteractor.Create(employeeDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &employee})
	}
}

func (router *employeeRouter) getAllEmployees(c *gin.Context) {
	params := c.Request.URL.Query()
	employees, paginationResult, err := router.employeeInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &employees, Previous: "localhost:9000/api/v1/employees/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/employees/?page=" + paginationResult["next"].(string)})
	}
}

func (router *employeeRouter) getEmployeeByEmail(c *gin.Context) {
	var requestEmailBody map[string]string

	if errBinding := c.ShouldBindJSON(&requestEmailBody); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	email, existEmail := requestEmailBody["email"]

	if !existEmail {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: common.ErrorBinding, Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := router.employeeInteractor.GetByEmail(email)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &employee})
	}
}

func (router *employeeRouter) updateEmployee(c *gin.Context) {
	requestBodyWithId := map[string]interface{}{}
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.Bind(&requestBodyWithId); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	id, errID := requestBodyWithId["id"]

	if !errID {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: common.ErrorBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := router.employeeInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &employee})
	}
}
