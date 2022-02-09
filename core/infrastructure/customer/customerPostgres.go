package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/customer/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	database *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) domain.CustomerRepository {
	return &CustomerRepositoryImpl{database: db}
}

func (r *CustomerRepositoryImpl) GetById(c *gin.Context) (*domain.Customer, error) {
	var customer domain.Customer

	r.database.Model(&domain.Customer{}).Preload("PawnOrders").Find(&customer, c.Param("id"))
	if customer.ID == 0 {
		return nil, errors.New("customer not found")
	}
	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAll() (*[]domain.Customer, error) {
	var customers []domain.Customer

	r.database.Model(domain.Customer{}).Preload("PawnOrders").Find(&customers)
	return &customers, nil
}

func (r *CustomerRepositoryImpl) Create(c *gin.Context) (*domain.Customer, error) {
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.Customer{}).Create(&customer)
	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(c *gin.Context) (*domain.Customer, error) {
	patch, customer := map[string]interface{}{}, domain.Customer{}
	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["id"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an ID in the JSON body")
	}

	result := r.database.Model(&domain.Customer{}).Where("id = ?", patch["id"]).Omit("id").Updates(&patch).Find(&customer)
	if result.RowsAffected == 0 {
		return nil, errors.New("customer not found or json data does not match ")
	}

	return &customer, nil
}
