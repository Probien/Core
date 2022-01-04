package application

import (
	"encoding/base64"
	"time"

	"github.com/JairDavid/Probien-Backend/core/employee/domain"
	"github.com/JairDavid/Probien-Backend/core/employee/infrastructure/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type EmployeeInteractor struct {
	repository domain.EmployeeRepository
}

func (EI *EmployeeInteractor) GenerateToken(data *domain.Employee) string {

	claims := &auth.AuthCustomClaims{
		Name:       data.Name,
		FirstName:  data.FirstName,
		SecondName: data.SecondName,
		Email:      data.Email,
		IsAdmin:    data.IsAdmin,
		CreatedAt:  data.CreatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "Probien",
			Subject:   data.Email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(base64.StdEncoding.EncodeToString([]byte("EQVJ7UM8xJNcfsaxs$aw3Es2Z@8ewegzxZ531C$^bhEoMq!%fe"))))
	if err != nil {
		panic(err)
	}
	return token
}

func (EI *EmployeeInteractor) Login(c *gin.Context) (domain.Employee, bool) {
	return EI.repository.Login(c)
}

func (EI *EmployeeInteractor) GetByEmail(c *gin.Context) (domain.Employee, error) {
	return EI.repository.GetByEmail(c)
}

func (EI *EmployeeInteractor) GetAll() ([]domain.Employee, error) {
	return EI.repository.GetAll()
}

func (EI *EmployeeInteractor) Create(c *gin.Context) (domain.Employee, error) {
	return EI.repository.Create(c)
}

func (EI *EmployeeInteractor) Update(c *gin.Context) (domain.Employee, error) {
	return EI.repository.Update(c)
}
