package domain

import "time"

type PaymentLog struct {
	ID         uint     `json:"id"`
	EmployeeID uint     `json:"employee_id" binding:"required"`
	Employee   Employee `json:"employee"`
	CustomerID uint     `json:"customer_id" binding:"required"`
	Customer   Customer `json:"customer"`
	Amount     float64  `json:"payment" binding:"required"`
	CreatedAt  time.Time
}
