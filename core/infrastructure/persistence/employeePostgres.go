package persistence

import (
	"encoding/json"

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
	employee, loginCredentials := domain.Employee{}, auth.LoginCredential{}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		return nil, ErrorBinding
	}

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", loginCredentials.Email).Preload("Profile").Preload("Roles.Role").Find(&employee).Error; err != nil {
		return nil, ErrorProcess
	}

	if employee.ID == 0 {
		return nil, EmployeeNotFound

	} else if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredentials.Password)); err != nil {
		return nil, InvalidCredentials
	}

	go r.database.Exec("CALL savesession(?)", employee.ID)

	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (*domain.Employee, error) {
	var employee domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		return nil, ErrorBinding
	}

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", employee.Email).Preload("Profile").Preload("Roles.Role").Preload("PawnOrdersDone.Customer").Preload("SessionLogs").Preload("Endorsements").Find(&employee).Error; err != nil {
		return nil, ErrorProcess
	}

	if employee.ID == 0 {
		return nil, EmployeeNotFound
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll() (*[]domain.Employee, error) {
	var employees []domain.Employee

	if err := r.database.Model(domain.Employee{}).Preload("Profile").Preload("Roles.Role").Find(&employees).Error; err != nil {
		return nil, ErrorProcess
	}

	return &employees, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (*domain.Employee, error) {
	crypt, employee := make(chan []byte, 1), domain.Employee{}

	if err := c.ShouldBindJSON(&employee); err != nil || employee.BranchOfficeID == 0 {
		return nil, ErrorBinding
	}

	auth.EncryptPassword([]byte(employee.Password), crypt)
	employee.Password = string(<-crypt)

	if err := r.database.Model(&domain.Employee{}).Omit("PawnOrdersDone").Omit("SessionLogs").Omit("EndorsementsDone").Omit("Roles").Create(&employee).Error; err != nil {
		return nil, ErrorProcess
	}

	for _, v := range employee.Roles {
		r.database.Exec("INSERT INTO employee_roles(role_id, employee_id) VALUES(?,?)", v.RoleID, employee.ID)
	}

	r.database.Model(&employee).Preload("Roles.Role").Find(&employee)

	data, _ := json.Marshal(&employee)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpInsert, SpNoPrevData, string(data[:]))
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (*domain.Employee, error) {
	patch, employee, employeeOld := map[string]interface{}{}, domain.Employee{}, domain.Employee{}

	if err := c.Bind(&patch); err != nil {
		return nil, err
	}

	_, errID := patch["id"]

	if !errID {
		return nil, ErrorBinding
	}

	r.database.Model(&domain.Employee{}).Find(&employeeOld, patch["id"])

	if err := r.database.Model(&domain.Employee{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&employee).Error; err != nil {
		return nil, ErrorProcess
	}

	if employee.ID == 0 {
		return nil, EmployeeNotFound
	}

	old, _ := json.Marshal(&employeeOld)
	current, _ := json.Marshal(&employee)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpUpdate, string(old[:]), string(current[:]))
	return &employee, nil
}
