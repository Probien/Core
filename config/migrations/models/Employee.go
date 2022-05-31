package models

import "time"

type Employee struct {
	ID             uint `gorm:"primaryKey"`
	Profile        Profile
	Email          string `gorm:"type:varchar(30);unique;not null"`
	Password       string `gorm:"type:varchar(80);not null"`
	IsAdmin        bool   `gorm:"not null"`
	IsActive       bool   `gorm:"not null"`
	BranchOfficeID uint   `gorm:"not null"`
	PawnOrdersDone []PawnOrder
	SessionLogs    []SessionLog
	Endorsements   []PawnOrder `gorm:"many2many:endorsements"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
