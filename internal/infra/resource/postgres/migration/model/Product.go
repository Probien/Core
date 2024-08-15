package model

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	PawnOrderID uint    `gorm:"not null"`
	CategoryID  uint    `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Name        string  `gorm:"type:varchar(20);not null"`
	Brand       string  `gorm:"type:varchar(20);not null"`
	Details     string  `gorm:"type:varchar(250);not null"`
	OnSale      bool    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
