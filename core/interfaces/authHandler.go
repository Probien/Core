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
		tokenizer, session := make(chan string, 1), make(chan auth.SessionCredentials, 1)
		employee, err := interactor.Login(c)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			go auth.GenerateToken(employee, tokenizer)
			go auth.GenerateSessionID(employee, session)
			sessionCoockie := <-session
			c.SetCookie("SID", sessionCoockie.ID, 60*30, "/", "localhost", true, true)
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &employee, Token: <-tokenizer})
		}
	})

	security.POST("/logout", func(c *gin.Context) {
		auth.ClearSessionID(c)
		c.SetCookie("SID", "", -1, "/", "localhost", true, true)
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.LOGOUT_DONE, Data: common.OUT})
	})

}
