package domain

import "time"

type Endorsement struct {
	ID          uint      `json:"id"`
	EmployeeID  uint      `json:"employee_id" binding:"required"`
	PawnOrderID uint      `json:"pawn_order_id" binding:"required"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}
