package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type BranchOfficeApp struct {
	port port.IBranchOfficeRepository
}

func NewBranchOfficeApp(repository port.IBranchOfficeRepository) BranchOfficeApp {
	return BranchOfficeApp{
		port: repository,
	}
}

func (b *BranchOfficeApp) GetAll(params url.Values) (*[]dto.BranchOffice, map[string]interface{}, error) {
	return b.port.GetAll(params)
}

func (b *BranchOfficeApp) GetById(id int) (*dto.BranchOffice, error) {
	return b.port.GetById(id)
}

func (b *BranchOfficeApp) Create(branchOfficeDto *dto.BranchOffice, userSessionId int) (*dto.BranchOffice, error) {
	return b.port.Create(branchOfficeDto, userSessionId)
}

func (b *BranchOfficeApp) Update(id int, patch map[string]interface{}, userSessionId int) (*dto.BranchOffice, error) {
	return b.port.Update(id, patch, userSessionId)
}
