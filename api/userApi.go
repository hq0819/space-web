package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"space-web/model"
	"space-web/service"
)

func InitUserApi(app *fiber.App) {
	group := app.Group("/user")
	group.Post("/login", Login)
	group.Post("/register", Register)
}

func Login(ctx *fiber.Ctx) error {
	user := new(model.User)
	err := ctx.BodyParser(user)
	if err != nil {
		fmt.Println(err)
	}
	res := service.GetUser(user)
	return ctx.JSON(res)
}

func Register(ctx *fiber.Ctx) error {
	m := new(model.User)
	_ = ctx.BodyParser(m)
	return ctx.JSON(service.RegisterUser(m))
}
