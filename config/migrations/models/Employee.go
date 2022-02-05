package models

import "time"

type Employee struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(20);not null"`
	FirstName      string `gorm:"type:varchar(20);not null"`
	SecondName     string `gorm:"type:varchar(20);not null"`
	Address        string `gorm:"type:varchar(50);not null"`
	Phone          string `gorm:"type:varchar(10);not null"`
	Email          string `gorm:"type:varchar(30);unique;not null"`
	Password       string `gorm:"type:varchar(80);not null"`
	IsAdmin        bool   `gorm:"not null"`
	IsActive       bool   `gorm:"not null"`
	BranchOfficeID uint   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
