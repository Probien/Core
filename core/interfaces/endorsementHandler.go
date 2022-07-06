package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func EndorsementHandler(v1 *gin.RouterGroup) {

	endorsementHandlerV1 := *v1.Group("/endorsements")
	endorsementHandlerV1.Use(auth.JwtAuth(false))
	interactor := application.EndorsemenInteractor{}

	// Maybe it's clearly if the func is declared as package function instead of here, it's more readible
	endorsementHandlerV1.POST("/", func(c *gin.Context) {
		endorsement, err := interactor.Create(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &endorsement})
		}
	})

	endorsementHandlerV1.GET("/", func(c *gin.Context) {
		endorsements, err := interactor.GetAll()

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &endorsements})
		}
	})

	endorsementHandlerV1.GET("/:id", func(c *gin.Context) {
		endorsement, err := interactor.GetById(c)

		if err != nil {
			c.JSON(
				http.StatusNotFound,
				common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &endorsement})
		}
	})

}
