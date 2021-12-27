package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/customer/domain"
	"github.com/JairDavid/Probien-Backend/core/customer/infrastructure/persistance"
)

func GetById() (domain.Customer, error) {
	return persistance.NewCustomerRepositoryImpl(config.GetDBInstance()).GetById()
}

func GetAll() ([]domain.Customer, error) {
	return persistance.NewCustomerRepositoryImpl(config.GetDBInstance()).GetAll()
}

func Create() (domain.Customer, error) {
	return persistance.NewCustomerRepositoryImpl(config.GetDBInstance()).Create()
}

func Update() (domain.Customer, error) {
	return persistance.NewCustomerRepositoryImpl(config.GetDBInstance()).Update()
}
