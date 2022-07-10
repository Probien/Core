package persistance

const (
	//define constants for store procedure actions
	SP_INSERT = "INSERT"
	SP_DELETE = "DELETE"
	SP_UPDATE = "UPDATE"

	//define constants for store procedure state values
	SP_NO_PREV_DATA = "No previous data"
	SP_NO_CURR_DATA = "No current data"

	//define general errors
	INVALID_ACTION      = "impossible to delete information with dependent data"
	INVALID_CREDENTIALS = "incorrect credentials"
	ERROR_PROCCESS      = "could not execute process in our database services"
	ERROR_BINDING       = "error binding JSON data (missing or invalid fields), verify JSON struct"

	//define query errors for each model
	BRANCH_NOT_FOUND      = "branch office not found"
	CATEGORY_NOT_FOUND    = "category not found"
	CUSTOMER_NOT_FOUND    = "customer not found"
	EMPLOYEE_NOT_FOUND    = "employee not foud"
	ENDORSEMENT_NOT_FOUND = "endorsement not found"
	PAWNORDER_NOT_FOUND   = "pawn order not found"
	PRODUCT_NOT_FOUND     = "product not found"
)
