package employee_domain

import "time"

type Employee struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsAdmin    bool   `json:"is_admin"`
	IsActive   bool   `json:"is_active"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
