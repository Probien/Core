package domain

import (
	"time"

	"github.com/JairDavid/Probien-Backend/core/product/domain"
)

type Category struct {
	ID            uint             `json:"id"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Interest_rate float64          `json:"interest_rate"`
	Products      []domain.Product `json:"products"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
