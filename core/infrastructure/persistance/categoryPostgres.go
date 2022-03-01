package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) repository.ICategoryRepository {
	return &CategoryRepositoryImpl{database: db}
}

func (r *CategoryRepositoryImpl) GetById(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	if err := r.database.Model(&domain.Category{}).Preload("Products").Find(&category, c.Param("id")).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if category.ID == 0 {
		return nil, errors.New("category not found")
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll() (*[]domain.Category, error) {
	var categories []domain.Category

	if err := r.database.Model(&domain.Category{}).Preload("Products").Find(&categories).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}
	return &categories, nil
}

func (r *CategoryRepositoryImpl) Create(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.Category{}).Create(&category).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) Delete(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	r.database.Model(&domain.Category{}).Find(&category, c.Param("id"))
	if category.ID == 0 {
		return nil, errors.New("category not found")
	} else if len(category.Products) > 0 {
		return nil, errors.New("you canot delete a category with related data")
	}

	if err := r.database.Model(&domain.Category{}).Unscoped().Delete(&category, &category.ID).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(c *gin.Context) (*domain.Category, error) {
	patch, category := map[string]interface{}{}, domain.Category{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New("error binding JSON data, verify json format")
	}

	if err := r.database.Model(&domain.Category{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&category).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if category.ID == 0 {
		return nil, errors.New("category not found")
	}

	return &category, nil
}
