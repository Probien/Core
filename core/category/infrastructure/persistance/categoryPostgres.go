package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/category/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) domain.CategoryRepository {
	return &CategoryRepositoryImpl{database: db}
}

func (r *CategoryRepositoryImpl) GetById(c *gin.Context) (domain.Category, error) {
	return domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) GetAll() ([]domain.Category, error) {
	return []domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) Create(c *gin.Context) (domain.Category, error) {
	return domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) Delete(c *gin.Context) (domain.Category, error) {
	return domain.Category{}, nil
}
func (r *CategoryRepositoryImpl) Update(c *gin.Context) (domain.Category, error) {
	return domain.Category{}, nil
}
