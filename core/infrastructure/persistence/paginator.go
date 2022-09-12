package persistence

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func Paginate(c *gin.Context, paginate map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if len(c.Query("page")) == 0 {
			paginate["previous"] = "1"
			paginate["next"] = "2"
			paginate["page"] = 1

		} else {
			paginate["page"], _ = strconv.Atoi(c.Query("page"))
			paginate["previous"] = strconv.Itoa(paginate["page"].(int) - 1)
			paginate["next"] = strconv.Itoa(paginate["page"].(int) + 1)
		}

		return db.Offset(paginate["page"].(int) * 10).Limit(10)
	}
}
