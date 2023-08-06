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
	group.Post("/list", GetArticleList)
	group.Post("/addArticle", AddArticle)
	group.Post("/queryRecommendArticles", QueryRecommendArticle)
	group.Post("/queryScanCount/:id", QueryScanCount)
	group.Post("/querySupportCount/:id", QuerySupportCount)
	group.Post("/queryCommentCount/:id", QueryCommentCount)
	group.Post("/addSupport/:id", AddSupport)
	group.Post("/addScanRecord/:id", AddScanRecord)

}

func GetArticleList(ctx *fiber.Ctx) error {
	info := utils.GetUserInfo(ctx)
	list, err := service.GetArticleListByUserId(info.RowId)
	if err != nil {
		return ctx.JSON(result.Fail("查询失败"))
	}
	return ctx.JSON(result.Success(list))
}

func QueryRecommendArticle(ctx *fiber.Ctx) error {
	pageInfo := new(model.PageInfo)
	_ = ctx.BodyParser(pageInfo)
	article := service.ArticleList(pageInfo)
	return ctx.JSON(article)
}
func QueryScanCount(ctx *fiber.Ctx) error {
	artID, _ := ctx.ParamsInt("id")
	return ctx.JSON(service.QueryScanCount(artID))
}

func QuerySupportCount(ctx *fiber.Ctx) error {
	artID, _ := ctx.ParamsInt("id")
	return ctx.JSON(service.QuerySupportCount(artID))
}
func QueryCommentCount(ctx *fiber.Ctx) error {
	artID, _ := ctx.ParamsInt("id")
	return ctx.JSON(service.QueryCommentCount(artID))
}

func AddSupport(ctx *fiber.Ctx) error {
	user, _ := utils.UserLocal.Get(ctx)
	info := user.Get("user").(model.UserInfo)
	artID, _ := ctx.ParamsInt("id")
	service.AddArticleSupport(int64(artID), int64(info.RowId))
	return ctx.JSON(result.Success(nil))
}
func AddScanRecord(ctx *fiber.Ctx) error {
	artID, _ := ctx.ParamsInt("id")
	service.AddArticleScan(artID)
	return ctx.JSON(result.Success(nil))
}

func AddArticle(ctx *fiber.Ctx) error {
	info := utils.GetUserInfo(ctx)
	m := new(model.Article)
	_ = ctx.BodyParser(m)
	m.UserId = info.RowId
	m.CreateTime = model.LocalDate(time.Now())
	m.UpdateTime = model.LocalDate(time.Now())
	err := service.AddArticle(m)
	if err != nil {
		return ctx.JSON(result.Fail("创建文章失败"))
	}
	return ctx.JSON(result.Success("成功"))
}
