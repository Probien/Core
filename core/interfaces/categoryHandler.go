package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type categoryRouter struct {
	categoryInteractor application.CategoryInteractor
}

func CategoryHandler(v1 *gin.RouterGroup) {
	var categoryRouter categoryRouter

	v1.POST("/", categoryRouter.CreateCategory)
	v1.GET("/", categoryRouter.getAllCategories)
	v1.GET("/:id", categoryRouter.getCategoryById)
	v1.PATCH("/", categoryRouter.updateCategory)
	v1.DELETE("/:id", categoryRouter.deleteCategory)
}

func (router *categoryRouter) CreateCategory(c *gin.Context) {
	category, err := router.categoryInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &category})
	}
}

func (router *categoryRouter) getAllCategories(c *gin.Context) {
	categories, err := router.categoryInteractor.GetAll(c)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &categories})
	}
}

func (router *categoryRouter) getCategoryById(c *gin.Context) {
	category, err := router.categoryInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &category})
	}
}

func (router *categoryRouter) updateCategory(c *gin.Context) {
	category, err := router.categoryInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &category})
	}
}

func (router *categoryRouter) deleteCategory(c *gin.Context) {
	category, err := router.categoryInteractor.Delete(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Deleted, Data: &category})
	}
}
