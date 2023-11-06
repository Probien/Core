package domain

import "time"

type ModerationLog struct {
	ID            uint      `json:"id"`
	TriggeredBy   uint      `json:"triggered_by"`
	Action        string    `json:"action"`
	PreviousValue string    `json:"previous_value"`
	CurrentValue  string    `json:"current_value"`
	CreatedAt     time.Time `json:"created_at"`
}
