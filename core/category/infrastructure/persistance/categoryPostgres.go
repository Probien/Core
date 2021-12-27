package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/category/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) domain.CategoryRepository {
	return &CategoryRepositoryImpl{database: db}
}

func (r *CategoryRepositoryImpl) GetById() (domain.Category, error) {
	return domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) GetAll() ([]domain.Category, error) {
	return []domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) Create() (domain.Category, error) {
	return domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) Delete() (domain.Category, error) {
	return domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) Update() (domain.Category, error) {
	return domain.Category{}, nil
}
