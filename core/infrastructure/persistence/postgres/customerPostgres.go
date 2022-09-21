package postgres

import (
	"encoding/json"
	"math"
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	database *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) repository.ICustomerRepository {
	return &CustomerRepositoryImpl{database: db}
}

func (r *CustomerRepositoryImpl) GetById(id int) (*domain.Customer, error) {
	var customer domain.Customer

	if err := r.database.Model(&domain.Customer{}).Preload("PawnOrders.Products").Preload("PawnOrders.Endorsements").Find(&customer, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if customer.ID == 0 {
		return nil, persistence.CustomerNotFound
	}
	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAll(params url.Values) (*[]domain.Customer, map[string]interface{}, error) {
	var customers []domain.Customer
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("customers").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(domain.Customer{}).Scopes(persistence.Paginate(params, paginationResult)).Preload("PawnOrders").Find(&customers).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &customers, paginationResult, nil
}

func (r *CustomerRepositoryImpl) Create(customerDto *domain.Customer, userSessionId int) (*domain.Customer, error) {

	if err := r.database.Model(&domain.Customer{}).Create(&customerDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&customerDto)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return customerDto, nil
}

func (r *CustomerRepositoryImpl) Update(id int, customerDto map[string]interface{}, userSessionId int) (*domain.Customer, error) {
	customer, customerOld := domain.Customer{}, domain.Customer{}

	r.database.Model(&domain.Customer{}).Find(&customerOld, id)

	if customer.ID == 0 {
		return nil, persistence.CustomerNotFound
	}

	if err := r.database.Model(&domain.Customer{}).Where("id = ?", id).Updates(&customerDto).Find(&customer).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	old, _ := json.Marshal(&customerOld)
	current, _ := json.Marshal(&customer)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpUpdate, string(old[:]), string(current[:]))
	return &customer, nil
}
