package port

import (
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/auth"
)

type IAuthRepository interface {
	Login(loginCredential auth.LoginCredential) (*dto.Employee, error)
	Logout()
	RecoverPassword()
}
