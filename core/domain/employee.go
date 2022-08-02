package domain

import "time"

type Employee struct {
	ID             uint           `json:"id"`
	Profile        *Profile       `json:"profile,omitempty"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	IsActive       bool           `json:"is_active"`
	BranchOfficeID uint           `json:"branch_office_id"`
	PawnOrdersDone *[]PawnOrder   `json:"pawn_orders_done,omitempty"`
	SessionLogs    *[]SessionLog  `json:"sessions,omitempty"`
	Endorsements   *[]Endorsement `json:"endorsements_done,omitempty"`
	Roles          []EmployeeRole `json:"roles,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
