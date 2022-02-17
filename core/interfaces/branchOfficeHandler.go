package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/gin-gonic/gin"
)

func BranchOfficeHandler(v1 *gin.RouterGroup) {
	branchOfficeHandlerv1 := *v1.Group("/branch-office")
	interactor := application.BranchOfficeInteractor{}

	branchOfficeHandlerv1.GET("/", func(c *gin.Context) {
		employees, err := interactor.GetAll()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": &employees})
	})

	branchOfficeHandlerv1.GET("/:id", func(c *gin.Context) {

	})

	branchOfficeHandlerv1.POST("/", func(c *gin.Context) {

	})

	branchOfficeHandlerv1.PATCH("/", func(c *gin.Context) {

	})
}
