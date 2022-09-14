package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
	"github.com/gin-gonic/gin"
)

type BranchOfficeInteractor struct {
}

func (BI *BranchOfficeInteractor) GetAll(c *gin.Context) (*[]domain.BranchOffice, map[string]interface{}, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.GetAll(c)
}

func (BI *BranchOfficeInteractor) GetById(c *gin.Context) (*domain.BranchOffice, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.GetById(c)

}

func (BI *BranchOfficeInteractor) Create(c *gin.Context) (*domain.BranchOffice, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.Create(c)

}
func (BI *BranchOfficeInteractor) Update(c *gin.Context) (*domain.BranchOffice, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.Update(c)

}
