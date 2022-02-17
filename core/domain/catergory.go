package domain

import (
	"time"
)

type Category struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Interest_rate float64   `json:"interest_rate"`
	Products      []Product `json:"products"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
