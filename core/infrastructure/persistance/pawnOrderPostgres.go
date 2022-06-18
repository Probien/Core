package persistance

import (
	"encoding/json"
	"errors"
	"log"

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
		return nil, errors.New(ERROR_PROCCESS)
	}

	if pawnOrder.ID == 0 {
		return nil, errors.New(PAWNORDER_NOT_FOUND)
	}

	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetAll() (*[]domain.PawnOrder, error) {
	var pawnOrders []domain.PawnOrder

	if err := r.database.Model(&domain.PawnOrder{}).Preload("Customer").Preload("Employee").Preload("Status").Find(&pawnOrders).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	return &pawnOrders, nil
}

func (r *PawnOrderRepositoryImpl) Create(c *gin.Context) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := c.ShouldBindJSON(&pawnOrder); err != nil || pawnOrder.CustomerID == 0 {
		log.Print(err)
		return nil, errors.New(ERROR_BINDING)
	}

	if err := r.database.Model(&domain.PawnOrder{}).Omit("Employee").Omit("Customer").Omit("Status").Create(&pawnOrder).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	data, _ := json.Marshal(&pawnOrder)
	//replace number 1 for employeeID session (JWT fix)
	go r.database.Exec("CALL savemovement(?,?,?,?)", 2, SP_INSERT, SP_NO_PREV_DATA, string(data[:]))
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) Update(c *gin.Context) (*domain.PawnOrder, error) {
	patch, pawnOrder, pawnOrderOld := map[string]interface{}{}, domain.PawnOrder{}, domain.PawnOrder{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New(ERROR_BINDING)
	}

	r.database.Model(&domain.PawnOrder{}).Find(&pawnOrderOld, patch["id"])

	if err := r.database.Model(&domain.PawnOrder{}).Where("id = ?", patch["id"]).Omit("Products").Omit("Endorsements").Updates(&patch).Find(&pawnOrder).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if pawnOrder.ID == 0 {
		return nil, errors.New(ERROR_BINDING)
	}

	old, _ := json.Marshal(&pawnOrderOld)
	new, _ := json.Marshal(&pawnOrder)
	//replace number 1 for employeeID session (JWT fix)
	go r.database.Exec("CALL savemovement(?,?,?,?)", 2, SP_UPDATE, string(old[:]), string(new[:]))
	return &pawnOrder, nil
}
