package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/customer/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	database *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) domain.CustomerRepository {
	return &CustomerRepositoryImpl{database: db}
}

func (r *CustomerRepositoryImpl) GetById(c *gin.Context) (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (r *CustomerRepositoryImpl) GetAll() ([]domain.Customer, error) {
	return []domain.Customer{}, nil
}

func (r *CustomerRepositoryImpl) Create(c *gin.Context) (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (r *CustomerRepositoryImpl) Update(c *gin.Context) (domain.Customer, error) {
	return domain.Customer{}, nil
}
