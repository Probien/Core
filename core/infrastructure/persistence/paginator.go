package persistence

import (
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(params url.Values, paginate map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		requestCurrentPage, _ := strconv.Atoi(params.Get("page"))
		pages := paginate["total_pages"].(float64)

		//option by defaul: Start with the page number 1
		paginate["previous"] = "1"
		paginate["page"] = 1

		//case 1: if there is not present the query param we set as default the page 1 as long as there are more than one page
		//It returns:
		//previous: 1
		//page: 1
		//next: 2
		if len(params.Get("page")) == 0 && pages > 1 {
			paginate["next"] = "2"
			return db.Offset(0).Limit(10)
		}

		//case 2: if there no more pages, we set same page as the next as long as there is the query param
		//It returns:
		//previous: 1
		//page: 1
		//next: 1
		if pages <= 1 {
			paginate["next"] = "1"
			return db.Offset(0).Limit(10)

		}

		//case 3: if current page is greater than the total pages, we redirect to start
		if float64(requestCurrentPage) > pages {
			paginate["previous"] = "1"
			paginate["next"] = "1"
			return db.Offset(0).Limit(10)
		}

		paginate["page"] = requestCurrentPage
		paginate["previous"] = strconv.Itoa(paginate["page"].(int) - 1)
		//case 4: if current page is the total page limit, we get current page as the next page, otherwise set other page
		if requestCurrentPage == int(pages) {
			paginate["next"] = strconv.Itoa(paginate["page"].(int))
		} else {
			paginate["next"] = strconv.Itoa(paginate["page"].(int) + 1)
		}

		return db.Offset(paginate["page"].(int) * 10).Limit(10)
	}
}
