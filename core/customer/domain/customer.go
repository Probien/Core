package domain

import (
	"time"

	"github.com/JairDavid/Probien-Backend/core/pawn_order/domain"
)

type Customer struct {
	ID         uint               `json:"id"`
	Name       string             `json:"name"`
	FirstName  string             `json:"first_name"`
	SecondName string             `json:"second_name"`
	Address    string             `json:"address"`
	Phone      string             `json:"phone"`
	PawnOrders []domain.PawnOrder `json:"pawn_orders"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
