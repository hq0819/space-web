package model

import "time"

type User struct {
	RowID        string    `json:"rowId" gorm:"primaryKey;type:int(10);autoincrement;comment:用户id " `
	Username     string    `form:"username" json:"username" gorm:"type:varchar(15);comment:用户名"`
	Password     string    `form:"password" json:"password" gorm:"type:varchar(50);comment:密码"`
	Gender       string    `json:"gender" gorm:"type:varchar(5);comment:性别"`
	PositionType string    `json:"positionType" gorm:"type:varchar(20);comment:职位类型"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
}

type Article struct {
	RowId      string    `json:"rowId" gorm:"primaryKey;type:int(10);autoincrement;comment:文章id"`
	UserId     string    `json:"userId" gorm:"type:varchar(10);comment:作者id t_user(row_id)"`
	Title      string    `json:"title" gorm:"type:varchar(100);comment:标题"`
	Content    string    `json:"content" gorm:"type:text;comment:文章内容"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:修改时间"`
}

func (*Article) TableName() string {
	return "T_ARTICLE"
}

func (*User) TableName() string {
	return "T_USER"
}
