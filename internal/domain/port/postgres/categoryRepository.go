package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type ICategoryRepository interface {
	GetById(id int) (*dto.Category, error)
	GetAll(params url.Values) (*[]dto.Category, map[string]interface{}, error)
	Create(categoryDto *dto.Category, userSessionId int) (*dto.Category, error)
	Delete(id int, userSessionId int) (*dto.Category, error)
	Update(id int, categoryDto map[string]interface{}, userSessionId int) (*dto.Category, error)
}
