package domain

import "time"

type PawnOrder struct {
	ID            uint           `json:"id"`
	EmployeeID    uint           `json:"employee_id" binding:"required"`
	Employee      *Employee      `json:"employee,omitempty"`
	CustomerID    uint           `json:"customer_id" binding:"required"`
	Customer      *Customer      `json:"customer,omitempty"`
	StatusID      uint           `json:"status_id" binding:"required"`
	Status        *Status        `json:"status,omitempty"`
	TotalAmount   float64        `json:"total_amount" binding:"required"`
	Monthly       bool           `json:"monthly" binding:"required"`
	Products      *[]Product     `json:"products,omitempty"`
	Endorsements  *[]Endorsement `json:"endorsements_done,omitempty"`
	CutOffDay     time.Time      `json:"cutoff_date"`
	ExtensionDate time.Time      `json:"extension_date"`
}
