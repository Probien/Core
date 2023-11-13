package application

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	authPort "github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	sessionPort "github.com/JairDavid/Probien-Backend/internal/domain/port/redis"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
)

type AuthApp struct {
	authPort    authPort.IAuthRepository
	sessionPort sessionPort.ISessionRepository
}

func NewAuthApp(authPort authPort.IAuthRepository, sessionPort sessionPort.ISessionRepository) AuthApp {
	return AuthApp{
		authPort:    authPort,
		sessionPort: sessionPort,
	}
}

func (a AuthApp) Login(loginCredential component.Credential) (*dto.Employee, error) {
	return a.authPort.Login(loginCredential)
}

func (a AuthApp) Logout(session string) (bool, error) {
	return a.authPort.Logout(session)
}

func (a AuthApp) RecoverPassword(email string) (bool, error) {
	return a.authPort.RecoverPassword(email)
}

func (a AuthApp) GenerateSessionID(employee *dto.Employee, session chan<- component.SessionCredential) {
	a.sessionPort.GenerateSessionID(employee, session)
}

func (a AuthApp) ClearSessionID(cookie string) error {
	return a.sessionPort.ClearSessionID(cookie)
}

func (a AuthApp) ExistCookie(cookie string, checker chan<- bool) {
	a.sessionPort.ExistCookie(cookie, checker)
}
