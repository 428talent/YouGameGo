package database

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     appConfig.String("redis.address"),
		Password: appConfig.String("redis.password"), // no password set
		DB:       0,  // use default DB
	})
	RedisClient = client
}
