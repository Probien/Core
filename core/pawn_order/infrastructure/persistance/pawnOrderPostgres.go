package persistance

import (
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

func (r *PawnOrderRepositoryImpl) GetById(c *gin.Context) (domain.PawnOrder, error) {
	return domain.PawnOrder{}, nil
}

func (r *PawnOrderRepositoryImpl) GetAll() ([]domain.PawnOrder, error) {
	return []domain.PawnOrder{}, nil
}

func (r *PawnOrderRepositoryImpl) Create(c *gin.Context) (domain.PawnOrder, error) {
	return domain.PawnOrder{}, nil
}

func (r *PawnOrderRepositoryImpl) Update(c *gin.Context) (domain.PawnOrder, error) {
	return domain.PawnOrder{}, nil
}
