package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	database *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) repository.IEmployeeRepository {
	return &EmployeeRepositoryImpl{database: db}
}

func (r *EmployeeRepositoryImpl) Login(c *gin.Context) (*domain.Employee, error) {
	employee, loginCredentials := domain.Employee{}, auth.LoginCredentials{}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", loginCredentials.Email).Find(&employee).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if employee.ID == 0 {
		return nil, errors.New("inexistent employee with that email")

	} else if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredentials.Password)); err != nil {
		return nil, errors.New("incorrect credentials")
	}

	go r.database.Exec("CALL savesession(?)", employee.ID)

	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (*domain.Employee, error) {
	var employee domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", employee.Email).Preload("Profile").Find(&employee).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if employee.ID == 0 {
		return nil, errors.New("employee with that email not found")
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll() (*[]domain.Employee, error) {
	var employees []domain.Employee

	if err := r.database.Model(domain.Employee{}).Preload("Profile").Preload("PawnOrdersDone").Preload("Endorsements").Find(&employees).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &employees, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (*domain.Employee, error) {
	crypt, employee := make(chan []byte, 1), domain.Employee{}

	if err := c.ShouldBindJSON(&employee); err != nil || employee.BranchOfficeID == 0 {
		return nil, errors.New("error binding JSON data, verify fields")
	}
	auth.EncryptPassword([]byte(employee.Password), crypt)
	employee.Password = string(<-crypt)

	if err := r.database.Model(&domain.Employee{}).Omit("PawnOrdersDone").Omit("SessionLogs").Omit("EndorsementsDone").Create(&employee).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	return &employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (*domain.Employee, error) {
	patch, employee := map[string]interface{}{}, domain.Employee{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New("error binding JSON data, verify json format")
	}

	if err := r.database.Model(&domain.Employee{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&employee).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if employee.ID == 0 {
		return nil, errors.New("employee not found or json data does not match ")
	}

	return &employee, nil
}
