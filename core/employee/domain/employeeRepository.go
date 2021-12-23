package domain

type EmployeeRepository interface {
	GetById() (Employee, error)
	GetAll() ([]Employee, error)
	Create() (Employee, error)
	Update() (Employee, error)
}
