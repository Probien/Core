package domain

import "time"

type ModerationLog struct {
	ID            uint     `json:"id"`
	EmployeeID    uint     `json:"employee_id"`
	Employee      Employee `json:"employee"`
	Action        string   `json:"action"`
	PreviousValue string   `json:"previous_value"`
	CurrentValue  string   `json:"current_value"`
	CreatedAt     time.Time
}
