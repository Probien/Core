package postgres

import (
	"encoding/json"
	"math"
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	database *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) repository.IEmployeeRepository {
	return &EmployeeRepositoryImpl{database: db}
}

func (r *EmployeeRepositoryImpl) Login(loginCredential auth.LoginCredential) (*domain.Employee, error) {
	var employee domain.Employee

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", loginCredential.Email).Preload("Profile").Preload("Roles.Role").Find(&employee).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if employee.ID == 0 {
		return nil, persistence.EmployeeNotFound

	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredential.Password)); err != nil {
		return nil, persistence.InvalidCredentials
	}

	go r.database.Exec("CALL savesession(?)", employee.ID)

	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(email string) (*domain.Employee, error) {
	var employee domain.Employee

	if err := r.database.Model(&domain.Employee{}).Where("email = ?", email).Preload("Profile").Preload("Roles.Role").Find(&employee).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if employee.ID == 0 {
		return nil, persistence.EmployeeNotFound
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll(params url.Values) (*[]domain.Employee, map[string]interface{}, error) {
	var employees []domain.Employee
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("employees").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(domain.Employee{}).Scopes(persistence.Paginate(params, paginationResult)).Preload("Profile").Preload("Roles.Role").Find(&employees).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &employees, paginationResult, nil
}

func (r *EmployeeRepositoryImpl) Create(employeeDto *domain.Employee, userSessionId int) (*domain.Employee, error) {
	crypt := make(chan []byte, 1)

	go auth.EncryptPassword([]byte(employeeDto.Password), crypt)
	employeeDto.Password = string(<-crypt)

	if err := r.database.Model(&domain.Employee{}).Omit("PawnOrdersDone").Omit("SessionLogs").Omit("EndorsementsDone").Omit("Roles").Create(&employeeDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	for _, v := range employeeDto.Roles {
		r.database.Exec("INSERT INTO employee_roles(role_id, employee_id) VALUES(?,?)", v.RoleID, employeeDto.ID)
	}

	r.database.Model(&employeeDto).Preload("Roles.Role").Find(&employeeDto)

	data, _ := json.Marshal(&employeeDto)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return employeeDto, nil
}

func (r *EmployeeRepositoryImpl) Update(id int, employeeDto map[string]interface{}, userSessionId int) (*domain.Employee, error) {
	employee, employeeOld := domain.Employee{}, domain.Employee{}

	r.database.Model(&domain.Employee{}).Find(&employeeOld, id)

	if employeeOld.ID == 0 {
		return nil, persistence.EmployeeNotFound
	}

	if err := r.database.Model(&domain.Employee{}).Where("id = ?", id).Updates(&employeeDto).Find(&employee).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	old, _ := json.Marshal(&employeeOld)
	current, _ := json.Marshal(&employee)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, persistence.SpUpdate, string(old[:]), string(current[:]))
	return &employee, nil
}
