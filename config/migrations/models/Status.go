package models

type Status struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(10);not null"`
	PawnOrders []PawnOrder
}
