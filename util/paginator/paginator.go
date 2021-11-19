package paginator

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetPageParams(ctx *gin.Context) (int, int) {
	var page, pageSize int
	pageStr, ok := ctx.GetQuery("page")
	if !ok {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
		if page == 0 {
			page = 1
		}
	}

	pageSizeStr, ok := ctx.GetQuery("page_size")
	if !ok {
		pageSize = 10
	} else {
		pageSize, _ = strconv.Atoi(pageSizeStr)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
	}
	return page, pageSize
}
