package domain

type CustomerRepository interface {
	GetById() (Customer, error)
	GetAll() ([]Customer, error)
	Create() (Customer, error)
	Update() (Customer, error)
}
