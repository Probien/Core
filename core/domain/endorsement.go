package domain

import "time"

type Endorsement struct {
	ID          uint    `json:"id"`
	PawnOrderID uint    `json:"pawn_order_id"`
	Payment     float64 `json:"payment"`
	Endorsement time.Time
	CutOffDay   time.Time
}
