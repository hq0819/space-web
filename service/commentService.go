package service

import (
	"space-web/model"
	"space-web/result"
)

func QueryCommentList(page *model.PageInfo) *result.Message {
	return result.Success(model.StartPage[model.Comment](page))
}
