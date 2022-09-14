package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence/postgres"

	"github.com/gin-gonic/gin"
)

type EmployeeInteractor struct {
}

func (EI *EmployeeInteractor) Login(c *gin.Context) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.Login(c)
}

func (EI *EmployeeInteractor) GetByEmail(c *gin.Context) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.GetByEmail(c)
}

func (EI *EmployeeInteractor) GetAll(c *gin.Context) (*[]domain.Employee, map[string]interface{}, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.GetAll(c)
}

func (EI *EmployeeInteractor) Create(c *gin.Context) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (EI *EmployeeInteractor) Update(c *gin.Context) (*domain.Employee, error) {
	repository := postgres.NewEmployeeRepositoryImpl(config.Database)
	return repository.Update(c)
}
