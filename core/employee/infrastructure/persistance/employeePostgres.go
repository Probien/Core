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

func (r *EmployeeRepositoryImpl) Login(c *gin.Context) (*domain.Employee, error) {
	employee, loginCredentials := domain.Employee{}, auth.LoginCredentials{}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		return &domain.Employee{}, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.Employee{}).Where("email = ?", loginCredentials.Email).Find(&employee)

	if employee.ID == 0 {
		return &domain.Employee{}, errors.New("inexistent employee with that email")

	} else if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredentials.Password)); err != nil {
		return &domain.Employee{}, errors.New("email or Password incorrect")
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (*domain.Employee, error) {
	var employee domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		return &domain.Employee{}, errors.New("error binding JSON data, verify fields")
	}

	r.database.Model(&domain.Employee{}).Where("email = ?", employee.Email).Find(&employee)

	if employee.ID == 0 {
		return &domain.Employee{}, errors.New("employee with that email not found")
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll() (*[]domain.Employee, error) {
	var employees []domain.Employee

	r.database.Model(domain.Employee{}).Find(&employees)
	return &employees, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (*domain.Employee, error) {
	crypt, employee := make(chan []byte, 1), domain.Employee{}

	if err := c.ShouldBindJSON(&employee); err != nil {
		return &domain.Employee{}, errors.New("error binding JSON data, verify fields")
	}
	auth.EncryptPassword([]byte(employee.Password), crypt)
	employee.Password = string(<-crypt)

	r.database.Model(&domain.Employee{}).Create(&employee)
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (*domain.Employee, error) {

	patch, employee := map[string]interface{}{}, domain.Employee{}

	if err := c.Bind(&patch); err != nil {
		return &domain.Employee{}, errors.New("error binding JSON data")
	} else if len(patch) == 0 {
		return &domain.Employee{}, errors.New("empty request body")
	} else if _, err := patch["email"]; !err {
		return &domain.Employee{}, errors.New("to perform this operation it is necessary to enter an email in the JSON body")
	}

	result := r.database.Model(&domain.Employee{}).Where("email = ?", &employee.Email).Omit("id").Updates(&patch).Find(&employee)
	if result.RowsAffected == 0 {
		return &domain.Employee{}, errors.New("employee not found or json data does not match ")
	}

	return &employee, nil
}
