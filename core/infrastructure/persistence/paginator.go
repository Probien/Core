package persistence

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context, paginate map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		pages := paginate["total_pages"].(float64)

		paginate["previous"] = "1"
		paginate["page"] = 1

		if pages <= 1 {
			paginate["next"] = "1"
			return db.Offset(0).Limit(10)

		} else if len(c.Query("page")) == 0 && pages > 1 {
			paginate["next"] = "2"
			return db.Offset(0).Limit(10)
		}

		paginate["page"], _ = strconv.Atoi(c.Query("page"))
		paginate["previous"] = strconv.Itoa(paginate["page"].(int) - 1)
		paginate["next"] = strconv.Itoa(paginate["page"].(int) + 1)

		return db.Offset(paginate["page"].(int) * 10).Limit(10)
	}
}
