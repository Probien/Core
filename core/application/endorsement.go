package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type EndorsemenInteractor struct {
}

func (EI *EndorsemenInteractor) GetById(id int) (*domain.Endorsement, error) {
	repository := postgres.NewEndorsementRepositoryImpl(config.Database)
	return repository.GetById(id)
}

func (EI *EndorsemenInteractor) GetAll(params url.Values) (*[]domain.Endorsement, map[string]interface{}, error) {
	repository := postgres.NewEndorsementRepositoryImpl(config.Database)
	return repository.GetAll(params)
}

func (EI *EndorsemenInteractor) Create(endorsementDto *domain.Endorsement, userSessionId int) (*domain.Endorsement, error) {
	repository := postgres.NewEndorsementRepositoryImpl(config.Database)
	return repository.Create(endorsementDto, userSessionId)
}
