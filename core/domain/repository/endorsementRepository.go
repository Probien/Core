package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type IEndorsementRepository interface {
	GetById(id int) (*domain.Endorsement, error)
	GetAll(params url.Values) (*[]domain.Endorsement, map[string]interface{}, error)
	Create(endorsementDto *domain.Endorsement, userSessionId int) (*domain.Endorsement, error)
}
