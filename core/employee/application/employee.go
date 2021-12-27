package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/employee/domain"
	"github.com/JairDavid/Probien-Backend/core/employee/infrastructure/persistance"
)

func GetById() (domain.Employee, error) {
	return persistance.NewEmployeeRepositoryImpl(config.GetDBInstance()).GetById()
}

func GetAll() ([]domain.Employee, error) {
	return persistance.NewEmployeeRepositoryImpl(config.GetDBInstance()).GetAll()
}

func Create() (domain.Employee, error) {
	return persistance.NewEmployeeRepositoryImpl(config.GetDBInstance()).Create()
}

func Update() (domain.Employee, error) {
	return persistance.NewEmployeeRepositoryImpl(config.GetDBInstance()).Update()
}
