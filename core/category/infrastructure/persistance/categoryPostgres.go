package persistance

import (
	"errors"

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
	var category domain.Category
	result := r.database.Model(&domain.Category{}).Find(&category, c.Param("id"))
	if result.RowsAffected == 0 {
		return domain.Category{}, errors.New("category not found")
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) GetAll() ([]domain.Category, error) {
	var categories []domain.Category

	r.database.Model(&domain.Category{}).Find(&categories)
	return []domain.Category{}, nil
}

func (r *CategoryRepositoryImpl) Create(c *gin.Context) (domain.Category, error) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		return domain.Category{}, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.Category{}).Create(&category)
	return category, nil
}

func (r *CategoryRepositoryImpl) Delete(c *gin.Context) (domain.Category, error) {
	return domain.Category{}, nil
}

func (r *CategoryRepositoryImpl) Update(c *gin.Context) (domain.Category, error) {
	return domain.Category{}, nil
}
