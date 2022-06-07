package domain

type Status struct {
	ID         uint         `json:"status_id"`
	Name       string       `json:"status_name"`
	PawnOrders *[]PawnOrder `json:"pawn-orders,omitempty"`
}
