package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type employeeRouter struct {
	employeeInteractor application.EmployeeInteractor
}

func EmployeeHandler(v1 *gin.RouterGroup) {

	var employeeRouter employeeRouter
	employeeHandlerV1 := *v1.Group("/employees")
	employeeHandlerV1.Use(auth.JwtAuth(true))

	employeeHandlerV1.POST("/", employeeRouter.createEmployee)
	employeeHandlerV1.GET("/", employeeRouter.getAllEmployees)
	employeeHandlerV1.GET("/byEmail/", employeeRouter.getEmployeeByEmail)
	employeeHandlerV1.PATCH("/", employeeRouter.updateEmployee)
}

func (ei *employeeRouter) createEmployee(c *gin.Context) {
	employee, err := ei.employeeInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &employee})
	}
}

func (ei *employeeRouter) getAllEmployees(c *gin.Context) {
	employees, err := ei.employeeInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employees})
	}
}

func (ei *employeeRouter) getEmployeeByEmail(c *gin.Context) {
	employee, err := ei.employeeInteractor.GetByEmail(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employee})
	}
}

func (ei *employeeRouter) updateEmployee(c *gin.Context) {
	employee, err := ei.employeeInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.UPDATED, Data: &employee})
	}
}
