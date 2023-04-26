package application

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"net/url"
)

type BranchOfficeInteractor struct {
	repository repository.IBranchOfficeRepository
}

func NewBranchOfficeInteractor(repository repository.IBranchOfficeRepository) BranchOfficeInteractor {
	return BranchOfficeInteractor{
		repository: repository,
	}
}

func (b *BranchOfficeInteractor) GetAll(params url.Values) (*[]domain.BranchOffice, map[string]interface{}, error) {
	return b.repository.GetAll(params)
}

func (b *BranchOfficeInteractor) GetById(id int) (*domain.BranchOffice, error) {
	return b.repository.GetById(id)
}

func (b *BranchOfficeInteractor) Create(branchOfficeDto *domain.BranchOffice, userSessionId int) (*domain.BranchOffice, error) {
	return b.repository.Create(branchOfficeDto, userSessionId)
}

func (b *BranchOfficeInteractor) Update(id int, patch map[string]interface{}, userSessionId int) (*domain.BranchOffice, error) {
	return b.repository.Update(id, patch, userSessionId)
}
