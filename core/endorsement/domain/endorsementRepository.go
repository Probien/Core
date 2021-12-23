package domain

type EndorsementRepository interface {
	GetById() (Endorsement, error)
	GetAll() ([]Endorsement, error)
	Create() (Endorsement, error)
}
