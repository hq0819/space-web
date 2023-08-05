package model

import (
	"gorm.io/gorm"
	"space-web/setting"
)

type PageInfo struct {
	PageSize  int64          `json:"pageSize"`
	PageNum   int64          `json:"pageNum"`
	OrderBy   string         `json:"orderBy"`
	Total     int64          `json:"total"`
	Condition map[string]any `json:"condition"`
}

type Page[T any] struct {
	PageSize int64 `json:"pageSize"`
	PageNum  int64 `json:"pageNum"`
	Rows     []*T  `json:"rows"`
	Total    int64 `json:"total"`
	LastPage bool  `json:"lastPage"`
}

func ToPage[T any](pageInfo *PageInfo, db *gorm.DB) *Page[T] {
	offset := (pageInfo.PageNum - 1) * pageInfo.PageSize
	var list []*T
	var total int64
	page := new(Page[T])
	page.PageNum = pageInfo.PageNum
	page.PageSize = pageInfo.PageSize
	page.Rows = make([]*T, 0)
	err := db.Count(&total).Error
	page.Total = total
	if err != nil || total == 0 {
		page.LastPage = true
		return page
	}
	db.Limit(int(pageInfo.PageSize)).Offset(int(offset)).Scan(&list)
	if len(list) == 0 {
		page.LastPage = true
		return page
	}
	page.Rows = list
	if page.PageSize*page.PageNum >= total {
		page.LastPage = true
	}
	return page
}

func StartPage[T any](pageInfo *PageInfo) *Page[T] {
	var list []*T
	var total int64
	var offset = int((pageInfo.PageNum - 1) * pageInfo.PageSize)
	page := new(Page[T])
	page.PageNum = pageInfo.PageNum
	page.PageSize = pageInfo.PageSize
	page.Rows = make([]*T, 0)
	err := setting.DB.Debug().Model(new(T)).Where(pageInfo.Condition).Count(&total).Error
	page.Total = total
	if err != nil || total == 0 {
		page.LastPage = true
		return page
	}
	setting.DB.Debug().Model(new(T)).Where(pageInfo.Condition).Order(pageInfo.OrderBy).Limit(int(pageInfo.PageSize)).Offset(offset).Scan(&list)
	if len(list) == 0 {
		page.LastPage = true
		return page
	}
	page.Rows = list
	if page.PageSize*page.PageNum >= total {
		page.LastPage = true
	}
	return page
}
