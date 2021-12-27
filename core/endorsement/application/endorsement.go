package application

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/endorsement/domain"
	"github.com/JairDavid/Probien-Backend/core/endorsement/infrastructure/persistance"
)

func GetById() (domain.Endorsement, error) {
	return persistance.NewEndorsementRepositoryImpl(config.GetDBInstance()).GetById()
}

func GetAll() ([]domain.Endorsement, error) {
	return persistance.NewEndorsementRepositoryImpl(config.GetDBInstance()).GetAll()
}

func Create() (domain.Endorsement, error) {
	return persistance.NewEndorsementRepositoryImpl(config.GetDBInstance()).Create()
}
