package models

import "time"

type Customer struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(20);not null"`
	FirstName  string `gorm:"type:varchar(20);not null"`
	SecondName string `gorm:"type:varchar(20);not null"`
	Address    string `gorm:"type:varchar(50);not null"`
	Phone      string `gorm:"type:varchar(10);not null"`
	PawnOrders []PawnOrder
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
