package postgres

import (
	"encoding/json"
	"math"
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) repository.ICategoryRepository {
	return &CategoryRepositoryImpl{database: db}
}

func (r *CategoryRepositoryImpl) GetById(id int) (*domain.Category, error) {
	var category domain.Category

	if err := r.database.Model(&domain.Category{}).Preload("Products").Find(&category, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if category.ID == 0 {
		return nil, persistence.CategoryNotFound
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll(params url.Values) (*[]domain.Category, map[string]interface{}, error) {
	var categories []domain.Category
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("categories").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.Category{}).Scopes(persistence.Paginate(params, paginationResult)).Find(&categories).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &categories, paginationResult, nil
}

func (r *CategoryRepositoryImpl) Create(categoryDto *domain.Category, userSessionId int) (*domain.Category, error) {

	if err := r.database.Model(&domain.Category{}).Create(&categoryDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&categoryDto)
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))

	return categoryDto, nil
}

func (r *CategoryRepositoryImpl) Delete(id int, userSessionId int) (*domain.Category, error) {
	var category domain.Category

	r.database.Model(&domain.Category{}).Find(&category, id)

	if category.ID == 0 {
		return nil, persistence.CategoryNotFound
	} else if len(category.Products) > 0 {
		return nil, persistence.InvalidAction
	}

	if err := r.database.Model(&domain.Category{}).Delete(&category, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	deleted, _ := json.Marshal(&category)
	r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpDelete, string(deleted[:]), persistence.SpNoCurrData)
	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(id int, categoryDto map[string]interface{}, userSessionId int) (*domain.Category, error) {
	category, categoryOld := domain.Category{}, domain.Category{}

	r.database.Model(&domain.Category{}).Find(&categoryOld, id)

	if categoryOld.ID == 0 {
		return nil, persistence.CategoryNotFound
	}

	if err := r.database.Model(&domain.Category{}).Where("id = ?", id).Updates(&categoryDto).Find(&category).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	old, _ := json.Marshal(&categoryOld)
	current, _ := json.Marshal(&category)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpUpdate, string(old[:]), string(current[:]))
	return &category, nil
}
