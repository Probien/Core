package adapter

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/adapter"
	"math"
	"net/url"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	database *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) port.ICustomerRepository {
	return &CustomerRepositoryImpl{database: db}
}

func (r *CustomerRepositoryImpl) GetById(id int) (*dto.Customer, error) {
	var customer dto.Customer

	if err := r.database.Model(&dto.Customer{}).Preload("PawnOrders.Products").Preload("PawnOrders.Endorsements").Find(&customer, id).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	if customer.ID == 0 {
		return nil, adapter.CustomerNotFound
	}
	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAll(params url.Values) (*[]dto.Customer, map[string]interface{}, error) {
	var customers []dto.Customer
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("customers").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(dto.Customer{}).Scopes(adapter.Paginate(params, paginationResult)).Preload("PawnOrders").Find(&customers).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &customers, paginationResult, nil
}

func (r *CustomerRepositoryImpl) Create(customerDto *dto.Customer, userSessionId int) (*dto.Customer, error) {

	if err := r.database.Model(&dto.Customer{}).Create(&customerDto).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	data, _ := json.Marshal(&customerDto)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpInsert, adapter.SpNoPrevData, string(data[:]))
	return customerDto, nil
}

func (r *CustomerRepositoryImpl) Update(id int, customerDto map[string]interface{}, userSessionId int) (*dto.Customer, error) {
	customer, customerOld := dto.Customer{}, dto.Customer{}

	r.database.Model(&dto.Customer{}).Find(&customerOld, id)

	if customerOld.ID == 0 {
		return nil, adapter.CustomerNotFound
	}

	if err := r.database.Model(&dto.Customer{}).Where("id = ?", id).Updates(&customerDto).Find(&customer).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	old, _ := json.Marshal(&customerOld)
	current, _ := json.Marshal(&customer)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpUpdate, string(old[:]), string(current[:]))
	return &customer, nil
}
