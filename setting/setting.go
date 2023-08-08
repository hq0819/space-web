package setting

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
import "github.com/spf13/viper"

var DB *gorm.DB
var Rdb *redis.Client
var FtpConf *FtpConfig

func InitDB(config *MysqlConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{QueryFields: true, NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		panic("数据库初始化失败")
	}
}

func InitRedis(config *redisConfig) {
	network := fmt.Sprintf(`%s:%s`, config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     network,
		Password: config.Auth,
	})
	Rdb = client
}
func InitFtp(config *FtpConfig) {
	FtpConf = config
}

func InitConfig() {
	vi := viper.New()
	vi.SetConfigFile("config.yml")
	mysqlConfig := new(MysqlConfig)
	redisConfig := new(redisConfig)
	ftpConf := new(FtpConfig)
	err := vi.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	_ = vi.UnmarshalKey("mysql", mysqlConfig, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
	})
	_ = vi.UnmarshalKey("redis", redisConfig, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
	})
	_ = vi.UnmarshalKey("ftp", ftpConf, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化数据库
	InitDB(mysqlConfig)
	InitRedis(redisConfig)
	InitFtp(ftpConf)
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type redisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Auth string `json:"auth"`
}

type FtpConfig struct {
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}
