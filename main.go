package main

import (
	"encoding/json"
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
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(filter.LoginFilter)
	logger := log.DefaultLogger()
	logger.SetLevel(log.LevelDebug)
	log.SetLogger(logger)
	//初始化配置
	setting.InitConfig()
	//初始化session
	utils.InitSession()
	//数据库自动迁移
	dao.MigrateModels()
	//初始化路由
	router.InitRoute(app)
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Print(string(data))
	err := app.Listen(":9001")
	if err != nil {
		_ = fmt.Errorf("服务启动失败%s", err.Error())
	}

}
