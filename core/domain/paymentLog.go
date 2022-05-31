package domain

import "time"

type PaymentLog struct {
	ID         uint     `json:"id"`
	EmployeeID uint     `json:"employee_id"`
	Employee   Employee `json:"employee"`
	CustomerID uint     `json:"customer_id"`
	Customer   Customer `json:"customer"`
	Amount     float64  `json:"payment"`
	CreatedAt  time.Time
}
