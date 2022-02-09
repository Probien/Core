package persistance

import (
	"errors"

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

func (r *EndorsementRepositoryImpl) GetById(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	r.database.Model(&domain.Endorsement{}).Preload("PawnOrderID").Find(&endorsement, c.Param("id"))
	if endorsement.ID == 0 {
		return nil, errors.New("endorsement not found")
	}
	return &endorsement, nil
}

func (r *EndorsementRepositoryImpl) GetAll() (*[]domain.Endorsement, error) {
	var endorsements []domain.Endorsement

	r.database.Model(&domain.Endorsement{}).Find(&endorsements)
	return &endorsements, nil
}

func (r *EndorsementRepositoryImpl) Create(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := c.ShouldBindJSON(&endorsement); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}
	r.database.Model(&domain.Endorsement{}).Create(&endorsement)

	return &endorsement, nil
}
