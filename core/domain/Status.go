package domain

type Status struct {
	ID         uint   `json:"status_id"`
	Name       string `gorm:"status_name"`
	PawnOrders []PawnOrder
}
