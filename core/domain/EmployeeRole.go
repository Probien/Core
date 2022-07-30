package domain

type EmployeeRole struct {
	EmployeeID uint `json:"-"`
	RoleID     uint `json:"role_id"`
	Role       Role `json:"role"`
}
