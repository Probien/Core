package adapter

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"gorm.io/gorm"
)

type AuthRepositoryImp struct {
	database *gorm.DB
}

func NewAuthRepositoryImp(db *gorm.DB) port.IAuthRepository {
	return AuthRepositoryImp{
		database: db,
	}
}

func (a AuthRepositoryImp) Login(loginCredential component.Credential) (*dto.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepositoryImp) Logout(session string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepositoryImp) RecoverPassword(email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
