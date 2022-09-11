package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

type CategoryInteractor struct {
}

func (CI *CategoryInteractor) GetById(c *gin.Context) (*domain.Category, error) {
	repository := persistence.NewCategoryRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (CI *CategoryInteractor) GetAll(c *gin.Context) (*[]domain.Category, error) {
	repository := persistence.NewCategoryRepositoryImpl(config.Database)
	return repository.GetAll(c)
}

func (CI *CategoryInteractor) Create(c *gin.Context) (*domain.Category, error) {
	repository := persistence.NewCategoryRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (CI *CategoryInteractor) Delete(c *gin.Context) (*domain.Category, error) {
	repository := persistence.NewCategoryRepositoryImpl(config.Database)
	return repository.Delete(c)
}

func (CI *CategoryInteractor) Update(c *gin.Context) (*domain.Category, error) {
	repository := persistence.NewCategoryRepositoryImpl(config.Database)
	return repository.Update(c)
}
