package persistance

import (
	"errors"
	"fmt"

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
	var employee domain.Employee
	var loginCredentials auth.LoginCredentials

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		return domain.Employee{}, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.Employee{}).Where("email = ?", loginCredentials.Email).Find(&employee)

	if employee == (domain.Employee{}) {
		return domain.Employee{}, errors.New("inexistent employee with that email")

	} else if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredentials.Password)); err != nil {
		return employee, errors.New("email or Password incorrect")
	}
	return employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (domain.Employee, error) {
	var employee domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		return domain.Employee{}, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.Employee{}).Where("email = ?", employee.Email).Find(&employee)

	if employee == (domain.Employee{}) {
		return domain.Employee{}, errors.New("inexistent employee with that email")
	}
	return employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll() ([]domain.Employee, error) {
	var employees []domain.Employee

	r.database.Model(domain.Employee{}).Find(&employees)
	return employees, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (domain.Employee, error) {
	crypt := make(chan []byte, 1)
	var employee domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		return domain.Employee{}, errors.New("error binding JSON data, verify fields")
	}
	auth.EncryptPassword([]byte(employee.Password), crypt)
	employee.Password = string(<-crypt)

	r.database.Model(&domain.Employee{}).Create(&employee)
	return employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (domain.Employee, error) {
	//PENDING: consider all error cases
	var employee domain.Employee
	var patch map[string]interface{}

	if err := c.Bind(&patch); err != nil {
		return domain.Employee{}, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return domain.Employee{}, errors.New("empty request body")
	} else if _, err := patch["email"]; !err {
		return domain.Employee{}, errors.New("to perform this operation it is necessary to enter an email in the JSON body")
	}

	r.database.Model(&domain.Employee{}).Where("email = ?", patch["email"]).Find(&employee)
	if employee.ID == 0 {
		return domain.Employee{}, errors.New("employee not found")
	}
	fmt.Print(employee)
	result := r.database.Model(&domain.Employee{}).Omit("password").Where("email = ?", &employee.Email).Updates(&patch)
	if result.RowsAffected == 0 {
		return domain.Employee{}, errors.New("json data does not match with the database entity")
	}

	return employee, nil
}
