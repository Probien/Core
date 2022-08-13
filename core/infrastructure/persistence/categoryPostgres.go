package persistence

import (
	"encoding/json"

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
		return nil, ErrorProcess
	}

	if category.ID == 0 {
		return nil, CategoryNotFound
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll() (*[]domain.Category, error) {
	var categories []domain.Category

	if err := r.database.Model(&domain.Category{}).Find(&categories).Error; err != nil {
		return nil, ErrorProcess
	}
	return &categories, nil
}

func (r *CategoryRepositoryImpl) Create(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		return nil, ErrorBinding
	}

	if err := r.database.Model(&domain.Category{}).Create(&category).Error; err != nil {
		return nil, ErrorProcess
	}

	data, _ := json.Marshal(&category)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", contextUserID.(int), SpInsert, SpNoPrevData, string(data[:]))
	return &category, nil
}

func (r *CategoryRepositoryImpl) Delete(c *gin.Context) (*domain.Category, error) {
	var category domain.Category

	r.database.Model(&domain.Category{}).Find(&category, c.Param("id"))
	if category.ID == 0 {
		return nil, CategoryNotFound
	} else if len(category.Products) > 0 {
		return nil, InvalidAction
	}

	if err := r.database.Model(&domain.Category{}).Delete(&category, &category.ID).Error; err != nil {
		return nil, ErrorProcess
	}

	deleted, _ := json.Marshal(&category)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpDelete, string(deleted[:]), SpNoCurrData)
	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(c *gin.Context) (*domain.Category, error) {
	patch, category, categoryOld := map[string]interface{}{}, domain.Category{}, domain.Category{}

	if err := c.Bind(&patch); err != nil {
		return nil, err
	}

	_, errID := patch["id"]

	if !errID {
		return nil, ErrorBinding
	}

	r.database.Model(&domain.Category{}).Find(&categoryOld, patch["id"])

	if err := r.database.Model(&domain.Category{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&category).Error; err != nil {
		return nil, ErrorProcess
	}

	if category.ID == 0 {
		return nil, CategoryNotFound
	}

	old, _ := json.Marshal(&categoryOld)
	current, _ := json.Marshal(&category)

	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpUpdate, string(old[:]), string(current[:]))
	return &category, nil
}
