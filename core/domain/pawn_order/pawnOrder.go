package pawn_order_domain

import (
	"time"

	endorsement_domain "github.com/JairDavid/Probien-Backend/core/domain/endorsement"
	product_domain "github.com/JairDavid/Probien-Backend/core/domain/product"
)

type PawnOrder struct {
	ID            uint                             `json:"id"`
	EmployeeID    uint                             `json:"employee_id"`
	StatusID      uint                             `json:"status_id"`
	Customer      uint                             `json:"customer"`
	TotalMount    float64                          `json:"total_mount"`
	Monthly       bool                             `json:"monthly"`
	Products      []product_domain.Product         `json:"products"`
	Endorsements  []endorsement_domain.Endorsement `json:"endorsements"`
	CutOffDay     time.Time
	ExtensionDate time.Time `json:"extension_date"`
}
