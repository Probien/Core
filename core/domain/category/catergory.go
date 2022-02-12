package category_domain

import (
	"time"

	product_domain "github.com/JairDavid/Probien-Backend/core/domain/product"
)

type Category struct {
	ID            uint                     `json:"id"`
	Name          string                   `json:"name"`
	Description   string                   `json:"description"`
	Interest_rate float64                  `json:"interest_rate"`
	Products      []product_domain.Product `json:"products"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
