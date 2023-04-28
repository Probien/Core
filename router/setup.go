package router

import (
	"github.com/JairDavid/Probien-Backend/core/interfaces"
	"github.com/gin-gonic/gin"
)

// Setup :Secure all routes with user authorities
func Setup(server *gin.Engine) {
	api := server.Group("/api/v1")
	{
		interfaces.NewAuthHandler().SetupRoutes(api)
		interfaces.NewProductHandler().SetupRoutesAndFilter(api)
		interfaces.NewCategoryHandler().SetupRoutesAndFilter(api)
		interfaces.NewBranchOfficeHandler().SetupRouterAndFilter(api)
		interfaces.NewEmployeeHandler().SetupRoutesAndFilter(api)
		interfaces.CustomerHandler().SetupRoutesAndFilter(api)
		interfaces.NewPawnOrderHandler().SetupRoutesAndFilter(api)
		interfaces.NewEndorsementHandler().SetupRoutesAndFilter(api)
		interfaces.LogHandler().SetupRoutesAndFilter(api)
	}
}
