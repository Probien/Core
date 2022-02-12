package persistance

import (
	"errors"

	product_domain "github.com/JairDavid/Probien-Backend/core/domain/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) product_domain.ProductRepository {
	return &ProductRepositoryImpl{database: db}
}

func (r *ProductRepositoryImpl) GetById(c *gin.Context) (*product_domain.Product, error) {
	var product product_domain.Product

	r.database.Model(&product_domain.Product{}).Find(&product, c.Param("id"))
	if product.ID == 0 {
		return nil, errors.New("product order not found")
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll() (*[]product_domain.Product, error) {
	var products []product_domain.Product

	r.database.Model(&product_domain.Product{}).Find(&products)
	return &products, nil
}

func (r *ProductRepositoryImpl) Create(c *gin.Context) (*product_domain.Product, error) {
	var product product_domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}
	r.database.Model(&product_domain.Product{}).Create(&product)
	return &product, nil
}

func (r *ProductRepositoryImpl) Update(c *gin.Context) (*product_domain.Product, error) {

	patch, product := map[string]interface{}{}, product_domain.Product{}

	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["id"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an ID in the JSON body")
	}

	result := r.database.Model(&product_domain.Product{}).Where("id = ?", patch["id"]).Omit("id").Updates(&patch).Find(&product)
	if result.RowsAffected == 0 {
		return nil, errors.New("product not found or json data does not match ")
	}

	return &product, nil
}
