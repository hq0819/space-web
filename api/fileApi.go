package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"space-web/result"
	"space-web/utils"
	"strings"
)

const static_resource_prefix = "http://114.115.213.117/selfspace/statics/"

func InitFileApi(app *fiber.App) {
	group := app.Group("/file")
	group.Post("/upload", Upload)
}

func Upload(ctx *fiber.Ctx) error {
	str := uuid.New().String()
	file, _ := ctx.FormFile("file")
	split := strings.Split(file.Filename, ".")
	fname := fmt.Sprintf("%s.%s", str, split[len(split)-1])
	open, _ := file.Open()
	err := utils.Upload(fname, open)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(result.Fail("上传失败"))
	}

	return ctx.JSON(result.Success(static_resource_prefix + fname))
}
