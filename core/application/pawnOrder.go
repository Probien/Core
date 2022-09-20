package application

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
)

type PawnOrderInteractor struct {
}

func (PI *PawnOrderInteractor) GetById(id int) (*domain.PawnOrder, error) {
	repository := postgres.NewPawnOrderRepositoryImpl(config.Database)
	return repository.GetById(id)
}

func (PI *PawnOrderInteractor) GetAll(params url.Values) (*[]domain.PawnOrder, map[string]interface{}, error) {
	repository := postgres.NewPawnOrderRepositoryImpl(config.Database)
	return repository.GetAll(params)
}

func (PI *PawnOrderInteractor) Create(pawnOrderDto *domain.PawnOrder, userSessionId int) (*domain.PawnOrder, error) {
	repository := postgres.NewPawnOrderRepositoryImpl(config.Database)
	return repository.Create(pawnOrderDto, userSessionId)
}

func (PI *PawnOrderInteractor) Update(id int, pawnOrderDto map[string]interface{}, userSessionId int) (*domain.PawnOrder, error) {
	repository := postgres.NewPawnOrderRepositoryImpl(config.Database)
	return repository.Update(id, pawnOrderDto, userSessionId)
}
