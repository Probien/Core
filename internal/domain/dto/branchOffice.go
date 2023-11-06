package dto

import "time"

type BranchOffice struct {
	ID         uint       `json:"branch_office_id"`
	BranchName string     `json:"branch_name" binding:"required"`
	Address    string     `json:"address" binding:"required"`
	ZipCode    string     `json:"zip_code" binding:"required"`
	Employees  []Employee `json:"employees"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
