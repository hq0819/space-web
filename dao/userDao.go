package dao

import (
	"errors"
	"space-web/model"
	. "space-web/setting"
)

func GetUser(u *model.User) (*model.UserInfo, error) {
	user := new(model.UserInfo)
	err := DB.Debug().Model(model.User{}).Select("row_id", "username", "gender", "avatar",
		"position_type", "DATE_FORMAT(create_time, '%Y/%m/%d %H:%i:%s') as create_time").
		Where("username=? and password=?", u.Username, u.Password).First(user).Error
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	return user, nil

}

func RegisterUser(u *model.User) error {
	return DB.Create(u).Error
}
