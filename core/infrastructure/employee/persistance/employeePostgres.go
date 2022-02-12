package persistance

import (
	"errors"

	employee_domain "github.com/JairDavid/Probien-Backend/core/domain/employee"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/employee/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	database *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) employee_domain.EmployeeRepository {
	return &EmployeeRepositoryImpl{database: db}
}

func (r *EmployeeRepositoryImpl) Login(c *gin.Context) (*employee_domain.Employee, error) {
	employee, loginCredentials := employee_domain.Employee{}, auth.LoginCredentials{}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&employee_domain.Employee{}).Where("email = ?", loginCredentials.Email).Find(&employee)

	if employee.ID == 0 {
		return nil, errors.New("inexistent employee with that email")

	} else if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredentials.Password)); err != nil {
		return nil, errors.New("email or Password incorrect")
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (*employee_domain.Employee, error) {
	var employee employee_domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&employee_domain.Employee{}).Where("email = ?", employee.Email).Find(&employee)

	if employee.ID == 0 {
		return nil, errors.New("employee with that email not found")
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll() (*[]employee_domain.Employee, error) {
	var employees []employee_domain.Employee

	r.database.Model(employee_domain.Employee{}).Find(&employees)
	return &employees, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (*employee_domain.Employee, error) {
	crypt, employee := make(chan []byte, 1), employee_domain.Employee{}

	if err := c.ShouldBindJSON(&employee); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}
	auth.EncryptPassword([]byte(employee.Password), crypt)
	employee.Password = string(<-crypt)

	r.database.Model(&employee_domain.Employee{}).Create(&employee)
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (*employee_domain.Employee, error) {

	patch, employee := map[string]interface{}{}, employee_domain.Employee{}

	if err := c.Bind(&patch); err != nil {
		return nil, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return nil, errors.New("empty request body")
	} else if _, err := patch["email"]; !err {
		return nil, errors.New("to perform this operation it is necessary to enter an email in the JSON body")
	}

	result := r.database.Model(&employee_domain.Employee{}).Where("email = ?", &employee.Email).Omit("id").Updates(&patch).Find(&employee)
	if result.RowsAffected == 0 {
		return nil, errors.New("employee not found or json data does not match ")
	}

	return &employee, nil
}
