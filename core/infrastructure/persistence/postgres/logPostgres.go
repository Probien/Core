package postgres

import (
	"math"

	"github.com/JairDavid/Probien-Backend/core/domain"
	"github.com/JairDavid/Probien-Backend/core/domain/repository"
	"github.com/JairDavid/Probien-Backend/core/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LogsRepositoryImp struct {
	database *gorm.DB
}

func NewLogsRepositoryImp(db *gorm.DB) repository.IlogsRepository {
	return &LogsRepositoryImp{database: db}
}

func (r *LogsRepositoryImp) GetAllSessions(c *gin.Context) (*[]domain.SessionLog, map[string]interface{}, error) {
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

func (r *LogsRepositoryImp) GetAllSessionsByEmployeeId(c *gin.Context) (*[]domain.SessionLog, map[string]interface{}, error) {
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

func (r *LogsRepositoryImp) GetAllMovements(c *gin.Context) (*[]domain.ModerationLog, map[string]interface{}, error) {
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

func (r *LogsRepositoryImp) GetAllMovementsByEmployeeId(c *gin.Context) (*[]domain.ModerationLog, map[string]interface{}, error) {
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
