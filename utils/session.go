package utils

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"space-web/model"
	"space-web/setting"
	"time"
)

var UserLocal *session.Store

const prefix = "user:"

func InitSession() {
	UserLocal = session.New(session.Config{KeyLookup: "cookie:access-token", Expiration: time.Minute * 30, Storage: RedisStorage{}})
	UserLocal.RegisterType(model.UserInfo{})
	UserLocal.RegisterType(model.LocalDate{})
}

func GetUserInfo(ctx *fiber.Ctx) model.UserInfo {
	ses, err := UserLocal.Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return ses.Get("user").(model.UserInfo)
}

type RedisStorage struct {
}

func (r RedisStorage) Get(key string) ([]byte, error) {
	return setting.Rdb.Get(context.Background(), prefix+key).Bytes()
}

func (r RedisStorage) Set(key string, val []byte, exp time.Duration) error {
	return setting.Rdb.Set(context.Background(), prefix+key, val, exp).Err()
}

func (r RedisStorage) Delete(key string) error {
	return setting.Rdb.Del(context.Background(), prefix+key).Err()
}

func (r RedisStorage) Reset() error {
	return setting.Rdb.Del(context.Background(), prefix+"*").Err()
}

func (r RedisStorage) Close() error {
	return setting.Rdb.Close()
}
