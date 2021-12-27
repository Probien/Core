package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/customer/domain"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	database *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) domain.CustomerRepository {
	return &CustomerRepositoryImpl{database: db}
}

func (r *CustomerRepositoryImpl) GetById() (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (r *CustomerRepositoryImpl) GetAll() ([]domain.Customer, error) {
	return []domain.Customer{}, nil
}

func (r *CustomerRepositoryImpl) Create() (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (r *CustomerRepositoryImpl) Update() (domain.Customer, error) {
	return domain.Customer{}, nil
}
