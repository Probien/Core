package postgres

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type BranchOfficeRepositoryImp struct {
	database *gorm.DB
}

func NewBranchOfficeRepositoryImp(db *gorm.DB) repository.IBranchOfficeRepository {
	return &BranchOfficeRepositoryImp{database: db}
}

func (r *BranchOfficeRepositoryImp) GetAll(c *gin.Context) (*[]domain.BranchOffice, error) {
	var branchOffices []domain.BranchOffice
	var totalRows int64
	paginationResult := map[string]interface{}{}

	go r.database.Table("branch_offices").Count(&totalRows)
	if len(c.Query("page")) == 0 || c.Query("page") == "0" {
		paginationResult["page"] = 1
	}

	paginationResult["page"], _ = strconv.Atoi(c.Query("page"))
	paginationResult["total_pages"] = math.Ceil(float64(totalRows / 10))
	paginationResult["previous"] = paginationResult["page"].(int) - 1
	paginationResult["next"] = paginationResult["page"].(int) + 1

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffices).Limit(10).Offset(paginationResult["page"].(int) * 10).Error; err != nil {
		return nil, persistence.ErrorProcess
	}
	return &branchOffices, nil
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
