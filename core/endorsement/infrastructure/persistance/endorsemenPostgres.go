package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/endorsement/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EndorsementRepositoryImpl struct {
	database *gorm.DB
}

func NewEndorsementRepositoryImpl(db *gorm.DB) domain.EndorsementRepository {
	return &EndorsementRepositoryImpl{database: db}
}

func (r *EndorsementRepositoryImpl) GetById(c *gin.Context) (domain.Endorsement, error) {
	return domain.Endorsement{}, nil
}

func (r *EndorsementRepositoryImpl) GetAll() ([]domain.Endorsement, error) {
	return []domain.Endorsement{}, nil
}

func (r *EndorsementRepositoryImpl) Create(c *gin.Context) (domain.Endorsement, error) {
	return domain.Endorsement{}, nil
}
