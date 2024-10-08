package adapter

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/adapter"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"math"
	"net/url"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	database *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) port.IEmployeeRepository {
	return &EmployeeRepositoryImpl{database: db}
}

func (r *EmployeeRepositoryImpl) Login(loginCredential component.Credential) (*dto.Employee, error) {
	var employee dto.Employee

	if err := r.database.Model(&dto.Employee{}).Where("email = ?", loginCredential.Email).Preload("Profile").Preload("Roles.Role").Find(&employee).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	if employee.ID == 0 {
		return nil, adapter.EmployeeNotFound

	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginCredential.Password)); err != nil {
		return nil, adapter.InvalidCredentials
	}

	go r.database.Exec("CALL savesession(?)", employee.ID)

	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetByEmail(email string) (*dto.Employee, error) {
	var employee dto.Employee

	if err := r.database.Model(&dto.Employee{}).Where("email = ?", email).Preload("Profile").Preload("Roles.Role").Find(&employee).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	if employee.ID == 0 {
		return nil, adapter.EmployeeNotFound
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAll(params url.Values) (*[]dto.Employee, map[string]interface{}, error) {
	var employees []dto.Employee
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("employees").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(dto.Employee{}).Scopes(adapter.Paginate(params, paginationResult)).Preload("Profile").Preload("Roles.Role").Find(&employees).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &employees, paginationResult, nil
}

func (r *EmployeeRepositoryImpl) Create(employeeDto *dto.Employee, userSessionId int) (*dto.Employee, error) {
	crypt := make(chan []byte, 1)
	auth := component.NewAuthenticator()

	go auth.EncryptPassword([]byte(employeeDto.Password), crypt)
	employeeDto.Password = string(<-crypt)

	if err := r.database.Model(&dto.Employee{}).Omit("PawnOrdersDone").Omit("SessionLogs").Omit("EndorsementsDone").Omit("Roles").Create(&employeeDto).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	for _, v := range employeeDto.Roles {
		r.database.Exec("INSERT INTO employee_roles(role_id, employee_id) VALUES(?,?)", v.RoleID, employeeDto.ID)
	}

	r.database.Model(&employeeDto).Preload("Roles.Role").Find(&employeeDto)

	data, _ := json.Marshal(&employeeDto)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpInsert, adapter.SpNoPrevData, string(data[:]))
	return employeeDto, nil
}

func (r *EmployeeRepositoryImpl) Update(id int, employeeDto map[string]interface{}, userSessionId int) (*dto.Employee, error) {
	employee, employeeOld := dto.Employee{}, dto.Employee{}

	r.database.Model(&dto.Employee{}).Find(&employeeOld, id)

	if employeeOld.ID == 0 {
		return nil, adapter.EmployeeNotFound
	}

	if err := r.database.Model(&dto.Employee{}).Where("id = ?", id).Updates(&employeeDto).Find(&employee).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	old, _ := json.Marshal(&employeeOld)
	current, _ := json.Marshal(&employee)

	go r.database.Exec("CALL savemovement(?,?,?,?)", userSessionId, adapter.SpUpdate, string(old[:]), string(current[:]))
	return &employee, nil
}
