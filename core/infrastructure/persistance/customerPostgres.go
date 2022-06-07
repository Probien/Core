package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	database *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) repository.ICustomerRepository {
	return &CustomerRepositoryImpl{database: db}
}

func (r *CustomerRepositoryImpl) GetById(c *gin.Context) (*domain.Customer, error) {
	var customer domain.Customer

	if err := r.database.Model(&domain.Customer{}).Preload("PawnOrders.Products").Preload("PawnOrders.Endorsements").Find(&customer, c.Param("id")).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if customer.ID == 0 {
		return nil, errors.New("customer not found")
	}
	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAll() (*[]domain.Customer, error) {
	var customers []domain.Customer

	if err := r.database.Model(domain.Customer{}).Preload("PawnOrders").Find(&customers).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &customers, nil
}

func (r *CustomerRepositoryImpl) Create(c *gin.Context) (*domain.Customer, error) {
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.Customer{}).Create(&customer).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(c *gin.Context) (*domain.Customer, error) {
	patch, customer := map[string]interface{}{}, domain.Customer{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New("error binding JSON data, verify json format")
	}

	if err := r.database.Model(&domain.Customer{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&customer).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if customer.ID == 0 {
		return nil, errors.New("customer not found")
	}

	return &customer, nil
}
