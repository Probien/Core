package models

import "time"

type Endorsement struct {
	EmployeeID  uint    `gorm:"primaryKey"`
	PawnOrderID uint    `gorm:"primaryKey"`
	Payment     float64 `gorm:"not null;"`
	Endorsement time.Time
	CutOffDay   time.Time
}
