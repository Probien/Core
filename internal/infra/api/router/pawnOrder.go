package router

import (
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/redis"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/gin-gonic/gin"
)

type IPawnOrderRouter interface {
	PawnOrderResource(g *gin.RouterGroup)
}

type PawnOrderRouter struct {
	auth          *component.Authenticator
	cookieManager port.ISessionRepository
	handler       handler.IPawnOrderHandler
}

func NewPawnOrderRouter(auth *component.Authenticator, cookieManager port.ISessionRepository, handler handler.IPawnOrderHandler) IPawnOrderRouter {
	return PawnOrderRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (p PawnOrderRouter) PawnOrderResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(p.auth, p.cookieManager, "ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE"))
	g.GET("/pawn-orders", p.handler.GetAll)
	g.GET("/pawn-orders/:id", p.handler.GetById)
	g.POST("/pawn-orders", p.handler.Create)
	g.PATCH("/pawn-orders", p.handler.Update)
}
