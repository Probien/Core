package domain

import "time"

type PawnOrder struct {
	ID            uint    `json:"id"`
	EmployeeID    uint    `json:"employee_id"`
	CustomerID    uint    `json:"customer_id"`
	StatusID      uint    `json:"status_id"`
	TotalMount    float64 `json:"total_mount"`
	Monthly       bool    `json:"monthly"`
	CutOffDay     time.Time
	ExtensionDate time.Time `json:"extension_date"`
}
