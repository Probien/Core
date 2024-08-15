package router

import (
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/redis"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/gin-gonic/gin"
)

type IEmployeeRouter interface {
	EmployeeResource(g *gin.RouterGroup)
}

type EmployeeRouter struct {
	auth          *component.Authenticator
	cookieManager port.ISessionRepository
	handler       handler.IEmployeeHandler
}

func NewEmployeeRouter(auth *component.Authenticator, cookieManager port.ISessionRepository, handler handler.IEmployeeHandler) IEmployeeRouter {
	return EmployeeRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (e EmployeeRouter) EmployeeResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(e.auth, e.cookieManager, "ROLE_ADMIN", "ROLE_MANAGER"))
	g.GET("/employees", e.handler.GetAll)
	g.GET("/employees/emails", e.handler.GetByEmail)
	g.POST("/employees", e.handler.Create)
	g.PATCH("/employees", e.handler.Update)
}
