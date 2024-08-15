package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"net/url"
)

type CategoryApp struct {
	port port.ICategoryRepository
}

func NewCategoryApp(repository port.ICategoryRepository) CategoryApp {
	return CategoryApp{
		port: repository,
	}
}

func (c *CategoryApp) GetById(id int) (*dto.Category, error) {
	return c.port.GetById(id)
}

func (c *CategoryApp) GetAll(params url.Values) (*[]dto.Category, map[string]interface{}, error) {
	return c.port.GetAll(params)
}

func (c *CategoryApp) Create(categoryDto *dto.Category, userSessionId int) (*dto.Category, error) {
	return c.port.Create(categoryDto, userSessionId)
}

func (c *CategoryApp) Delete(id int, userSessionId int) (*dto.Category, error) {
	return c.port.Delete(id, userSessionId)
}

func (c *CategoryApp) Update(id int, categoryDto map[string]interface{}, userSessionId int) (*dto.Category, error) {
	return c.port.Update(id, categoryDto, userSessionId)
}
