package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
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
	pawnOrder, err := router.pawnOrderInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &pawnOrder})
	}
}

func (router *pawnOrderRouter) getAllPawnOrders(c *gin.Context) {
	pawnOrders, err := router.pawnOrderInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &pawnOrders})
	}
}

func (router *pawnOrderRouter) getPawnOrderById(c *gin.Context) {
	pawnOrder, err := router.pawnOrderInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &pawnOrder})
	}
}

func (router *pawnOrderRouter) updatePawnOrder(c *gin.Context) {
	pawnOrder, err := router.pawnOrderInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &pawnOrder})
	}
}
