package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type branchRouter struct {
	branchInteractor application.BranchOfficeInteractor
}

func BranchOfficeHandler(v1 *gin.RouterGroup) {
	var branchRouter branchRouter

	v1.POST("/", branchRouter.createBranch)
	v1.GET("/", branchRouter.getAllBranches)
	v1.GET("/:id", branchRouter.getBranchById)
	v1.PATCH("/", branchRouter.updateBranch)

}

func (router *branchRouter) createBranch(c *gin.Context) {
	branchOffice, err := router.branchInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &branchOffice})
	}
}

func (router *branchRouter) getAllBranches(c *gin.Context) {
	branchOffices, err := router.branchInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"})
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &branchOffices})
	}
}

func (router *branchRouter) getBranchById(c *gin.Context) {
	branchOffice, err := router.branchInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &branchOffice})
	}
}

func (router *branchRouter) updateBranch(c *gin.Context) {
	branchOffice, err := router.branchInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.UPDATED, Data: &branchOffice})
	}
}
