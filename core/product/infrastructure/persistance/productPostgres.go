package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/product/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) domain.ProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	r.database.Model(&domain.Product{}).Preload("Products").Preload("Endorsements").Find(&product, c.Param("id"))
	if product.ID == 0 {
		return nil, errors.New("pawn order not found")
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll() (*[]domain.Product, error) {
	var products []domain.Product

	r.database.Model(&domain.Product{}).Preload("Products").Preload("Endorsements").Find(&products)
	return &[]domain.Product{}, nil
}

func (r *ProductRepositoryImpl) Create(c *gin.Context) (*domain.Product, error) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}
	r.database.Model(&domain.Product{}).Create(&product)
	return &product, nil
}

func (r *ProductRepositoryImpl) Update(c *gin.Context) (*domain.Product, error) {

	patch, product := map[string]interface{}{}, domain.Product{}

	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["id"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an ID in the JSON body")
	}

	result := r.database.Model(&domain.Product{}).Where("id = ?", patch["id"]).Omit("id").Updates(&patch).Find(&product)
	if result.RowsAffected == 0 {
		return nil, errors.New("category not found or json data does not match ")
	}

	return &product, nil
}
