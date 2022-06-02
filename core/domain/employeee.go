package domain

import "time"

type Employee struct {
	ID               uint         `json:"id"`
	Profile          Profile      `json:"profile"`
	Email            string       `json:"email"`
	Password         string       `json:"password"`
	IsAdmin          bool         `json:"is_admin"`
	IsActive         bool         `json:"is_active"`
	BranchOfficeID   uint         `json:"branch_office_id"`
	PawnOrdersDone   []PawnOrder  `json:"pawn_orders_done"`
	SessionLogs      []SessionLog `json:"sessions"`
	EndorsementsDone []PawnOrder  `json:"endorsements_done"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
