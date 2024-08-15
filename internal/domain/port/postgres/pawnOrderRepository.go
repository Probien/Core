package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"net/url"
)

type IPawnOrderRepository interface {
	GetById(id int) (*dto.PawnOrder, error)
	GetByIdForUpdate(id uint) (*dto.PawnOrder, error)
	GetAll(params url.Values) (*[]dto.PawnOrder, map[string]interface{}, error)
	Create(pawnOrderDto *dto.PawnOrder, userSessionId int) (*dto.PawnOrder, error)
	Update(id int, pawnOrderDto map[string]interface{}, userSessionId int) (*dto.PawnOrder, error)
}
