package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"space-web/dao"
	"space-web/filter"
	"space-web/router"
	"space-web/setting"
	"space-web/utils"
)

func main() {
	fmt.Println("服务启动中...")
	app := fiber.New()
	app.Use(recover.New(recover.Config{EnableStackTrace: true, StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
		fmt.Printf("url:%s,参数:%s", c.OriginalURL(), c.Request().Body())
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
	//初始化路由
	router.InitRoute(app)
	err := app.Listen(":9001")
	if err != nil {
		_ = fmt.Errorf("服务启动失败%s", err.Error())
	}

}
