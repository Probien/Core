package application

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"net/url"
)

type PawnOrderInteractor struct {
	repository repository.IPawnOrderRepository
}

func NewPawnOrderInteractor(repository repository.IPawnOrderRepository) PawnOrderInteractor {
	return PawnOrderInteractor{
		repository: repository,
	}
}

func (p *PawnOrderInteractor) GetById(id int) (*domain.PawnOrder, error) {
	return p.repository.GetById(id)
}

func (p *PawnOrderInteractor) GetAll(params url.Values) (*[]domain.PawnOrder, map[string]interface{}, error) {
	return p.repository.GetAll(params)
}

func (p *PawnOrderInteractor) Create(pawnOrderDto *domain.PawnOrder, userSessionId int) (*domain.PawnOrder, error) {
	return p.repository.Create(pawnOrderDto, userSessionId)
}

func (p *PawnOrderInteractor) Update(id int, pawnOrderDto map[string]interface{}, userSessionId int) (*domain.PawnOrder, error) {
	return p.repository.Update(id, pawnOrderDto, userSessionId)
}
