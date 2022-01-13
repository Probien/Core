package application

import (
	"time"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/employee/domain"
	"github.com/JairDavid/Probien-Backend/core/employee/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/employee/infrastructure/persistance"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type EmployeeInteractor struct {
}

func (EI *EmployeeInteractor) GenerateToken(data *domain.Employee, tokenizer chan<- string) {

	claims := &auth.AuthCustomClaims{
		Name:       data.Name,
		FirstName:  data.FirstName,
		SecondName: data.SecondName,
		Email:      data.Email,
		IsAdmin:    data.IsAdmin,
		CreatedAt:  data.CreatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    "Probien",
			Subject:   data.Email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte("DPzN3tMBaKsAPxvq8hWfaBHu5oeoj4bioNMQ6NzBSifkTthYAcoM67NzWTaZbPSDhGTkZhsdxyvYmNALanSoa3MH8CBW6Auv"))
	if err != nil {
		panic(err)
	}
	tokenizer <- token
}

func (EI *EmployeeInteractor) Login(c *gin.Context) (domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.GetDBInstance())
	return repository.Login(c)
}

func (EI *EmployeeInteractor) GetByEmail(c *gin.Context) (domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.GetDBInstance())
	return repository.GetByEmail(c)
}

func (EI *EmployeeInteractor) GetAll() ([]domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.GetDBInstance())
	return repository.GetAll()
}

func (EI *EmployeeInteractor) Create(c *gin.Context) (domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.GetDBInstance())
	return repository.Create(c)
}

func (EI *EmployeeInteractor) Update(c *gin.Context) (domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.GetDBInstance())
	return repository.Update(c)
}
