package service

import (
	"space-web/dao"
	"space-web/model"
	"space-web/result"
)

func GetUser(u *model.User) (*model.UserInfo, error) {
	return dao.GetUser(u)
}

func RegisterUser(u *model.User) *result.Message {
	err := dao.RegisterUser(u)
	if err != nil {
		return result.Fail("注册失败")
	}
	return result.Success(nil)
}
