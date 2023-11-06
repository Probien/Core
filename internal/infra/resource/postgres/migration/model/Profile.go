package model

type Profile struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Name       string `gorm:"type:varchar(20);not null"`
	FirstName  string `gorm:"type:varchar(20);not null"`
	SecondName string `gorm:"type:varchar(20);not null"`
	Address    string `gorm:"type:varchar(50);not null"`
	Phone      string `gorm:"type:varchar(10);not null"`
}
