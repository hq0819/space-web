package api

import (
	"github.com/gofiber/fiber/v2"
)

func InitArticleApi(app *fiber.App) {
	group := app.Group("/article")
	group.Get("/list", GetArticleList)
}

func GetArticleList(ctx *fiber.Ctx) error {
	return ctx.SendString(`成功`)

}
