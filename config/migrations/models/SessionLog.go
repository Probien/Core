package models

import "time"

type SessionLog struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
	CreatedAt  time.Time
}
