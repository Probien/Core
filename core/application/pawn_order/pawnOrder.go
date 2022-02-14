package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	pawn_order_domain "github.com/JairDavid/Probien-Backend/core/domain/pawn_order"
	pawn_order_infra "github.com/JairDavid/Probien-Backend/core/infrastructure/pawn_order"
	"github.com/gin-gonic/gin"
)

type PawnOrderInteractor struct {
}

func (PI *PawnOrderInteractor) GetById(c *gin.Context) (*pawn_order_domain.PawnOrder, error) {
	repository := pawn_order_infra.NewPawnOrderRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (PI *PawnOrderInteractor) GetAll() (*[]pawn_order_domain.PawnOrder, error) {
	repository := pawn_order_infra.NewPawnOrderRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (PI *PawnOrderInteractor) Create(c *gin.Context) (*pawn_order_domain.PawnOrder, error) {
	repository := pawn_order_infra.NewPawnOrderRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (PI *PawnOrderInteractor) Update(c *gin.Context) (*pawn_order_domain.PawnOrder, error) {
	repository := pawn_order_infra.NewPawnOrderRepositoryImpl(config.Database)
	return repository.Update(c)
}
