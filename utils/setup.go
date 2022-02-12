package utils

import (
	category_interface "github.com/JairDavid/Probien-Backend/core/interfaces/category"
	customer_interface "github.com/JairDavid/Probien-Backend/core/interfaces/customer"
	employee_interface "github.com/JairDavid/Probien-Backend/core/interfaces/employee"
	endorsement_interface "github.com/JairDavid/Probien-Backend/core/interfaces/endorsement"
	pawn_order_interface "github.com/JairDavid/Probien-Backend/core/interfaces/pawn_order"
	product_interface "github.com/JairDavid/Probien-Backend/core/interfaces/product"
	"github.com/gin-gonic/gin"
)

func Setup(s *gin.Engine) {

	v1 := *s.Group("/probien/api/v1")
	//v1.Use(authenticator.AuthJWT())
	{
		category_interface.CategoryHandler(&v1)
		customer_interface.CustomerHandler(&v1)
		employee_interface.EmployeeHandler(&v1)
		endorsement_interface.EndorsementHandler(&v1)
		pawn_order_interface.PawnOrderHandler(&v1)
		product_interface.ProductHandler(&v1)
	}

}
