package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type IBranchOfficeRepository interface {
	GetAll(params url.Values) (*[]domain.BranchOffice, map[string]interface{}, error)
	GetById(id int) (*domain.BranchOffice, error)
	Create(branchOfficeDto *domain.BranchOffice, userSessionId int) (*domain.BranchOffice, error)
	Update(id int, patch map[string]interface{}, userSessionId int) (*domain.BranchOffice, error)
}
