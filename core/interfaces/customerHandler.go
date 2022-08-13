package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type customerRouter struct {
	customerInteractor application.CustomerInteractor
}

func CustomerHandler(v1 *gin.RouterGroup) {
	var customerRouter customerRouter

	v1.POST("/", customerRouter.createCustomer)
	v1.GET("/", customerRouter.GetAllCustomers)
	v1.GET("/:id", customerRouter.getCustomerById)
	v1.PATCH("/", customerRouter.updateCustomer)
}

func (router *customerRouter) createCustomer(c *gin.Context) {
	customer, err := router.customerInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &customer})
	}
}

func (router *customerRouter) GetAllCustomers(c *gin.Context) {
	customers, err := router.customerInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &customers})
	}
}

func (router *customerRouter) getCustomerById(c *gin.Context) {
	customer, err := router.customerInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &customer})
	}
}

func (router *customerRouter) updateCustomer(c *gin.Context) {
	customer, err := router.customerInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &customer})
	}
}
