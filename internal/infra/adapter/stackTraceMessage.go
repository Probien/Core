package adapter

import "errors"

const (
	//define constants for store procedure actions
	SpInsert = "INSERT"
	SpDelete = "DELETE"
	SpUpdate = "UPDATE"

	//define constants for store procedure state values
	SpNoPrevData = "No previous data"
	SpNoCurrData = "No current data"
)

var (
	//define general errors
	InvalidAction      = errors.New("[persistence] impossible to delete information with dependent data")
	ErrorProcess       = errors.New("[persistence] could not execute process")
	InvalidCredentials = errors.New("[authorization] incorrect credentials")
	ErrorBinding       = errors.New("[validator] missing or invalid fields: verify JSON struct")

	//define query errors for each model
	BranchNotFound      = errors.New("[persistence] branch office not found")
	CategoryNotFound    = errors.New("[persistence] category not found")
	CustomerNotFound    = errors.New("[persistence] customer not found")
	EmployeeNotFound    = errors.New("[persistence] employee not found")
	EndorsementNotFound = errors.New("[persistence] endorsement not found")
	PawnOrderNotFound   = errors.New("[persistence] pawn order not found")
	ProductNotFound     = errors.New("[persistence] product not found")
)
