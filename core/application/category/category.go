package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	category_domain "github.com/JairDavid/Probien-Backend/core/domain/category"
	"github.com/gin-gonic/gin"
)

type CategoryInteractor struct {
}

func (CI *CategoryInteractor) GetById(c *gin.Context) (*category_domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.Database)
	return repository.GetById(c)
}

func (CI *CategoryInteractor) GetAll() (*[]category_domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (CI *CategoryInteractor) Create(c *gin.Context) (*category_domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (CI *CategoryInteractor) Delete(c *gin.Context) (*category_domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.Database)
	return repository.Delete(c)
}

func (CI *CategoryInteractor) Update(c *gin.Context) (*category_domain.Category, error) {
	repository := persistance.NewCategoryRepositoryImpl(config.Database)
	return repository.Update(c)
}
