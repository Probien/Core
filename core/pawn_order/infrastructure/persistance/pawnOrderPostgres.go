package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/pawn_order/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PawnOrderRepositoryImpl struct {
	database *gorm.DB
}

func NewPawnOrderRepositoryImpl(db *gorm.DB) domain.PawnOrderRepository {
	return &PawnOrderRepositoryImpl{database: db}
}

func (r *PawnOrderRepositoryImpl) GetById(c *gin.Context) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	r.database.Model(&domain.PawnOrder{}).Preload("Products").Preload("Endorsements").Find(&pawnOrder, c.Param("id"))
	if pawnOrder.ID == 0 {
		return nil, errors.New("pawn order not found")
	}
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetAll() (*[]domain.PawnOrder, error) {
	var pawnOrders []domain.PawnOrder

	r.database.Model(&domain.PawnOrder{}).Preload("Products").Preload("Endorsements").Find(&pawnOrders)
	return &pawnOrders, nil
}

func (r *PawnOrderRepositoryImpl) Create(c *gin.Context) (*domain.PawnOrder, error) {
	var pawnOrder domain.PawnOrder

	if err := c.ShouldBindJSON(&pawnOrder); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.PawnOrder{}).Create(&pawnOrder)
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) Update(c *gin.Context) (*domain.PawnOrder, error) {
	patch, pawnOrder := map[string]interface{}{}, domain.PawnOrder{}

	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["id"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an ID in the JSON body")
	}

	result := r.database.Model(&domain.PawnOrder{}).Where("id = ?", patch["id"]).Omit("id").Updates(&patch).Find(&pawnOrder)
	if result.RowsAffected == 0 {
		return nil, errors.New("category not found or json data does not match ")
	}

	return &domain.PawnOrder{}, nil
}
