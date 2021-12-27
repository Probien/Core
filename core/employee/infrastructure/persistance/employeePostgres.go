package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/employee/domain"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	database *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) domain.EmployeeRepository {
	return &EmployeeRepositoryImpl{database: db}
}

func (r *EmployeeRepositoryImpl) GetById() (domain.Employee, error) {
	return domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) GetAll() ([]domain.Employee, error) {
	return []domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) Create() (domain.Employee, error) {
	return domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) Update() (domain.Employee, error) {
	return domain.Employee{}, nil
}
