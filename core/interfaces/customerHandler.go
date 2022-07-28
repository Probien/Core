package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type customerRouter struct {
	customerInteractor application.CustomerInteractor
}

func CustomerHandler(v1 *gin.RouterGroup) {

	var customerRouter customerRouter
	customerHandlerV1 := *v1.Group("/customers")
	customerHandlerV1.Use(auth.JwtAuth(false))

	customerHandlerV1.POST("/", customerRouter.createCustomer)
	customerHandlerV1.GET("/", customerRouter.GetAllCustomers)
	customerHandlerV1.GET("/:id", customerRouter.getCustomerById)
	customerHandlerV1.PATCH("/", customerRouter.updateCustomer)
}

func (router *customerRouter) createCustomer(c *gin.Context) {
	customer, err := router.customerInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &customer})
	}
}

func (router *customerRouter) GetAllCustomers(c *gin.Context) {
	customers, err := router.customerInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &customers})
	}
}

func (router *customerRouter) getCustomerById(c *gin.Context) {
	customer, err := router.customerInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &customer})
	}
}

func (router *customerRouter) updateCustomer(c *gin.Context) {
	customer, err := router.customerInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.UPDATED, Data: &customer})
	}
}
