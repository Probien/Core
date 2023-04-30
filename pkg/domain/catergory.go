package domain

import (
	"time"
)

type Category struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	Interest_rate float64   `json:"interest_rate" binding:"required"`
	Products      []Product `json:"products,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
