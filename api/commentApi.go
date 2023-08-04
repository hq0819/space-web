package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"space-web/model"
	"space-web/service"
)

func InitCommentApi(app *fiber.App) {
	group := app.Group("/comment")
	group.Post("/list", CommentList)

}

func CommentList(ctx *fiber.Ctx) error {
	pageInfo := new(model.PageInfo)
	err := ctx.BodyParser(pageInfo)
	if err != nil {
		fmt.Println(err)
	}
	service.QueryCommentList(pageInfo)
	return ctx.JSON(nil)
}
