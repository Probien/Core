package domain

import "time"

type Product struct {
	ID          uint    `json:"id"`
	PawnOrderID uint    `json:"pawn_order_id"`
	CategoryID  uint    `json:"category_id"`
	Price       float64 `json:"price"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Details     string  `json:"details"`
	OnSale      bool    `json:"on_sale"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
