package postgres

import (
	"encoding/json"
	"math"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BranchOfficeRepositoryImp struct {
	database *gorm.DB
}

func NewBranchOfficeRepositoryImp(db *gorm.DB) repository.IBranchOfficeRepository {
	return &BranchOfficeRepositoryImp{database: db}
}

func (r *BranchOfficeRepositoryImp) GetAll(c *gin.Context) (*[]domain.BranchOffice, map[string]interface{}, error) {
	var branchOffices []domain.BranchOffice
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("branch_offices").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.BranchOffice{}).Scopes(persistence.Paginate(c, paginationResult)).Preload("Employees").Find(&branchOffices).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &branchOffices, paginationResult, nil
}

func (r *BranchOfficeRepositoryImp) GetById(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffice, c.Param("id")).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if branchOffice.ID == 0 {
		return nil, persistence.BranchNotFound
	}
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Create(c *gin.Context) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := c.ShouldBindJSON(&branchOffice); err != nil {
		return nil, persistence.ErrorBinding
	}

	if err := r.database.Model(&domain.BranchOffice{}).Omit("Employees").Create(&branchOffice).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&branchOffice)
	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", contextUserID.(int), persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Update(c *gin.Context) (*domain.BranchOffice, error) {
	patch, branchOffice, branchOfficeOld := map[string]interface{}{}, domain.BranchOffice{}, domain.BranchOffice{}

	if err := c.Bind(&patch); err != nil {
		return nil, err
	}

	_, errID := patch["id"]

	if !errID {
		return nil, persistence.ErrorBinding
	}

	r.database.Model(&domain.BranchOffice{}).Find(&branchOfficeOld, patch["id"])

	if err := r.database.Model(&domain.BranchOffice{}).Where("id = ?", patch["id"]).Updates(&patch).Find(&branchOffice).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if branchOffice.ID == 0 {
		return nil, persistence.BranchNotFound
	}

	old, _ := json.Marshal(&branchOfficeOld)
	current, _ := json.Marshal(&branchOffice)

	contextUserID, _ := c.Get("user_id")
	//context user id, is the userID comming from jwt decoded
	go r.database.Exec("CALL savemovement(?,?,?,?)", contextUserID.(int), persistence.SpUpdate, string(old[:]), string(current[:]))
	return &branchOffice, nil
}
