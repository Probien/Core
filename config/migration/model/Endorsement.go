package model

import "time"

type Endorsement struct {
	ID          uint    `gorm:"primaryKey"`
	EmployeeID  uint    `gorm:"not null"`
	PawnOrderID uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null;"`
	CreatedAt   time.Time
}
