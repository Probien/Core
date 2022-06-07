package domain

import "time"

type Endorsement struct {
	ID          uint    `json:"id"`
	EmployeeID  uint    `json:"employee_id"`
	PawnOrderID uint    `json:"pawn_order_id"`
	Amount      float64 `json:"amount"`
	CreatedAt   time.Time
}
