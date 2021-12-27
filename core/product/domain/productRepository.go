package domain

type ProductRepository interface {
	GetById() (Product, error)
	GetAll() ([]Product, error)
	Create() ([]Product, error)
	Update() (Product, error)
}
