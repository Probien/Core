package domain

import "time"

type BranchOffice struct {
	ID        uint       `json:"branch_office_id"`
	Employees []Employee `json:"employees"`
	Payment   float64    `json:"payment"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
