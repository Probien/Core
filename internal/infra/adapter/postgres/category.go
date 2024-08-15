package adapter

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/adapter"
	"math"
	"net/url"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) port.ICategoryRepository {
	return &CategoryRepositoryImpl{database: db}
}

func (r *CategoryRepositoryImpl) GetById(id int) (*dto.Category, error) {
	var category dto.Category

	if err := r.database.Model(&dto.Category{}).Preload("Products").Find(&category, id).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	if category.ID == 0 {
		return nil, adapter.CategoryNotFound
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll(params url.Values) (*[]dto.Category, map[string]interface{}, error) {
	var categories []dto.Category
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("categories").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.Category{}).Scopes(adapter.Paginate(params, paginationResult)).Find(&categories).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &categories, paginationResult, nil
}

func (r *CategoryRepositoryImpl) Create(categoryDto *dto.Category, userSessionId int) (*dto.Category, error) {

	if err := r.database.Model(&dto.Category{}).Create(&categoryDto).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	data, _ := json.Marshal(&categoryDto)
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", userSessionId, adapter.SpInsert, adapter.SpNoPrevData, string(data[:]))

	return categoryDto, nil
}

func (r *CategoryRepositoryImpl) Delete(id int, userSessionId int) (*dto.Category, error) {
	var category dto.Category

	r.database.Model(&dto.Category{}).Find(&category, id)

	if category.ID == 0 {
		return nil, adapter.CategoryNotFound
	} else if len(category.Products) > 0 {
		return nil, adapter.InvalidAction
	}

	if err := r.database.Model(&dto.Category{}).Delete(&category, id).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	deleted, _ := json.Marshal(&category)
	r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpDelete, string(deleted[:]), adapter.SpNoCurrData)
	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(id int, categoryDto map[string]interface{}, userSessionId int) (*dto.Category, error) {
	category, categoryOld := dto.Category{}, dto.Category{}

	r.database.Model(&dto.Category{}).Find(&categoryOld, id)

	if categoryOld.ID == 0 {
		return nil, adapter.CategoryNotFound
	}

	if err := r.database.Model(&dto.Category{}).Where("id = ?", id).Updates(&categoryDto).Find(&category).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	old, _ := json.Marshal(&categoryOld)
	current, _ := json.Marshal(&category)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpUpdate, string(old[:]), string(current[:]))
	return &category, nil
}
