package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type branchRouter struct {
	branchInteractor application.BranchOfficeInteractor
}

func BranchOfficeHandler(v1 *gin.RouterGroup) {

	var branchRouter branchRouter
	branchOfficeHandlerv1 := *v1.Group("/branch-offices")
	branchOfficeHandlerv1.Use(auth.JwtAuth(false))

	branchOfficeHandlerv1.POST("/", branchRouter.createBranch)
	branchOfficeHandlerv1.GET("/", branchRouter.getAllBranches)
	branchOfficeHandlerv1.GET("/:id", branchRouter.getBranchById)
	branchOfficeHandlerv1.PATCH("/", branchRouter.updateBranch)
}

func (bi *branchRouter) createBranch(c *gin.Context) {
	branchOffice, err := bi.branchInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &branchOffice})
	}
}

func (bi *branchRouter) getAllBranches(c *gin.Context) {
	branchOffices, err := bi.branchInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"})
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &branchOffices})
	}
}

func (bi *branchRouter) getBranchById(c *gin.Context) {
	branchOffice, err := bi.branchInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &branchOffice})
	}
}

func (bi *branchRouter) updateBranch(c *gin.Context) {
	branchOffice, err := bi.branchInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.CONSULTED, Data: &branchOffice})
	}
}
