package router

import (
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/redis"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/gin-gonic/gin"
)

type IEndorsementRouter interface {
	EndorsementResource(g *gin.RouterGroup)
}

type EndorsementRouter struct {
	auth          *component.Authenticator
	cookieManager port.ISessionRepository
	handler       handler.IEndorsementHandler
}

func NewEndorsementRouter(auth *component.Authenticator, cookieManager port.ISessionRepository, handler handler.IEndorsementHandler) IEndorsementRouter {
	return EndorsementRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (e EndorsementRouter) EndorsementResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(e.auth, e.cookieManager, "ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE"))
	g.GET("/endorsements", e.handler.GetAll)
	g.GET("/endorsements/:id", e.handler.GetById)
	g.POST("/endorsements", e.handler.Create)
}
