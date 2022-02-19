package router

import (
	"github.com/JairDavid/Probien-Backend/core/interfaces"
	"github.com/gin-gonic/gin"
)

func Setup(ge *gin.Engine) {

	v1 := *ge.Group("/probien/api/v1")
	//v1.Use(authenticator.AuthJWT())
	{
		interfaces.CategoryHandler(&v1)
		interfaces.CustomerHandler(&v1)
		interfaces.EmployeeHandler(&v1)
		interfaces.EndorsementHandler(&v1)
		interfaces.PawnOrderHandler(&v1)
		interfaces.ProductHandler(&v1)
	}

}
