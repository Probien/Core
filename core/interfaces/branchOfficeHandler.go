package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

func BranchOfficeHandler(v1 *gin.RouterGroup) {

	branchOfficeHandlerv1 := *v1.Group("/branch-offices")
	branchOfficeHandlerv1.Use(auth.RoutesAndAuthority(false))
	interactor := application.BranchOfficeInteractor{}

	branchOfficeHandlerv1.POST("/", func(c *gin.Context) {
		branchOffice, err := interactor.Create(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: "successfully created", Data: &branchOffice})
		}
	})

	branchOfficeHandlerv1.GET("/", func(c *gin.Context) {
		branchOffices, err := interactor.GetAll()

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.Response{Status: http.StatusInternalServerError, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"})
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &branchOffices})
		}
	})

	branchOfficeHandlerv1.GET("/:id", func(c *gin.Context) {
		branchOffice, err := interactor.GetById(c)

		if err != nil {
			c.JSON(
				http.StatusNotFound,
				common.Response{Status: http.StatusNotFound, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: "successfully consulted", Data: &branchOffice})
		}
	})

	branchOfficeHandlerv1.PATCH("/", func(c *gin.Context) {
		branchOffice, err := interactor.Update(c)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.Response{Status: http.StatusBadRequest, Message: "failed operation", Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
			)
		} else {
			c.JSON(http.StatusOK, common.Response{Status: http.StatusAccepted, Message: "successfully updated", Data: &branchOffice})
		}
	})
}
