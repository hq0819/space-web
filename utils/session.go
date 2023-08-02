package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"space-web/model"
	"time"
)

var UserLocal *session.Store

func InitSession() {
	UserLocal = session.New(session.Config{Expiration: time.Minute * 3})
	UserLocal.RegisterType(model.UserInfo{})
}

func GetUserInfo(ctx *fiber.Ctx) model.UserInfo {
	ses, err := UserLocal.Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return ses.Get("user").(model.UserInfo)
}
