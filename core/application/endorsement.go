package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"
	"github.com/gin-gonic/gin"
)

type EndorsemenInteractor struct {
}

func (EI *EndorsemenInteractor) GetById(c *gin.Context) (*domain.Endorsement, error) {
	repository := postgres.NewEndorsementRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (EI *EndorsemenInteractor) GetAll(c *gin.Context) (*[]domain.Endorsement, error) {
	repository := postgres.NewEndorsementRepositoryImpl(config.Database)
	return repository.GetAll(c)
}

func (EI *EndorsemenInteractor) Create(c *gin.Context) (*domain.Endorsement, error) {
	repository := postgres.NewEndorsementRepositoryImpl(config.Database)
	return repository.Create(c)
}
