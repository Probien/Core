package persistance

import (
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BranchOfficeRepositoryImp struct {
	database *gorm.DB
}

func NewBranchOfficeRepositoryImp(db *gorm.DB) repository.BranchOfficeRepository {
	return &BranchOfficeRepositoryImp{database: db}
}

func (r *BranchOfficeRepositoryImp) GetAll() (*[]domain.BranchOffice, error) {
	return &[]domain.BranchOffice{}, nil
}

func (r *BranchOfficeRepositoryImp) GetById(c *gin.Context) (*domain.BranchOffice, error) {
	return &domain.BranchOffice{}, nil
}

func (r *BranchOfficeRepositoryImp) Create(c *gin.Context) (*domain.BranchOffice, error) {
	return &domain.BranchOffice{}, nil
}

func (r *BranchOfficeRepositoryImp) Update(c *gin.Context) (*domain.BranchOffice, error) {
	return &domain.BranchOffice{}, nil
}
