package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type endorsementRouter struct {
	endorsementInteractor application.EndorsemenInteractor
}

func EndorsementHandler(v1 *gin.RouterGroup) {
	var endorsementRouter endorsementRouter

	v1.POST("/", endorsementRouter.createEndorsement)
	v1.GET("/", endorsementRouter.getAllEndorsements)
	v1.GET("/:id", endorsementRouter.getEndorsementById)
}

func (router *endorsementRouter) createEndorsement(c *gin.Context) {
	endorsement, err := router.endorsementInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &endorsement})
	}
}

func (router *endorsementRouter) getAllEndorsements(c *gin.Context) {
	endorsements, paginationResult, err := router.endorsementInteractor.GetAll(c)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &endorsements, Previous: "localhost:9000/probien/api/v1/endorsements/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/probien/api/v1/endorsements/?page=" + paginationResult["next"].(string)})
	}
}

func (router *endorsementRouter) getEndorsementById(c *gin.Context) {
	endorsement, err := router.endorsementInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &endorsement})
	}
}
