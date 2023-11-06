package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type IEndorsementRepository interface {
	GetById(id int) (*dto.Endorsement, error)
	GetAll(params url.Values) (*[]dto.Endorsement, map[string]interface{}, error)
	Create(endorsementDto *dto.Endorsement, userSessionId int) (*dto.Endorsement, error)
}
