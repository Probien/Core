package application

import (
	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/domain/repository"
	"net/url"
)

type EndorsementInteractor struct {
	repository repository.IEndorsementRepository
}

func NewEndorsementInteractor(repository repository.IEndorsementRepository) EndorsementInteractor {
	return EndorsementInteractor{
		repository: repository,
	}
}

func (e *EndorsementInteractor) GetById(id int) (*domain.Endorsement, error) {
	return e.repository.GetById(id)
}

func (e *EndorsementInteractor) GetAll(params url.Values) (*[]domain.Endorsement, map[string]interface{}, error) {
	return e.repository.GetAll(params)
}

func (e *EndorsementInteractor) Create(endorsementDto *domain.Endorsement, userSessionId int) (*domain.Endorsement, error) {
	return e.repository.Create(endorsementDto, userSessionId)
}
