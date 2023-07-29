package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"space-web/model"
	"space-web/service"
	"space-web/utils"
	"time"
)

func InitUserApi(gin *gin.Engine) {
	group := gin.Group("/user")
	group.GET("/login", Login)
	group.POST("/register", Register)
	group.GET("/getSystemTime", GetSystemTime)
}

func Login(ctx *gin.Context) {
	user := new(model.User)
	_ = ctx.BindQuery(user)
	res := service.GetUser(user)
	session := sessions.Default(ctx)
	session.Set("user", utils.ModelToMap(res.Data, `2006/01/02 15:04:05`, "password"))
	_ = session.Save()
	ctx.JSON(200, res)
	return
}

func Register(ctx *gin.Context) {
	m := new(model.User)
	_ = ctx.ShouldBind(m)
	ctx.JSON(200, service.RegisterUser(m))
	return
}

func GetSystemTime(ctx *gin.Context) {
	ctx.JSON(200, time.Now())
	return
}
