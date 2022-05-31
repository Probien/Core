package models

import "time"

type Endorsement struct {
	EmployeeID  uint    `gorm:"primaryKey"`
	PawnOrderID uint    `gorm:"primaryKey"`
	Amount      float64 `gorm:"not null;"`
	PaidAt      time.Time
}
