package postgres

import (
	"encoding/json"
	"math"
	"net/url"

	"github.com/JairDavid/Probien-Backend/pkg/domain"
	"github.com/JairDavid/Probien-Backend/pkg/domain/repository"
	"github.com/JairDavid/Probien-Backend/pkg/infrastructure/persistence"
	"gorm.io/gorm"
)

type BranchOfficeRepositoryImp struct {
	database *gorm.DB
}

func NewBranchOfficeRepositoryImp(db *gorm.DB) repository.IBranchOfficeRepository {
	return &BranchOfficeRepositoryImp{database: db}
}

func (r *BranchOfficeRepositoryImp) GetAll(params url.Values) (*[]domain.BranchOffice, map[string]interface{}, error) {
	var branchOffices []domain.BranchOffice
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("branch_offices").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.BranchOffice{}).Scopes(persistence.Paginate(params, paginationResult)).Preload("Employees").Find(&branchOffices).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &branchOffices, paginationResult, nil
}

func (r *BranchOfficeRepositoryImp) GetById(id int) (*domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice

	if err := r.database.Model(&domain.BranchOffice{}).Preload("Employees").Find(&branchOffice, id).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	if branchOffice.ID == 0 {
		return nil, persistence.BranchNotFound
	}
	return &branchOffice, nil
}

func (r *BranchOfficeRepositoryImp) Create(branchOfficeDto *domain.BranchOffice, userSessionId int) (*domain.BranchOffice, error) {

	if err := r.database.Model(&domain.BranchOffice{}).Omit("Employees").Create(&branchOfficeDto).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	data, _ := json.Marshal(&branchOfficeDto)
	go r.database.Exec("CALL savemovement(?, ?, ?, ?)", userSessionId, persistence.SpInsert, persistence.SpNoPrevData, string(data[:]))

	return branchOfficeDto, nil
}

func (r *BranchOfficeRepositoryImp) Update(id int, branchOfficeDto map[string]interface{}, userSessionId int) (*domain.BranchOffice, error) {
	branchOffice, branchOfficeOld := domain.BranchOffice{}, domain.BranchOffice{}

	r.database.Model(&domain.BranchOffice{}).Find(&branchOfficeOld, id)

	if branchOfficeOld.ID == 0 {
		return nil, persistence.BranchNotFound
	}

	if err := r.database.Model(&domain.BranchOffice{}).Where("id = ?", id).Updates(&branchOfficeDto).Find(&branchOffice).Error; err != nil {
		return nil, persistence.ErrorProcess
	}

	old, _ := json.Marshal(&branchOfficeOld)
	current, _ := json.Marshal(&branchOffice)

	go r.database.Exec("CALL savemovement(?,?,?,?)", id, persistence.SpUpdate, string(old[:]), string(current[:]))
	return &branchOffice, nil
}
