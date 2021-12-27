package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/category/domain"
	"github.com/JairDavid/Probien-Backend/core/category/infrastructure/persistance"
)

func GetById() (domain.Category, error) {
	return persistance.NewCategoryRepositoryImpl(config.GetDBInstance()).GetById()
}

func GetAll() ([]domain.Category, error) {
	return persistance.NewCategoryRepositoryImpl(config.GetDBInstance()).GetAll()
}

func Create() (domain.Category, error) {
	return persistance.NewCategoryRepositoryImpl(config.GetDBInstance()).Create()
}

func Delete() (domain.Category, error) {
	return persistance.NewCategoryRepositoryImpl(config.GetDBInstance()).Delete()
}

func Update() (domain.Category, error) {
	return persistance.NewCategoryRepositoryImpl(config.GetDBInstance()).Update()
}
