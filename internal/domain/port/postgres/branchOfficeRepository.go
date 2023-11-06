package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type IBranchOfficeRepository interface {
	GetAll(params url.Values) (*[]dto.BranchOffice, map[string]interface{}, error)
	GetById(id int) (*dto.BranchOffice, error)
	Create(branchOfficeDto *dto.BranchOffice, userSessionId int) (*dto.BranchOffice, error)
	Update(id int, patch map[string]interface{}, userSessionId int) (*dto.BranchOffice, error)
}
