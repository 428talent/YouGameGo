package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"you_game_go/redis"
)

type Token struct {
	UserId    int
	LoginTime time.Time
}

func InsertTokenToRedis(data *Token, key string) error {
	var insertMap = map[string]interface{}{
		"UserId":    data.UserId,
		"LoginTime": data.LoginTime,
	}
	beego.Info(data)
	beego.Info(key)
	redis.RedisConn.HMSet(fmt.Sprintf("token:%s", key), insertMap)
	redis.RedisConn.Expire(fmt.Sprintf("token:%s", key), time.Duration(time.Hour * 144))
	return nil
}
