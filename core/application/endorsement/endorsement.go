package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	endorsement_domain "github.com/JairDavid/Probien-Backend/core/domain/endorsement"
	endorsement_infra "github.com/JairDavid/Probien-Backend/core/infrastructure/endorsement"
	"github.com/gin-gonic/gin"
)

type EndorsemenInteractor struct {
}

func (EI *EndorsemenInteractor) GetById(c *gin.Context) (*endorsement_domain.Endorsement, error) {
	repository := endorsement_infra.NewEndorsementRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (EI *EndorsemenInteractor) GetAll() (*[]endorsement_domain.Endorsement, error) {
	repository := endorsement_infra.NewEndorsementRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (EI *EndorsemenInteractor) Create(c *gin.Context) (*endorsement_domain.Endorsement, error) {
	repository := endorsement_infra.NewEndorsementRepositoryImpl(config.Database)
	return repository.Create(c)
}
