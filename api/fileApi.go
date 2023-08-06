package api

import "github.com/gofiber/fiber/v2"

func InitFileApi(app *fiber.App) {
	group := app.Group("/file")
	group.Post("/upload", Upload)
}

func Upload(ctx *fiber.Ctx) error {

	return nil

}
