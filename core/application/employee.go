package application

import (
	"strconv"
	"time"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistance"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type EmployeeInteractor struct {
}

func (EI *EmployeeInteractor) GenerateToken(employee *domain.Employee, tokenizer chan<- string) {
	claims := &auth.AuthCustomClaims{
		Name:      employee.Profile.Name,
		IsAdmin:   employee.IsAdmin,
		CreatedAt: time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			Issuer:    "Probien",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(int(employee.ID)),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte("DPzN3tMBaKsAPxvq8hWfaBHu5oeoj4bioNMQ6NzBSifkTthYAcoM67NzWTaZbPSDhGTkZhsdxyvYmNALanSoa3MH8CBW6Auv"))
	if err != nil {
		panic(err)
	}
	tokenizer <- token
}

func (EI *EmployeeInteractor) Login(c *gin.Context) (*domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.Database)
	return repository.Login(c)
}

func (EI *EmployeeInteractor) GetByEmail(c *gin.Context) (*domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.Database)
	return repository.GetByEmail(c)
}

func (EI *EmployeeInteractor) GetAll() (*[]domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.Database)
	return repository.GetAll()
}

func (EI *EmployeeInteractor) Create(c *gin.Context) (*domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.Database)
	return repository.Create(c)
}

func (EI *EmployeeInteractor) Update(c *gin.Context) (*domain.Employee, error) {
	repository := persistance.NewEmployeeRepositoryImpl(config.Database)
	return repository.Update(c)
}
