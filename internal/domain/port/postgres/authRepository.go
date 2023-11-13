package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
)

type IAuthRepository interface {
	Login(loginCredential component.Credential) (*dto.Employee, error)
	Logout(session string) (bool, error)
	RecoverPassword(email string) (bool, error)
}
