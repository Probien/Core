package persistance

import (
	"errors"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BranchOfficeRepositoryImp struct {
	database *gorm.DB
}

func NewBranchOfficeRepositoryImp(db *gorm.DB) repository.IBranchOfficeRepository {
	return &BranchOfficeRepositoryImp{database: db}
}

func (r *BranchOfficeRepositoryImp) GetAll() (*[]domain.BranchOffice, error) {
	var branchOffices []domain.BranchOffice

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffices).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}
	return &branchOffices, nil
}

func (r *BranchOfficeRepositoryImp) GetById(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffice, c.Param("id")).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if branchOffice.ID == 0 {
		return nil, errors.New("branch office not found")
	}
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Create(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := c.ShouldBindJSON(&branchOffice); err != nil {
		return nil, errors.New("error binding JSON data, verify fields")
	}

	if err := r.database.Model(&domain.BranchOffice{}).Omit("Employees").Create(&branchOffice).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Update(c *gin.Context) (*domain.BranchOffice, error) {
	patch, branchOffice := map[string]interface{}{}, domain.BranchOffice{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New("error binding JSON data, verify json format")
	}

	if err := r.database.Model(&domain.BranchOffice{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&branchOffice).Error; err != nil {
		return nil, errors.New("failed to establish a connection with our database services")
	}

	if branchOffice.ID == 0 {
		return nil, errors.New("branch office not found")
	}
	return &branchOffice, nil
}
