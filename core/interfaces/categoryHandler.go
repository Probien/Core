package interfaces

import (
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/core/application"
	"github.com/JairDavid/Probien-Backend/core/domain"
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
	var categoryDto domain.Category
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&categoryDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	category, err := router.categoryInteractor.Create(&categoryDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, common.Response{Status: http.StatusCreated, Message: common.Created, Data: &category})
	}
}

func (router *categoryRouter) getAllCategories(c *gin.Context) {
	params := c.Request.URL.Query()
	categories, paginationResult, err := router.categoryInteractor.GetAll(params)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			common.Response{Status: http.StatusInternalServerError, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &categories, Previous: "localhost:9000/api/v1/categories/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/categories/?page=" + paginationResult["next"].(string)})
	}
}

func (router *categoryRouter) getCategoryById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := router.categoryInteractor.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			common.Response{Status: http.StatusNotFound, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, common.Response{Status: http.StatusOK, Message: common.Consulted, Data: &category})
	}
}

func (router *categoryRouter) updateCategory(c *gin.Context) {
	requestBodyWithId := map[string]interface{}{}
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.Bind(&requestBodyWithId); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	id, errID := requestBodyWithId["id"]

	if !errID {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: common.ErrorBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	category, err := router.categoryInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Updated, Data: &category})
	}
}

func (router *categoryRouter) deleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")
	category, err := router.categoryInteractor.Delete(id, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.Response{Status: http.StatusBadRequest, Message: common.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, common.Response{Status: http.StatusAccepted, Message: common.Deleted, Data: &category})
	}
}
