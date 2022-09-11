package persistence

import (
	"encoding/json"

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

func (r *BranchOfficeRepositoryImp) GetAll(c *gin.Context) (*[]domain.BranchOffice, error) {
	var branchOffices []domain.BranchOffice

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffices).Error; err != nil {
		return nil, ErrorProcess
	}
	return &branchOffices, nil
}

func (r *BranchOfficeRepositoryImp) GetById(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffice, c.Param("id")).Error; err != nil {
		return nil, ErrorProcess
	}

	if branchOffice.ID == 0 {
		return nil, BranchNotFound
	}
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Create(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := c.ShouldBindJSON(&branchOffice); err != nil {
		return nil, ErrorBinding
	}

	if err := r.database.Model(&domain.BranchOffice{}).Omit("Employees").Create(&branchOffice).Error; err != nil {
		return nil, ErrorProcess
	}

	data, _ := json.Marshal(&branchOffice)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", contextUserID.(int), SpInsert, SpNoPrevData, string(data[:]))
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Update(c *gin.Context) (*domain.BranchOffice, error) {
	patch, branchOffice, branchOfficeOld := map[string]interface{}{}, domain.BranchOffice{}, domain.BranchOffice{}

	if err := c.Bind(&patch); err != nil {
		return nil, err
	}

	_, errID := patch["id"]

	if !errID {
		return nil, ErrorBinding
	}

	r.database.Model(&domain.BranchOffice{}).Find(&branchOfficeOld, patch["id"])

	if err := r.database.Model(&domain.BranchOffice{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&branchOffice).Error; err != nil {
		return nil, ErrorProcess
	}

	if branchOffice.ID == 0 {
		return nil, BranchNotFound
	}

	old, _ := json.Marshal(&branchOfficeOld)
	current, _ := json.Marshal(&branchOffice)

	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SpUpdate, string(old[:]), string(current[:]))
	return &branchOffice, nil
}
