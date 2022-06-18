package persistance

import (
	"encoding/json"
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
		return nil, errors.New(ERROR_PROCCESS)
	}

	if customer.ID == 0 {
		return nil, errors.New(CUSTOMER_NOT_FOUND)
	}
	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAll() (*[]domain.Customer, error) {
	var customers []domain.Customer

	if err := r.database.Model(domain.Customer{}).Preload("PawnOrders").Find(&customers).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	return &customers, nil
}

func (r *CustomerRepositoryImpl) Create(c *gin.Context) (*domain.Customer, error) {
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		return nil, errors.New(ERROR_BINDING)
	}

	if err := r.database.Model(&domain.Customer{}).Create(&customer).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	data, _ := json.Marshal(&customer)
	//replace number 1 for employeeID session (JWT fix)
	go r.database.Exec("CALL savemovement(?,?,?,?)", 2, SP_INSERT, SP_NO_PREV_DATA, string(data[:]))
	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(c *gin.Context) (*domain.Customer, error) {
	patch, customer, customerOld := map[string]interface{}{}, domain.Customer{}, domain.Customer{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New(ERROR_BINDING)
	}

	r.database.Model(&domain.Customer{}).Find(&customerOld, patch["id"])

	if err := r.database.Model(&domain.Customer{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&customer).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if customer.ID == 0 {
		return nil, errors.New(CUSTOMER_NOT_FOUND)
	}

	old, _ := json.Marshal(&customerOld)
	new, _ := json.Marshal(&customer)
	//replace number 1 for employeeID session (JWT fix)
	go r.database.Exec("CALL savemovement(?,?,?,?)", 2, SP_UPDATE, string(old[:]), string(new[:]))
	return &customer, nil
}
