package application

import (
	"github.com/JairDavid/Probien-Backend/core/pawn_order/domain"
	"github.com/gin-gonic/gin"
)

type PawnOrderInteractor struct {
	repository domain.PawnOrderRepository
}

func (PI *PawnOrderInteractor) GetById(c *gin.Context) (domain.PawnOrder, error) {
	return PI.repository.GetById(c)
}

func (PI *PawnOrderInteractor) GetAll() ([]domain.PawnOrder, error) {
	return PI.repository.GetAll()
}

func (PI *PawnOrderInteractor) Create(c *gin.Context) (domain.PawnOrder, error) {
	return PI.repository.Create(c)
}

func (PI *PawnOrderInteractor) Update(c *gin.Context) (domain.PawnOrder, error) {
	return PI.repository.Update(c)
}
