package router

import (
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/redis"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/middleware"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/gin-gonic/gin"
)

type IBranchOfficeRouter interface {
	BranchOfficeResource(g *gin.RouterGroup)
}

type BranchOfficeRouter struct {
	auth          *component.Authenticator
	cookieManager port.ISessionRepository
	handler       handler.IBranchOfficeHandler
}

func NewBranchOfficeRouter(auth *component.Authenticator, cookieManager port.ISessionRepository, handler handler.IBranchOfficeHandler) IBranchOfficeRouter {
	return &BranchOfficeRouter{
		auth:          auth,
		cookieManager: cookieManager,
		handler:       handler,
	}
}

func (b BranchOfficeRouter) BranchOfficeResource(g *gin.RouterGroup) {
	g.Use(middleware.JwtRbac(b.auth, b.cookieManager, "ROLE_ADMIN", "ROLE_MANAGER"))
	g.POST("/branch-offices", b.handler.Create)
	g.GET("/branch-offices", b.handler.GetAll)
	g.GET("/branch-offices/:id", b.handler.GetById)
	g.PATCH("/branch-offices", b.handler.Update)
}
