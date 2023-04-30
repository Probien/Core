package postgres

import (
	"encoding/json"
	"math"
	"net/url"

	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/persistence"

	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/domain/repository"
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

func (r *EndorsementRepositoryImpl) GetById(id int) (*domain.Endorsement, error) {
	var endorsement domain.Endorsement

	if err := r.database.Model(&domain.Endorsement{}).Find(&endorsement, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if endorsement.ID == 0 {
		return nil, persistence.EndorsementNotFound
	}

	return &endorsement, nil
}

func (r *EndorsementRepositoryImpl) GetAll(params url.Values) (*[]domain.Endorsement, map[string]interface{}, error) {
	var endorsements []domain.Endorsement
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("endorsements").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.Endorsement{}).Scopes(persistence.Paginate(params, paginationResult)).Find(&endorsements).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &endorsements, paginationResult, nil
}

func (r *EndorsementRepositoryImpl) Create(endorsementDto *domain.Endorsement, userSessionId int) (*domain.Endorsement, error) {

	if _, err := r.pawnOrderRepository.GetByIdForUpdate(endorsementDto.PawnOrderID); err != nil {
		return nil, err
	}

	if err := r.database.Model(&domain.Endorsement{}).Create(&endorsementDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&endorsementDto)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return endorsementDto, nil
}
