package utils

import (
	category "github.com/JairDavid/Probien-Backend/core/category/interfaces"
	customer "github.com/JairDavid/Probien-Backend/core/customer/interfaces"
	employee "github.com/JairDavid/Probien-Backend/core/employee/interfaces"
	endorsement "github.com/JairDavid/Probien-Backend/core/endorsement/interfaces"
	pawnOrder "github.com/JairDavid/Probien-Backend/core/pawn_order/interfaces"
	product "github.com/JairDavid/Probien-Backend/core/product/interfaces"
	"github.com/gin-gonic/gin"
)

func Setup(s *gin.Engine) {

	v1 := *s.Group("/probien/api/v1")
	{
		category.CategoryHandler(&v1)
		customer.CustomerHandler(&v1)
		employee.EmployeeHandler(&v1)
		endorsement.EndorsementHandler(&v1)
		pawnOrder.PawnOrderHandler(&v1)
		product.ProductHandler(&v1)
	}

}
