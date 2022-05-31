package domain

import "time"

type BranchOffice struct {
	ID         uint       `json:"branch_office_id"`
	BranchName string     `json:"branch_name"`
	Address    string     `json:"address"`
	ZipCode    string     `json:"zip_code"`
	Employees  []Employee `json:"employees"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
