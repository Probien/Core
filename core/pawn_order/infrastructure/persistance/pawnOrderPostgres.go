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
		return &domain.PawnOrder{}, errors.New("pawn order not found")
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
		return &domain.PawnOrder{}, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.PawnOrder{}).Create(&pawnOrder)
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) Update(c *gin.Context) (*domain.PawnOrder, error) {
	return &domain.PawnOrder{}, nil
}
