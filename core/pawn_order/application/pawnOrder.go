package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/pawn_order/domain"
	"github.com/JairDavid/Probien-Backend/core/pawn_order/infrastructure/persistance"
)

func GetById() (domain.PawnOrder, error) {
	return persistance.NewPawnOrderRepositoryImpl(config.GetDBInstance()).GetById()
}

func GetAll() ([]domain.PawnOrder, error) {
	return persistance.NewPawnOrderRepositoryImpl(config.GetDBInstance()).GetAll()
}

func Create() (domain.PawnOrder, error) {
	return persistance.NewPawnOrderRepositoryImpl(config.GetDBInstance()).Create()
}

func Update() (domain.PawnOrder, error) {
	return persistance.NewPawnOrderRepositoryImpl(config.GetDBInstance()).Update()
}
