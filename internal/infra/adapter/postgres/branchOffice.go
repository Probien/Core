package adapter

import (
	"encoding/json"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/domain/port/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/adapter"
	"gorm.io/gorm"
	"math"
	"net/url"
)

type BranchOfficeRepositoryImp struct {
	database *gorm.DB
}

func NewBranchOfficeRepositoryImp(db *gorm.DB) port.IBranchOfficeRepository {
	return &BranchOfficeRepositoryImp{database: db}
}

func (r *BranchOfficeRepositoryImp) GetAll(params url.Values) (*[]dto.BranchOffice, map[string]interface{}, error) {
	var branchOffices []dto.BranchOffice
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("branch_offices").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&dto.BranchOffice{}).Scopes(adapter.Paginate(params, paginationResult)).Preload("Employees").Find(&branchOffices).Error; err != nil {
		return nil, nil, adapter.ErrorProcess
	}

	return &branchOffices, paginationResult, nil
}

func (r *BranchOfficeRepositoryImp) GetById(id int) (*dto.BranchOffice, error) {
	var branchOffice dto.BranchOffice

	if err := r.database.Model(&dto.BranchOffice{}).Preload("Employees").Find(&branchOffice, id).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	if branchOffice.ID == 0 {
		return nil, adapter.BranchNotFound
	}
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Create(branchOfficeDto *dto.BranchOffice, userSessionId int) (*dto.BranchOffice, error) {

	if err := r.database.Model(&dto.BranchOffice{}).Omit("Employees").Create(&branchOfficeDto).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	data, _ := json.Marshal(&branchOfficeDto)
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", userSessionId, adapter.SpInsert, adapter.SpNoPrevData, string(data[:]))

	return branchOfficeDto, nil
}

func (r *BranchOfficeRepositoryImp) Update(id int, branchOfficeDto map[string]interface{}, userSessionId int) (*dto.BranchOffice, error) {
	branchOffice, branchOfficeOld := dto.BranchOffice{}, dto.BranchOffice{}

	r.database.Model(&dto.BranchOffice{}).Find(&branchOfficeOld, id)

	if branchOfficeOld.ID == 0 {
		return nil, adapter.BranchNotFound
	}

	if err := r.database.Model(&dto.BranchOffice{}).Where("id = ?", id).Updates(&branchOfficeDto).Find(&branchOffice).Error; err != nil {
		return nil, adapter.ErrorProcess
	}

	old, _ := json.Marshal(&branchOfficeOld)
	current, _ := json.Marshal(&branchOffice)

	go r.database.Exec("CALL savemovement(?,?,?,?)", id, adapter.SpUpdate, string(old[:]), string(current[:]))
	return &branchOffice, nil
}
