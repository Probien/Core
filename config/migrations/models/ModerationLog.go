package models

import "time"

type ModerationLog struct {
	ID            uint `gorm:"primaryKey"`
	EmployeeID    uint
	Employee      Employee `gorm:"foreignKey:EmployeeID"`
	Action        string
	PreviousValue string
	CurrentValue  string
	CreatedAt     time.Time
}
