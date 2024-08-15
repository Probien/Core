package model

import "time"

type ModerationLog struct {
	ID            uint `gorm:"primaryKey"`
	TriggeredBy   uint
	Action        string
	PreviousValue string
	CurrentValue  string
	CreatedAt     time.Time
}
