package model

type Role struct {
	ID          uint        `gorm:"primaryKey"`
	RoleName    string      `gorm:"type:varchar(30);not null"`
	Description string      `gorm:"type:varchar(50);not null"`
	Employees   *[]Employee `gorm:"many2many:employee_roles;"`
}
