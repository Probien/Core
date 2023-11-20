package router

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/JairDavid/Probien-Backend/internal/infra/resource/redis"
	"github.com/gin-gonic/gin"
)

type ICustomerRouter interface {
	CustomerResource(g *gin.RouterGroup)
}

type CustomerRouter struct {
	auth          *component.Authenticator
	cookieManager *redis.Client
	handler       handler.ICustomerHandler
}

func NewCustomerRouter(auth *component.Authenticator, cookieManager *redis.Client, handler handler.ICustomerHandler) ICustomerRouter {
	return CustomerRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (c CustomerRouter) CustomerResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(c.auth, c.cookieManager, "ROLE_ADMIN", "ROLE_MANAGER"))
	g.GET("/customers", c.handler.GetAll)
	g.GET("/customers/:id", c.handler.GetById)
	g.POST("/customers", c.handler.Create)
	g.PATCH("/customers", c.handler.Update)
}
