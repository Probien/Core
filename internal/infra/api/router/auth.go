package router

import "github.com/gin-gonic/gin"

type IAuthRouter interface {
	AuthResource(g *gin.RouterGroup)
}

type AuthRouter struct {
}

func NewAuth() IAuthRouter {
	return &AuthRouter{}
}

func (a *AuthRouter) AuthResource(g *gin.RouterGroup) {
	//TODO implement me
	panic("implement me")
}
