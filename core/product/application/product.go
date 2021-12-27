package application

import (
	"github.com/JairDavid/Probien-Backend/config"

	"github.com/JairDavid/Probien-Backend/core/product/domain"
	"github.com/JairDavid/Probien-Backend/core/product/infrastructure/persistance"
)

func GetById() (domain.Product, error) {
	return persistance.NewProductRepositoryImpl(config.GetDBInstance()).GetById()
}

func GetAll() ([]domain.Product, error) {
	return persistance.NewProductRepositoryImpl(config.GetDBInstance()).GetAll()
}

func Create() ([]domain.Product, error) {
	return persistance.NewProductRepositoryImpl(config.GetDBInstance()).Create()
}

func Update() (domain.Product, error) {
	return persistance.NewProductRepositoryImpl(config.GetDBInstance()).Update()
}
