package persistance

import (
	"errors"

	endorsement_domain "github.com/JairDavid/Probien-Backend/core/domain/endorsement"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EndorsementRepositoryImpl struct {
	database *gorm.DB
}

func NewEndorsementRepositoryImpl(db *gorm.DB) endorsement_domain.EndorsementRepository {
	return &EndorsementRepositoryImpl{database: db}
}

func (r *EndorsementRepositoryImpl) GetById(c *gin.Context) (*endorsement_domain.Endorsement, error) {
	var endorsement endorsement_domain.Endorsement

	r.database.Model(&endorsement_domain.Endorsement{}).Preload("PawnOrderID").Find(&endorsement, c.Param("id"))
	if endorsement.ID == 0 {
		return nil, errors.New("endorsement not found")
	}
	return &endorsement, nil
}

func (r *EndorsementRepositoryImpl) GetAll() (*[]endorsement_domain.Endorsement, error) {
	var endorsements []endorsement_domain.Endorsement

	r.database.Model(&endorsement_domain.Endorsement{}).Find(&endorsements)
	return &endorsements, nil
}

func (r *EndorsementRepositoryImpl) Create(c *gin.Context) (*endorsement_domain.Endorsement, error) {
	var endorsement endorsement_domain.Endorsement

	if err := c.ShouldBindJSON(&endorsement); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}
	r.database.Model(&endorsement_domain.Endorsement{}).Create(&endorsement)

	return &endorsement, nil
}
