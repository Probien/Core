package dto

import "time"

type Category struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	InterestRate float64   `json:"interest_rate" binding:"required"`
	Products     []Product `json:"products"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
