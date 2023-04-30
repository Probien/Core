package response

import "errors"

const (
	//define auth messages
	LoginDone  = "successfully logged in"
	LogoutDone = "successfully logged out"

	//define http errors
	FailedHttpOperation = "failed operation"

	//define CRUD operation messages
	Created   = "successfully created"
	Consulted = "successfully consulted"
	Deleted   = "successfully deleted"
	Updated   = "successfully updated"

	PaginationUrl = "localhost:9000/probien/api/v1/branch-offices/?page="
)

var (
	ErrorBinding = errors.New("[validator] missing or invalid fields: verify JSON struct")
)
