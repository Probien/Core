package models

import "time"

type Endorsement struct {
	ID          uint    `gorm:"primaryKey"`
	PawnOrderID uint    `gorm:"not null"`
	Payment     float64 `gorm:"not null"`
	Endorsement time.Time
	CutOffDay   time.Time
}
