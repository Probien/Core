package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

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
			auth.GenerateToken(employee, tokenizer)
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employee, Token: <-tokenizer})
		}
	})

	security.POST("/logout", func(c *gin.Context) {

		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.LOGOUT_DONE, Data: common.OUT})
	})

}
