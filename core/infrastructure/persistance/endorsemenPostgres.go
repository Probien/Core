package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EndorsementRepositoryImpl struct {
	database *gorm.DB
}

func NewEndorsementRepositoryImpl(db *gorm.DB) repository.IEndorsementRepository {
	return &EndorsementRepositoryImpl{database: db}
}

func (r *EndorsementRepositoryImpl) GetById(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := r.database.Model(&domain.Endorsement{}).Find(&endorsement, c.Param("id")).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &endorsement, nil
}

func (r *EndorsementRepositoryImpl) GetAll() (*[]domain.Endorsement, error) {
	var endorsements []domain.Endorsement

	if err := r.database.Model(&domain.Endorsement{}).Find(&endorsements).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &endorsements, nil
}

func (r *EndorsementRepositoryImpl) Create(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := c.ShouldBindJSON(&endorsement); err != nil || endorsement.PawnOrderID == 0 {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.Endorsement{}).Create(&endorsement).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &endorsement, nil
}
