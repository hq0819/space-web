package filter

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"space-web/result"
	"space-web/utils"
)

func LoginFilter(ctx *fiber.Ctx) error {
	url := ctx.OriginalURL()
	fmt.Println(url)
	if url == `/user/login` || url == `/user/register` {
		return ctx.Next()
	}
	get, err := utils.UserLocal.Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	u := get.Get("user")
	if u != nil {
		return ctx.Next()
	}
	return ctx.JSON(result.Fail("请先登录"))
}
