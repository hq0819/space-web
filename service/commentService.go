package service

import (
	"space-web/model"
	"space-web/setting"
)

func QueryCommentList(page *model.PageInfo) *model.PageInfo {
	var list *[]model.Comment
	setting.DB.Debug().Model(model.Comment{}).Where(page.Condition).Scan(list)
	return nil
}
