package domain

import (
	"time"

	"github.com/JairDavid/Probien-Backend/core/endorsement/domain"
	product "github.com/JairDavid/Probien-Backend/core/product/domain"
)

type PawnOrder struct {
	ID            uint                 `json:"id"`
	EmployeeID    uint                 `json:"employee_id"`
	StatusID      uint                 `json:"status_id"`
	Customer      uint                 `json:"customer"`
	TotalMount    float64              `json:"total_mount"`
	Monthly       bool                 `json:"monthly"`
	Products      []product.Product    `json:"products"`
	Endorsements  []domain.Endorsement `json:"endorsements"`
	CutOffDay     time.Time
	ExtensionDate time.Time `json:"extension_date"`
}
