package interfaces

import (
	"net/http"

	"context"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

var bgctx = context.Background()

func AuthHandler(v1 *gin.RouterGroup) {

	security := *v1.Group("/auth")
	interactor := application.EmployeeInteractor{}

	security.POST("/login", func(c *gin.Context) {
		tokenizer := make(chan string, 1)
		employee, err := interactor.Login(c)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			interactor.GenerateToken(employee, tokenizer)
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employee, Token: <-tokenizer})
		}
	})

	security.POST("/logout", func(ctx *gin.Context) {
		config.Client.Get(bgctx, "")
	})

}
