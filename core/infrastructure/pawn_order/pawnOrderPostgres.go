package persistance

import (
	"errors"

	pawn_order_domain "github.com/JairDavid/Probien-Backend/core/domain/pawn_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PawnOrderRepositoryImpl struct {
	database *gorm.DB
}

func NewPawnOrderRepositoryImpl(db *gorm.DB) pawn_order_domain.PawnOrderRepository {
	return &PawnOrderRepositoryImpl{database: db}
}

func (r *PawnOrderRepositoryImpl) GetById(c *gin.Context) (*pawn_order_domain.PawnOrder, error) {
	var pawnOrder pawn_order_domain.PawnOrder

	r.database.Model(&pawn_order_domain.PawnOrder{}).Find(&pawnOrder, c.Param("id"))
	if pawnOrder.ID == 0 {
		return nil, errors.New("pawn order not found")
	}
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) GetAll() (*[]pawn_order_domain.PawnOrder, error) {
	var pawnOrders []pawn_order_domain.PawnOrder

	r.database.Model(&pawn_order_domain.PawnOrder{}).Find(&pawnOrders)
	return &pawnOrders, nil
}

func (r *PawnOrderRepositoryImpl) Create(c *gin.Context) (*pawn_order_domain.PawnOrder, error) {
	var pawnOrder pawn_order_domain.PawnOrder

	if err := c.ShouldBindJSON(&pawnOrder); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&pawn_order_domain.PawnOrder{}).Omit("Endorsements").Create(&pawnOrder)
	return &pawnOrder, nil
}

func (r *PawnOrderRepositoryImpl) Update(c *gin.Context) (*pawn_order_domain.PawnOrder, error) {
	patch, pawnOrder := map[string]interface{}{}, pawn_order_domain.PawnOrder{}

	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["id"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an ID in the JSON body")
	}

	result := r.database.Model(&pawn_order_domain.PawnOrder{}).Where("id = ?", patch["id"]).Omit("id").Updates(&patch).Find(&pawnOrder)
	if result.RowsAffected == 0 {
		return nil, errors.New("pawn order not found or json data does not match ")
	}

	return &pawnOrder, nil
}
