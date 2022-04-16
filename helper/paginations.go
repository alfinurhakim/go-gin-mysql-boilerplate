package helper

import (
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit  string `form:"limit" json:"limit"`
	Page   string `form:"page" json:"page"`
	Search string `form:"search" json:"search"`
}

func Paginate(limit string, page string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(limit)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
