package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
)

type ISessionRepository interface {
	GenerateSessionID(employee *dto.Employee, session chan<- component.SessionCredential)
	ClearSessionID(cookie string) error
	ExistCookie(cookie string, checker chan<- bool)
}
