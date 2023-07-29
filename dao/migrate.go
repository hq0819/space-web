package dao

import (
	"space-web/model"
	"space-web/setting"
)

func MigrateModels() {
	err := setting.DB.Migrator().AutoMigrate(new(model.User), new(model.Article))
	if err != nil {
		panic(err)
	}
}
