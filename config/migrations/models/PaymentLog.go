package models

import "time"

type PaymentLog struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
	CustomerID uint
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	Payment    float64  `gorm:"not null"`
	CreatedAt  time.Time
}
