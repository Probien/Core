package domain

import (
	"time"
)

type PawnOrder struct {
	ID         uint      `json:"id"`
	EmployeeID uint      `json:"employee_id"`
	Employee   *Employee `json:"employee"`
	CustomerID uint      `json:"customer_id"`
	Customer   *Customer `json:"customer"`
	StatusID   uint      `json:"status_id"`

	TotalMount    float64   `json:"total_mount"`
	Monthly       bool      `json:"monthly"`
	Products      []Product `json:"products"`
	CutOffDay     time.Time
	ExtensionDate time.Time `json:"extension_date"`
}
