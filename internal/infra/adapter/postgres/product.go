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

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) port.IProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById(id int) (*dto.Product, error) {
	var product dto.Product

	if err := r.database.Model(&dto.Product{}).Find(&product, id).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	if product.ID == 0 {
		return nil, adapter.ProductNotFound
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll(params url.Values) (*[]dto.Product, map[string]interface{}, error) {
	var products []dto.Product
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("products").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.Product{}).Scopes(adapter.Paginate(params, paginationResult)).Find(&products).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &products, paginationResult, nil
}

func (r *ProductRepositoryImpl) Create(productDto *dto.Product, userSessionId int) (*dto.Product, error) {

	if err := r.database.Model(&dto.Product{}).Create(&productDto).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	data, _ := json.Marshal(&productDto)
	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpInsert, adapter.SpNoPrevData, string(data[:]))
	return productDto, nil
}

func (r *ProductRepositoryImpl) Update(id int, productDto map[string]interface{}, userSessionId int) (*dto.Product, error) {
	product, productOld := dto.Product{}, dto.Product{}

	r.database.Model(&dto.Product{}).Find(&productOld, id)

	if productOld.ID == 0 {
		return nil, adapter.ProductNotFound
	}

	if err := r.database.Model(&dto.Product{}).Where("id = ?", id).Updates(&productDto).Find(&product).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	old, _ := json.Marshal(&productOld)
	current, _ := json.Marshal(&product)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpUpdate, string(old[:]), string(current[:]))
	return &product, nil
}
