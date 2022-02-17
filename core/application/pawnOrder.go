package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

type PawnOrderInteractor struct {
}

func (PI *PawnOrderInteractor) GetById(c *gin.Context) (*domain.PawnOrder, error) {
	repository := persistance.NewPawnOrderRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (PI *PawnOrderInteractor) GetAll() (*[]domain.PawnOrder, error) {
	repository := persistance.NewPawnOrderRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (PI *PawnOrderInteractor) Create(c *gin.Context) (*domain.PawnOrder, error) {
	repository := persistance.NewPawnOrderRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (PI *PawnOrderInteractor) Update(c *gin.Context) (*domain.PawnOrder, error) {
	repository := persistance.NewPawnOrderRepositoryImpl(config.Database)
	return repository.Update(c)
}
