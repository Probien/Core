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

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) repository.IProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	if err := r.database.Model(&domain.Product{}).Find(&product, c.Param("id")).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if product.ID == 0 {
		return nil, persistence.ProductNotFound
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll(c *gin.Context) (*[]domain.Product, map[string]interface{}, error) {
	var products []domain.Product
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("products").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.Product{}).Scopes(persistence.Paginate(c, paginationResult)).Find(&products).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &products, paginationResult, nil
}

func (r *ProductRepositoryImpl) Create(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return nil, persistence.ErrorBinding
	}

	if err := r.database.Model(&domain.Product{}).Create(&product).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&product)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return &product, nil
}

func (r *ProductRepositoryImpl) Update(c *gin.Context) (*domain.Product, error) {
	patch, product, productOld := map[string]interface{}{}, domain.Product{}, domain.Product{}

	if err := c.Bind(&patch); err != nil {
		return nil, persistence.ErrorBinding
	}

	_, errID := patch["id"]

	if !errID {
		return nil, persistence.ErrorBinding
	}

	r.database.Model(&domain.Product{}).Find(&productOld, patch["id"])

	if err := r.database.Model(&domain.Product{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&product).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if product.ID == 0 {
		return nil, persistence.ProductNotFound
	}

	old, _ := json.Marshal(&productOld)
	current, _ := json.Marshal(&product)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpUpdate, string(old[:]), string(current[:]))
	return &product, nil
}
