package router

import (
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/redis"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/gin-gonic/gin"
)

type ILogRouter interface {
	LogResource(g *gin.RouterGroup)
}

type LogRouter struct {
	auth          *component.Authenticator
	cookieManager port.ISessionRepository
	handler       handler.ILogHandler
}

func NewLogRouter(auth *component.Authenticator, cookieManager port.ISessionRepository, handler handler.ILogHandler) ILogRouter {
	return LogRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (l LogRouter) LogResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(l.auth, l.cookieManager, "ROLE_ADMIN"))
	g.GET("/logs/movements", l.handler.GetAllMovements)
	g.GET("/logs/movements/employees", l.handler.GetAllMovementsByEmployeeId)
	g.POST("/logs/sessions", l.handler.GetAllSessions)
	g.PATCH("/logs/sessions/employees", l.handler.GetAllSessionsByEmployeeId)
}
