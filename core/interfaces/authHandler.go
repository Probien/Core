package interfaces

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/response"
	"github.com/gin-gonic/gin"
)

type authRouter struct {
	loginInteractor application.EmployeeInteractor
}

func NewAuthHandler() *authRouter {
	//dependency injection
	return &authRouter{
		loginInteractor: application.NewEmployeeInteractor(postgres.NewEmployeeRepositoryImpl(config.GetConnection())),
	}
}

func (a *authRouter) SetupRoutes(v1 *gin.RouterGroup) {
	v1.POST("/login", a.login)
	v1.POST("/logout", a.logout)
}

func (a *authRouter) login(c *gin.Context) {
	tokenizer, session := make(chan string, 1), make(chan auth.SessionCredential, 1)
	var loginCredentials auth.LoginCredential

	if errBinding := c.ShouldBindJSON(&loginCredentials); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	employee, err := a.loginInteractor.Login(loginCredentials)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		go auth.GenerateToken(employee, tokenizer)
		go auth.GenerateSessionID(employee, session)
		sessionCoockie := <-session
		c.SetCookie("SID", sessionCoockie.ID, 60*30, "/", "localhost", true, true)
		c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.LoginDone, Data: &employee, Token: <-tokenizer})
	}
}

func (a *authRouter) logout(c *gin.Context) {
	err := auth.ClearSessionID(c)
	c.SetCookie("SID", "", -1, "/", "localhost", true, true)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.LogoutDone, Data: "Done"})
}
