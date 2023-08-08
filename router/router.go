package router

import (
	"github.com/gofiber/fiber/v2"
	"space-web/api"
)

func InitRoute(app *fiber.App) {
	//用户api
	api.InitUserApi(app)
	//文章api
	api.InitArticleApi(app)
	//评论api
	api.InitCommentApi(app)

	api.InitFileApi(app)
}
