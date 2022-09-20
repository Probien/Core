package postgres

import (
	"math"
	"net/url"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"
	"gorm.io/gorm"
)

type LogsRepositoryImp struct {
	database *gorm.DB
}

func NewLogsRepositoryImp(db *gorm.DB) repository.IlogsRepository {
	return &LogsRepositoryImp{database: db}
}

func (r *LogsRepositoryImp) GetAllSessions(params url.Values) (*[]domain.SessionLog, map[string]interface{}, error) {
	var sessions []domain.SessionLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("session_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.SessionLog{}).Scopes(persistence.Paginate(c, paginationResult)).Preload("Employee").Find(&sessions).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &sessions, paginationResult, nil
}

func (r *LogsRepositoryImp) GetAllSessionsByEmployeeId(id int) (*[]domain.SessionLog, map[string]interface{}, error) {
	var sessions []domain.SessionLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("session_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.SessionLog{}).Scopes(persistence.Paginate(c, paginationResult)).Where("employee_id = ?", c.Param("id")).Preload("Employee").Find(&sessions).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &sessions, paginationResult, nil
}

func (r *LogsRepositoryImp) GetAllMovements(params url.Values) (*[]domain.ModerationLog, map[string]interface{}, error) {
	var movements []domain.ModerationLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	r.database.Table("moderation_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.ModerationLog{}).Scopes(persistence.Paginate(c, paginationResult)).Find(&movements).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &movements, paginationResult, nil
}

func (r *LogsRepositoryImp) GetAllMovementsByEmployeeId(id int) (*[]domain.ModerationLog, map[string]interface{}, error) {
	var movements []domain.ModerationLog
	var totalRows int64
	paginationResult := map[string]interface{}{}

	go r.database.Table("moderation_logs").Count(&totalRows)
	paginationResult["total_pages"] = math.Floor(float64(totalRows) / 10)

	if err := r.database.Model(&domain.ModerationLog{}).Scopes(persistence.Paginate(c, paginationResult)).Where("user_id", c.Param("id")).Find(&movements).Error; err != nil {
		return nil, nil, persistence.ErrorProcess
	}

	return &movements, paginationResult, nil
}
