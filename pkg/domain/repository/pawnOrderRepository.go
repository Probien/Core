package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/pkg/domain"
)

type IPawnOrderRepository interface {
	GetById(id int) (*domain.PawnOrder, error)
	GetByIdForUpdate(id uint) (*domain.PawnOrder, error)
	GetAll(params url.Values) (*[]domain.PawnOrder, map[string]interface{}, error)
	Create(pawnOrderDto *domain.PawnOrder, userSessionId int) (*domain.PawnOrder, error)
	Update(id int, pawnOrderDto map[string]interface{}, userSessionId int) (*domain.PawnOrder, error)
}
