package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/product/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) domain.ProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById() (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *ProductRepositoryImpl) GetAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *ProductRepositoryImpl) Create() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *ProductRepositoryImpl) Update() (domain.Product, error) {
	return domain.Product{}, nil
}
