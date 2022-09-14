package postgres

import (
	"encoding/json"
	"math"

	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

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
		return nil, persistence.ErrorProcess
	}

	if category.ID == 0 {
		return nil, persistence.CategoryNotFound
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll(c *gin.Context) (*[]domain.Category, map[string]interface{}, error) {
	var categories []domain.Category
	var totalRows int64
	paginationResult := map[string]interface{}{}

	go r.database.Table("categories").Count(&totalRows)
	if err := r.database.Model(&domain.Category{}).Scopes(persistence.Paginate(c, paginationResult)).Find(&categories).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	paginationResult["total_pages"] = math.Ceil(float64(totalRows / 10))
	return &categories, paginationResult, nil
}

func (r *CategoryRepositoryImpl) Create(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		return nil, persistence.ErrorBinding
	}

	if err := r.database.Model(&domain.Category{}).Create(&category).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&category)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", contextUserID.(int), persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return &category, nil
}

func (r *CategoryRepositoryImpl) Delete(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	r.database.Model(&domain.Category{}).Find(&category, c.Param("id"))
	if category.ID == 0 {
		return nil, persistence.CategoryNotFound
	} else if len(category.Products) > 0 {
		return nil, persistence.InvalidAction
	}

	if err := r.database.Model(&domain.Category{}).Delete(&category, &category.ID).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	deleted, _ := json.Marshal(&category)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpDelete, string(deleted[:]), persistence.SpNoCurrData)
	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(c *gin.Context) (*domain.Category, error) {
	patch, category, categoryOld := map[string]interface{}{}, domain.Category{}, domain.Category{}

	if err := c.Bind(&patch); err != nil {
		return nil, err
	}

	_, errID := patch["id"]

	if !errID {
		return nil, persistence.ErrorBinding
	}

	r.database.Model(&domain.Category{}).Find(&categoryOld, patch["id"])

	if err := r.database.Model(&domain.Category{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&category).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if category.ID == 0 {
		return nil, persistence.CategoryNotFound
	}

	old, _ := json.Marshal(&categoryOld)
	current, _ := json.Marshal(&category)

	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpUpdate, string(old[:]), string(current[:]))
	return &category, nil
}
