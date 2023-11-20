package router

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/JairDavid/Probien-Backend/internal/infra/resource/redis"
	"github.com/gin-gonic/gin"
)

type IProductRouter interface {
	ProductResource(g *gin.RouterGroup)
}

type ProductRouter struct {
	auth          *component.Authenticator
	cookieManager *redis.Client
	handler       handler.IProductHandler
}

func NewProductRouter(auth *component.Authenticator, cookieManager *redis.Client, handler handler.IProductHandler) IProductRouter {
	return ProductRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (p ProductRouter) ProductResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(p.auth, p.cookieManager, "ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE"))
	g.GET("/products", p.handler.GetAll)
	g.GET("/products/:id", p.handler.GetById)
	g.POST("/products", p.handler.Create)
	g.PATCH("/products", p.handler.Update)
}
