package interfaces

import (
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/domain"
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
	var productDto *domain.Product
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&productDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	product, err := router.productInteractor.Create(productDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &product})
	}
}

func (router *productRouter) getAllProducts(c *gin.Context) {
	params := c.Request.URL.Query()
	products, paginationResult, err := router.productInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &products, Previous: "localhost:9000/api/v1/products/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/products/?page=" + paginationResult["next"].(string)})
	}
}

func (router *productRouter) getProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := router.productInteractor.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &product})
	}
}

func (router *productRouter) updateProduct(c *gin.Context) {
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

	product, err := router.productInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &product})
	}
}
