package router

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/gin-gonic/gin"
)

type IAuthRouter interface {
	AuthResource(g *gin.RouterGroup)
}

type AuthRouter struct {
	authHandler handler.IAuthHandler
}

func NewAuthRouter(authHandler handler.IAuthHandler) IAuthRouter {
	return &AuthRouter{
		authHandler: authHandler,
	}
}

func (a *AuthRouter) AuthResource(g *gin.RouterGroup) {
	g.POST("/login", a.authHandler.Login)
	g.POST("/logout", a.authHandler.Logout)
	g.GET("/password-recovery", a.authHandler.RecoverPassword)
}
