package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func EmployeeHandler(v1 *gin.RouterGroup) {

	interactor := application.EmployeeInteractor{}
	employeeHandlerV1 := *v1.Group("/employees")

	employeeHandlerV1.POST("/login", func(c *gin.Context) {
		tokenizer := make(chan string, 1)
		employee, err := interactor.Login(c)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			interactor.GenerateToken(employee, tokenizer)
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully logged in", Data: &employee, Token: <-tokenizer})
		}
	})

	employeeHandlerV1.POST("/", func(c *gin.Context) {
		employee, err := interactor.Create(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusCreated, Message: "successfully created", Data: &employee})
		}
	})

	employeeHandlerV1.GET("/", func(c *gin.Context) {
		employees, err := interactor.GetAll()

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.Response{Status: http.StatusInternalServerError, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &employees})
		}
	})

	employeeHandlerV1.GET("/byEmail/", func(c *gin.Context) {
		employee, err := interactor.GetByEmail(c)

		if err != nil {
			c.JSON(
				http.StatusNotFound,
				common.Response{Status: http.StatusNotFound, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &employee})
		}
	})

	employeeHandlerV1.PATCH("/", func(c *gin.Context) {
		employee, err := interactor.Update(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusAccepted, Message: "successfully updated", Data: &employee})
		}
	})
}
