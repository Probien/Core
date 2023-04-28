package interfaces

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/interfaces/response"
	"github.com/gin-gonic/gin"
)

type productRouter struct {
	productInteractor application.ProductInteractor
}

func NewProductHandler() *productRouter {
	//dependency injection
	return &productRouter{
		productInteractor: application.NewProductInteractor(postgres.NewProductRepositoryImpl(config.GetConnection())),
	}
}

func (p *productRouter) SetupRoutesAndFilter(v1 *gin.RouterGroup) {
	v1.Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE"))
	v1.POST("/products", p.createProduct)
	v1.GET("/products", p.getAllProducts)
	v1.GET("/products/:id", p.getProductById)
	v1.PATCH("/products", p.updateProduct)
}

func (p *productRouter) createProduct(c *gin.Context) {
	var productDto *domain.Product
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&productDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	product, err := p.productInteractor.Create(productDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &product})
	}
}

func (p *productRouter) getAllProducts(c *gin.Context) {
	params := c.Request.URL.Query()
	products, paginationResult, err := p.productInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &products, Previous: "localhost:9000/api/v1/products/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/products/?page=" + paginationResult["next"].(string)})
	}
}

func (p *productRouter) getProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := p.productInteractor.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &product})
	}
}

func (p *productRouter) updateProduct(c *gin.Context) {
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

	product, err := p.productInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Updated, Data: &product})
	}
}
