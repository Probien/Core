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

func (ei *endorsementRouter) createEndorsement(c *gin.Context) {
	endorsement, err := ei.endorsementInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &endorsement})
	}
}

func (ei *endorsementRouter) getAllEndorsements(c *gin.Context) {
	endorsements, err := ei.endorsementInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &endorsements})
	}
}

func (ei *endorsementRouter) getEndorsementById(c *gin.Context) {
	endorsement, err := ei.endorsementInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &endorsement})
	}
}
