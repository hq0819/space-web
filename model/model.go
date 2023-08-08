package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type User struct {
	RowID        int       `json:"rowId" gorm:"primaryKey;type:int(10);autoincrement;comment:用户id " `
	Username     string    `form:"username" json:"username" gorm:"type:varchar(15);comment:用户名"`
	Password     string    `form:"password" json:"password" gorm:"type:varchar(50);comment:密码"`
	Gender       string    `json:"gender" gorm:"type:varchar(5);comment:性别"`
	Avatar       string    `json:"avatar" gorm:"type varchar(200);comment:头像"`
	PositionType string    `json:"positionType" gorm:"type:varchar(20);comment:职位类型"`
	CreateTime   LocalDate `json:"createTime" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
}

type Article struct {
	RowId      int       `json:"rowId" gorm:"primaryKey;type:int(10);autoincrement;comment:文章id"`
	UserId     int       `json:"userId" gorm:"type:varchar(10);comment:作者id t_user(row_id)"`
	Title      string    `json:"title" gorm:"type:varchar(100);comment:标题"`
	Content    string    `json:"content" gorm:"type:text;comment:文章内容"`
	PicUrl     string    `json:"picUrl" gorm:"type varchar(500);comment:图片"`
	CreateTime LocalDate `json:"createTime" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime LocalDate `json:"UpdateTime" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:修改时间"`
	CoverImg   string    `json:"coverImg" gorm:"type:varchar(500);comment:封面图片"`
}

type Comment struct {
	RowId      int       `json:"rowId" gorm:"primaryKey;type:int(10);autoincrement;comment:主键"`
	ArticleId  string    `json:"articleId" gorm:"type:int(10);comment:文章id"`
	Content    string    `json:"content" gorm:"type:text;comment:内容"`
	CreateTime LocalDate `json:"createTime" gorm:"type:timestamp;comment:创建时间"`
}

func (*Comment) TableName() string {
	return "T_COMMENT"
}

func (*Article) TableName() string {
	return "T_ARTICLE"
}

func (*User) TableName() string {
	return "T_USER"
}

type LocalDate time.Time

func (l *LocalDate) MarshalJSON() ([]byte, error) {
	t := time.Time(*l)
	output := fmt.Sprintf("\"%s\"", t.Format("2006/01/02 15:04:05"))
	return []byte(output), nil
}

func (l LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	tt := time.Time(l)
	if tt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tt, nil

}
func (l *LocalDate) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*l = LocalDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type UserInfo struct {
	RowId        int    `json:"rowId"`
	Username     string `json:"username"`
	Gender       string `json:"gender"`
	Avatar       string `json:"avatar"`
	PositionType string `json:"positionType"`
	CreateTime   string `json:"createTime"`
}
