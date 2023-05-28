package interfaces

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/persistence/postgres"
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/pkg/application"
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/interfaces/response"
	"github.com/gin-gonic/gin"
)

type pawnOrderRouter struct {
	pawnOrderInteractor application.PawnOrderInteractor
}

func NewPawnOrderHandler() *pawnOrderRouter {
	//dependency injection
	return &pawnOrderRouter{
		pawnOrderInteractor: application.NewPawnOrderInteractor(postgres.NewPawnOrderRepositoryImpl(config.GetConnection())),
	}
}

func (p *pawnOrderRouter) SetupRoutesAndFilter(v1 *gin.RouterGroup) {
	pr := v1.Group("/").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE"))
	pr.POST("pawn-orders", p.createPawnOrder)
	pr.GET("pawn-orders", p.getAllPawnOrders)
	pr.GET("pawn-orders/:id", p.getPawnOrderById)
	pr.PATCH("pawn-orders", p.updatePawnOrder)
}

func (p *pawnOrderRouter) createPawnOrder(c *gin.Context) {
	var pawnOrderDto domain.PawnOrder
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&pawnOrderDto); errBinding != nil || pawnOrderDto.CustomerID == 0 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	pawnOrder, err := p.pawnOrderInteractor.Create(&pawnOrderDto, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &pawnOrder})

}

func (p *pawnOrderRouter) getAllPawnOrders(c *gin.Context) {
	params := c.Request.URL.Query()
	pawnOrders, paginationResult, err := p.pawnOrderInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &pawnOrders, Previous: "localhost:9000/api/v1/pawn-orders/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/pawn-orders/?page=" + paginationResult["next"].(string)})
}

func (p *pawnOrderRouter) getPawnOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pawnOrder, err := p.pawnOrderInteractor.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &pawnOrder})
}

func (p *pawnOrderRouter) updatePawnOrder(c *gin.Context) {
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

	pawnOrder, err := p.pawnOrderInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Updated, Data: &pawnOrder})
}
