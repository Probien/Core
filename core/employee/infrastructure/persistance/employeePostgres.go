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

	var employee, clientEmployeeData domain.Employee
	c.ShouldBindJSON(&clientEmployeeData)
	r.database.Model(domain.Employee{}).Where("email = ?", clientEmployeeData).Find(&employee)

	if employee == (domain.Employee{}) {
		return domain.Employee{}, errors.New("inexistent employee with that email")

	} else if err := bcrypt.CompareHashAndPassword(auth.EncryptPassword([]byte(clientEmployeeData.Password)), []byte(employee.Password)); err != nil {
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
	var employee domain.Employee
	err := c.ShouldBindJSON(&employee)

	employee.Password = string(auth.EncryptPassword([]byte(employee.Password)))
	if err != nil {
		return domain.Employee{}, errors.New("error binding JSON data")
	}
	r.database.Model(&domain.Employee{}).Create(&employee)
	return employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (domain.Employee, error) {
	return domain.Employee{}, nil
}
