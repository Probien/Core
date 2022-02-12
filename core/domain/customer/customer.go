package customer_domain

import (
	"time"

	pawn_order_domain "github.com/JairDavid/Probien-Backend/core/domain/pawn_order"
)

type Customer struct {
	ID         uint                          `json:"id"`
	Name       string                        `json:"name"`
	FirstName  string                        `json:"first_name"`
	SecondName string                        `json:"second_name"`
	Address    string                        `json:"address"`
	Phone      string                        `json:"phone"`
	PawnOrders []pawn_order_domain.PawnOrder `json:"pawn_orders"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
