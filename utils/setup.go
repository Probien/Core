package utils

import (
	"github.com/JairDavid/Probien-Backend/core/interfaces"
	"github.com/gin-gonic/gin"
)

func Setup(s *gin.Engine) {

	v1 := *s.Group("/probien/api/v1")
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
