package domain

type Profile struct {
	ID         uint   `json:"id"`
	EmployeeID uint   `json:"employee_id"`
	Name       string `json:"name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
}
