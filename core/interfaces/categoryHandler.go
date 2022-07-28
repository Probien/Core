package interfaces

import (
	"net/http"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/interfaces/common"
	"github.com/gin-gonic/gin"
)

type categoryRouter struct {
	categoryInteractor application.CategoryInteractor
}

func CategoryHandler(v1 *gin.RouterGroup) {

	var categoryRouter categoryRouter
	categoryHandlerV1 := *v1.Group("/categories")
	categoryHandlerV1.Use(auth.JwtAuth(false))

	categoryHandlerV1.POST("/", categoryRouter.CreateCategory)
	categoryHandlerV1.GET("/", categoryRouter.getAllCategories)
	categoryHandlerV1.GET("/:id", categoryRouter.getCategoryById)
	categoryHandlerV1.PATCH("/", categoryRouter.updateCategory)
	categoryHandlerV1.DELETE("/", categoryRouter.deleteCategory)
}

func (ci *categoryRouter) CreateCategory(c *gin.Context) {
	category, err := ci.categoryInteractor.Create(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.CREATED, Data: &category})
	}
}

func (ci *categoryRouter) getAllCategories(c *gin.Context) {
	categories, err := ci.categoryInteractor.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &categories})
	}
}

func (ci *categoryRouter) getCategoryById(c *gin.Context) {
	category, err := ci.categoryInteractor.GetById(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.CONSULTED, Data: &category})
	}
}

func (ci *categoryRouter) updateCategory(c *gin.Context) {
	category, err := ci.categoryInteractor.Update(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.UPDATED, Data: &category})
	}
}

func (ci *categoryRouter) deleteCategory(c *gin.Context) {
	category, err := ci.categoryInteractor.Delete(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FAILED_HTTP_OPERATION, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.DELETED, Data: &category})
	}
}
