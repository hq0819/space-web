package api

import (
	"github.com/gofiber/fiber/v2"
	"space-web/model"
	"space-web/result"
	"space-web/service"
	"space-web/utils"
	"time"
)

func InitArticleApi(app *fiber.App) {
	group := app.Group("/article")
	group.Get("/list", GetArticleList)
	group.Post("/addArticle", AddArticle)
}

func GetArticleList(ctx *fiber.Ctx) error {
	info := utils.GetUserInfo(ctx)
	list, err := service.GetArticleListByUserId(info.RowId)
	if err != nil {
		return ctx.JSON(result.Fail("查询失败"))
	}
	return ctx.JSON(result.Success(list))
}

func AddArticle(ctx *fiber.Ctx) error {
	info := utils.GetUserInfo(ctx)
	m := new(model.Article)
	_ = ctx.BodyParser(m)
	m.UserId = info.RowId
	m.CreateTime = time.Now()
	err := service.AddArticle(m)
	if err != nil {
		return ctx.JSON(result.Fail("创建文章失败"))
	}
	return ctx.JSON(result.Success("成功"))
}
