package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type PawnOrderApp struct {
	port port.IPawnOrderRepository
}

func NewPawnOrderApp(repository port.IPawnOrderRepository) PawnOrderApp {
	return PawnOrderApp{
		port: repository,
	}
}

func (p *PawnOrderApp) GetById(id int) (*dto.PawnOrder, error) {
	return p.port.GetById(id)
}

func (p *PawnOrderApp) GetAll(params url.Values) (*[]dto.PawnOrder, map[string]interface{}, error) {
	return p.port.GetAll(params)
}

func (p *PawnOrderApp) Create(pawnOrderDto *dto.PawnOrder, userSessionId int) (*dto.PawnOrder, error) {
	return p.port.Create(pawnOrderDto, userSessionId)
}

func (p *PawnOrderApp) Update(id int, pawnOrderDto map[string]interface{}, userSessionId int) (*dto.PawnOrder, error) {
	return p.port.Update(id, pawnOrderDto, userSessionId)
}
