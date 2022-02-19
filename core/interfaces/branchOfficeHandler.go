package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/gin-gonic/gin"
)

func BranchOfficeHandler(v1 *gin.RouterGroup) {

	branchOfficeHandlerv1 := *v1.Group("/branch-office")
	interactor := application.BranchOfficeInteractor{}

	branchOfficeHandlerv1.POST("/", func(c *gin.Context) {
		branchOffice, err := interactor.Create(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": &branchOffice})
	})

	branchOfficeHandlerv1.GET("/", func(c *gin.Context) {
		branchOffices, err := interactor.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": "something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &branchOffices})
		}
	})

	branchOfficeHandlerv1.GET("/:id", func(c *gin.Context) {
		branchOffice, err := interactor.GetById(c)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &branchOffice})
		}
	})

	branchOfficeHandlerv1.PATCH("/", func(c *gin.Context) {
		branchOffice, err := interactor.Update(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": &branchOffice})
		}
	})
}
