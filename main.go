package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"space-web/api"
	"space-web/model"
	"time"
)

func main() {
	fmt.Println("服务启动中...")
	app := fiber.New()
	store := session.New(session.Config{Expiration: 10 * time.Minute})
	store.RegisterType(model.User{})

	app.Use(store)
	//初始化配置
	//setting.InitConfig()
	//数据库自动迁移
	//dao.MigrateModels()
	api.InitUserApi(app)
	api.InitArticleApi(app)

	err := app.Listen(":9001")
	if err != nil {
		_ = fmt.Errorf("服务启动失败%s", err.Error())
	}

}
