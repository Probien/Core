package router

import (
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces"
	"github.com/gin-gonic/gin"
)

func Setup(server *gin.Engine) {
	api := server.Group("/probien/api/v1")
	{
		interfaces.AuthHandler(api.Group("/auth"))

		interfaces.ProductHandler(api.Group("/products").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE")).(*gin.RouterGroup))

		interfaces.CategoryHandler(api.Group("/categories").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE")).(*gin.RouterGroup))

		interfaces.BranchOfficeHandler(api.Group("/branch-offices").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER")).(*gin.RouterGroup))

		interfaces.EmployeeHandler(api.Group("/employees").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER")).(*gin.RouterGroup))

		interfaces.CustomerHandler(api.Group("/customers").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE")).(*gin.RouterGroup))

		interfaces.PawnOrderHandler(api.Group("/pawn-orders").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE")).(*gin.RouterGroup))

		interfaces.EndorsementHandler(api.Group("/endorsements").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER", "ROLE_EMPLOYEE")).(*gin.RouterGroup))

		interfaces.LogHandler(api.Group("/logs").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER")).(*gin.RouterGroup))
	}
}
