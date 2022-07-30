package domain

type Role struct {
	ID          uint   `json:"-"`
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
}
