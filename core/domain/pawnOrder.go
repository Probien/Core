package domain

import "time"

type PawnOrder struct {
	ID            uint          `json:"id"`
	EmployeeID    uint          `json:"employee_id"`
	Employee      *Employee     `json:"employee,omitempty"`
	CustomerID    uint          `json:"customer_id"`
	Customer      *Customer     `json:"customer,omitempty"`
	StatusID      uint          `json:"status_id"`
	Status        *Status       `json:"status,omitempty"`
	TotalAmount   float64       `json:"total_amount"`
	Monthly       bool          `json:"monthly"`
	Products      []Product     `json:"products"`
	Endorsements  []Endorsement `json:"endorsements_done,omitempty"`
	CutOffDay     time.Time     `json:"cutoff_date"`
	ExtensionDate time.Time     `json:"extension_date"`
}
