package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"space-web/model"
	"space-web/service"
	"space-web/utils"
)

func InitUserApi(app *fiber.App) {
	group := app.Group("/user")
	group.Post("/login", Login)
	group.Post("/register", Register)
}

func Login(ctx *fiber.Ctx) error {
	session, err2 := utils.UserLocal.Get(ctx)
	if err2 != nil {
		fmt.Println(err2)
	}
	user := new(model.User)
	err := ctx.BodyParser(user)
	if err != nil {
		fmt.Println(err)
	}
	res := service.GetUser(user)
	if err != nil {
		fmt.Println(err)
	}
	session.Set("user", res.Data)
	err = session.Save()
	if err != nil {
		fmt.Println(err)
	}
	return ctx.JSON(res)
}

func Register(ctx *fiber.Ctx) error {
	m := new(model.User)
	_ = ctx.BodyParser(m)
	return ctx.JSON(service.RegisterUser(m))
}
