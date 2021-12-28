package application

import (
	"github.com/JairDavid/Probien-Backend/core/endorsement/domain"
	"github.com/gin-gonic/gin"
)

type EndorsemenInteractor struct {
	repository domain.EndorsementRepository
}

func (EI *EndorsemenInteractor) GetById(c *gin.Context) (domain.Endorsement, error) {
	return EI.repository.GetById(c)
}

func (EI *EndorsemenInteractor) GetAll() ([]domain.Endorsement, error) {
	return EI.repository.GetAll()
}

func (EI *EndorsemenInteractor) Create(c *gin.Context) (domain.Endorsement, error) {
	return EI.repository.Create(c)
}
