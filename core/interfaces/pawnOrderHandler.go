package interfaces

import (
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type pawnOrderRouter struct {
	pawnOrderInteractor application.PawnOrderInteractor
}

func PawnOrderHandler(v1 *gin.RouterGroup) {
	var pawnOrderRouter pawnOrderRouter

	v1.POST("/", pawnOrderRouter.createPawnOrder)
	v1.GET("/", pawnOrderRouter.getAllPawnOrders)
	v1.GET("/:id", pawnOrderRouter.getPawnOrderById)
	v1.PATCH("/", pawnOrderRouter.updatePawnOrder)
}

func (router *pawnOrderRouter) createPawnOrder(c *gin.Context) {
	var pawnOrderDto domain.PawnOrder
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&pawnOrderDto); errBinding != nil || pawnOrderDto.CustomerID == 0 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	pawnOrder, err := router.pawnOrderInteractor.Create(&pawnOrderDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &pawnOrder})
	}
}

func (router *pawnOrderRouter) getAllPawnOrders(c *gin.Context) {
	params := c.Request.URL.Query()
	pawnOrders, paginationResult, err := router.pawnOrderInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &pawnOrders, Previous: "localhost:9000/probien/api/v1/pawn-orders/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/probien/api/v1/pawn-orders/?page=" + paginationResult["next"].(string)})
	}
}

func (router *pawnOrderRouter) getPawnOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pawnOrder, err := router.pawnOrderInteractor.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &pawnOrder})
	}
}

func (router *pawnOrderRouter) updatePawnOrder(c *gin.Context) {
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

	pawnOrder, err := router.pawnOrderInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &pawnOrder})
	}
}
