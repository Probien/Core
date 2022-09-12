package postgres

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

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
		return nil, persistence.ErrorProcess
	}

	if customer.ID == 0 {
		return nil, persistence.CustomerNotFound
	}
	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAll(c *gin.Context) (*[]domain.Customer, error) {
	var customers []domain.Customer

	if err := r.database.Model(domain.Customer{}).Preload("PawnOrders").Find(&customers).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	return &customers, nil
}

func (r *CustomerRepositoryImpl) Create(c *gin.Context) (*domain.Customer, error) {
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		return nil, persistence.ErrorBinding
	}

	if err := r.database.Model(&domain.Customer{}).Create(&customer).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&customer)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return &customer, nil
}

func (r *CustomerRepositoryImpl) Update(c *gin.Context) (*domain.Customer, error) {
	patch, customer, customerOld := map[string]interface{}{}, domain.Customer{}, domain.Customer{}

	if err := c.Bind(&patch); err != nil {
		return nil, err
	}

	_, errID := patch["id"]

	if !errID {
		return nil, persistence.ErrorBinding
	}

	r.database.Model(&domain.Customer{}).Find(&customerOld, patch["id"])

	if err := r.database.Model(&domain.Customer{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&customer).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if customer.ID == 0 {
		return nil, persistence.CustomerNotFound
	}

	old, _ := json.Marshal(&customerOld)
	current, _ := json.Marshal(&customer)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpUpdate, string(old[:]), string(current[:]))
	return &customer, nil
}
