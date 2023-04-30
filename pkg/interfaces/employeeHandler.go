package interfaces

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/persistence/postgres"
	"net/http"

	"github.com/JairDavid/Probien-Backend/pkg/application"
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/interfaces/response"
	"github.com/gin-gonic/gin"
)

type employeeRouter struct {
	employeeInteractor application.EmployeeInteractor
}

func NewEmployeeHandler() *employeeRouter {
	//dependency injection
	return &employeeRouter{
		employeeInteractor: application.NewEmployeeInteractor(postgres.NewEmployeeRepositoryImpl(config.GetConnection())),
	}
}

func (e *employeeRouter) SetupRoutesAndFilter(v1 *gin.RouterGroup) {
	er := v1.Group("/").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE"))
	er.POST("employees", e.createEmployee)
	er.GET("employees", e.getAllEmployees)
	er.GET("employees/byEmail", e.getEmployeeByEmail)
	er.PATCH("employees", e.updateEmployee)
}

func (e *employeeRouter) createEmployee(c *gin.Context) {
	var employeeDto *domain.Employee
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&employeeDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := e.employeeInteractor.Create(employeeDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &employee})
	}
}

func (e *employeeRouter) getAllEmployees(c *gin.Context) {
	params := c.Request.URL.Query()
	employees, paginationResult, err := e.employeeInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &employees, Previous: "localhost:9000/api/v1/employees/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/employees/?page=" + paginationResult["next"].(string)})
	}
}

func (e *employeeRouter) getEmployeeByEmail(c *gin.Context) {
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

	employee, err := e.employeeInteractor.GetByEmail(email)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &employee})
	}
}

func (e *employeeRouter) updateEmployee(c *gin.Context) {
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

	employee, err := e.employeeInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Updated, Data: &employee})
	}
}
