package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"space-web/api"
	"space-web/dao"
	"space-web/filter"
	"space-web/setting"
	"space-web/utils"
)

func main() {
	fmt.Println("服务启动中...")
	app := fiber.New()
	app.Use(recover.New(recover.Config{StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
		log.Errorf("url%s:", c.OriginalURL())
	}}))
	app.Use(filter.LoginFilter)
	logger := log.DefaultLogger()
	logger.SetLevel(log.LevelDebug)
	log.SetLogger(logger)
	utils.InitSession()
	//初始化配置
	setting.InitConfig()
	//数据库自动迁移
	dao.MigrateModels()
	api.InitUserApi(app)
	api.InitArticleApi(app)

	err := app.Listen(":9001")
	if err != nil {
		_ = fmt.Errorf("服务启动失败%s", err.Error())
	}

}
