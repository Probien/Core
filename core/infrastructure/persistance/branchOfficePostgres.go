package persistance

import (
	"encoding/json"
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
		return nil, errors.New(ERROR_PROCCESS)
	}
	return &branchOffices, nil
}

func (r *BranchOfficeRepositoryImp) GetById(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffice, c.Param("id")).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if branchOffice.ID == 0 {
		return nil, errors.New(BRANCH_NOT_FOUND)
	}
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Create(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := c.ShouldBindJSON(&branchOffice); err != nil {
		return nil, errors.New(ERROR_BINDING)
	}

	if err := r.database.Model(&domain.BranchOffice{}).Omit("Employees").Create(&branchOffice).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	data, _ := json.Marshal(&branchOffice)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", contextUserID.(uint), SP_INSERT, SP_NO_PREV_DATA, string(data[:]))
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Update(c *gin.Context) (*domain.BranchOffice, error) {
	patch, branchOffice, branchOfficeOld := map[string]interface{}{}, domain.BranchOffice{}, domain.BranchOffice{}
	_, errID := patch["id"]

	if err := c.Bind(&patch); err != nil && !errID {
		return nil, errors.New(ERROR_BINDING)
	}

	r.database.Model(&domain.BranchOffice{}).Find(&branchOfficeOld, patch["id"])

	if err := r.database.Model(&domain.BranchOffice{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&branchOffice).Error; err != nil {
		return nil, errors.New(ERROR_PROCCESS)
	}

	if branchOffice.ID == 0 {
		return nil, errors.New(BRANCH_NOT_FOUND)
	}

	old, _ := json.Marshal(&branchOfficeOld)
	new, _ := json.Marshal(&branchOffice)

	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), SP_UPDATE, string(old[:]), string(new[:]))
	return &branchOffice, nil
}
