package database

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/go-redis/redis"
	"os"
)

var RedisClient *redis.Client

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	redisHost := os.Getenv("APPLICATION_REDIS_HOST")
	if len(redisHost) == 0{
		redisHost = appConfig.DefaultString("redis.address","localhost:6379")
	}
	redisPassword := os.Getenv("APPLICATION_REDIS_PASSWORD")
	if len(redisPassword) == 0{
		redisPassword = appConfig.DefaultString("redis.address","")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword, // no password set
		DB:       0,  // use default DB
	})
	RedisClient = client
}
