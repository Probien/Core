package domain

import "time"

type BranchOffice struct {
	ID         uint       `json:"branch_office_id"`
	Employees  []Employee `json:"employees"`
	BranchName string     `json:"branch_name"`
	Address    string     `json:"address"`
	ZipCode    string     `json:"zip_code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
