package persistence

import (
	"encoding/json"
	"time"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PawnOrderRepositoryImpl struct {
	database *gorm.DB
}

func NewPawnOrderRepositoryImpl(db *gorm.DB) repository.IPawnOrderRepository {
	return &PawnOrderRepositoryImpl{database: db}
}

func (r *PawnOrderRepositoryImpl) GetById(c *gin.Context) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Preload("Employee").Preload("Customer").Preload("Status").Preload("Endorsements").Preload("Products").Find(&pawnOrder, c.Param("id")).Error; err != nil {
		return nil, ErrorProcess
	}

	if pawnOrder.ID == 0 {
		return nil, PawnOrderNotFound
	}

	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetByIdForUpdate(id uint) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Find(&pawnOrder, id).Error; err != nil {
		return nil, ErrorProcess
	}

	if pawnOrder.ID == 0 {
		return nil, PawnOrderNotFound
	}

	pawnOrder.CutOffDay = pawnOrder.CutOffDay.AddDate(0, 0, 7)
	if pawnOrder.Monthly {
		pawnOrder.ExtensionDate = pawnOrder.CutOffDay.AddDate(0, 0, 3)
	} else {
		pawnOrder.ExtensionDate = pawnOrder.CutOffDay.AddDate(0, 0, 1)
	}

	if err := r.database.Model(&pawnOrder).Where("id = ?", pawnOrder.ID).Updates(map[string]interface{}{"cut_off_day": pawnOrder.CutOffDay, "extension_date": pawnOrder.ExtensionDate, "status_id": 1}).Error; err != nil {
		return nil, ErrorProcess
	}

	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetAll() (*[]domain.PawnOrder, error) {
	var pawnOrders []domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Preload("Customer").Preload("Employee").Preload("Status").Find(&pawnOrders).Error; err != nil {
		return nil, ErrorProcess
	}

	return &pawnOrders, nil
}

func (r *PawnOrderRepositoryImpl) Create(c *gin.Context) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := c.ShouldBindJSON(pawnOrder); err != nil || pawnOrder.CustomerID == 0 {
		return nil, ErrorBinding
	}
	pawnOrder.CutOffDay = time.Now().AddDate(0, 0, 7)

	if pawnOrder.Monthly {
		pawnOrder.ExtensionDate = time.Now().AddDate(0, 0, 10)
	} else {
		pawnOrder.ExtensionDate = time.Now().AddDate(0, 0, 8)
	}

	if err := r.database.Model(&domain.PawnOrder{}).Omit("Employee").Omit("Customer").Omit("Status").Create(&pawnOrder).Error; err != nil {
		return nil, ErrorProcess
	}

	data, _ := json.Marshal(&pawnOrder)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpInsert, SpNoPrevData, string(data[:]))
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) Update(c *gin.Context) (*domain.PawnOrder, error) {
	patch, pawnOrder, pawnOrderOld := map[string]interface{}{}, domain.PawnOrder{}, domain.PawnOrder{}
	_, errID := patch["id"]

	if err := c.Bind(patch); err != nil && !errID {
		return nil, ErrorBinding
	}

	r.database.Model(&domain.PawnOrder{}).Find(&pawnOrderOld, patch["id"])

	if err := r.database.Model(&domain.PawnOrder{}).Where("id = ?", patch["id"]).Omit("Products").Omit("Endorsements").Updates(&patch).Find(&pawnOrder).Error; err != nil {
		return nil, ErrorProcess
	}

	if pawnOrder.ID == 0 {
		return nil, PawnOrderNotFound
	}

	old, _ := json.Marshal(&pawnOrderOld)
	current, _ := json.Marshal(&pawnOrder)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpUpdate, string(old[:]), string(current[:]))
	return &pawnOrder, nil
}
