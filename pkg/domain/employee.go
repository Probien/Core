package domain

import "time"

type Employee struct {
	ID             uint           `json:"id"`
	Profile        *Profile       `json:"profile,omitempty"`
	Email          string         `json:"email" binding:"required"`
	Password       string         `json:"password" binding:"required"`
	IsActive       bool           `json:"is_active" binding:"required"`
	BranchOfficeID uint           `json:"branch_office_id" binding:"required"`
	PawnOrdersDone []PawnOrder    `json:"pawn_orders_done,omitempty"`
	SessionLogs    []SessionLog   `json:"sessions,omitempty"`
	Endorsements   []Endorsement  `json:"endorsements_done,omitempty"`
	Roles          []EmployeeRole `json:"roles,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
