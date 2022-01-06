package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/employee/domain"
	"github.com/JairDavid/Probien-Backend/core/employee/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	database *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) domain.EmployeeRepository {
	return &EmployeeRepositoryImpl{database: db}
}

func (r *EmployeeRepositoryImpl) Login(c *gin.Context) (domain.Employee, error) {
	crypt := make(chan []byte, 1)
	var employee domain.Employee
	var loginCredentials auth.LoginCredentials

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		return domain.Employee{}, errors.New("error binding JSON data")
	}

	r.database.Model(domain.Employee{}).Where("email = ?", loginCredentials.Email).Find(&employee)

	go auth.EncryptPassword([]byte(loginCredentials.Password), crypt)

	if employee == (domain.Employee{}) {
		return domain.Employee{}, errors.New("inexistent employee with that email")

		//fix this :D, watch COmpareHashAndPassword data type return
	} else if err := bcrypt.CompareHashAndPassword(<-crypt, []byte(employee.Password)); err != nil {
		return employee, errors.New("email or Password incorrect")
	}
	return employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (domain.Employee, error) {
	return domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) GetAll() ([]domain.Employee, error) {
	return []domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (domain.Employee, error) {
	crypt := make(chan []byte, 1)
	var employee domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		return domain.Employee{}, errors.New("error binding JSON data")
	}
	auth.EncryptPassword([]byte(employee.Password), crypt)
	employee.Password = string(<-crypt)

	r.database.Model(&domain.Employee{}).Create(&employee)
	return employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (domain.Employee, error) {
	return domain.Employee{}, nil
}
