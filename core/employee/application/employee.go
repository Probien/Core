package application

import (
	"github.com/JairDavid/Probien-Backend/core/employee/domain"
	"github.com/gin-gonic/gin"
)

type EmployeeInteractor struct {
	repository domain.EmployeeRepository
}

func (EI *EmployeeInteractor) GetById(c *gin.Context) (domain.Employee, error) {
	return EI.repository.GetByEmail(c)
}

func (EI *EmployeeInteractor) GetAll() ([]domain.Employee, error) {
	return EI.repository.GetAll()
}

func (EI *EmployeeInteractor) Create(c *gin.Context) (domain.Employee, error) {
	return EI.repository.Create(c)
}

func (EI *EmployeeInteractor) Update(c *gin.Context) (domain.Employee, error) {
	return EI.repository.Update(c)
}
