package redis

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var RedisConn *redis.Client

func init() {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ping := RedisConn.Ping()
	beego.Info(ping.Result())
}
