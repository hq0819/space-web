package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"space-web/model"
)

func InitArticleApi(gin *gin.Engine) {
	group := gin.Group("/article")
	group.GET("/list", GetArticleList)
}

func GetArticleList(ctx *gin.Context) {
	u, ok := sessions.Default(ctx).Get("user").(*model.User)
	if ok && u != nil {
		fmt.Println(u)
	}
}
