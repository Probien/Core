package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
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
	employee, err := router.employeeInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &employee})
	}
}

func (router *employeeRouter) getAllEmployees(c *gin.Context) {
	employees, err := router.employeeInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employees})
	}
}

func (router *employeeRouter) getEmployeeByEmail(c *gin.Context) {
	employee, err := router.employeeInteractor.GetByEmail(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employee})
	}
}

func (router *employeeRouter) updateEmployee(c *gin.Context) {
	employee, err := router.employeeInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.UPDATED, Data: &employee})
	}
}
