package postgres

import (
	"encoding/json"
	"math"

	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

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
		return nil, persistence.ErrorBinding
	}

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", loginCredentials.Email).Preload("Profile").Preload("Roles.Role").Find(&employee).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if employee.ID == 0 {
		return nil, persistence.EmployeeNotFound

	} else if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredentials.Password)); err != nil {
		return nil, persistence.InvalidCredentials
	}

	go r.database.Exec("CALL savesession(?)", employee.ID)

	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (*domain.Employee, error) {
	body, employee := map[string]interface{}{}, domain.Employee{}

	if err := c.ShouldBindJSON(&body); err != nil {
		return nil, err
	}

	_, email := body["email"]

	if !email {
		return nil, persistence.ErrorBinding
	}

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", body["email"]).Preload("Profile").Preload("Roles.Role").Find(&employee).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if employee.ID == 0 {
		return nil, persistence.EmployeeNotFound
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll(c *gin.Context) (*[]domain.Employee, map[string]interface{}, error) {
	var employees []domain.Employee
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("employees").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(domain.Employee{}).Scopes(persistence.Paginate(c, paginationResult)).Preload("Profile").Preload("Roles.Role").Find(&employees).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &employees, paginationResult, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (*domain.Employee, error) {
	crypt, employee := make(chan []byte, 1), domain.Employee{}

	if err := c.ShouldBindJSON(&employee); err != nil {
		return nil, persistence.ErrorBinding
	}

	go auth.EncryptPassword([]byte(employee.Password), crypt)
	employee.Password = string(<-crypt)

	if err := r.database.Model(&domain.Employee{}).Omit("PawnOrdersDone").Omit("SessionLogs").Omit("EndorsementsDone").Omit("Roles").Create(&employee).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	for _, v := range employee.Roles {
		r.database.Exec("INSERT INTO employee_roles(role_id, employee_id) VALUES(?,?)", v.RoleID, employee.ID)
	}

	r.database.Model(&employee).Preload("Roles.Role").Find(&employee)

	data, _ := json.Marshal(&employee)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (*domain.Employee, error) {
	patch, employee, employeeOld := map[string]interface{}{}, domain.Employee{}, domain.Employee{}

	if err := c.Bind(&patch); err != nil {
		return nil, err
	}

	_, errID := patch["id"]

	if !errID {
		return nil, persistence.ErrorBinding
	}

	r.database.Model(&domain.Employee{}).Find(&employeeOld, patch["id"])

	if err := r.database.Model(&domain.Employee{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&employee).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if employee.ID == 0 {
		return nil, persistence.EmployeeNotFound
	}

	old, _ := json.Marshal(&employeeOld)
	current, _ := json.Marshal(&employee)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpUpdate, string(old[:]), string(current[:]))
	return &employee, nil
}
