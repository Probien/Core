package domain

import "time"

type Endorsement struct {
	EmployeeID  uint    `json:"employee_id"`
	PawnOrderID uint    `json:"pawn_order_id"`
	Amount      float64 `json:"amount"`
	PaidAt      time.Time
}
