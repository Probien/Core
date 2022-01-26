package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/endorsement/domain"
	"github.com/JairDavid/Probien-Backend/core/endorsement/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

type EndorsemenInteractor struct {
}

func (EI *EndorsemenInteractor) GetById(c *gin.Context) (domain.Endorsement, error) {
	repository := persistance.NewEndorsementRepositoryImpl(config.GetDBInstance())
	return repository.GetById(c)
}

func (EI *EndorsemenInteractor) GetAll() ([]domain.Endorsement, error) {
	repository := persistance.NewEndorsementRepositoryImpl(config.GetDBInstance())
	return repository.GetAll()
}

func (EI *EndorsemenInteractor) Create(c *gin.Context) (domain.Endorsement, error) {
	repository := persistance.NewEndorsementRepositoryImpl(config.GetDBInstance())
	return repository.Create(c)
}
