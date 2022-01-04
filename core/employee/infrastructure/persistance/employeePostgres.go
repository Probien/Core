package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/employee/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	database *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) domain.EmployeeRepository {
	return &EmployeeRepositoryImpl{database: db}
}

func (r *EmployeeRepositoryImpl) Login(c *gin.Context) (domain.Employee, bool) {
	return domain.Employee{}, true
}

func (r *EmployeeRepositoryImpl) GetByEmail(c *gin.Context) (domain.Employee, error) {
	return domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) GetAll() ([]domain.Employee, error) {
	return []domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) Create(c *gin.Context) (domain.Employee, error) {
	return domain.Employee{}, nil
}

func (r *EmployeeRepositoryImpl) Update(c *gin.Context) (domain.Employee, error) {
	return domain.Employee{}, nil
}
