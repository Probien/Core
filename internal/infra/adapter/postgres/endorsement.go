package adapter

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	port "github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/adapter"
	"math"
	"net/url"

	"gorm.io/gorm"
)

type EndorsementRepositoryImpl struct {
	database            *gorm.DB
	pawnOrderRepository PawnOrderRepositoryImpl
}

func NewEndorsementRepositoryImpl(db *gorm.DB) port.IEndorsementRepository {
	return &EndorsementRepositoryImpl{
		database:            db,
		pawnOrderRepository: PawnOrderRepositoryImpl{database: db},
	}
}

func (r *EndorsementRepositoryImpl) GetById(id int) (*dto.Endorsement, error) {
	var endorsement dto.Endorsement

	if err := r.database.Model(&dto.Endorsement{}).Find(&endorsement, id).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	if endorsement.ID == 0 {
		return nil, adapter.EndorsementNotFound
	}

	return &endorsement, nil
}

func (r *EndorsementRepositoryImpl) GetAll(params url.Values) (*[]dto.Endorsement, map[string]interface{}, error) {
	var endorsements []dto.Endorsement
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("endorsements").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.Endorsement{}).Scopes(adapter.Paginate(params, paginationResult)).Find(&endorsements).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &endorsements, paginationResult, nil
}

func (r *EndorsementRepositoryImpl) Create(endorsementDto *dto.Endorsement, userSessionId int) (*dto.Endorsement, error) {

	if _, err := r.pawnOrderRepository.GetByIdForUpdate(endorsementDto.PawnOrderID); err != nil {
		return nil, err
	}

	if err := r.database.Model(&dto.Endorsement{}).Create(&endorsementDto).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	data, _ := json.Marshal(&endorsementDto)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpInsert, adapter.SpNoPrevData, string(data[:]))
	return endorsementDto, nil
}
