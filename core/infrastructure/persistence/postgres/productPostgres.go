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

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) repository.IProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById(id int) (*domain.Product, error) {
	var product domain.Product

	if err := r.database.Model(&domain.Product{}).Find(&product, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if product.ID == 0 {
		return nil, persistence.ProductNotFound
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll(params url.Values) (*[]domain.Product, map[string]interface{}, error) {
	var products []domain.Product
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("products").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.Product{}).Scopes(persistence.Paginate(params, paginationResult)).Find(&products).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &products, paginationResult, nil
}

func (r *ProductRepositoryImpl) Create(productDto *domain.Product, userSessionId int) (*domain.Product, error) {

	if err := r.database.Model(&domain.Product{}).Create(&productDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&productDto)
	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return productDto, nil
}

func (r *ProductRepositoryImpl) Update(id int, productDto map[string]interface{}, userSessionId int) (*domain.Product, error) {
	product, productOld := domain.Product{}, domain.Product{}

	r.database.Model(&domain.Product{}).Find(&productOld, id)

	if productOld.ID == 0 {
		return nil, persistence.ProductNotFound
	}

	if err := r.database.Model(&domain.Product{}).Where("id = ?", id).Updates(&productDto).Find(&product).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	old, _ := json.Marshal(&productOld)
	current, _ := json.Marshal(&product)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpUpdate, string(old[:]), string(current[:]))
	return &product, nil
}
