package domain

import "time"

type SessionLog struct {
	ID         uint     `json:"id"`
	EmployeeID uint     `json:"employee_id"`
	Employee   Employee `json:"employee"`
	CreatedAt  time.Time
}
