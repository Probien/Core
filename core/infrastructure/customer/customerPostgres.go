package persistance

import (
	"errors"

	customer_domain "github.com/JairDavid/Probien-Backend/core/domain/customer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	database *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) customer_domain.CustomerRepository {
	return &CustomerRepositoryImpl{database: db}
}

func (r *CustomerRepositoryImpl) GetById(c *gin.Context) (*customer_domain.Customer, error) {
	var customer customer_domain.Customer

	r.database.Model(&customer_domain.Customer{}).Preload("PawnOrders").Find(&customer, c.Param("id"))
	if customer.ID == 0 {
		return nil, errors.New("customer not found")
	}
	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAll() (*[]customer_domain.Customer, error) {
	var customers []customer_domain.Customer

	r.database.Model(customer_domain.Customer{}).Preload("PawnOrders").Find(&customers)
	return &customers, nil
}

func (r *CustomerRepositoryImpl) Create(c *gin.Context) (*customer_domain.Customer, error) {
	var customer customer_domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&customer_domain.Customer{}).Create(&customer)
	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(c *gin.Context) (*customer_domain.Customer, error) {
	patch, customer := map[string]interface{}{}, customer_domain.Customer{}
	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["id"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an ID in the JSON body")
	}

	result := r.database.Model(&customer_domain.Customer{}).Where("id = ?", patch["id"]).Omit("id").Updates(&patch).Find(&customer)
	if result.RowsAffected == 0 {
		return nil, errors.New("customer not found or json data does not match ")
	}

	return &customer, nil
}
