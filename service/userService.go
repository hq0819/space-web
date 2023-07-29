package service

import (
	"space-web/dao"
	"space-web/model"
	"space-web/result"
)

func GetUser(u *model.User) *result.Message {
	user, err := dao.GetUser(u)
	if err != nil {
		return result.Fail("用户名或密码错误!!!")
	}
	return result.Success(user)
}

func RegisterUser(u *model.User) *result.Message {
	err := dao.RegisterUser(u)
	if err != nil {
		return result.Fail("注册失败")
	}
	return result.Success(nil)
}
