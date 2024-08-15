package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type IProductRepository interface {
	GetById(id int) (*dto.Product, error)
	GetAll(params url.Values) (*[]dto.Product, map[string]interface{}, error)
	Create(productDto *dto.Product, userSessionId int) (*dto.Product, error)
	Update(id int, productDto map[string]interface{}, userSessionId int) (*dto.Product, error)
}
