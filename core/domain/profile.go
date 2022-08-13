package domain

type Profile struct {
	ID         uint   `json:"id"`
	EmployeeID uint   `json:"employee_id"`
	Name       string `json:"name" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
}
