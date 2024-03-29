package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type authRouter struct {
	loginInteractor application.EmployeeInteractor
}

func AuthHandler(v1 *gin.RouterGroup) {
	var authRouter authRouter

	v1.POST("/login", authRouter.login)
	v1.POST("/logout", authRouter.logout)
}

func (router *authRouter) login(c *gin.Context) {
	tokenizer, session := make(chan string, 1), make(chan auth.SessionCredential, 1)
	var loginCredentials auth.LoginCredential

	if errBinding := c.ShouldBindJSON(&loginCredentials); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := router.loginInteractor.Login(loginCredentials)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		go auth.GenerateToken(employee, tokenizer)
		go auth.GenerateSessionID(employee, session)
		sessionCoockie := <-session
		c.SetCookie("SID", sessionCoockie.ID, 60*30, "/", "localhost", true, true)
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.LoginDone, Data: &employee, Token: <-tokenizer})
	}
}

func (router *authRouter) logout(c *gin.Context) {
	err := auth.ClearSessionID(c)
	c.SetCookie("SID", "", -1, "/", "localhost", true, true)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.LogoutDone, Data: "Done"})
}
