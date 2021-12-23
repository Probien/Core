package domain

type CategoryRepository interface {
	GetById() (Category, error)
	GetAll() ([]Category, error)
	Create() (Category, error)
	Delete() (Category, error)
	Update() (Category, error)
}
