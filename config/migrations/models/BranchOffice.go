package models

import "time"

type BranchOffice struct {
	ID         uint `gorm:"primaryKey"`
	Employees  []Employee
	BranchName string  `gorm:"type:varchar(30);not null"`
	Payment    float64 `gorm:"not null"`
	Address    string  `gorm:"type:varchar(60);not null"`
	ZipCode    string  `gorm:"type:varchar(5);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
