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

func NewCategoryRepositoryImpl(db *gorm.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{database: db}
}

func (r *CategoryRepositoryImpl) GetById(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	r.database.Model(&domain.Category{}).Preload("Products").Find(&category, c.Param("id"))
	if category.ID == 0 {
		return nil, errors.New("category not found")
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll() (*[]domain.Category, error) {
	var categories []domain.Category
	r.database.Raw("SELECT * FROM categories").Scan(&categories)
	//r.database.Model(&category_domain.Category{}).Preload("Products").Find(&categories)
	return &categories, nil
}

func (r *CategoryRepositoryImpl) Create(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.Category{}).Create(&category)
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

	r.database.Model(&domain.Category{}).Unscoped().Delete(&category, &category.ID)

	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(c *gin.Context) (*domain.Category, error) {
	patch, category := map[string]interface{}{}, domain.Category{}

	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["id"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an ID in the JSON body")
	}

	result := r.database.Model(&domain.Category{}).Where("id = ?", patch["id"]).Omit("id").Updates(&patch).Find(&category)
	if result.RowsAffected == 0 {
		return nil, errors.New("category not found or json data does not match ")
	}

	return &category, nil
}
