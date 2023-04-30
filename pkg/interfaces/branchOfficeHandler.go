package interfaces

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/persistence/postgres"
	"net/http"
	"strconv"

	"github.com/JairDavid/Probien-Backend/pkg/application"
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/interfaces/response"
	"github.com/gin-gonic/gin"
)

type branchRouter struct {
	branchInteractor application.BranchOfficeInteractor
}

func NewBranchOfficeHandler() *branchRouter {
	//dependency injection
	return &branchRouter{
		branchInteractor: application.NewBranchOfficeInteractor(postgres.NewBranchOfficeRepositoryImp(config.GetConnection())),
	}
}

func (b *branchRouter) SetupRouterAndFilter(v1 *gin.RouterGroup) {
	br := v1.Group("/").Use(auth.JwtRbac("ROLE_ADMIN", "ROLE_MANAGER"))
	br.POST("branch-offices", b.createBranch)
	br.GET("branch-offices", b.getAllBranches)
	br.GET("branch-offices/:id", b.getBranchById)
	br.PATCH("branch-offices", b.updateBranch)
}

func (b *branchRouter) createBranch(c *gin.Context) {
	var branchOfficeDto domain.BranchOffice
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.ShouldBindJSON(&branchOfficeDto); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	branchOffice, err := b.branchInteractor.Create(&branchOfficeDto, userSessionId.(int))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: response.Created, Data: &branchOffice})
	}
}

func (b *branchRouter) getAllBranches(c *gin.Context) {
	params := c.Request.URL.Query()
	branchOffices, paginationResult, err := b.branchInteractor.GetAll(params)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Response{Status: http.StatusInternalServerError, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"})
	} else {
		c.JSON(http.StatusOK, response.PaginatedResponse{Status: http.StatusOK, ItemsPerPage: 10, TotalPages: int(paginationResult["total_pages"].(float64)), CurrentPage: paginationResult["page"].(int), Data: &branchOffices, Previous: "localhost:9000/api/v1/branch-offices/?page=" + paginationResult["previous"].(string), Next: "localhost:9000/api/v1/branch-offices/?page=" + paginationResult["next"].(string)})
	}
}

func (b *branchRouter) getBranchById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	branchOffice, err := b.branchInteractor.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: response.Consulted, Data: &branchOffice})
	}
}

func (b *branchRouter) updateBranch(c *gin.Context) {
	requestBodyWithId := map[string]interface{}{}
	//Obtained from decoded token (middleware)
	userSessionId, _ := c.Get("user_id")

	if errBinding := c.Bind(&requestBodyWithId); errBinding != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: errBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	id, errID := requestBodyWithId["id"]

	if !errID {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: response.ErrorBinding.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
		return
	}

	branchOffice, err := b.branchInteractor.Update(int(id.(float64)), requestBodyWithId, userSessionId.(int))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Response{Status: http.StatusBadRequest, Message: response.FailedHttpOperation, Data: err.Error(), Help: "https://probien/api/v1/swagger-ui.html"},
		)
	} else {
		c.JSON(http.StatusAccepted, response.Response{Status: http.StatusAccepted, Message: response.Updated, Data: &branchOffice})
	}
}
