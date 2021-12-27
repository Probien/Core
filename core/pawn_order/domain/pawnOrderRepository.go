package domain

type PawnOrderRepository interface {
	GetById() (PawnOrder, error)
	GetAll() ([]PawnOrder, error)
	Create() (PawnOrder, error)
	Update() (PawnOrder, error)
}
