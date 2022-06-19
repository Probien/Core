package persistance

import (
	"encoding/json"
	"errors"

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
		return nil, errors.New(ERROR_PROCCESS)
	}

	if product.ID == 0 {
		return nil, errors.New(PRODUCT_NOT_FOUND)
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll() (*[]domain.Product, error) {
	var products []domain.Product

	if err := r.database.Model(&domain.Product{}).Find(&products).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	return &products, nil
}

func (r *ProductRepositoryImpl) Create(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return nil, errors.New(ERROR_BINDING)
	}

	if err := r.database.Model(&domain.Product{}).Create(&product).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	data, _ := json.Marshal(&product)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SP_INSERT, SP_NO_PREV_DATA, string(data[:]))
	return &product, nil
}

func (r *ProductRepositoryImpl) Update(c *gin.Context) (*domain.Product, error) {
	patch, product, productOld := map[string]interface{}{}, domain.Product{}, domain.Product{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New(ERROR_BINDING)
	}

	r.database.Model(&domain.Product{}).Find(&productOld, patch["id"])

	if err := r.database.Model(&domain.Product{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&product).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if product.ID == 0 {
		return nil, errors.New(ERROR_BINDING)
	}

	old, _ := json.Marshal(&productOld)
	new, _ := json.Marshal(&product)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SP_UPDATE, string(old[:]), string(new[:]))
	return &product, nil
}
