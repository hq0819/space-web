package setting

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
import "github.com/spf13/viper"

var DB *gorm.DB

func InitDB(config *MysqlConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库初始化失败")
	}
}

func InitConfig() {
	vi := viper.New()
	vi.SetConfigFile("config.yml")
	mysqlConfig := new(MysqlConfig)
	err := vi.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	err = vi.UnmarshalKey("mysql", mysqlConfig, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
	})
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	//初始化数据库
	InitDB(mysqlConfig)

}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}
