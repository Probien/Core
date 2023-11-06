package adapter

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"math"
	"net/url"

	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/persistence"

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
		return nil, persistence.ErrorProcess
	}

	if category.ID == 0 {
		return nil, persistence.CategoryNotFound
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll(params url.Values) (*[]dto.Category, map[string]interface{}, error) {
	var categories []dto.Category
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("categories").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.Category{}).Scopes(persistence.Paginate(params, paginationResult)).Find(&categories).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &categories, paginationResult, nil
}

func (r *CategoryRepositoryImpl) Create(categoryDto *dto.Category, userSessionId int) (*dto.Category, error) {

	if err := r.database.Model(&dto.Category{}).Create(&categoryDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&categoryDto)
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))

	return categoryDto, nil
}

func (r *CategoryRepositoryImpl) Delete(id int, userSessionId int) (*dto.Category, error) {
	var category dto.Category

	r.database.Model(&dto.Category{}).Find(&category, id)

	if category.ID == 0 {
		return nil, persistence.CategoryNotFound
	} else if len(category.Products) > 0 {
		return nil, persistence.InvalidAction
	}

	if err := r.database.Model(&dto.Category{}).Delete(&category, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	deleted, _ := json.Marshal(&category)
	r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpDelete, string(deleted[:]), persistence.SpNoCurrData)
	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(id int, categoryDto map[string]interface{}, userSessionId int) (*dto.Category, error) {
	category, categoryOld := dto.Category{}, dto.Category{}

	r.database.Model(&dto.Category{}).Find(&categoryOld, id)

	if categoryOld.ID == 0 {
		return nil, persistence.CategoryNotFound
	}

	if err := r.database.Model(&dto.Category{}).Where("id = ?", id).Updates(&categoryDto).Find(&category).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	old, _ := json.Marshal(&categoryOld)
	current, _ := json.Marshal(&category)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpUpdate, string(old[:]), string(current[:]))
	return &category, nil
}
