package domain

import "time"

type Employee struct {
	ID             uint           `json:"id"`
	Profile        *Profile       `json:"profile,omitempty"`
	Email          string         `json:"email" binding:"required"`
	Password       string         `json:"password" binding:"required"`
	IsActive       bool           `json:"is_active" binding:"required"`
	BranchOfficeID uint           `json:"branch_office_id" binding:"required"`
	PawnOrdersDone []PawnOrder    `json:"pawn_orders_done"`
	SessionLogs    []SessionLog   `json:"sessions"`
	Endorsements   []Endorsement  `json:"endorsements_done"`
	Roles          []EmployeeRole `json:"roles"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
