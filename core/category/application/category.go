package application

import (
	"github.com/JairDavid/Probien-Backend/core/category/domain"
	"github.com/gin-gonic/gin"
)

type CategoryInteractor struct {
	repository domain.CategoryRepository
}

func (CI *CategoryInteractor) GetById(c *gin.Context) (domain.Category, error) {
	return CI.repository.GetById(c)
}

func (CI *CategoryInteractor) GetAll() ([]domain.Category, error) {
	return CI.repository.GetAll()
}

func (CI *CategoryInteractor) Create(c *gin.Context) (domain.Category, error) {
	return CI.repository.Create(c)
}

func (CI *CategoryInteractor) Delete(c *gin.Context) (domain.Category, error) {
	return CI.repository.Delete(c)
}

func (CI *CategoryInteractor) Update(c *gin.Context) (domain.Category, error) {
	return CI.repository.Update(c)
}
