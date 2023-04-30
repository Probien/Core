package domain

import "time"

type Product struct {
	ID          uint    `json:"id"`
	PawnOrderID uint    `json:"pawn_order_id" binding:"required"`
	CategoryID  uint    `json:"category_id" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Brand       string  `json:"brand" binding:"required"`
	Details     string  `json:"details" binding:"required"`
	OnSale      bool    `json:"on_sale" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
