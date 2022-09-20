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

	if err := r.database.Model(&domain.PawnOrder{}).Preload("Employee").Preload("Customer").Preload("Status").Preload("Endorsements").Preload("Products").Find(&pawnOrder, c.Param("id")).Error; err != nil {
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

	if err := r.database.Model(&domain.PawnOrder{}).Scopes(persistence.Paginate(c, paginationResult)).Preload("Customer").Preload("Employee").Preload("Status").Find(&pawnOrders).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &pawnOrders, paginationResult, nil
}

func (r *PawnOrderRepositoryImpl) Create(pawnOrderDto *domain.PawnOrder) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := c.ShouldBindJSON(&pawnOrder); err != nil || pawnOrder.CustomerID == 0 {
		return nil, persistence.ErrorBinding
	}
	pawnOrder.CutOffDay = time.Now().AddDate(0, 0, 7)

	if pawnOrder.Monthly {
		pawnOrder.ExtensionDate = time.Now().AddDate(0, 0, 10)
	} else {
		pawnOrder.ExtensionDate = time.Now().AddDate(0, 0, 8)
	}

	if err := r.database.Model(&domain.PawnOrder{}).Omit("Employee").Omit("Customer").Omit("Status").Create(&pawnOrder).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&pawnOrder)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) Update(pawnOrderDto map[string]interface{}) (*domain.PawnOrder, error) {
	patch, pawnOrder, pawnOrderOld := map[string]interface{}{}, domain.PawnOrder{}, domain.PawnOrder{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, persistence.ErrorBinding
	}

	r.database.Model(&domain.PawnOrder{}).Find(&pawnOrderOld, patch["id"])

	if err := r.database.Model(&domain.PawnOrder{}).Where("id = ?", patch["id"]).Omit("Products").Omit("Endorsements").Updates(&patch).Find(&pawnOrder).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if pawnOrder.ID == 0 {
		return nil, persistence.PawnOrderNotFound
	}

	old, _ := json.Marshal(&pawnOrderOld)
	current, _ := json.Marshal(&pawnOrder)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpUpdate, string(old[:]), string(current[:]))
	return &pawnOrder, nil
}
