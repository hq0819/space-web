package main

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"space-web/api"
	"space-web/filter"
	"space-web/model"
	"space-web/setting"
)

func main() {
	fmt.Println("服务启动中...")
	engine := gin.Default()
	//初始化配置
	setting.InitConfig()
	//数据库自动迁移
	//dao.MigrateModels()
	gob.Register(new(model.User))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 3600 * 2})
	sessFunc := sessions.Sessions("mysession", store)
	//登录校验
	engine.Use(sessFunc, filter.LoginFilter)
	api.InitUserApi(engine)
	api.InitArticleApi(engine)
	err := engine.Run(":9001")
	if err != nil {
		_ = fmt.Errorf("服务启动失败%s", err.Error())
	}

}
