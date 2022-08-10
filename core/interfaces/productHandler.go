package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type productRouter struct {
	productInteractor application.ProductInteractor
}

func ProductHandler(v1 *gin.RouterGroup) {

	var productRouter productRouter

	v1.POST("/", productRouter.createProduct)
	v1.GET("/", productRouter.getAllProducts)
	v1.GET("/:id", productRouter.getProductById)
	v1.PATCH("/", productRouter.updateProduct)
}

func (router *productRouter) createProduct(c *gin.Context) {
	product, err := router.productInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &product})
	}
}

func (router *productRouter) getAllProducts(c *gin.Context) {
	products, err := router.productInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &products})
	}
}

func (router *productRouter) getProductById(c *gin.Context) {
	product, err := router.productInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &product})
	}
}

func (router *productRouter) updateProduct(c *gin.Context) {
	product, err := router.productInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &product})
	}
}
