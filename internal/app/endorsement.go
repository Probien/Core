package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type EndorsementApp struct {
	port port.IEndorsementRepository
}

func NewEndorsementApp(repository port.IEndorsementRepository) EndorsementApp {
	return EndorsementApp{
		port: repository,
	}
}

func (e *EndorsementApp) GetById(id int) (*dto.Endorsement, error) {
	return e.port.GetById(id)
}

func (e *EndorsementApp) GetAll(params url.Values) (*[]dto.Endorsement, map[string]interface{}, error) {
	return e.port.GetAll(params)
}

func (e *EndorsementApp) Create(endorsementDto *dto.Endorsement, userSessionId int) (*dto.Endorsement, error) {
	return e.port.Create(endorsementDto, userSessionId)
}
