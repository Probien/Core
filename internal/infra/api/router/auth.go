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
	//TODO implement me
	panic("implement me")
}
