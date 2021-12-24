package models

import "time"

type PawnOrder struct {
	ID            uint    `gorm:"primaryKey"`
	EmployeeID    uint    `gorm:"not null"`
	CustomerID    uint    `gorm:"not null"`
	StatusID      uint    `gorm:"not null"`
	TotalMount    float64 `gorm:"not null"`
	Monthly       bool    `gorm:"not null"`
	CutOffDay     time.Time
	ExtensionDate time.Time
}
