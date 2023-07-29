package filter

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"space-web/model"
	"space-web/result"
)

func LoginFilter(ctx *gin.Context) {
	path := ctx.FullPath()
	if path == `/user/login` || path == `/user/register` {
		ctx.Next()
		return
	}
	session := sessions.Default(ctx)
	get := session.Get("user")
	fmt.Println(get)
	user, ok := get.(*model.User)
	if ok && user != nil {
		ctx.Next()
		return
	}
	ctx.JSON(200, result.Fail("请先登录"))
	return
}
