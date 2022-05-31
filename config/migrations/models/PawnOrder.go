package models

import "time"

type PawnOrder struct {
	ID            uint `gorm:"primaryKey"`
	EmployeeID    uint
	Employee      Employee `gorm:"foreignKey:EmployeeID;"`
	CustomerID    uint
	Customer      Customer `gorm:"foreignKey:CustomerID;"`
	StatusID      uint
	Status        Status  `gorm:"foreignKey:StatusID;"`
	TotalMount    float64 `gorm:"not null"`
	Monthly       bool    `gorm:"not null"`
	Products      []Product
	CutOffDay     time.Time
	ExtensionDate time.Time
}
