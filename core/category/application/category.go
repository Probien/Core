package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/category/domain"
	"github.com/JairDavid/Probien-Backend/core/category/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

type CategoryInteractor struct {
}

func (CI *CategoryInteractor) GetById(c *gin.Context) (domain.Category, error) {
	repo := persistance.NewCategoryRepositoryImpl(config.GetDBInstance())
	return repo.GetById(c)
}

func (CI *CategoryInteractor) GetAll() ([]domain.Category, error) {
	repo := persistance.NewCategoryRepositoryImpl(config.GetDBInstance())
	return repo.GetAll()
}

func (CI *CategoryInteractor) Create(c *gin.Context) (domain.Category, error) {
	repo := persistance.NewCategoryRepositoryImpl(config.GetDBInstance())
	return repo.Create(c)
}

func (CI *CategoryInteractor) Delete(c *gin.Context) (domain.Category, error) {
	repo := persistance.NewCategoryRepositoryImpl(config.GetDBInstance())
	return repo.Delete(c)
}

func (CI *CategoryInteractor) Update(c *gin.Context) (domain.Category, error) {
	repo := persistance.NewCategoryRepositoryImpl(config.GetDBInstance())
	return repo.Update(c)
}
