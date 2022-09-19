package repository

import (
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
)

type IPawnOrderRepository interface {
	GetById(id int) (*domain.PawnOrder, error)
	GetByIdForUpdate(id uint) (*domain.PawnOrder, error)
	GetAll(params url.Values) (*[]domain.PawnOrder, map[string]interface{}, error)
	Create(pawnOrderDto *domain.PawnOrder) (*domain.PawnOrder, error)
	Update(pawnOrderDto map[string]interface{}) (*domain.PawnOrder, error)
}
