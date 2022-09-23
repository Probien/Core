package postgres

import (
	"encoding/json"
	"math"
	"net/url"
	"time"

	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"gorm.io/gorm"
)

type PawnOrderRepositoryImpl struct {
	database *gorm.DB
}

func NewPawnOrderRepositoryImpl(db *gorm.DB) repository.IPawnOrderRepository {
	return &PawnOrderRepositoryImpl{database: db}
}

func (r *PawnOrderRepositoryImpl) GetById(id int) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Preload("Employee").Preload("Customer").Preload("Status").Preload("Endorsements").Preload("Products").Find(&pawnOrder, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if pawnOrder.ID == 0 {
		return nil, persistence.PawnOrderNotFound
	}

	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetByIdForUpdate(id uint) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Find(&pawnOrder, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if pawnOrder.ID == 0 {
		return nil, persistence.PawnOrderNotFound
	}

	pawnOrder.CutOffDay = pawnOrder.CutOffDay.AddDate(0, 0, 7)
	if pawnOrder.Monthly {
		pawnOrder.ExtensionDate = pawnOrder.CutOffDay.AddDate(0, 0, 3)
	} else {
		pawnOrder.ExtensionDate = pawnOrder.CutOffDay.AddDate(0, 0, 1)
	}

	if err := r.database.Model(&pawnOrder).Where("id = ?", pawnOrder.ID).Updates(map[string]interface{}{"cut_off_day": pawnOrder.CutOffDay, "extension_date": pawnOrder.ExtensionDate, "status_id": 1}).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetAll(params url.Values) (*[]domain.PawnOrder, map[string]interface{}, error) {
	var pawnOrders []domain.PawnOrder
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("pawn_orders").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.PawnOrder{}).Scopes(persistence.Paginate(params, paginationResult)).Preload("Customer").Preload("Employee").Preload("Status").Find(&pawnOrders).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &pawnOrders, paginationResult, nil
}

func (r *PawnOrderRepositoryImpl) Create(pawnOrderDto *domain.PawnOrder, userSessionId int) (*domain.PawnOrder, error) {
	pawnOrderDto.CutOffDay = time.Now().AddDate(0, 0, 7)

	if pawnOrderDto.Monthly {
		pawnOrderDto.ExtensionDate = time.Now().AddDate(0, 0, 10)
	} else {
		pawnOrderDto.ExtensionDate = time.Now().AddDate(0, 0, 8)
	}

	if err := r.database.Model(&domain.PawnOrder{}).Omit("Employee").Omit("Customer").Omit("Status").Create(&pawnOrderDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&pawnOrderDto)
	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return pawnOrderDto, nil
}

func (r *PawnOrderRepositoryImpl) Update(id int, pawnOrderDto map[string]interface{}, userSessionId int) (*domain.PawnOrder, error) {
	pawnOrder, pawnOrderOld := domain.PawnOrder{}, domain.PawnOrder{}

	r.database.Model(&domain.PawnOrder{}).Find(&pawnOrderOld, id)

	if pawnOrderOld.ID == 0 {
		return nil, persistence.PawnOrderNotFound
	}

	if err := r.database.Model(&domain.PawnOrder{}).Where("id = ?", id).Omit("Products").Omit("Endorsements").Updates(&pawnOrderDto).Find(&pawnOrder).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	old, _ := json.Marshal(&pawnOrderOld)
	current, _ := json.Marshal(&pawnOrder)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpUpdate, string(old[:]), string(current[:]))
	return &pawnOrder, nil
}
