package domain

import "time"

type Customer struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name" binding:"required"`
	FirstName  string      `json:"first_name" binding:"required"`
	SecondName string      `json:"second_name" binding:"required"`
	Address    string      `json:"address" binding:"required"`
	Phone      string      `json:"phone" binding:"required"`
	PawnOrders []PawnOrder `json:"pawn_orders"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
