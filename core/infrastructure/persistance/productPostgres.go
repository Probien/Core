package persistance

import (
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
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if product.ID == 0 {
		return nil, errors.New("product order not found")
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll() (*[]domain.Product, error) {
	var products []domain.Product

	if err := r.database.Model(&domain.Product{}).Find(&products).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &products, nil
}

func (r *ProductRepositoryImpl) Create(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.Product{}).Create(&product).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) Update(c *gin.Context) (*domain.Product, error) {
	patch, product := map[string]interface{}{}, domain.Product{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New("error binding JSON data, verify json format")
	}

	if err := r.database.Model(&domain.Product{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&product).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if product.ID == 0 {
		return nil, errors.New("product not found or json data does not match ")
	}

	return &product, nil
}
