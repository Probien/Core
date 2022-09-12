package postgres

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EndorsementRepositoryImpl struct {
	database            *gorm.DB
	pawnOrderRepository PawnOrderRepositoryImpl
}

func NewEndorsementRepositoryImpl(db *gorm.DB) repository.IEndorsementRepository {
	return &EndorsementRepositoryImpl{
		database:            db,
		pawnOrderRepository: PawnOrderRepositoryImpl{database: db},
	}
}

func (r *EndorsementRepositoryImpl) GetById(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := r.database.Model(&domain.Endorsement{}).Find(&endorsement, c.Param("id")).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if endorsement.ID == 0 {
		return nil, persistence.EndorsementNotFound
	}

	return &endorsement, nil
}

func (r *EndorsementRepositoryImpl) GetAll(c *gin.Context) (*[]domain.Endorsement, error) {
	var endorsements []domain.Endorsement

	if err := r.database.Model(&domain.Endorsement{}).Find(&endorsements).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	return &endorsements, nil
}

func (r *EndorsementRepositoryImpl) Create(c *gin.Context) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := c.ShouldBindJSON(&endorsement); err != nil || endorsement.PawnOrderID == 0 || endorsement.EmployeeID == 0 {
		return nil, persistence.ErrorBinding
	}

	if _, err := r.pawnOrderRepository.GetByIdForUpdate(endorsement.PawnOrderID); err != nil {
		return nil, err
	}

	if err := r.database.Model(&domain.Endorsement{}).Create(&endorsement).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&endorsement)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return &endorsement, nil
}
