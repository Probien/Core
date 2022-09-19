package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type BranchOfficeInteractor struct {
}

func (BI *BranchOfficeInteractor) GetAll(params url.Values) (*[]domain.BranchOffice, map[string]interface{}, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.GetAll(params)
}

func (BI *BranchOfficeInteractor) GetById(id int) (*domain.BranchOffice, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.GetById(id)

}

func (BI *BranchOfficeInteractor) Create(branchOfficeDto *domain.BranchOffice, userSessionId int) (*domain.BranchOffice, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.Create(branchOfficeDto, userSessionId)

}
func (BI *BranchOfficeInteractor) Update(id int, patch map[string]interface{}, userSessionId int) (*domain.BranchOffice, error) {
	repository := postgres.NewBranchOfficeRepositoryImp(config.Database)
	return repository.Update(id, patch, userSessionId)

}
