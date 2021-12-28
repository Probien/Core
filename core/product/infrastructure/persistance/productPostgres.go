package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/product/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) domain.ProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById(c *gin.Context) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *ProductRepositoryImpl) GetAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *ProductRepositoryImpl) Create(c *gin.Context) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *ProductRepositoryImpl) Update(c *gin.Context) (domain.Product, error) {
	return domain.Product{}, nil
}
