package models

import "time"

type BranchOffice struct {
	ID        uint `gorm:"primaryKey"`
	Employees []Employee
	Payment   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
