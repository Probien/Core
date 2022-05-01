package models

import "time"

type BranchOffice struct {
	ID         uint   `gorm:"primaryKey"`
	BranchName string `gorm:"type:varchar(30);not null"`
	Address    string `gorm:"type:varchar(60);not null"`
	ZipCode    string `gorm:"type:varchar(5);not null"`
	Employees  []Employee
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
