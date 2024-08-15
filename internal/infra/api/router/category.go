package router

import (
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/redis"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/gin-gonic/gin"
)

type ICategoryRouter interface {
	CategoryResource(g *gin.RouterGroup)
}

type CategoryRouter struct {
	auth          *component.Authenticator
	cookieManager port.ISessionRepository
	handler       handler.ICategoryHandler
}

func NewCategoryHandler(auth *component.Authenticator, cookieManager port.ISessionRepository, handler handler.ICategoryHandler) ICategoryRouter {
	return CategoryRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (c CategoryRouter) CategoryResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(c.auth, c.cookieManager, "ROLE_ADMIN", "ROLE_MANAGER"))
	g.GET("/categories", c.handler.GetAll)
	g.GET("/categories/:id", c.handler.GetById)
	g.POST("/categories", c.handler.Create)
	g.DELETE("/categories", c.handler.Delete)
	g.PATCH("/categories", c.handler.Update)
}
