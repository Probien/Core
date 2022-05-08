package router

import (
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces"
	"github.com/gin-gonic/gin"
)

func Setup(server *gin.Engine) {
	v1 := *server.Group("/probien/api/v1")
	interfaces.AuthHandler(&v1)

	v1.Use(auth.AuthJWT())
	{
		interfaces.EmployeeHandler(&v1)
		interfaces.BranchOfficeHandler(&v1)
		interfaces.CategoryHandler(&v1)
		interfaces.CustomerHandler(&v1)
		interfaces.EndorsementHandler(&v1)
		interfaces.PawnOrderHandler(&v1)
		interfaces.ProductHandler(&v1)
		interfaces.LogHandler(&v1)
	}
}
