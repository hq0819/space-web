package dao

import (
	"errors"
	"space-web/model"
	"space-web/setting"
)

func GetUser(u *model.User) (*model.User, error) {
	user := new(model.User)
	err := setting.DB.Where("username=? and  password = ?", u.Username, u.Password).First(user).Error
	if err == nil {
		user.Password = ""
		return user, nil
	}
	return nil, errors.New("用户名或密码错误")

}

func RegisterUser(u *model.User) error {
	return setting.DB.Create(u).Error
}
